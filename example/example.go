package example

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gitlab.com/osaki-lab/reguerr"
	"net/http"
)

func UserHandleFunc(w http.ResponseWriter, r *http.Request) {
	if err := somethingUserOperation("id1"); err != nil {
		errorLog(err)
		w.WriteHeader(httpStatus(err))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func errorLog(err error) {
	rerr, ok := reguerr.ErrorOf(err)
	if ok {
		level, err := zerolog.ParseLevel(rerr.Level().String())
		if err != nil {
			log.Error().Str("code", rerr.Code()).Msgf(rerr.Error())
		} else {
			log.WithLevel(level).Str("code", rerr.Code()).Msgf(rerr.Error())
		}
		return
	}
	log.Printf("unexpected error: %v\n", err)
}

func httpStatus(err error) int {
	code, ok := reguerr.StatusOf(err)
	if ok {
		return code
	} else {
		return http.StatusInternalServerError
	}
}

func somethingUserOperation(key string) error {
	// some operation for user
	return NewNotFoundOperationIDErr(fmt.Errorf("key=%v", key))
}

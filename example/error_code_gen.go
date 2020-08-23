// generated by reguerr; DO NOT EDIT
package example

import (
	"errors"
	"gitlab.com/osaki-lab/reguerr"
)

func NewPermissionDeniedErr(err error) *reguerr.CodeError {
	return PermissionDeniedErr.WithError(err)
}

func IsPermissionDeniedErr(err error) bool {
	var cerr *reguerr.CodeError
	if as := errors.As(err, &cerr); as {
		if cerr.Code == PermissionDeniedErr.Code {
			return true
		}
	}
	return false
}

func NewUpdateConflictErr(err error, arg1 interface{}) *reguerr.CodeError {
	return UpdateConflictErr.WithError(err).WithArgs(arg1)
}

func IsUpdateConflictErr(err error) bool {
	var cerr *reguerr.CodeError
	if as := errors.As(err, &cerr); as {
		if cerr.Code == UpdateConflictErr.Code {
			return true
		}
	}
	return false
}

func NewInvalidInputParameterErr(err error, payload map[string]interface{}) *reguerr.CodeError {
	return InvalidInputParameterErr.WithError(err).WithArgs(payload)
}

func IsInvalidInputParameterErr(err error) bool {
	var cerr *reguerr.CodeError
	if as := errors.As(err, &cerr); as {
		if cerr.Code == InvalidInputParameterErr.Code {
			return true
		}
	}
	return false
}

func NewUserTypeUnregisterErr() *reguerr.CodeError {
	return UserTypeUnregisterErr
}

func IsUserTypeUnregisterErr(err error) bool {
	var cerr *reguerr.CodeError
	if as := errors.As(err, &cerr); as {
		if cerr.Code == UserTypeUnregisterErr.Code {
			return true
		}
	}
	return false
}

func NewNotFoundOperationIDErr(err error) *reguerr.CodeError {
	return NotFoundOperationIDErr.WithError(err)
}

func IsNotFoundOperationIDErr(err error) bool {
	var cerr *reguerr.CodeError
	if as := errors.As(err, &cerr); as {
		if cerr.Code == NotFoundOperationIDErr.Code {
			return true
		}
	}
	return false
}

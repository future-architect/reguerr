package example

import (
	"gitlab.com/osaki-lab/reguerr"
)

// reguerr targets is below.
// $GOPATH/gitlab.com/osaki-lab/reguerr>go run cmd/reguerr/main.go -f example/error_code.go
var (
	// No message arguments
	PermissionDeniedErr = reguerr.New("1001", "permission denied")

	// One message arguments
	UpdateConflictErr = reguerr.New("1002", "other user updated: key=%s")

	// Message arguments with label
	InvalidInputParameterErr = reguerr.New("1003", "invalid input parameter: %v").
		Label(0,"payload", map[string]interface{}{})

	// Disable default error argument
	UserTypeUnregisterErr = reguerr.New("1005", "not found user type").DisableError()

	// Change log level and exitCode
	NotFoundOperationIDErr = reguerr.New("1004", "not found operation id").
		WarnLevel().WithStatusCode(404)
)

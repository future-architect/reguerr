package example

import (
	"gitlab.com/future-architect/reguerr"
)

// reguerr targets is below.
// cd $GOPATH/src/gitlab.com/future-architect/reguerr && go run cmd/reguerr/main.go generate -f example/error_code.go
var (
	// No message arguments
	PermissionDeniedErr = reguerr.New("1001", "permission denied").Build()

	// One message arguments
	UpdateConflictErr = reguerr.New("1002", "other user updated: key=%s").Build()

	// Message arguments with label
	InvalidInputParameterErr = reguerr.New("1003", "invalid input parameter: %v").
		Label(0,"payload", map[string]interface{}{}).
		Build()

	// Disable default error argument
	UserTypeUnregisterErr = reguerr.New("1005", "not found user type").DisableError().Build()

	// Change log level and exitCode
	NotFoundOperationIDErr = reguerr.New("1004", "not found operation id").
		Warn().WithStatusCode(404).
		Build()
)

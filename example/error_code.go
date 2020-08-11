package example

import (
	"gitlab.com/osaki-lab/errcdgen"
)

// errcdgen targets is below.
var (
	// No message arguments
	PermissionDeniedErr = errcdgen.NewCodeError("1001", "permission denied")

	// One message arguments
	UpdateConflictErr = errcdgen.NewCodeError("1002", "other user updated: key=%s")

	// Message arguments with label
	InvalidInputParameterErr = errcdgen.NewCodeError("1003", "invalid input parameter: %v").
		Label("payload", map[string]interface{}{})

	// Disable default error argument
	UserTypeUnregisterErr = errcdgen.NewCodeError("1005", "not found user type").DisableError()

	// Change log level and exitCode
	NotFoundOperationIDErr = errcdgen.NewCodeError("1004", "not found operation id").
		WarnLevel().WithStatusCode(404)
)

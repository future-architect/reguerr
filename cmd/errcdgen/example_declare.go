package main

import (
	"gitlab.com/osaki-lab/errcdgen"
)

var (
	InvalidInputParameterErr = errcdgen.NewCodeError("1003", "invalid input parameter: %v").Arg("payload", map[string]interface{}{}) // ラベル付き
	UpdateConflictErr        = errcdgen.NewCodeError("1003", "other user updated: key=%s")                                           // 引数1つ
	NotFoundOperationIDErr   = errcdgen.NewCodeError("1004", "not found operation id").WithLevel(errcdgen.WarnLevel)                 // ログレベルを変更
	UserTypeUnregisterErr    = errcdgen.NewCodeError("1005", "not found user type").DisableError()                                   // 引数なし
)

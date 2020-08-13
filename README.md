# errcdgen

errcdgen - error code generator

## Motivation

業務用システムにおいて、何か不具合が発生した場合に問題切り分けをスムーズにするため、エラーコードを付与することが求められます。

一方でGoにはerrorという仕組みがあり、コレを用いて異常な状態を上位に伝播させる仕組みを持ちますが、これと上記のエラーコードをつなげる良い仕組みを、見つけることができませんでした。

このツールを用いると、一貫的なエラーコード付きの error を生成することができます。

また、ログレベルやレスポンスコードも個別設定できるため、エラーハンドリングの箇所（ログ出力やメトリクス通知）に活かす事もできます。

## Installation

```sh
go get -u gitlab.com/osaki-lab/errcdgen/cmd/errcdgen
```

## Options

```sh
>errcdgen -h

Usage:
  code generator for error handling with message code [flags]

Flags:
  -f, --file string   input go file
  -h, --help          help for code
```

## Usage

```sh
# target file
cat <<EOF > example.go
package example

import (
	"gitlab.com/osaki-lab/errcdgen"
)

var (
	// No message arguments
	PermissionDeniedErr = errcdgen.NewCodeError("1001", "permission denied")

	// One message arguments
	UpdateConflictErr = errcdgen.NewCodeError("1002", "other user updated: key=%s")

	// Message arguments with label
	InvalidInputParameterErr = errcdgen.NewCodeError("1003", "invalid input parameter: %v").
		Label(0,"payload", map[string]interface{}{})
)
EOF

# START errcdgen
./errcdgen example.go
```

Output is bellow format.

```go example_gen.go
// Code generated by errcdgen; DO NOT EDIT.
package example

import (
	"gitlab.com/osaki-lab/errcdgen"
)

func NewPermissionDeniedErr(err error) *errcdgen.CodeError {
	return PermissionDeniedErr.WithError(err)
}

func NewUpdateConflictErr(err error, arg1 interface{}) *errcdgen.CodeError {
	return UpdateConflictErr.WithError(err).WithArgs(arg1)
}

func NewInvalidInputParameterErr(err error, payload map[string]interface{}) *errcdgen.CodeError {
	return InvalidInputParameterErr.WithError(err).WithArgs(payload)
}
```

And errcdgen generated markdown table.

+------+--------------------------+----------+------------+-----------------------------+
| CODE |           NAME           | LOGLEVEL | STATUSCODE |           FORMAT            |
+------+--------------------------+----------+------------+-----------------------------+
| 1001 | PermissionDeniedErr      | Error    |        500 | permission denied           |
| 1002 | UpdateConflictErr        | Error    |        500 | other user updated: key=%s  |
| 1003 | InvalidInputParameterErr | Error    |        500 | invalid input parameter: %v |
+------+--------------------------+----------+------------+-----------------------------+


If you want to see more examples, you can get [example](./example).


## Use Case errcdgen 

TODO

## License

Apache License Version 2.0
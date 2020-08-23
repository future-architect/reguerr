### reguerr

reguerr - Code generator for systematic error handling

## Motivation

業務用システムにおいて、何か不具合が発生した場合に問題切り分けをスムーズにするため、エラーコードを付与することが求められます。

一方でGoにはerrorという仕組みがあり、コレを用いて異常な状態を上位に伝播させる仕組みを持ちますが、これと上記のエラーコードをつなげる良い仕組みを、見つけることができませんでした。

このツールを用いると、一貫的なエラーコード付きの error を生成することができます。

また、ログレベルやレスポンスコードも個別設定できるため、エラーハンドリングの箇所（ログ出力やメトリクス通知）に活かす事もできます。

## Installation

```sh
go get -u gitlab.com/osaki-lab/reguerr/cmd/reguerr
```

## Options

```sh
>reguerr -h

Usage:
  code generator for systematic error handling. [flags]

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
	"gitlab.com/osaki-lab/reguerr"
)

var (
	// No message arguments
	PermissionDeniedErr = reguerr.NewCodeError("1001", "permission denied")

	// One message arguments
	UpdateConflictErr = reguerr.NewCodeError("1002", "other user updated: key=%s")

	// Message arguments with label
	InvalidInputParameterErr = reguerr.NewCodeError("1003", "invalid input parameter: %v").
		Label(0,"payload", map[string]interface{}{})
)
EOF

# START reguerr
./reguerr example.go
```

Output is bellow format.

```go example_gen.go
// Code generated by reguerr; DO NOT EDIT.
package example

import (
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
```

Then reguerr also generated markdown table.

| CODE |           NAME           | LOGLEVEL | STATUSCODE |           FORMAT            |
|------|--------------------------|----------|------------|-----------------------------|
| 1001 | PermissionDeniedErr      | Error    |        500 | permission denied           |
| 1002 | UpdateConflictErr        | Error    |        500 | other user updated: key=%s  |
| 1003 | InvalidInputParameterErr | Error    |        500 | invalid input parameter: %v |


If you want to see more examples, you can get [example](./example).


## Use Case reguerr 

TODO

## License

Apache License Version 2.0
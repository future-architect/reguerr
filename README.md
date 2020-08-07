# errcdgen

errcdgen - error code generator

## Motivation

業務用システムにおいて、何か不具合が発生した場合に問題切り分けをスムーズにするため、エラーコードを付与することが求められます。

一方でGoにはerrorという仕組みがあり、コレを用いて異常な状態を上位に伝播させる仕組みを持ちますが、これと上記のエラーコードをつなげる良い仕組みを、見つけることがdけいませんでした。

このツールを用いると、一貫的なエラーコード付きの error を生成することができます。

## Installation

```sh
go get -u gitlab.com/osaki-lab/errcdgen
```

## Usage

```sh
# データ準備
cat <<EOF > error_list.jsonl
{"code": "10001", "type":"InvalidInputParameter", "format":"invalid input parameter: %v"}
{"code": "10002", "type":"UpdateConflict",        "format":"other user updated: [key:%s]"}
{"code": "10003", "type":"UserTypeUnregister",    "format":"not found user type: [%v]"}
EOF

# 生成
./errcdgen error_list.jsonl
```

Output is bellow format.

```go
const (
	// error codes
	InvalidInputParameter    = 10001
	UpdateConflict           = 10002
	UserTypeUnregister       = 10003

    InvalidInputParameterErr = NewCodeError(InvalidInputParameter, "invalid input parameter: %v")
    UpdateConflictErr        = NewCodeError(UpdateConflict, "other user updated: [key:%s]")
    UserTypeUnregisterErr    = NewCodeError(UserTypeUnregister, "not found user type: [%v]")
)

type CodeError struct {
	code    string
    format  string
    args    []interface
	err     error
}

func (e CodeError) Error() string {
    if e.err != nil {
    	return fmt.Sprintf("[%s] %s: %v", e.code, fmt.Sprintf(e.format, e.args), e.err)
    }
	return fmt.Sprintf("[%s] %s", e.code, fmt.Sprintf(e.format, e.args))
}

func NewCodeError(code, format string) CodeError {
    return &CodeError{
        code:   code,
        format: format,
    }
}

func (e CodeError)WithErr(err error) CodeError {
    return &CodeError{
        code: e.code,
        format: e.format,
        err: err,
    }
}
```


## Development with errcdgen


```go

```


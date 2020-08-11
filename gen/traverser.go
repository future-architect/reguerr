package gen

import (
	"gitlab.com/osaki-lab/errcdgen"
	"go/ast"
	"strconv"
	"strings"
)

type File struct {
	PkgName string
	Decls   []*Decl
}

type Decl struct {
	Name             string
	Code             string
	Format           string
	LogLevelEnable   bool
	LogLevel         errcdgen.Level
	StatusCodeEnable bool
	StatusCode       int
	DisableErr       bool
	Labels           []Label
}

type Label struct {
	Index  int
	Name   string
	GoType string
}

func (f File) Bindings() []Binding {
	var resp []Binding
	for _, d := range f.Decls {
		resp = append(resp, Binding{
			Name:             d.Name,
			DisableErr:       d.DisableErr,
			LogLevelEnable:   d.LogLevelEnable,
			LogLevel:         d.LogLevel,
			StatusCodeEnable: d.StatusCodeEnable,
			StatusCode:       d.StatusCode,
		})
	}
	return resp
}

func Traverse(n *ast.File) (*File, error) {

	var resp []*Decl
	for _, d := range n.Decls {
		decls, err := traverseAst(d)
		if err != nil {
			return nil, err
		}
		if decls != nil {
			resp = append(resp, decls...)
		}
	}

	return &File{
		PkgName: n.Name.Name,
		Decls:   resp,
	}, nil
}

func traverseAst(n ast.Node) ([]*Decl, error) {
	var resp []*Decl

	switch n := n.(type) {
	case *ast.DeclStmt:
		decls, err := traverseAst(n.Decl)
		if err != nil {
			return nil, err
		}

		if decls != nil {
			resp = append(resp, decls...)
		}
	case *ast.GenDecl:
		if n.Tok.String() == "var" {
			for _, spec := range n.Specs {
				decls, err := traverseAst(spec)
				if err != nil {
					return nil, err
				}
				if decls != nil {
					resp = append(resp, decls...)
				}
			}
		}
	case *ast.ValueSpec:
		// カンマ区切りで左辺に複数宣言するのには対応しない
		declare := traverseDeclareBlock(n.Values[0])
		if declare != nil {
			declare.Name = n.Names[0].Name
			resp = append(resp, declare)
		}
	}

	return resp, nil
}

func traverseDeclareBlock(v ast.Expr) *Decl {
	switch v := v.(type) {
	case *ast.CallExpr:
		decl := traverseDeclareBlock(v.Fun)
		if decl == nil {
			return nil
		}

		if decl.StatusCodeEnable && len(v.Args) == 1 {
			arg0, ok := v.Args[0].(*ast.BasicLit)
			if !ok {
				return decl
			}
			decl.StatusCode, _ = strconv.Atoi(arg0.Value)
			return decl
		}

		if len(v.Args) < 2 {
			return decl
		}

		arg0, ok := v.Args[0].(*ast.BasicLit)
		if !ok {
			return decl
		}
		arg1, ok := v.Args[1].(*ast.BasicLit)
		if !ok {
			return decl
		}

		decl.Code = strings.Trim(arg0.Value, `"`)
		decl.Format = strings.Trim(arg1.Value, `"`)
		return decl

	case *ast.SelectorExpr:
		if vi, ok := v.X.(*ast.Ident); ok {
			if vi.Name == "errcdgen" && v.Sel.Name == "NewCodeError" {
				return &Decl{} // 空で返す
			}
		}

		decl := traverseDeclareBlock(v.X)
		if decl == nil {
			return nil
		}

		if v.Sel.Name == "DisableError" {
			decl.DisableErr = true
		}

		if v.Sel.Name == "TraceLevel" {
			decl.LogLevelEnable = true
			decl.LogLevel = errcdgen.TraceLevel
		}

		if v.Sel.Name == "DebugLevel" {
			decl.LogLevelEnable = true
			decl.LogLevel = errcdgen.DebugLevel
		}

		if v.Sel.Name == "InfoLevel" {
			decl.LogLevelEnable = true
			decl.LogLevel = errcdgen.InfoLevel
		}

		if v.Sel.Name == "WarnLevel" {
			decl.LogLevelEnable = true
			decl.LogLevel = errcdgen.WarnLevel
		}

		if v.Sel.Name == "ErrorLevel" {
			decl.LogLevelEnable = true
			decl.LogLevel = errcdgen.ErrorLevel
		}

		if v.Sel.Name == "FatalLevel" {
			decl.LogLevelEnable = true
			decl.LogLevel = errcdgen.FatalLevel
		}

		if v.Sel.Name == "WithStatusCode" {
			decl.LogLevelEnable = true
			decl.StatusCodeEnable = true
		}

		return decl
	}

	return nil
}

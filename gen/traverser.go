package gen

import (
	"gitlab.com/osaki-lab/errcdgen"
	"go/ast"
	"strings"
)

type File struct {
	PkgName  string
	Declares []*DeclareErr
}

type DeclareErr struct {
	Name       string
	Code       string
	Format     string
	LogLevel   errcdgen.Level
	StatusCode int
	ErrDisable bool
}

func (f File) Bindings() []Binding {
	var params []Binding
	for _, declare := range f.Declares {
		//fmt.Printf("%+v\n", declare)
		params = append(params, Binding{
			Name: declare.Name,
		})
	}
	return params
}

func Traverse(n *ast.File) (*File, error) {

	var resp []*DeclareErr
	for _, decl := range n.Decls {
		decls, err := traverseAst(decl)
		if err != nil {
			return nil, err
		}
		if decls != nil {
			resp = append(resp, decls...)
		}
	}

	return &File{
		PkgName:  n.Name.Name,
		Declares: resp,
	}, nil
}

func traverseAst(n ast.Node) ([]*DeclareErr, error) {
	var resp []*DeclareErr

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

func traverseDeclareBlock(v ast.Expr) *DeclareErr {
	if v, ok := v.(*ast.CallExpr); ok {
		if vFun, ok := v.Fun.(*ast.SelectorExpr); ok {

			vFunX, ok := vFun.X.(*ast.Ident)
			if !ok || vFunX.Name != "errcdgen" {
				return nil
			}

			if vFun.Sel.Name != "NewCodeError" {
				return nil
			}

			arg0, ok := v.Args[0].(*ast.BasicLit)
			if !ok {
				return nil
			}
			arg1, ok := v.Args[1].(*ast.BasicLit)
			if !ok {
				return nil
			}
			return &DeclareErr{
				Name:       "",
				Code:       strings.Trim(arg0.Value, `"`),
				Format:     strings.Trim(arg1.Value, `"`),
				LogLevel:   0,
				StatusCode: 0,
				ErrDisable: false,
			}
		}
	}

	// TODO メソッドチェーンの解析

	return nil
}

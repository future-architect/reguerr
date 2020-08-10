package gen

import (
	"go/ast"
	"strings"
)

func Traverse(n ast.Node) ([]*DeclareErr, error) {
	var resp []*DeclareErr

	switch n := n.(type) {
	case *ast.File:
		for _, decl := range n.Decls {
			decls, err := Traverse(decl)
			if err != nil {
				return nil, err
			}
			if decls != nil {
				resp = append(resp, decls...)
			}
		}
	case *ast.DeclStmt:
		decls, err := Traverse(n.Decl)
		if err != nil {
			return nil, err
		}

		if decls != nil {
			resp = append(resp, decls...)
		}
	case *ast.GenDecl:
		if n.Tok.String() == "var" {
			for _, spec := range n.Specs {
				decls, err := Traverse(spec)
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
				ExitCode:   0,
				ErrDisable: false,
			}
		}
	}

	// TODO メソッドチェーンの解析

	return nil
}

/*
Copyright © 2020 osaki-lab mano.junki@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	gen "gitlab.com/osaki-lab/errcdgen/gen"

	//"gitlab.com/osaki-lab/errcdgen/gen"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

var (
	// チェック対象のファイル
	file string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "コード付きエラー生成",
	Short: `errcdgenのルールに則った変数から、初期化関数を生成するコマンドです`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("test", file)

		wd, err := os.Getwd()
		if err != nil {
			return err
		}
		filename := file
		if !strings.HasSuffix(filename, ".go") {
			filename = filename + ".go"
		}

		path := filepath.Join(wd, filename)

		fset := token.NewFileSet()
		f, err := parser.ParseFile(fset, path, nil, 0)
		if err != nil {
			fmt.Printf("Failed to parse file: %v\n", err)
			return nil
		}
		//ast.Print(fset, f)

		traverse, err := traverse(f)
		if err != nil {
			return err
		}

		for _, declare := range traverse {
			fmt.Printf("%+v\n", declare)
		}

		return nil
	},
}

func traverse(n ast.Node) ([]*gen.DeclareErr, error) {
	var resp []*gen.DeclareErr

	switch n := n.(type) {

	case *ast.File:
		for _, decl := range n.Decls {
			decls, err := traverse(decl)
			if err != nil {
				return nil, err
			}
			resp = append(resp, decls...)
		}
	case *ast.DeclStmt:
		declares, err := traverse(n.Decl)
		if err != nil {
			return nil, err
		}
		resp = append(resp, declares...)

	case *ast.GenDecl:
		if n.Tok.String() == "var" {
			for _, spec := range n.Specs {
				declares, err := traverse(spec)
				if err != nil {
					return nil, err
				}
				resp = append(resp, declares...)
			}
		}

	case *ast.ValueSpec:
		// カンマ区切りで左辺に複数宣言するのには対応しない
		declare := DeclareCdErr(n.Values[0])
		if declare != nil {
			declare.Name = n.Names[0].Name
			resp = append(resp, declare)
		}
	}

	return resp, nil
}

func DeclareCdErr(v ast.Expr) *gen.DeclareErr {
	ve, ok := v.(*ast.CallExpr)
	if !ok {
		return nil
	}

	fe, ok := ve.Fun.(*ast.SelectorExpr)
	if !ok {
		return nil
	}

	xe, ok := fe.X.(*ast.CallExpr)
	if !ok {
		return nil
	}

	xf, ok := xe.Fun.(*ast.SelectorExpr)
	if !ok {
		return nil
	}

	if xf.X.(*ast.Ident).Name != "errcdgen" || xf.Sel.Name != "NewCodeError" {
		return nil
	}

	arg0, ok := xe.Args[0].(*ast.BasicLit)
	if !ok {
		// 発生しないはず
		return nil
	}
	errCode := arg0.Value

	arg1, ok := xe.Args[1].(*ast.BasicLit)
	if !ok {
		return nil
	}
	format := arg1.Value

	return &gen.DeclareErr{
		Name:       "",
		Code:       errCode,
		Format:     format,
		LogLevel:   0,
		ExitCode:   0,
		ErrDisable: false,
	}

	// TODO メソッドチェーンの解析

	return nil
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Options
	rootCmd.Flags().StringVarP(&file, "file", "f", "", "input go file")
	_ = rootCmd.MarkFlagRequired("file")
}

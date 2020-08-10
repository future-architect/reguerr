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
	"bytes"
	"fmt"
	"github.com/spf13/cobra"
	"gitlab.com/osaki-lab/errcdgen/gen"
	"go/parser"
	"go/token"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var (
	// チェック対象のファイル
	file string
)

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

		f, err := parser.ParseFile(token.NewFileSet(), path, nil, 0)
		if err != nil {
			fmt.Printf("Failed to parse file: %v\n", err)
			return nil
		}
		//ast.Print(fset, f)

		traverse, err := gen.Traverse(f)
		if err != nil {
			return err
		}

		var params []gen.Binding
		for _, declare := range traverse.Declares {
			//fmt.Printf("%+v\n", declare)
			params = append(params, gen.Binding{
				Name: declare.Name,
			})
		}

		content, err := gen.Generate(traverse.PkgName, params)
		if err != nil {
			return err
		}

		outFile := strings.Replace(filename, ".go", "_errcdgen.go", 1)

		out, err := os.Create(outFile)
		if err != nil {
			return err
		}
		defer out.Close()

		_, err = io.Copy(out, bytes.NewReader(content))

		return err
	},
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

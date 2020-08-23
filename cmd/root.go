/*
Copyright © 2020 reguerr

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
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"gitlab.com/osaki-lab/reguerr/gen"
	"go/parser"
	"go/token"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var (
	// input file path
	file string
)

var rootCmd = &cobra.Command{
	Use:           "code generator for error handling with message code",
	Short:         `code generator for error handling with message code`,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		wd, err := os.Getwd()
		if err != nil {
			return err
		}
		path := filepath.Join(wd, file)

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

		if err := gen.Validate(traverse.Decls); err != nil {
			return fmt.Errorf("input file contains invalid content: %v\n", err)
		}

		content, err := gen.Generate(traverse.PkgName, traverse.Decls)
		if err != nil {
			return err
		}

		out, err := os.Create(strings.Replace(file, ".go", "_gen.go", 1))
		if err != nil {
			return err
		}
		defer out.Close()

		_, err = io.Copy(out, bytes.NewReader(content))
		if err != nil {
			return err
		}

		doc, err := os.Create(strings.Replace(file, ".go", "_gen.md", 1))
		if err != nil {
			return err
		}
		defer doc.Close()

		if err := gen.GenerateMarkdown(doc, traverse.Decls); err != nil {
			return fmt.Errorf("generate markdown: %w", err)
		}

		return err
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		color.Red(err.Error())
		os.Exit(1)
	}
}

func init() {
	// Options
	rootCmd.Flags().StringVarP(&file, "file", "f", "", "input go file")
	_ = rootCmd.MarkFlagRequired("file")
}

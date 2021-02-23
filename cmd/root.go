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
	"gitlab.com/osaki-lab/reguerr"
	"gitlab.com/osaki-lab/reguerr/gen"
	"go/parser"
	"go/token"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var (
	file       string
	errLevel   string
	statusCode int
)

var rootCmd = &cobra.Command{
	Use: "reguerr is code generator for systematic error handling with message code",
}

var generateCmd = &cobra.Command{
	Use:           "generate",
	Short:         "generate reguerr code",
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		wd, err := os.Getwd()
		if err != nil {
			return err
		}
		path := filepath.Join(wd, file)

		f, err := parser.ParseFile(token.NewFileSet(), path, nil, 0)
		if err != nil {
			return fmt.Errorf("failed to parse file: %v\n", err)
		}
		//ast.Print(fset, f)

		traverse, err := gen.Traverse(f)
		if err != nil {
			return err
		}

		if err := gen.Validate(traverse.Decls); err != nil {
			return fmt.Errorf("input file contains invalid content: %v\n", err)
		}

		var opts []gen.Option
		if errLevel != "" {
			level, err := reguerr.NewLevel(errLevel + "Level")
			if err != nil{
				return err
			}
			opts = append(opts, gen.DefaultErrorLevel(level))
		}
		if statusCode != -1 {
			opts = append(opts, gen.DefaultStatusCode(statusCode))
		}

		content, err := gen.GenerateCode(traverse, opts...)
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

		if err := gen.GenerateMarkdown(doc, traverse.Decls, opts...); err != nil {
			return fmt.Errorf("generate markdown: %w", err)
		}

		return err
	},
}

var validateCmd = &cobra.Command{
	Use:           "validate",
	Short:         "validate input file",
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		wd, err := os.Getwd()
		if err != nil {
			return err
		}
		path := filepath.Join(wd, file)

		f, err := parser.ParseFile(token.NewFileSet(), path, nil, 0)
		if err != nil {
			return fmt.Errorf("failed to parse file: %v\n", err)
		}

		traverse, err := gen.Traverse(f)
		if err != nil {
			return err
		}

		if err := gen.Validate(traverse.Decls); err != nil {
			return fmt.Errorf("input file contains invalid content: %v\n", err)
		}

		return nil
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
	generateCmd.Flags().StringVarP(&file, "file", "f", "", "input go file")
	_ = generateCmd.MarkFlagRequired("file")
	generateCmd.Flags().StringVarP(&errLevel, "defaultErrorLevel", "", "", "change default log level(Trace,Debug,Info,Warn,Error,Fatal)")
	generateCmd.Flags().IntVarP(&statusCode, "defaultStatusCode", "", -1, "change default status code")

	validateCmd.Flags().StringVarP(&file, "file", "f", "", "input go file")
	_ = validateCmd.MarkFlagRequired("file")

	rootCmd.AddCommand(generateCmd)
	rootCmd.AddCommand(validateCmd)

}

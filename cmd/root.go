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
	"os"
)

var (
	// チェック対象のファイル
	file string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "コード付きエラー生成",
	Short: `errcdgenのルールに則った変数から、初期化関数を生成するコマンドです`,
	RunE: func(cmd *cobra.Command, args []string) error{
		fmt.Println("test", file)






		return nil
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
	rootCmd.Flags().StringVarP(&file, "file", "f","", "input go file")
	_ = rootCmd.MarkFlagRequired("file")
}

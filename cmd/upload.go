// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	//"strings"
	"github.com/spf13/cobra"
)

var local_path string
var bucket string
var prefix string
var level int
// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "upload files",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		local_path = args[0]
		fmt.Println(local_path)
		fmt.Println(bucket)
		fmt.Println(prefix)
		
	},
}
func init() {
	RootCmd.AddCommand(uploadCmd)
	uploadCmd.Flags().SortFlags = false
	uploadCmd.Flags().StringVarP(&bucket, "bucket", "b", "", "bucket name")
	uploadCmd.Flags().StringVarP(&prefix, "prefix", "p", "", "bucket prefix")
	uploadCmd.Flags().IntVarP(&level, "level", "l", 1, "check level")
}

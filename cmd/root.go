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
	"os"
	"github.com/Unknwon/goconfig"
	"github.com/spf13/cobra"
)

var cfgFile string
var verbose bool
var configpath="cos.conf"

var secret_id string
var secret_key string
var appid string
var region string
// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "cosutil",
	Short: "A cos tool",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) { 
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() { 
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.CDM.yaml)")
	RootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	conf, err := goconfig.LoadConfigFile(configpath)
	if err != nil {
        fmt.Println("failed to load config file")
		os.Create(configpath)
	}else{
		secret_id, _ = conf.GetValue(goconfig.DEFAULT_SECTION,"secret_id")
		secret_key, _ = conf.GetValue(goconfig.DEFAULT_SECTION,"secret_key")
		appid, _ = conf.GetValue(goconfig.DEFAULT_SECTION,"appid")
		region, _ = conf.GetValue(goconfig.DEFAULT_SECTION,"region")
		fmt.Println(secret_id)
	}
}

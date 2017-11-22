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
//	"fmt"
	"os"
	"github.com/Unknwon/goconfig"
	"github.com/spf13/cobra"
)
// configCmd represents the config command


var configCmd = &cobra.Command{
	Use:   "config",
	Short: "config your information",
	Run: func(cmd *cobra.Command, args []string) {
		os.Create(configpath)
		conf, _ := goconfig.LoadConfigFile(configpath)
		//fmt.Println(err)
		conf.SetValue(goconfig.DEFAULT_SECTION, "secret_id", secret_id)
		conf.SetValue(goconfig.DEFAULT_SECTION, "secret_key", secret_key)
		conf.SetValue(goconfig.DEFAULT_SECTION, "appid", appid)
		conf.SetValue(goconfig.DEFAULT_SECTION, "region", region)
		
		
		goconfig.SaveConfigFile(conf,configpath)
		//fmt.Println(conf.GetValue(goconfig.DEFAULT_SECTION,"ip"))
	},
}

func init() {
	RootCmd.AddCommand(configCmd)
	configCmd.Flags().StringVarP(&secret_id, "secret_id", "a", "", "secret_id")
	configCmd.Flags().StringVarP(&secret_key, "secret_key", "k", "", "secret_key")
	configCmd.Flags().StringVarP(&appid, "appid", "u", "", "appid")
	configCmd.Flags().StringVarP(&region, "region", "r", "", "region")

}

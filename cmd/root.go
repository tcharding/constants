// Copyright © 2017 Tobin C. Harding <me@tobin.cc>
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

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

const long = `Searches a Golang package for constants.

If package name is not provided, searches all Go files in $PWD.

Example Usage:

# Searches package 'accounts' in $PWD
$ constants accounts

# Searches package 'github.com/USER/accounts' in $GOPATH
$ constants github.com/USER/accounts

# Lists duplicate constants (including filename).
$ constants accounts --duplicates`

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "constants [PACKAGE]",
	Short: "'greps' for Golang constants",
	Long:  long,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("bare command not yet implemented")
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.Flags().Bool("dup", false, "Show duplicate constants.")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".constants") // name of config file (without extension)
	viper.AddConfigPath("$HOME")      // adding home directory as first search path
	viper.AutomaticEnv()              // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

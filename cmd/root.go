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
	"path/filepath"
	"regexp"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tcharding/constants/adt"
	"github.com/tcharding/constants/search"
)

var cfgFile string
var dup bool

const long = `Searches all Go files in current directory for constants.

Example Usage:

# Prints all constants to stderr.
constants

# Lists duplicate constants including filename.
constants --duplicates`

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "constants",
	Short: "'greps' for Golang constants",
	Long:  long,
	Run:   run,
}

// Execute adds all child commands to the root command sets flags appropriately.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.Flags().BoolVar(&dup, "dup", false, "Show duplicate constants.")
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

func run(cmd *cobra.Command, args []string) {
	paths := getPathsToSearch()
	adt := adt.NewADT()

	for _, path := range paths {
		f, err := os.Open(path)
		if err != nil {
			fmt.Errorf("Failed to open %s\n", path)
			continue
		}
		raw := search.ExtractConsants(f)

		adt.AddRawConstants(path, raw)
	}

	if dup {
		adt.Duplicates()
	} else {
		adt.Dump()
	}
}

func getPathsToSearch() []string {
	const (
		ext = ".go"
	)
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	var paths []string
	filepath.Walk(cwd, func(path string, f os.FileInfo, _ error) error {
		if !f.IsDir() {
			r, err := regexp.MatchString(ext, f.Name())
			if err == nil && r {
				paths = append(paths, path)
			}
		}
		return nil
	})
	return paths
}

func printSlice(slice []string) {
	for _, s := range slice {
		fmt.Println(s)
	}
}

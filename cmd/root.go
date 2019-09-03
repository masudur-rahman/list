/*
Copyright © 2019 Masudur Rahman

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
	"io/ioutil"
	"os"
	"sort"
	"strings"

	"github.com/masudur-rahman/list/lib"
	"github.com/masudur-rahman/list/types"
	"github.com/spf13/cobra"
)

var (
	showFile, showFolder, showHidden, showAll bool
)

func isHidden(filename string) bool {
	return filename[0] == '.'
}

var rootCmd = &cobra.Command{
	Use:   "list",
	Short: "List all the files and directories under current directory",
	RunE: func(cmd *cobra.Command, args []string) error {
		cwd, err := os.Getwd()
		if err != nil {
			return err
		}
		fi, err := ioutil.ReadDir(cwd)
		if err != nil {
			return err
		}

		showAll = !(showFile || showFolder)
		items := make([]types.File, 0, len(fi))
		for _, f := range fi {
			if !showHidden && isHidden(f.Name()) {
				continue
			}
			if showFile && !f.IsDir() {
				items = append(items, types.File{Name: f.Name(), IsDir: f.IsDir()})
			}
			if showFolder && f.IsDir() {
				items = append(items, types.File{Name: f.Name(), IsDir: f.IsDir()})
			}
			if showAll {
				items = append(items, types.File{Name: f.Name(), IsDir: f.IsDir()})
			}
		}

		sort.SliceStable(items, func(i, j int) bool {
			return strings.ToLower(items[i].Name) < strings.ToLower(items[j].Name)
		})

		lib.Println(items)
		return nil
	},
}

func init() {
	rootCmd.Flags().BoolVarP(&showFile, "file", "f", false, "If file flag is set only files are shown")
	rootCmd.Flags().BoolVarP(&showFolder, "dir", "d", false, "If dir flag is set only directories are shown")
	rootCmd.Flags().BoolVar(&showHidden, "show-hidden", false, "If show-hidden flag is set hidden files or folders are shown")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

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
package generate

import (
	"os"
	"time"

	"github.com/bruyss/go-object-generator/obwriter"
	"github.com/spf13/cobra"
)

// generateMeasmonsCmd represents the generateMeasmons command
var generateMeasmonsCmd = &cobra.Command{
	Use:   "measmons",
	Short: "Generate measmon objects",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		now := time.Now().Format("20060102_150405")
		obwriter.GenFolderName = obwriter.GenFolderRoot + "/" + now + "_measmons"
		err := os.MkdirAll(obwriter.GenFolderName, 0666)
		if err != nil && !os.IsExist(err) {
			return err
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	GenerateCmd.AddCommand(generateMeasmonsCmd)
}

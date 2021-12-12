/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/vladcostea/toolbox/io"
)

// formatCmd represents the format command
var jsonFormatCmd = &cobra.Command{
	Use:   "format",
	Short: "Format a JSON blob",
	Long:  `Format and indent a JSON blob from stdin or clipboard.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		data, err := io.ReadAll()
		if err != nil {
			return err
		}

		m := map[string]interface{}{}
		if err := json.Unmarshal(data, &m); err != nil {
			return err
		}

		b, err := json.MarshalIndent(m, "", "  ")
		if err != nil {
			return err
		}

		return io.WriteAll(b)
	},
}

func init() {
	jsonCmd.AddCommand(jsonFormatCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// formatCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// formatCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

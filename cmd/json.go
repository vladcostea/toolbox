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
	"fmt"
	"strings"

	tbio "github.com/vladcostea/toolbox/io"

	"github.com/spf13/cobra"
)

var jsonFormatIndent int

// jsonCmd represents the json command
var jsonCmd = &cobra.Command{
	Use:   "json",
	Short: "Validate JSON blob",
	Long:  `Validate a JSON blob from stdin or clipboard`,
	RunE: func(cmd *cobra.Command, args []string) error {
		data, err := tbio.ReadAll()
		if err != nil {
			return err
		}

		var m interface{}
		err = json.Unmarshal(data, &m)
		if err != nil {
			if jsonError, ok := err.(*json.SyntaxError); ok {
				line, character, _ := lineAndCharacter(string(data), int(jsonError.Offset))
				return fmt.Errorf("syntax error at line %d, character %d: %v", line, character, jsonError.Error())
			}
		}

		if jsonFormatIndent > 0 {
			b, err := json.MarshalIndent(m, "", strings.Repeat(" ", jsonFormatIndent))
			if err != nil {
				return err
			}

			return tbio.WriteAll(b)
		}

		return tbio.WriteAll(nil)
	},
}

func lineAndCharacter(input string, offset int) (line int, character int, err error) {
	lf := rune(0x0A)

	if offset > len(input) || offset < 0 {
		return 0, 0, fmt.Errorf("couldn't find offset %d within the input", offset)
	}

	// Humans tend to count from 1.
	line = 1

	for i, b := range input {
		if b == lf {
			line++
			character = 0
		}
		character++
		if i == offset {
			break
		}
	}

	return line, character, nil
}

func init() {
	jsonCmd.Flags().IntVarP(&jsonFormatIndent, "format", "f", 2, "--format [optional size of indent]")
	rootCmd.AddCommand(jsonCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// jsonCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// jsonCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

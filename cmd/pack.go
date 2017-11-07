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
	"errors"

	"github.com/sajayantony/cpacker/packer"
	"github.com/spf13/cobra"
)

var packDirectory string
var cntName string

// packCmd represents the pack command
var packCmd = &cobra.Command{
	Use:   "pack SourceDirectory ImageName",
	Short: "Pack contents from the specified directory",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return errors.New("The source directory and container image name are required")
		}
		packDirectory = args[0]
		cntName = args[1]
		return packer.Pack(packDirectory, cntName)
	},
}

func init() {
	RootCmd.AddCommand(packCmd)
}

/*
Copyright 2022 Nethermind

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
package cli

import (
	"fmt"

	"github.com/Melvillian/sedge/internal/utils"

	"github.com/spf13/cobra"
)

func VersionCmd() *cobra.Command {
	// Build command
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print sedge version",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintln(cmd.OutOrStdout(), "sedge "+utils.CurrentVersion())
		},
	}
	return cmd
}

// Copyright © 2020 Humio Ltd.
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
)

func newAlertsRemoveCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove [flags] <view> <name>",
		Short: "Removes an alert.",
		Long:  `Removes the alert with name '<name>' in the view with name '<view>'.`,
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			view := args[0]
			name := args[1]

			// Get the HTTP client
			client := NewApiClient(cmd)

			err := client.Alerts().Delete(view, name)
			if err != nil {
				cmd.Println(fmt.Errorf("error removing ingest token: %s", err))
				os.Exit(1)
			}

			cmd.Println("Alert removed")
		},
	}

	return cmd
}

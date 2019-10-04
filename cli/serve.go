// Copyright © 2016 Abcum Ltd
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

package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/abcum/cirrius/cnf"
	"github.com/abcum/cirrius/log"
	"github.com/abcum/cirrius/web"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve the current directory",
	PreRun: func(cmd *cobra.Command, args []string) {

		if cnf.Settings.Logging.Output != "none" {
			fmt.Print(logo)
		}

	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {

		if err = web.Setup(); err != nil {
			log.Fatal(err)
			return
		}

		return

	},
	PostRunE: func(cmd *cobra.Command, args []string) (err error) {

		if err = web.Exit(); err != nil {
			log.Fatal(err)
			return
		}

		return

	},
}

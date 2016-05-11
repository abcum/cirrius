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

	"github.com/abcum/magnifio/web"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the magnifio server",
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Print(logo)
	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		return web.Setup(opts)
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		web.Exit()
	},
}

func init() {

	startCmd.PersistentFlags().IntVar(&opts.Port.Http, "port-http", 80, "The port on which to serve the http server.")
	startCmd.PersistentFlags().IntVar(&opts.Port.Https, "port-https", 443, "The port on which to serve the https server.")

}

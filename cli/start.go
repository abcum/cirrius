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
	"os"

	"github.com/spf13/cobra"

	"github.com/abcum/cirrius/log"
	"github.com/abcum/cirrius/web"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the cirrius server",
	PreRun: func(cmd *cobra.Command, args []string) {

		if opts.Logging.Output != "none" {
			fmt.Print(logo)
		}

	},
	RunE: func(cmd *cobra.Command, args []string) (err error) {

		if err = web.Setup(opts); err != nil {
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

func init() {

	host, _ := os.Hostname()

	startCmd.PersistentFlags().StringVarP(&opts.Surreal, "surreal", "s", "https://api.surreal.io", "URL of the SurrealDB server endpoint.")

	startCmd.PersistentFlags().StringVar(&opts.Cert.Crt, "cert-crt", "", "Path to the server certificate. Needed when running in secure mode.")
	startCmd.PersistentFlags().StringVar(&opts.Cert.Key, "cert-key", "", "Path to the server private key. Needed when running in secure mode.")

	startCmd.PersistentFlags().StringVarP(&opts.Node.Host, "bind", "b", "0.0.0.0", "The hostname or ip address to listen for connections on.")
	startCmd.PersistentFlags().StringVarP(&opts.Node.Name, "name", "n", host, "The name of this node, used for logs and statistics.")

	startCmd.PersistentFlags().IntVarP(&opts.Port.Web, "port", "p", 80, "The port on which to serve the server.")

}

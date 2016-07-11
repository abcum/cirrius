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

	"github.com/abcum/cirrius/web"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the cirrius server",
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

	startCmd.PersistentFlags().StringVar(&opts.Cert.Crt, "cert-crt", "", "Path to the server certificate. Needed in secure mode.")
	startCmd.PersistentFlags().StringVar(&opts.Cert.Key, "cert-key", "", "Path to the server private key. Needed in secure mode.")
	startCmd.PersistentFlags().StringVar(&opts.Cert.Pem, "cert-pem", "", "The PEM encoded certificate and private key data. Use this as an alternative to the --cert-crt and --cert-key flags.")

	startCmd.PersistentFlags().IntVar(&opts.Port.Web, "port-web", 80, "The port on which to serve the web server.")

}

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
	"os"

	"github.com/spf13/cobra"

	"github.com/abcum/magnifio/cnf"
)

var opts *cnf.Options

var mainCmd = &cobra.Command{
	Use:   "magnifio",
	Short: "Magnifio server",
}

func init() {

	mainCmd.AddCommand(
		startCmd,
		versionCmd,
	)

	opts = &cnf.Options{}

	mainCmd.PersistentFlags().StringVar(&opts.Logging.Level, "log-level", "error", "Specify log verbosity")
	mainCmd.PersistentFlags().StringVar(&opts.Logging.Output, "log-output", "stderr", "Specify log output destination")
	mainCmd.PersistentFlags().StringVar(&opts.Logging.Format, "log-format", "text", "Specify log output format (text, json)")
	mainCmd.PersistentFlags().StringVar(&opts.Logging.Newrelic, "log-newrelic", "", "Log to Newrelic using the specified license key")
	mainCmd.PersistentFlags().MarkHidden("log-newrelic")

	cobra.OnInitialize(setup)

}

// Init runs the cli app
func Init() {
	if err := mainCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}

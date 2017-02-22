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

	"github.com/abcum/cirrius/cnf"
)

var opts *cnf.Options

var mainCmd = &cobra.Command{
	Use:   "cirrius",
	Short: "Cirrius server",
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

	mainCmd.PersistentFlags().StringVar(&opts.Logging.Google.Name, "log-driver-google-name", "cirrius", "Specify a log name")
	mainCmd.PersistentFlags().StringVar(&opts.Logging.Google.Project, "log-driver-google-project", "", "Specify a project id")
	mainCmd.PersistentFlags().StringVar(&opts.Logging.Google.Credentials, "log-driver-google-credentials", "", "Specify the path to a credentials file")
	mainCmd.PersistentFlags().MarkHidden("log-driver-google-name")
	mainCmd.PersistentFlags().MarkHidden("log-driver-google-project")
	mainCmd.PersistentFlags().MarkHidden("log-driver-google-credentials")

	mainCmd.PersistentFlags().StringVar(&opts.Logging.Syslog.Tag, "log-driver-syslog-tag", "cirrius", "Specify a syslog name tag")
	mainCmd.PersistentFlags().StringVar(&opts.Logging.Syslog.Host, "log-driver-syslog-host", "localhost:514", "Specify a remote host:port")
	mainCmd.PersistentFlags().StringVar(&opts.Logging.Syslog.Protocol, "log-driver-syslog-protocol", "", "Specify the syslog protocol to use")
	mainCmd.PersistentFlags().StringVar(&opts.Logging.Syslog.Priority, "log-driver-syslog-priority", "debug", "Specify the syslog priority")
	mainCmd.PersistentFlags().MarkHidden("log-driver-syslog-tag")
	mainCmd.PersistentFlags().MarkHidden("log-driver-syslog-host")
	mainCmd.PersistentFlags().MarkHidden("log-driver-syslog-protocol")
	mainCmd.PersistentFlags().MarkHidden("log-driver-syslog-priority")

	cobra.OnInitialize(setup)

}

// Init runs the cli app
func Init() {
	if err := mainCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}

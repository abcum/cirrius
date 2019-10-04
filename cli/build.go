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
	"github.com/abcum/cirrius/log"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build and package the current directory",
	RunE: func(cmd *cobra.Command, args []string) (err error) {

		// Ensure that the command has a build type
		// specified in the command arguments.

		switch cnf.Settings.Kind {
		case "aws", "as3", "linux-amd64":
		default:
			log.Fatalln("Invalid build type provided.")
		}

		// Remove any previous server build which
		// was downloaded in a previous build.

		os.RemoveAll("main")

		// Remove any previous zip file which may
		// have been created in a previous build.

		os.RemoveAll("main.zip")

		// Download the server for the specified
		// type so we can bundle it together.

		if err = saver(cnf.Settings.Kind); err != nil {
			log.Fatalln("Unable to build package.")
			return
		}

		// Zip the contents of the current folder
		// and bundle it with the specified server.

		switch cnf.Settings.Kind {

		case "aws", "as3":

			defer os.RemoveAll("main")

			if err = zipper("main.zip"); err != nil {
				log.Fatalln("Unable to build package.")
				return
			}

		}

		return

	},
}

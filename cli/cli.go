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

var mainCmd = &cobra.Command{
	Use:   "cirrius",
	Short: "Cirrius",
}

func init() {

	mainCmd.AddCommand(
		buildCmd,
		serveCmd,
		versionCmd,
	)

	cobra.OnInitialize(cnf.Setup)

}

// Init runs the cli app
func Init() {
	if err := mainCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}

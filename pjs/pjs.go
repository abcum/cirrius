// Copyright © 2016 Abcum Ltd
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//,
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pjs

import (
	"os/exec"
	"runtime"

	"github.com/abcum/cirrius/cnf"
	"github.com/abcum/cirrius/log"
)

// Setup sets up the server for remote connections
func Setup(opts *cnf.Options) (err error) {

	log.WithPrefix("pjs").Infof("Starting pjs server on %s", "localhost:8080")

	go func() {
		cmd := exec.Command("./pjs/phantom-"+runtime.GOOS, "pjs/phantom.js")
		cmd.Start()
		cmd.Wait()
	}()

	return nil

}

// Exit tears down the server gracefully
func Exit() {
	log.WithPrefix("pjs").Infof("Gracefully shutting down %s protocol", "pjs")
}

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

package web

import (
	"github.com/abcum/fibre"
	"github.com/abcum/fibre/mw"

	"github.com/abcum/cirrius/cnf"
	"github.com/abcum/cirrius/log"
)

// Setup sets up the server for remote connections
func Setup() (err error) {

	log.WithPrefix("web").Infof("Starting web server on %s", cnf.Settings.Conn)

	s := fibre.Server()

	routes(s)
	s.SetName("web")
	s.SetIdleTimeout("60s")
	s.SetHTTPErrorHandler(errors)
	s.Logger().SetLogger(log.Instance())

	// Setup middleware

	s.Use(mw.Fail()) // Catch panics
	s.Use(mw.Logs()) // Log requests

	// Log successful start

	log.WithPrefix("web").Infof("Started web server on %s", cnf.Settings.Conn)

	// Run the server

	if len(cnf.Settings.Cert.Crt) == 0 || len(cnf.Settings.Cert.Key) == 0 {
		log.WithPrefix("web").Infof("Running at http://%s", cnf.Settings.Conn)
		s.Run(cnf.Settings.Conn)
	}

	if len(cnf.Settings.Cert.Crt) != 0 && len(cnf.Settings.Cert.Key) != 0 {
		log.WithPrefix("web").Infof("Running at https://%s", cnf.Settings.Conn)
		s.Run(cnf.Settings.Conn, cnf.Settings.Cert.Crt, cnf.Settings.Cert.Key)
	}

	return nil

}

// Exit tears down the server gracefully
func Exit() (err error) {

	log.WithPrefix("web").Infof("Gracefully shutting down %s protocol", "web")

	return

}

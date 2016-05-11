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

// +build !appengine

package web

import (
	"github.com/abcum/fibre"
	"github.com/abcum/fibre/mw"
	"github.com/abcum/magnifio/cnf"
	"github.com/abcum/magnifio/log"
)

// Setup sets up the server for remote connections
func Setup(opts *cnf.Options) (err error) {

	log.WithPrefix("web").Infof("Starting web server on %s", opts.Conn.Http)

	s := fibre.Server(opts)

	routes(s)
	s.SetWait(60)
	s.SetName("web")
	s.SetHTTPErrorHandler(errors)
	s.Logger().SetLogger(log.Instance())

	// Setup middleware

	s.Use(mw.Logs()) // Log requests
	s.Use(mw.Fail()) // Catch panics
	s.Use(mw.Gzip()) // Gzip responses
	s.Use(mw.Uniq()) // Add uniq headers
	s.Use(mw.Cors()) // Add cors headers

	// Check body size

	s.Use(mw.Size(&mw.SizeOpts{
		AllowedLength: 1000000,
	}))

	// Setup newrelic integration

	s.Use(mw.Newrelic(&mw.NewrelicOpts{
		Name:    []byte("Magnifio"),
		License: []byte(opts.Logging.Newrelic),
	}))

	// Run the server

	s.Run(opts.Conn.Http)

	return nil

}

// Exit tears down the server gracefully
func Exit() {
	log.WithPrefix("web").Infof("Gracefully shutting down %s protocol", "web")
}
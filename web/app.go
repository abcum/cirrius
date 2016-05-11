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

// +build appengine

package web

import (
	"net/http"

	"github.com/abcum/fibre"
	"github.com/abcum/fibre/mw"
)

func init() {

	s := fibre.Server(nil)

	routes(s)
	s.SetName("web")
	s.SetHTTPErrorHandler(errors)

	// Setup middleware

	s.Use(mw.Logs()) // Log requests
	s.Use(mw.Fail()) // Catch panics
	s.Use(mw.Gzip()) // Gzip responses
	s.Use(mw.Uniq()) // Add uniq headers
	s.Use(mw.Cors()) // Add cors headers

	// Run the server

	http.Handle("/", s)

}

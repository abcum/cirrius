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

	"github.com/abcum/cirrius/log"
	"github.com/abcum/cirrius/util/cert"
)

func setup() {

	// --------------------------------------------------
	// Ports
	// --------------------------------------------------

	// Specify default port
	if opts.Port.Web == 0 {
		opts.Port.Web = 80
	}

	// Ensure port number is valid
	if opts.Port.Web < 0 || opts.Port.Web > 65535 {
		log.Fatal("Please specify a valid port number for --port-web")
	}

	// Store the ports in host:port string format
	opts.Conn.Web = fmt.Sprintf(":%d", opts.Port.Web)

	// --------------------------------------------------
	// Certs
	// --------------------------------------------------

	if opts.Cert.Pem != "" {

		if opts.Cert.Key != "" || opts.Cert.Crt != "" {
			log.Fatal("You can not specify --cert-pem with --cert-key or --cert-crt")
		}

		err := cert.Extract(opts.Cert.Pem, "cert.key", "cert.crt")
		if err != nil {
			log.Fatal(err)
		}

		opts.Cert.Key = "cert.key"
		opts.Cert.Crt = "cert.crt"

	}

	// --------------------------------------------------
	// Logging
	// --------------------------------------------------

	// Ensure that the specified
	// logging level is allowed

	if opts.Logging.Level != "" {

		chk := map[string]bool{
			"debug":   true,
			"info":    true,
			"warning": true,
			"error":   true,
			"fatal":   true,
			"panic":   true,
		}

		if _, ok := chk[opts.Logging.Level]; !ok {
			log.Fatal("Incorrect log level specified")
		}

		log.SetLevel(opts.Logging.Level)

	}

	// Ensure that the specified
	// logging output is allowed

	if opts.Logging.Output != "" {

		chk := map[string]bool{
			"stdout": true,
			"stderr": true,
		}

		if _, ok := chk[opts.Logging.Output]; !ok {
			log.Fatal("Incorrect log output specified")
		}

		log.SetOutput(opts.Logging.Output)

	}

	// Ensure that the specified
	// logging format is allowed

	if opts.Logging.Format != "" {

		chk := map[string]bool{
			"text": true,
			"json": true,
		}

		if _, ok := chk[opts.Logging.Format]; !ok {
			log.Fatal("Incorrect log format specified")
		}

		log.SetFormat(opts.Logging.Format)

	}

}

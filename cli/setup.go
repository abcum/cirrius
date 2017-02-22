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
	"os"
	"strings"

	"github.com/abcum/cirrius/cnf"
	"github.com/abcum/cirrius/log"
	"github.com/abcum/cirrius/util/uuid"
)

func setup() {

	// --------------------------------------------------
	// Nodes
	// --------------------------------------------------

	// Ensure that the default
	// node details are defined

	if opts.Node.Host == "" {
		opts.Node.Host, _ = os.Hostname()
	}

	if opts.Node.Name == "" {
		opts.Node.Name = opts.Node.Host
	}

	if opts.Node.UUID == "" {
		opts.Node.UUID = opts.Node.Name + "-" + uuid.NewV4().String()
	}

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
	opts.Conn.Web = fmt.Sprintf("%s:%d", opts.Node.Host, opts.Port.Web)

	// --------------------------------------------------
	// Certs
	// --------------------------------------------------

	if strings.HasPrefix(opts.Cert.Crt, "-----") {
		var err error
		var doc *os.File
		if doc, err = os.Create("cert.crt"); err != nil {
			log.Fatal("Can not decode PEM encoded certificate into cert.crt")
		}
		doc.Write([]byte(opts.Cert.Crt))
		doc.Close()
		opts.Cert.Crt = "cert.crt"
	}

	if strings.HasPrefix(opts.Cert.Key, "-----") {
		var err error
		var doc *os.File
		if doc, err = os.Create("cert.key"); err != nil {
			log.Fatal("Can not decode PEM encoded private key into cert.key")
		}
		doc.Write([]byte(opts.Cert.Key))
		doc.Close()
		opts.Cert.Key = "cert.key"
	}

	// --------------------------------------------------
	// Logging
	// --------------------------------------------------

	var chk map[string]bool

	// Ensure that the specified
	// logging level is allowed

	if opts.Logging.Level != "" {

		chk = map[string]bool{
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
	// logging format is allowed

	if opts.Logging.Format != "" {

		chk = map[string]bool{
			"text": true,
			"json": true,
		}

		if _, ok := chk[opts.Logging.Format]; !ok {
			log.Fatal("Incorrect log format specified")
		}

		log.SetFormat(opts.Logging.Format)

	}

	// Ensure that the specified
	// logging output is allowed

	if opts.Logging.Output != "" {

		chk = map[string]bool{
			"none":   true,
			"stdout": true,
			"stderr": true,
			"google": true,
			"syslog": true,
		}

		if _, ok := chk[opts.Logging.Output]; !ok {
			log.Fatal("Incorrect log output specified")
		}

		if opts.Logging.Output == "google" {

			hook, err := log.NewGoogleHook(
				opts.Logging.Level,
				opts.Logging.Google.Name,
				opts.Logging.Google.Project,
				opts.Logging.Google.Credentials,
			)

			if err != nil {
				log.Fatal("There was a problem configuring the google logging driver")
			}

			log.Instance().Hooks.Add(hook)

			opts.Logging.Output = "none"

		}

		if opts.Logging.Output == "syslog" {

			chk = map[string]bool{
				"":    true,
				"tcp": true,
				"udp": true,
			}

			if _, ok := chk[opts.Logging.Syslog.Protocol]; !ok {
				log.Fatal("Incorrect log protocol specified for syslog logging driver")
			}

			chk = map[string]bool{
				"debug":   true,
				"info":    true,
				"notice":  true,
				"warning": true,
				"err":     true,
				"crit":    true,
				"alert":   true,
				"emerg":   true,
			}

			if _, ok := chk[opts.Logging.Syslog.Priority]; !ok {
				log.Fatal("Incorrect log priority specified for syslog logging driver")
			}

			hook, err := log.NewSyslogHook(
				opts.Logging.Level,
				opts.Logging.Syslog.Protocol,
				opts.Logging.Syslog.Host,
				opts.Logging.Syslog.Priority,
				opts.Logging.Syslog.Tag,
			)

			if err != nil {
				log.Fatal("There was a problem configuring the syslog logging driver")
			}

			log.Instance().Hooks.Add(hook)

			opts.Logging.Output = "none"

		}

		log.SetOutput(opts.Logging.Output)

	}

	// Enable global options object

	cnf.Settings = opts

}

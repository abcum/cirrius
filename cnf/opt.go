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

package cnf

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"strings"

	"gopkg.in/yaml.v2"

	"github.com/abcum/cirrius/log"
)

func Setup() {

	opts := &Options{}

	// --------------------------------------------------
	// Yaml
	// --------------------------------------------------

	if bit, err := ioutil.ReadFile("config.yaml"); err == nil {
		if err := yaml.Unmarshal(bit, &opts); err != nil {
			log.Fatalf("Yaml configuration file is invalid: %s", err)
		}
	}

	// --------------------------------------------------
	// Nodes
	// --------------------------------------------------

	// Specify default port
	if opts.Port == 0 {
		opts.Port = 1111
	}

	// Specify default host
	if opts.Host == "" {
		opts.Host = "localhost"
	}

	// Ensure port number is valid
	if opts.Port < 0 || opts.Port > 65535 {
		log.Fatal("Please specify a valid port number for --port")
	}

	// Store the ports in host:port string format
	opts.Conn = fmt.Sprintf("%s:%d", opts.Host, opts.Port)

	// --------------------------------------------------
	// Certs
	// --------------------------------------------------

	if strings.HasPrefix(opts.Cert.Crt, "-----") {
		var err error
		var doc *os.File
		var out string = path.Join(os.TempDir(), "cirrius.crt")
		if doc, err = os.Create(out); err != nil {
			log.Fatalf("Can not decode PEM encoded certificate into %s", out)
		}
		doc.Write([]byte(opts.Cert.Crt))
		doc.Close()
		opts.Cert.Crt = out
	}

	if strings.HasPrefix(opts.Cert.Key, "-----") {
		var err error
		var doc *os.File
		var out string = path.Join(os.TempDir(), "cirrius.key")
		if doc, err = os.Create(out); err != nil {
			log.Fatalf("Can not decode PEM encoded private key into %s", out)
		}
		doc.Write([]byte(opts.Cert.Key))
		doc.Close()
		opts.Cert.Key = out
	}

	if opts.Cert.Key != "" {
		if _, err := os.Open(opts.Cert.Key); err != nil {
			log.Fatalf("Can not open certificate key file %s", opts.Cert.Key)
		}
	}

	if opts.Cert.Crt != "" {
		if _, err := os.Open(opts.Cert.Crt); err != nil {
			log.Fatalf("Can not open certificate crt file %s", opts.Cert.Crt)
		}
	}

	// --------------------------------------------------
	// Chrome
	// --------------------------------------------------

	// Specify default endpoint
	if opts.Chrome.Endpoint == "" {
		opts.Chrome.Endpoint = "http://127.0.0.1:9515"
	}

	// Specify default executable
	if opts.Chrome.Executable == "" {
		switch runtime.GOOS {
		case "windows":
			opts.Chrome.Executable = "C:\\Program Files (x86)\\Google\\Chrome\\Application\\chrome.exe"
		case "darwin":
			opts.Chrome.Executable = "/Applications/Google Chrome.app/Contents/MacOS/Google Chrome"
		case "linux":
			opts.Chrome.Executable = "/usr/bin/chrome"
		}
	}

	// --------------------------------------------------
	// Logging
	// --------------------------------------------------

	var chk map[string]bool

	// Ensure that the specified
	// logging level is allowed

	if opts.Logging.Level == "" {
		opts.Logging.Level = "debug"
	}

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

	if opts.Logging.Format == "" {
		opts.Logging.Format = "text"
	}

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

	if opts.Logging.Output == "" {
		opts.Logging.Output = "stderr"
	}

	if opts.Logging.Output != "" {

		chk = map[string]bool{
			"none":   true,
			"stdout": true,
			"stderr": true,
		}

		if _, ok := chk[opts.Logging.Output]; !ok {
			log.Fatal("Incorrect log output specified")
		}

		log.SetOutput(opts.Logging.Output)

	}

	// Enable global options object

	Settings = opts

}

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

var Settings *Options

// Options defines global configuration options
type Options struct {
	Surreal string // Surreal host:port to connect to

	Port struct {
		Web int // Web port as an number
	}

	Conn struct {
		Web string // Web port as a string
	}

	Cert struct {
		Crt string // File location of server crt
		Key string // File location of server key
	}

	Node struct {
		Host string // Node hostname
		Name string // Name of this node
		UUID string // UUID of this node
	}

	Chrome struct {
		Endpoint   string // URL endpoint of chromedriver
		Executable string // Path to Google Chrome executable
	}

	Logging struct {
		Level  string // Stores the configured logging level
		Output string // Stores the configured logging output
		Format string // Stores the configured logging format
		Google struct {
			Name        string // Stores the GCE logging name
			Project     string // Stores the GCE logging project
			Credentials string // Store the path to the credentials file
		}
		Syslog struct {
			Tag      string // Stores the syslog tag name
			Host     string // Stores the syslog remote host:port
			Protocol string // Stores the syslog protocol to use
			Priority string // Stores the syslog logging priority
		}
	}
}

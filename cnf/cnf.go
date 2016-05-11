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

// Options defines global configuration options
type Options struct {
	DB struct {
		Host string // Surreal host to connect to
		Port string // Surreal port to connect to
		Base string // Base key to use in KV stores
	}

	Node struct {
		Host   string // Node hostname
		Name   string // Name of this node
		Region string // Region that this node is located
	}

	Port struct {
		Http  int // Http port
		Https int // Https port
	}

	Conn struct {
		Http  string // Http port
		Https string // Https port
	}

	Logging struct {
		Level    string // Stores the configured logging level
		Output   string // Stores the configured logging output
		Format   string // Stores the configured logging format
		Newrelic string // Stores the configured newrelic license key
	}
}
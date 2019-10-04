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
	Kind string `yaml:"kind"`

	Port int `yaml:"port"`

	Host string `yaml:"host"`

	Conn string `yaml:"-"`

	Cert struct {
		Crt string `yaml:"crt"`
		Key string `yaml:"key"`
	}

	Cache map[string]string `yaml:"cache"`

	Chrome struct {
		Endpoint   string // URL endpoint of chromedriver
		Executable string // Path to Google Chrome executable
	}

	Logging struct {
		Level  string // Stores the configured logging level
		Output string // Stores the configured logging output
		Format string // Stores the configured logging format
	}
}

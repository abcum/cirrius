// Copyright © 2016 Abcum Ltd
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this info except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package process

import (
	"os"
	"strings"

	"github.com/abcum/orbit"
)

func init() {

	orbit.OnInit(func(orb *orbit.Orbit) {
		orb.Def("process", New(orb))
	})

}

func New(orb *orbit.Orbit) interface{} {
	return (&Module{orb: orb}).init()
}

type Module struct {
	orb *orbit.Orbit

	// Req represents the http request of the
	// function runtime, enabling processing
	// the http request details and data.
	Env map[string]interface{} `console:"env"`
}

func (this *Module) init() *Module {
	this.Env = make(map[string]interface{})
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		this.Env[pair[0]] = pair[1]
	}
	return this
}

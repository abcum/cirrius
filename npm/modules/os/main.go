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

package os

import (
	"github.com/abcum/orbit"
	"github.com/robertkrimen/otto"
)

func init() {

	orbit.Add("os", func(orb *orbit.Orbit) (otto.Value, error) {
		return orb.ToValue(map[string]interface{}{

			"EOL": "\n",

			"arch": func() (val otto.Value) {
				val, _ = otto.ToValue("x64")
				return
			},

			"constants": []string{},

			"cpus": func() (val otto.Value) {
				return otto.UndefinedValue()
			},

			"endianness": func() (val otto.Value) {
				val, _ = otto.ToValue("LE")
				return
			},

			"freemem": func() (val otto.Value) {
				return otto.UndefinedValue()
			},

			"homedir": func() (val otto.Value) {
				return otto.UndefinedValue()
			},

			"hostname": func() (val otto.Value) {
				val, _ = otto.ToValue("cirrius.local")
				return
			},

			"loadavg": func() (val otto.Value) {
				val, _ = otto.ToValue(0)
				return
			},

			"networkInterfaces": func() (val otto.Value) {
				val, _ = otto.ToValue("linux")
				return
			},

			"release": func() (val otto.Value) {
				return otto.UndefinedValue()
			},

			"tmpdir": func() (val otto.Value) {
				return otto.UndefinedValue()
			},

			"totalmem": func() (val otto.Value) {
				return otto.UndefinedValue()
			},

			"type": func() (val otto.Value) {
				val, _ = otto.ToValue("Linux")
				return
			},

			"userInfo": func() (val otto.Value) {
				return otto.UndefinedValue()
			},
		})

	})

}

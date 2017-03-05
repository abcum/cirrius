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

package shortid

import (
	"github.com/SKAhack/go-shortid"
	"github.com/abcum/orbit"
	"github.com/robertkrimen/otto"
)

func init() {

	orbit.Add("shortid", func(orb *orbit.Orbit) (otto.Value, error) {

		generator := shortid.Generator()

		return orb.ToValue(map[string]interface{}{
			"generate": func() (val otto.Value) {
				val, _ = orb.ToValue(generator.Generate())
				return
			},
			"seed": func(seed float64) (val otto.Value) {
				generator.SetSeed(seed)
				return
			},
			"characters": func(char string) (val otto.Value) {
				generator.SetCharacters(char)
				val, _ = orb.ToValue(char)
				return
			},
		})

	})

}

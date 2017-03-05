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

package ws

import (
	"github.com/abcum/fibre"
	"github.com/abcum/orbit"
	"github.com/robertkrimen/otto"
)

type task struct {
	callback func()
}

type callback func(interface{})

func (t *task) Startup(orb *orbit.Orbit) {

}

func (t *task) Cleanup(orb *orbit.Orbit) {
	session := orb.Context().Value("fibre").(*fibre.Context)
	session.Socket().Close(1000)
}

func (t *task) Execute(orb *orbit.Orbit) (err error) {
	t.callback()
	return
}

func init() {

	orbit.Add("ws", func(orb *orbit.Orbit) (otto.Value, error) {

		wsocket := &task{}

		session := orb.Context().Value("fibre").(*fibre.Context)

		return orb.ToValue(map[string]interface{}{

			"init": func() {
				if session.Upgrade() == nil {
					orb.Push(wsocket)
				}
			},

			"exit": func() {
				if session.Socket() != nil {
					orb.Pull(wsocket)
				}
			},

			"recv": func(callback callback) {
				if session.Socket() != nil {
					go func() {
						defer orb.Pull(wsocket)
						for {
							if _, req, err := session.Socket().Read(); err == nil {
								orb.Next(&task{func() { callback(string(req)) }})
								continue
							}
							break
						}
					}()
				}
			},

			"send": map[string]interface{}{
				"text": func(data string) {
					if session.Socket() != nil {
						session.Socket().SendText(data)
					}
				},
				"json": func(data interface{}) {
					if session.Socket() != nil {
						session.Socket().SendJSON(data)
					}
				},
				"cbor": func(data interface{}) {
					if session.Socket() != nil {
						session.Socket().SendCBOR(data)
					}
				},
				"binc": func(data interface{}) {
					if session.Socket() != nil {
						session.Socket().SendBINC(data)
					}
				},
				"pack": func(data interface{}) {
					if session.Socket() != nil {
						session.Socket().SendPACK(data)
					}
				},
			},
		})

	})

}

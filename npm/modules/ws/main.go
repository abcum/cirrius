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

func (t *task) Startup(ctx *orbit.Orbit) {

}

func (t *task) Cleanup(ctx *orbit.Orbit) {
	ctx.Vars["fibre"].(*fibre.Context).Socket().Close(1000)
}

func (t *task) Execute(ctx *orbit.Orbit) (err error) {
	t.callback()
	return
}

func init() {

	orbit.Add("ws", func(ctx *orbit.Orbit) (otto.Value, error) {

		wsocket := &task{}

		session := ctx.Vars["fibre"].(*fibre.Context)

		return ctx.ToValue(map[string]interface{}{

			"recv": func(callback callback) {

				if err := session.Upgrade(); err != nil {
					return
				}

				ctx.Push(wsocket)

				go func() {
					defer ctx.Pull(wsocket)
					for {
						if _, req, err := session.Socket().Read(); err == nil {
							ctx.Next(&task{func() { callback(req) }})
							continue
						}
						break
					}
				}()

			},

			"exit": func() {
				ctx.Pull(wsocket)
			},

			"send": map[string]interface{}{
				"text": func(data string) {
					session.Socket().SendText(data)
				},
				"json": func(data interface{}) {
					session.Socket().SendJSON(data)
				},
				"pack": func(data interface{}) {
					session.Socket().SendPACK(data)
				},
			},
		})

	})

}

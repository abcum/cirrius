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

package allot

import (
	"fmt"

	"github.com/abcum/orbit"
	"github.com/robertkrimen/otto"
	"github.com/sbstjn/allot"
)

type cmd struct {
	cmd allot.Command
	fnc otto.Value
}

func init() {
	orbit.Add("text/allot", New)
}

func New(orb *orbit.Orbit) interface{} {
	return &Module{
		orb: orb,
	}
}

type Module struct {
	orb  *orbit.Orbit
	cmds []*cmd
}

func (this *Module) Add(value string, cb otto.Value) {

	this.cmds = append(this.cmds, &cmd{
		cmd: allot.New(value),
		fnc: cb,
	})

}

func (this *Module) Run(value interface{}) {

	var str string

	switch v := value.(type) {
	case fmt.Stringer:
		str = v.String()
	case string:
		str = v
	}

	for _, cmd := range this.cmds {

		if match, err := cmd.cmd.Match(str); err == nil {
			_, err := cmd.fnc.Call(cmd.fnc, NewMatch(this.orb, match))
			if err != nil {
				this.orb.Quit(err)
			}
		}

	}

}

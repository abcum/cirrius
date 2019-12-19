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

package surreal

import (
	"github.com/abcum/orbit"
	"github.com/abcum/otto"

	"github.com/abcum/cirrius/util/args"
)

func init() {
	orbit.Add("db/surreal", New)
}

func New(orb *orbit.Orbit) interface{} {
	return (&Module{orb: orb}).Init
}

type Module struct {
	orb *orbit.Orbit
}

func (this *Module) Init(call otto.FunctionCall) otto.Value {

	args.Size(this.orb, call, 1, 2)

	url := args.String(this.orb, call, 0)
	opt := args.Object(this.orb, call, 1)

	api := NewDB(this.orb, url, opt)

	return args.Value(this.orb, api)

}

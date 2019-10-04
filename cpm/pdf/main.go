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

// +build cgo

/*

The pdf package enables creating and manipulating PDF documents.

	var lib = require('pdf');
	var pdf = lib("xxxxxxx-xxxxxx-xxxxxx-xxxxxx-xxxxxx");
	var doc = new pdf({
		masterpassword: '123456',
		permissions: "noprint nohiresprint nocopy nomodify noaccessible noassemble",
	});

	doc.page(595, 842);

	doc.pipe(context.res.body);

*/
package pdf

import (
	"github.com/abcum/orbit"
	"github.com/robertkrimen/otto"

	"github.com/abcum/cirrius/util/args"
)

func init() {
	orbit.Add("pdf", New)
}

func New(orb *orbit.Orbit) interface{} {
	return (&Module{orb: orb}).Init
}

type Module struct {
	orb *orbit.Orbit
}

func (this *Module) Init(call otto.FunctionCall) otto.Value {

	args.Size(this.orb, call, 0, 1)

	key := args.String(this.orb, call, 0)

	lib := NewLib(this.orb, key).New

	return args.Value(this.orb, lib)

}

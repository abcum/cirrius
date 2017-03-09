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

package routes

import (
	"github.com/abcum/fibre"
	"github.com/abcum/orbit"
	"github.com/robertkrimen/otto"
)

func init() {
	orbit.Add("routes", New)
}

func New(orb *orbit.Orbit) interface{} {
	return &Module{
		orb: orb,
		rtr: &router{routes: make(map[string][]*route)},
		ctx: orb.Context().Value("fibre").(*fibre.Context),
	}
}

type Module struct {
	orb *orbit.Orbit
	rtr *router
	ctx *fibre.Context
}

func (this *Module) add(meth, path string, call otto.Value) {
	this.rtr.add(meth, path, call)
}

func (this *Module) Get(path string, call otto.Value) {
	this.add("GET", path, call)
}

func (this *Module) Put(path string, call otto.Value) {
	this.add("PUT", path, call)
}

func (this *Module) Post(path string, call otto.Value) {
	this.add("POST", path, call)
}

func (this *Module) Patch(path string, call otto.Value) {
	this.add("PATCH", path, call)
}

func (this *Module) Trace(path string, call otto.Value) {
	this.add("TRACE", path, call)
}

func (this *Module) Delete(path string, call otto.Value) {
	this.add("DELETE", path, call)
}

func (this *Module) Options(path string, call otto.Value) {
	this.add("OPTIONS", path, call)
}

func (this *Module) Connect(path string, call otto.Value) {
	this.add("CONNECT", path, call)
}

func (this *Module) Any(path string, call otto.Value) {
	this.add("GET", path, call)
	this.add("PUT", path, call)
	this.add("POST", path, call)
	this.add("PATCH", path, call)
	this.add("TRACE", path, call)
	this.add("DELETE", path, call)
	this.add("OPTIONS", path, call)
	this.add("CONNECT", path, call)
}

func (this *Module) Run() {
	meth := this.ctx.Request().Method
	path := this.ctx.Request().URL().Path
	for _, route := range this.rtr.routes[meth] {
		if vars, ok := route.test(path); ok {
			route.handler.Call(route.handler, vars)
			return
		}
	}
	return
}

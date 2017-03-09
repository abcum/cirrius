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

package context

import (
	"github.com/abcum/fibre"
	"github.com/abcum/orbit"
)

type Status struct {
	orb *orbit.Orbit
	val int
	fib *fibre.Context
}

func NewStatus(orb *orbit.Orbit, val int) *Status {
	return &Status{
		orb: orb,
		val: val,
		fib: orb.Context().Value("fibre").(*fibre.Context),
	}
}

func (this *Status) Xml(data interface{}) {
	this.fib.XML(this.val, data)
	panic(nil)
	return
}

func (this *Status) Text(data interface{}) {
	this.fib.Text(this.val, data)
	panic(nil)
	return
}

func (this *Status) Html(data interface{}) {
	this.fib.HTML(this.val, data)
	panic(nil)
	return
}

func (this *Status) Json(data interface{}) {
	this.fib.JSON(this.val, data)
	panic(nil)
	return
}

func (this *Status) Cbor(data interface{}) {
	this.fib.CBOR(this.val, data)
	panic(nil)
	return
}

func (this *Status) Pack(data interface{}) {
	this.fib.PACK(this.val, data)
	panic(nil)
	return
}

func (this *Status) Send(data interface{}) {
	this.fib.Send(this.val, data)
	panic(nil)
	return
}

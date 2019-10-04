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
	"github.com/abcum/orbit"
)

type Status struct {
	orb *orbit.Orbit
	val int
	req *Request
	res *Response
}

func NewStatus(orb *orbit.Orbit, req *Request, res *Response, val int) *Status {
	return &Status{
		orb: orb,
		val: val,
		req: req,
		res: res,
	}
}

func (this *Status) Done() {
	this.res.res.Code(this.val)
	panic(nil)
	return
}

func (this *Status) Xml(data interface{}) {
	this.res.res.Xml(this.val, data)
	panic(nil)
	return
}

func (this *Status) Text(data interface{}) {
	this.res.res.Text(this.val, data)
	panic(nil)
	return
}

func (this *Status) Html(data interface{}) {
	this.res.res.Html(this.val, data)
	panic(nil)
	return
}

func (this *Status) Json(data interface{}) {
	this.res.res.Json(this.val, data)
	panic(nil)
	return
}

func (this *Status) Cbor(data interface{}) {
	this.res.res.Cbor(this.val, data)
	panic(nil)
	return
}

func (this *Status) Pack(data interface{}) {
	this.res.res.Pack(this.val, data)
	panic(nil)
	return
}

func (this *Status) Head(key, val string) *Status {
	this.res.res.Head(key, val)
	return this
}

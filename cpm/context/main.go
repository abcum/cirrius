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
	"fmt"

	"github.com/abcum/orbit"
	"github.com/aymerick/raymond"
)

func init() {

	orbit.OnInit(func(orb *orbit.Orbit) {
		orb.Def("context", New(orb))
	})

}

func New(orb *orbit.Orbit) interface{} {
	return &Module{
		orb: orb,
		Req: NewRequest(orb),
		Res: NewResponse(orb),
	}
}

type Module struct {
	orb *orbit.Orbit

	// Req represents the http request of the
	// function runtime, enabling processing
	// the http request details and data.
	Req *Request `console:"req"`

	// Res represents the http response of the
	// function runtime, enabling sending data
	// back to the requesting http client.
	Res *Response `console:"res"`
}

// Success enables sending only an http 200
// success status code to the http client.
func (this *Module) Success() {
	this.Res.res.Code(200)
	panic(nil)
	return
}

// Failure enables sending only an http 400
// failure status code to the http client.
func (this *Module) Failure() {
	this.Res.res.Code(400)
	panic(nil)
	return
}

// Head enables defining the http headers
// which are to be sent with the response.
func (this *Module) Head(key, val string) {
	this.Res.res.Head(key, val)
}

// Status enables defining the http status
// code, with a chained method to send the
// data using different encoding types.
func (this *Module) Status(code int) *Status {
	return NewStatus(this.orb, this.Req, this.Res, code)
}

// The following methods enable sending the
// response data using different encoding
// methods, with an http status code 200.

func (this *Module) Xml(data interface{}) {
	this.Res.res.Xml(200, data)
	panic(nil)
	return
}

func (this *Module) Text(data interface{}) {
	this.Res.res.Text(200, data)
	panic(nil)
	return
}

func (this *Module) Html(data interface{}) {
	this.Res.res.Html(200, data)
	panic(nil)
	return
}

func (this *Module) Json(data interface{}) {
	this.Res.res.Json(200, data)
	panic(nil)
	return
}

func (this *Module) Cbor(data interface{}) {
	this.Res.res.Cbor(200, data)
	panic(nil)
	return
}

func (this *Module) Pack(data interface{}) {
	this.Res.res.Pack(200, data)
	panic(nil)
	return
}

// Render compiles and outputs a handlebars
// template, and additional partial templates
// included within the main template file.

func (this *Module) Render(file string, args ...interface{}) {

	var vars interface{}

	if len(args) > 0 {
		vars = args[0]
	}

	data, _, err := this.orb.File(file, "hbs")
	if err != nil {
		err := fmt.Errorf("Unable to find handlebars template '%s'.", file)
		this.orb.Quit(err)
	}

	code := fmt.Sprintf("%s", data)

	tmpl, err := raymond.Parse(code)
	if err != nil {
		err := fmt.Errorf("Unable to parse handlebars template '%s'.", file)
		this.orb.Quit(err)
	}

	priv := raymond.NewDataFrame()
	priv.Set("orb", this.orb)

	html, err := tmpl.ExecWith(vars, priv)
	if err != nil {
		err := fmt.Errorf("Unable to render handlebars template '%s'.", file)
		this.orb.Quit(err)
	}

	this.Res.res.Html(200, html)
	panic(nil)
	return

}

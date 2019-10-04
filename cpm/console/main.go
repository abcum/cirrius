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

package console

import (
	"log"
	"path"
	"time"

	"reflect"

	"unicode"
	"unicode/utf8"

	"github.com/abcum/orbit"
	"github.com/fatih/structs"
	"github.com/robertkrimen/otto"
)

type entry struct {
	kind string
	fold string
	file string
	line int
	char int
	args []interface{}
}

func init() {

	orbit.OnInit(func(orb *orbit.Orbit) {
		orb.Def("console", New(orb))
	})

}

func New(orb *orbit.Orbit) interface{} {
	return &Module{
		orb: orb,
		stw: make(map[string]time.Time),
	}
}

type Module struct {
	orb *orbit.Orbit
	stw map[string]time.Time
}

func (this *Module) Log(call otto.FunctionCall) otto.Value {
	this.output("log", this.input(call.ArgumentList)...)
	return otto.UndefinedValue()
}

func (this *Module) Info(call otto.FunctionCall) otto.Value {
	this.output("info", this.input(call.ArgumentList)...)
	return otto.UndefinedValue()
}

func (this *Module) Warn(call otto.FunctionCall) otto.Value {
	this.output("warn", this.input(call.ArgumentList)...)
	return otto.UndefinedValue()
}

func (this *Module) Error(call otto.FunctionCall) otto.Value {
	this.output("error", this.input(call.ArgumentList)...)
	return otto.UndefinedValue()
}

func (this *Module) Debug(call otto.FunctionCall) otto.Value {
	this.output("debug", this.input(call.ArgumentList)...)
	return otto.UndefinedValue()
}

func (this *Module) Trace(call otto.FunctionCall) otto.Value {
	this.output("trace", this.input(call.ArgumentList)...)
	return otto.UndefinedValue()
}

func (this *Module) Time(call otto.FunctionCall) otto.Value {
	this.stw[call.Argument(0).String()] = time.Now()
	return otto.UndefinedValue()
}

func (this *Module) Stop(call otto.FunctionCall) otto.Value {
	name := call.Argument(0).String()
	amnt := time.Since(this.stw[name])
	delete(this.stw, name)
	this.output("time", name, amnt)
	return otto.UndefinedValue()
}

func (this *Module) TimeEnd(call otto.FunctionCall) otto.Value {
	name := call.Argument(0).String()
	amnt := time.Since(this.stw[name])
	delete(this.stw, name)
	this.output("time", name, amnt)
	return otto.UndefinedValue()
}

func (this *Module) camel(s string) string {
	if s == "" {
		return ""
	}
	r, n := utf8.DecodeRuneInString(s)
	return string(unicode.ToLower(r)) + s[n:]
}

func (this *Module) input(args []otto.Value) (vars []interface{}) {
	for _, v := range args {
		switch {
		default:
			vars = append(vars, v.String())
		case v.IsFunction():
			vars = append(vars, v)
		case v.IsNull():
			vars = append(vars, "null")
		case v.IsUndefined():
			vars = append(vars, "undefined")
		case v.IsNumber():
			vars = append(vars, v.String())
		case v.IsString():
			vars = append(vars, v.String())
		case v.IsBoolean():
			vars = append(vars, v.String())
		case v.IsObject():
			o, _ := v.Export()
			p := reflect.ValueOf(o)
			r := reflect.ValueOf(o)
			for r.Kind() == reflect.Ptr {
				r = r.Elem()
			}
			if r.Kind() != reflect.Struct {
				vars = append(vars, o)
			} else {
				if d, ok := p.Interface().(Debuggable); ok {
					vars = append(vars, d.Debug())
				} else {
					structs.DefaultTagName = "console"
					strt := structs.Map(p.Interface())
					for i := 0; i < p.NumMethod(); i++ {
						m := p.Type().Method(i)
						strt[this.camel(m.Name)] = "function()"
					}
					vars = append(vars, strt)
				}
			}
		}
	}
	return
}

func (this *Module) output(kind string, args ...interface{}) {

	fold, file := path.Split(this.orb.Otto.Context().Filename)
	line := this.orb.Otto.Context().Line
	char := this.orb.Otto.Context().Column

	data := entry{kind, fold, file, line, char, args}

	log.Printf("console.%s: %v in %s%s at %d:%d\n", data.kind, data.args, data.fold, data.file, data.line, data.char)

}

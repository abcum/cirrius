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
	"github.com/robertkrimen/otto"
	"github.com/surrealdb/surreal.go"

	"github.com/abcum/cirrius/util/args"
)

type DB struct {
	orb *orbit.Orbit
	val *surreal.DB
	url string
	opt map[string]interface{}
}

func NewDB(orb *orbit.Orbit, url string, opt map[string]interface{}) *DB {
	return (&DB{
		orb: orb,
		url: url,
		opt: opt,
	}).init()
}

func (this *DB) init() *DB {

	var err error

	cnf := &surreal.Config{}

	if this.opt != nil {

		if val, ok := this.opt["ns"].(string); ok {
			cnf.NS = val
		}

		if val, ok := this.opt["db"].(string); ok {
			cnf.DB = val
		}

		if val, ok := this.opt["user"].(string); ok {
			cnf.User = val
		}

		if val, ok := this.opt["pass"].(string); ok {
			cnf.Pass = val
		}

	}

	this.val, err = surreal.New(this.url, cnf)
	if err != nil {
		this.orb.Quit(err)
	}

	this.orb.Once("exit", func() {
		this.val.Close()
	})

	return this

}

func (this *DB) Close() {
	this.val.Close()
}

func (this *DB) Query(call otto.FunctionCall) otto.Value {

	args.Size(this.orb, call, 1, 2)

	sql := args.String(this.orb, call, 0)
	opt := args.Object(this.orb, call, 1)

	res, err := this.val.Query(sql, opt)
	if err != nil {
		this.orb.Quit(err)
	}

	return args.Value(this.orb, res)

}

func (this *DB) Select(call otto.FunctionCall) otto.Value {

	args.Size(this.orb, call, 1, 2)

	tb := args.String(this.orb, call, 0)
	id := args.Any(this.orb, call, 1)

	res, err := this.val.Select(tb, id)
	if err != nil {
		this.orb.Quit(err)
	}

	return args.Value(this.orb, res)

}

func (this *DB) Create(call otto.FunctionCall) otto.Value {

	args.Size(this.orb, call, 1, 2)

	tb := args.String(this.orb, call, 0)
	id := args.Any(this.orb, call, 1)
	ob := args.Object(this.orb, call, 2)

	res, err := this.val.Create(tb, id, ob)
	if err != nil {
		this.orb.Quit(err)
	}

	return args.Value(this.orb, res)

}

func (this *DB) Update(call otto.FunctionCall) otto.Value {

	args.Size(this.orb, call, 1, 2)

	tb := args.String(this.orb, call, 0)
	id := args.Any(this.orb, call, 1)
	ob := args.Object(this.orb, call, 2)

	res, err := this.val.Update(tb, id, ob)
	if err != nil {
		this.orb.Quit(err)
	}

	return args.Value(this.orb, res)

}

func (this *DB) Change(call otto.FunctionCall) otto.Value {

	args.Size(this.orb, call, 1, 2)

	tb := args.String(this.orb, call, 0)
	id := args.Any(this.orb, call, 1)
	ob := args.Object(this.orb, call, 2)

	res, err := this.val.Change(tb, id, ob)
	if err != nil {
		this.orb.Quit(err)
	}

	return args.Value(this.orb, res)

}

func (this *DB) Modify(call otto.FunctionCall) otto.Value {

	args.Size(this.orb, call, 1, 2)

	tb := args.String(this.orb, call, 0)
	id := args.Any(this.orb, call, 1)
	ob := args.Object(this.orb, call, 2)

	res, err := this.val.Modify(tb, id, ob)
	if err != nil {
		this.orb.Quit(err)
	}

	return args.Value(this.orb, res)

}

func (this *DB) Delete(call otto.FunctionCall) otto.Value {

	args.Size(this.orb, call, 1, 2)

	tb := args.String(this.orb, call, 0)
	id := args.Any(this.orb, call, 1)

	res, err := this.val.Delete(tb, id)
	if err != nil {
		this.orb.Quit(err)
	}

	return args.Value(this.orb, res)

}

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

package web

import (
	"fmt"

	"github.com/abcum/fibre"
	"github.com/abcum/orbit"
	"github.com/robertkrimen/otto"

	// Load general
	_ "github.com/abcum/cirrius/npm"

	// Load globals
	_ "github.com/abcum/cirrius/npm/globals/console"
	_ "github.com/abcum/cirrius/npm/globals/context"
	_ "github.com/abcum/cirrius/npm/globals/process"

	// Load modules
	_ "github.com/abcum/cirrius/npm/modules/binc"
	_ "github.com/abcum/cirrius/npm/modules/cbor"
	_ "github.com/abcum/cirrius/npm/modules/colour"
	_ "github.com/abcum/cirrius/npm/modules/dns"
	_ "github.com/abcum/cirrius/npm/modules/http"
	_ "github.com/abcum/cirrius/npm/modules/https"
	_ "github.com/abcum/cirrius/npm/modules/json"
	_ "github.com/abcum/cirrius/npm/modules/msgpack"
	_ "github.com/abcum/cirrius/npm/modules/os"
	_ "github.com/abcum/cirrius/npm/modules/path"
	_ "github.com/abcum/cirrius/npm/modules/routes"
	_ "github.com/abcum/cirrius/npm/modules/shortid"
	_ "github.com/abcum/cirrius/npm/modules/uuid"
	_ "github.com/abcum/cirrius/npm/modules/ws"
)

func init() {

	orbit.OnInit(func(ctx *orbit.Orbit) {

		session := ctx.Vars["fibre"].(*fibre.Context)

		fmt.Println("INIT", session.Response().Header().Get("Request-Id"))

	})

	orbit.OnExit(func(ctx *orbit.Orbit) {

		session := ctx.Vars["fibre"].(*fibre.Context)

		fmt.Println("EXIT", session.Response().Header().Get("Request-Id"))

	})

	orbit.OnFail(func(ctx *orbit.Orbit, err error) {
		if tmp, ok := err.(*otto.Error); ok {
			fmt.Println("FAIL", tmp.String())
		} else {
			fmt.Println("FAIL", err.Error())
		}
		// Log runtime error on context
	})

	orbit.OnFile(func(ctx *orbit.Orbit, files []string) (code interface{}, file string, err error) {
		info, err := find(files...) // TODO Need to use fibre request context here
		if err != nil {
			return code, file, err
		}
		return info.data, info.path, err
	})

}

func processNode(ctx *fibre.Context, info *info) (err error) {

	// New orbit instance
	orb := orbit.New(0)

	// Set fibre instance
	orb.Vars["fibre"] = ctx

	// Run orbit instance
	return orb.Exec(info.path, info.data)

}

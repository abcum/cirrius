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
	"github.com/abcum/fibre"
	"github.com/abcum/orbit"

	// Load globals
	_ "github.com/abcum/cirrius/npm/globals/console"
	_ "github.com/abcum/cirrius/npm/globals/context"
	_ "github.com/abcum/cirrius/npm/globals/process"

	// Load modules
	_ "github.com/abcum/cirrius/npm/modules/accounting"
	_ "github.com/abcum/cirrius/npm/modules/async"
	_ "github.com/abcum/cirrius/npm/modules/babel"
	_ "github.com/abcum/cirrius/npm/modules/bluebird"
	_ "github.com/abcum/cirrius/npm/modules/chance"
	_ "github.com/abcum/cirrius/npm/modules/classnames"
	_ "github.com/abcum/cirrius/npm/modules/co"
	_ "github.com/abcum/cirrius/npm/modules/cookie"
	_ "github.com/abcum/cirrius/npm/modules/dateformat"
	_ "github.com/abcum/cirrius/npm/modules/deep-extend"
	_ "github.com/abcum/cirrius/npm/modules/deepmerge"
	_ "github.com/abcum/cirrius/npm/modules/dns"
	_ "github.com/abcum/cirrius/npm/modules/events"
	_ "github.com/abcum/cirrius/npm/modules/extend"
	_ "github.com/abcum/cirrius/npm/modules/hashids"
	_ "github.com/abcum/cirrius/npm/modules/http"
	_ "github.com/abcum/cirrius/npm/modules/https"
	_ "github.com/abcum/cirrius/npm/modules/immutable"
	_ "github.com/abcum/cirrius/npm/modules/jsondiffpatch"
	_ "github.com/abcum/cirrius/npm/modules/jsonic"
	_ "github.com/abcum/cirrius/npm/modules/jsonwebtoken"
	_ "github.com/abcum/cirrius/npm/modules/lodash"
	_ "github.com/abcum/cirrius/npm/modules/merge"
	_ "github.com/abcum/cirrius/npm/modules/moment"
	_ "github.com/abcum/cirrius/npm/modules/ms"
	_ "github.com/abcum/cirrius/npm/modules/odiff"
	_ "github.com/abcum/cirrius/npm/modules/once"
	_ "github.com/abcum/cirrius/npm/modules/os"
	_ "github.com/abcum/cirrius/npm/modules/path"
	_ "github.com/abcum/cirrius/npm/modules/pdfkit"
	_ "github.com/abcum/cirrius/npm/modules/promise"
	_ "github.com/abcum/cirrius/npm/modules/promiz"
	_ "github.com/abcum/cirrius/npm/modules/q"
	_ "github.com/abcum/cirrius/npm/modules/qs"
	_ "github.com/abcum/cirrius/npm/modules/react"
	_ "github.com/abcum/cirrius/npm/modules/routes"
	_ "github.com/abcum/cirrius/npm/modules/rx"
	_ "github.com/abcum/cirrius/npm/modules/semver"
	_ "github.com/abcum/cirrius/npm/modules/shortid"
	_ "github.com/abcum/cirrius/npm/modules/string"
	_ "github.com/abcum/cirrius/npm/modules/sugar"
	_ "github.com/abcum/cirrius/npm/modules/traverse"
	_ "github.com/abcum/cirrius/npm/modules/underscore"
	_ "github.com/abcum/cirrius/npm/modules/uuid"
	_ "github.com/abcum/cirrius/npm/modules/validator"
	_ "github.com/abcum/cirrius/npm/modules/when"
	_ "github.com/abcum/cirrius/npm/modules/wrappy"
	_ "github.com/abcum/cirrius/npm/modules/xtend"
)

func init() {

	orbit.OnInit(func(ctx *orbit.Orbit) {
		// Create unique runtime id for logs and events
	})

	orbit.OnExit(func(ctx *orbit.Orbit) {
		// Finish runtime id
	})

	orbit.OnFail(func(ctx *orbit.Orbit, err error) {
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

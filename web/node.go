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
	_ "github.com/abcum/cirrius/npm/globals/output"
	_ "github.com/abcum/cirrius/npm/globals/process"
	_ "github.com/abcum/cirrius/npm/globals/promise"

	// Load modules
	_ "github.com/abcum/cirrius/npm/modules/accounting"
	_ "github.com/abcum/cirrius/npm/modules/async"
	_ "github.com/abcum/cirrius/npm/modules/bluebird"
	_ "github.com/abcum/cirrius/npm/modules/chance"
	_ "github.com/abcum/cirrius/npm/modules/deepmerge"
	_ "github.com/abcum/cirrius/npm/modules/hashids"
	_ "github.com/abcum/cirrius/npm/modules/immutable"
	_ "github.com/abcum/cirrius/npm/modules/jsondiffpatch"
	_ "github.com/abcum/cirrius/npm/modules/jsonwebtoken"
	_ "github.com/abcum/cirrius/npm/modules/lodash"
	_ "github.com/abcum/cirrius/npm/modules/merge"
	_ "github.com/abcum/cirrius/npm/modules/moment"
	_ "github.com/abcum/cirrius/npm/modules/odiff"
	_ "github.com/abcum/cirrius/npm/modules/promiz"
	_ "github.com/abcum/cirrius/npm/modules/q"
	_ "github.com/abcum/cirrius/npm/modules/shortid"
	_ "github.com/abcum/cirrius/npm/modules/sugar"
	_ "github.com/abcum/cirrius/npm/modules/underscore"
	_ "github.com/abcum/cirrius/npm/modules/uuid"
	_ "github.com/abcum/cirrius/npm/modules/validator"
)

func init() {

	orbit.OnInit(func(ctx *orbit.Orbit) {
		// Create unique runtime id for logs and events
	})

	orbit.OnExit(func(ctx *orbit.Orbit) {
		// Finish runtime id
	})

	orbit.OnFile(func(ctx *orbit.Orbit, files []string) (code interface{}, file string, err error) {
		info, err := find(files...) // TODO Need to use fibre request context here
		if err != nil {
			return code, file, err
		}
		return info.data, info.fold, err
	})

}

func processNode(c *fibre.Context, input []byte) (val interface{}, err error) {

	ret, err := orbit.Run(input) // TODO Need to pass in fibre request context here
	if err != nil {
		return nil, err
	}

	return ret.Export()

}

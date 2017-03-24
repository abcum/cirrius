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
	"context"

	"github.com/abcum/fibre"
	"github.com/abcum/orbit"

	_ "github.com/abcum/cirrius/cpm"
)

func init() {

	orbit.OnFile(func(orb *orbit.Orbit, files []string) (code interface{}, file string, err error) {
		info, err := find(files...) // TODO Need to use fibre request context here
		if err != nil {
			return code, file, err
		}
		return info.data, info.path, err
	})

}

func processJS(fib *fibre.Context, info *info) (err error) {

	// New orbit instance
	orb := orbit.New(0)

	// New orbit context
	ctx := context.WithValue(orb.Context(), "fibre", fib)

	// Set orbit context
	orb = orb.WithContext(ctx)

	// Run orbit instance
	return orb.Exec(info.path, info.data)

}

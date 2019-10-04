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

package aws

import (
	"regexp"
	"sort"

	"github.com/abcum/cirrius/cnf"
	"github.com/abcum/cirrius/exe"
)

func headers(req *Req, res *Res, file *exe.File) {

	res.Head("X-Powered-By", "Cirrius")

	if _, ok := res.res.Headers["Content-Type"]; !ok && file.Path != "main.js" {
		res.Head("Content-Type", file.Mime)
	}

	if _, ok := res.res.Headers["Cache-Control"]; !ok {
		res.Head("Cache-Control", "private, no-store, no-cache, must-revalidate, proxy-revalidate")
	}

	if typ, ok := res.res.Headers["Content-Type"]; ok {

		var keys sort.StringSlice

		for k := range cnf.Settings.Cache {
			keys = append(keys, k)
		}

		sort.Sort(sort.Reverse(keys))

		for _, v := range keys {
			if reg, err := regexp.Compile(v); err == nil {
				if reg.MatchString(typ) {
					res.Head("Cache-Control", cnf.Settings.Cache[v])
					return
				}
			}
		}

	}

}

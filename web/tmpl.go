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
	"regexp"
)

func tmpl(input []byte, filetype string) []byte {

	if filetype == "ssi" {

		regex := regexp.MustCompile(`<!--#template file=['|"]([^'|"]+)['|"] -->`)

		index := regex.FindSubmatchIndex(input)

		if index != nil {

			match := input[index[2]:index[3]]

			found, _ := find(string(match))

			if found != nil {

				found.data = tmpl(found.data, filetype)

				input = append(input[:index[0]], input[index[1]:]...)

				input = regexp.MustCompile(`<!--#content -->`).ReplaceAll(found.data, input)

			}

		}

	}

	return input

}

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

func incl(input []byte, filetype string) []byte {

	if filetype == "css" {

		regex := regexp.MustCompile(`@import (url\()?['|"]([^'|"]+)['|"](\))?`)

		input = regex.ReplaceAllFunc(input, func(found []byte) []byte {
			info, err := find(string(regex.FindSubmatch(found)[2]))
			if err != nil {
				return nil
			}
			return incl(info.data, filetype)
		})

	}

	if filetype == "ssi" {

		regex := regexp.MustCompile(`<!--#include file=['|"]([^'|"]+)['|"] -->`)

		input = regex.ReplaceAllFunc(input, func(found []byte) []byte {
			info, err := find(string(regex.FindSubmatch(found)[1]))
			if err != nil {
				return nil
			}
			return incl(info.data, filetype)
		})

	}

	return input

}

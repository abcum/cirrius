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

package exe

import (
	"context"
	"fmt"
)

func (e *Executable) incl(ctx context.Context, file *File, filetype string) *File {

	if filetype == "css" {

		file.Data = regexCSS.ReplaceAllFunc(file.Data, func(found []byte) []byte {
			file, err := e.find(ctx, file, string(regexCSS.FindSubmatch(found)[2]))
			if err != nil {
				return nil
			}
			return e.incl(ctx, file, filetype).Data
		})

	}

	if filetype == "ssi" {

		file.Data = regexSSI.ReplaceAllFunc(file.Data, func(found []byte) []byte {
			file, err := e.find(ctx, file, string(regexSSI.FindSubmatch(found)[1]))
			if err != nil {
				return nil
			}
			return e.incl(ctx, file, filetype).Data
		})

	}

	return file

}

func (e *Executable) tmpl(ctx context.Context, file *File, filetype string) *File {

	if filetype == "ssi" {

		idx := regexTPL.FindSubmatchIndex(file.Data)

		if idx != nil {

			path := string(file.Data[idx[2]:idx[3]])

			tmpl, _ := e.find(ctx, file, path)

			if tmpl != nil {

				tmpl = e.tmpl(ctx, tmpl, filetype)

				file.Data = regexHTM.ReplaceAll(tmpl.Data, file.Data)

			}

		}

	}

	return file

}

func (e *Executable) bust(ctx context.Context, file *File, filetype string) *File {

	file.Data = regexIMG.ReplaceAllFunc(file.Data, func(found []byte) []byte {

		mtch := regexIMG.FindSubmatch(found)[1]

		path := string(mtch)

		file, err := e.find(ctx, file, path)
		if err != nil {
			return mtch
		}

		bust := fmt.Sprintf("%s-v%d%s",
			path[:len(path)-len(file.Extn)],
			file.Time.Unix(),
			file.Extn,
		)

		return []byte(bust)

	})

	return file

}

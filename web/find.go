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
	"io/ioutil"
	"mime"
	"os"
	"path"

	"github.com/abcum/fibre"
)

type info struct {
	path string
	fold string
	file string
	extn string
	mime string
	data []byte
}

func load(url string) (*info, error) {

	var files []string

	// Clean the url
	url = path.Clean(url)

	// Check for the exact file name
	files = append(files, url)

	// Check for the default file extensions in a folder
	if path.Base(url) == "/" || path.Ext(url) == "" {
		files = append(files, path.Join(url, "index.html"))
	}

	// Check for the file with default extensions
	if path.Base(url) != "/" && path.Ext(url) == "" {
		files = append(files, url+".html")
	}

	// Check for a single page app
	if path.Ext(url) == "" {
		files = append(files, "index.html")
	}

	// Check for a main js file
	files = append(files, "main.js")

	// Check for a custom 404 file
	files = append(files, "404.html")

	return find(files...)

}

func find(files ...string) (*info, error) {

	for _, file := range files {

		full := path.Join("dev", file)

		if hand, err := os.Stat(full); err == nil {

			if !hand.IsDir() && path.Base(full)[0] != '.' {

				info := &info{}
				info.path = path.Clean(file)
				info.fold = path.Dir(info.path)
				info.file = path.Base(info.path)
				info.extn = path.Ext(info.path)
				info.mime = mime.TypeByExtension(info.extn)
				info.data, _ = ioutil.ReadFile(full)

				return info, err

			}

		}

	}

	return nil, fibre.NewHTTPError(404)

}

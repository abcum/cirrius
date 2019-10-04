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
	"io/ioutil"
	"mime"
	"os"
	"path/filepath"
	"regexp"

	"github.com/abcum/fibre"

	"github.com/abcum/cirrius/exe"
)

var cacher = regexp.MustCompile("-v[0-9]+")

func find(ctx context.Context, prev *exe.File, url string) (file *exe.File, err error) {

	var paths []string

	// Clean the url
	url = filepath.Clean(url)

	// Check for the exact file name
	if filepath.Base(url) != "." {
		paths = append(paths, url)
	}

	// Check for the file with default extensions
	if filepath.Base(url) != "." && filepath.Ext(url) == "" {
		paths = append(paths, url+".html")
	}

	// Check for the default file extensions in a folder
	if filepath.Base(url) == "." || filepath.Ext(url) == "" {
		paths = append(paths, filepath.Join(url, "index.html"))
	}

	// Check for a main js file
	paths = append(paths, "main.js")

	// Check for a custom 404 file
	paths = append(paths, "404.html")

	// Fetch the files
	return load(ctx, prev, paths...)

}

func load(ctx context.Context, prev *exe.File, paths ...string) (file *exe.File, err error) {

	for _, path := range paths {

		switch {
		case path[0] == '/':
			path = filepath.Join(".", path)
		case prev == nil:
			path = filepath.Join(".", path)
		case prev != nil:
			path = filepath.Join(prev.Fold, path)
		}

		if cacher.MatchString(path) {
			path = cacher.ReplaceAllString(path, "")
		}

		if hand, err := os.Stat(path); err == nil {

			// Ignore if the path is a folder

			if hand.IsDir() {
				continue
			}

			// Ignore if the path is invisible

			if filepath.Base(path)[0] == '.' {
				continue
			}

			// Otherwise load the file

			file = &exe.File{}
			file.Path = filepath.Clean(path)
			file.Fold = filepath.Dir(file.Path)
			file.File = filepath.Base(file.Path)
			file.Extn = filepath.Ext(file.Path)
			file.Name = file.File[:len(file.File)-len(file.Extn)]
			file.Mime = mime.TypeByExtension(file.Extn)
			file.Time = hand.ModTime()
			file.Data, _ = ioutil.ReadFile(path)

			return file, nil

		}

	}

	return nil, fibre.NewHTTPError(404)

}

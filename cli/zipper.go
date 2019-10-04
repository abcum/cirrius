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

package cli

import (
	"archive/zip"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func saver(name string) error {

	path := "https://pkg.cirrius.io/" + name

	// Create the file

	file, err := os.Create("main")
	if err != nil {
		return err
	}

	defer file.Close()

	// Download the file data

	resp, err := http.Get(path)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	// Copy the file data to the file

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	return os.Chmod("main", 0755)

}

func zipper(name string) error {

	// Create the zip file

	file, err := os.Create(name)
	if err != nil {
		return err
	}

	defer file.Close()

	// Setup the archive format

	arch := zip.NewWriter(file)

	defer arch.Close()

	// Walk through the folder files

	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		if path == name {
			return nil
		}

		if path[0] == '.' {
			return nil
		}

		if info.Name()[0] == '.' {
			return nil
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Name = path
			header.Method = zip.Deflate
		}

		writer, err := arch.CreateHeader(header)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		temp, err := os.Open(path)
		if err != nil {
			return err
		}

		_, err = io.Copy(writer, temp)

		temp.Close()

		return err

	})

	return err

}

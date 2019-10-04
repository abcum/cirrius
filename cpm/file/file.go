// Copyright © 2016 Abcum Ltd
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this info except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package file

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/abcum/orbit"
	"github.com/spf13/afero"

	"github.com/abcum/cirrius/util/uuid"
)

type File struct {
	afero.File
	orb *orbit.Orbit
}

func NewTemp(orb *orbit.Orbit) *File {
	fs := orb.Context().Value("mfs").(afero.Fs)
	fil, err := fs.OpenFile(uuid.New().String(), os.O_RDWR|os.O_CREATE, 0744)
	if err != nil {
		orb.Quit(err)
	}
	return (&File{orb: orb, File: fil}).init()
}

func NewFile(orb *orbit.Orbit, name string) *File {
	fs := orb.Context().Value("mfs").(afero.Fs)
	fil, err := fs.OpenFile(name, os.O_RDWR|os.O_CREATE, 0744)
	if err != nil {
		orb.Quit(err)
	}
	return (&File{orb: orb, File: fil}).init()
}

func (this *File) init() *File {
	this.orb.Once("exit", func() {
		this.Close()
	})
	return this
}

func (this *File) End() {
	this.Close()
}

func (this *File) Close() error {
	return this.File.Close()
}

func (this *File) Pipe(w io.Writer) {
	io.Copy(w, this.File)
}

func (this *File) Consume(r io.Reader) {
	io.Copy(this.File, r)
}

func (this *File) String() string {
	this.File.Seek(0, 0)
	res, err := ioutil.ReadAll(this)
	if err != nil {
		this.orb.Quit(err)
	}
	return string(res)
}

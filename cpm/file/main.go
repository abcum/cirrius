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
	"errors"
	"os"
	"path"

	"github.com/abcum/fibre"
	"github.com/abcum/orbit"

	"github.com/abcum/cirrius/cpm/stream"
)

func init() {
	orbit.Add("file", New)
}

var pathErr = errors.New("Path not allowed")

func New(orb *orbit.Orbit) interface{} {
	return (&Module{
		orb: orb,
	}).init()
}

type Module struct {
	orb *orbit.Orbit
	dir string
}

func (this *Module) init() *Module {

	this.orb.Once("exit", func() {
		os.RemoveAll(this.dir)
	})

	fib := this.orb.Context().Value("fibre").(*fibre.Context)

	dir := path.Join(os.TempDir(), fib.Get("id").(string))

	os.Mkdir(dir, 0744)

	this.dir = dir

	return this

}

func (this *Module) Read(file string) *stream.ReadCloser {

	file = path.Join(this.dir, path.Clean(file))

	if ok, _ := path.Match(path.Join(this.dir, "*"), file); !ok {
		this.orb.Quit(pathErr)
	}

	stm, err := os.OpenFile(file, os.O_RDONLY, 0)
	if err != nil {
		this.orb.Quit(err)
	}

	return stream.NewReadCloser(this.orb, stm)

}

func (this *Module) Write(file string) *stream.WriteCloser {

	file = path.Join(this.dir, path.Clean(file))

	if ok, _ := path.Match(path.Join(this.dir, "*"), file); !ok {
		this.orb.Quit(pathErr)
	}

	stm, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0744)
	if err != nil {
		this.orb.Quit(err)
	}

	return stream.NewWriteCloser(this.orb, stm)

}

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
	"os"

	"github.com/abcum/orbit"
	"github.com/spf13/afero"
)

func init() {
	orbit.Add("file", New)
}

var fs = afero.NewOsFs()

func New(orb *orbit.Orbit) interface{} {
	return (&Module{
		orb: orb,
	}).init()
}

type Module struct {
	orb *orbit.Orbit
}

func (this *Module) init() *Module {
	return this
}

func (this *Module) Temp() *File {
	return NewTemp(this.orb)
}

func (this *Module) Open(file string) *File {
	return NewFile(this.orb, file)
}

func (this *Module) Load(file string) *File {
	fil, err := fs.OpenFile(file, os.O_RDONLY, 0)
	if err != nil {
		this.orb.Quit(err)
	}
	return (&File{orb: this.orb, File: fil}).init()
}

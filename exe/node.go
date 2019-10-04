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

// +build !web

package exe

import (
	"context"

	"github.com/abcum/orbit"
	"github.com/spf13/afero"

	_ "github.com/abcum/cirrius/cpm"
)

func init() {

	orbit.OnFile(func(orb *orbit.Orbit, files []string) (interface{}, string, error) {

		e := orb.Context().Value("exe").(*Executable)

		file, err := e.find(orb.Context(), nil, files...)
		if err != nil {
			return nil, "", err
		}

		return file.Data, file.Path, err

	})

}

func (e *Executable) LoadJS(ctx context.Context, file *File) (err error) {

	// Register the filesystem
	ctx = context.WithValue(ctx, "mfs", afero.NewMemMapFs())

	// Register the executable
	ctx = context.WithValue(ctx, "exe", e)

	// New orbit instance
	orb := orbit.New(0)

	// Set orbit context
	orb = orb.WithContext(ctx)

	// Run orbit instance
	return orb.Exec(file.Path, file.Data)

}

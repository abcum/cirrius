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
	"bytes"
	"context"

	"github.com/yosssi/gcss"
)

func (e *Executable) LoadGCSS(ctx context.Context, file *File) (_ *File, err error) {

	file = e.incl(ctx, file, "css")

	file = e.bust(ctx, file, "css")

	w := bytes.NewBuffer(nil)
	r := bytes.NewReader(file.Data)
	_, err = gcss.Compile(w, r)
	file.Data = w.Bytes()

	return file, err

}

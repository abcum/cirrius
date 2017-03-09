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

package html

import (
	"bytes"
	"io"

	"github.com/abcum/orbit"

	"github.com/abcum/cirrius/cpm/minify"
	"github.com/abcum/cirrius/cpm/stream"
)

var mime = "text/html"

func init() {
	orbit.Add("minify/html", New)
}

func New(orb *orbit.Orbit) interface{} {
	return &Module{
		orb: orb,
	}
}

type Module struct {
	orb *orbit.Orbit
}

func (this *Module) Reader(r io.Reader) *stream.Reader {
	return stream.NewReader(this.orb, minify.Minifier.Reader(mime, r))
}

func (this *Module) Writer(w io.Writer) *stream.WriteCloser {
	return stream.NewWriteCloser(this.orb, minify.Minifier.Writer(mime, w))
}

func (this *Module) Minify(src interface{}) (out string) {
	switch v := src.(type) {
	case string:
		out, err := minify.Minifier.String(mime, v)
		if err != nil {
			this.orb.Quit(err)
		}
		return out
	case io.Reader:
		var buf bytes.Buffer
		err := minify.Minifier.Minify(mime, &buf, v)
		if err != nil {
			this.orb.Quit(err)
		}
		return buf.String()
	}
	return
}

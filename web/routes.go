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

package web

import (
	"mime"

	"context"

	"github.com/abcum/cirrius/exe"

	"github.com/abcum/fibre"
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/css"
	"github.com/tdewolff/minify/html"
	"github.com/tdewolff/minify/js"
	"github.com/tdewolff/minify/json"
	"github.com/tdewolff/minify/svg"
	"github.com/tdewolff/minify/xml"
)

func routes(s *fibre.Fibre) {

	mime.AddExtensionType(".md", "text/html")
	mime.AddExtensionType(".htm", "text/html")
	mime.AddExtensionType(".tpl", "text/html")
	mime.AddExtensionType(".hbs", "text/html")
	mime.AddExtensionType(".less", "text/css")
	mime.AddExtensionType(".sass", "text/css")
	mime.AddExtensionType(".scss", "text/css")
	mime.AddExtensionType(".gcss", "text/css")

	s.Any("/*", func(c *fibre.Context) (err error) {

		// --------------------------------------------------
		// Specify context
		// --------------------------------------------------

		ctx := c.Context()
		ctx = context.WithValue(ctx, "req", &Req{c})
		ctx = context.WithValue(ctx, "res", &Res{c})

		// --------------------------------------------------
		// Find the file
		// --------------------------------------------------

		file, err := find(ctx, nil, c.Request().URL().Path)
		if err != nil {
			return err
		}

		// --------------------------------------------------
		// Setup the runtime
		// --------------------------------------------------

		run := exe.NewExecutable(load)

		// --------------------------------------------------
		// Execute the file
		// --------------------------------------------------

		if file.Path == "main.js" {
			return run.LoadJS(ctx, file)
		}

		switch file.Extn {
		case ".css":
			file, err = run.LoadCSS(ctx, file)
		case ".less":
			file, err = run.LoadLESS(ctx, file)
		case ".scss":
			file, err = run.LoadSCSS(ctx, file)
		case ".html":
			file, err = run.LoadHTML(ctx, file)
		}

		if err != nil {
			return err
		}

		// --------------------------------------------------
		// Minify the output
		// --------------------------------------------------

		m := minify.New()
		m.AddFunc("text/css", css.Minify)
		m.AddFunc("text/html", html.Minify)
		m.AddFunc("text/javascript", js.Minify)
		m.AddFunc("image/svg+xml", svg.Minify)
		m.AddFunc("application/xml", xml.Minify)
		m.AddFunc("application/json", json.Minify)
		m.AddFunc("application/javascript", js.Minify)
		file.Data, _ = m.Bytes(file.Mime, file.Data)

		return c.Data(200, file.Data, file.Mime)

	})

}

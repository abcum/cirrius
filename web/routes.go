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

	"cloud.google.com/go/trace"

	"github.com/abcum/fibre"
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/css"
	"github.com/tdewolff/minify/html"
	"github.com/tdewolff/minify/js"
	"github.com/tdewolff/minify/json"
	"github.com/tdewolff/minify/svg"
	"github.com/tdewolff/minify/xml"

	"github.com/abcum/cirrius/util/build"
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

	// --------------------------------------------------
	// Endpoint for health checks
	// --------------------------------------------------

	s.Get("/info", func(c *fibre.Context) error {
		return c.Send(200, build.GetInfo())
	})

	// --------------------------------------------------
	// Endpoint for all other requests
	// --------------------------------------------------

	s.Any("/*", func(c *fibre.Context) (err error) {

		span := trace.FromContext(c.Context()).NewChild("route")
		nctx := trace.NewContext(c.Context(), span)
		c = c.WithContext(nctx)
		defer span.Finish()

		if err, ok := render(c); ok {
			return err
		}

		info, err := load(c.Request().URL().Path)
		if err != nil {
			return err
		}

		if info.path == "main.js" {
			return processJS(c, info)
		}

		switch info.extn {
		case ".md":
			info.data, err = processMD(info.data)
		case ".css":
			info.data, err = processCSS(info.data)
		case ".less":
			info.data, err = processLESS(info.data)
		case ".scss":
			info.data, err = processSCSS(info.data)
		case ".sass":
			info.data, err = processSASS(info.data)
		case ".gcss":
			info.data, err = processGCSS(info.data)
		case ".html":
			info.data, err = processHTML(info.data)
		}

		if err != nil {
			return err
		}

		switch info.extn {
		case ".md", ".tpl", ".hbs", ".htm", ".html":
			m := minify.New()
			m.AddFunc("text/css", css.Minify)
			m.AddFunc("text/html", html.Minify)
			m.AddFunc("text/javascript", js.Minify)
			m.AddFunc("image/svg+xml", svg.Minify)
			m.AddFunc("application/xml", xml.Minify)
			m.AddFunc("application/json", json.Minify)
			m.AddFunc("application/javascript", js.Minify)
			info.data, _ = m.Bytes(info.mime, info.data)
		}

		return c.Data(200, info.data, info.mime)

	})

}

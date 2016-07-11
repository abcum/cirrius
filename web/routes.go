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

	mime.AddExtensionType(".htm", "text/html")
	mime.AddExtensionType(".tpl", "text/html")
	mime.AddExtensionType(".hbs", "text/html")
	mime.AddExtensionType(".less", "text/css")
	mime.AddExtensionType(".sass", "text/css")
	mime.AddExtensionType(".scss", "text/css")
	mime.AddExtensionType(".gcss", "text/css")
	mime.AddExtensionType(".mock", "application/json")

	// --------------------------------------------------
	// Endpoint for health checks
	// --------------------------------------------------

	s.Get("/info", func(c *fibre.Context) error {
		return c.Code(200)
	})

	// --------------------------------------------------
	// Endpoint for all other requests
	// --------------------------------------------------

	s.Any("/*", func(c *fibre.Context) (err error) {

		// if c.Request().URL().Host == "127.0.0.1" {
		// 	return c.File("app", c.Param("*"))
		// }

		// if c.Request().URL().Host == "app.magnif.io" {
		// 	return c.File("app", c.Param("*"))
		// }

		// info, err := load("main.node")
		// if err != nil {
		// 	return err
		// }

		info, err := load(c.Request().URL().Path)
		if err != nil {
			return err
		}

		if info.extn == ".node" {

			_, err := processNode(c, info.path, info.data)
			return err

		} else {

			switch info.extn {
			case ".md":
				info.data, err = processMd(info.data)
			case ".css":
				info.data, err = processCss(info.data)
			case ".less":
				info.data, err = processLess(info.data)
			case ".scss":
				info.data, err = processScss(info.data)
			case ".sass":
				info.data, err = processSass(info.data)
			case ".gcss":
				info.data, err = processGcss(info.data)
			case ".mock":
				info.data, err = processMock(info.data)
			case ".html":
				info.data, err = processHtml(info.data)
			}

		}

		if err != nil {
			return err
		}

		m := minify.New()

		m.AddFunc("text/css", css.Minify)
		m.AddFunc("text/html", html.Minify)
		m.AddFunc("text/javascript", js.Minify)
		m.AddFunc("image/svg+xml", svg.Minify)
		m.AddFunc("application/xml", xml.Minify)
		m.AddFunc("application/json", json.Minify)
		m.AddFunc("application/javascript", js.Minify)

		info.data, _ = m.Bytes(info.mime, info.data)

		return c.Data(200, info.data, info.mime)

	})

}

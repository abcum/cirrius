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
	"fmt"
	"io"
	"os/exec"
	"path"
	"runtime"
	"strings"

	"github.com/abcum/fibre"
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/css"
	"github.com/tdewolff/minify/html"
	"github.com/tdewolff/minify/js"
	"github.com/tdewolff/minify/json"
	"github.com/tdewolff/minify/svg"
	"github.com/tdewolff/minify/xml"
)

var agents = [...]string{
	"baiduspider",
	"facebookexternalhit",
	"twitterbot",
	"rogerbot",
	"linkedinbot",
	"embedly",
	"quora link preview",
	"showyoubot",
	"outbrain",
	"pinterest",
	"developers.google.com/+/web/snippet",
	"slackbot",
}

func render(c *fibre.Context) (err error, ok bool) {

	usr := strings.ToLower(c.Request().Header().Get("User-Agent"))

	ext := path.Ext(c.Request().URL().Path)

	met := c.Request().Method

	if usr == "" {
		return
	}

	if met != "GET" {
		return
	}

	if ext != "" && ext != ".htm" && ext != ".html" {
		return
	}

	for _, agent := range agents {
		if strings.Contains(agent, usr) {
			return prerender(c)
		}
	}

	return

}

func prerender(c *fibre.Context) (err error, ok bool) {

	var std io.ReadCloser

	url := fmt.Sprintf("%s://%s%s", c.Request().URL().Scheme, c.Request().Host, c.Request().RequestURI)
	bin := "./pjs/phantom-" + runtime.GOOS
	cnf := "pjs/phantom.js"

	cmd := exec.Command(bin, cnf, url)

	if std, err = cmd.StdoutPipe(); err != nil {
		return
	}

	defer std.Close()

	if err = cmd.Start(); err != nil {
		return
	}

	m := minify.New()
	m.AddFunc("text/css", css.Minify)
	m.AddFunc("text/html", html.Minify)
	m.AddFunc("text/javascript", js.Minify)
	m.AddFunc("image/svg+xml", svg.Minify)
	m.AddFunc("application/xml", xml.Minify)
	m.AddFunc("application/json", json.Minify)
	m.AddFunc("application/javascript", js.Minify)

	return m.Minify("text/html", c.Response().ResponseWriter, std), true

}

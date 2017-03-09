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

package minify

import (
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/css"
	"github.com/tdewolff/minify/html"
	"github.com/tdewolff/minify/js"
	"github.com/tdewolff/minify/json"
	"github.com/tdewolff/minify/svg"
	"github.com/tdewolff/minify/xml"
)

var Minifier *minify.M

func init() {
	Minifier = minify.New()
	Minifier.AddFunc("text/css", css.Minify)
	Minifier.AddFunc("text/html", html.Minify)
	Minifier.AddFunc("text/javascript", js.Minify)
	Minifier.AddFunc("image/svg+xml", svg.Minify)
	Minifier.AddFunc("application/xml", xml.Minify)
	Minifier.AddFunc("application/json", json.Minify)
	Minifier.AddFunc("application/ld+json", json.Minify)
	Minifier.AddFunc("application/javascript", js.Minify)
}

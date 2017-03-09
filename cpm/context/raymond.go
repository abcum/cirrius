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

package context

import (
	"fmt"

	"github.com/abcum/orbit"
	"github.com/aymerick/raymond"
)

func init() {

	raymond.RegisterHelper("partial", func(name string, options *raymond.Options) raymond.SafeString {

		orb := options.DataFrame().Get("orb").(*orbit.Orbit)

		data, _, err := orb.File(name, "hbs")
		if err != nil {
			orb.Quit(fmt.Errorf("Unable to find handlebars template '%s'.", name))
		}

		code := fmt.Sprintf("%s", data)

		tmpl, err := raymond.Parse(code)
		if err != nil {
			orb.Quit(fmt.Errorf("Unable to parse handlebars template '%s'.", name))
		}

		priv := options.DataFrame()

		html, err := tmpl.ExecWith(options.Ctx(), priv)
		if err != nil {
			orb.Quit(fmt.Errorf("Unable to render handlebars template '%s'.", name))
		}

		return raymond.SafeString(html)

	})

}

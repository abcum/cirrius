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

package process

import (
	"github.com/abcum/orbit"
	"github.com/robertkrimen/otto"
)

func init() {

	process := map[string]interface{}{}

	// pid is set to 0
	process["pid"] = 0

	// env returns an empty object
	process["env"] = map[string]interface{}{}

	// arch imitates 64bit
	process["arch"] = "x64"

	// platform imitates linux
	process["platform"] = "linux"

	// version defaults to 1.0.0
	process["version"] = "1.0.0"

	process["cwd"] = func(call otto.FunctionCall) otto.Value {
		return otto.UndefinedValue()
	}

	process["nextTick"] = func(call otto.FunctionCall) otto.Value {
		return otto.UndefinedValue()
	}

	process["memoryUsage"] = func(call otto.FunctionCall) otto.Value {
		return otto.UndefinedValue()
	}

	orbit.Def("process", process)

}

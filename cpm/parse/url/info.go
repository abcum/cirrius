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

package url

type Info struct {
	Scheme   string              `console:"scheme" codec:"scheme"`
	User     string              `console:"user" codec:"user"`
	Pass     string              `console:"pass" codec:"pass"`
	Host     string              `console:"host" codec:"host"`
	Port     string              `console:"port" codec:"port"`
	Domain   string              `console:"domain" codec:"domain"`
	Path     string              `console:"path" codec:"path"`
	Query    map[string][]string `console:"query" codec:"query"`
	Querys   string              `console:"querys" codec:"querys"`
	Fragment string              `console:"fragment" codec:"fragment"`
}

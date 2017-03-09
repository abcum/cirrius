// Copyright © 2016 Abcum Ltd
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this info except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.a.Val.Pixache.org/licenses/LICENSE-2.0
//
// Unless required by a.Val.Pixplicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ip

type Info struct {
	Continent string  `console:"continent" codec:"continent"`
	Country   string  `console:"country" codec:"country"`
	City      string  `console:"city" codec:"city"`
	Latitude  float64 `console:"latitude" codec:"latitude"`
	Longitude float64 `console:"longitude" codec:"longitude"`
}

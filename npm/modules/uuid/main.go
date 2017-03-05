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

package uuid

import (
	"github.com/abcum/orbit"
	"github.com/robertkrimen/otto"
	"github.com/satori/go.uuid"
)

func init() {

	orbit.Add("uuid", func(orb *orbit.Orbit) (otto.Value, error) {
		return orb.ToValue(map[string]interface{}{
			"namespace": map[string]interface{}{
				"dns":  uuid.NamespaceDNS,
				"url":  uuid.NamespaceURL,
				"oid":  uuid.NamespaceOID,
				"x500": uuid.NamespaceX500,
			},
			"v1": func() (val otto.Value) {
				val, _ = orb.ToValue(uuid.NewV1().String())
				return
			},
			"v2": func(domain byte) (val otto.Value) {
				val, _ = orb.ToValue(uuid.NewV2(domain).String())
				return
			},
			"v3": func(from, name string) (val otto.Value) {
				val, _ = orb.ToValue(uuid.NewV3(uuid.FromStringOrNil(from), name).String())
				return
			},
			"v4": func() (val otto.Value) {
				val, _ = orb.ToValue(uuid.NewV4().String())
				return
			},
			"v5": func(from, name string) (val otto.Value) {
				val, _ = orb.ToValue(uuid.NewV5(uuid.FromStringOrNil(from), name).String())
				return
			},
		})
	})

}

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
	"github.com/satori/go.uuid"
)

func init() {
	orbit.Add("uuid", New)
}

func New(orb *orbit.Orbit) interface{} {
	return &Module{
		orb:  orb,
		DNS:  uuid.NamespaceDNS,
		URL:  uuid.NamespaceURL,
		OID:  uuid.NamespaceOID,
		X500: uuid.NamespaceX500,
	}
}

type Module struct {
	orb  *orbit.Orbit
	DNS  uuid.UUID
	URL  uuid.UUID
	OID  uuid.UUID
	X500 uuid.UUID
}

func (this *Module) V1() string {
	return uuid.NewV1().String()
}

func (this *Module) V2(domain byte) string {
	return uuid.NewV2(domain).String()
}

func (this *Module) V3(from, name string) string {
	return uuid.NewV3(uuid.FromStringOrNil(from), name).String()
}

func (this *Module) V4() string {
	return uuid.NewV4().String()
}

func (this *Module) V5(from, name string) string {
	return uuid.NewV5(uuid.FromStringOrNil(from), name).String()
}

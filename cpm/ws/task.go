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

// +build !aws,!gce,!as3

package ws

import (
	"github.com/abcum/orbit"
)

type task struct {
	callback func()
}

func (t *task) Startup(orb *orbit.Orbit) {

}

func (t *task) Cleanup(orb *orbit.Orbit) {
	if req, ok := orb.Context().Value("req").(fibrer); ok {
		if req.Fibre().Socket() != nil {
			req.Fibre().Socket().Close(1000)
		}
	}
}

func (t *task) Execute(orb *orbit.Orbit) (err error) {
	t.callback()
	return
}

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

package time

import (
	"time"

	"github.com/abcum/orbit"
	"github.com/abcum/otto"
)

type timer struct {
	timer    *time.Timer
	repeater bool
	callback otto.Value
	argument []interface{}
	duration time.Duration
}

func (t *timer) Startup(orb *orbit.Orbit) {
}

func (t *timer) Cleanup(orb *orbit.Orbit) {
	t.timer.Stop()
}

func (t *timer) Execute(orb *orbit.Orbit) (err error) {

	_, err = t.callback.Call(t.callback, t.argument...)

	if t.repeater {
		t.timer.Reset(t.duration)
	} else {
		orb.Pull(t)
	}

	return

}

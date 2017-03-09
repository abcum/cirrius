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
)

type Duration struct {
	orb *orbit.Orbit
	Val time.Duration `console:"-"`
}

func NewDuration(orb *orbit.Orbit, val time.Duration) *Duration {
	return &Duration{
		orb: orb,
		Val: val,
	}
}

func (this *Duration) Debug() string {
	return this.Val.String()
}

func (this *Duration) Hours() float64 {
	return this.Val.Hours()
}

func (this *Duration) Minutes() float64 {
	return this.Val.Minutes()
}

func (this *Duration) Nanoseconds() int64 {
	return this.Val.Nanoseconds()
}

func (this *Duration) Seconds() float64 {
	return this.Val.Seconds()
}

func (this *Duration) String() string {
	return this.Val.String()
}

func (this *Duration) MarshalText() ([]byte, error) {
	return []byte(this.Val.String()), nil
}

func (this *Duration) MarshalBinary() ([]byte, error) {
	return []byte(this.Val.String()), nil
}

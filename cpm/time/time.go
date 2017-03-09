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

type Time struct {
	orb *orbit.Orbit
	Val time.Time `console:"-"`
}

func NewTime(orb *orbit.Orbit, val time.Time) *Time {
	return &Time{
		orb: orb,
		Val: val,
	}
}

func (this *Time) After(other *Time) bool {
	return this.Val.After(other.Val)
}

func (this *Time) Before(other *Time) bool {
	return this.Val.Before(other.Val)
}

func (this *Time) Debug() string {
	return this.Val.String()
}

func (this *Time) Day() int {
	return this.Val.Day()
}

func (this *Time) Equal(other *Time) bool {
	return this.Val.Equal(other.Val)
}

func (this *Time) Hour() int {
	return this.Val.Hour()
}

func (this *Time) Minute() int {
	return this.Val.Minute()
}

func (this *Time) Month() time.Month {
	return this.Val.Month()
}

func (this *Time) Nanosecond() int {
	return this.Val.Nanosecond()
}

func (this *Time) Second() int {
	return this.Val.Second()
}

func (this *Time) Since(other *Time) *Duration {
	return NewDuration(this.orb, this.Val.Sub(other.Val))
}

func (this *Time) String() string {
	return this.Val.String()
}

func (this *Time) Unix() int64 {
	return this.Val.Unix()
}

func (this *Time) UnixNano() int64 {
	return this.Val.UnixNano()
}

func (this *Time) Until(other *Time) *Duration {
	return NewDuration(this.orb, other.Val.Sub(this.Val))
}

func (this *Time) Year() int {
	return this.Val.Year()
}

func (this *Time) MarshalText() ([]byte, error) {
	out, err := this.Val.MarshalText()
	if err != nil {
		this.orb.Quit(err)
	}
	return out, err
}

func (this *Time) MarshalBinary() ([]byte, error) {
	out, err := this.Val.MarshalBinary()
	if err != nil {
		this.orb.Quit(err)
	}
	return out, err
}

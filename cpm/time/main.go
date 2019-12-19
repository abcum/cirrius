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
	"github.com/olebedev/when"
	"github.com/olebedev/when/rules/common"
	"github.com/olebedev/when/rules/en"
	"github.com/olebedev/when/rules/ru"
)

func init() {
	orbit.Add("time", New)
}

func New(orb *orbit.Orbit) interface{} {
	return &Module{
		orb:      orb,
		RFC822:   time.RFC822,
		RFC822Z:  time.RFC822Z,
		RFC850:   time.RFC850,
		RFC1123:  time.RFC1123,
		RFC1123Z: time.RFC1123Z,
		RFC3339:  time.RFC3339,
		RFC3339N: time.RFC3339Nano,
		KITCHEN:  time.Kitchen,
	}
}

type Module struct {
	orb      *orbit.Orbit
	RFC822   string
	RFC822Z  string
	RFC850   string
	RFC1123  string
	RFC1123Z string
	RFC3339  string
	RFC3339N string
	KITCHEN  string
}

func (this *Module) Now() *Time {
	return NewTime(this.orb, time.Now())
}

func (this *Module) Unix(sec, nsec int64) *Time {
	return NewTime(this.orb, time.Unix(sec, nsec))
}

func (this *Module) Since(val *Time) *Duration {
	return this.Now().Since(val)
}

func (this *Module) Until(val *Time) *Duration {
	return this.Now().Until(val)
}

// Parse parses a formatted string and returns the time value
// it represents. The format defines which input format should
// be used to parse the string.
//
//	var time = require('time');
//	var when = time.Parse("2016-22-04T15:04:05Z07:00", time.RFC3339)
//
func (this *Module) Parse(text, format string) *Time {
	val, err := time.Parse(format, text)
	if err != nil {
		this.orb.Quit(err)
	}
	return NewTime(this.orb, val)
}

// Detect detects a date or time within a piece of text using
// natural language processing, and returns a time object,
// relative to the `base` time argument.
//
//	var time = require('time');
//	var when = time.Detect("Call me next wednesday at 2:25 p.m", time.Now())
//
func (this *Module) Detect(text string, base *Time) *Time {
	w := when.New(nil)
	w.Add(en.All...)
	w.Add(ru.All...)
	w.Add(common.All...)
	res, err := w.Parse(text, base.Val)
	if err != nil {
		this.orb.Quit(err)
	}
	return NewTime(this.orb, res.Time)
}

// Duration detects a time duration parsed from a string. A
// duration string is a possibly signed sequence of decimal
// numbers, each with optional fraction and a unit suffix,
// such as "300ms", "-1.5h" or "2h45m". Valid time units are
// "ns", "us" (or "µs"), "ms", "s", "m", "h".
//
//	var time = require('time');
//	var duration = time.Duration("300ms")
//
func (this *Module) Duration(text string) *Duration {
	dur, err := time.ParseDuration(text)
	if err != nil {
		this.orb.Quit(err)
	}
	return NewDuration(this.orb, dur)
}

// Sleep pauses the execution of the current code for
// the specified number of milliseconds.
//
//	var time = require('time');
// 	time.Sleep(3000);
//	console.log("Finally here");
//
func (this *Module) Sleep(millis int) {
	time.Sleep(time.Duration(millis) * time.Millisecond)
}

// Every schedules the execution of a callback function after
// the specified number of milliseconds.
//
//	var time = require('time');
// 	time.after(1000, function() {
//		console.log("Hi there!")
// 	})
//
// To pass arguments to the callback function, pass them in
// as subsequent arguments to the function.
//
//	var time = require('time');
//	time.every(1000, function() {
//		console.log(arguments);
//	}, 1, 2, "three")
//
func (this *Module) After(after int, call otto.Value, args ...interface{}) *timer {

	timer := &timer{
		callback: call,
		argument: args,
		repeater: false,
		duration: time.Duration(after) * time.Millisecond,
	}

	this.orb.Push(timer)

	timer.timer = time.AfterFunc(timer.duration, func() {
		this.orb.Next(timer)
	})

	return timer

}

// Every schedules repeated execution of a callback function
// after the specified number of milliseconds.
//
//	var time = require('time');
// 	time.every(1000, function() {
//		console.log("Hi there!")
// 	})
//
// To pass arguments to the callback function, pass them in
// as subsequent arguments to the function.
//
//	var time = require('time');
//	time.every(1000, function() {
//		console.log(arguments);
//	}, 1, 2, "three")
//
func (this *Module) Every(every int, call otto.Value, args ...interface{}) *timer {

	timer := &timer{
		callback: call,
		argument: args,
		repeater: true,
		duration: time.Duration(every) * time.Millisecond,
	}

	this.orb.Push(timer)

	timer.timer = time.AfterFunc(timer.duration, func() {
		this.orb.Next(timer)
	})

	return timer

}

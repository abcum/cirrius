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
	"github.com/abcum/fibre"
	"github.com/abcum/orbit"
	"github.com/gorilla/websocket"

	"github.com/abcum/cirrius/cnf"
)

func init() {
	orbit.Add("ws", New)
}

type fibrer interface {
	Fibre() *fibre.Context
}

func New(orb *orbit.Orbit) interface{} {
	return (&Module{
		orb:    orb,
		job:    new(task),
		req:    orb.Context().Value("req").(cnf.Request),
		res:    orb.Context().Value("res").(cnf.Response),
		TEXT:   websocket.TextMessage,
		BINARY: websocket.BinaryMessage,
	}).init()
}

type Module struct {
	orb    *orbit.Orbit
	job    *task
	req    cnf.Request
	res    cnf.Response
	ctx    *fibre.Context
	TEXT   int
	BINARY int
}

func (this *Module) init() *Module {
	if req, ok := this.req.(fibrer); ok {
		this.ctx = req.Fibre()
	}
	return this
}

func (this *Module) Exit() {
	if this.ctx != nil && this.ctx.Socket() != nil {
		this.orb.Pull(this.job)
	}
}

func (this *Module) Init() {
	if this.ctx != nil && this.ctx.Upgrade() == nil {
		this.orb.Push(this.job)
	}
}

func (this *Module) Upgrade() {
	if this.ctx != nil && this.ctx.Upgrade() == nil {
		this.orb.Push(this.job)
	}
}

func (this *Module) Send(kind int, data string) {
	if this.ctx != nil && this.ctx.Socket() != nil {
		this.ctx.Socket().Send(kind, []byte(data))
	}
}

func (this *Module) Recv(callback func(string)) {
	if this.ctx != nil && this.ctx.Socket() != nil {
		go func() {
			defer this.orb.Pull(this.job)
			for {
				if _, req, err := this.ctx.Socket().Read(); err == nil {
					this.job.callback = func() { callback(string(req)) }
					this.orb.Next(this.job)
					continue
				}
				break
			}
		}()
	}
}

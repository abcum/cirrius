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

package stream

import (
	"bytes"
	"io"
	"io/ioutil"

	"github.com/abcum/orbit"
)

func init() {
	orbit.Add("stream", New)
}

func New(orb *orbit.Orbit) interface{} {
	return &Module{
		orb: orb,
	}
}

type Module struct {
	orb *orbit.Orbit
}

func (this *Module) Buffer() *Buffer {
	return NewBuffer(this.orb, &bytes.Buffer{})
}

// ------------------------------
//
// ------------------------------

type Reader struct {
	orb *orbit.Orbit
	stm io.Reader
}

func NewReader(orb *orbit.Orbit, stm io.Reader) *Reader {
	return &Reader{
		orb: orb,
		stm: stm,
	}
}

func (this *Reader) End() {
	// Does nothing
}

func (this *Reader) Close() error {
	return nil // Does nothing
}

func (this *Reader) String() string {
	res, err := ioutil.ReadAll(this.stm)
	if err != nil {
		this.orb.Quit(err)
	}
	return string(res)
}

func (this *Reader) Pipe(w io.Writer) {
	io.Copy(w, this.stm)
}

func (this *Reader) Read(p []byte) (n int, e error) {
	return this.stm.Read(p)
}

// ------------------------------
//
// ------------------------------

type ReadCloser struct {
	orb *orbit.Orbit
	stm io.ReadCloser
}

func NewReadCloser(orb *orbit.Orbit, stm io.ReadCloser) *ReadCloser {
	return (&ReadCloser{
		orb: orb,
		stm: stm,
	}).init()
}

func (this *ReadCloser) init() *ReadCloser {
	this.orb.Once("exit", func() {
		this.Close()
	})
	return this
}

func (this *ReadCloser) End() {
	this.stm.Close()
}

func (this *ReadCloser) Close() error {
	return this.stm.Close()
}

func (this *ReadCloser) String() string {
	res, err := ioutil.ReadAll(this.stm)
	if err != nil {
		this.orb.Quit(err)
	}
	return string(res)
}

func (this *ReadCloser) Pipe(w io.Writer) {
	io.Copy(w, this.stm)
}

func (this *ReadCloser) Read(p []byte) (n int, e error) {
	return this.stm.Read(p)
}

// ------------------------------
//
// ------------------------------

type Writer struct {
	orb *orbit.Orbit
	stm io.Writer
}

func NewWriter(orb *orbit.Orbit, stm io.Writer) *Writer {
	return &Writer{
		orb: orb,
		stm: stm,
	}
}

func (this *Writer) End() {
	// Does nothing
}

func (this *Writer) Close() error {
	return nil // Does nothing
}

func (this *Writer) Consume(r io.Reader) {
	io.Copy(this.stm, r)
}

func (this *Writer) Write(p []byte) (n int, e error) {
	return this.stm.Write(p)
}

// ------------------------------
//
// ------------------------------

type WriteCloser struct {
	orb *orbit.Orbit
	stm io.WriteCloser
}

func NewWriteCloser(orb *orbit.Orbit, stm io.WriteCloser) *WriteCloser {
	return (&WriteCloser{
		orb: orb,
		stm: stm,
	}).init()
}

func (this *WriteCloser) init() *WriteCloser {
	this.orb.Once("exit", func() {
		this.Close()
	})
	return this
}

func (this *WriteCloser) End() {
	this.stm.Close()
}

func (this *WriteCloser) Close() error {
	return this.stm.Close()
}

func (this *WriteCloser) Consume(r io.Reader) {
	io.Copy(this.stm, r)
}

func (this *WriteCloser) Write(p []byte) (n int, e error) {
	return this.stm.Write(p)
}

// ------------------------------
//
// ------------------------------

type Buffer struct {
	orb *orbit.Orbit
	buf *bytes.Buffer
}

func NewBuffer(orb *orbit.Orbit, buf *bytes.Buffer) *Buffer {
	return &Buffer{
		orb: orb,
		buf: buf,
	}
}

func (this *Buffer) End() {
	// Does nothing
}

func (this *Buffer) Close() error {
	return nil // Does nothing
}

func (this *Buffer) Pipe(w io.Writer) {
	io.Copy(w, this.buf)
}

func (this *Buffer) Consume(r io.Reader) {
	io.Copy(this.buf, r)
}

func (this *Buffer) Read(p []byte) (n int, e error) {
	return this.buf.Read(p)
}

func (this *Buffer) Write(p []byte) (n int, e error) {
	return this.buf.Write(p)
}

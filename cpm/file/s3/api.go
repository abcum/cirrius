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

package s3

import (
	"bytes"
	"io"
	"strings"

	"github.com/abcum/orbit"
	"github.com/abcum/otto"

	"github.com/abcum/cirrius/cpm/file"
	"github.com/abcum/cirrius/cpm/stream"
	"github.com/abcum/cirrius/util/args"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type API struct {
	orb *orbit.Orbit
	api *s3.S3
	opt map[string]interface{}
}

func NewAPI(orb *orbit.Orbit, opt map[string]interface{}) *API {
	return (&API{
		orb: orb,
		opt: opt,
	}).init()
}

func (this *API) init() *API {

	cnf := &aws.Config{}

	if val, ok := this.opt["region"].(string); ok {
		cnf.Region = aws.String(val)
	}

	if val, ok := this.opt["insecure"].(bool); ok {
		cnf.DisableSSL = aws.Bool(val)
	}

	ses := session.Must(session.NewSession())

	this.api = s3.New(ses, cnf)

	return this

}

func (this *API) Del(bucket, name string) {

	ctx := this.orb.Context()

	cnf := &s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(name),
	}

	_, err := this.api.DeleteObjectWithContext(ctx, cnf)
	if err != nil {
		this.orb.Quit(err)
	}

	return

}

func (this *API) Get(bucket, name string) *stream.ReadCloser {

	ctx := this.orb.Context()

	cnf := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(name),
	}

	res, err := this.api.GetObjectWithContext(ctx, cnf)
	if err != nil {
		this.orb.Quit(err)
	}

	return stream.NewReadCloser(this.orb, res.Body)

}

func (this *API) Put(call otto.FunctionCall) otto.Value {

	var rdr io.Reader

	args.Size(this.orb, call, 3, 4)

	buck := args.String(this.orb, call, 0)
	name := args.String(this.orb, call, 1)
	body := args.Any(this.orb, call, 2)
	opts := args.Object(this.orb, call, 3)

	switch val := body.(type) {
	case io.ReadSeeker:
		val.Seek(0, 0)
		rdr = val
	case io.Reader:
		fil := file.NewTemp(this.orb)
		fil.Consume(val)
		fil.Seek(0, 0)
		rdr = fil
	case []byte:
		rdr = bytes.NewReader(val)
	case string:
		rdr = strings.NewReader(val)
	}

	ctx := this.orb.Context()

	cnf := &s3.PutObjectInput{
		Bucket: aws.String(buck),
		Key:    aws.String(name),
		Body:   aws.ReadSeekCloser(rdr),
	}

	if val, ok := opts["ACL"].(string); ok {
		cnf.ACL = aws.String(val)
	}

	if val, ok := opts["Tagging"].(string); ok {
		cnf.Tagging = aws.String(val)
	}

	if val, ok := opts["StorageClass"].(string); ok {
		cnf.StorageClass = aws.String(val)
	}

	if val, ok := opts["ServerSideEncryption"].(string); ok {
		cnf.StorageClass = aws.String(val)
	}

	if val, ok := opts["CacheControl"].(string); ok {
		cnf.CacheControl = aws.String(val)
	}

	if val, ok := opts["ContentType"].(string); ok {
		cnf.ContentType = aws.String(val)
	}

	if val, ok := opts["ContentEncoding"].(string); ok {
		cnf.ContentEncoding = aws.String(val)
	}

	if val, ok := opts["ContentDisposition"].(string); ok {
		cnf.ContentDisposition = aws.String(val)
	}

	_, err := this.api.PutObjectWithContext(ctx, cnf)
	if err != nil {
		this.orb.Quit(err)
	}

	return otto.UndefinedValue()

}

func (this *API) File(call otto.FunctionCall) otto.Value {

	args.Size(this.orb, call, 2, 3)

	buck := args.String(this.orb, call, 0)
	name := args.String(this.orb, call, 1)
	opts := args.Object(this.orb, call, 2)

	file := NewFile(this.orb, this.api, buck, name, opts)

	return args.Value(this.orb, file)

}

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
	"github.com/abcum/orbit"

	"github.com/abcum/cirrius/cpm/file"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

type File struct {
	*file.File
	orb  *orbit.Orbit
	api  *s3.S3
	buck string
	name string
	opts map[string]interface{}
}

func NewFile(orb *orbit.Orbit, api *s3.S3, buck, name string, opts map[string]interface{}) *File {
	return (&File{
		orb:  orb,
		api:  api,
		buck: buck,
		name: name,
		opts: opts,
	}).init()
}

func (this *File) init() *File {
	this.File = file.NewTemp(this.orb)
	this.orb.Once("exit", func() {
		this.Save()
	})
	return this
}

func (this *File) Save() {

	this.File.Seek(0, 0)

	ctx := this.orb.Context()

	cnf := &s3.PutObjectInput{
		Bucket: aws.String(this.buck),
		Key:    aws.String(this.name),
		Body:   aws.ReadSeekCloser(this.File),
	}

	if val, ok := this.opts["ACL"].(string); ok {
		cnf.ACL = aws.String(val)
	}

	if val, ok := this.opts["Tagging"].(string); ok {
		cnf.Tagging = aws.String(val)
	}

	if val, ok := this.opts["StorageClass"].(string); ok {
		cnf.StorageClass = aws.String(val)
	}

	if val, ok := this.opts["ServerSideEncryption"].(string); ok {
		cnf.StorageClass = aws.String(val)
	}

	if val, ok := this.opts["CacheControl"].(string); ok {
		cnf.CacheControl = aws.String(val)
	}

	if val, ok := this.opts["ContentType"].(string); ok {
		cnf.ContentType = aws.String(val)
	}

	if val, ok := this.opts["ContentEncoding"].(string); ok {
		cnf.ContentEncoding = aws.String(val)
	}

	if val, ok := this.opts["ContentDisposition"].(string); ok {
		cnf.ContentDisposition = aws.String(val)
	}

	_, err := this.api.PutObjectWithContext(ctx, cnf)
	if err != nil {
		this.orb.Quit(err)
	}

}

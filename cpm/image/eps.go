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

package image

import (
	"bytes"
	"errors"
	"image"
	"io"
	"os/exec"

	"github.com/disintegration/imaging"
)

func init() {
	image.RegisterFormat("eps", `%!PS`, epsDecode, epsConfig)
}

func epsConfig(r io.Reader) (image.Config, error) {
	return image.Config{}, errors.New("Not implemented")
}

func epsDecode(r io.Reader) (image.Image, error) {

	var buf bytes.Buffer

	cmd := exec.Command(
		"gs",
		"-q",
		"-r300",
		"-dSAFER",
		"-dNOPAUSE",
		"-dLastPage=1",
		"-dFirstPage=1",
		"-sDEVICE=png16m",
		"-dTextAlphaBits=4",
		"-dGraphicsAlphaBits=4",
		`-sstdout=%stderr`,
		`-sOutputFile=%stdout`,
		"-",
	)

	cmd.Stdin, cmd.Stdout = r, &buf

	if err := cmd.Run(); err != nil {
		return nil, err
	}

	return imaging.Decode(&buf)

}

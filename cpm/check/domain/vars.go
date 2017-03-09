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

package domain

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"time"
)

var config = &tls.Config{InsecureSkipVerify: true}

var addrError = errors.New("Domain name is not valid.")

type deprecated struct {
	name    string
	expired time.Time
}

var deprecateds = map[x509.SignatureAlgorithm]deprecated{
	x509.MD2WithRSA: {
		name:    "MD2 with RSA",
		expired: time.Date(2008, 1, 1, 0, 0, 0, 0, time.UTC),
	},
	x509.MD5WithRSA: {
		name:    "MD5 with RSA",
		expired: time.Date(2012, 1, 1, 0, 0, 0, 0, time.UTC),
	},
	x509.SHA1WithRSA: {
		name:    "SHA1 with RSA",
		expired: time.Date(2017, 1, 1, 0, 0, 0, 0, time.UTC),
	},
	x509.DSAWithSHA1: {
		name:    "DSA with SHA1",
		expired: time.Date(2017, 1, 1, 0, 0, 0, 0, time.UTC),
	},
	x509.ECDSAWithSHA1: {
		name:    "ECDSA with SHA1",
		expired: time.Date(2017, 1, 1, 0, 0, 0, 0, time.UTC),
	},
}

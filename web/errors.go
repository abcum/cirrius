// Copyright © 2016 Abcum Ltd
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package web

import (
	"github.com/abcum/fibre"
)

func errors(err error, c *fibre.Context) {

	code := 500
	text := ""

	if e, ok := err.(*fibre.HTTPError); ok {
		code = e.Code()
		text = e.Error()
	}

	switch c.Type() {
	default:
		c.Text(code, text)
	case "text/html":
		c.JSON(code, errs[code])
	case "application/json":
		c.JSON(code, errs[code])
	case "application/msgpack":
		c.PACK(code, errs[code])
	}

}

var errs = map[int]interface{}{

	400: map[string]interface{}{
		"code":        400,
		"details":     "Request problems detected",
		"information": "There is a problem with your request. Ensure that the request is valid.",
	},

	401: map[string]interface{}{
		"code":        401,
		"details":     "Authentication failed",
		"information": "Your authentication details are invalid.",
	},

	403: map[string]interface{}{
		"code":        403,
		"details":     "Request resource forbidden",
		"information": "Your request was forbidden. You don't have the necessary permissions to access this resource.",
	},

	404: map[string]interface{}{
		"code":        404,
		"details":     "Request resource not found",
		"information": "The requested resource does not exist. Check that you have entered the url correctly.",
	},

	405: map[string]interface{}{
		"code":        405,
		"details":     "This method is not allowed",
		"information": "The requested http method is not allowed for this resource. Refer to the documentation for allowed methods.",
	},

	409: map[string]interface{}{
		"code":        409,
		"details":     "Request conflict detected",
		"information": "The request could not be processed because of a conflict in the request.",
	},

	413: map[string]interface{}{
		"code":        413,
		"details":     "Request content length too large",
		"information": "All requests to the database must not exceed the predefined content length.",
	},

	415: map[string]interface{}{
		"code":        415,
		"details":     "Unsupported content type requested",
		"information": "The request needs to adhere to certain constraints. Check your request settings and try again.",
	},

	422: map[string]interface{}{
		"code":        422,
		"details":     "Request problems detected",
		"information": "There is a problem with your request. The request appears to contain invalid data.",
	},

	500: map[string]interface{}{
		"code":    500,
		"details": "There was a problem with our servers, and we have been notified",
	},
}

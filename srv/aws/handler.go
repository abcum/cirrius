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

package aws

import (
	"mime"

	"context"

	"github.com/abcum/cirrius/cnf"
	"github.com/abcum/cirrius/exe"

	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/css"
	"github.com/tdewolff/minify/html"
	"github.com/tdewolff/minify/js"
	"github.com/tdewolff/minify/json"
	"github.com/tdewolff/minify/svg"
	"github.com/tdewolff/minify/xml"

	"github.com/aws/aws-lambda-go/events"
)

func Handler(ctx context.Context, req *events.APIGatewayProxyRequest) (res *events.APIGatewayProxyResponse, err error) {

	cnf.Setup()

	request, response := &Req{req: req}, &Res{res: &events.APIGatewayProxyResponse{}}

	// --------------------------------------------------
	// Set mime types
	// --------------------------------------------------

	mime.AddExtensionType(".md", "text/html")
	mime.AddExtensionType(".htm", "text/html")
	mime.AddExtensionType(".tpl", "text/html")
	mime.AddExtensionType(".hbs", "text/html")
	mime.AddExtensionType(".less", "text/css")
	mime.AddExtensionType(".sass", "text/css")
	mime.AddExtensionType(".scss", "text/css")
	mime.AddExtensionType(".gcss", "text/css")

	// --------------------------------------------------
	// Specify context
	// --------------------------------------------------

	ctx = context.WithValue(ctx, "req", request)
	ctx = context.WithValue(ctx, "res", response)

	// --------------------------------------------------
	// Find the file
	// --------------------------------------------------

	file, err := find(ctx, nil, request.Path())
	if err != nil {
		response.Code(404)
		return response.done(), nil
	}

	// --------------------------------------------------
	// Setup the runtime
	// --------------------------------------------------

	run := exe.NewExecutable(load)

	// --------------------------------------------------
	// Execute the file
	// --------------------------------------------------

	if file.Path == "main.js" {

		err = run.LoadJS(ctx, file)

		headers(request, response, file)

		return response.done(), err

	}

	// --------------------------------------------------
	// Execute the file
	// --------------------------------------------------

	switch file.Extn {
	case ".css":
		file, err = run.LoadCSS(ctx, file)
	case ".less":
		file, err = run.LoadLESS(ctx, file)
	case ".scss":
		file, err = run.LoadSCSS(ctx, file)
	case ".html":
		file, err = run.LoadHTML(ctx, file)
	}

	if err != nil {
		response.Code(500)
		return response.done(), nil
	}

	// --------------------------------------------------
	// Minify the output
	// --------------------------------------------------

	m := minify.New()
	m.AddFunc("text/css", css.Minify)
	m.AddFunc("text/html", html.Minify)
	m.AddFunc("text/javascript", js.Minify)
	m.AddFunc("image/svg+xml", svg.Minify)
	m.AddFunc("application/xml", xml.Minify)
	m.AddFunc("application/json", json.Minify)
	m.AddFunc("application/javascript", js.Minify)

	w := m.Writer(file.Mime, response.Body())
	w.Write(file.Data)
	w.Close()

	// --------------------------------------------------
	// Specify any headers
	// --------------------------------------------------

	headers(request, response, file)

	// --------------------------------------------------
	// Submit the response
	// --------------------------------------------------

	response.Code(200)
	return response.done(), nil

}

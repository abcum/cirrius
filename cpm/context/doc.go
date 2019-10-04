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

/*

The global context package enables reading the request and sending a response.

	// Pipe the request body to the response.
	context.req.body.pipe(context.res.body);

	// Or let the response consume the request.
	context.res.body.consume(context.req.body);

	// And set http response headers.
	context.head('Content-Type', 'application/pdf');
	context.head('Content-Description', 'File Transfer');
	context.head('Content-Transfer-Encoding', 'binary');
	context.head('Content-Disposition', 'inline; filename="demo.pdf"');

*/
package context

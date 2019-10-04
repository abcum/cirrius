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

The api/mailgun enables composing and sending emails.

	var fs = require('file');
	var lib = require('api/mailgun');
	var api = lib("example.com", "key-xxxxx", "pubkey-xxxxx");

	var img = fs.load('logo.png');
	var pdf = fs.load('invoice.pdf');

	var msg = api.compose();
	msg.subject("Hi there");
	msg.to("info@abcum.com");
	msg.from("info@example.com");
	msg.replyto("info@example.com");
	msg.text("This is just a test");
	msg.attach("invoice.pdf", pdf);
	msg.inline("logo.png", img);
	msg.send();

*/
package mailgun

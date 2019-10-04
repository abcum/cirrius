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

package as3

import (
	"context"

	"github.com/abcum/cirrius/exe"
	"github.com/aws/aws-lambda-go/events"
)

func Handler(ctx context.Context, evt *events.S3Event) (err error) {

	// IMPORTANT event := &Event{evt}

	// --------------------------------------------------
	// Find the file
	// --------------------------------------------------

	file, err := find(ctx, nil, "main.js")
	if err != nil {
		return err
	}

	// --------------------------------------------------
	// Execute the file
	// --------------------------------------------------

	return exe.NewExecutable(load).LoadJS(ctx, file)

}

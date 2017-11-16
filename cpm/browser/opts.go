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

package browser

import (
	"fmt"

	"github.com/abcum/cirrius/cnf"
)

func opts(d *device) map[string]interface{} {

	var w, h int

	switch d {
	case nil:
		w = macbookPro15.Width
		h = macbookPro15.Height
	default:
		w = d.Width
		h = d.Height
	}

	args := []string{
		fmt.Sprintf("--window-size=%d,%d", w, h),
		"--disable-background-networking",
		"--disable-new-tab-first-run",
		"--disable-component-update",
		"--disable-boot-animation",
		"--disable-add-to-shelf",
		"--disable-default-apps",
		"--disable-cloud-import",
		"--disable-web-security",
		"--disable-appcontainer",
		"--disable-file-system",
		"--disable-extensions",
		"--disable-translate",
		"--disable-infobars",
		"--disable-breakpad",
		"--disable-sync",
		"--disable-gpu",
		"--no-default-browser-check",
		"--no-first-run",
		"--no-sandbox",
		"--headless",
		"--single-process",
	}

	opts := map[string]interface{}{
		"binary": cnf.Settings.Chrome.Executable,
		"detach": false,
		"args":   args,
	}

	if d != nil {
		opts["mobileEmulation"] = map[string]interface{}{
			"deviceMetrics": map[string]interface{}{
				"width":      d.Width,
				"height":     d.Height,
				"pixelRatio": d.Ratio,
			},
			"userAgent": d.Agent,
		}
	}

	return map[string]interface{}{
		"chromeOptions": opts,
	}

}

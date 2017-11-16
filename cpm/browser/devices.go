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

type device struct {
	Width  int
	Height int
	Ratio  float64
	Agent  string
}

// Apple devices

var iPhone4 = &device{
	Width:  320,
	Height: 480,
	Ratio:  2.0,
	Agent:  "Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_1 like Mac OS X) AppleWebKit/603.1.30 (KHTML, like Gecko) Version/10.0 Mobile/14E8301 Safari/602.1",
}

var iPhone5 = &device{
	Width:  320,
	Height: 568,
	Ratio:  2.0,
	Agent:  "Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_1 like Mac OS X) AppleWebKit/603.1.30 (KHTML, like Gecko) Version/10.0 Mobile/14E8301 Safari/602.1",
}

var iPhone6 = &device{
	Width:  375,
	Height: 667,
	Ratio:  2.0,
	Agent:  "Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_1 like Mac OS X) AppleWebKit/603.1.30 (KHTML, like Gecko) Version/10.0 Mobile/14E8301 Safari/602.1",
}

var iPhone6Plus = &device{
	Width:  414,
	Height: 736,
	Ratio:  3.0,
	Agent:  "Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_1 like Mac OS X) AppleWebKit/603.1.30 (KHTML, like Gecko) Version/10.0 Mobile/14E8301 Safari/602.1",
}

var iPad = &device{
	Width:  768,
	Height: 1024,
	Ratio:  2.0,
	Agent:  "Mozilla/5.0 (iPad; CPU OS 10_3_1 like Mac OS X) AppleWebKit/603.1.30 (KHTML, like Gecko) Version/10.0 Mobile/14E8301 Safari/602.1",
}

var iPadPro = &device{
	Width:  1366,
	Height: 1024,
	Ratio:  2.0,
	Agent:  "Mozilla/5.0 (iPad; CPU OS 10_3_1 like Mac OS X) AppleWebKit/603.1.30 (KHTML, like Gecko) Version/10.0 Mobile/14E8301 Safari/602.1",
}

var macbook = &device{
	Width:  1280,
	Height: 800,
	Ratio:  2.0,
	Agent:  "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/603.2.4 (KHTML, like Gecko) Version/10.1.1 Safari/603.2.4",
}

var macbookAir11 = &device{
	Width:  1366,
	Height: 768,
	Ratio:  1.0,
	Agent:  "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/603.2.4 (KHTML, like Gecko) Version/10.1.1 Safari/603.2.4",
}

var macbookAir13 = &device{
	Width:  1440,
	Height: 900,
	Ratio:  1.0,
	Agent:  "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/603.2.4 (KHTML, like Gecko) Version/10.1.1 Safari/603.2.4",
}

var macbookPro13 = &device{
	Width:  1280,
	Height: 800,
	Ratio:  2.0,
	Agent:  "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/603.2.4 (KHTML, like Gecko) Version/10.1.1 Safari/603.2.4",
}

var macbookPro15 = &device{
	Width:  1440,
	Height: 900,
	Ratio:  2.0,
	Agent:  "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/603.2.4 (KHTML, like Gecko) Version/10.1.1 Safari/603.2.4",
}

var iMac27 = &device{
	Width:  2560,
	Height: 1440,
	Ratio:  1.0,
	Agent:  "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/603.2.4 (KHTML, like Gecko) Version/10.1.1 Safari/603.2.4",
}

// Nexus devices

var nexus4 = &device{
	Width:  384,
	Height: 640,
	Ratio:  2.0,
	Agent:  "Mozilla/5.0 (Linux; Android 4.4.2; Nexus 4 Build/KOT49H) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Mobile Safari/537.36",
}

var nexus5 = &device{
	Width:  360,
	Height: 640,
	Ratio:  3.0,
	Agent:  "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Mobile Safari/537.36",
}

var nexus6 = &device{
	Width:  411,
	Height: 731,
	Ratio:  3.5,
	Agent:  "Mozilla/5.0 (Linux; Android 5.1.1; Nexus 6 Build/LYZ28E) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Mobile Safari/537.36",
}

var nexus7 = &device{
	Width:  600,
	Height: 960,
	Ratio:  2.0,
	Agent:  "Mozilla/5.0 (Linux; Android 4.3; Nexus 7 Build/JSS15Q) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Safari/537.36",
}

var nexus9 = &device{
	Width:  768,
	Height: 1024,
	Ratio:  2.0,
	Agent:  "Mozilla/5.0 (Linux; Android 4.3; Nexus 9 Build/JSS15Q) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Safari/537.36",
}

var nexus10 = &device{
	Width:  800,
	Height: 1280,
	Ratio:  2.0,
	Agent:  "Mozilla/5.0 (Linux; Android 4.3; Nexus 10 Build/JSS15Q) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Safari/537.36",
}

// Google devices

var pixel = &device{
	Width:  411,
	Height: 731,
	Ratio:  2.6,
	Agent:  "Mozilla/5.0 (Linux; Android 7.1.2; Pixel Build/N2G47O) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.83 Mobile Safari/537.36",
}

var pixelXL = &device{
	Width:  411,
	Height: 731,
	Ratio:  3.5,
	Agent:  "Mozilla/5.0 (Linux; Android 7.1.2; Pixel Build/N2G47O) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.83 Mobile Safari/537.36",
}

// Samsung devices

var galaxyS5 = &device{
	Width:  411,
	Height: 731,
	Ratio:  2.6,
	Agent:  "Mozilla/5.0 (Linux; Android 5.0; SM-G900P Build/LRX21T) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Mobile Safari/537.36",
}

// Windows devices

var surface = &device{
	Width:  1366,
	Height: 768,
	Ratio:  1.0,
	Agent:  "Mozilla/5.0 (Windows NT 10.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36 Edge/12.10136",
}

var surface2 = &device{
	Width:  1280,
	Height: 720,
	Ratio:  1.5,
	Agent:  "Mozilla/5.0 (Windows NT 10.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36 Edge/12.10136",
}

var surface3 = &device{
	Width:  1280,
	Height: 720,
	Ratio:  1.5,
	Agent:  "Mozilla/5.0 (Windows NT 10.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36 Edge/12.10136",
}

var surfaceBook = &device{
	Width:  1500,
	Height: 1000,
	Ratio:  2.0,
	Agent:  "Mozilla/5.0 (Windows NT 10.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36 Edge/12.10136",
}

var surfacePro = &device{
	Width:  1280,
	Height: 720,
	Ratio:  1.5,
	Agent:  "Mozilla/5.0 (Windows NT 10.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36 Edge/12.10136",
}

var surfacePro3 = &device{
	Width:  1440,
	Height: 960,
	Ratio:  1.5,
	Agent:  "Mozilla/5.0 (Windows NT 10.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36 Edge/12.10136",
}

var surfacePro4 = &device{
	Width:  1440,
	Height: 960,
	Ratio:  2.0,
	Agent:  "Mozilla/5.0 (Windows NT 10.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36 Edge/12.10136",
}

var surfaceHub55 = &device{
	Width:  1920,
	Height: 1080,
	Ratio:  2.0,
	Agent:  "Mozilla/5.0 (Windows NT 10.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36 Edge/12.10136",
}

var surfaceHub84 = &device{
	Width:  3840,
	Height: 2160,
	Ratio:  2.0,
	Agent:  "Mozilla/5.0 (Windows NT 10.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36 Edge/12.10136",
}

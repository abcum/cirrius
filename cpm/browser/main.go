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
	"github.com/abcum/orbit"
	"github.com/abcum/webdriver"

	"github.com/abcum/cirrius/cnf"
)

func init() {
	orbit.Add("browser", New)
}

func New(orb *orbit.Orbit) interface{} {
	return &Module{
		orb: orb,
		wed: webdriver.NewDriver(cnf.Settings.Chrome.Endpoint),
		// Device metrics
		IPhone4:      iPhone4,
		IPhone5:      iPhone5,
		IPhone6:      iPhone6,
		IPhone6Plus:  iPhone6Plus,
		IPad:         iPad,
		IPadPro:      iPadPro,
		Macbook:      macbook,
		MacbookAir11: macbookAir11,
		MacbookAir13: macbookAir13,
		MacbookPro13: macbookPro13,
		MacbookPro15: macbookPro15,
		IMac27:       iMac27,
		Nexus4:       nexus4,
		Nexus5:       nexus5,
		Nexus6:       nexus6,
		Nexus7:       nexus7,
		Nexus9:       nexus9,
		Nexus10:      nexus10,
		Pixel:        pixel,
		PixelXL:      pixelXL,
		GalaxyS5:     galaxyS5,
		Surface:      surface,
		Surface2:     surface2,
		Surface3:     surface3,
		SurfaceBook:  surfaceBook,
		SurfacePro:   surfacePro,
		SurfacePro3:  surfacePro3,
		SurfacePro4:  surfacePro4,
		SurfaceHub55: surfaceHub55,
		SurfaceHub84: surfaceHub84,
	}
}

type Module struct {
	orb          *orbit.Orbit
	wed          *webdriver.Driver
	IPhone4      *device
	IPhone5      *device
	IPhone6      *device
	IPhone6Plus  *device
	IPad         *device
	IPadPro      *device
	Macbook      *device
	MacbookAir11 *device
	MacbookAir13 *device
	MacbookPro13 *device
	MacbookPro15 *device
	IMac27       *device
	Nexus4       *device
	Nexus5       *device
	Nexus6       *device
	Nexus7       *device
	Nexus9       *device
	Nexus10      *device
	Pixel        *device
	PixelXL      *device
	GalaxyS5     *device
	Surface      *device
	Surface2     *device
	Surface3     *device
	SurfaceBook  *device
	SurfacePro   *device
	SurfacePro3  *device
	SurfacePro4  *device
	SurfaceHub55 *device
	SurfaceHub84 *device
}

func (this *Module) New(d ...*device) *Page {
	var dev *device
	if len(d) > 0 {
		dev = d[0]
	}
	ses, err := this.wed.Session(opts(dev), nil)
	if err != nil {
		this.orb.Quit(err)
	}
	return NewPage(this.orb, ses)
}

// Copyright 2011 go-gdal. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gdal

import "testing"

func TestInit(t *testing.T) {
	if gdalReady == false  {
	      t.Errorf("Library not initialized")
	}
}

func TestTiffDriver(t *testing.T) {
	driver := gdal.GetDriverByName("GTiff")
	if driver == nil {
			t.Errorf("GeoTIFF driver not found")
	}
}

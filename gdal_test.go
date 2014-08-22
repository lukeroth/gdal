// Copyright 2011 go-gdal. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gdal

import "testing"

func TestTiffDriver(t *testing.T) {
	_, err := GetDriverByName("GTiff")
	if err != nil {
		t.Errorf(err.Error())
	}
}

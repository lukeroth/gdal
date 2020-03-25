// Copyright 2011 go-gdal. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gdal

import (
	"testing"
)

func TestTiffDriver(t *testing.T) {
	_, err := GetDriverByName("GTiff")
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestMissingMetadata(t *testing.T) {
	ds, err := Open("testdata/tiles.gpkg", ReadOnly)
	if err != nil {
		t.Fatalf("failed to open test file: %v", err)
	}

	metadata := ds.Metadata("something-that-wont-exist")
	if len(metadata) != 0 {
		t.Errorf("got %d items, want 0", len(metadata))
	}
}

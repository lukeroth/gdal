package gdal

import (
	"fmt"
	"testing"
)

func TestWarp(t *testing.T) {
	srcDS, err := Open("testdata/tiles.gpkg", ReadOnly)
	if err != nil {
		t.Errorf("Open: %v", err)
	}

	opts := []string{"-t_srs", "epsg:3857", "-of", "GPKG"}

	dstDS, err := Warp("/tmp/tiles-3857.gpkg", []Dataset{srcDS}, opts)
	if err != nil {
		t.Errorf("Warp: %v", err)
	}
	fmt.Printf("dst size: %d, %d", dstDS.RasterXSize(), dstDS.RasterYSize())

	pngdriver, err := GetDriverByName("PNG")
	pngdriver.CreateCopy("/tmp/foo.png", dstDS, 0, nil, nil, nil)
	dstDS.Close()
}

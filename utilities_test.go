package gdal

import (
	"testing"
)

func TestVectorTranslate(t *testing.T) {
	srcDS, err := OpenEx("testdata/test.shp", OFReadOnly, nil, nil, nil)
	if err != nil {
		t.Errorf("Open: %v", err)
	}

	opts := []string{"-t_srs", "epsg:4326", "-f", "GeoJSON"}

	dstDS, err := VectorTranslate("/tmp/test4326.geojson", []Dataset{srcDS}, opts)
	if err != nil {
		t.Errorf("Warp: %v", err)
	}
	dstDS.Close()
	dstDS, err = OpenEx("/tmp/test4326.geojson", OFReadOnly|OFVector, []string{"geojson"}, nil, nil)
	if err != nil {
		t.Errorf("Open after translate: %v", err)
	}
	dstDS.Close()

}
func TestRasterize(t *testing.T) {
	srcDS, err := OpenEx("testdata/test.shp", OFReadOnly, nil, nil, nil)
	if err != nil {
		t.Errorf("Open: %v", err)
	}

	opts := []string{"-a", "code", "-tr", "10", "10"}

	dstDS, err := Rasterize("/tmp/rasterized.tif", srcDS, opts)
	if err != nil {
		t.Errorf("Warp: %v", err)
	}
	dstDS.Close()
	dstDS, err = Open("/tmp/rasterized.tif", ReadOnly)
	if err != nil {
		t.Errorf("Open after vector translate: %v", err)
	}
	dstDS.Close()

}

func TestWarp(t *testing.T) {
	srcDS, err := Open("testdata/tiles.gpkg", ReadOnly)
	if err != nil {
		t.Errorf("Open: %v", err)
	}

	opts := []string{"-t_srs", "epsg:3857", "-of", "GPKG"}

	dstDS, err := Warp("/tmp/tiles-3857.gpkg", nil, []Dataset{srcDS}, opts)
	if err != nil {
		t.Errorf("Warp: %v", err)
	}

	pngdriver, err := GetDriverByName("PNG")
	pngdriver.CreateCopy("/tmp/foo.png", dstDS, 0, nil, nil, nil)
	dstDS.Close()
}

func TestTranslate(t *testing.T) {
	srcDS, err := Open("testdata/tiles.gpkg", ReadOnly)
	if err != nil {
		t.Errorf("Open: %v", err)
	}

	opts := []string{"-of", "GTiff"}

	dstDS, err := Translate("/tmp/tiles.tif", srcDS, opts)
	if err != nil {
		t.Errorf("Warp: %v", err)
	}
	dstDS.Close()

	dstDS, err = Open("/tmp/tiles.tif", ReadOnly)
	if err != nil {
		t.Errorf("Open after raster translate: %v", err)
	}
	dstDS.Close()
}

func TestDEMProcessing(t *testing.T) {
	srcDS, err := Open("testdata/demproc.tif", ReadOnly)
	if err != nil {
		t.Errorf("Open: %v", err)
	}

	opts := []string{"-of", "GTiff"}

	dstDS, err := DEMProcessing("/tmp/demproc_output.tif", srcDS, "color-relief", "testdata/demproc_colors.txt", opts)
	if err != nil {
		t.Errorf("DEMProcessing: %v", err)
	}
	dstDS.Close()

	dstDS, err = Open("/tmp/demproc_output.tif", ReadOnly)
	if err != nil {
		t.Errorf("Open after raster DEM Processing: %v", err)
	}
	dstDS.Close()
}

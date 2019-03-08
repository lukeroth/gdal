package main

import (
	"flag"
	"fmt"
	"github.com/lukeroth/gdal"
)

func main() {
	flag.Parse()
	filename := flag.Arg(0)
	if filename == "" {
		fmt.Printf("Usage: test_tiff [filename]\n")
		return
	}
	fmt.Printf("Filename: %s\n", filename)

	fmt.Printf("Allocating buffer\n")
	var buffer [256 * 256]uint8
	//	buffer := make([]uint8, 256 * 256)

	fmt.Printf("Computing values\n")
	for x := 0; x < 256; x++ {
		for y := 0; y < 256; y++ {
			loc := x + y*256
			val := x + y
			if val >= 256 {
				val -= 256
			}
			buffer[loc] = uint8(val)
		}
	}

	fmt.Printf("%d drivers available\n", gdal.GetDriverCount())
	for x := 0; x < gdal.GetDriverCount(); x++ {
		driver := gdal.GetDriver(x)
		fmt.Printf("%s: %s\n", driver.ShortName(), driver.LongName())
	}

	fmt.Printf("Loading driver\n")
	driver, err := gdal.GetDriverByName("GTiff")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("Creating dataset\n")
	dataset := driver.Create(filename, 256, 256, 1, gdal.Byte, nil)
	defer dataset.Close()

	fmt.Printf("Creating projection\n")
	spatialRef := gdal.CreateSpatialReference("")

	fmt.Printf("Setting EPSG code\n")
	spatialRef.FromEPSG(3857)

	fmt.Printf("Converting to WKT\n")
	srString, err := spatialRef.ToWKT()

	fmt.Printf("Assigning projection: %s\n", srString)
	dataset.SetProjection(srString)

	fmt.Printf("Setting geotransform\n")
	dataset.SetGeoTransform([6]float64{444720, 30, 0, 3751320, 0, -30})

	fmt.Printf("Getting raster band\n")
	raster := dataset.RasterBand(1)

	fmt.Printf("Writing to raster band\n")
	raster.IO(gdal.Write, 0, 0, 256, 256, buffer, 256, 256, 0, 0)

	fmt.Printf("Reading geotransform:")
	geoTransform := dataset.GeoTransform()
	fmt.Printf("%v, %v, %v, %v, %v, %v\n",
		geoTransform[0], geoTransform[1], geoTransform[2], geoTransform[3], geoTransform[4], geoTransform[5])

	fmt.Printf("Reading projection:")
	wkt := dataset.Projection()
	fmt.Printf("%s\n", wkt)

	fmt.Printf("End program\n")
}

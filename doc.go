/*
Package gdal provides a wrapper for GDAL, the Geospatial Data Abstraction Library.  This C/C++ library provides access to a large number of geospatial data formats.  It also contains a wrapper for the related OGR library, although coverage of this library is currently much more limited.

Limitations

The majority of the GDAL C API is covered.  Areas that are not currently covered include the asynchronous readers, color and raster attribute tables, and quite a few less oftenly used functions.  Spatial reference support is sufficient to support assigning spatial reference values to GDAL datasets, but little more.

Usage

A simple program to create a georeferenced blank 256x256 GeoTIFF:
	package main

	import (
		"fmt"
		"flag"
		"github.com/lukeroth/gdal"
	)

	func main() {
		flag.Parse()
		filename := flag.Arg(0)
		if filename == "" {
			fmt.Printf("Usage: test_tiff [filename]\n")
			return
		}
		buffer := make([]uint8, 256 * 256)
		
		driver, err := gdal.GetDriverByName("GTiff")
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		dataset := driver.Create(filename, 256, 256, 1, gdal.Byte, nil)
		defer dataset.Close()
		
		spatialRef := gdal.CreateSpatialReference("")
		spatialRef.FromEPSG(3857)
		srString, err := spatialRef.ToWKT()
		dataset.SetProjection(srString)
		dataset.SetGeoTransform([]float64{444720, 30, 0, 3751320, 0, -30})
		raster := dataset.RasterBand(1)
		raster.IO(gdal.Write, 0, 0, 256, 256, buffer, 256, 256, 0, 0)
	}
More examples can be found in the ./examples subdirectory.

Recent changes

10/2/2012: Initial OGR and OSR support added
*/
package gdal
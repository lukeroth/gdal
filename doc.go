/*
Package gdal provides a wrapper for GDAL, the Geospatial Data Abstraction Library.  This C/C++ library provides access to a large number of geospatial raster data formats.  It also contains a wrapper for the related OGR Simple Feature Library which provides similar functionality for vector formats.

Limitations

The majority of the GDAL C API is covered, with the exception of asynchronous readers and a number of less oftenly used functions.  OGR support is currently much less complete.  Spatial reference support is sufficient to support assigning spatial reference values to GDAL datasets from WKT and Proj.4 definitions, but does not cover the majority of spatial reference formats yet.

This wrapper has only been tested on 64-bit Ubuntu Linux, with version 1.9.1 of the GDAL library.

Usage

A simple program to create a georeferenced blank 256x256 GeoTIFF:
	package main

	import (
		"fmt"
		"flag"
		gdal "github.com/lukeroth/gdal_go"
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

A more complete changelog can be found on Github.
10/4/2012: Renamed project to work better with go tools, additional OGR code
10/3/2012: Restructed OGR code, added initial algorithm functions and placeholders
10/2/2012: Initial OGR and OSR support added; color table and raster attribute table support added.
*/
package gdal

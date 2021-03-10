/*
Package gdal provides a wrapper for GDAL, the Geospatial Data Abstraction Library.  This C/C++ library provides access to a large number of geospatial raster data formats.  It also contains a wrapper for the related OGR Simple Feature Library which provides similar functionality for vector formats.

Limitations

Some less oftenly used functions are not yet implemented.  The majoriry of these involve style tables, asynchronous I/O, and GCPs.

The documentation is fairly limited, but the functionality fairly closely matches that of the C++ api.

This wrapper has most recently been tested on Windows7, using the MinGW32_x64 compiler and GDAL version 1.11.

Usage

A simple program to create a georeferenced blank 256x256 GeoTIFF:
	package main

	import (
		"fmt"
		"flag"
		gdal "github.com/lukeroth/gdal "
	)

	func main() {
		flag.Parse()
		filename := flag.Arg(0)
		if filename == "" {
			fmt.Printf("Usage: tiff [filename]\n")
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

*/
package gdal

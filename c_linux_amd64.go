// +build linux,amd64 darwin,amd64

package gdal

// #cgo linux pkg-config: gdal

/*
#cgo LDFLAGS: -lgdal
#cgo CFLAGS: -I/usr/include/gdal
*/
import "C"

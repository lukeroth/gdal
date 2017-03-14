package gdal

/*
#include "go_gdal.h"
#include "gdal_version.h"

#cgo linux  pkg-config: gdal
#cgo darwin pkg-config: gdal
#cgo windows LDFLAGS: -Lc:/gdal/release-1600-x64/lib -lgdal_i
#cgo windows CFLAGS: -IC:/gdal/release-1600-x64/include
*/
import "C"
import (
	"fmt"
	"unsafe"
)

var _ = fmt.Println

/* --------------------------------------------- */
/* Command line utility wrapper functions        */
/* --------------------------------------------- */

// gdalwarp

func Warp(dstDS string, sourceDS []Dataset, options []string) (Dataset, error) {
	length := len(options)
	opts := make([]*C.char, length+1)
	for i := 0; i < length; i++ {
		opts[i] = C.CString(options[i])
		defer C.free(unsafe.Pointer(opts[i]))
	}
	opts[length] = (*C.char)(unsafe.Pointer(nil))
	warpopts := C.GDALWarpAppOptionsNew(
		(**C.char)(unsafe.Pointer(&opts[0])),
		(*C.GDALWarpAppOptionsForBinary)(unsafe.Pointer(nil)))
	defer C.GDALWarpAppOptionsFree(warpopts)

	srcDS := make([]C.GDALDatasetH, len(sourceDS))
	for i, ds := range sourceDS {
		srcDS[i] = ds.cval
	}
	var cerr C.int
	if dstDS == "" {
		dstDS = "MEM:::"
	}
	cdstDS := C.CString(dstDS)
	defer C.free(unsafe.Pointer(cdstDS))
	ds := C.GDALWarp(cdstDS, nil,
		C.int(len(sourceDS)),
		(*C.GDALDatasetH)(unsafe.Pointer(&srcDS[0])),
		warpopts, &cerr)
	if cerr != 0 {
		return Dataset{}, fmt.Errorf("warp failed with code %d", cerr)
	}
	return Dataset{ds}, nil

}

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

func stringArrayContains(array []string, needle string) bool {
	for _, s := range array {
		if s == needle {
			return true
		}
	}
	return false
}

func Warp(dstDS string, sourceDS []Dataset, options []string) (Dataset, error) {
	if dstDS == "" {
		dstDS = "MEM:::"
		if !stringArrayContains(options, "-of") {
			options = append([]string{"-of", "MEM"}, options...)
		}
	}
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

func Translate(dstDS string, sourceDS Dataset, options []string) (Dataset, error) {
	if dstDS == "" {
		dstDS = "MEM:::"
		if !stringArrayContains(options, "-of") {
			options = append([]string{"-of", "MEM"}, options...)
		}
	}
	length := len(options)
	opts := make([]*C.char, length+1)
	for i := 0; i < length; i++ {
		opts[i] = C.CString(options[i])
		defer C.free(unsafe.Pointer(opts[i]))
	}
	opts[length] = (*C.char)(unsafe.Pointer(nil))
	translateopts := C.GDALTranslateOptionsNew(
		(**C.char)(unsafe.Pointer(&opts[0])),
		(*C.GDALTranslateOptionsForBinary)(unsafe.Pointer(nil)))
	defer C.GDALTranslateOptionsFree(translateopts)

	var cerr C.int
	cdstDS := C.CString(dstDS)
	defer C.free(unsafe.Pointer(cdstDS))
	ds := C.GDALTranslate(cdstDS,
		sourceDS.cval,
		translateopts, &cerr)
	if cerr != 0 {
		return Dataset{}, fmt.Errorf("translate failed with code %d", cerr)
	}
	return Dataset{ds}, nil

}

func VectorTranslate(dstDS string, sourceDS []Dataset, options []string) (Dataset, error) {
	if dstDS == "" {
		dstDS = "MEM:::"
		if !stringArrayContains(options, "-f") {
			options = append([]string{"-f", "MEM"}, options...)
		}
	}
	length := len(options)
	opts := make([]*C.char, length+1)
	for i := 0; i < length; i++ {
		opts[i] = C.CString(options[i])
		defer C.free(unsafe.Pointer(opts[i]))
	}
	opts[length] = (*C.char)(unsafe.Pointer(nil))
	translateopts := C.GDALVectorTranslateOptionsNew(
		(**C.char)(unsafe.Pointer(&opts[0])),
		(*C.GDALVectorTranslateOptionsForBinary)(unsafe.Pointer(nil)))
	defer C.GDALVectorTranslateOptionsFree(translateopts)

	srcDS := make([]C.GDALDatasetH, len(sourceDS))
	for i, ds := range sourceDS {
		srcDS[i] = ds.cval
	}

	var cerr C.int
	cdstDS := C.CString(dstDS)
	defer C.free(unsafe.Pointer(cdstDS))
	ds := C.GDALVectorTranslate(cdstDS, nil,
		C.int(len(sourceDS)),
		(*C.GDALDatasetH)(unsafe.Pointer(&srcDS[0])),
		translateopts, &cerr)
	if cerr != 0 {
		return Dataset{}, fmt.Errorf("vector translate failed with code %d", cerr)
	}
	return Dataset{ds}, nil

}

func Rasterize(dstDS string, sourceDS Dataset, options []string) (Dataset, error) {
	if dstDS == "" {
		dstDS = "MEM:::"
		if !stringArrayContains(options, "-f") {
			options = append([]string{"-of", "MEM"}, options...)
		}
	}
	length := len(options)
	opts := make([]*C.char, length+1)
	for i := 0; i < length; i++ {
		opts[i] = C.CString(options[i])
		defer C.free(unsafe.Pointer(opts[i]))
	}
	opts[length] = (*C.char)(unsafe.Pointer(nil))
	rasterizeopts := C.GDALRasterizeOptionsNew(
		(**C.char)(unsafe.Pointer(&opts[0])),
		(*C.GDALRasterizeOptionsForBinary)(unsafe.Pointer(nil)))
	defer C.GDALRasterizeOptionsFree(rasterizeopts)

	var cerr C.int
	cdstDS := C.CString(dstDS)
	defer C.free(unsafe.Pointer(cdstDS))
	ds := C.GDALRasterize(cdstDS, nil,
		sourceDS.cval,
		rasterizeopts, &cerr)
	if cerr != 0 {
		return Dataset{}, fmt.Errorf("rasterize failed with code %d", cerr)
	}
	return Dataset{ds}, nil

}

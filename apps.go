package gdal

/*
#include "go_gdal.h"
#include "gdal_version.h"
#include "gdal_utils.h"

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
/* GDAL utilities                                */
/* --------------------------------------------- */

type GDALTranslateOptions struct {
	cval *C.GDALTranslateOptions
}
type GDALWarpAppOptions struct {
	cval *C.GDALWarpAppOptions
}

func GDALTranslate(
	destName string,
	srcDS Dataset,
	options []string,
) Dataset {

	var err C.int

	length := len(options)
	cOptions := make([]*C.char, length+1)
	for i := 0; i < length; i++ {
		cOptions[i] = C.CString(options[i])
		defer C.free(unsafe.Pointer(cOptions[i]))
	}
	cOptions[length] = (*C.char)(unsafe.Pointer(nil))

	gdalTranslateOptions := GDALTranslateOptions{C.GDALTranslateOptionsNew((**C.char)(unsafe.Pointer(&cOptions[0])), nil)}

	outputDs := C.GDALTranslate(
		C.CString(destName),
		srcDS.cval,
		gdalTranslateOptions.cval,
		&err,
	)

	return Dataset{outputDs}

}

func GDALWarp(
	destName string,
	dstDs Dataset,
	srcDs []Dataset,
	options []string,
) Dataset {

	var err C.int

	length := len(options)
	cOptions := make([]*C.char, length+1)
	for i := 0; i < length; i++ {
		cOptions[i] = C.CString(options[i])
		defer C.free(unsafe.Pointer(cOptions[i]))
	}
	cOptions[length] = (*C.char)(unsafe.Pointer(nil))

	gdalWarpOptions := GDALWarpAppOptions{C.GDALWarpAppOptionsNew((**C.char)(unsafe.Pointer(&cOptions[0])), nil)}

	pahSrcDs := make([]C.GDALDatasetH, len(srcDs)+1)
	for i := 0; i < len(srcDs); i++ {
		pahSrcDs[i] = srcDs[i].cval
	}
	pahSrcDs[len(srcDs)] = (C.GDALDatasetH)(unsafe.Pointer(nil))

	outputDs := C.GDALWarp(
		C.CString(destName),
		dstDs.cval,
		C.int(len(srcDs)),
		(*C.GDALDatasetH)(unsafe.Pointer(&pahSrcDs[0])),
		gdalWarpOptions.cval,
		&err,
	)

	return Dataset{outputDs}

}

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
	"errors"
	"fmt"
	"math"
	"unsafe"
)

var _ = fmt.Println

/* --------------------------------------------- */
/* Misc functions                                */
/* --------------------------------------------- */

// Compute optimal PCT for RGB image
func ComputeMedianCutPCT(
	red, green, blue RasterBand,
	colors int,
	ct ColorTable,
	progress ProgressFunc,
	data interface{},
) int {
	arg := &goGDALProgressFuncProxyArgs{
		progress, data,
	}

	err := C.GDALComputeMedianCutPCT(
		red.cval,
		green.cval,
		blue.cval,
		nil,
		C.int(colors),
		ct.cval,
		C.goGDALProgressFuncProxyB(),
		unsafe.Pointer(arg),
	)
	return int(err)
}

// 24bit to 8bit conversion with dithering
func DitherRGB2PCT(
	red, green, blue, target RasterBand,
	ct ColorTable,
	progress ProgressFunc,
	data interface{},
) int {
	arg := &goGDALProgressFuncProxyArgs{
		progress, data,
	}

	err := C.GDALDitherRGB2PCT(
		red.cval,
		green.cval,
		blue.cval,
		target.cval,
		ct.cval,
		C.goGDALProgressFuncProxyB(),
		unsafe.Pointer(arg),
	)
	return int(err)
}

// Compute checksum for image region
func (rb RasterBand) Checksum(xOff, yOff, xSize, ySize int) int {
	sum := C.GDALChecksumImage(rb.cval, C.int(xOff), C.int(yOff), C.int(xSize), C.int(ySize))
	return int(sum)
}

// Compute the proximity of all pixels in the image to a set of pixels in the source image
func (src RasterBand) ComputeProximity(
	dest RasterBand,
	options []string,
	progress ProgressFunc,
	data interface{},
) error {
	arg := &goGDALProgressFuncProxyArgs{
		progress, data,
	}

	length := len(options)
	opts := make([]*C.char, length+1)
	for i := 0; i < length; i++ {
		opts[i] = C.CString(options[i])
		defer C.free(unsafe.Pointer(opts[i]))
	}
	opts[length] = (*C.char)(unsafe.Pointer(nil))

	return C.GDALComputeProximity(
		src.cval,
		dest.cval,
		(**C.char)(unsafe.Pointer(&opts[0])),
		C.goGDALProgressFuncProxyB(),
		unsafe.Pointer(arg),
	).Err()
}

// Fill selected raster regions by interpolation from the edges
func (src RasterBand) FillNoData(
	mask RasterBand,
	distance float64,
	iterations int,
	options []string,
	progress ProgressFunc,
	data interface{},
) error {
	arg := &goGDALProgressFuncProxyArgs{
		progress, data,
	}

	length := len(options)
	opts := make([]*C.char, length+1)
	for i := 0; i < length; i++ {
		opts[i] = C.CString(options[i])
		defer C.free(unsafe.Pointer(opts[i]))
	}
	opts[length] = (*C.char)(unsafe.Pointer(nil))

	return C.GDALFillNodata(
		src.cval,
		mask.cval,
		C.double(distance),
		0,
		C.int(iterations),
		(**C.char)(unsafe.Pointer(&opts[0])),
		C.goGDALProgressFuncProxyB(),
		unsafe.Pointer(arg),
	).Err()
}

// Create polygon coverage from raster data using an integer buffer
func (src RasterBand) Polygonize(
	mask RasterBand,
	layer Layer,
	fieldIndex int,
	options []string,
	progress ProgressFunc,
	data interface{},
) error {
	arg := &goGDALProgressFuncProxyArgs{
		progress, data,
	}

	length := len(options)
	opts := make([]*C.char, length+1)
	for i := 0; i < length; i++ {
		opts[i] = C.CString(options[i])
		defer C.free(unsafe.Pointer(opts[i]))
	}
	opts[length] = (*C.char)(unsafe.Pointer(nil))

	return C.GDALPolygonize(
		src.cval,
		mask.cval,
		layer.cval,
		C.int(fieldIndex),
		(**C.char)(unsafe.Pointer(&opts[0])),
		C.goGDALProgressFuncProxyB(),
		unsafe.Pointer(arg),
	).Err()
}

// Create polygon coverage from raster data using a floating point buffer
func (src RasterBand) FPolygonize(
	mask RasterBand,
	layer Layer,
	fieldIndex int,
	options []string,
	progress ProgressFunc,
	data interface{},
) error {
	arg := &goGDALProgressFuncProxyArgs{
		progress, data,
	}

	length := len(options)
	opts := make([]*C.char, length+1)
	for i := 0; i < length; i++ {
		opts[i] = C.CString(options[i])
		defer C.free(unsafe.Pointer(opts[i]))
	}
	opts[length] = (*C.char)(unsafe.Pointer(nil))

	return C.GDALFPolygonize(
		src.cval,
		mask.cval,
		layer.cval,
		C.int(fieldIndex),
		(**C.char)(unsafe.Pointer(&opts[0])),
		C.goGDALProgressFuncProxyB(),
		unsafe.Pointer(arg),
	).Err()
}

// Removes small raster polygons
func (src RasterBand) SieveFilter(
	mask, dest RasterBand,
	threshold, connectedness int,
	options []string,
	progress ProgressFunc,
	data interface{},
) error {
	arg := &goGDALProgressFuncProxyArgs{
		progress, data,
	}

	length := len(options)
	opts := make([]*C.char, length+1)
	for i := 0; i < length; i++ {
		opts[i] = C.CString(options[i])
		defer C.free(unsafe.Pointer(opts[i]))
	}
	opts[length] = (*C.char)(unsafe.Pointer(nil))

	return C.GDALSieveFilter(
		src.cval,
		mask.cval,
		dest.cval,
		C.int(threshold),
		C.int(connectedness),
		(**C.char)(unsafe.Pointer(&opts[0])),
		C.goGDALProgressFuncProxyB(),
		unsafe.Pointer(arg),
	).Err()
}

/* --------------------------------------------- */
/* Warp functions                                */
/* --------------------------------------------- */

//Unimplemented: CreateGenImgProjTransformer
//Unimplemented: CreateGenImgProjTransformer2
//Unimplemented: CreateGenImgProjTransformer3
//Unimplemented: SetGenImgProjTransformerDstGeoTransform
//Unimplemented: DestroyGenImgProjTransformer
//Unimplemented: GenImgProjTransform

//Unimplemented: CreateReprojectionTransformer
//Unimplemented: DestroyReprojection
//Unimplemented: ReprojectionTransform
//Unimplemented: CreateGCPTransformer
//Unimplemented: CreateGCPRefineTransformer
//Unimplemented: DestroyGCPTransformer
//Unimplemented: GCPTransform

//Unimplemented: CreateTPSTransformer
//Unimplemented: DestroyTPSTransformer
//Unimplemented: TPSTransform

//Unimplemented: CreateRPCTransformer
//Unimplemented: DestroyRPCTransformer
//Unimplemented: RPCTransform

//Unimplemented: CreateGeoLocTransformer
//Unimplemented: DestroyGeoLocTransformer
//Unimplemented: GeoLocTransform

//Unimplemented: CreateApproxTransformer
//Unimplemented: DestroyApproxTransformer
//Unimplemented: ApproxTransform

//Unimplemented: SimpleImageWarp
//Unimplemented: SuggestedWarpOutput
//Unimplemented: SuggsetedWarpOutput2
//Unimplemented: SerializeTransformer
//Unimplemented: DeserializeTransformer

//Unimplemented: TransformGeolocations

/* --------------------------------------------- */
/* Contour line functions                        */
/* --------------------------------------------- */

//Unimplemented: CreateContourGenerator
//Unimplemented: FeedLine
//Unimplemented: Destroy
//Unimplemented: ContourWriter
//Unimplemented: ContourGenerate

/* --------------------------------------------- */
/* Rasterizer functions                          */
/* --------------------------------------------- */

// Burn geometries into raster
//Unimplmemented: RasterizeGeometries

// Burn geometries from the specified list of layers into the raster
//Unimplemented: RasterizeLayers

// Burn geometries from the specified list of layers into the raster
//Unimplemented: RasterizeLayersBuf

/* --------------------------------------------- */
/* Gridding functions                            */
/* --------------------------------------------- */

type GDALGridAlgorithm int

const (
	GGA_InverseDistancetoAPower                = GDALGridAlgorithm(C.GGA_InverseDistanceToAPower)
	GGA_MovingAverage                          = GDALGridAlgorithm(C.GGA_MovingAverage)
	GGA_NearestNeighbor                        = GDALGridAlgorithm(C.GGA_NearestNeighbor)
	GGA_MetricMinimum                          = GDALGridAlgorithm(C.GGA_MetricMinimum)
	GGA_MetricMaximum                          = GDALGridAlgorithm(C.GGA_MetricMaximum)
	GGA_MetricRange                            = GDALGridAlgorithm(C.GGA_MetricRange)
	GGA_MetricCount                            = GDALGridAlgorithm(C.GGA_MetricCount)
	GGA_MetricAverageDistance                  = GDALGridAlgorithm(C.GGA_MetricAverageDistance)
	GGA_MetricAverageDistancePts               = GDALGridAlgorithm(C.GGA_MetricAverageDistancePts)
	GGA_Linear                                 = GDALGridAlgorithm(C.GGA_Linear)
	GGA_InverseDistanceToAPowerNearestNeighbor = GDALGridAlgorithm(C.GGA_InverseDistanceToAPowerNearestNeighbor)
)

// CPLErr GDALGridCreate(
// - GDALGridAlgorithm eAlgorithm,
// - const void *poOptions,
// - GUInt32 nPoints,
// - const double *padfX,
// - const double *padfY,
// - const double *padfZ,
// - double dfXMin,
// - double dfXMax,
// - double dfYMin,
// - double dfYMax,
// - GUInt32 nXSize,
// - GUInt32 nYSize,
// - GDALDataType eType,
// - void *pData,
// - GDALProgressFunc pfnProgress,
// - void *pProgressArg
// )
func CreateGrid(
	algo GDALGridAlgorithm,
	options []string,
	x, y, z []float64,
	nX, nY uint,
	buffer interface{}, // should be pre-initialized slice!
	progress ProgressFunc,
	data interface{},
) error {
	// options
	length := len(options)
	ooptions := make([]*C.char, length+1)
	for i := 0; i < length; i++ {
		ooptions[i] = C.CString(options[i])
		defer C.free(unsafe.Pointer(ooptions[i]))
	}
	ooptions[length] = (*C.char)(unsafe.Pointer(nil))

	if len(x) != len(y) || len(x) != len(z) {
		return errors.New("lengths of x, y, z should equal")
	}

	var lx, hx, ly, hy = math.MaxFloat64, -math.MaxFloat64, math.MaxFloat64, -math.MaxFloat64
	for i := range x {
		if x[i] < lx {
			lx = x[i]
		}
		if x[i] > hx {
			hx = x[i]
		}
		if y[i] < ly {
			ly = y[i]
		}
		if y[i] > hy {
			hy = y[i]
		}
	}

	dataType, dataPtr, err := determineBufferType(buffer)
	if err != nil {
		return err
	}

	arg := &goGDALProgressFuncProxyArgs{progress, data}

	return C.GDALGridCreate(
		C.GDALGridAlgorithm(algo),
		unsafe.Pointer(&ooptions[0]),
		C.uint(uint(len(x))),
		(*C.double)(unsafe.Pointer(&x[0])),
		(*C.double)(unsafe.Pointer(&y[0])),
		(*C.double)(unsafe.Pointer(&z[0])),
		C.double(lx),
		C.double(hx),
		C.double(ly),
		C.double(hy),
		C.uint(nX),
		C.uint(nY),
		C.GDALDataType(dataType),
		dataPtr,
		C.goGDALProgressFuncProxyB(),
		unsafe.Pointer(arg),
	).Err()
}

//Unimplemented: ComputeMatchingPoints

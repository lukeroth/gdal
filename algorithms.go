package gdal

/*
#include "go_gdal.h"
#include "gdal_version.h"
*/
import "C"
import (
	"errors"
	"fmt"
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

// GridAlgorithm represents Grid Algorithm code
type GridAlgorithm int

//
const (
	GA_InverseDistancetoAPower                = GridAlgorithm(C.GGA_InverseDistanceToAPower)
	GA_MovingAverage                          = GridAlgorithm(C.GGA_MovingAverage)
	GA_NearestNeighbor                        = GridAlgorithm(C.GGA_NearestNeighbor)
	GA_MetricMinimum                          = GridAlgorithm(C.GGA_MetricMinimum)
	GA_MetricMaximum                          = GridAlgorithm(C.GGA_MetricMaximum)
	GA_MetricRange                            = GridAlgorithm(C.GGA_MetricRange)
	GA_MetricCount                            = GridAlgorithm(C.GGA_MetricCount)
	GA_MetricAverageDistance                  = GridAlgorithm(C.GGA_MetricAverageDistance)
	GA_MetricAverageDistancePts               = GridAlgorithm(C.GGA_MetricAverageDistancePts)
	GA_Linear                                 = GridAlgorithm(C.GGA_Linear)
	GA_InverseDistanceToAPowerNearestNeighbor = GridAlgorithm(C.GGA_InverseDistanceToAPowerNearestNeighbor)
)

// GridLinearOptions: Linear method control options.
type GridLinearOptions struct {
	// Radius: in case the point to be interpolated does not fit into a triangle of the Delaunay triangulation,
	// use that maximum distance to search a nearest neighbour, or use nodata otherwise. If set to -1, the search
	// distance is infinite. If set to 0, nodata value will be always used.
	Radius float64
	// NoDataValue: no data marker to fill empty points.
	NoDataValue float64
}

// GridInverseDistanceToAPowerOptions: Inverse distance to a power method control options.
type GridInverseDistanceToAPowerOptions struct {
	// Power: Weighting power
	Power float64
	// Smoothing: Smoothing parameter
	Smoothing float64
	// AnisotropyRatio: Reserved for future use
	AnisotropyRatio float64
	// AnisotropyAngle: Reserved for future use
	AnisotropyAngle float64
	// Radius1: The first radius (X axis if rotation angle is 0) of search ellipse.
	Radius1 float64
	// Radius2: The second radius (Y axis if rotation angle is 0) of search ellipse.
	Radius2 float64
	// Angle: Angle of ellipse rotation in degrees. Ellipse rotated counter clockwise.
	Angle float64
	// MaxPoints: Maximum number of data points to use.
	// Do not search for more points than this number. If less amount of points found the grid node
	// considered empty and will be filled with NODATA marker.
	MaxPoints uint32
	// MinPoints: Minimum number of data points to use.
	// If less amount of points found the grid node considered empty and will be filled with NODATA marker.
	MinPoints uint32
	// NoDataValue: No data marker to fill empty points.
	NoDataValue float64
}

// GridInverseDistanceToAPowerNearestNeighborOptions: Inverse distance to a power, with nearest neighbour search,
// control options
type GridInverseDistanceToAPowerNearestNeighborOptions struct {
	// Power: Weighting power
	Power float64
	// Radius: The radius of search circle
	Radius float64
	// Smoothing: Smoothing parameter
	Smoothing float64
	// MaxPoints: Maximum number of data points to use.
	// Do not search for more points than this number. If less amount of points found the grid node
	// considered empty and will be filled with NODATA marker.
	MaxPoints uint32
	// MinPoints: Minimum number of data points to use.
	// If less amount of points found the grid node considered empty and will be filled with NODATA marker.
	MinPoints uint32
	// NoDataValue: No data marker to fill empty points.
	NoDataValue float64
}

// GridMovingAverageOptions: Moving average method control options
type GridMovingAverageOptions struct {
	// Radius1: The first radius (X axis if rotation angle is 0) of search ellipse.
	Radius1 float64
	// Radius2: The second radius (Y axis if rotation angle is 0) of search ellipse.
	Radius2 float64
	// Angle: Angle of ellipse rotation in degrees. Ellipse rotated counter clockwise.
	Angle float64
	// MinPoints: Minimum number of data points to use.
	// If less amount of points found the grid node considered empty and will be filled with NODATA marker.
	MinPoints uint32
	// NoDataValue: No data marker to fill empty points.
	NoDataValue float64
}

// GridNearestNeighborOptions: Nearest neighbor method control options.
type GridNearestNeighborOptions struct {
	// Radius1: The first radius (X axis if rotation angle is 0) of search ellipse.
	Radius1 float64
	// Radius2: The second radius (Y axis if rotation angle is 0) of search ellipse.
	Radius2 float64
	// Angle: Angle of ellipse rotation in degrees. Ellipse rotated counter clockwise.
	Angle float64
	// NoDataValue: No data marker to fill empty points.
	NoDataValue float64
}

// GridDataMetricsOptions: Data metrics method control options
type GridDataMetricsOptions struct {
	// Radius1: The first radius (X axis if rotation angle is 0) of search ellipse.
	Radius1 float64
	// Radius2: The second radius (Y axis if rotation angle is 0) of search ellipse.
	Radius2 float64
	// Angle: Angle of ellipse rotation in degrees. Ellipse rotated counter clockwise.
	Angle float64
	// MinPoints: Minimum number of data points to use.
	// If less amount of points found the grid node considered empty and will be filled with NODATA marker.
	MinPoints uint32
	// NoDataValue: No data marker to fill empty points.
	NoDataValue float64
}

var errInvalidOptionsTypeWasPassed = errors.New("invalid options type was passed")

// GridCreate: Create regular grid from the scattered data.
// This function takes the arrays of X and Y coordinates and corresponding Z values as input and computes
// regular grid (or call it a raster) from these scattered data. You should supply geometry and extent of the
// output grid.
func GridCreate(
	algorithm GridAlgorithm,
	options interface{},
	x, y, z []float64,
	xMin, xMax, yMin, yMax float64,
	nX, nY uint,
	progress ProgressFunc,
	data interface{},
) ([]float64, error) {
	if len(x) != len(y) || len(x) != len(z) {
		return nil, errors.New("lengths of x, y, z should equal")
	}

	poptions := unsafe.Pointer(nil)
	switch algorithm {
	case GA_InverseDistancetoAPower:
		soptions, ok := options.(GridInverseDistanceToAPowerOptions)
		if !ok {
			return nil, errInvalidOptionsTypeWasPassed
		}
		poptions = unsafe.Pointer(&C.GDALGridInverseDistanceToAPowerOptions{
			dfPower:           C.double(soptions.Power),
			dfSmoothing:       C.double(soptions.Smoothing),
			dfAnisotropyRatio: C.double(soptions.AnisotropyRatio),
			dfAnisotropyAngle: C.double(soptions.AnisotropyAngle),
			dfRadius1:         C.double(soptions.Radius1),
			dfRadius2:         C.double(soptions.Radius2),
			dfAngle:           C.double(soptions.Angle),
			nMaxPoints:        C.uint(soptions.MaxPoints),
			nMinPoints:        C.uint(soptions.MinPoints),
			dfNoDataValue:     C.double(soptions.NoDataValue),
		})
	case GA_InverseDistanceToAPowerNearestNeighbor:
		soptions, ok := options.(GridInverseDistanceToAPowerNearestNeighborOptions)
		if !ok {
			return nil, errInvalidOptionsTypeWasPassed
		}
		poptions = unsafe.Pointer(&C.GDALGridInverseDistanceToAPowerNearestNeighborOptions{
			dfPower:       C.double(soptions.Power),
			dfRadius:      C.double(soptions.Radius),
			dfSmoothing:   C.double(soptions.Smoothing),
			nMaxPoints:    C.uint(soptions.MaxPoints),
			nMinPoints:    C.uint(soptions.MinPoints),
			dfNoDataValue: C.double(soptions.NoDataValue),
		})
	case GA_MovingAverage:
		soptions, ok := options.(GridMovingAverageOptions)
		if !ok {
			return nil, errInvalidOptionsTypeWasPassed
		}
		poptions = unsafe.Pointer(&C.GDALGridMovingAverageOptions{
			dfRadius1:     C.double(soptions.Radius1),
			dfRadius2:     C.double(soptions.Radius2),
			dfAngle:       C.double(soptions.Angle),
			nMinPoints:    C.uint(soptions.MinPoints),
			dfNoDataValue: C.double(soptions.NoDataValue),
		})
	case GA_NearestNeighbor:
		soptions, ok := options.(GridNearestNeighborOptions)
		if !ok {
			return nil, errInvalidOptionsTypeWasPassed
		}
		poptions = unsafe.Pointer(&C.GDALGridNearestNeighborOptions{
			dfRadius1:     C.double(soptions.Radius1),
			dfRadius2:     C.double(soptions.Radius2),
			dfAngle:       C.double(soptions.Angle),
			dfNoDataValue: C.double(soptions.NoDataValue),
		})
	case GA_MetricMinimum, GA_MetricMaximum, GA_MetricCount, GA_MetricRange,
		GA_MetricAverageDistance, GA_MetricAverageDistancePts:
		soptions, ok := options.(GridDataMetricsOptions)
		if !ok {
			return nil, errInvalidOptionsTypeWasPassed
		}
		poptions = unsafe.Pointer(&C.GDALGridDataMetricsOptions{
			dfRadius1:     C.double(soptions.Radius1),
			dfRadius2:     C.double(soptions.Radius2),
			dfAngle:       C.double(soptions.Angle),
			nMinPoints:    C.uint(soptions.MinPoints),
			dfNoDataValue: C.double(soptions.NoDataValue),
		})
	case GA_Linear:
		soptions, ok := options.(GridLinearOptions)
		if !ok {
			return nil, errInvalidOptionsTypeWasPassed
		}
		poptions = unsafe.Pointer(&C.GDALGridLinearOptions{
			dfRadius:      C.double(soptions.Radius),
			dfNoDataValue: C.double(soptions.NoDataValue),
		})
	}

	buffer := make([]float64, nX*nY)
	arg := &goGDALProgressFuncProxyArgs{progress, data}
	err := C.GDALGridCreate(
		C.GDALGridAlgorithm(algorithm),
		poptions,
		C.uint(uint(len(x))),
		(*C.double)(unsafe.Pointer(&x[0])),
		(*C.double)(unsafe.Pointer(&y[0])),
		(*C.double)(unsafe.Pointer(&z[0])),
		C.double(xMin),
		C.double(xMax),
		C.double(yMin),
		C.double(yMax),
		C.uint(nX),
		C.uint(nY),
		C.GDALDataType(Float64),
		unsafe.Pointer(&buffer[0]),
		C.goGDALProgressFuncProxyB(),
		unsafe.Pointer(arg),
	).Err()
	return buffer, err
}

//Unimplemented: ComputeMatchingPoints

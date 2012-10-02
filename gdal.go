package gdal

/*
#include "go_gdal.h"
#include "gdal_version.h"

#cgo linux  pkg-config: gdal
#cgo darwin pkg-config: gdal
#cgo windows LDFLAGS: -lgdal.dll
*/
import "C"
import (
	"unsafe"
	"fmt"
)
var _ = fmt.Println

func init() {
	C.GDALAllRegister()
}

/* -------------------------------------------------------------------- */
/*      Significant constants.                                          */
/* -------------------------------------------------------------------- */

const (
	VERSION_MAJOR	= int(C.GDAL_VERSION_MAJOR)
	VERSION_MINOR	= int(C.GDAL_VERSION_MINOR)
	VERSION_REV	= int(C.GDAL_VERSION_REV)
	VERSION_BUILD	= int(C.GDAL_VERSION_BUILD)
	VERSION_NUM	= int(C.GDAL_VERSION_NUM)
	RELEASE_DATE	= int(C.GDAL_RELEASE_DATE)
	RELEASE_NAME	= string(C.GDAL_RELEASE_NAME)
)

// Error handling.  The following is bare-bones, and needs to be replaced with something more useful.
func (err _Ctype_CPLErr) Error() string {
	switch err {
	case 0:
		return "No Error"
	case 1:
		return "Debug Error"
	case 2:
		return "Warning Error"
	case 3:
		return "Failure Error"
	case 4:
		return "Fatal Error"
	}
	return "Illegal error"
}

func (err _Ctype_OGRErr) Error() string {
	switch err {
	case 0:
		return "No Error"
	case 1:
		return "Debug Error"
	case 2:
		return "Warning Error"
	case 3:
		return "Failure Error"
	case 4:
		return "Fatal Error"
	}
	return "Illegal error"
}

// Pixel data types
type DataType int

const (
	Unknown		= DataType(C.GDT_Unknown)
	Byte		= DataType(C.GDT_Byte)
	UInt16		= DataType(C.GDT_UInt16)
	Int16		= DataType(C.GDT_Int16)
	UInt32		= DataType(C.GDT_UInt32)
	Int32		= DataType(C.GDT_Int32)
	Float32		= DataType(C.GDT_Float32)
	Float64		= DataType(C.GDT_Float64)
	CInt16		= DataType(C.GDT_CInt16)
	CInt32		= DataType(C.GDT_CInt32)
	CFloat32	= DataType(C.GDT_CFloat32)
	CFloat64	= DataType(C.GDT_CFloat64)
)

// Get data type size in bits.
func (dataType DataType) Size() int {
	return int(C.GDALGetDataTypeSize(C.GDALDataType(dataType)))
}

func (dataType DataType) IsComplex() int {
	return int(C.GDALDataTypeIsComplex(C.GDALDataType(dataType)))
}

func (dataType DataType) Name() string {
	return C.GoString(C.GDALGetDataTypeName(C.GDALDataType(dataType)))
}

func (dataType DataType) Union(dataTypeB DataType) DataType {
	return DataType(
		C.GDALDataTypeUnion(C.GDALDataType(dataType), C.GDALDataType(dataTypeB)),
	)
}

// status of the asynchronous stream
type AsyncStatusType int

const (
	AR_Pending   = AsyncStatusType(C.GARIO_PENDING)
	AR_Update    = AsyncStatusType(C.GARIO_UPDATE)
	AR_Error     = AsyncStatusType(C.GARIO_ERROR)
	AR_Complete  = AsyncStatusType(C.GARIO_COMPLETE)
)

func (statusType AsyncStatusType) Name() string {
	return C.GoString(C.GDALGetAsyncStatusTypeName(C.GDALAsyncStatusType(statusType)))
}

func GetAsyncStatusTypeByName(statusTypeName string) AsyncStatusType {
	name := C.CString(statusTypeName)
	defer C.free(unsafe.Pointer(name))
	return AsyncStatusType(C.GDALGetAsyncStatusTypeByName(name))
}

// Flag indicating read/write, or read-only access to data.
type Access int

const (
	// Read only (no update) access
	ReadOnly = Access(C.GA_ReadOnly)
	// Read/write access.
	Update = Access(C.GA_Update)
)

// Read/Write flag for RasterIO() method
type RWFlag int

const (
	// Read data
	Read = RWFlag(C.GF_Read)
	// Write data
	Write = RWFlag(C.GF_Write)
)

// Types of color interpretation for raster bands.
type ColorInterp int

const (
	CI_Undefined		= ColorInterp(C.GCI_Undefined)
	CI_GrayIndex		= ColorInterp(C.GCI_GrayIndex)
	CI_PaletteIndex	= ColorInterp(C.GCI_PaletteIndex)
	CI_RedBand			= ColorInterp(C.GCI_RedBand)
	CI_GreenBand		= ColorInterp(C.GCI_GreenBand)
	CI_BlueBand		= ColorInterp(C.GCI_BlueBand)
	CI_AlphaBand		= ColorInterp(C.GCI_AlphaBand)
	CI_HueBand			= ColorInterp(C.GCI_HueBand)
	CI_SaturationBand	= ColorInterp(C.GCI_SaturationBand)
	CI_LightnessBand	= ColorInterp(C.GCI_LightnessBand)
	CI_CyanBand		= ColorInterp(C.GCI_CyanBand)
	CI_MagentaBand		= ColorInterp(C.GCI_MagentaBand)
	CI_YellowBand		= ColorInterp(C.GCI_YellowBand)
	CI_BlackBand		= ColorInterp(C.GCI_BlackBand)
	CI_YCbCr_YBand		= ColorInterp(C.GCI_YCbCr_YBand)
	CI_YCbCr_CbBand	= ColorInterp(C.GCI_YCbCr_CbBand)
	CI_YCbCr_CrBand	= ColorInterp(C.GCI_YCbCr_CrBand)
	CI_Max				= ColorInterp(C.GCI_Max)
)

func (colorInterp ColorInterp) Name() string {
	return C.GoString(C.GDALGetColorInterpretationName(C.GDALColorInterp(colorInterp)))
}

func GetColorInterpretationByName(name string) ColorInterp {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	return ColorInterp(C.GDALGetColorInterpretationByName(cName))
}

// Types of color interpretations for a GDALColorTable.
type PaletteInterp int

const (
	// Grayscale (in GDALColorEntry.c1)
	PI_Gray	= PaletteInterp(C.GPI_Gray)
	// Red, Green, Blue and Alpha in (in c1, c2, c3 and c4)
	PI_RGB		= PaletteInterp(C.GPI_RGB)
	// Cyan, Magenta, Yellow and Black (in c1, c2, c3 and c4)
	PI_CMYK	= PaletteInterp(C.GPI_CMYK)
	// Hue, Lightness and Saturation (in c1, c2, and c3)
	PI_HLS		= PaletteInterp(C.GPI_HLS)
)

func (paletteInterp PaletteInterp) Name() string {
	return C.GoString(C.GDALGetPaletteInterpretationName(C.GDALPaletteInterp(paletteInterp)))
}

// "well known" metadata items.
const (
	MD_AREA_OR_POINT = string(C.GDALMD_AREA_OR_POINT)
	MD_AOP_AREA      = string(C.GDALMD_AOP_AREA)
	MD_AOP_POINT     = string(C.GDALMD_AOP_POINT)
)

/* -------------------------------------------------------------------- */
/*      Define handle types related to various internal classes.        */
/* -------------------------------------------------------------------- */

type MajorObject struct {
	cval C.GDALMajorObjectH
}

type Dataset struct {
	cval C.GDALDatasetH
}

type RasterBand struct {
	cval C.GDALRasterBandH
}

type Driver struct {
	cval C.GDALDriverH
}

type ColorTable struct {
	cval C.GDALColorTableH
}

type RasterAttributeTable struct {
	cval C.GDALRasterAttributeTableH
}

type AsyncReader struct {
	cval C.GDALAsyncReaderH
}

type ColorEntry struct {
	cval C.GDALColorEntry
}

/* -------------------------------------------------------------------- */
/*      Callback "progress" function.                                   */
/* -------------------------------------------------------------------- */

type ProgressFunc func(complete float64, message string, progressArg interface{}) int

func DummyProgress(complete float64, message string, data interface{}) int {
	msg := C.CString(message)
	defer C.free(unsafe.Pointer(msg))

	retval := C.GDALDummyProgress(C.double(complete), msg, unsafe.Pointer(nil))
	return int(retval)
}

func TermProgress(complete float64, message string, data interface{}) int {
	msg := C.CString(message)
	defer C.free(unsafe.Pointer(msg))

	retval := C.GDALTermProgress(C.double(complete), msg, unsafe.Pointer(nil))
	return int(retval)
}

func ScaledProgress(complete float64, message string, data interface{}) int {
	msg := C.CString(message)
	defer C.free(unsafe.Pointer(msg))

	retval := C.GDALScaledProgress(C.double(complete), msg, unsafe.Pointer(nil))
	return int(retval)
}

func CreateScaledProgress(min, max float64, progress ProgressFunc, data unsafe.Pointer) unsafe.Pointer {
	panic("not implemented!")
	return nil
}

func DestroyScaledProgress(data unsafe.Pointer) {
	C.GDALDestroyScaledProgress(data)
}

// -----------------------------------------------------------------------

type goGDALProgressFuncProxyArgs struct {
	progresssFunc ProgressFunc
	data         interface{}
}

//export goGDALProgressFuncProxyA
func goGDALProgressFuncProxyA(complete C.double, message *C.char, data *interface{}) int {
	if arg, ok := (*data).(goGDALProgressFuncProxyArgs); ok {
		return arg.progresssFunc(
			float64(complete), C.GoString(message), arg.data,
		)
	}
	return 0
}

/* ==================================================================== */
/*      Registration/driver related.                                    */
/* ==================================================================== */

const (
	DMD_LONGNAME			= string(C.GDAL_DMD_LONGNAME)
	DMD_HELPTOPIC			= string(C.GDAL_DMD_HELPTOPIC)
	DMD_MIMETYPE			= string(C.GDAL_DMD_MIMETYPE)
	DMD_EXTENSION			= string(C.GDAL_DMD_EXTENSION)
	DMD_CREATIONOPTIONLIST	= string(C.GDAL_DMD_CREATIONOPTIONLIST)
	DMD_CREATIONDATATYPES 	= string(C.GDAL_DMD_CREATIONDATATYPES)

	DCAP_CREATE     		= string(C.GDAL_DCAP_CREATE)
	DCAP_CREATECOPY			= string(C.GDAL_DCAP_CREATECOPY)
	DCAP_VIRTUALIO			= string(C.GDAL_DCAP_VIRTUALIO)
)

// Create a new dataset with this driver.
func (driver Driver) Create(
	filename string,
	xSize, ySize, bands int,
	dataType DataType,
	options []string,
) Dataset {
	name := C.CString(filename)
	defer C.free(unsafe.Pointer(name))

	length := len(options)
	opts := make([]*C.char, length + 1)
	for i := 0; i < length; i++ {
		opts[i] = C.CString(options[i])
		defer C.free(unsafe.Pointer(opts[i]))
	}
	opts[length] = (*C.char)(unsafe.Pointer(nil))

	h := C.GDALCreate(
		driver.cval,
		name,
		C.int(xSize), C.int(ySize), C.int(bands),
		C.GDALDataType(dataType),
		(**C.char)(unsafe.Pointer(&opts[0])),
	)
	return Dataset{h}
}

// Create a copy of a dataset.
func (driver Driver) CreateCopy(
	filename string,
	sourceDataset Dataset,
	strict int, 
	options []string,
	progress ProgressFunc,
	data interface{},
) Dataset {
	name := C.CString(filename)
	defer C.free(unsafe.Pointer(name))

	length := len(options)
	opts := make([]*C.char, length + 1)
	for i := 0; i < length; i++ {
		opts[i] = C.CString(options[i])
		defer C.free(unsafe.Pointer(opts[i]))
	}
	opts[length] = (*C.char)(unsafe.Pointer(nil))

	arg := &goGDALProgressFuncProxyArgs{
		progress, data,
	}

	h := C.GDALCreateCopy(
		driver.cval, name,
		sourceDataset.cval,
		C.int(strict), (**C.char)(unsafe.Pointer(&opts[0])),
		C.goGDALProgressFuncProxyB(),
		unsafe.Pointer(arg),
	)
	return Dataset{h}
}

// Return the driver needed to access the provided dataset name.
func IdentifyDriver(filename string, filenameList []string) Driver {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))

	length := len(filenameList)
	cFilenameList := make([]*C.char, length + 1)
	for i := 0; i < length; i++ {
		cFilenameList[i] = C.CString(filenameList[i])
		defer C.free(unsafe.Pointer(cFilenameList[i]))
	}
	cFilenameList[length] = (*C.char)(unsafe.Pointer(nil))

	driver := C.GDALIdentifyDriver(cFilename, (**C.char)(unsafe.Pointer(&cFilenameList[0])))
	return Driver{driver}
}

// Open an existing dataset
func Open(filename string, access Access) Dataset {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))

	dataset := C.GDALOpen(cFilename, C.GDALAccess(access))
	return Dataset{dataset}
}

// Open a shared existing dataset
func OpenShared(filename string, access Access) Dataset {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))

	dataset := C.GDALOpenShared(cFilename, C.GDALAccess(access))
	return Dataset{dataset}
}

// Unimplemented: DumpOpenDatasets

// Return the driver by short name
func GetDriverByName(driverName string) (Driver, error) {
	cName := C.CString(driverName)
	defer C.free(unsafe.Pointer(cName))

	driver := C.GDALGetDriverByName(cName)
	if driver == nil {
		return Driver{driver}, fmt.Errorf("Error: driver '%s' not found", driverName)
	}
	return Driver{driver}, nil
}

// Fetch the number of registered drivers.
func GetDriverCount() int {
	nDrivers := C.GDALGetDriverCount()
	return int(nDrivers)
}

// Fetch driver by index
func GetDriver(index int) Driver {
	driver := C.GDALGetDriver(C.int(index))
	return Driver{driver}
}

// Destroy a GDAL driver
func (driver Driver) Destroy() {
	C.GDALDestroyDriver(driver.cval)
}

// Registers a driver for use
func (driver Driver) Register() int {
	index := C.GDALRegisterDriver(driver.cval)
	return int(index)
}

// Reregister the driver
func (driver Driver) Deregister() {
	C.GDALDeregisterDriver(driver.cval)
}

// Destroy the driver manager
func DestroyDriverManager() {
	C.GDALDestroyDriverManager()
}

// Delete named dataset
func (driver Driver) DeleteDataset(name string) error {
	cDriver := driver.cval
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	err := C.GDALDeleteDataset(cDriver, cName)
	return error(err)
}

// Rename named dataset
func (driver Driver) RenameDataset(newName, oldName string) error {
	cDriver := driver.cval
	cNewName := C.CString(newName)
	defer C.free(unsafe.Pointer(cNewName))
	cOldName := C.CString(oldName)
	defer C.free(unsafe.Pointer(cOldName))
	err := C.GDALRenameDataset(cDriver, cNewName, cOldName)
	return error(err)
}

// Copy all files associated with the named dataset
func (driver Driver) CopyDatasetFiles(newName, oldName string) error {
	cDriver := driver.cval
	cNewName := C.CString(newName)
	defer C.free(unsafe.Pointer(cNewName))
	cOldName := C.CString(oldName)
	defer C.free(unsafe.Pointer(cOldName))
	err := C.GDALCopyDatasetFiles(cDriver, cNewName, cOldName)
	return error(err)
}

/* ==================================================================== */
/*      GDAL_GCP                                                        */
/* ==================================================================== */

// Unimplemented: InitGCPs
// Unimplemented: DeinitGCPs
// Unimplemented: DuplicateGCPs
// Unimplemented: GCPsToGeoTransform
// Unimplemented: InvGeoTransform
// Unimplemented: ApplyGeoTransform

/* ==================================================================== */
/*      major objects (dataset, and, driver, drivermanager).            */
/* ==================================================================== */

// Fetch object description
func (object MajorObject) Description() string {
	cObject := object.cval
	desc := C.GoString(C.GDALGetDescription(cObject))
	return desc
}

// Set object description
func (object MajorObject) SetDescription(desc string) {
	cObject := object.cval
	cDesc := C.CString(desc)
	defer C.free(unsafe.Pointer(cDesc))
	C.GDALSetDescription(cObject, cDesc)
}

// Fetch metadata
func (object MajorObject) Metadata(domain string) []string {
	panic("not implemented!")
	return nil
}

// Set metadata
func (object MajorObject) SetMetadata(metadata []string, domain string) {
	panic("not implemented!")
	return
}

// Fetch a single metadata item
func (object MajorObject) MetadataItem(name, domain string) string {
	panic("not implemented!")
	return ""
}

// Set a single metadata item
func (object MajorObject) SetMetadataItem(name, value, domain string) {
	panic("not implemented!")
	return
}

/* ==================================================================== */
/*      GDALDataset class ... normally this represents one file.        */
/* ==================================================================== */

// Get the driver to which this dataset relates
func (dataset Dataset) Driver() Driver {
	driver := Driver{C.GDALGetDatasetDriver(dataset.cval)}
	return driver
}

// Unimplemented: GDALGetFileList

// Close the dataset
func (dataset Dataset) Close() {
	C.GDALClose(dataset.cval)
	return
}

// Fetch X size of raster
func (dataset Dataset) RasterXSize() int {
	xSize := int(C.GDALGetRasterXSize(dataset.cval))
	return xSize
}

// Fetch Y size of raster
func (dataset Dataset) RasterYSize() int {
	ySize := int(C.GDALGetRasterYSize(dataset.cval))
	return ySize
}

// Fetch the number of raster bands in the dataset
func (dataset Dataset) RasterCount() int {
	count := int(C.GDALGetRasterCount(dataset.cval))
	return count
}

// Fetch a raster band object from a dataset
func (dataset Dataset) RasterBand(band int) RasterBand {
	rasterBand := RasterBand{C.GDALGetRasterBand(dataset.cval, C.int(band))}
	return rasterBand
}

// Add a band to a dataset
func (dataset Dataset) AddBand(dataType DataType, options []string) error {
	length := len(options)
	cOptions := make([]*C.char, length + 1)
	for i := 0; i < length; i++ {
		cOptions[i] = C.CString(options[i])
		defer C.free(unsafe.Pointer(cOptions[i]))
	}
	cOptions[length] = (*C.char)(unsafe.Pointer(nil))

	err := C.GDALAddBand(
		dataset.cval, 
		C.GDALDataType(dataType), 		
		(**C.char)(unsafe.Pointer(&cOptions[0])))

	return error(err)
}

// Unimplemented: GDALBeginAsyncReader
// Unimplemented: GDALEndAsyncReader

// Read / write a region of image data from multiple bands	  
func (dataset Dataset) IO(
	rwFlag RWFlag,
	xOff, yOff, xSize, ySize int,
	buffer interface{},
	bufXSize, bufYSize int,
	bandCount int,
	bandMap []int,
	pixelSpace, lineSpace, bandSpace int,
) error {
	var dataType DataType
	var dataPtr unsafe.Pointer
	switch data := buffer.(type) {
	case []int8:
		dataType = Byte
		dataPtr = unsafe.Pointer(&data[0])
	case []uint8:
		dataType = Byte
		dataPtr = unsafe.Pointer(&data[0])
	case []int16:
		dataType = Int16
		dataPtr = unsafe.Pointer(&data[0])
	case []uint16:
		dataType = UInt16
		dataPtr = unsafe.Pointer(&data[0])
	case []int32:
		dataType = Int32
		dataPtr = unsafe.Pointer(&data[0])
	case []uint32:
		dataType = UInt32
		dataPtr = unsafe.Pointer(&data[0])
	default:
		return fmt.Errorf("Error: buffer is not a valid data type (must be a valid numeric slice)")
	}

	err := C.GDALDatasetRasterIO(
		dataset.cval,
		C.GDALRWFlag(rwFlag),
		C.int(xOff), C.int(yOff), C.int(xSize), C.int(ySize),
		dataPtr,
		C.int(bufXSize), C.int(bufYSize),
		C.GDALDataType(dataType),
		C.int(bandCount),
		(*C.int)(unsafe.Pointer(&bandMap[0])),
		C.int(pixelSpace), C.int(lineSpace), C.int(bandSpace))
	return error(err)
}

// Advise driver of upcoming read requests
func (dataset Dataset) AdviseRead(
	rwFlag RWFlag,
	xOff, yOff, xSize, ySize, bufXSize, bufYSize int,
	dataType DataType,
	bandCount int,
	bandMap []int,
	options []string,
) error {
	length := len(options)
	cOptions := make([]*C.char, length + 1)
	for i := 0; i < length; i++ {
		cOptions[i] = C.CString(options[i])
		defer C.free(unsafe.Pointer(cOptions[i]))
	}
	cOptions[length] = (*C.char)(unsafe.Pointer(nil))

	err := C.GDALDatasetAdviseRead(
		dataset.cval,
		C.int(xOff), C.int(yOff), C.int(xSize), C.int(ySize),
		C.int(bufXSize), C.int(bufYSize),
		C.GDALDataType(dataType),
		C.int(bandCount),
		(*C.int)(unsafe.Pointer(&bandMap[0])),
		(**C.char)(unsafe.Pointer(&cOptions[0])))
	return error(err)
}

// Fetch the projection definition string for this dataset
func (dataset Dataset) Projection() string {
	proj := C.GoString(C.GDALGetProjectionRef(dataset.cval))
	return proj
}

// Set the projection reference string
func (dataset Dataset) SetProjection(proj string) error {
	cProj := C.CString(proj)
	defer C.free(unsafe.Pointer(cProj))

	err := C.GDALSetProjection(dataset.cval, cProj)
	return error(err)
}

// Get the affine transformation coefficients
func (dataset Dataset) GeoTransform() []float64 {
	var transform []float64
	C.GDALGetGeoTransform(dataset.cval, (*C.double)(unsafe.Pointer(&transform[0])))
	return transform
}

// Set the affine transformation coefficients
func (dataset Dataset) SetGeoTransform(transform []float64) error {
	err := C.GDALSetGeoTransform(dataset.cval, (*C.double)(unsafe.Pointer(&transform[0])))
	return error(err)
}

// Get number of GCPs
func (dataset Dataset) GDALGetGCPCount() int {
	count := C.GDALGetGCPCount(dataset.cval)
	return int(count)
}

// Unimplemented: GDALGetGCPProjection
// Unimplemented: GDALGetGCPs
// Unimplemented: GDALSetGCPs

// Fetch a format specific internally meaningful handle
func (dataset Dataset) GDALGetInternalHandle(request string) unsafe.Pointer {
	cRequest := C.CString(request)
	defer C.free(unsafe.Pointer(cRequest))
	
	ptr := C.GDALGetInternalHandle(dataset.cval, cRequest)
	return ptr
}

// Add one to dataset reference count
func (dataset Dataset) GDALReferenceDataset() int {
	count := C.GDALReferenceDataset(dataset.cval)
	return int(count)
}

// Subtract one from dataset reference count
func (dataset Dataset) GDALDereferenceDataset() int {
	count := C.GDALDereferenceDataset(dataset.cval)
	return int(count)
}

// Build raster overview(s)
func (dataset Dataset) BuildOverviews(
	resampling string,
	nOverviews int,
	overviewList []int,
	nBands int,
	bandList []int,
	progress ProgressFunc,
	data interface{},
) error {
	cResampling := C.CString(resampling)
	defer C.free(unsafe.Pointer(cResampling))

	arg := &goGDALProgressFuncProxyArgs{progress, data}

	err := C.GDALBuildOverviews(
		dataset.cval,
		cResampling,
		C.int(nOverviews),
		(*C.int)(unsafe.Pointer(&overviewList[0])),
		C.int(nBands),
		(*C.int)(unsafe.Pointer(&bandList[0])),
		C.goGDALProgressFuncProxyB(),
		unsafe.Pointer(arg),
	)
	return error(err)
}

// Unimplemented: GDALGetOpenDatasets

// Return access flag
func (dataset Dataset) Access() Access {
	accessVal := C.GDALGetAccess(dataset.cval)
	return Access(accessVal)
}

// Write all write cached data to disk
func (dataset Dataset) FlushCache() {
	C.GDALFlushCache(dataset.cval)
	return
}

// Adds a mask band to the dataset
func (dataset Dataset) CreateMaskBand(flags int) error {
	err := C.GDALCreateDatasetMaskBand(dataset.cval, C.int(flags))
	return error(err)
}

// Copy all dataset raster data
func (sourceDataset Dataset) CopyWholeRaster(
	destDataset Dataset,
	options []string,
	progress ProgressFunc,
	data interface{},
) error {
	arg := &goGDALProgressFuncProxyArgs{progress, data}

	length := len(options)
	cOptions := make([]*C.char, length + 1)
	for i := 0; i < length; i++ {
		cOptions[i] = C.CString(options[i])
		defer C.free(unsafe.Pointer(cOptions[i]))
	}
	cOptions[length] = (*C.char)(unsafe.Pointer(nil))

	err := C.GDALDatasetCopyWholeRaster(
		sourceDataset.cval,
		destDataset.cval,
		(**C.char)(unsafe.Pointer(&cOptions[0])),
		C.goGDALProgressFuncProxyB(),
		unsafe.Pointer(arg),
	)
	return error(err)
}

/* ==================================================================== */
/*      GDALRasterBand ... one band/channel in a dataset.               */
/* ==================================================================== */

// Fetch the pixel data type for this band
func (rasterBand RasterBand) RasterDataType() DataType {
	dataType := C.GDALGetRasterDataType(rasterBand.cval)
	return DataType(dataType)
}

// Fetch the "natural" block size of this band
func (rasterBand RasterBand) BlockSize() (int, int) {
	var xSize, ySize int
	C.GDALGetBlockSize(rasterBand.cval, (*C.int)(unsafe.Pointer(&xSize)), (*C.int)(unsafe.Pointer(&ySize)))
	return xSize, ySize
}

// Advise driver of upcoming read requests
func (rasterBand RasterBand) AdviseRead(
	xOff, yOff, xSize, ySize, bufXSize, bufYSize int,
	dataType DataType,
	options []string,
) error {
	length := len(options)
	cOptions := make([]*C.char, length + 1)
	for i := 0; i < length; i++ {
		cOptions[i] = C.CString(options[i])
		defer C.free(unsafe.Pointer(cOptions[i]))
	}
	cOptions[length] = (*C.char)(unsafe.Pointer(nil))

	err := C.GDALRasterAdviseRead(
		rasterBand.cval,
		C.int(xOff), C.int(yOff), C.int(xSize), C.int(ySize), C.int(bufXSize), C.int(bufYSize),
		C.GDALDataType(dataType),
		(**C.char)(unsafe.Pointer(&cOptions[0])),
	)
	return error(err)
}

// Read / Write a region of image data for this band
func (rasterBand RasterBand) IO(
	rwFlag RWFlag,
	xOff, yOff, xSize, ySize int,
	buffer interface{},
	bufXSize, bufYSize int,
	pixelSpace, lineSpace int,
) error {
	var dataType DataType
	var dataPtr unsafe.Pointer
	switch data := buffer.(type) {
	case []int8:
		dataType = Byte
		dataPtr = unsafe.Pointer(&data[0])
	case []uint8:
		dataType = Byte
		dataPtr = unsafe.Pointer(&data[0])
	case []int16:
		dataType = Int16
		dataPtr = unsafe.Pointer(&data[0])
	case []uint16:
		dataType = UInt16
		dataPtr = unsafe.Pointer(&data[0])
	case []int32:
		dataType = Int32
		dataPtr = unsafe.Pointer(&data[0])
	case []uint32:
		dataType = UInt32
		dataPtr = unsafe.Pointer(&data[0])
	default:
		return fmt.Errorf("Error: buffer is not a valid data type (must be a valid numeric slice)")
	}

	err := C.GDALRasterIO(
		rasterBand.cval,
		C.GDALRWFlag(rwFlag),
		C.int(xOff), C.int(yOff), C.int(xSize), C.int(ySize),
		dataPtr,
		C.int(bufXSize), C.int(bufYSize),
		C.GDALDataType(dataType),
		C.int(pixelSpace), C.int(lineSpace))
	return error(err)
}

// Read a block of image data efficiently
func (rasterBand RasterBand) ReadBlock(xOff, yOff int, dataPtr unsafe.Pointer) error {
	err := C.GDALReadBlock(rasterBand.cval, C.int(xOff), C.int(yOff), dataPtr)
	return error(err)
}

// Write a block of image data efficiently
func (rasterBand RasterBand) WriteBlock(xOff, yOff int, dataPtr unsafe.Pointer) error {
	err := C.GDALWriteBlock(rasterBand.cval, C.int(xOff), C.int(yOff), dataPtr)
	return error(err)
}

// Fetch X size of raster
func (rasterBand RasterBand) XSize() int {
	xSize := C.GDALGetRasterBandXSize(rasterBand.cval)
	return int(xSize)
}

// Fetch Y size of raster
func (rasterBand RasterBand) YSize() int {
	ySize := C.GDALGetRasterBandYSize(rasterBand.cval)
	return int(ySize)
}

// Find out if we have update permission for this band
func (rasterBand RasterBand) GetAccess() Access {
	access := C.GDALGetRasterAccess(rasterBand.cval)
	return Access(access)
}

// Fetch the band number of this raster band
func (rasterBand RasterBand) BandNumber() int {
	bandNumber := C.GDALGetBandNumber(rasterBand.cval)
	return int(bandNumber)
}

// Fetch the owning dataset handle
func (rasterBand RasterBand) GetDataset() Dataset {
	dataset := C.GDALGetBandDataset(rasterBand.cval)
	return Dataset{dataset}
}

// How should this band be interpreted as color?
func (rasterBand RasterBand) ColorInterp() ColorInterp {
	colorInterp := C.GDALGetRasterColorInterpretation(rasterBand.cval)
	return ColorInterp(colorInterp)
}

// Set color interpretation of the raster band
func (rasterBand RasterBand) SetColorInterp(colorInterp ColorInterp) error {
	err := C.GDALSetRasterColorInterpretation(rasterBand.cval, C.GDALColorInterp(colorInterp))
	return error(err)
}

// Fetch the color table associated with this raster band
func (rasterBand RasterBand) ColorTable() ColorTable {
	colorTable := C.GDALGetRasterColorTable(rasterBand.cval)
	return ColorTable{colorTable}
}

// Set the raster color table for this raster band
func (rasterBand RasterBand) SetColorTable(colorTable ColorTable) error {
	err := C.GDALSetRasterColorTable(rasterBand.cval, colorTable.cval)
	return error(err)
}

// Check for arbitrary overviews
func (rasterBand RasterBand) HasArbitraryOverviews() int {
	yes := C.GDALHasArbitraryOverviews(rasterBand.cval)
	return int(yes)
}

// Return the number of overview layers available
func (rasterBand RasterBand) OverviewCount() int {
	count := C.GDALGetOverviewCount(rasterBand.cval)
	return int(count)
}

// Fetch overview raster band object
func (rasterBand RasterBand) Overview(level int) RasterBand {
	overview := C.GDALGetOverview(rasterBand.cval, C.int(level))
	return RasterBand{overview}
}

// Fetch the no data value for this band
func (rasterBand RasterBand) NoDataValue() (val float64, valid bool) {
	var success int
	noDataVal := C.GDALGetRasterNoDataValue(rasterBand.cval, (*C.int)(unsafe.Pointer(&success)))
	return float64(noDataVal), (success != 0)
}

// Set the no data value for this band
func (rasterBand RasterBand) SetNoDataValue(val float64) error {
	err := C.GDALSetRasterNoDataValue(rasterBand.cval, C.double(val))
	return error(err)
}

// Fetch the list of category names for this raster
func (rasterBand RasterBand) CategoryNames() []string {
	p := C.GDALGetRasterCategoryNames(rasterBand.cval)
	var strings []string
	q := uintptr(unsafe.Pointer(p))
	for {
		p = (**C.char)(unsafe.Pointer(q))
		if *p == nil {
			break
		}
		strings = append(strings, C.GoString(*p))
		q += unsafe.Sizeof(q)
	}
	
	return strings
}

// Set the category names for this band
func (rasterBand RasterBand) SetRasterCategoryNames(names []string) error {
	length := len(names)
	cStrings := make([]*C.char, length + 1)
	for i := 0; i < length; i++ {
		cStrings[i] = C.CString(names[i])
		defer C.free(unsafe.Pointer(cStrings[i]))
	}
	cStrings[length] = (*C.char)(unsafe.Pointer(nil))

	err := C.GDALSetRasterCategoryNames(rasterBand.cval, (**C.char)(unsafe.Pointer(&cStrings[0])))
	
	return error(err)
}

// Fetch the minimum value for this band
func (rasterBand RasterBand) GetMinimum() (val float64, valid bool) {
	var success int
	min := C.GDALGetRasterMinimum(rasterBand.cval, (*C.int)(unsafe.Pointer(&success)))
	return float64(min), (success != 0)
}

// Fetch the maximum value for this band
func (rasterBand RasterBand) GetMaximum() (val float64, valid bool) {
	var success int
	max := C.GDALGetRasterMaximum(rasterBand.cval, (*C.int)(unsafe.Pointer(&success)))
	return float64(max), (success != 0)
}

// Fetch image statistics
func (rasterBand RasterBand) GetStatistics(approxOK, force int) (min, max, mean, stdDev float64) {
	C.GDALGetRasterStatistics(
		rasterBand.cval,
		C.int(approxOK),
		C.int(force),
		(*C.double)(unsafe.Pointer(&min)),
		(*C.double)(unsafe.Pointer(&max)),
		(*C.double)(unsafe.Pointer(&mean)),
		(*C.double)(unsafe.Pointer(&stdDev)),
	)
	return min, max, mean, stdDev
}

// Compute image statistics
func (rasterBand RasterBand) ComputeStatistics(
	approxOK int,
	progress ProgressFunc,
	data interface{},
) (min, max, mean, stdDev float64) {
	arg := &goGDALProgressFuncProxyArgs{progress, data}

	C.GDALComputeRasterStatistics(
		rasterBand.cval,
		C.int(approxOK),
		(*C.double)(unsafe.Pointer(&min)),
		(*C.double)(unsafe.Pointer(&max)),
		(*C.double)(unsafe.Pointer(&mean)),
		(*C.double)(unsafe.Pointer(&stdDev)),
		C.goGDALProgressFuncProxyB(),
		unsafe.Pointer(arg),
	)
	return min, max, mean, stdDev
}

// Set statistics on raster band
func (rasterBand RasterBand) SetStatistics(min, max, mean, stdDev float64) error {
	err := C.GDALSetRasterStatistics(
		rasterBand.cval,
		C.double(min),
		C.double(max),
		C.double(mean),
		C.double(stdDev))
	return error(err)
}

// Return raster unit type
func (rasterBand RasterBand) GetUnitType() string {
	cString := C.GDALGetRasterUnitType(rasterBand.cval)
	return C.GoString(cString)
}

// Set unit type
func (rasterBand RasterBand) SetUnitType(unit string) error {
	cString := C.CString(unit)
	defer C.free(unsafe.Pointer(cString))

	err := C.GDALSetRasterUnitType(rasterBand.cval, cString)
	return error(err)
}

// Fetch the raster value offset
func (rasterBand RasterBand) GetOffset() (float64, bool) {	
	var success int
	val := C.GDALGetRasterOffset(rasterBand.cval, (*C.int)(unsafe.Pointer(&success)))
	return float64(val), (success != 0)
}

// Set scaling offset
func (rasterBand RasterBand) SetOffset(offset float64) error {
	err := C.GDALSetRasterOffset(rasterBand.cval, C.double(offset))
	return error(err)
}

// Fetch the raster value scale
func (rasterBand RasterBand) GetScale() (float64, bool) {
	var success int
	val := C.GDALGetRasterScale(rasterBand.cval, (*C.int)(unsafe.Pointer(&success)))
	return float64(val), (success != 0)
}

// Set scaling ratio
func (rasterBand RasterBand) SetScale(scale float64) error {
	err := C.GDALSetRasterScale(rasterBand.cval, C.double(scale))
	return error(err)
}

// Compute the min / max values for a band
func (rasterBand RasterBand) ComputeMinMax(approxOK int) (min, max float64) {
	var minmax [2]float64
	C.GDALComputeRasterMinMax(
		rasterBand.cval,
		C.int(approxOK),
		(*C.double)(unsafe.Pointer(&minmax[0])))
	return minmax[0], minmax[1]
}

// Flush raster data cache
func (rasterBand RasterBand) FlushCache() {
	C.GDALFlushRasterCache(rasterBand.cval)
}

// Unimplemented: GetRasterHistogram

// Unimplemented: GetDefaultHistogram

// Unimplemented: GetRandomRasterSample

// Unimplemented: GetRasterSampleOverview

// Fill this band with a constant value
func (rasterBand RasterBand) Fill(real, imaginary float64) error {
	err := C.GDALFillRaster(rasterBand.cval, C.double(real), C.double(imaginary))
	return error(err)
}

// Unimplemented: ComputeBandStats

// Unimplemented: OverviewMagnitudeCorrection

// Fetch default Raster Attribute Table
func (rasterBand RasterBand) GetDefaultRAT() RasterAttributeTable {
	rat := C.GDALGetDefaultRAT(rasterBand.cval)
	return RasterAttributeTable{rat}
}

// Set default Raster Attribute Table
func (rasterBand RasterBand) SetDefaultRAT(rat RasterAttributeTable) error {
	err := C.GDALSetDefaultRAT(rasterBand.cval, rat.cval)
	return error(err)
}

// Unimplemented: AddDerivedBandPixelFunc

// Return the mask band associated with the band
func (rasterBand RasterBand) GetMaskBand() RasterBand {
	mask := C.GDALGetMaskBand(rasterBand.cval)
	return RasterBand{mask}
}

// Return the status flags of the mask band associated with the band
func (rasterBand RasterBand) GetMaskFlags() int {
	flags := C.GDALGetMaskFlags(rasterBand.cval)
	return int(flags)
}

// Adds a mask band to the current band
func (rasterBand RasterBand) CreateMaskBand(flags int) error {
	err := C.GDALCreateMaskBand(rasterBand.cval, C.int(flags))
	return error(err)
}

// Copy all raster band raster data
func (sourceRaster RasterBand) RasterBandCopyWholeRaster(
	destRaster RasterBand,
	options []string,
	progress ProgressFunc,
	data interface{},
) error {
	arg := &goGDALProgressFuncProxyArgs{progress, data}

	length := len(options)
	cOptions := make([]*C.char, length + 1)
	for i := 0; i < length; i++ {
		cOptions[i] = C.CString(options[i])
		defer C.free(unsafe.Pointer(cOptions[i]))
	}
	cOptions[length] = (*C.char)(unsafe.Pointer(nil))

	err := C.GDALRasterBandCopyWholeRaster(
		sourceRaster.cval,
		destRaster.cval,
		(**C.char)(unsafe.Pointer(&cOptions[0])),
		C.goGDALProgressFuncProxyB(),
		unsafe.Pointer(arg),
	)
	return error(err)
}

// Generate downsampled overviews
// Unimplemented: RegenerateOverviews

/* ==================================================================== */
/*     GDALAsyncReader                                                  */
/* ==================================================================== */

// Unimplemented: GetNextUpdatedRegion
// Unimplemented: LockBuffer
// Unimplemented: UnlockBuffer

/* ==================================================================== */
/*      Color tables.                                                   */
/* ==================================================================== */

// Unimplemented: CreateColorTable
// Unimplemented: DestroyColorTable
// Unimplemented: CloneColorTable
// Unimplemented: GetPaletteInterpretation
// Unimplemented: GetColorEntryCount
// Unimplemented: GetColorEntry
// Unimplemented: GetColorEntryAsRGB
// Unimplemented: SetColorEntry
// Unimplemented: CreateColorRamp

/* ==================================================================== */
/*      Raster Attribute Table                                          */
/* ==================================================================== */

type RATFieldType int

const (
	GFT_Integer = RATFieldType(C.GFT_Integer)
	GFT_Real    = RATFieldType(C.GFT_Real)
	GFT_String  = RATFieldType(C.GFT_String)
)

type RATFieldUsage int

const (
	GFU_Generic    = RATFieldUsage(C.GFU_Generic)
	GFU_PixelCount = RATFieldUsage(C.GFU_PixelCount)
	GFU_Name       = RATFieldUsage(C.GFU_Name)
	GFU_Min        = RATFieldUsage(C.GFU_Min)
	GFU_Max        = RATFieldUsage(C.GFU_Max)
	GFU_MinMax     = RATFieldUsage(C.GFU_MinMax)
	GFU_Red        = RATFieldUsage(C.GFU_Red)
	GFU_Green      = RATFieldUsage(C.GFU_Green)
	GFU_Blue       = RATFieldUsage(C.GFU_Blue)
	GFU_Alpha      = RATFieldUsage(C.GFU_Alpha)
	GFU_RedMin     = RATFieldUsage(C.GFU_RedMin)
	GFU_GreenMin   = RATFieldUsage(C.GFU_GreenMin)
	GFU_BlueMin    = RATFieldUsage(C.GFU_BlueMin)
	GFU_AlphaMin   = RATFieldUsage(C.GFU_AlphaMin)
	GFU_RedMax     = RATFieldUsage(C.GFU_RedMax)
	GFU_GreenMax   = RATFieldUsage(C.GFU_GreenMax)
	GFU_BlueMax    = RATFieldUsage(C.GFU_BlueMax)
	GFU_AlphaMax   = RATFieldUsage(C.GFU_AlphaMax)
	GFU_MaxCount   = RATFieldUsage(C.GFU_MaxCount)
)

// Unimplemented: CreateRasterAttributeTable
// Unimplemented: DestroyRasterAttributeTable
// Unimplemented: GetColumnCount
// Unimplemented: GetNameOfCol
// Unimplemented: GetUsageOfCol
// Unimplemented: GetTypeOfCol
// Unimplemented: GetColOfUsage
// Unimplemented: GetRowCount
// Unimplemented: GetValueAsString
// Unimplemented: GetValueAsInt
// Unimplemented: GetValueAsDouble
// Unimplemented: SetValueAsString
// Unimplemented: SetValueAsInt
// Unimplemented: SetValueAsDouble
// Unimplemented: SetRowCount
// Unimplemented: CreateColumn
// Unimplemented: SetLinearBinning
// Unimplemented: GetLinearBinning
// Unimplemented: InitializeFromColorTable
// Unimplemented: TranslateToColorTable
// Unimplemented: DumpReadable
// Unimplemented: GetRowOfValue

/* ==================================================================== */
/*      GDAL Cache Management                                           */
/* ==================================================================== */

// Set maximum cache memory
func SetCacheMax(bytes int) {
	C.GDALSetCacheMax64(C.GIntBig(bytes))
}

// Get maximum cache memory
func GetCacheMax() int {
	bytes := C.GDALGetCacheMax64()
	return int(bytes)
}

// Get cache memory used
func GetCacheUsed() int {
	bytes := C.GDALGetCacheUsed64()
	return int(bytes)
}

// Try to flush one cached raster block
func FlushCacheBlock() bool {
	flushed := C.GDALFlushCacheBlock()
	return (flushed != 0)
}

/* -------------------------------------------------------------------- */
/*      Helper functions.                                               */
/* -------------------------------------------------------------------- */

// Unimplemented: GeneralCmdLineProcessor
// Unimplemented: SwapWords
// Unimplemented: CopyWords
// Unimplemented: CopyBits
// Unimplemented: LoadWorldFile
// Unimplemented: ReadWorldFile
// Unimplemented: WriteWorldFile
// Unimplemented: LoadTabFile
// Unimplemented: ReadTabFile
// Unimplemented: LoadOziMapFile
// Unimplemented: ReadOziMapFile
// Unimplemented: LoadRPBFile
// Unimplemented: LoadRPCFile
// Unimplemented: WriteRPBFile
// Unimplemented: LoadIMDFile
// Unimplemented: WriteIMDFile
// Unimplemented: DecToDMS
// Unimplemented: PackedDMSToDec
// Unimplemented: DecToPackedDMS

/* -------------------------------------------------------------------- */
/*      Spatial reference functions.                                    */
/* -------------------------------------------------------------------- */

type SpatialReference struct {
	cval C.OGRSpatialReferenceH
}

// Create a new SpatialReference
func CreateSpatialReference(wkt string) SpatialReference {
	cString := C.CString(wkt)
	defer C.free(unsafe.Pointer(cString))
	sr := C.OSRNewSpatialReference(cString)
	return SpatialReference{sr}
}

// Export coordinate system to WKT
func (sr SpatialReference) ToWkt() (string, error) {
	var p *C.char
	err := C.OSRExportToWkt(sr.cval, &p)
	wkt := C.GoString(p)
	return wkt, error(err)
}

// Initialize SRS based on EPSG code
func (sr SpatialReference) FromEPSG(code int) error {
	err := C.OSRImportFromEPSG(sr.cval, C.int(code))
	return error(err)
}

/* -------------------------------------------------------------------- */
/*      Coordinate transformation functions.                            */
/* -------------------------------------------------------------------- */

type CoordinateTransformation struct {
	cval C.OGRCoordinateTransformationH
}


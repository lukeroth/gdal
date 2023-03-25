package gdal

/*
#include "go_gdal.h"
#include "gdal_version.h"
*/
import "C"
import (
	"errors"
	"fmt"
	"reflect"
	"unsafe"
)

var _ = fmt.Println

func init() {
	C.GDALAllRegister()
}

/* -------------------------------------------------------------------- */
/*      Significant constants.                                          */
/* -------------------------------------------------------------------- */

const (
	VERSION_MAJOR = int(C.GDAL_VERSION_MAJOR)
	VERSION_MINOR = int(C.GDAL_VERSION_MINOR)
	VERSION_REV   = int(C.GDAL_VERSION_REV)
	VERSION_BUILD = int(C.GDAL_VERSION_BUILD)
	VERSION_NUM   = int(C.GDAL_VERSION_NUM)
	RELEASE_DATE  = int(C.GDAL_RELEASE_DATE)
	RELEASE_NAME  = string(C.GDAL_RELEASE_NAME)
)

var (
	ErrDebug   = errors.New("Debug Error")
	ErrWarning = errors.New("Warning Error")
	ErrFailure = errors.New("Failure Error")
	ErrFatal   = errors.New("Fatal Error")
	ErrIllegal = errors.New("Illegal Error")
)

// Error handling.  The following is bare-bones, and needs to be replaced with something more useful.
func (err C.CPLErr) Err() error {
	switch err {
	case 0:
		return nil
	case 1:
		return ErrDebug
	case 2:
		return ErrWarning
	case 3:
		return ErrFailure
	case 4:
		return ErrFailure
	}
	return ErrIllegal
}

func (err C.OGRErr) Err() error {
	switch err {
	case 0:
		return nil
	case 1:
		return ErrDebug
	case 2:
		return ErrWarning
	case 3:
		return ErrFailure
	case 4:
		return ErrFailure
	}
	return ErrIllegal
}

// Pixel data types
type DataType int

const (
	Unknown  = DataType(C.GDT_Unknown)
	Byte     = DataType(C.GDT_Byte)
	UInt16   = DataType(C.GDT_UInt16)
	Int16    = DataType(C.GDT_Int16)
	UInt32   = DataType(C.GDT_UInt32)
	Int32    = DataType(C.GDT_Int32)
	Float32  = DataType(C.GDT_Float32)
	Float64  = DataType(C.GDT_Float64)
	CInt16   = DataType(C.GDT_CInt16)
	CInt32   = DataType(C.GDT_CInt32)
	CFloat32 = DataType(C.GDT_CFloat32)
	CFloat64 = DataType(C.GDT_CFloat64)
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

// Safe array conversion
func IntSliceToCInt(data []int) []C.int {
	sliceSz := len(data)
	result := make([]C.int, sliceSz)
	for i := 0; i < sliceSz; i++ {
		result[i] = C.int(data[i])
	}
	return result
}

// Safe array conversion
func CIntSliceToInt(data []C.int) []int {
	sliceSz := len(data)
	result := make([]int, sliceSz)
	for i := 0; i < sliceSz; i++ {
		result[i] = int(data[i])
	}
	return result
}
func CUIntBigSliceToInt(data []C.GUIntBig) []int {
	sliceSz := len(data)
	result := make([]int, sliceSz)
	for i := 0; i < sliceSz; i++ {
		result[i] = int(data[i])
	}
	return result
}

// status of the asynchronous stream
type AsyncStatusType int

const (
	AR_Pending  = AsyncStatusType(C.GARIO_PENDING)
	AR_Update   = AsyncStatusType(C.GARIO_UPDATE)
	AR_Error    = AsyncStatusType(C.GARIO_ERROR)
	AR_Complete = AsyncStatusType(C.GARIO_COMPLETE)
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

type OpenFlag uint

const (
	OFReadOnly      = OpenFlag(C.GDAL_OF_READONLY)
	OFUpdate        = OpenFlag(C.GDAL_OF_UPDATE)
	OFShared        = OpenFlag(C.GDAL_OF_SHARED)
	OFVector        = OpenFlag(C.GDAL_OF_VECTOR)
	OFRaster        = OpenFlag(C.GDAL_OF_RASTER)
	OFVerbose_Error = OpenFlag(C.GDAL_OF_VERBOSE_ERROR)
)

// Types of color interpretation for raster bands.
type ColorInterp int

const (
	CI_Undefined      = ColorInterp(C.GCI_Undefined)
	CI_GrayIndex      = ColorInterp(C.GCI_GrayIndex)
	CI_PaletteIndex   = ColorInterp(C.GCI_PaletteIndex)
	CI_RedBand        = ColorInterp(C.GCI_RedBand)
	CI_GreenBand      = ColorInterp(C.GCI_GreenBand)
	CI_BlueBand       = ColorInterp(C.GCI_BlueBand)
	CI_AlphaBand      = ColorInterp(C.GCI_AlphaBand)
	CI_HueBand        = ColorInterp(C.GCI_HueBand)
	CI_SaturationBand = ColorInterp(C.GCI_SaturationBand)
	CI_LightnessBand  = ColorInterp(C.GCI_LightnessBand)
	CI_CyanBand       = ColorInterp(C.GCI_CyanBand)
	CI_MagentaBand    = ColorInterp(C.GCI_MagentaBand)
	CI_YellowBand     = ColorInterp(C.GCI_YellowBand)
	CI_BlackBand      = ColorInterp(C.GCI_BlackBand)
	CI_YCbCr_YBand    = ColorInterp(C.GCI_YCbCr_YBand)
	CI_YCbCr_CbBand   = ColorInterp(C.GCI_YCbCr_CbBand)
	CI_YCbCr_CrBand   = ColorInterp(C.GCI_YCbCr_CrBand)
	CI_Max            = ColorInterp(C.GCI_Max)
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
	PI_Gray = PaletteInterp(C.GPI_Gray)
	// Red, Green, Blue and Alpha in (in c1, c2, c3 and c4)
	PI_RGB = PaletteInterp(C.GPI_RGB)
	// Cyan, Magenta, Yellow and Black (in c1, c2, c3 and c4)
	PI_CMYK = PaletteInterp(C.GPI_CMYK)
	// Hue, Lightness and Saturation (in c1, c2, and c3)
	PI_HLS = PaletteInterp(C.GPI_HLS)
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

func (ce *ColorEntry) Set(c1, c2, c3, c4 uint) {
	ce.cval.c1 = C.short(c1)
	ce.cval.c2 = C.short(c2)
	ce.cval.c3 = C.short(c3)
	ce.cval.c4 = C.short(c4)
}

func (ce *ColorEntry) Get() (c1, c2, c3, c4 uint8) {

	return *(*uint8)(unsafe.Pointer(&ce.cval.c1)), *(*uint8)(unsafe.Pointer(&ce.cval.c2)), *(*uint8)(unsafe.Pointer(&ce.cval.c3)), *(*uint8)(unsafe.Pointer(&ce.cval.c4))
}

type VSILFILE struct {
	cval *C.VSILFILE
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
	data          interface{}
}

//export goGDALProgressFuncProxyA
func goGDALProgressFuncProxyA(complete C.double, message *C.char, data unsafe.Pointer) int {
	arg := (*goGDALProgressFuncProxyArgs)(data)
	return arg.progresssFunc(
		float64(complete), C.GoString(message), arg.data,
	)
}

// CPLSetConfigOption
func CPLSetConfigOption(key, val string) {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))
	cVal := C.CString(val)
	defer C.free(unsafe.Pointer(cVal))
	C.CPLSetConfigOption(cKey, cVal)
}

// CPLGetConfigOption
func CPLGetConfigOption(key, val string) string {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))

	cVal := C.CString(val)
	defer C.free(unsafe.Pointer(cVal))
	return C.GoString(C.CPLGetConfigOption(cKey, cVal))
}

/* ==================================================================== */
/*      Registration/driver related.                                    */
/* ==================================================================== */

const (
	DMD_LONGNAME           = string(C.GDAL_DMD_LONGNAME)
	DMD_HELPTOPIC          = string(C.GDAL_DMD_HELPTOPIC)
	DMD_MIMETYPE           = string(C.GDAL_DMD_MIMETYPE)
	DMD_EXTENSION          = string(C.GDAL_DMD_EXTENSION)
	DMD_CREATIONOPTIONLIST = string(C.GDAL_DMD_CREATIONOPTIONLIST)
	DMD_CREATIONDATATYPES  = string(C.GDAL_DMD_CREATIONDATATYPES)

	DCAP_CREATE     = string(C.GDAL_DCAP_CREATE)
	DCAP_CREATECOPY = string(C.GDAL_DCAP_CREATECOPY)
	DCAP_VIRTUALIO  = string(C.GDAL_DCAP_VIRTUALIO)
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
	opts := make([]*C.char, length+1)
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

// Create a copy of a dataset
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
	opts := make([]*C.char, length+1)
	for i := 0; i < length; i++ {
		opts[i] = C.CString(options[i])
		defer C.free(unsafe.Pointer(opts[i]))
	}
	opts[length] = (*C.char)(unsafe.Pointer(nil))

	var h C.GDALDatasetH

	if progress == nil {
		h = C.GDALCreateCopy(
			driver.cval, name,
			sourceDataset.cval,
			C.int(strict),
			(**C.char)(unsafe.Pointer(&opts[0])),
			nil,
			nil,
		)
	} else {
		arg := &goGDALProgressFuncProxyArgs{
			progress, data,
		}
		h = C.GDALCreateCopy(
			driver.cval, name,
			sourceDataset.cval,
			C.int(strict), (**C.char)(unsafe.Pointer(&opts[0])),
			C.goGDALProgressFuncProxyB(),
			unsafe.Pointer(arg),
		)
	}

	return Dataset{h}
}

// Return the driver needed to access the provided dataset name.
func IdentifyDriver(filename string, filenameList []string) Driver {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))

	length := len(filenameList)
	cFilenameList := make([]*C.char, length+1)
	for i := 0; i < length; i++ {
		cFilenameList[i] = C.CString(filenameList[i])
		defer C.free(unsafe.Pointer(cFilenameList[i]))
	}
	cFilenameList[length] = (*C.char)(unsafe.Pointer(nil))

	driver := C.GDALIdentifyDriver(cFilename, (**C.char)(unsafe.Pointer(&cFilenameList[0])))
	return Driver{driver}
}

// Open an existing dataset
func Open(filename string, access Access) (Dataset, error) {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))

	dataset := C.GDALOpen(cFilename, C.GDALAccess(access))
	if dataset == nil {
		return Dataset{nil}, fmt.Errorf("Error: dataset '%s' open error", filename)
	}
	return Dataset{dataset}, nil
}

// Open an existing dataset
func OpenEx(filename string, flags OpenFlag, allowedDrivers []string,
	openOptions []string, siblingFiles []string) (Dataset, error) {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))

	var driversA, ooptionsA, siblingsA **C.char
	if allowedDrivers != nil {
		length := len(allowedDrivers)
		drivers := make([]*C.char, length+1)
		for i := 0; i < length; i++ {
			drivers[i] = C.CString(allowedDrivers[i])
			defer C.free(unsafe.Pointer(drivers[i]))
		}
		drivers[length] = (*C.char)(unsafe.Pointer(nil))
		driversA = (**C.char)(unsafe.Pointer(&drivers[0]))
	}
	if openOptions != nil {
		length := len(openOptions)
		ooptions := make([]*C.char, length+1)
		for i := 0; i < length; i++ {
			ooptions[i] = C.CString(openOptions[i])
			defer C.free(unsafe.Pointer(ooptions[i]))
		}
		ooptions[length] = (*C.char)(unsafe.Pointer(nil))
		ooptionsA = (**C.char)(unsafe.Pointer(&ooptions[0]))
	}
	if siblingFiles != nil {
		length := len(siblingFiles)
		siblings := make([]*C.char, length+1)
		for i := 0; i < length; i++ {
			siblings[i] = C.CString(siblingFiles[i])
			defer C.free(unsafe.Pointer(siblings[i]))
		}
		siblings[length] = (*C.char)(unsafe.Pointer(nil))
		siblingsA = (**C.char)(unsafe.Pointer(&siblings[0]))
	}

	dataset := C.GDALOpenEx(cFilename, C.uint(flags), driversA, ooptionsA, siblingsA)
	if dataset == nil {
		return Dataset{nil}, fmt.Errorf("Error: dataset '%s' openEx error", filename)
	}
	return Dataset{dataset}, nil
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
	return C.GDALDeleteDataset(cDriver, cName).Err()
}

// Rename named dataset
func (driver Driver) RenameDataset(newName, oldName string) error {
	cDriver := driver.cval
	cNewName := C.CString(newName)
	defer C.free(unsafe.Pointer(cNewName))
	cOldName := C.CString(oldName)
	defer C.free(unsafe.Pointer(cOldName))
	return C.GDALRenameDataset(cDriver, cNewName, cOldName).Err()
}

// Copy all files associated with the named dataset
func (driver Driver) CopyDatasetFiles(newName, oldName string) error {
	cDriver := driver.cval
	cNewName := C.CString(newName)
	defer C.free(unsafe.Pointer(cNewName))
	cOldName := C.CString(oldName)
	defer C.free(unsafe.Pointer(cOldName))
	return C.GDALCopyDatasetFiles(cDriver, cNewName, cOldName).Err()
}

// Get the short name associated with this driver
func (driver Driver) ShortName() string {
	cDriver := driver.cval
	return C.GoString(C.GDALGetDriverShortName(cDriver))
}

// Get the long name associated with this driver
func (driver Driver) LongName() string {
	cDriver := driver.cval
	return C.GoString(C.GDALGetDriverLongName(cDriver))
}

/* ==================================================================== */
/*      GDAL_GCP                                                        */
/* ==================================================================== */

// Unimplemented: InitGCPs
// Unimplemented: DeinitGCPs
// Unimplemented: DuplicateGCPs
// Unimplemented: GCPsToGeoTransform
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
func (dataset *Dataset) Metadata(domain string) []string {
	cDomain := C.CString(domain)
	defer C.free(unsafe.Pointer(cDomain))

	p := C.GDALGetMetadata(
		C.GDALMajorObjectH(unsafe.Pointer(dataset.cval)),
		cDomain,
	)
	var strings []string
	q := uintptr(unsafe.Pointer(p))
	for {
		p = (**C.char)(unsafe.Pointer(q))
		if p == nil || *p == nil {
			break
		}
		strings = append(strings, C.GoString(*p))
		q += unsafe.Sizeof(q)
	}

	return strings
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

// TODO: Make correct class hirerarchy via interfaces

func (rasterBand *RasterBand) SetMetadataItem(name, value, domain string) error {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_value := C.CString(value)
	defer C.free(unsafe.Pointer(c_value))

	c_domain := C.CString(domain)
	defer C.free(unsafe.Pointer(c_domain))

	return C.GDALSetMetadataItem(
		C.GDALMajorObjectH(unsafe.Pointer(rasterBand.cval)),
		c_name, c_value, c_domain,
	).Err()
}

// TODO: Make korrekt class hirerarchy via interfaces

func (object *Dataset) SetMetadataItem(name, value, domain string) error {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_value := C.CString(value)
	defer C.free(unsafe.Pointer(c_value))

	c_domain := C.CString(domain)
	defer C.free(unsafe.Pointer(c_domain))

	return C.GDALSetMetadataItem(
		C.GDALMajorObjectH(unsafe.Pointer(object.cval)),
		c_name, c_value, c_domain,
	).Err()
}

// Fetch single metadata item.
func (object *Driver) MetadataItem(name, domain string) string {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_domain := C.CString(domain)
	defer C.free(unsafe.Pointer(c_domain))

	return C.GoString(
		C.GDALGetMetadataItem(
			C.GDALMajorObjectH(unsafe.Pointer(object.cval)),
			c_name, c_domain,
		),
	)
}
func (object *Dataset) MetadataItem(name, domain string) string {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_domain := C.CString(domain)
	defer C.free(unsafe.Pointer(c_domain))

	return C.GoString(
		C.GDALGetMetadataItem(
			C.GDALMajorObjectH(unsafe.Pointer(object.cval)),
			c_name, c_domain,
		),
	)
}

func (object *RasterBand) MetadataItem(name, domain string) string {
	c_name := C.CString(name)
	defer C.free(unsafe.Pointer(c_name))

	c_domain := C.CString(domain)
	defer C.free(unsafe.Pointer(c_domain))

	return C.GoString(
		C.GDALGetMetadataItem(
			C.GDALMajorObjectH(unsafe.Pointer(object.cval)),
			c_name, c_domain,
		),
	)
}

/* ==================================================================== */
/*      GDALDataset class ... normally this represents one file.        */
/* ==================================================================== */

// Get the driver to which this dataset relates
func (dataset Dataset) Driver() Driver {
	driver := Driver{C.GDALGetDatasetDriver(dataset.cval)}
	return driver
}

// Fetch files forming the dataset.
func (dataset Dataset) FileList() []string {
	p := C.GDALGetFileList(dataset.cval)
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
	cOptions := make([]*C.char, length+1)
	for i := 0; i < length; i++ {
		cOptions[i] = C.CString(options[i])
		defer C.free(unsafe.Pointer(cOptions[i]))
	}
	cOptions[length] = (*C.char)(unsafe.Pointer(nil))

	return C.GDALAddBand(
		dataset.cval,
		C.GDALDataType(dataType),
		(**C.char)(unsafe.Pointer(&cOptions[0])),
	).Err()
}

type ResampleAlg int

const (
	GRA_NearestNeighbour = ResampleAlg(0)
	GRA_Bilinear         = ResampleAlg(1)
	GRA_Cubic            = ResampleAlg(2)
	GRA_CubicSpline      = ResampleAlg(3)
	GRA_Lanczos          = ResampleAlg(4)
)

func (dataset Dataset) AutoCreateWarpedVRT(srcWKT, dstWKT string, resampleAlg ResampleAlg) (Dataset, error) {
	c_srcWKT := C.CString(srcWKT)
	defer C.free(unsafe.Pointer(c_srcWKT))
	c_dstWKT := C.CString(dstWKT)
	defer C.free(unsafe.Pointer(c_dstWKT))
	/*

	 */
	h := C.GDALAutoCreateWarpedVRT(dataset.cval, c_srcWKT, c_dstWKT, C.GDALResampleAlg(resampleAlg), 0.0, nil)
	d := Dataset{h}
	if h == nil {
		return d, fmt.Errorf("AutoCreateWarpedVRT failed")
	}
	return d, nil

}

// Unimplemented: GDALBeginAsyncReader
// Unimplemented: GDALEndAsyncReader

func determineBufferType(buffer interface{}) (dataType DataType, dataPtr unsafe.Pointer, err error) {
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
	case []float32:
		dataType = Float32
		dataPtr = unsafe.Pointer(&data[0])
	case []float64:
		dataType = Float64
		dataPtr = unsafe.Pointer(&data[0])
	default:
		err = fmt.Errorf("error: buffer is not a valid data type (must be a valid numeric slice)")
	}
	return
}

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
	dataType, dataPtr, err := determineBufferType(buffer)
	if err != nil {
		return err
	}

	return C.GDALDatasetRasterIO(
		dataset.cval,
		C.GDALRWFlag(rwFlag),
		C.int(xOff), C.int(yOff), C.int(xSize), C.int(ySize),
		dataPtr,
		C.int(bufXSize), C.int(bufYSize),
		C.GDALDataType(dataType),
		C.int(bandCount),
		(*C.int)(unsafe.Pointer(&IntSliceToCInt(bandMap)[0])),
		C.int(pixelSpace), C.int(lineSpace), C.int(bandSpace),
	).Err()
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
	cOptions := make([]*C.char, length+1)
	for i := 0; i < length; i++ {
		cOptions[i] = C.CString(options[i])
		defer C.free(unsafe.Pointer(cOptions[i]))
	}
	cOptions[length] = (*C.char)(unsafe.Pointer(nil))

	return C.GDALDatasetAdviseRead(
		dataset.cval,
		C.int(xOff), C.int(yOff), C.int(xSize), C.int(ySize),
		C.int(bufXSize), C.int(bufYSize),
		C.GDALDataType(dataType),
		C.int(bandCount),
		(*C.int)(unsafe.Pointer(&IntSliceToCInt(bandMap)[0])),
		(**C.char)(unsafe.Pointer(&cOptions[0])),
	).Err()
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

	return C.GDALSetProjection(dataset.cval, cProj).Err()
}

// Get the affine transformation coefficients
func (dataset Dataset) GeoTransform() [6]float64 {
	var transform [6]float64
	C.GDALGetGeoTransform(dataset.cval, (*C.double)(unsafe.Pointer(&transform[0])))
	return transform
}

// Set the affine transformation coefficients
func (dataset Dataset) SetGeoTransform(transform [6]float64) error {
	return C.GDALSetGeoTransform(
		dataset.cval,
		(*C.double)(unsafe.Pointer(&transform[0])),
	).Err()
}

// Return the inverted transform
func (dataset Dataset) InvGeoTransform() [6]float64 {
	return InvGeoTransform(dataset.GeoTransform())
}

// Invert the supplied transform
func InvGeoTransform(transform [6]float64) [6]float64 {
	var result [6]float64
	C.GDALInvGeoTransform((*C.double)(unsafe.Pointer(&transform[0])), (*C.double)(unsafe.Pointer(&result[0])))
	return result
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

	return C.GDALBuildOverviews(
		dataset.cval,
		cResampling,
		C.int(nOverviews),
		(*C.int)(unsafe.Pointer(&IntSliceToCInt(overviewList)[0])),
		C.int(nBands),
		(*C.int)(unsafe.Pointer(&IntSliceToCInt(bandList)[0])),
		C.goGDALProgressFuncProxyB(),
		unsafe.Pointer(arg),
	).Err()
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
	return C.GDALCreateDatasetMaskBand(dataset.cval, C.int(flags)).Err()
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
	cOptions := make([]*C.char, length+1)
	for i := 0; i < length; i++ {
		cOptions[i] = C.CString(options[i])
		defer C.free(unsafe.Pointer(cOptions[i]))
	}
	cOptions[length] = (*C.char)(unsafe.Pointer(nil))

	return C.GDALDatasetCopyWholeRaster(
		sourceDataset.cval,
		destDataset.cval,
		(**C.char)(unsafe.Pointer(&cOptions[0])),
		C.goGDALProgressFuncProxyB(),
		unsafe.Pointer(arg),
	).Err()
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
	cOptions := make([]*C.char, length+1)
	for i := 0; i < length; i++ {
		cOptions[i] = C.CString(options[i])
		defer C.free(unsafe.Pointer(cOptions[i]))
	}
	cOptions[length] = (*C.char)(unsafe.Pointer(nil))

	return C.GDALRasterAdviseRead(
		rasterBand.cval,
		C.int(xOff), C.int(yOff), C.int(xSize), C.int(ySize), C.int(bufXSize), C.int(bufYSize),
		C.GDALDataType(dataType),
		(**C.char)(unsafe.Pointer(&cOptions[0])),
	).Err()
}

// Read / Write a region of image data for this band
func (rasterBand RasterBand) IO(
	rwFlag RWFlag,
	xOff, yOff, xSize, ySize int,
	buffer interface{},
	bufXSize, bufYSize int,
	pixelSpace, lineSpace int,
) error {
	dataType, dataPtr, err := determineBufferType(buffer)
	if err != nil {
		return err
	}

	return C.GDALRasterIO(
		rasterBand.cval,
		C.GDALRWFlag(rwFlag),
		C.int(xOff), C.int(yOff), C.int(xSize), C.int(ySize),
		dataPtr,
		C.int(bufXSize), C.int(bufYSize),
		C.GDALDataType(dataType),
		C.int(pixelSpace), C.int(lineSpace),
	).Err()
}

// Read a block of image data efficiently
func (rasterBand RasterBand) ReadBlock(xOff, yOff int, dataPtr unsafe.Pointer) error {
	return C.GDALReadBlock(rasterBand.cval, C.int(xOff), C.int(yOff), dataPtr).Err()
}

// Write a block of image data efficiently
func (rasterBand RasterBand) WriteBlock(xOff, yOff int, dataPtr unsafe.Pointer) error {
	return C.GDALWriteBlock(rasterBand.cval, C.int(xOff), C.int(yOff), dataPtr).Err()
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
	return C.GDALSetRasterColorInterpretation(rasterBand.cval, C.GDALColorInterp(colorInterp)).Err()
}

// Fetch the color table associated with this raster band
func (rasterBand RasterBand) ColorTable() ColorTable {
	colorTable := C.GDALGetRasterColorTable(rasterBand.cval)
	return ColorTable{colorTable}
}

// Set the raster color table for this raster band
func (rasterBand RasterBand) SetColorTable(colorTable ColorTable) error {
	return C.GDALSetRasterColorTable(rasterBand.cval, colorTable.cval).Err()
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
	return float64(noDataVal), success != 0
}

// Set the no data value for this band
func (rasterBand RasterBand) SetNoDataValue(val float64) error {
	return C.GDALSetRasterNoDataValue(rasterBand.cval, C.double(val)).Err()
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
	cStrings := make([]*C.char, length+1)
	for i := 0; i < length; i++ {
		cStrings[i] = C.CString(names[i])
		defer C.free(unsafe.Pointer(cStrings[i]))
	}
	cStrings[length] = (*C.char)(unsafe.Pointer(nil))

	return C.GDALSetRasterCategoryNames(rasterBand.cval, (**C.char)(unsafe.Pointer(&cStrings[0]))).Err()
}

// Fetch the minimum value for this band
func (rasterBand RasterBand) GetMinimum() (val float64, valid bool) {
	var success int
	min := C.GDALGetRasterMinimum(rasterBand.cval, (*C.int)(unsafe.Pointer(&success)))
	return float64(min), success != 0
}

// Fetch the maximum value for this band
func (rasterBand RasterBand) GetMaximum() (val float64, valid bool) {
	var success int
	max := C.GDALGetRasterMaximum(rasterBand.cval, (*C.int)(unsafe.Pointer(&success)))
	return float64(max), success != 0
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
	return C.GDALSetRasterStatistics(
		rasterBand.cval,
		C.double(min),
		C.double(max),
		C.double(mean),
		C.double(stdDev),
	).Err()
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

	return C.GDALSetRasterUnitType(rasterBand.cval, cString).Err()
}

// Fetch the raster value offset
func (rasterBand RasterBand) GetOffset() (float64, bool) {
	var success int
	val := C.GDALGetRasterOffset(rasterBand.cval, (*C.int)(unsafe.Pointer(&success)))
	return float64(val), success != 0
}

// Set scaling offset
func (rasterBand RasterBand) SetOffset(offset float64) error {
	return C.GDALSetRasterOffset(rasterBand.cval, C.double(offset)).Err()
}

// Fetch the raster value scale
func (rasterBand RasterBand) GetScale() (float64, bool) {
	var success int
	val := C.GDALGetRasterScale(rasterBand.cval, (*C.int)(unsafe.Pointer(&success)))
	return float64(val), success != 0
}

// Set scaling ratio
func (rasterBand RasterBand) SetScale(scale float64) error {
	return C.GDALSetRasterScale(rasterBand.cval, C.double(scale)).Err()
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

// Compute raster histogram
func (rasterBand RasterBand) Histogram(
	min, max float64,
	buckets int,
	includeOutOfRange, approxOK int,
	progress ProgressFunc,
	data interface{},
) ([]int, error) {
	arg := &goGDALProgressFuncProxyArgs{
		progress, data,
	}

	histogram := make([]C.GUIntBig, buckets)

	if err := C.GDALGetRasterHistogramEx(
		rasterBand.cval,
		C.double(min),
		C.double(max),
		C.int(buckets),
		(*C.GUIntBig)(unsafe.Pointer(&histogram[0])),
		C.int(includeOutOfRange),
		C.int(approxOK),
		C.goGDALProgressFuncProxyB(),
		unsafe.Pointer(arg),
	).Err(); err != nil {
		return nil, err
	} else {
		return CUIntBigSliceToInt(histogram), nil
	}
}

// Fetch default raster histogram
func (rasterBand RasterBand) DefaultHistogram(
	force int,
	progress ProgressFunc,
	data interface{},
) (min, max float64, buckets int, histogram []int, err error) {
	arg := &goGDALProgressFuncProxyArgs{
		progress, data,
	}

	var cHistogram *C.GUIntBig

	err = C.GDALGetDefaultHistogramEx(
		rasterBand.cval,
		(*C.double)(&min),
		(*C.double)(&max),
		(*C.int)(unsafe.Pointer(&buckets)),
		&cHistogram,
		C.int(force),
		C.goGDALProgressFuncProxyB(),
		unsafe.Pointer(arg),
	).Err()

	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&histogram))
	sliceHeader.Cap = buckets
	sliceHeader.Len = buckets
	sliceHeader.Data = uintptr(unsafe.Pointer(cHistogram))

	return min, max, buckets, histogram, err
}

// Set default raster histogram
// Unimplemented: SetDefaultHistogram

// Unimplemented: GetRandomRasterSample

// Fetch best sampling overviews
// Unimplemented: GetRasterSampleOverview

// Fill this band with a constant value
func (rasterBand RasterBand) Fill(real, imaginary float64) error {
	return C.GDALFillRaster(rasterBand.cval, C.double(real), C.double(imaginary)).Err()
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
	return C.GDALSetDefaultRAT(rasterBand.cval, rat.cval).Err()
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
	return C.GDALCreateMaskBand(rasterBand.cval, C.int(flags)).Err()
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
	cOptions := make([]*C.char, length+1)
	for i := 0; i < length; i++ {
		cOptions[i] = C.CString(options[i])
		defer C.free(unsafe.Pointer(cOptions[i]))
	}
	cOptions[length] = (*C.char)(unsafe.Pointer(nil))

	return C.GDALRasterBandCopyWholeRaster(
		sourceRaster.cval,
		destRaster.cval,
		(**C.char)(unsafe.Pointer(&cOptions[0])),
		C.goGDALProgressFuncProxyB(),
		unsafe.Pointer(arg),
	).Err()
}

// Generate downsampled overviews
func (sourceRaster RasterBand) RegenerateOverviews(
	overviewCount int,
	destRasterBands *RasterBand,
	resampling string,
	progress ProgressFunc,
	data interface{},
) error {
	arg := &goGDALProgressFuncProxyArgs{progress, data}
	cVal := C.CString(resampling)
	defer C.free(unsafe.Pointer(cVal))
	return C.GDALRegenerateOverviews(
		sourceRaster.cval,
		C.int(overviewCount),
		&destRasterBands.cval,
		cVal,
		C.goGDALProgressFuncProxyB(),
		unsafe.Pointer(arg),
	).Err()
}

/* ==================================================================== */
/*     GDALAsyncReader                                                  */
/* ==================================================================== */

// Unimplemented: GetNextUpdatedRegion
// Unimplemented: LockBuffer
// Unimplemented: UnlockBuffer

/* ==================================================================== */
/*      Color tables.                                                   */
/* ==================================================================== */

// Construct a new color table
func CreateColorTable(interp PaletteInterp) ColorTable {
	ct := C.GDALCreateColorTable(C.GDALPaletteInterp(interp))
	return ColorTable{ct}
}

// Destroy the color table
func (ct ColorTable) Destroy() {
	C.GDALDestroyColorTable(ct.cval)
}

// Make a copy of the color table
func (ct ColorTable) Clone() ColorTable {
	newCT := C.GDALCloneColorTable(ct.cval)
	return ColorTable{newCT}
}

// Fetch palette interpretation
func (ct ColorTable) PaletteInterpretation() PaletteInterp {
	pi := C.GDALGetPaletteInterpretation(ct.cval)
	return PaletteInterp(pi)
}

// Get number of color entries in table
func (ct ColorTable) EntryCount() int {
	count := C.GDALGetColorEntryCount(ct.cval)
	return int(count)
}

// Fetch a color entry from table
func (ct ColorTable) Entry(index int) ColorEntry {
	entry := C.GDALGetColorEntry(ct.cval, C.int(index))
	return ColorEntry{*entry}
}

// Unimplemented: EntryAsRGB

// Set entry in color table
func (ct ColorTable) SetEntry(index int, entry ColorEntry) {
	C.GDALSetColorEntry(ct.cval, C.int(index), &entry.cval)
}

// Create color ramp
func (ct ColorTable) CreateColorRamp(start, end int, startColor, endColor ColorEntry) {
	C.GDALCreateColorRamp(ct.cval, C.int(start), &startColor.cval, C.int(end), &endColor.cval)
}

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

// Construct empty raster attribute table
func CreateRasterAttributeTable() RasterAttributeTable {
	rat := C.GDALCreateRasterAttributeTable()
	return RasterAttributeTable{rat}
}

// Destroy a RAT
func (rat RasterAttributeTable) Destroy() {
	C.GDALDestroyRasterAttributeTable(rat.cval)
}

// Fetch table column count
func (rat RasterAttributeTable) ColumnCount() int {
	count := C.GDALRATGetColumnCount(rat.cval)
	return int(count)
}

// Fetch the name of indicated column
func (rat RasterAttributeTable) NameOfCol(index int) string {
	name := C.GDALRATGetNameOfCol(rat.cval, C.int(index))
	return C.GoString(name)
}

// Fetch the usage of indicated column
func (rat RasterAttributeTable) UsageOfCol(index int) RATFieldUsage {
	rfu := C.GDALRATGetUsageOfCol(rat.cval, C.int(index))
	return RATFieldUsage(rfu)
}

// Fetch the type of indicated column
func (rat RasterAttributeTable) TypeOfCol(index int) RATFieldType {
	rft := C.GDALRATGetTypeOfCol(rat.cval, C.int(index))
	return RATFieldType(rft)
}

// Fetch column index for indicated usage
func (rat RasterAttributeTable) ColOfUsage(rfu RATFieldUsage) int {
	index := C.GDALRATGetColOfUsage(rat.cval, C.GDALRATFieldUsage(rfu))
	return int(index)
}

// Fetch row count
func (rat RasterAttributeTable) RowCount() int {
	count := C.GDALRATGetRowCount(rat.cval)
	return int(count)
}

// Fetch field value as string
func (rat RasterAttributeTable) ValueAsString(row, field int) string {
	cString := C.GDALRATGetValueAsString(rat.cval, C.int(row), C.int(field))
	return C.GoString(cString)
}

// Fetch field value as integer
func (rat RasterAttributeTable) ValueAsInt(row, field int) int {
	val := C.GDALRATGetValueAsInt(rat.cval, C.int(row), C.int(field))
	return int(val)
}

// Fetch field value as float64
func (rat RasterAttributeTable) ValueAsFloat64(row, field int) float64 {
	val := C.GDALRATGetValueAsDouble(rat.cval, C.int(row), C.int(field))
	return float64(val)
}

// Set field value from string
func (rat RasterAttributeTable) SetValueAsString(row, field int, val string) {
	cVal := C.CString(val)
	defer C.free(unsafe.Pointer(cVal))
	C.GDALRATSetValueAsString(rat.cval, C.int(row), C.int(field), cVal)
}

// Set field value from integer
func (rat RasterAttributeTable) SetValueAsInt(row, field, val int) {
	C.GDALRATSetValueAsInt(rat.cval, C.int(row), C.int(field), C.int(val))
}

// Set field value from float64
func (rat RasterAttributeTable) SetValueAsFloat64(row, field int, val float64) {
	C.GDALRATSetValueAsDouble(rat.cval, C.int(row), C.int(field), C.double(val))
}

// Set row count
func (rat RasterAttributeTable) SetRowCount(count int) {
	C.GDALRATSetRowCount(rat.cval, C.int(count))
}

// Create new column
func (rat RasterAttributeTable) CreateColumn(name string, rft RATFieldType, rfu RATFieldUsage) error {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	return C.GDALRATCreateColumn(rat.cval, cName, C.GDALRATFieldType(rft), C.GDALRATFieldUsage(rfu)).Err()
}

// Set linear binning information
func (rat RasterAttributeTable) SetLinearBinning(row0min, binsize float64) error {
	return C.GDALRATSetLinearBinning(rat.cval, C.double(row0min), C.double(binsize)).Err()
}

// Fetch linear binning information
func (rat RasterAttributeTable) LinearBinning() (row0min, binsize float64, exists bool) {
	success := C.GDALRATGetLinearBinning(rat.cval, (*C.double)(&row0min), (*C.double)(&binsize))
	return row0min, binsize, success != 0
}

// Initialize RAT from color table
func (rat RasterAttributeTable) FromColorTable(ct ColorTable) error {
	return C.GDALRATInitializeFromColorTable(rat.cval, ct.cval).Err()
}

// Translate RAT to a color table
func (rat RasterAttributeTable) ToColorTable(count int) ColorTable {
	ct := C.GDALRATTranslateToColorTable(rat.cval, C.int(count))
	return ColorTable{ct}
}

// Dump RAT in readable form to a file
// Unimplemented: DumpReadable

// Get row for pixel value
func (rat RasterAttributeTable) RowOfValue(val float64) (int, bool) {
	row := C.GDALRATGetRowOfValue(rat.cval, C.double(val))
	return int(row), row != -1
}

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
	return flushed != 0
}

/* ==================================================================== */
/*      GDAL VSI Virtual File System                                    */
/* ==================================================================== */

// List VSI files
func VSIReadDirRecursive(filename string) []string {
	name := C.CString(filename)
	defer C.free(unsafe.Pointer(name))

	p := C.VSIReadDirRecursive(name)
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

// Open file.
func VSIFOpenL(fileName string, fileAccess string) (VSILFILE, error) {
	cFileName := C.CString(fileName)
	defer C.free(unsafe.Pointer(cFileName))
	cFileAccess := C.CString(fileAccess)
	defer C.free(unsafe.Pointer(cFileAccess))
	file := C.VSIFOpenL(cFileName, cFileAccess)

	if file == nil {
		return VSILFILE{nil}, fmt.Errorf("Error: VSILFILE '%s' open error", fileName)
	}
	return VSILFILE{file}, nil
}

// Close file.
func VSIFCloseL(file VSILFILE) {
	C.VSIFCloseL(file.cval)
	return
}

// Read bytes from file.
func VSIFReadL(nSize, nCount int, file VSILFILE) []byte {
	data := make([]byte, nSize*nCount)
	p := unsafe.Pointer(&data[0])
	C.VSIFReadL(p, C.size_t(nSize), C.size_t(nCount), file.cval)

	return data
}

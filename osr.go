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
	"reflect"
)

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

// Initialize SRS based on WKT string
func (sr SpatialReference) FromWKT(wkt string) error {
	cString := C.CString(wkt)
	defer C.free(unsafe.Pointer(cString))
	err := C.OSRImportFromWkt(sr.cval, &cString)
	return error(err)
}

// Export coordinate system to WKT
func (sr SpatialReference) ToWKT() (string, error) {
	var p *C.char
	err := C.OSRExportToWkt(sr.cval, &p)
	wkt := C.GoString(p)
	return wkt, error(err)
}

// Export coordinate system to a nicely formatted WKT string
func (sr SpatialReference) ToPrettyWKT(simplify bool) (string, error) {
	var p *C.char
	var cBool int
	if simplify {
		cBool = 1
	} else {
		cBool = 0
	}
	err := C.OSRExportToPrettyWkt(sr.cval, &p, C.int(cBool))
	wkt := C.GoString(p)
	return wkt, error(err)
}

// Initialize SRS based on EPSG code
func (sr SpatialReference) FromEPSG(code int) error {
	err := C.OSRImportFromEPSG(sr.cval, C.int(code))
	return error(err)
}

// Initialize SRS based on EPSG code, using EPSG lat/long ordering
func (sr SpatialReference) FromEPSGA(code int) error {
	err := C.OSRImportFromEPSGA(sr.cval, C.int(code))
	return error(err)
}

// Destroy the spatial reference
func (sr SpatialReference) Destroy() {
	C.OSRDestroySpatialReference(sr.cval)
}

// Make a duplicate of this spatial reference
func (sr SpatialReference) Clone() SpatialReference {
	newSR := C.OSRClone(sr.cval)
	return SpatialReference{newSR}
}

// Make a duplicate of the GEOGCS node of this spatial reference
func (sr SpatialReference) CloneGeogCS() SpatialReference {
	newSR := C.OSRCloneGeogCS(sr.cval)
	return SpatialReference{newSR}
}

// Increments the reference count by one, returning reference count
func (sr SpatialReference) Reference() int {
	count := C.OSRReference(sr.cval)
	return int(count)
}

// Decrements the reference count by one, returning reference count
func (sr SpatialReference) Dereference() int {
	count := C.OSRDereference(sr.cval)
	return int(count)
}

// Decrements the reference count by one and destroy if zero
func (sr SpatialReference) Release() {
	C.OSRRelease(sr.cval)
}

// Validate spatial reference tokens
func (sr SpatialReference) Validate() error {
	err := C.OSRValidate(sr.cval)
	return error(err)
}

// Correct parameter ordering to match CT specification
func (sr SpatialReference) FixupOrdering() error {
	err := C.OSRFixupOrdering(sr.cval)
	return error(err)
}

// Fix up spatial reference as needed
func (sr SpatialReference) Fixup() error {
	err := C.OSRFixup(sr.cval)
	return error(err)
}

// Strip OGC CT parameters
func (sr SpatialReference) StripCTParams() error {
	err := C.OSRStripCTParms(sr.cval)
	return error(err)
}

// Import PROJ.4 coordinate string
func (sr SpatialReference) FromProj4(input string) error {
	cString := C.CString(input)
	defer C.free(unsafe.Pointer(cString))
	err := C.OSRImportFromProj4(sr.cval, cString)
	return error(err)
}

// Export coordinate system in PROJ.4 format
func (sr SpatialReference) ToProj4() (string, error) {
	var p *C.char
	err := C.OSRExportToProj4(sr.cval, &p)
	proj4 := C.GoString(p)
	return proj4, error(err)
}

// Import coordinate system from ESRI .prj formats
func (sr SpatialReference) FromESRI(input string) error {
	cString := C.CString(input)
	defer C.free(unsafe.Pointer(cString))
	err := C.OSRImportFromProj4(sr.cval, cString)
	return error(err)
}

// Import coordinate system from PCI projection definition
func (sr SpatialReference) FromPCI(proj, units string, params []float64) error {
	cProj := C.CString(proj)
	defer C.free(unsafe.Pointer(cProj))
	cUnits := C.CString(units)
	defer C.free(unsafe.Pointer(cUnits))

	err := C.OSRImportFromPCI(
		sr.cval,
		cProj,
		cUnits,
		(*C.double)(unsafe.Pointer(&params[0])))
	return error(err)
}

// Import coordinate system from USGS projection definition
func (sr SpatialReference) FromUSGS(projsys, zone int, params []float64, datum int) error {
	err := C.OSRImportFromUSGS(
		sr.cval,
		C.long(projsys),
		C.long(zone),
		(*C.double)(unsafe.Pointer(&params[0])),
		C.long(datum))
	return error(err)
}

// Import coordinate system from XML format (GML only currently)
func (sr SpatialReference) FromXML(xml string) error {
	cXml := C.CString(xml)
	defer C.free(unsafe.Pointer(cXml))
	err := C.OSRImportFromXML(sr.cval, cXml)
	return error(err)
}

// Import coordinate system from ERMapper projection definitions
func (sr SpatialReference) FromERM(proj, datum, units string) error {
	cProj := C.CString(proj)
	defer C.free(unsafe.Pointer(cProj))
	cDatum := C.CString(datum)
	defer C.free(unsafe.Pointer(cDatum))
	cUnits := C.CString(units)
	defer C.free(unsafe.Pointer(cUnits))

	err := C.OSRImportFromERM(sr.cval, cProj, cDatum, cUnits)
	return error(err)
}

// Import coordinate system from a URL
func (sr SpatialReference) FromURL(url string) error {
	cURL := C.CString(url)
	defer C.free(unsafe.Pointer(cURL))
	err := C.OSRImportFromXML(sr.cval, cURL)
	return error(err)
}

// Export coordinate system in PCI format
func (sr SpatialReference) ToPCI() (proj, units string, params []float64, errVal error) {
	var p, u *C.char
	err := C.OSRExportToPCI(sr.cval, &p, &u, (**C.double)(unsafe.Pointer(&params[0])))
	header := (*reflect.SliceHeader)((unsafe.Pointer(&params)))
	header.Cap = 17
	header.Len = 17
	defer C.free(unsafe.Pointer(p))
	defer C.free(unsafe.Pointer(u))
	return C.GoString(p), C.GoString(u), params, error(err)
}

// Export coordinate system to USGS GCTP projection definition
func (sr SpatialReference) ToUSGS() (proj, zone int, params []float64, datum int, errVal error) {
	err := C.OSRExportToUSGS(
		sr.cval, 
		(*C.long)(unsafe.Pointer(&proj)),
		(*C.long)(unsafe.Pointer(&zone)),
		(**C.double)(unsafe.Pointer(&params[0])),
		(*C.long)(unsafe.Pointer(&datum)))

	header := (*reflect.SliceHeader)((unsafe.Pointer(&params)))
	header.Cap = 15
	header.Len = 15
		
	return proj, zone, params, datum, error(err)
}

// Export coordinate system in XML format
func (sr SpatialReference) ToXML() (xml string, errVal error) {
	var x *C.char
	err := C.OSRExportToXML(sr.cval, &x, nil)
	defer C.free(unsafe.Pointer(x))
	return C.GoString(x), error(err)
}

// Export coordinate system in Mapinfo style CoordSys format
func (sr SpatialReference) ToMICoordSys() (output string, errVal error) {
	var x *C.char
	err := C.OSRExportToMICoordSys(sr.cval, &x)
	defer C.free(unsafe.Pointer(x))
	return C.GoString(x), error(err)
}

// Export coordinate system in ERMapper format
// Unimplemented: ToERM

// Convert in place to ESRI WKT format
func (sr SpatialReference) MorphToESRI() error {
	err := C.OSRMorphToESRI(sr.cval)
	return error(err)
}

// Convert in place from ESRI WKT format
func (sr SpatialReference) MorphFromESRI() error {
	err := C.OSRMorphFromESRI(sr.cval)
	return error(err)
}

// Fetch indicated attribute of named node
func (sr SpatialReference) AttrValue(key string, child int) (value string, ok bool) {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))
	val := C.OSRGetAttrValue(sr.cval, cKey, C.int(child))
	return C.GoString(val), val != nil
}

// Set attribute value in spatial reference
func (sr SpatialReference) SetAttrValue(path, value string) error {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))
	err := C.OSRSetAttrValue(sr.cval, cPath, cValue)
	return error(err)
}

// Set the angular units for the geographic coordinate system
func (sr SpatialReference) SetAngularUnits(units string, radians float64) error {
	cUnits := C.CString(units)
	defer C.free(unsafe.Pointer(cUnits))
	err := C.OSRSetAngularUnits(sr.cval, cUnits, C.double(radians))
	return error(err)
}

// Fetch the angular units for the geographic coordinate system
func (sr SpatialReference) AngularUnits() (string, float64) {
	var x *C.char
	factor := C.OSRGetAngularUnits(sr.cval, &x)
	defer C.free(unsafe.Pointer(x))
	return C.GoString(x), float64(factor)
}

// Set the linear units for the projection
func (sr SpatialReference) SetLinearUnits(name string, toMeters float64) error {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	err := C.OSRSetLinearUnits(sr.cval, cName, C.double(toMeters))
	return error(err)
}

// Set the linear units for the target node
func (sr SpatialReference) SetTargetLinearUnits(target, units string, toMeters float64) error {
	cTarget := C.CString(target)
	defer C.free(unsafe.Pointer(cTarget))
	cUnits := C.CString(units)
	defer C.free(unsafe.Pointer(cUnits))
	err := C.OSRSetTargetLinearUnits(sr.cval, cTarget, cUnits, C.double(toMeters))
	return error(err)
}

// Set the linear units for the target node and update all existing linear parameters
func (sr SpatialReference) SetLinearUnitsAndUpdateParameters(name string, toMeters float64) error {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	err := C.OSRSetLinearUnitsAndUpdateParameters(sr.cval, cName, C.double(toMeters))
	return error(err)
}

// Fetch linear projection units
func (sr SpatialReference) LinearUnits() (string, float64) {
	var x *C.char
	factor := C.OSRGetLinearUnits(sr.cval, &x)
	defer C.free(unsafe.Pointer(x))
	return C.GoString(x), float64(factor)
}

// Fetch linear units for target
func (sr SpatialReference) TargetLinearUnits(target string) (string, float64) {
	cTarget := C.CString(target)
	defer C.free(unsafe.Pointer(cTarget))
	var x *C.char
	factor := C.OSRGetTargetLinearUnits(sr.cval, cTarget, &x)
	defer C.free(unsafe.Pointer(x))
	return C.GoString(x), float64(factor)
}

// Fetch prime meridian information
func (sr SpatialReference) PrimeMeridian() (string, float64) {
	var x *C.char
	offset := C.OSRGetPrimeMeridian(sr.cval, &x)
	defer C.free(unsafe.Pointer(x))
	return C.GoString(x), float64(offset)
}

// Return true if geographic coordinate system
func (sr SpatialReference) IsGeographic() bool {
	val := C.OSRIsGeographic(sr.cval)
	return val != 0
}

// Return true if local coordinate system
func (sr SpatialReference) IsLocal() bool {
	val := C.OSRIsLocal(sr.cval)
	return val != 0
}

// Return true if projected coordinate system
func (sr SpatialReference) IsProjected() bool {
	val := C.OSRIsProjected(sr.cval)
	return val != 0
}

// Return true if compound coordinate system
func (sr SpatialReference) IsCompound() bool {
	val := C.OSRIsCompound(sr.cval)
	return val != 0
}

// Return true if geocentric coordinate system
func (sr SpatialReference) IsGeocentric() bool {
	val := C.OSRIsGeocentric(sr.cval)
	return val != 0
}

// Return true if vertical coordinate system
func (sr SpatialReference) IsVertical() bool {
	val := C.OSRIsVertical(sr.cval)
	return val != 0
}

// Return true if the geographic coordinate systems match
func (sr SpatialReference) IsSameGeographicCS(other SpatialReference) bool {
	val := C.OSRIsSameGeogCS(sr.cval, other.cval)
	return val != 0
}

// Return true if the vertical coordinate systems match
func (sr SpatialReference) IsSameVerticalCS(other SpatialReference) bool {
	val := C.OSRIsSameVertCS(sr.cval, other.cval)
	return val != 0
}

// Return true if the coordinate systems describe the same system
func (sr SpatialReference) IsSame(other SpatialReference) bool {
	val := C.OSRIsSame(sr.cval, other.cval)
	return val != 0
}

// Set the user visible local CS name
func (sr SpatialReference) SetLocalCS(name string) error {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	err := C.OSRSetLocalCS(sr.cval, cName)
	return error(err)
}

// Set the user visible projected CS name
func (sr SpatialReference) SetProjectedCS(name string) error {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	err := C.OSRSetProjCS(sr.cval, cName)
	return error(err)
}

// Set the user visible geographic CS name
func (sr SpatialReference) SetGeocentricCS(name string) error {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	err := C.OSRSetGeocCS(sr.cval, cName)
	return error(err)
}

// Set geographic CS based on well known name
func (sr SpatialReference) SetWellKnownGeographicCS(name string) error {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	err := C.OSRSetWellKnownGeogCS(sr.cval, cName)
	return error(err)
}

// Set spatial reference from various text formats
func (sr SpatialReference) SetFromUserInput(name string) error {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	err := C.OSRSetFromUserInput(sr.cval, cName)
	return error(err)
}

// Copy geographic CS from another spatial reference
func (sr SpatialReference) CopyGeographicCSFrom(other SpatialReference) error {
	err := C.OSRCopyGeogCSFrom(sr.cval, other.cval)
	return error(err)
}

// Set the Bursa-Wolf conversion to WGS84
func (sr SpatialReference) SetTOWGS84(dx, dy, dz, ex, ey, ez, ppm float64) error {
	err := C.OSRSetTOWGS84(
		sr.cval,
		C.double(dx),
		C.double(dy),
		C.double(dz),
		C.double(ex),
		C.double(ey),
		C.double(ez),
		C.double(ppm))
	return error(err)
}

// Fetch the TOWGS84 parameters if available
func (sr SpatialReference) TOWGS84() (coeff [7]float64, errVal error) {
	err := C.OSRGetTOWGS84(sr.cval, (*C.double)(unsafe.Pointer(&coeff[0])), 7)
	return coeff, error(err)
}

// Setup a compound coordinate system
func (sr SpatialReference) SetCompoundCS(
	name string,
	horizontal, vertical SpatialReference,
) error {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	err := C.OSRSetCompoundCS(sr.cval, cName, horizontal.cval, vertical.cval)
	return error(err)
}

// Unimplemented: SetGeogCS

// Unimplemented: SetVertCS

// Unimplemented: SemiMajor

// Unimplemented: SemiMinor

// Unimplemented: InvFlattening

// Unimplemented: SetAuthority

// Unimplemented: AuthorityCode

// Unimplemented: AuthorityName

// Unimplemented: SetProjection

// Unimplemented: SetProjParm

// Unimplemented: ProjParm

// Unimplemented: SetNormProjParm

// Unimplemented: NormProjParm

// Unimplemented: SetUTM

// Unimplemented: UTMZone

// Unimplemented: SetStatePlane

// Unimplemented: SetStatePlaneWithUnits

// Unimplemented: AutoIdentifyEPSG

// Unimplemented: EPSGTreatsAsLatLong

// Unimplemented: Axis

// Unimplemented: SetACEA

// Unimplemented: SetAE

// Unimplemented: SetBonne

// Unimplemented: SetCEA

// Unimplemented: SetCS

// Unimplemented: SetEC

// Unimplemented: SetEckert

// Unimplemented: SetEckertIV

// Unimplemented: SetEckertVI

// Unimplemented: SetEquirectangular

// Unimplemented: SetEquirectangular2

// Unimplemented: SetGS

// Unimplemented: SetGH

// Unimplemented: SetIGH

// Unimplemented: SetGEOS

// Unimplemented: SetGaussSchreiberTMercator

// Unimplemented: SetGnomonic

// Unimplemented: SetOM

// Unimplemented: SetHOM

// Unimplemented: SetHOM2PNO

// Unimplemented: SetIWMPolyconic

// Unimplemented: SetKrovak

// Unimplemented: SetLAEA

// Unimplemented: SetLCC

// Unimplemented: SetLCC1SP

// Unimplemented: SetLCCB

// Unimplemented: SetMC

// Unimplemented: SetMercator

// Unimplemented: SetMollweide

// Unimplemented: SetNZMG

// Unimplemented: SetOS

// Unimplemented: SetOrthographic

// Unimplemented: SetPolyconic

// Unimplemented: SetPS

// Unimplemented: SetRobinson

// Unimplemented: SetSinusoidal

// Unimplemented: SetStereographic

// Unimplemented: SetSOC

// Unimplemented: SetTM

// Unimplemented: SetTMVariant

// Unimplemented: SetTMG

// Unimplemented: SetTMSO

// Unimplemented: SetVDG

// Unimplemented: SetWagner

// Cleanup cached SRS related memory
func CleanupSR() {
	C.OSRCleanup()
}

/* -------------------------------------------------------------------- */
/*      Coordinate transformation functions.                            */
/* -------------------------------------------------------------------- */

type CoordinateTransform struct {
	cval C.OGRCoordinateTransformationH
}

// Create a new CoordinateTransform
func CreateCoordinateTransform(
	source SpatialReference,
	dest SpatialReference,
) CoordinateTransform {
	ct := C.OCTNewCoordinateTransformation(source.cval, dest.cval)
	return CoordinateTransform{ct}
}

// Destroy CoordinateTransform
func (ct CoordinateTransform) Destroy() {
	C.OCTDestroyCoordinateTransformation(ct.cval)
}

// Fetch list of possible projection methods
func ProjectionMethods() []string {
	p := C.OPTGetProjectionMethods()
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

// Fetch the parameters for a given projection method
func ParameterList(method string) (params []string, name string) {
	cMethod := C.CString(method)
	defer C.free(unsafe.Pointer(cMethod))

	var cName *C.char

	p := C.OPTGetParameterList(cMethod, &cName)

	name = C.GoString(cName)

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

	return strings, name
}

// Fetch information about a single parameter of a projection method
func ParameterInfo(
	projectionMethod, parameterName string,
) (
	username, paramType string, 
	defaultValue float64,
	ok bool,
) {
	cMethod := C.CString(projectionMethod)
	defer C.free(unsafe.Pointer(cMethod))

	cName := C.CString(parameterName)
	defer C.free(unsafe.Pointer(cName))

	var cUserName *C.char
	var cParamType *C.char
	var cDefaultValue C.double	
	
	success := C.OPTGetParameterInfo(
		cMethod,
		cName,
		&cUserName,
		&cParamType,
		&cDefaultValue)
	return C.GoString(cUserName), C.GoString(cParamType), float64(cDefaultValue), success != 0
}

package gdal

/*
#include "go_gdal.h"
#include "gdal_version.h"
#include "ogr_srs_api.h"
*/
import "C"
import (
	"reflect"
	"unsafe"
)

type AxisMappingStrategy uint32

const (
	OAMS_TraditionalGisOrder = AxisMappingStrategy(C.OAMS_TRADITIONAL_GIS_ORDER)
	OAMS_AuthorityCompliant  = AxisMappingStrategy(C.OAMS_AUTHORITY_COMPLIANT)
	OAMS_Custom              = AxisMappingStrategy(C.OAMS_CUSTOM)
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

func (sr SpatialReference) SetAxisMappingStrategy(strategy AxisMappingStrategy) {
	C.OSRSetAxisMappingStrategy(sr.cval, C.OSRAxisMappingStrategy(strategy))
}

// Initialize SRS based on WKT string
func (sr SpatialReference) FromWKT(wkt string) error {
	cString := C.CString(wkt)
	defer C.free(unsafe.Pointer(cString))
	return C.OSRImportFromWkt(sr.cval, &cString).Err()
}

// Export coordinate system to WKT
func (sr SpatialReference) ToWKT() (string, error) {
	var p *C.char
	err := C.OSRExportToWkt(sr.cval, &p).Err()
	wkt := C.GoString(p)
	return wkt, err
}

// Export coordinate system to a nicely formatted WKT string
func (sr SpatialReference) ToPrettyWKT(simplify bool) (string, error) {
	var p *C.char
	err := C.OSRExportToPrettyWkt(
		sr.cval, &p, BoolToCInt(simplify),
	).Err()
	wkt := C.GoString(p)
	return wkt, err
}

// Initialize SRS based on EPSG code
func (sr SpatialReference) FromEPSG(code int) error {
	return C.OSRImportFromEPSG(sr.cval, C.int(code)).Err()
}

// Initialize SRS based on EPSG code, using EPSG lat/long ordering
func (sr SpatialReference) FromEPSGA(code int) error {
	return C.OSRImportFromEPSGA(sr.cval, C.int(code)).Err()
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
	return C.OSRValidate(sr.cval).Err()
}

// Import PROJ.4 coordinate string
func (sr SpatialReference) FromProj4(input string) error {
	cString := C.CString(input)
	defer C.free(unsafe.Pointer(cString))
	return C.OSRImportFromProj4(sr.cval, cString).Err()
}

// Export coordinate system in PROJ.4 format
func (sr SpatialReference) ToProj4() (string, error) {
	var p *C.char
	err := C.OSRExportToProj4(sr.cval, &p).Err()
	proj4 := C.GoString(p)
	return proj4, err
}

// Import coordinate system from ESRI .prj formats
func (sr SpatialReference) FromESRI(input string) error {
	cString := C.CString(input)
	defer C.free(unsafe.Pointer(cString))
	return C.OSRImportFromProj4(sr.cval, cString).Err()
}

// Import coordinate system from PCI projection definition
func (sr SpatialReference) FromPCI(proj, units string, params []float64) error {
	cProj := C.CString(proj)
	defer C.free(unsafe.Pointer(cProj))
	cUnits := C.CString(units)
	defer C.free(unsafe.Pointer(cUnits))

	return C.OSRImportFromPCI(
		sr.cval,
		cProj,
		cUnits,
		(*C.double)(unsafe.Pointer(&params[0])),
	).Err()
}

// Import coordinate system from USGS projection definition
func (sr SpatialReference) FromUSGS(projsys, zone int, params []float64, datum int) error {
	return C.OSRImportFromUSGS(
		sr.cval,
		C.long(projsys),
		C.long(zone),
		(*C.double)(unsafe.Pointer(&params[0])),
		C.long(datum),
	).Err()
}

// Import coordinate system from XML format (GML only currently)
func (sr SpatialReference) FromXML(xml string) error {
	cXml := C.CString(xml)
	defer C.free(unsafe.Pointer(cXml))
	return C.OSRImportFromXML(sr.cval, cXml).Err()
}

// Import coordinate system from ERMapper projection definitions
func (sr SpatialReference) FromERM(proj, datum, units string) error {
	cProj := C.CString(proj)
	defer C.free(unsafe.Pointer(cProj))
	cDatum := C.CString(datum)
	defer C.free(unsafe.Pointer(cDatum))
	cUnits := C.CString(units)
	defer C.free(unsafe.Pointer(cUnits))

	return C.OSRImportFromERM(sr.cval, cProj, cDatum, cUnits).Err()
}

// Import coordinate system from a URL
func (sr SpatialReference) FromURL(url string) error {
	cURL := C.CString(url)
	defer C.free(unsafe.Pointer(cURL))
	return C.OSRImportFromXML(sr.cval, cURL).Err()
}

// Export coordinate system in PCI format
func (sr SpatialReference) ToPCI() (proj, units string, params []float64, errVal error) {
	var p, u *C.char
	err := C.OSRExportToPCI(
		sr.cval, &p, &u, (**C.double)(unsafe.Pointer(&params[0])),
	).Err()
	header := (*reflect.SliceHeader)(unsafe.Pointer(&params))
	header.Cap = 17
	header.Len = 17
	defer C.free(unsafe.Pointer(p))
	defer C.free(unsafe.Pointer(u))
	return C.GoString(p), C.GoString(u), params, err
}

// Export coordinate system to USGS GCTP projection definition
func (sr SpatialReference) ToUSGS() (proj, zone int, params []float64, datum int, errVal error) {
	err := C.OSRExportToUSGS(
		sr.cval,
		(*C.long)(unsafe.Pointer(&proj)),
		(*C.long)(unsafe.Pointer(&zone)),
		(**C.double)(unsafe.Pointer(&params[0])),
		(*C.long)(unsafe.Pointer(&datum)),
	).Err()

	header := (*reflect.SliceHeader)(unsafe.Pointer(&params))
	header.Cap = 15
	header.Len = 15

	return proj, zone, params, datum, err
}

// Export coordinate system in XML format
func (sr SpatialReference) ToXML() (xml string, errVal error) {
	var x *C.char
	err := C.OSRExportToXML(sr.cval, &x, nil).Err()
	defer C.free(unsafe.Pointer(x))
	return C.GoString(x), err
}

// Export coordinate system in Mapinfo style CoordSys format
func (sr SpatialReference) ToMICoordSys() (output string, errVal error) {
	var x *C.char
	err := C.OSRExportToMICoordSys(sr.cval, &x).Err()
	defer C.free(unsafe.Pointer(x))
	return C.GoString(x), err
}

// Export coordinate system in ERMapper format
// Unimplemented: ToERM

// Convert in place to ESRI WKT format
func (sr SpatialReference) MorphToESRI() error {
	return C.OSRMorphToESRI(sr.cval).Err()
}

// Convert in place from ESRI WKT format
func (sr SpatialReference) MorphFromESRI() error {
	return C.OSRMorphFromESRI(sr.cval).Err()
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
	return C.OSRSetAttrValue(sr.cval, cPath, cValue).Err()
}

// Set the angular units for the geographic coordinate system
func (sr SpatialReference) SetAngularUnits(units string, radians float64) error {
	cUnits := C.CString(units)
	defer C.free(unsafe.Pointer(cUnits))
	return C.OSRSetAngularUnits(sr.cval, cUnits, C.double(radians)).Err()
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
	return C.OSRSetLinearUnits(sr.cval, cName, C.double(toMeters)).Err()
}

// Set the linear units for the target node
func (sr SpatialReference) SetTargetLinearUnits(target, units string, toMeters float64) error {
	cTarget := C.CString(target)
	defer C.free(unsafe.Pointer(cTarget))
	cUnits := C.CString(units)
	defer C.free(unsafe.Pointer(cUnits))
	return C.OSRSetTargetLinearUnits(sr.cval, cTarget, cUnits, C.double(toMeters)).Err()
}

// Set the linear units for the target node and update all existing linear parameters
func (sr SpatialReference) SetLinearUnitsAndUpdateParameters(name string, toMeters float64) error {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	return C.OSRSetLinearUnitsAndUpdateParameters(sr.cval, cName, C.double(toMeters)).Err()
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
	return C.OSRSetLocalCS(sr.cval, cName).Err()
}

// Set the user visible projected CS name
func (sr SpatialReference) SetProjectedCS(name string) error {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	return C.OSRSetProjCS(sr.cval, cName).Err()
}

// Set the user visible geographic CS name
func (sr SpatialReference) SetGeocentricCS(name string) error {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	return C.OSRSetGeocCS(sr.cval, cName).Err()
}

// Set geographic CS based on well known name
func (sr SpatialReference) SetWellKnownGeographicCS(name string) error {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	return C.OSRSetWellKnownGeogCS(sr.cval, cName).Err()
}

// Set spatial reference from various text formats
func (sr SpatialReference) SetFromUserInput(name string) error {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	return C.OSRSetFromUserInput(sr.cval, cName).Err()
}

// Copy geographic CS from another spatial reference
func (sr SpatialReference) CopyGeographicCSFrom(other SpatialReference) error {
	return C.OSRCopyGeogCSFrom(sr.cval, other.cval).Err()
}

// Set the Bursa-Wolf conversion to WGS84
func (sr SpatialReference) SetTOWGS84(dx, dy, dz, ex, ey, ez, ppm float64) error {
	return C.OSRSetTOWGS84(
		sr.cval,
		C.double(dx),
		C.double(dy),
		C.double(dz),
		C.double(ex),
		C.double(ey),
		C.double(ez),
		C.double(ppm),
	).Err()
}

// Fetch the TOWGS84 parameters if available
func (sr SpatialReference) TOWGS84() (coeff [7]float64, err error) {
	err = C.OSRGetTOWGS84(sr.cval, (*C.double)(unsafe.Pointer(&coeff[0])), 7).Err()
	return
}

// Setup a compound coordinate system
func (sr SpatialReference) SetCompoundCS(
	name string,
	horizontal, vertical SpatialReference,
) error {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	return C.OSRSetCompoundCS(sr.cval, cName, horizontal.cval, vertical.cval).Err()
}

// Set geographic coordinate system
func (sr SpatialReference) SetGeographicCS(
	geogName, datumName, spheroidName string,
	semiMajor, flattening float64,
	pmName string,
	offset float64,
	angularUnits string,
	toRadians float64,
) error {
	cGeogName := C.CString(geogName)
	defer C.free(unsafe.Pointer(cGeogName))
	cDatumName := C.CString(datumName)
	defer C.free(unsafe.Pointer(cDatumName))
	cSpheroidName := C.CString(spheroidName)
	defer C.free(unsafe.Pointer(cSpheroidName))
	cPMName := C.CString(pmName)
	defer C.free(unsafe.Pointer(cPMName))
	cAngularUnits := C.CString(angularUnits)
	defer C.free(unsafe.Pointer(cAngularUnits))
	return C.OSRSetGeogCS(
		sr.cval,
		cGeogName,
		cDatumName,
		cSpheroidName,
		C.double(semiMajor),
		C.double(flattening),
		cPMName,
		C.double(offset),
		cAngularUnits,
		C.double(toRadians),
	).Err()
}

// Set up the vertical coordinate system
func (sr SpatialReference) SetVerticalCS(csName, datumName string, datumType int) error {
	cCSName := C.CString(csName)
	defer C.free(unsafe.Pointer(cCSName))
	cDatumName := C.CString(datumName)
	defer C.free(unsafe.Pointer(cDatumName))
	return C.OSRSetVertCS(sr.cval, cCSName, cDatumName, C.int(datumType)).Err()
}

// Get spheroid semi-major axis
func (sr SpatialReference) SemiMajorAxis() (float64, error) {
	var cErr C.OGRErr
	axis := C.OSRGetSemiMajor(sr.cval, &cErr)
	return float64(axis), cErr.Err()
}

// Get spheroid semi-minor axis
func (sr SpatialReference) SemiMinorAxis() (float64, error) {
	var cErr C.OGRErr
	axis := C.OSRGetSemiMinor(sr.cval, &cErr)
	return float64(axis), cErr.Err()
}

// Get spheroid inverse flattening axis
func (sr SpatialReference) InverseFlattening() (float64, error) {
	var cErr C.OGRErr
	flat := C.OSRGetInvFlattening(sr.cval, &cErr)
	return float64(flat), cErr.Err()
}

// Sets the authority for a node
func (sr SpatialReference) SetAuthority(target, authority string, code int) error {
	cTarget := C.CString(target)
	defer C.free(unsafe.Pointer(cTarget))
	cAuthority := C.CString(authority)
	defer C.free(unsafe.Pointer(cAuthority))
	return C.OSRSetAuthority(sr.cval, cTarget, cAuthority, C.int(code)).Err()
}

// Get the authority code for a node
func (sr SpatialReference) AuthorityCode(target string) string {
	cTarget := C.CString(target)
	defer C.free(unsafe.Pointer(cTarget))
	code := C.OSRGetAuthorityCode(sr.cval, cTarget)
	return C.GoString(code)
}

// Get the authority name for a node
func (sr SpatialReference) AuthorityName(target string) string {
	cTarget := C.CString(target)
	defer C.free(unsafe.Pointer(cTarget))
	code := C.OSRGetAuthorityName(sr.cval, cTarget)
	return C.GoString(code)
}

// Set a projection by name
func (sr SpatialReference) SetProjectionByName(name string) error {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	return C.OSRSetProjection(sr.cval, cName).Err()
}

// Set a projection parameter value
func (sr SpatialReference) SetProjectionParameter(name string, value float64) error {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	return C.OSRSetProjParm(sr.cval, cName, C.double(value)).Err()
}

// Fetch a projection parameter value
func (sr SpatialReference) ProjectionParameter(name string, defaultValue float64) (float64, error) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	var cErr C.OGRErr
	value := C.OSRGetProjParm(sr.cval, cName, C.double(defaultValue), &cErr)
	return float64(value), cErr.Err()
}

// Set a projection parameter with a normalized value
func (sr SpatialReference) SetNormalizedProjectionParameter(name string, value float64) error {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	return C.OSRSetNormProjParm(sr.cval, cName, C.double(value)).Err()
}

// Fetch a normalized projection parameter value
func (sr SpatialReference) NormalizedProjectionParameter(
	name string, defaultValue float64,
) (float64, error) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	var cErr C.OGRErr
	value := C.OSRGetProjParm(sr.cval, cName, C.double(defaultValue), &cErr)
	return float64(value), cErr.Err()
}

// Set UTM projection definition
func (sr SpatialReference) SetUTM(zone int, north bool) error {
	return C.OSRSetUTM(sr.cval, C.int(zone), BoolToCInt(north)).Err()
}

// Get UTM zone information
func (sr SpatialReference) UTMZone() (zone int, north bool) {
	var northInt C.int
	cZone := C.OSRGetUTMZone(sr.cval, &northInt)
	return int(cZone), northInt != 0
}

// Set State Plane projection definition
func (sr SpatialReference) SetStatePlane(zone int, nad83 bool) error {
	return C.OSRSetStatePlane(sr.cval, C.int(zone), BoolToCInt(nad83)).Err()
}

// Set State Plane projection definition
func (sr SpatialReference) SetStatePlaneWithUnits(
	zone int,
	nad83 bool,
	unitName string,
	factor float64,
) error {
	cUnitName := C.CString(unitName)
	defer C.free(unsafe.Pointer(cUnitName))
	return C.OSRSetStatePlaneWithUnits(
		sr.cval,
		C.int(zone),
		BoolToCInt(nad83),
		cUnitName,
		C.double(factor),
	).Err()
}

// Set EPSG authority info if possible
func (sr SpatialReference) AutoIdentifyEPSG() error {
	return C.OSRAutoIdentifyEPSG(sr.cval).Err()
}

// Return true if EPSG feels this coordinate system should be treated as having lat/long coordinate ordering
func (sr SpatialReference) EPSGTreatsAsLatLong() bool {
	val := C.OSREPSGTreatsAsLatLong(sr.cval)
	return val != 0
}

// Fetch the orientation of one axis
// Unimplemented: Axis

// Set to Albers Conic Equal Area
func (sr SpatialReference) SetACEA(
	stdp1, stdp2, centerLat, centerLong, falseEasting, falseNorthing float64,
) error {
	return C.OSRSetACEA(
		sr.cval,
		C.double(stdp1),
		C.double(stdp2),
		C.double(centerLat),
		C.double(centerLong),
		C.double(falseEasting),
		C.double(falseNorthing),
	).Err()
}

// Set to Azimuthal Equidistant
func (sr SpatialReference) SetAE(centerLat, centerLong, falseEasting, falseNorthing float64) error {
	return C.OSRSetAE(
		sr.cval,
		C.double(centerLat),
		C.double(centerLong),
		C.double(falseEasting),
		C.double(falseNorthing),
	).Err()
}

// Set to Bonne
func (sr SpatialReference) SetBonne(standardParallel, centralMeridian, falseEasting, falseNorthing float64) error {
	return C.OSRSetBonne(
		sr.cval,
		C.double(standardParallel),
		C.double(centralMeridian),
		C.double(falseEasting),
		C.double(falseNorthing),
	).Err()
}

// Set to Cylindrical Equal Area
func (sr SpatialReference) SetCEA(stdp1, centralMeridian, falseEasting, falseNorthing float64) error {
	return C.OSRSetCEA(
		sr.cval,
		C.double(stdp1),
		C.double(centralMeridian),
		C.double(falseEasting),
		C.double(falseNorthing),
	).Err()
}

// Set to Cassini-Soldner
func (sr SpatialReference) SetCS(centerLat, centerLong, falseEasting, falseNorthing float64) error {
	return C.OSRSetCS(
		sr.cval,
		C.double(centerLat),
		C.double(centerLong),
		C.double(falseEasting),
		C.double(falseNorthing),
	).Err()
}

// Set to Equidistant Conic
func (sr SpatialReference) SetEC(
	stdp1, stdp2, centerLat, centerLong, falseEasting, falseNorthing float64,
) error {
	return C.OSRSetEC(
		sr.cval,
		C.double(stdp1),
		C.double(stdp2),
		C.double(centerLat),
		C.double(centerLong),
		C.double(falseEasting),
		C.double(falseNorthing),
	).Err()
}

// Set to Eckert I-VI
func (sr SpatialReference) SetEckert(variation int, centralMeridian, falseEasting, falseNorthing float64) error {
	return C.OSRSetEckert(
		sr.cval,
		C.int(variation),
		C.double(centralMeridian),
		C.double(falseEasting),
		C.double(falseNorthing),
	).Err()
}

// Set to Equirectangular
func (sr SpatialReference) SetEquirectangular(
	centerLat, centerLong, falseEasting, falseNorthing float64,
) error {
	return C.OSRSetEquirectangular(
		sr.cval,
		C.double(centerLat),
		C.double(centerLong),
		C.double(falseEasting),
		C.double(falseNorthing),
	).Err()
}

// Set to Equirectangular (generalized form)
func (sr SpatialReference) SetEquirectangularGeneralized(
	centerLat, centerLong, psuedoStdParallel, falseEasting, falseNorthing float64,
) error {
	return C.OSRSetEquirectangular2(
		sr.cval,
		C.double(centerLat),
		C.double(centerLong),
		C.double(psuedoStdParallel),
		C.double(falseEasting),
		C.double(falseNorthing),
	).Err()
}

// Set to Gall Stereographic
func (sr SpatialReference) SetGS(centralMeridian, falseEasting, falseNorthing float64) error {
	return C.OSRSetGS(
		sr.cval,
		C.double(centralMeridian),
		C.double(falseEasting),
		C.double(falseNorthing),
	).Err()
}

// Set to Goode Homolosine
func (sr SpatialReference) SetGH(centralMeridian, falseEasting, falseNorthing float64) error {
	return C.OSRSetGH(
		sr.cval,
		C.double(centralMeridian),
		C.double(falseEasting),
		C.double(falseNorthing),
	).Err()
}

// Set to Interrupted Goode Homolosine
func (sr SpatialReference) SetIGH() error {
	return C.OSRSetIGH(sr.cval).Err()
}

// Set to GEOS - Geostationary Satellite View
func (sr SpatialReference) SetGEOS(
	centralMeridian, satelliteHeight, falseEasting, falseNorthing float64,
) error {
	return C.OSRSetGEOS(
		sr.cval,
		C.double(centralMeridian),
		C.double(satelliteHeight),
		C.double(falseEasting),
		C.double(falseNorthing),
	).Err()
}

// Set to Gauss Schreiber Transverse Mercator
func (sr SpatialReference) SetGSTM(
	centerLat, centerLong, scale, falseEasting, falseNorthing float64,
) error {
	return C.OSRSetGaussSchreiberTMercator(
		sr.cval,
		C.double(centerLat),
		C.double(centerLong),
		C.double(scale),
		C.double(falseEasting),
		C.double(falseNorthing),
	).Err()
}

// Set to gnomonic
func (sr SpatialReference) SetGnomonic(
	centerLat, centerLong, falseEasting, falseNorthing float64,
) error {
	return C.OSRSetGnomonic(
		sr.cval,
		C.double(centerLat),
		C.double(centerLong),
		C.double(falseEasting),
		C.double(falseNorthing),
	).Err()
}

// Set to Hotine Oblique Mercator projection using azimuth angle
func (sr SpatialReference) SetHOM(
	centerLat, centerLong, azimuth, rectToSkew, scale, falseEasting, falseNorthing float64,
) error {
	return C.OSRSetHOM(
		sr.cval,
		C.double(centerLat),
		C.double(centerLong),
		C.double(azimuth),
		C.double(rectToSkew),
		C.double(scale),
		C.double(falseEasting),
		C.double(falseNorthing),
	).Err()
}

// Set to Hotine Oblique Mercator projection using two points on projection centerline
func (sr SpatialReference) SetHOM2PNO(
	centerLat, lat1, long1, lat2, long2, scale, falseEasting, falseNorthing float64,
) error {
	return C.OSRSetHOM2PNO(
		sr.cval,
		C.double(centerLat),
		C.double(lat1),
		C.double(long1),
		C.double(lat2),
		C.double(long2),
		C.double(scale),
		C.double(falseEasting),
		C.double(falseNorthing),
	).Err()
}

// Set to International Map of the World Polyconic
func (sr SpatialReference) SetIWMPolyconic(
	lat1, lat2, centerLong, falseEasting, falseNorthing float64,
) error {
	return C.OSRSetIWMPolyconic(
		sr.cval,
		C.double(lat1),
		C.double(lat2),
		C.double(centerLong),
		C.double(falseEasting),
		C.double(falseNorthing),
	).Err()
}

// Set to Krovak Oblique Conic Conformal
func (sr SpatialReference) SetKrovak(
	centerLat, centerLong, azimuth, psuedoStdParallel, scale, falseEasting, falseNorthing float64,
) error {
	return C.OSRSetKrovak(
		sr.cval,
		C.double(centerLat),
		C.double(centerLong),
		C.double(azimuth),
		C.double(psuedoStdParallel),
		C.double(scale),
		C.double(falseEasting),
		C.double(falseNorthing),
	).Err()
}

// Set to Lambert Azimuthal Equal Area
func (sr SpatialReference) SetLAEA(
	centerLat, centerLong, falseEasting, falseNorthing float64,
) error {
	return C.OSRSetLAEA(
		sr.cval,
		C.double(centerLat),
		C.double(centerLong),
		C.double(falseEasting),
		C.double(falseNorthing),
	).Err()
}

// Set to Lambert Conformal Conic
func (sr SpatialReference) SetLCC(
	stdp1, stdp2, centerLat, centerLong, falseEasting, falseNorthing float64,
) error {
	return C.OSRSetLCC(
		sr.cval,
		C.double(stdp1),
		C.double(stdp2),
		C.double(centerLat),
		C.double(centerLong),
		C.double(falseEasting),
		C.double(falseNorthing),
	).Err()
}

// Set to Lambert Conformal Conic (1 standard parallel)
func (sr SpatialReference) SetLCC1SP(
	centerLat, centerLong, scale, falseEasting, falseNorthing float64,
) error {
	return C.OSRSetLCC1SP(
		sr.cval,
		C.double(centerLat),
		C.double(centerLong),
		C.double(scale),
		C.double(falseEasting),
		C.double(falseNorthing),
	).Err()
}

// Set to Lambert Conformal Conic (Belgium)
func (sr SpatialReference) SetLCCB(
	stdp1, stdp2, centerLat, centerLong, falseEasting, falseNorthing float64,
) error {
	return C.OSRSetLCCB(
		sr.cval,
		C.double(stdp1),
		C.double(stdp2),
		C.double(centerLat),
		C.double(centerLong),
		C.double(falseEasting),
		C.double(falseNorthing),
	).Err()
}

// Set to Miller Cylindrical
func (sr SpatialReference) SetMC(
	centerLat, centerLong, falseEasting, falseNorthing float64,
) error {
	return C.OSRSetMC(
		sr.cval,
		C.double(centerLat),
		C.double(centerLong),
		C.double(falseEasting),
		C.double(falseNorthing),
	).Err()
}

// Set to Mercator
func (sr SpatialReference) SetMercator(
	centerLat, centerLong, scale, falseEasting, falseNorthing float64,
) error {
	return C.OSRSetMercator(
		sr.cval,
		C.double(centerLat),
		C.double(centerLong),
		C.double(scale),
		C.double(falseEasting),
		C.double(falseNorthing),
	).Err()
}

// Set tp Mollweide
func (sr SpatialReference) SetMollweide(
	centralMeridian, falseEasting, falseNorthing float64,
) error {
	return C.OSRSetMollweide(
		sr.cval,
		C.double(centralMeridian),
		C.double(falseEasting),
		C.double(falseNorthing),
	).Err()
}

// Set to New Zealand Map Grid
func (sr SpatialReference) SetNZMG(
	centerLat, centerLong, falseEasting, falseNorthing float64,
) error {
	return C.OSRSetNZMG(
		sr.cval,
		C.double(centerLat),
		C.double(centerLong),
		C.double(falseEasting),
		C.double(falseNorthing),
	).Err()
}

// Set to Oblique Stereographic
func (sr SpatialReference) SetOS(
	originLat, meridian, scale, falseEasting, falseNorthing float64,
) error {
	return C.OSRSetOS(
		sr.cval,
		C.double(originLat),
		C.double(meridian),
		C.double(scale),
		C.double(falseEasting),
		C.double(falseNorthing),
	).Err()
}

// Set to Orthographic
func (sr SpatialReference) SetOrthographic(
	centerLat, centerLong, falseEasting, falseNorthing float64,
) error {
	return C.OSRSetOrthographic(
		sr.cval,
		C.double(centerLat),
		C.double(centerLong),
		C.double(falseEasting),
		C.double(falseNorthing),
	).Err()
}

// Set to Polyconic
func (sr SpatialReference) SetPolyconic(
	centerLat, centerLong, falseEasting, falseNorthing float64,
) error {
	return C.OSRSetPolyconic(
		sr.cval,
		C.double(centerLat),
		C.double(centerLong),
		C.double(falseEasting),
		C.double(falseNorthing),
	).Err()
}

// Set to Polar Stereographic
func (sr SpatialReference) SetPS(
	centerLat, centerLong, scale, falseEasting, falseNorthing float64,
) error {
	return C.OSRSetPS(
		sr.cval,
		C.double(centerLat),
		C.double(centerLong),
		C.double(scale),
		C.double(falseEasting),
		C.double(falseNorthing),
	).Err()
}

// Set to Robinson
func (sr SpatialReference) SetRobinson(
	centerLong, falseEasting, falseNorthing float64,
) error {
	return C.OSRSetRobinson(
		sr.cval,
		C.double(centerLong),
		C.double(falseEasting),
		C.double(falseNorthing),
	).Err()
}

// Set to Sinusoidal
func (sr SpatialReference) SetSinusoidal(
	centerLong, falseEasting, falseNorthing float64,
) error {
	return C.OSRSetSinusoidal(
		sr.cval,
		C.double(centerLong),
		C.double(falseEasting),
		C.double(falseNorthing),
	).Err()
}

// Set to Stereographic
func (sr SpatialReference) SetStereographic(
	centerLat, centerLong, scale, falseEasting, falseNorthing float64,
) error {
	return C.OSRSetStereographic(
		sr.cval,
		C.double(centerLat),
		C.double(centerLong),
		C.double(scale),
		C.double(falseEasting),
		C.double(falseNorthing),
	).Err()
}

// Set to Swiss Oblique Cylindrical
func (sr SpatialReference) SetSOC(
	latitudeOfOrigin, centralMeridian, falseEasting, falseNorthing float64,
) error {
	return C.OSRSetSOC(
		sr.cval,
		C.double(latitudeOfOrigin),
		C.double(centralMeridian),
		C.double(falseEasting),
		C.double(falseNorthing),
	).Err()
}

// Set to Transverse Mercator
func (sr SpatialReference) SetTM(
	centerLat, centerLong, scale, falseEasting, falseNorthing float64,
) error {
	return C.OSRSetTM(
		sr.cval,
		C.double(centerLat),
		C.double(centerLong),
		C.double(scale),
		C.double(falseEasting),
		C.double(falseNorthing),
	).Err()
}

// Set to Transverse Mercator variant
func (sr SpatialReference) SetTMVariant(
	variantName string, centerLat, centerLong, scale, falseEasting, falseNorthing float64,
) error {
	cName := C.CString(variantName)
	defer C.free(unsafe.Pointer(cName))
	return C.OSRSetTMVariant(
		sr.cval,
		cName,
		C.double(centerLat),
		C.double(centerLong),
		C.double(scale),
		C.double(falseEasting),
		C.double(falseNorthing),
	).Err()
}

// Set to Tunisia Mining Grid
func (sr SpatialReference) SetTMG(
	centerLat, centerLong, falseEasting, falseNorthing float64,
) error {
	return C.OSRSetTMG(
		sr.cval,
		C.double(centerLat),
		C.double(centerLong),
		C.double(falseEasting),
		C.double(falseNorthing),
	).Err()
}

// Set to Transverse Mercator (South Oriented)
func (sr SpatialReference) SetTMSO(
	centerLat, centerLong, scale, falseEasting, falseNorthing float64,
) error {
	return C.OSRSetTMSO(
		sr.cval,
		C.double(centerLat),
		C.double(centerLong),
		C.double(scale),
		C.double(falseEasting),
		C.double(falseNorthing),
	).Err()
}

// Set to VanDerGrinten
func (sr SpatialReference) SetVDG(
	centerLong, falseEasting, falseNorthing float64,
) error {
	return C.OSRSetVDG(
		sr.cval,
		C.double(centerLong),
		C.double(falseEasting),
		C.double(falseNorthing),
	).Err()
}

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

func (ct CoordinateTransform) Transform(numPoints int, xPoints []float64, yPoints []float64, zPoints []float64) bool {
	val := C.OCTTransform(ct.cval, C.int(numPoints), (*C.double)(unsafe.Pointer(&xPoints[0])), (*C.double)(unsafe.Pointer(&yPoints[0])), (*C.double)(unsafe.Pointer(&zPoints[0])))
	return int(val) != 0
}

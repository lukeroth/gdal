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

// Unimplemented: ToProj4
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

// Unimplemented: FromPCI
// Unimplemented: FromUSGS
// Unimplemented: FromXML
// Unimplemented: FromERM
// Unimplemented: FromURL
// Unimplemented: ToPCI
// Unimplemented: ToUSGS
// Unimplemented: ToXML
// Unimplemented: ToMICoordSys
// Unimplemented: ToERM
// Unimplemented: MorphToESRI
// Unimplemented: MorphFromESRI
// Unimplemented: SetAttrValue
// Unimplemented: AttrValue
// Unimplemented: SetAngularUnits
// Unimplemented: AngularUnits
// Unimplemented: SetLinearUnits
// Unimplemented: SetTargetLinearUnits
// Unimplemented: SetLinearUnitsAndUpdateParameters
// Unimplemented: LinearUnits
// Unimplemented: TargetLinearUnits
// Unimplemented: PrimeMeridian
// Unimplemented: IsGeographic
// Unimplemented: IsLocal
// Unimplemented: IsProjected
// Unimplemented: IsCompound
// Unimplemented: IsGeocentric
// Unimplemented: IsVertical
// Unimplemented: IsSameGeogCS
// Unimplemented: IsSameVertCS
// Unimplemented: IsSame
// Unimplemented: SetLocalCS
// Unimplemented: SetProjCS
// Unimplemented: SetGeocCS
// Unimplemented: SetWellKnownGeogCS
// Unimplemented: SetFromUserInput
// Unimplemented: CopyGeogCSFrom
// Unimplemented: SetTOWGS84
// Unimplemented: TOWGS84
// Unimplemented: SetCompoundCS
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

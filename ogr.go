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

func init() {
	C.OGRRegisterAll()
}

/* -------------------------------------------------------------------- */
/*      Significant constants.                                          */
/* -------------------------------------------------------------------- */

// List of well known binary geometry types
type GeometryType int

const (
	GT_Unknown               = GeometryType(C.wkbUnknown)
	GT_Point                 = GeometryType(C.wkbPoint)
	GT_LineString            = GeometryType(C.wkbLineString)
	GT_Polygon               = GeometryType(C.wkbPolygon)
	GT_MultiPoint            = GeometryType(C.wkbMultiPoint)
	GT_MultiLineString       = GeometryType(C.wkbMultiLineString)
	GT_MultiPolygon          = GeometryType(C.wkbMultiPolygon)
	GT_GeometryCollection    = GeometryType(C.wkbGeometryCollection)
	GT_None                  = GeometryType(C.wkbNone)
	GT_LinearRing            = GeometryType(C.wkbLinearRing)
	GT_Point25D              = GeometryType(C.wkbPoint25D)
	GT_LineString25D         = GeometryType(C.wkbLineString25D)
	GT_Polygon25D            = GeometryType(C.wkbPolygon25D)
	GT_MultiPoint25D         = GeometryType(C.wkbMultiPoint25D)
	GT_MultiLineString25D    = GeometryType(C.wkbMultiLineString25D)
	GT_MultiPolygon25D       = GeometryType(C.wkbMultiPolygon25D)
	GT_GeometryCollection25D = GeometryType(C.wkbGeometryCollection25D)
)

/* -------------------------------------------------------------------- */
/*      Geometry functions                                              */
/* -------------------------------------------------------------------- */

type Geometry struct {
	cval C.OGRGeometryH
}

//Create a geometry object from its well known binary representation
func CreateFromWKB(wkb []uint8, srs SpatialReference, bytes int) (Geometry, error) {
	cString := (*C.uchar)(unsafe.Pointer(&wkb[0]))
	var newGeom Geometry
	err := C.OGR_G_CreateFromWkb(cString, srs.cval, &newGeom.cval, C.int(bytes))
	return newGeom, error(err)
}

//Create a geometry object from its well known text representation
func CreateFromWKT(wkt string, srs SpatialReference) (Geometry, error) {
	cString := C.CString(wkt)
	defer C.free(unsafe.Pointer(cString))
	var newGeom Geometry
	err := C.OGR_G_CreateFromWkt(&cString, srs.cval, &newGeom.cval)
	return newGeom, error(err)
}

// Destroy geometry object
func (geometry Geometry) Destroy() {
	C.OGR_G_DestroyGeometry(geometry.cval)
}

// Create an empty geometry of the desired type
func Create(geomType GeometryType) Geometry {
	geom := C.OGR_G_CreateGeometry(C.OGRwkbGeometryType(geomType))
	return Geometry{geom}
}

// Unimplemented: ApproximateArcAngles

// Convert to polygon
func (geom Geometry) ForceToPolygon() Geometry {
	newGeom := C.OGR_G_ForceToPolygon(geom.cval)
	return Geometry{newGeom}
}

// Convert to multipolygon
func (geom Geometry) ForceToMultiPolygon() Geometry {
	newGeom := C.OGR_G_ForceToMultiPolygon(geom.cval)
	return Geometry{newGeom}
}

// Convert to multipoint
func (geom Geometry) ForceToMultiPoint() Geometry {
	newGeom := C.OGR_G_ForceToMultiPoint(geom.cval)
	return Geometry{newGeom}
}

// Convert to multilinestring
func (geom Geometry) ForceToMultiLineString() Geometry {
	newGeom := C.OGR_G_ForceToMultiLineString(geom.cval)
	return Geometry{newGeom}
}

// Get the dimension of this geometry
func (geom Geometry) Dimension() int {
	dim := C.OGR_G_GetDimension(geom.cval)
	return int(dim)
}

// Get the dimension of the coordinates in this geometry
func (geom Geometry) CoordinateDimension() int {
	dim := C.OGR_G_GetCoordinateDimension(geom.cval)
	return int(dim)
}

// Set the dimension of the coordinates in this geometry
func (geom Geometry) SetCoordinateDimension(dim int) {
	C.OGR_G_SetCoordinateDimension(geom.cval, C.int(dim))
}

// Create a copy of this geometry
func (geom Geometry) Clone() Geometry {
	newGeom := C.OGR_G_Clone(geom.cval)
	return Geometry{newGeom}
}

// Unimplemented: GetEnvelope
// Unimplemented: GetEnvelope3D

// Assign a geometry from well known binary data
func (geom Geometry) FromWKB(wkb []uint8, bytes int) error {
	cString := (*C.uchar)(unsafe.Pointer(&wkb[0]))
	err := C.OGR_G_ImportFromWkb(geom.cval, cString, C.int(bytes))
	return error(err)
}

// Convert a geometry to well known binary data
// Unimplemented: ExportToWkb

// Returns size of related binary representation
func (geom Geometry) WKBSize() int {
	size := C.OGR_G_WkbSize(geom.cval)
	return int(size)
}

// Assign geometry object from its well known text representation
func (geom Geometry) FromWKT(wkt string) error {
	cString := C.CString(wkt)
	defer C.free(unsafe.Pointer(cString))
	err := C.OGR_G_ImportFromWkt(geom.cval, &cString)
	return error(err)
}

// Unimplemented: ExportToWkt
func (geom Geometry) ToWKT() (string, error) {
	var p *C.char
	err := C.OGR_G_ExportToWkt(geom.cval, &p)
	wkt := C.GoString(p)
	return wkt, error(err)
}

// Fetch geometry type
func (geom Geometry) Type() GeometryType {
	gt := C.OGR_G_GetGeometryType(geom.cval)
	return GeometryType(gt)
}

// Fetch geometry name
func (geom Geometry) Name() string {
	name := C.OGR_G_GetGeometryName(geom.cval)
	return C.GoString(name)
}

// Unimplemented: DumpReadable

// Convert geometry to strictly 2D
func (geom Geometry) FlattenTo2D() {
	C.OGR_G_FlattenTo2D(geom.cval)
}

// Force rings to be closed
func (geom Geometry) CloseRings() {
	C.OGR_G_CloseRings(geom.cval)
}

// Unimplemented: CreateFromGML
// Unimplemented: ExportToGML
// Unimplemented: ExportToGMLEx
// Unimplemented: ExportToKML
// Unimplemented: ExportToJson
// Unimplemented: ExportToJsonEx
// Unimplemented: SetSpatialReference
// Unimplemented: SpatialReference
// Unimplemented: Transform
// Unimplemented: TransformTo
// Unimplemented: Simplify
// Unimplemented: SimplifyPreserveTopology
// Unimplemented: Segmentize
// Unimplemented: Intersects
// Unimplemented: Equals
// Unimplemented: Disjoint
// Unimplemented: Touches
// Unimplemented: Crosses
// Unimplemented: Within
// Unimplemented: Contains
// Unimplemented: Overlaps
// Unimplemented: Boundary
// Unimplemented: ConvexHull
// Unimplemented: Buffer
// Unimplemented: Intersection
// Unimplemented: Union
// Unimplemented: UnionCascaded
// Unimplemented: PointOnSurface
// Unimplemented: Difference
// Unimplemented: SymDifference
// Unimplemented: Distance
// Unimplemented: Length
// Unimplemented: Area
// Unimplemented: Centroid
// Unimplemented: Empty
// Unimplemented: IsEmpty
// Unimplemented: IsValid
// Unimplemented: IsSimple
// Unimplemented: IsRing
// Unimplemented: Polygonize
// Unimplemented: SymmetricDifference
// Unimplemented: PointCount
// Unimplemented: Points
// Unimplemented: X
// Unimplemented: Y
// Unimplemented: Z
// Unimplemented: Point
// Unimplemented: SetPoint
// Unimplemented: SetPoint_2D
// Unimplemented: AddPoint
// Unimplemented: AddPoint_2D
// Unimplemented: GeometryCount
// Unimplemented: GeometryRef
// Unimplemented: AddGeometry
// Unimplemented: AddGeometryDirectly
// Unimplemented: RemoveGeometry
// Unimplemented: BuildPolygonFromEdges

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

/* -------------------------------------------------------------------- */
/*      Field definition functions                                      */
/* -------------------------------------------------------------------- */

type FieldDefn struct {
	cval C.OGRFieldDefnH
}

// Unimplemented: Create
// Unimplemented: Destroy
// Unimplemented: Name
// Unimplemented: SetName
// Unimplemented: Type
// Unimplemented: SetType
// Unimplemented: Justify
// Unimplemented: SetJustify
// Unimplemented: Width
// Unimplemented: SetWidth
// Unimplemented: Precision
// Unimplemented: SetPrecision
// Unimplemented: Set
// Unimplemented: IsIgnorned
// Unimplemented: SetIgnored
// Unimplemented: FieldTypeName

/* -------------------------------------------------------------------- */
/*      Feature definition functions                                    */
/* -------------------------------------------------------------------- */

type FeatureDefn struct {
	cval C.OGRFeatureDefnH
}

// Unimplemented: Create
// Unimplemented: Destroy
// Unimplemented: Release
// Unimplemented: Name
// Unimplemented: FieldCount
// Unimplemented: FieldDefn
// Unimplemented: FieldIndex
// Unimplemented: AddFieldDefn
// Unimplemented: DeleteFieldDefn
// Unimplemented: GeomType
// Unimplemented: SetGeomType
// Unimplemented: IsGeometryIgnored
// Unimplemented: SetGeometryIgnored
// Unimplemented: IsStyleIgnored
// Unimplemented: SetStyleIgnored
// Unimplemented: Reference
// Unimplemented: Dereference
// Unimplemented: ReferenceCount

/* -------------------------------------------------------------------- */
/*      Feature functions                                               */
/* -------------------------------------------------------------------- */

type Feature struct {
	cval C.OGRFeatureH
}

// Unimplemented: Create
// Unimplemented: Destroy
// Unimplemented: DefnRef
// Unimplemented: SetGeometryDirectly
// Unimplemented: SetGeometryDirectly
// Unimplemented: GeometryRef
// Unimplemented: StealGeometry
// Unimplemented: Clone
// Unimplemented: Equal
// Unimplemented: FieldCount
// Unimplemented: FieldDefnRef
// Unimplemented: FieldIndex
// Unimplemented: IsFieldSet
// Unimplemented: UnsetField
// Unimplemented: RawFieldRef
// Unimplemented: FieldAsInteger
// Unimplemented: FieldAsDouble
// Unimplemented: FieldAsString
// Unimplemented: FieldAsIntegerList
// Unimplemented: FieldAsDoubleList
// Unimplemented: FieldAsStringList
// Unimplemented: FieldAsBinary
// Unimplemented: FieldAsDateTime
// Unimplemented: SetFieldInteger
// Unimplemented: SetFieldDouble
// Unimplemented: SetFieldString
// Unimplemented: SetFieldIntegerList
// Unimplemented: SetFieldDoubleList
// Unimplemented: SetFieldStringList
// Unimplemented: SetFieldRaw
// Unimplemented: SetFieldBinary
// Unimplemented: SetFieldDateTime
// Unimplemented: FID
// Unimplemented: SetFID
// Unimplemented: DumpReadable
// Unimplemented: SetFrom
// Unimplemented: SetFromWithMap
// Unimplemented: StyleString
// Unimplemented: SetStyleString
// Unimplemented: SetStyleStringDirectly

/* -------------------------------------------------------------------- */
/*      Layer functions                                                 */
/* -------------------------------------------------------------------- */

type Layer struct {
	cval C.OGRLayerH
}

// Unimplemented: Name
// Unimplemented: GeomType
// Unimplemented: SpatialFilter
// Unimplemented: SetSpatialFilter
// Unimplemented: SetSpatialFilterRect
// Unimplemented: SetAttributeFilter
// Unimplemented: ResetReading
// Unimplemented: NextFeature
// Unimplemented: SetNextByIndex
// Unimplemented: Feature
// Unimplemented: SetFeature
// Unimplemented: CreateFeature
// Unimplemented: DeleteFeature
// Unimplemented: LayerDefn
// Unimplemented: SpatialRef
// Unimplemented: FeatureCount
// Unimplemented: Extent
// Unimplemented: TestCapability
// Unimplemented: CreateField
// Unimplemented: DeleteField
// Unimplemented: ReorderFields
// Unimplemented: ReorderField
// Unimplemented: AlterFieldDefn
// Unimplemented: StartTransaction
// Unimplemented: CommitTransaction
// Unimplemented: RollbackTransaction
// Unimplemented: SyncToDisk
// Unimplemented: FIDColumn
// Unimplemented: GeometryColumn
// Unimplemented: SetIgnoredFields
// Unimplemented: Intersection
// Unimplemented: Union
// Unimplemented: SymDifference
// Unimplemented: Identity
// Unimplemented: Update
// Unimplemented: Clip
// Unimplemented: Erase

/* -------------------------------------------------------------------- */
/*      Data source functions                                           */
/* -------------------------------------------------------------------- */

type DataSource struct {
	cval C.OGRDataSourceH
}

// Unimplemented: Open
// Unimplemented: Release
// Unimplemented: Destroy
// Unimplemented: Name
// Unimplemented: LayerCount
// Unimplemented: Layer
// Unimplemented: LayerByName
// Unimplemented: DeleteLayer
// Unimplemented: Driver
// Unimplemented: CreateLayer
// Unimplemented: CopyLayer
// Unimplemented: TestCapability
// Unimplemented: ExecuteSQL
// Unimplemented: ReleaseResultSet
// Unimplemented: SyncToDisk

/* -------------------------------------------------------------------- */
/*      Driver functions                                                */
/* -------------------------------------------------------------------- */

type OGRDriver struct {
	cval C.OGRSFDriverH
}

// Unimplemented: Name
// Unimplemented: Open
// Unimplemented: TestCapability
// Unimplemented: CreateDataSource
// Unimplemented: CopyDataSource
// Unimplemented: DeleteDataSource
// Unimplemented: Register
// Unimplemented: Deregister
// Unimplemented: DriverCount
// Unimplemented: DriverByIndex
// Unimplemented: DriverByName

/* -------------------------------------------------------------------- */
/*      Style manager functions                                         */
/* -------------------------------------------------------------------- */

type StyleMgr struct {
	cval C.OGRStyleMgrH
}

type StyleTool struct {
	cval C.OGRStyleToolH
}

type StyleTable struct {
	cval C.OGRStyleTableH
}

// Unimplemented: CreateStyleManager
// Unimplemented: Destroy
// Unimplemented: InitFromFeature
// Unimplemented: InitStyleString
// Unimplemented: PartCount
// Unimplemented: PartCount
// Unimplemented: AddPart
// Unimplemented: AddStyle

// Unimplemented: CreateStyleTool
// Unimplemented: Destroy
// Unimplemented: Type
// Unimplemented: Unit
// Unimplemented: SetUnit
// Unimplemented: ParamStr
// Unimplemented: ParamNum
// Unimplemented: ParamDbl
// Unimplemented: SetParamStr
// Unimplemented: SetParamNum
// Unimplemented: SetParamDbl
// Unimplemented: StyleString
// Unimplemented: RGBFromString

// Unimplemented: CreateStyleTable
// Unimplemented: Destroy
// Unimplemented: Save
// Unimplemented: Load
// Unimplemented: Find
// Unimplemented: ResetStyleStringReading
// Unimplemented: NextStyle
// Unimplemented: LastStyleName

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

// Unimplemented: ParameterInfo

/* -------------------------------------------------------------------- */
/*      Misc functions                                                  */
/* -------------------------------------------------------------------- */

// Unimplemented: OpenDSCount
// Unimplemented: OpenDS
// Unimplemented: CleanupAll

package ogr

/*
#include "ogr_core.h"
#include "ogr_api.h"

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
	Unknown					= GeometryType(C.wkbUnknown)
	Point					= GeometryType(C.wkbPoint)
	LineString				= GeometryType(C.wkbLineString)
	Polygon					= GeometryType(C.wkbPolygon)
	MultiPoint				= GeometryType(C.wkbMultiPoint)
	MultiLineString			= GeometryType(C.wkbMultiLineString)
	MultiPolygon			= GeometryType(C.wkbMultiPolygon)
	GeometryCollection		= GeometryType(C.wkbGeometryCollection)
	None					= GeometryType(C.wkbNone)
	LinearRing				= GeometryType(C.wkbLinearRing)
	Point25D				= GeometryType(C.wkbPoint25D)
	LineString25D			= GeometryType(C.wkbLineString25D)
	Polygon25D				= GeometryType(C.wkbPolygon25D)
	MultiPoint25D			= GeometryType(C.wkbMultiPoint25D)
	MultiLineString25D		= GeometryType(C.wkbMultiLineString25D)
	MultiPolygon25D			= GeometryType(C.wkbMultiPolygon25D)
	GeometryCollection25D	= GeometryType(C.wkbGeometryCollection25D)
)

// Error handling.  The following is bare-bones, and needs to be replaced with something more useful.
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

/* -------------------------------------------------------------------- */
/*      Geometry functions                                              */
/* -------------------------------------------------------------------- */

type Geometry struct {
	cval C.OGRGeometryH
}

//Create a geometry object from its well known binary representation
func CreateFromWkb(wkb []uint8, srs SpatialReference, bytes int) (Geometry, error) {
	cString := (*C.uchar)(unsafe.Pointer(&wkb[0]))
	var newGeom Geometry
	err := C.OGR_G_CreateFromWkb(cString, srs.cval, &newGeom.cval, C.int(bytes))
	return newGeom, error(err)
}

//Create a geometry object from its well known text representation
func CreateFromWkt(wkt string, srs SpatialReference) (Geometry, error) {
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
// Unimplemented: ForceToPolygon
// Unimplemented: ForceToMultiPolygon
// Unimplemented: ForceToMultiPoint
// Unimplemented: ForceToMultiLineString
// Unimplemented: GetDimension
// Unimplemented: GetCoordinateDimension
// Unimplemented: SetCoordinateDimension
// Unimplemented: Clone
// Unimplemented: GetEnvelope
// Unimplemented: GetEnvelope3D
// Unimplemented: ImportFromWkb
// Unimplemented: ExportToWkb
// Unimplemented: WkbSize
// Unimplemented: ImportFromWkt
// Unimplemented: ExportToWkt
// Unimplemented: Type
// Unimplemented: Name
// Unimplemented: DumpReadable
// Unimplemented: FlattenTo3D
// Unimplemented: CloseRings
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
/*      Spatial reference functions                                     */
/* -------------------------------------------------------------------- */

type SpatialReference struct {
	cval C.OGRSpatialReferenceH
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

type Driver struct {
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
/*      Misc functions                                                  */
/* -------------------------------------------------------------------- */

// Unimplemented: OpenDSCount
// Unimplemented: OpenDS
// Unimplemented: CleanupAll

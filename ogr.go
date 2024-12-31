package gdal

/*
#include "go_gdal.h"
#include "go_ogr_wkb.h"
#include "gdal_version.h"
*/
import "C"
import (
	"reflect"
	"time"
	"unsafe"
)

func init() {
	C.OGRRegisterAll()
}

/* -------------------------------------------------------------------- */
/*      Significant constants.                                          */
/* -------------------------------------------------------------------- */

// List of well known binary geometry types
type GeometryType uint32

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
/*      Envelope functions                                              */
/* -------------------------------------------------------------------- */

type Envelope struct {
	cval C.OGREnvelope
}

func (env Envelope) MinX() float64 {
	return float64(env.cval.MinX)
}

func (env Envelope) MaxX() float64 {
	return float64(env.cval.MaxX)
}

func (env Envelope) MinY() float64 {
	return float64(env.cval.MinY)
}

func (env Envelope) MaxY() float64 {
	return float64(env.cval.MaxY)
}

func (env *Envelope) SetMinX(val float64) {
	env.cval.MinX = C.double(val)
}

func (env *Envelope) SetMaxX(val float64) {
	env.cval.MaxX = C.double(val)
}

func (env *Envelope) SetMinY(val float64) {
	env.cval.MinY = C.double(val)
}

func (env *Envelope) SetMaxY(val float64) {
	env.cval.MaxY = C.double(val)
}

func (env Envelope) IsInit() bool {
	return env.cval.MinX != 0 || env.cval.MinY != 0 || env.cval.MaxX != 0 || env.cval.MaxY != 0
}

func min(a, b C.double) C.double {
	if a < b {
		return a
	}
	return b
}

func max(a, b C.double) C.double {
	if a > b {
		return a
	}
	return b
}

// Return the union of this envelope with another one
func (env Envelope) Union(other Envelope) Envelope {
	if env.IsInit() {
		env.cval.MinX = min(env.cval.MinX, other.cval.MinX)
		env.cval.MinY = min(env.cval.MinY, other.cval.MinY)
		env.cval.MaxX = max(env.cval.MaxX, other.cval.MaxX)
		env.cval.MaxY = max(env.cval.MaxY, other.cval.MaxY)
	} else {
		env.cval.MinX = other.cval.MinX
		env.cval.MinY = other.cval.MinY
		env.cval.MaxX = other.cval.MaxX
		env.cval.MaxY = other.cval.MaxY
	}
	return env
}

// Return the intersection of this envelope with another
func (env Envelope) Intersect(other Envelope) Envelope {
	if env.Intersects(other) {
		if env.IsInit() {
			env.cval.MinX = max(env.cval.MinX, other.cval.MinX)
			env.cval.MinY = max(env.cval.MinY, other.cval.MinY)
			env.cval.MaxX = min(env.cval.MaxX, other.cval.MaxX)
			env.cval.MaxY = min(env.cval.MaxY, other.cval.MaxY)
		} else {
			env.cval.MinX = other.cval.MinX
			env.cval.MinY = other.cval.MinY
			env.cval.MaxX = other.cval.MaxX
			env.cval.MaxY = other.cval.MaxY
		}
	} else {
		env.cval.MinX = 0
		env.cval.MinY = 0
		env.cval.MaxX = 0
		env.cval.MaxY = 0
	}
	return env
}

// Test if one envelope intersects another
func (env Envelope) Intersects(other Envelope) bool {
	return env.cval.MinX <= other.cval.MaxX &&
		env.cval.MaxX >= other.cval.MinX &&
		env.cval.MinY <= other.cval.MaxY &&
		env.cval.MaxY >= other.cval.MinY
}

// Test if one envelope completely contains another
func (env Envelope) Contains(other Envelope) bool {
	return env.cval.MinX <= other.cval.MinX &&
		env.cval.MaxX >= other.cval.MaxX &&
		env.cval.MinY <= other.cval.MinY &&
		env.cval.MaxY >= other.cval.MaxY
}

/* -------------------------------------------------------------------- */
/*      Envelope3D functions                                              */
/* -------------------------------------------------------------------- */

type Envelope3D struct {
	cval C.OGREnvelope3D
}

func (env Envelope3D) MinX() float64 {
	return float64(env.cval.MinX)
}

func (env Envelope3D) MaxX() float64 {
	return float64(env.cval.MaxX)
}

func (env Envelope3D) MinY() float64 {
	return float64(env.cval.MinY)
}

func (env Envelope3D) MaxY() float64 {
	return float64(env.cval.MaxY)
}

func (env Envelope3D) MinZ() float64 {
	return float64(env.cval.MinZ)
}

func (env Envelope3D) MaxZ() float64 {
	return float64(env.cval.MaxZ)
}

func (env *Envelope3D) SetMinX(val float64) {
	env.cval.MinX = C.double(val)
}

func (env *Envelope3D) SetMaxX(val float64) {
	env.cval.MaxX = C.double(val)
}

func (env *Envelope3D) SetMinY(val float64) {
	env.cval.MinY = C.double(val)
}

func (env *Envelope3D) SetMaxY(val float64) {
	env.cval.MaxY = C.double(val)
}

func (env *Envelope3D) SetMinZ(val float64) {
	env.cval.MinZ = C.double(val)
}

func (env *Envelope3D) SetMaxZ(val float64) {
	env.cval.MaxZ = C.double(val)
}

func (env Envelope3D) IsInit() bool {
	return env.cval.MinX != 0 || env.cval.MinY != 0 || env.cval.MinZ != 0 || env.cval.MaxX != 0 || env.cval.MaxY != 0 || env.cval.MaxZ != 0
}

// Return the union of this envelope3D with another one
func (env Envelope3D) Union(other Envelope3D) Envelope3D {
	if env.IsInit() {
		env.cval.MinX = min(env.cval.MinX, other.cval.MinX)
		env.cval.MinY = min(env.cval.MinY, other.cval.MinY)
		env.cval.MinZ = min(env.cval.MinZ, other.cval.MinZ)
		env.cval.MaxX = max(env.cval.MaxX, other.cval.MaxX)
		env.cval.MaxY = max(env.cval.MaxY, other.cval.MaxY)
		env.cval.MaxZ = max(env.cval.MaxZ, other.cval.MaxZ)
	} else {
		env.cval.MinX = other.cval.MinX
		env.cval.MinY = other.cval.MinY
		env.cval.MinZ = other.cval.MinY
		env.cval.MaxX = other.cval.MaxX
		env.cval.MaxY = other.cval.MaxY
		env.cval.MaxZ = other.cval.MaxZ
	}
	return env
}

// Return the intersection of this envelope3D with another
func (env Envelope3D) Intersect(other Envelope3D) Envelope3D {
	if env.Intersects(other) {
		if env.IsInit() {
			env.cval.MinX = max(env.cval.MinX, other.cval.MinX)
			env.cval.MinY = max(env.cval.MinY, other.cval.MinY)
			env.cval.MinZ = max(env.cval.MinZ, other.cval.MinZ)
			env.cval.MaxX = min(env.cval.MaxX, other.cval.MaxX)
			env.cval.MaxY = min(env.cval.MaxY, other.cval.MaxY)
			env.cval.MaxZ = min(env.cval.MaxZ, other.cval.MaxZ)
		} else {
			env.cval.MinX = other.cval.MinX
			env.cval.MinY = other.cval.MinY
			env.cval.MinZ = other.cval.MinZ
			env.cval.MaxX = other.cval.MaxX
			env.cval.MaxY = other.cval.MaxY
			env.cval.MaxZ = other.cval.MaxZ
		}
	} else {
		env.cval.MinX = 0
		env.cval.MinY = 0
		env.cval.MinZ = 0
		env.cval.MaxX = 0
		env.cval.MaxY = 0
		env.cval.MaxZ = 0
	}
	return env
}

// Test if one envelope3D intersects another
func (env Envelope3D) Intersects(other Envelope3D) bool {
	return env.cval.MinX <= other.cval.MaxX &&
		env.cval.MaxX >= other.cval.MinX &&
		env.cval.MinY <= other.cval.MaxY &&
		env.cval.MaxY >= other.cval.MinY &&
		env.cval.MinZ <= other.cval.MaxZ &&
		env.cval.MaxZ >= other.cval.MinZ

}

// Test if one envelope3D completely contains another
func (env Envelope3D) Contains(other Envelope3D) bool {
	return env.cval.MinX <= other.cval.MinX &&
		env.cval.MaxX >= other.cval.MaxX &&
		env.cval.MinY <= other.cval.MinY &&
		env.cval.MaxY >= other.cval.MaxY &&
		env.cval.MinZ <= other.cval.MinZ &&
		env.cval.MaxZ >= other.cval.MaxZ

}

/* -------------------------------------------------------------------- */
/*      Misc functions                                                  */
/* -------------------------------------------------------------------- */

// Clean up all OGR related resources
func CleanupOGR() {
	C.OGRCleanupAll()
}

// Convert a go bool to a C int
func BoolToCInt(in bool) (out C.int) {
	if in {
		out = 1
	} else {
		out = 0
	}
	return
}

/* -------------------------------------------------------------------- */
/*      Geometry functions                                              */
/* -------------------------------------------------------------------- */

type Geometry struct {
	cval C.OGRGeometryH
}

//Create a geometry object from its well known binary representation
func CreateFromWKB(wkb []uint8, srs SpatialReference, bytes int) (Geometry, error) {
	cString := unsafe.Pointer(&wkb[0])
	var newGeom Geometry
	return newGeom, C.go_CreateFromWkb(
		cString, srs.cval, &newGeom.cval, C.int(bytes),
	).Err()
}

//Create a geometry object from its well known text representation
func CreateFromWKT(wkt string, srs SpatialReference) (Geometry, error) {
	cString := C.CString(wkt)
	defer C.free(unsafe.Pointer(cString))
	var newGeom Geometry
	return newGeom, C.OGR_G_CreateFromWkt(
		&cString, srs.cval, &newGeom.cval,
	).Err()
}

//Create a geometry object from its GeoJSON representation
func CreateFromJson(_json string) Geometry {
	cString := C.CString(_json)
	defer C.free(unsafe.Pointer(cString))
	var newGeom Geometry
	newGeom.cval = C.OGR_G_CreateGeometryFromJson(cString)
	return newGeom
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

// Stroke arc to linestring
func ApproximateArcAngles(
	x, y, z,
	primaryRadius,
	secondaryRadius,
	rotation,
	startAngle,
	endAngle,
	stepSizeDegrees float64,
) Geometry {
	geom := C.OGR_G_ApproximateArcAngles(
		C.double(x),
		C.double(y),
		C.double(z),
		C.double(primaryRadius),
		C.double(secondaryRadius),
		C.double(rotation),
		C.double(startAngle),
		C.double(endAngle),
		C.double(stepSizeDegrees))
	return Geometry{geom}
}

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

// Compute and return the bounding envelope for this geometry
func (geom Geometry) Envelope() Envelope {
	var env Envelope
	C.OGR_G_GetEnvelope(geom.cval, &env.cval)
	return env
}

// Compute and return the 3D bounding envelope for this geometry
func (geom Geometry) Envelope3D() Envelope3D {
	var env Envelope3D
	C.OGR_G_GetEnvelope3D(geom.cval, &env.cval)
	return env
}

// Assign a geometry from well known binary data
func (geom Geometry) FromWKB(wkb []uint8, bytes int) error {
	cString := unsafe.Pointer(&wkb[0])
	return C.go_ImportFromWkb(geom.cval, cString, C.int(bytes)).Err()
}

// Convert a geometry to well known binary data
func (geom Geometry) ToWKB() ([]uint8, error) {
	b := make([]uint8, geom.WKBSize())
	cString := (*C.uchar)(unsafe.Pointer(&b[0]))
	err := C.go_ExportToWkb(geom.cval, C.OGRwkbByteOrder(C.wkbNDR), cString).Err()
	return b, err
}

// Returns size of related binary representation
func (geom Geometry) WKBSize() int {
	size := C.OGR_G_WkbSize(geom.cval)
	return int(size)
}

// Assign geometry object from its well known text representation
func (geom Geometry) FromWKT(wkt string) error {
	cString := C.CString(wkt)
	defer C.free(unsafe.Pointer(cString))
	return C.OGR_G_ImportFromWkt(geom.cval, &cString).Err()
}

// Fetch geometry as WKT
func (geom Geometry) ToWKT() (string, error) {
	var p *C.char
	err := C.OGR_G_ExportToWkt(geom.cval, &p).Err()
	wkt := C.GoString(p)
	defer C.free(unsafe.Pointer(p))
	return wkt, err
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

// Create a geometry from its GML representation
func CreateFromGML(gml string) Geometry {
	cString := C.CString(gml)
	defer C.free(unsafe.Pointer(cString))
	geom := C.OGR_G_CreateFromGML(cString)
	return Geometry{geom}
}

// Convert a geometry to GML format
func (geom Geometry) ToGML() string {
	val := C.OGR_G_ExportToGML(geom.cval)
	return C.GoString(val)
}

// Convert a geometry to GML format with options
func (geom Geometry) ToGML_Ex(options []string) string {
	length := len(options)
	opts := make([]*C.char, length+1)
	for i := 0; i < length; i++ {
		opts[i] = C.CString(options[i])
		defer C.free(unsafe.Pointer(opts[i]))
	}
	opts[length] = (*C.char)(unsafe.Pointer(nil))

	val := C.OGR_G_ExportToGMLEx(geom.cval, (**C.char)(unsafe.Pointer(&opts[0])))
	return C.GoString(val)
}

// Convert a geometry to KML format
func (geom Geometry) ToKML() string {
	val := C.OGR_G_ExportToKML(geom.cval, nil)
	result := C.GoString(val)
	C.free(unsafe.Pointer(val))
	return result
}

// Convert a geometry to JSON format
func (geom Geometry) ToJSON() string {
	val := C.OGR_G_ExportToJson(geom.cval)
	return C.GoString(val)
}

// Convert a geometry to JSON format with options
func (geom Geometry) ToJSON_ex(options []string) string {
	length := len(options)
	opts := make([]*C.char, length+1)
	for i := 0; i < length; i++ {
		opts[i] = C.CString(options[i])
		defer C.free(unsafe.Pointer(opts[i]))
	}
	opts[length] = (*C.char)(unsafe.Pointer(nil))

	val := C.OGR_G_ExportToJsonEx(geom.cval, (**C.char)(unsafe.Pointer(&opts[0])))
	return C.GoString(val)
}

// Fetch the spatial reference associated with this geometry
func (geom Geometry) SpatialReference() SpatialReference {
	spatialRef := C.OGR_G_GetSpatialReference(geom.cval)
	return SpatialReference{spatialRef}
}

// Assign a spatial reference to this geometry
func (geom Geometry) SetSpatialReference(spatialRef SpatialReference) {
	C.OGR_G_AssignSpatialReference(geom.cval, spatialRef.cval)
}

// Apply coordinate transformation to geometry
func (geom Geometry) Transform(ct CoordinateTransform) error {
	return C.OGR_G_Transform(geom.cval, ct.cval).Err()
}

// Transform geometry to new spatial reference system
func (geom Geometry) TransformTo(sr SpatialReference) error {
	return C.OGR_G_TransformTo(geom.cval, sr.cval).Err()
}

// Simplify the geometry
func (geom Geometry) Simplify(tolerance float64) Geometry {
	newGeom := C.OGR_G_Simplify(geom.cval, C.double(tolerance))
	return Geometry{newGeom}
}

// Simplify the geometry while preserving topology
func (geom Geometry) SimplifyPreservingTopology(tolerance float64) Geometry {
	newGeom := C.OGR_G_SimplifyPreserveTopology(geom.cval, C.double(tolerance))
	return Geometry{newGeom}
}

// Modify the geometry such that it has no line segment longer than the given distance
func (geom Geometry) Segmentize(distance float64) {
	C.OGR_G_Segmentize(geom.cval, C.double(distance))
}

// Return true if these features intersect
func (geom Geometry) Intersects(other Geometry) bool {
	val := C.OGR_G_Intersects(geom.cval, other.cval)
	return val != 0
}

// Return true if these features are equal
func (geom Geometry) Equals(other Geometry) bool {
	val := C.OGR_G_Equals(geom.cval, other.cval)
	return val != 0
}

// Return true if the features are disjoint
func (geom Geometry) Disjoint(other Geometry) bool {
	val := C.OGR_G_Disjoint(geom.cval, other.cval)
	return val != 0
}

// Return true if this feature touches the other
func (geom Geometry) Touches(other Geometry) bool {
	val := C.OGR_G_Touches(geom.cval, other.cval)
	return val != 0
}

// Return true if this feature crosses the other
func (geom Geometry) Crosses(other Geometry) bool {
	val := C.OGR_G_Crosses(geom.cval, other.cval)
	return val != 0
}

// Return true if this geometry is within the other
func (geom Geometry) Within(other Geometry) bool {
	val := C.OGR_G_Within(geom.cval, other.cval)
	return val != 0
}

// Return true if this geometry contains the other
func (geom Geometry) Contains(other Geometry) bool {
	val := C.OGR_G_Contains(geom.cval, other.cval)
	return val != 0
}

// Return true if this geometry overlaps the other
func (geom Geometry) Overlaps(other Geometry) bool {
	val := C.OGR_G_Overlaps(geom.cval, other.cval)
	return val != 0
}

// Compute boundary for the geometry
func (geom Geometry) Boundary() Geometry {
	newGeom := C.OGR_G_Boundary(geom.cval)
	return Geometry{newGeom}
}

// Compute convex hull for the geometry
func (geom Geometry) ConvexHull() Geometry {
	newGeom := C.OGR_G_ConvexHull(geom.cval)
	return Geometry{newGeom}
}

// Compute buffer of the geometry
func (geom Geometry) Buffer(distance float64, segments int) Geometry {
	newGeom := C.OGR_G_Buffer(geom.cval, C.double(distance), C.int(segments))
	return Geometry{newGeom}
}

// Compute intersection of this geometry with the other
func (geom Geometry) Intersection(other Geometry) Geometry {
	newGeom := C.OGR_G_Intersection(geom.cval, other.cval)
	return Geometry{newGeom}
}

// Compute union of this geometry with the other
func (geom Geometry) Union(other Geometry) Geometry {
	newGeom := C.OGR_G_Union(geom.cval, other.cval)
	return Geometry{newGeom}
}

func (geom Geometry) UnionCascaded() Geometry {
	newGeom := C.OGR_G_UnionCascaded(geom.cval)
	return Geometry{newGeom}
}

// Unimplemented: PointOn Surface (until 2.0)
// Return a point guaranteed to lie on the surface
// func (geom Geometry) PointOnSurface() Geometry {
//  newGeom := C.OGR_G_PointOnSurface(geom.cval)
//  return Geometry{newGeom}
// }

// Compute difference between this geometry and the other
func (geom Geometry) Difference(other Geometry) Geometry {
	newGeom := C.OGR_G_Difference(geom.cval, other.cval)
	return Geometry{newGeom}
}

// Compute symmetric difference between this geometry and the other
func (geom Geometry) SymmetricDifference(other Geometry) Geometry {
	newGeom := C.OGR_G_SymDifference(geom.cval, other.cval)
	return Geometry{newGeom}
}

// Compute distance between thie geometry and the other
func (geom Geometry) Distance(other Geometry) float64 {
	dist := C.OGR_G_Distance(geom.cval, other.cval)
	return float64(dist)
}

// Compute 3D distance between thie geometry and the other.
// This method is built on the SFCGAL library,
// check it for the definition of the geometry operation.
// If OGR is built without the SFCGAL library, this method will always return -1.0
func (geom Geometry) Distance3D(other Geometry) float64 {
	dist := C.OGR_G_Distance3D(geom.cval, other.cval)
	return float64(dist)
}

// Compute length of geometry
func (geom Geometry) Length() float64 {
	length := C.OGR_G_Length(geom.cval)
	return float64(length)
}

// Compute area of geometry
func (geom Geometry) Area() float64 {
	area := C.OGR_G_Area(geom.cval)
	return float64(area)
}

// Compute centroid of geometry
func (geom Geometry) Centroid() Geometry {
	centroid := Geometry{C.OGR_G_CreateGeometry(C.wkbPoint)}
	C.OGR_G_Centroid(geom.cval, centroid.cval)
	return centroid
}

// Clear the geometry to its uninitialized state
func (geom Geometry) Empty() {
	C.OGR_G_Empty(geom.cval)
}

// Test if the geometry is empty
func (geom Geometry) IsEmpty() bool {
	val := C.OGR_G_IsEmpty(geom.cval)
	return val != 0
}

// Test if the geometry is null
func (geom Geometry) IsNull() bool {
	return geom.cval == nil
}

// Test if the geometry is valid
func (geom Geometry) IsValid() bool {
	val := C.OGR_G_IsValid(geom.cval)
	return val != 0
}

// Test if the geometry is simple
func (geom Geometry) IsSimple() bool {
	val := C.OGR_G_IsSimple(geom.cval)
	return val != 0
}

// Test if the geometry is a ring
func (geom Geometry) IsRing() bool {
	val := C.OGR_G_IsRing(geom.cval)
	return val != 0
}

// Polygonize a set of sparse edges
func (geom Geometry) Polygonize() Geometry {
	newGeom := C.OGR_G_Polygonize(geom.cval)
	return Geometry{newGeom}
}

// Fetch number of points in the geometry
func (geom Geometry) PointCount() int {
	count := C.OGR_G_GetPointCount(geom.cval)
	return int(count)
}

// Unimplemented: Points

// Fetch the X coordinate of a point in the geometry
func (geom Geometry) X(index int) float64 {
	x := C.OGR_G_GetX(geom.cval, C.int(index))
	return float64(x)
}

// Fetch the Y coordinate of a point in the geometry
func (geom Geometry) Y(index int) float64 {
	y := C.OGR_G_GetY(geom.cval, C.int(index))
	return float64(y)
}

// Fetch the Z coordinate of a point in the geometry
func (geom Geometry) Z(index int) float64 {
	z := C.OGR_G_GetZ(geom.cval, C.int(index))
	return float64(z)
}

// Fetch the coordinates of a point in the geometry
func (geom Geometry) Point(index int) (x, y, z float64) {
	C.OGR_G_GetPoint(
		geom.cval,
		C.int(index),
		(*C.double)(&x),
		(*C.double)(&y),
		(*C.double)(&z))
	return
}

// Set the coordinates of a point in the geometry
func (geom Geometry) SetPoint(index int, x, y, z float64) {
	C.OGR_G_SetPoint(
		geom.cval,
		C.int(index),
		C.double(x),
		C.double(y),
		C.double(z))
}

// Set the coordinates of a point in the geometry, ignoring the 3rd dimension
func (geom Geometry) SetPoint2D(index int, x, y float64) {
	C.OGR_G_SetPoint_2D(geom.cval, C.int(index), C.double(x), C.double(y))
}

// Add a new point to the geometry (line string or polygon only)
func (geom Geometry) AddPoint(x, y, z float64) {
	C.OGR_G_AddPoint(geom.cval, C.double(x), C.double(y), C.double(z))
}

// Add a new point to the geometry (line string or polygon only), ignoring the 3rd dimension
func (geom Geometry) AddPoint2D(x, y float64) {
	C.OGR_G_AddPoint_2D(geom.cval, C.double(x), C.double(y))
}

// Fetch the number of elements in the geometry, or number of geometries in the container
func (geom Geometry) GeometryCount() int {
	count := C.OGR_G_GetGeometryCount(geom.cval)
	return int(count)
}

// Fetch geometry from a geometry container
func (geom Geometry) Geometry(index int) Geometry {
	newGeom := C.OGR_G_GetGeometryRef(geom.cval, C.int(index))
	return Geometry{newGeom}
}

// Add a geometry to a geometry container
func (geom Geometry) AddGeometry(other Geometry) error {
	return C.OGR_G_AddGeometry(geom.cval, other.cval).Err()
}

// Add a geometry to a geometry container and assign ownership to that container
func (geom Geometry) AddGeometryDirectly(other Geometry) error {
	return C.OGR_G_AddGeometryDirectly(geom.cval, other.cval).Err()
}

// Remove a geometry from the geometry container
func (geom Geometry) RemoveGeometry(index int, delete bool) error {
	return C.OGR_G_RemoveGeometry(geom.cval, C.int(index), BoolToCInt(delete)).Err()
}

// Build a polygon / ring from a set of lines
func (geom Geometry) BuildPolygonFromEdges(autoClose bool, tolerance float64) (Geometry, error) {
	var cErr C.OGRErr
	newGeom := C.OGRBuildPolygonFromEdges(
		geom.cval,
		0,
		BoolToCInt(autoClose),
		C.double(tolerance),
		&cErr,
	)
	return Geometry{newGeom}, cErr.Err()
}

/* -------------------------------------------------------------------- */
/*      Field definition functions                                      */
/* -------------------------------------------------------------------- */

// List of well known binary geometry types
type FieldType int

const (
	FT_Integer       = FieldType(C.OFTInteger)
	FT_IntegerList   = FieldType(C.OFTIntegerList)
	FT_Real          = FieldType(C.OFTReal)
	FT_RealList      = FieldType(C.OFTRealList)
	FT_String        = FieldType(C.OFTString)
	FT_StringList    = FieldType(C.OFTStringList)
	FT_Binary        = FieldType(C.OFTBinary)
	FT_Date          = FieldType(C.OFTDate)
	FT_Time          = FieldType(C.OFTTime)
	FT_DateTime      = FieldType(C.OFTDateTime)
	FT_Integer64     = FieldType(C.OFTInteger64)
	FT_Integer64List = FieldType(C.OFTInteger64List)
)

type Justification int

const (
	J_Undefined = Justification(C.OJUndefined)
	J_Left      = Justification(C.OJLeft)
	J_Right     = Justification(C.OJRight)
)

type FieldDefinition struct {
	cval C.OGRFieldDefnH
}

type Field struct {
	cval *C.OGRField
}

// Create a new field definition
func CreateFieldDefinition(name string, fieldType FieldType) FieldDefinition {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	fieldDef := C.OGR_Fld_Create(cName, C.OGRFieldType(fieldType))
	return FieldDefinition{fieldDef}
}

// Destroy the field definition
func (fd FieldDefinition) Destroy() {
	C.OGR_Fld_Destroy(fd.cval)
}

// Fetch the name of the field
func (fd FieldDefinition) Name() string {
	name := C.OGR_Fld_GetNameRef(fd.cval)
	return C.GoString(name)
}

// Set the name of the field
func (fd FieldDefinition) SetName(name string) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	C.OGR_Fld_SetName(fd.cval, cName)
}

// Fetch the type of this field
func (fd FieldDefinition) Type() FieldType {
	fType := C.OGR_Fld_GetType(fd.cval)
	return FieldType(fType)
}

// Set the type of this field
func (fd FieldDefinition) SetType(fType FieldType) {
	C.OGR_Fld_SetType(fd.cval, C.OGRFieldType(fType))
}

// Fetch the justification for this field
func (fd FieldDefinition) Justification() Justification {
	justify := C.OGR_Fld_GetJustify(fd.cval)
	return Justification(justify)
}

// Set the justification for this field
func (fd FieldDefinition) SetJustification(justify Justification) {
	C.OGR_Fld_SetJustify(fd.cval, C.OGRJustification(justify))
}

// Fetch the formatting width for this field
func (fd FieldDefinition) Width() int {
	width := C.OGR_Fld_GetWidth(fd.cval)
	return int(width)
}

// Set the formatting width for this field
func (fd FieldDefinition) SetWidth(width int) {
	C.OGR_Fld_SetWidth(fd.cval, C.int(width))
}

// Fetch the precision for this field
func (fd FieldDefinition) Precision() int {
	precision := C.OGR_Fld_GetPrecision(fd.cval)
	return int(precision)
}

// Set the precision for this field
func (fd FieldDefinition) SetPrecision(precision int) {
	C.OGR_Fld_SetPrecision(fd.cval, C.int(precision))
}

// Set defining parameters of field in a single call
func (fd FieldDefinition) Set(
	name string,
	fType FieldType,
	width, precision int,
	justify Justification,
) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	C.OGR_Fld_Set(
		fd.cval,
		cName,
		C.OGRFieldType(fType),
		C.int(width),
		C.int(precision),
		C.OGRJustification(justify),
	)
}

// Fetch whether this field should be ignored when fetching features
func (fd FieldDefinition) IsIgnored() bool {
	ignore := C.OGR_Fld_IsIgnored(fd.cval)
	return ignore != 0
}

// Set whether this field should be ignored when fetching features
func (fd FieldDefinition) SetIgnored(ignore bool) {
	C.OGR_Fld_SetIgnored(fd.cval, BoolToCInt(ignore))
}

// Fetch human readable name for the field type
func (ft FieldType) Name() string {
	name := C.OGR_GetFieldTypeName(C.OGRFieldType(ft))
	return C.GoString(name)
}

/* -------------------------------------------------------------------- */
/*      Feature definition functions                                    */
/* -------------------------------------------------------------------- */

type FeatureDefinition struct {
	cval C.OGRFeatureDefnH
}

// Create a new feature definition object
func CreateFeatureDefinition(name string) FeatureDefinition {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	fd := C.OGR_FD_Create(cName)
	return FeatureDefinition{fd}
}

// Destroy a feature definition object
func (fd FeatureDefinition) Destroy() {
	C.OGR_FD_Destroy(fd.cval)
}

// Drop a reference, and delete object if no references remain
func (fd FeatureDefinition) Release() {
	C.OGR_FD_Release(fd.cval)
}

// Fetch the name of this feature definition
func (fd FeatureDefinition) Name() string {
	name := C.OGR_FD_GetName(fd.cval)
	return C.GoString(name)
}

// Fetch the number of fields in the feature definition
func (fd FeatureDefinition) FieldCount() int {
	count := C.OGR_FD_GetFieldCount(fd.cval)
	return int(count)
}

// Fetch the definition of the indicated field
func (fd FeatureDefinition) FieldDefinition(index int) FieldDefinition {
	fieldDefn := C.OGR_FD_GetFieldDefn(fd.cval, C.int(index))
	return FieldDefinition{fieldDefn}
}

// Fetch the index of the named field
func (fd FeatureDefinition) FieldIndex(name string) int {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	index := C.OGR_FD_GetFieldIndex(fd.cval, cName)
	return int(index)
}

// Add a new field definition to this feature definition
func (fd FeatureDefinition) AddFieldDefinition(fieldDefn FieldDefinition) {
	C.OGR_FD_AddFieldDefn(fd.cval, fieldDefn.cval)
}

// Delete a field definition from this feature definition
func (fd FeatureDefinition) DeleteFieldDefinition(index int) error {
	return C.OGR_FD_DeleteFieldDefn(fd.cval, C.int(index)).Err()
}

// Fetch the geometry base type of this feature definition
func (fd FeatureDefinition) GeometryType() GeometryType {
	gt := C.OGR_FD_GetGeomType(fd.cval)
	return GeometryType(gt)
}

// Set the geometry base type for this feature definition
func (fd FeatureDefinition) SetGeometryType(geomType GeometryType) {
	C.OGR_FD_SetGeomType(fd.cval, C.OGRwkbGeometryType(geomType))
}

// Fetch if the geometry can be ignored when fetching features
func (fd FeatureDefinition) IsGeometryIgnored() bool {
	isIgnored := C.OGR_FD_IsGeometryIgnored(fd.cval)
	return isIgnored != 0
}

// Set whether the geometry can be ignored when fetching features
func (fd FeatureDefinition) SetGeometryIgnored(val bool) {
	C.OGR_FD_SetGeometryIgnored(fd.cval, BoolToCInt(val))
}

// Fetch if the style can be ignored when fetching features
func (fd FeatureDefinition) IsStyleIgnored() bool {
	isIgnored := C.OGR_FD_IsStyleIgnored(fd.cval)
	return isIgnored != 0
}

// Set whether the style can be ignored when fetching features
func (fd FeatureDefinition) SetStyleIgnored(val bool) {
	C.OGR_FD_SetStyleIgnored(fd.cval, BoolToCInt(val))
}

// Increment the reference count by one
func (fd FeatureDefinition) Reference() int {
	count := C.OGR_FD_Reference(fd.cval)
	return int(count)
}

// Decrement the reference count by one
func (fd FeatureDefinition) Dereference() int {
	count := C.OGR_FD_Dereference(fd.cval)
	return int(count)
}

// Fetch the current reference count
func (fd FeatureDefinition) ReferenceCount() int {
	count := C.OGR_FD_GetReferenceCount(fd.cval)
	return int(count)
}

/* -------------------------------------------------------------------- */
/*      Feature functions                                               */
/* -------------------------------------------------------------------- */

type Feature struct {
	cval C.OGRFeatureH
}

// Create a feature from this feature definition
func (fd FeatureDefinition) Create() Feature {
	feature := C.OGR_F_Create(fd.cval)
	return Feature{feature}
}

// Destroy this feature
func (feature Feature) Destroy() {
	C.OGR_F_Destroy(feature.cval)
}

// Fetch feature definition
func (feature Feature) Definition() FeatureDefinition {
	fd := C.OGR_F_GetDefnRef(feature.cval)
	return FeatureDefinition{fd}
}

// Set feature geometry
func (feature Feature) SetGeometry(geom Geometry) error {
	return C.OGR_F_SetGeometry(feature.cval, geom.cval).Err()
}

// Set feature geometry, passing ownership to the feature
func (feature Feature) SetGeometryDirectly(geom Geometry) error {
	return C.OGR_F_SetGeometryDirectly(feature.cval, geom.cval).Err()
}

// Fetch geometry of this feature
func (feature Feature) Geometry() Geometry {
	geom := C.OGR_F_GetGeometryRef(feature.cval)
	return Geometry{geom}
}

// Fetch geometry of this feature and assume ownership
func (feature Feature) StealGeometry() Geometry {
	geom := C.OGR_F_StealGeometry(feature.cval)
	return Geometry{geom}
}

// Duplicate feature
func (feature Feature) Clone() Feature {
	newFeature := C.OGR_F_Clone(feature.cval)
	return Feature{newFeature}
}

// Test if two features are the same
func (f1 Feature) Equal(f2 Feature) bool {
	equal := C.OGR_F_Equal(f1.cval, f2.cval)
	return equal != 0
}

// Fetch number of fields on this feature
func (feature Feature) FieldCount() int {
	count := C.OGR_F_GetFieldCount(feature.cval)
	return int(count)
}

// Fetch definition for the indicated field
func (feature Feature) FieldDefinition(index int) FieldDefinition {
	defn := C.OGR_F_GetFieldDefnRef(feature.cval, C.int(index))
	return FieldDefinition{defn}
}

// Fetch the field index for the given field name
func (feature Feature) FieldIndex(name string) int {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	index := C.OGR_F_GetFieldIndex(feature.cval, cName)
	return int(index)
}

// Return if a field has ever been assigned a value
func (feature Feature) IsFieldSet(index int) bool {
	set := C.OGR_F_IsFieldSet(feature.cval, C.int(index))
	return set != 0
}

// Clear a field and mark it as unset
func (feature Feature) UnnsetField(index int) {
	C.OGR_F_UnsetField(feature.cval, C.int(index))
}

// Fetch a reference to the internal field value
func (feature Feature) RawField(index int) Field {
	field := C.OGR_F_GetRawFieldRef(feature.cval, C.int(index))
	return Field{field}
}

// Fetch field value as integer
func (feature Feature) FieldAsInteger(index int) int {
	val := C.OGR_F_GetFieldAsInteger(feature.cval, C.int(index))
	return int(val)
}

// Fetch field value as 64-bit integer
func (feature Feature) FieldAsInteger64(index int) int64 {
	val := C.OGR_F_GetFieldAsInteger64(feature.cval, C.int(index))
	return int64(val)
}

// Fetch field value as float64
func (feature Feature) FieldAsFloat64(index int) float64 {
	val := C.OGR_F_GetFieldAsDouble(feature.cval, C.int(index))
	return float64(val)
}

// Fetch field value as string
func (feature Feature) FieldAsString(index int) string {
	val := C.OGR_F_GetFieldAsString(feature.cval, C.int(index))
	return C.GoString(val)
}

// Fetch field as list of integers
func (feature Feature) FieldAsIntegerList(index int) []int {
	var count int
	cArray := C.OGR_F_GetFieldAsIntegerList(feature.cval, C.int(index), (*C.int)(unsafe.Pointer(&count)))
	var goSlice []int
	header := (*reflect.SliceHeader)(unsafe.Pointer(&goSlice))
	header.Cap = count
	header.Len = count
	header.Data = uintptr(unsafe.Pointer(cArray))
	return goSlice
}

// Fetch field as list of 64-bit integers
func (feature Feature) FieldAsInteger64List(index int) []int64 {
	var count int
	cArray := C.OGR_F_GetFieldAsInteger64List(feature.cval, C.int(index), (*C.int)(unsafe.Pointer(&count)))
	var goSlice []int64
	header := (*reflect.SliceHeader)(unsafe.Pointer(&goSlice))
	header.Cap = count
	header.Len = count
	header.Data = uintptr(unsafe.Pointer(cArray))
	return goSlice
}

// Fetch field as list of float64
func (feature Feature) FieldAsFloat64List(index int) []float64 {
	var count int
	cArray := C.OGR_F_GetFieldAsDoubleList(feature.cval, C.int(index), (*C.int)(unsafe.Pointer(&count)))
	var goSlice []float64
	header := (*reflect.SliceHeader)(unsafe.Pointer(&goSlice))
	header.Cap = count
	header.Len = count
	header.Data = uintptr(unsafe.Pointer(cArray))
	return goSlice
}

// Fetch field as list of strings
func (feature Feature) FieldAsStringList(index int) []string {
	p := C.OGR_F_GetFieldAsStringList(feature.cval, C.int(index))

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

// Fetch field as binary data
func (feature Feature) FieldAsBinary(index int) []uint8 {
	var count int
	cArray := C.OGR_F_GetFieldAsBinary(feature.cval, C.int(index), (*C.int)(unsafe.Pointer(&count)))
	var goSlice []uint8
	header := (*reflect.SliceHeader)(unsafe.Pointer(&goSlice))
	header.Cap = count
	header.Len = count
	header.Data = uintptr(unsafe.Pointer(cArray))
	return goSlice
}

// Fetch field as date and time
func (feature Feature) FieldAsDateTime(index int) (time.Time, bool) {
	var year, month, day, hour, minute, second, tzFlag int
	success := C.OGR_F_GetFieldAsDateTime(
		feature.cval,
		C.int(index),
		(*C.int)(unsafe.Pointer(&year)),
		(*C.int)(unsafe.Pointer(&month)),
		(*C.int)(unsafe.Pointer(&day)),
		(*C.int)(unsafe.Pointer(&hour)),
		(*C.int)(unsafe.Pointer(&minute)),
		(*C.int)(unsafe.Pointer(&second)),
		(*C.int)(unsafe.Pointer(&tzFlag)),
	)
	t := time.Date(year, time.Month(month), day, hour, minute, second, 0, time.UTC)
	return t, success != 0
}

// Set field to integer value
func (feature Feature) SetFieldInteger(index, value int) {
	C.OGR_F_SetFieldInteger(feature.cval, C.int(index), C.int(value))
}

// Set field to 64-bit integer value
func (feature Feature) SetFieldInteger64(index int, value int64) {
	C.OGR_F_SetFieldInteger64(feature.cval, C.int(index), C.GIntBig(value))
}

// Set field to float64 value
func (feature Feature) SetFieldFloat64(index int, value float64) {
	C.OGR_F_SetFieldDouble(feature.cval, C.int(index), C.double(value))
}

// Set field to string value
func (feature Feature) SetFieldString(index int, value string) {
	cVal := C.CString(value)
	defer C.free(unsafe.Pointer(cVal))
	C.OGR_F_SetFieldString(feature.cval, C.int(index), cVal)
}

// Set field to list of integers
func (feature Feature) SetFieldIntegerList(index int, value []int) {
	C.OGR_F_SetFieldIntegerList(
		feature.cval,
		C.int(index),
		C.int(len(value)),
		(*C.int)(unsafe.Pointer(&value[0])),
	)
}

// Set field to list of 64-bit integers
func (feature Feature) SetFieldInteger64List(index int, value []int64) {
	C.OGR_F_SetFieldIntegerList(
		feature.cval,
		C.int(index),
		C.int(len(value)),
		(*C.int)(unsafe.Pointer(&value[0])),
	)
}

// Set field to list of float64
func (feature Feature) SetFieldFloat64List(index int, value []float64) {
	C.OGR_F_SetFieldDoubleList(
		feature.cval,
		C.int(index),
		C.int(len(value)),
		(*C.double)(unsafe.Pointer(&value[0])),
	)
}

// Set field to list of strings
func (feature Feature) SetFieldStringList(index int, value []string) {
	length := len(value)
	cValue := make([]*C.char, length+1)
	for i := 0; i < length; i++ {
		cValue[i] = C.CString(value[i])
		defer C.free(unsafe.Pointer(cValue[i]))
	}
	cValue[length] = (*C.char)(unsafe.Pointer(nil))

	C.OGR_F_SetFieldStringList(
		feature.cval,
		C.int(index),
		(**C.char)(unsafe.Pointer(&cValue[0])),
	)
}

// Set field from the raw field pointer
func (feature Feature) SetFieldRaw(index int, field Field) {
	C.OGR_F_SetFieldRaw(feature.cval, C.int(index), field.cval)
}

// Set field as binary data
func (feature Feature) SetFieldBinary(index int, value []uint8) {
	C.OGR_F_SetFieldBinary(
		feature.cval,
		C.int(index),
		C.int(len(value)),
		unsafe.Pointer(&value[0]),
	)
}

// Set field as date / time
func (feature Feature) SetFieldDateTime(index int, dt time.Time) {
	C.OGR_F_SetFieldDateTime(
		feature.cval,
		C.int(index),
		C.int(dt.Year()),
		C.int(dt.Month()),
		C.int(dt.Day()),
		C.int(dt.Hour()),
		C.int(dt.Minute()),
		C.int(dt.Second()),
		C.int(1),
	)
}

// Fetch feature indentifier
func (feature Feature) FID() int64 {
	fid := C.OGR_F_GetFID(feature.cval)
	return int64(fid)
}

// Set feature identifier
func (feature Feature) SetFID(fid int64) error {
	return C.OGR_F_SetFID(feature.cval, C.GIntBig(fid)).Err()
}

// Unimplemented: DumpReadable

// Set one feature from another
func (this Feature) SetFrom(other Feature, forgiving int) error {
	return C.OGR_F_SetFrom(this.cval, other.cval, C.int(forgiving)).Err()
}

// Set one feature from another, using field map
func (this Feature) SetFromWithMap(other Feature, forgiving int, fieldMap []int) error {
	return C.OGR_F_SetFromWithMap(
		this.cval,
		other.cval,
		C.int(forgiving),
		(*C.int)(unsafe.Pointer(&fieldMap[0])),
	).Err()
}

// Fetch style string for this feature
func (feature Feature) StlyeString() string {
	style := C.OGR_F_GetStyleString(feature.cval)
	return C.GoString(style)
}

// Set style string for this feature
func (feature Feature) SetStyleString(style string) {
	cStyle := C.CString(style)
	C.OGR_F_SetStyleStringDirectly(feature.cval, cStyle)
}

// Returns true if this contains a null pointer
func (feature Feature) IsNull() bool {
	return feature.cval == nil
}

/* -------------------------------------------------------------------- */
/*      Layer functions                                                 */
/* -------------------------------------------------------------------- */

type Layer struct {
	cval C.OGRLayerH
}

// Return the layer name
func (layer Layer) Name() string {
	name := C.OGR_L_GetName(layer.cval)
	return C.GoString(name)
}

// Return the layer geometry type
func (layer Layer) Type() GeometryType {
	gt := C.OGR_L_GetGeomType(layer.cval)
	return GeometryType(gt)
}

// Return the current spatial filter for this layer
func (layer Layer) SpatialFilter() Geometry {
	geom := C.OGR_L_GetSpatialFilter(layer.cval)
	return Geometry{geom}
}

// Set a new spatial filter for this layer
func (layer Layer) SetSpatialFilter(filter Geometry) {
	C.OGR_L_SetSpatialFilter(layer.cval, filter.cval)
}

// Set a new rectangular spatial filter for this layer
func (layer Layer) SetSpatialFilterRect(minX, minY, maxX, maxY float64) {
	C.OGR_L_SetSpatialFilterRect(
		layer.cval,
		C.double(minX), C.double(minY), C.double(maxX), C.double(maxY),
	)
}

// Set a new attribute query filter
func (layer Layer) SetAttributeFilter(filter string) error {
	cFilter := C.CString(filter)
	defer C.free(unsafe.Pointer(cFilter))
	return C.OGR_L_SetAttributeFilter(layer.cval, cFilter).Err()
}

// Reset reading to start on the first featre
func (layer Layer) ResetReading() {
	C.OGR_L_ResetReading(layer.cval)
}

// Fetch the next available feature from this layer
func (layer Layer) NextFeature() *Feature {
	feature := C.OGR_L_GetNextFeature(layer.cval)
	if feature == nil {
		return nil
	}
	return &Feature{feature}
}

// Move read cursor to the provided index
func (layer Layer) SetNextByIndex(index int64) error {
	return C.OGR_L_SetNextByIndex(layer.cval, C.GIntBig(index)).Err()
}

// Fetch a feature by its index
func (layer Layer) Feature(index int64) Feature {
	feature := C.OGR_L_GetFeature(layer.cval, C.GIntBig(index))
	return Feature{feature}
}

// Rewrite the provided feature
func (layer Layer) SetFeature(feature Feature) error {
	return C.OGR_L_SetFeature(layer.cval, feature.cval).Err()
}

// Create and write a new feature within a layer
func (layer Layer) Create(feature Feature) error {
	return C.OGR_L_CreateFeature(layer.cval, feature.cval).Err()
}

// Delete indicated feature from layer
func (layer Layer) Delete(index int64) error {
	return C.OGR_L_DeleteFeature(layer.cval, C.GIntBig(index)).Err()
}

// Fetch the schema information for this layer
func (layer Layer) Definition() FeatureDefinition {
	defn := C.OGR_L_GetLayerDefn(layer.cval)
	return FeatureDefinition{defn}
}

// Fetch the spatial reference system for this layer
func (layer Layer) SpatialReference() SpatialReference {
	sr := C.OGR_L_GetSpatialRef(layer.cval)
	return SpatialReference{sr}
}

// Fetch the feature count for this layer
func (layer Layer) FeatureCount(force bool) (count int, ok bool) {
	count = int(C.OGR_L_GetFeatureCount(layer.cval, BoolToCInt(force)))
	return count, count != -1
}

// Fetch the extent of this layer
func (layer Layer) Extent(force bool) (env Envelope, err error) {
	err = C.OGR_L_GetExtent(layer.cval, &env.cval, BoolToCInt(force)).Err()
	return
}

// Test if this layer supports the named capability
func (layer Layer) TestCapability(capability string) bool {
	cString := C.CString(capability)
	defer C.free(unsafe.Pointer(cString))
	val := C.OGR_L_TestCapability(layer.cval, cString)
	return val != 0
}

// Create a new field on a layer
func (layer Layer) CreateField(fd FieldDefinition, approxOK bool) error {
	return C.OGR_L_CreateField(layer.cval, fd.cval, BoolToCInt(approxOK)).Err()
}

// Delete a field from the layer
func (layer Layer) DeleteField(index int) error {
	return C.OGR_L_DeleteField(layer.cval, C.int(index)).Err()
}

// Reorder all the fields of a layer
func (layer Layer) ReorderFields(layerMap []int) error {
	return C.OGR_L_ReorderFields(layer.cval, (*C.int)(unsafe.Pointer(&layerMap[0]))).Err()
}

// Reorder an existing field of a layer
func (layer Layer) ReorderField(oldIndex, newIndex int) error {
	return C.OGR_L_ReorderField(layer.cval, C.int(oldIndex), C.int(newIndex)).Err()
}

// Alter the definition of an existing field of a layer
func (layer Layer) AlterFieldDefn(index int, newDefn FieldDefinition, flags int) error {
	return C.OGR_L_AlterFieldDefn(layer.cval, C.int(index), newDefn.cval, C.int(flags)).Err()
}

// Begin a transation on data sources which support it
func (layer Layer) StartTransaction() error {
	return C.OGR_L_StartTransaction(layer.cval).Err()
}

// Commit a transaction on data sources which support it
func (layer Layer) CommitTransaction() error {
	return C.OGR_L_CommitTransaction(layer.cval).Err()
}

// Roll back the current transaction on data sources which support it
func (layer Layer) RollbackTransaction() error {
	return C.OGR_L_RollbackTransaction(layer.cval).Err()
}

// Flush pending changes to the layer
func (layer Layer) Sync() error {
	return C.OGR_L_SyncToDisk(layer.cval).Err()
}

// Fetch the name of the FID column
func (layer Layer) FIDColumn() string {
	name := C.OGR_L_GetFIDColumn(layer.cval)
	return C.GoString(name)
}

// Fetch the name of the geometry column
func (layer Layer) GeometryColumn() string {
	name := C.OGR_L_GetGeometryColumn(layer.cval)
	return C.GoString(name)
}

// Set which fields can be ignored when retrieving features from the layer
func (layer Layer) SetIgnoredFields(names []string) error {
	length := len(names)
	cNames := make([]*C.char, length+1)
	for i := 0; i < length; i++ {
		cNames[i] = C.CString(names[i])
		defer C.free(unsafe.Pointer(cNames[i]))
	}
	cNames[length] = (*C.char)(unsafe.Pointer(nil))

	return C.OGR_L_SetIgnoredFields(layer.cval, (**C.char)(unsafe.Pointer(&cNames[0]))).Err()
}

// Return the intersection of two layers
// Unimplemented: Intersection
// Will be new in 2.0

// Return the union of two layers
// Unimplemented: Union
// Will be new in 2.0

// Return the symmetric difference of two layers
// Unimplemented: SymDifference
// Will be new in 2.0

// Identify features in this layer with ones from the provided layer
// Unimplemented: Identity
// Will be new in 2.0

// Update this layer with features from the provided layer
// Unimplemented: Update
// Will be new in 2.0

// Clip off areas that are not covered by the provided layer
// Unimplemented: Clip
// Will be new in 2.0

// Remove areas that are covered by the provided layer
// Unimplemented: Erase
// Will be new in 2.0

/* -------------------------------------------------------------------- */
/*      Data source functions                                           */
/* -------------------------------------------------------------------- */

type DataSource struct {
	cval C.OGRDataSourceH
}

// Open a file / data source with one of the registered drivers
func OpenDataSource(name string, update int) DataSource {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	ds := C.OGROpen(cName, C.int(update), nil)
	return DataSource{ds}
}

// Open a shared file / data source with one of the registered drivers
func OpenSharedDataSource(name string, update int) DataSource {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	ds := C.OGROpenShared(cName, C.int(update), nil)
	return DataSource{ds}
}

// Drop a reference to this datasource and destroy if reference is zero
func (ds DataSource) Release() error {
	return C.OGRReleaseDataSource(ds.cval).Err()
}

// Return the number of opened data sources
func OpenDataSourceCount() int {
	count := C.OGRGetOpenDSCount()
	return int(count)
}

// Return the i'th datasource opened
func OpenDataSourceByIndex(index int) DataSource {
	ds := C.OGRGetOpenDS(C.int(index))
	return DataSource{ds}
}

// Closes datasource and releases resources
func (ds DataSource) Destroy() {
	C.OGR_DS_Destroy(ds.cval)
}

// Fetch the name of the data source
func (ds DataSource) Name() string {
	name := C.OGR_DS_GetName(ds.cval)
	return C.GoString(name)
}

// Fetch the number of layers in this data source
func (ds DataSource) LayerCount() int {
	count := C.OGR_DS_GetLayerCount(ds.cval)
	return int(count)
}

// Fetch a layer of this data source by index
func (ds DataSource) LayerByIndex(index int) Layer {
	layer := C.OGR_DS_GetLayer(ds.cval, C.int(index))
	return Layer{layer}
}

// Fetch a layer of this data source by name
func (ds DataSource) LayerByName(name string) Layer {
	cString := C.CString(name)
	defer C.free(unsafe.Pointer(cString))
	layer := C.OGR_DS_GetLayerByName(ds.cval, cString)
	return Layer{layer}
}

// Delete the layer from the data source
func (ds DataSource) Delete(index int) error {
	return C.OGR_DS_DeleteLayer(ds.cval, C.int(index)).Err()
}

// Fetch the driver that the data source was opened with
func (ds DataSource) Driver() OGRDriver {
	driver := C.OGR_DS_GetDriver(ds.cval)
	return OGRDriver{driver}
}

// Create a new layer on the data source
func (ds DataSource) CreateLayer(
	name string,
	sr SpatialReference,
	geomType GeometryType,
	options []string,
) Layer {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	length := len(options)
	opts := make([]*C.char, length+1)
	for i := 0; i < length; i++ {
		opts[i] = C.CString(options[i])
		defer C.free(unsafe.Pointer(opts[i]))
	}
	opts[length] = (*C.char)(unsafe.Pointer(nil))

	layer := C.OGR_DS_CreateLayer(
		ds.cval,
		cName,
		sr.cval,
		C.OGRwkbGeometryType(geomType),
		(**C.char)(unsafe.Pointer(&opts[0])),
	)
	return Layer{layer}
}

// Duplicate an existing layer
func (ds DataSource) CopyLayer(
	source Layer,
	name string,
	options []string,
) Layer {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	length := len(options)
	opts := make([]*C.char, length+1)
	for i := 0; i < length; i++ {
		opts[i] = C.CString(options[i])
		defer C.free(unsafe.Pointer(opts[i]))
	}
	opts[length] = (*C.char)(unsafe.Pointer(nil))

	layer := C.OGR_DS_CopyLayer(
		ds.cval,
		source.cval,
		cName,
		(**C.char)(unsafe.Pointer(&opts[0])),
	)
	return Layer{layer}
}

// Test if the data source has the indicated capability
func (ds DataSource) TestCapability(capability string) bool {
	cString := C.CString(capability)
	defer C.free(unsafe.Pointer(cString))
	val := C.OGR_DS_TestCapability(ds.cval, cString)
	return val != 0
}

// Execute an SQL statement against the data source
func (ds DataSource) ExecuteSQL(sql string, filter Geometry, dialect string) Layer {
	cSQL := C.CString(sql)
	defer C.free(unsafe.Pointer(cSQL))
	cDialect := C.CString(dialect)
	defer C.free(unsafe.Pointer(cDialect))

	layer := C.OGR_DS_ExecuteSQL(ds.cval, cSQL, filter.cval, cDialect)
	return Layer{layer}
}

// Release the results of ExecuteSQL
func (ds DataSource) ReleaseResultSet(layer Layer) {
	C.OGR_DS_ReleaseResultSet(ds.cval, layer.cval)
}

// Flush pending changes to the data source
func (ds DataSource) Sync() error {
	return C.OGR_DS_SyncToDisk(ds.cval).Err()
}

/* -------------------------------------------------------------------- */
/*      Driver functions                                                */
/* -------------------------------------------------------------------- */

type OGRDriver struct {
	cval C.OGRSFDriverH
}

// Fetch name of driver (file format)
func (driver OGRDriver) Name() string {
	name := C.OGR_Dr_GetName(driver.cval)
	return C.GoString(name)
}

// Attempt to open file with this driver
func (driver OGRDriver) Open(filename string, update int) (newDS DataSource, ok bool) {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))
	ds := C.OGR_Dr_Open(driver.cval, cFilename, C.int(update))
	return DataSource{ds}, ds != nil
}

// Test if this driver supports the named capability
func (driver OGRDriver) TestCapability(capability string) bool {
	cString := C.CString(capability)
	defer C.free(unsafe.Pointer(cString))
	val := C.OGR_Dr_TestCapability(driver.cval, cString)
	return val != 0
}

// Create a new data source based on this driver
func (driver OGRDriver) Create(name string, options []string) (newDS DataSource, ok bool) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	length := len(options)
	opts := make([]*C.char, length+1)
	for i := 0; i < length; i++ {
		opts[i] = C.CString(options[i])
		defer C.free(unsafe.Pointer(opts[i]))
	}
	opts[length] = (*C.char)(unsafe.Pointer(nil))

	ds := C.OGR_Dr_CreateDataSource(driver.cval, cName, (**C.char)(unsafe.Pointer(&opts[0])))
	return DataSource{ds}, ds != nil
}

// Create a new datasource with this driver by copying all layers of the existing datasource
func (driver OGRDriver) Copy(source DataSource, name string, options []string) (newDS DataSource, ok bool) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	length := len(options)
	opts := make([]*C.char, length+1)
	for i := 0; i < length; i++ {
		opts[i] = C.CString(options[i])
		defer C.free(unsafe.Pointer(opts[i]))
	}
	opts[length] = (*C.char)(unsafe.Pointer(nil))

	ds := C.OGR_Dr_CopyDataSource(driver.cval, source.cval, cName, (**C.char)(unsafe.Pointer(&opts[0])))
	return DataSource{ds}, ds != nil
}

// Delete a data source
func (driver OGRDriver) Delete(filename string) error {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))
	return C.OGR_Dr_DeleteDataSource(driver.cval, cFilename).Err()
}

// Add a driver to the list of registered drivers
func (driver OGRDriver) Register() {
	C.OGRRegisterDriver(driver.cval)
}

// Remove a driver from the list of registered drivers
func (driver OGRDriver) Deregister() {
	C.OGRDeregisterDriver(driver.cval)
}

// Fetch the number of registered drivers
func OGRDriverCount() int {
	count := C.OGRGetDriverCount()
	return int(count)
}

// Fetch the indicated driver by index
func OGRDriverByIndex(index int) OGRDriver {
	driver := C.OGRGetDriver(C.int(index))
	return OGRDriver{driver}
}

// Fetch the indicated driver by name
func OGRDriverByName(name string) OGRDriver {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	driver := C.OGRGetDriverByName(cName)
	return OGRDriver{driver}
}

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

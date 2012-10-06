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

func (env Envelope) SetMinX(val float64) {
	env.cval.MinX = C.double(val)
}

func (env Envelope) SetMaxX(val float64) {
	env.cval.MaxX = C.double(val)
}

func (env Envelope) SetMinY(val float64) {
	env.cval.MinY = C.double(val)
}

func (env Envelope) SetMaxY(val float64) {
	env.cval.MaxY = C.double(val)
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
/*      Field definition functions                                      */
/* -------------------------------------------------------------------- */

type FieldDefinition struct {
	cval C.OGRFieldDefnH
}

type Field struct {
	cval *C.OGRField
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

type FeatureDefinition struct {
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
	err := C.OGR_F_SetGeometry(feature.cval, geom.cval)
	return error(err)
}

// Set feature geometry, passing ownership to the feature
func (feature Feature) SetGeometryDirectly(geom Geometry) error {
	err := C.OGR_F_SetGeometryDirectly(feature.cval, geom.cval)
	return error(err)
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
	err := C.OGR_L_SetAttributeFilter(layer.cval, cFilter)
	return error(err)
}

// Reset reading to start on the first featre
func (layer Layer) ResetReading() {
	C.OGR_L_ResetReading(layer.cval)
}

// Fetch the next available feature from this layer
func (layer Layer) NextFeature() Feature {
	feature := C.OGR_L_GetNextFeature(layer.cval)
	return Feature{feature}
}

// Move read cursor to the provided index
func (layer Layer) SetNextByIndex(index int) error {
	err := C.OGR_L_SetNextByIndex(layer.cval, C.long(index))
	return error(err)
}

// Fetch a feature by its index
func (layer Layer) Feature(index int) Feature {
	feature := C.OGR_L_GetFeature(layer.cval, C.long(index))
	return Feature{feature}
}

// Rewrite the provided feature
func (layer Layer) SetFeature(feature Feature) error {
	err := C.OGR_L_SetFeature(layer.cval, feature.cval)
	return error(err)
}

// Create and write a new feature within a layer
func (layer Layer) Create(feature Feature) error {
	err := C.OGR_L_CreateFeature(layer.cval, feature.cval)
	return error(err)
}

// Delete indicated feature from layer
func (layer Layer) Delete(index int) error {
	err := C.OGR_L_DeleteFeature(layer.cval, C.long(index))
	return error(err)
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
	err = error(C.OGR_L_GetExtent(layer.cval, &env.cval, BoolToCInt(force)))
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
	err := C.OGR_L_CreateField(layer.cval, fd.cval, BoolToCInt(approxOK))
	return error(err)
}

// Delete a field from the layer
func (layer Layer) DeleteField(index int) error {
	err := C.OGR_L_DeleteField(layer.cval, C.int(index))
	return error(err)
}

// Reorder all the fields of a layer
// Unimplemented: ReorderFields

// Reorder an existing field of a layer
// Unimplemented: ReorderField

// Alter the definition of an existing field of a layer
// Unimplemented: AlterFieldDefn

// Begin a transation on data sources which support it
func (layer Layer) StartTransaction() error {
	err := C.OGR_L_StartTransaction(layer.cval)
	return error(err)
}

// Commit a transaction on data sources which support it
func (layer Layer) CommitTransaction() error {
	err := C.OGR_L_CommitTransaction(layer.cval)
	return error(err)
}

// Roll back the current transaction on data sources which support it
func (layer Layer) RollbackTransaction() error {
	err := C.OGR_L_RollbackTransaction(layer.cval)
	return error(err)
}

// Flush pending changes to the layer
func (layer Layer) Sync() error {
	err := C.OGR_L_SyncToDisk(layer.cval)
	return error(err)
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

	err := C.OGR_L_SetIgnoredFields(layer.cval, (**C.char)(unsafe.Pointer(&cNames[0])))
	return error(err)
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
	err := C.OGRReleaseDataSource(ds.cval);
	return error(err)
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
	err := C.OGR_DS_DeleteLayer(ds.cval, C.int(index))
	return error(err)
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
	err := C.OGR_DS_SyncToDisk(ds.cval)
	return error(err)
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
	err := C.OGR_Dr_DeleteDataSource(driver.cval, cFilename)
	return error(err)
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

package gdal

/*
#include "gdal.h"
#include "ogr_core.h"
#include "ogr_api.h"

#cgo linux  pkg-config: gdal
#cgo darwin pkg-config: gdal
#cgo windows LDFLAGS: -lgdal.dll
*/
import "C"
import (
	_ "unsafe"
)

func init() {
	C.OGRRegisterAll()
}

/* -------------------------------------------------------------------- */
/*      Significant constants.                                          */
/* -------------------------------------------------------------------- */

/* -------------------------------------------------------------------- */
/*      Define handle types related to various internal classes.        */
/* -------------------------------------------------------------------- */

/* -------------------------------------------------------------------- */
/*      Geometry functions                                              */
/* -------------------------------------------------------------------- */

type Geometry struct {
	cval C.OGRGeometryH
}

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

/* -------------------------------------------------------------------- */
/*      Feature definition functions                                    */
/* -------------------------------------------------------------------- */

type FeatureDefn struct {
	cval C.OGRFeatureDefnH
}

/* -------------------------------------------------------------------- */
/*      Feature functions                                               */
/* -------------------------------------------------------------------- */

type Feature struct {
	cval C.OGRFeatureH
}

/* -------------------------------------------------------------------- */
/*      Layer functions                                                 */
/* -------------------------------------------------------------------- */

type Layer struct {
	cval C.OGRLayerH
}

/* -------------------------------------------------------------------- */
/*      Data source functions                                           */
/* -------------------------------------------------------------------- */

type DataSource struct {
	cval C.OGRDataSourceH
}

type SFDriver struct {
	cval C.OGRSFDriverH
}


// Copyright 2019 AirMap Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
#ifndef GO_OGR_WKB_H_
#define GO_OGR_WKB_H_

#include <ogr_api.h>

// go_CreateFromWkb helps us in handling an API/ABI break introduced by gdal 2.3.0.
OGRErr go_CreateFromWkb(void *pabyData, OGRSpatialReferenceH hSRS, OGRGeometryH* phGeometry, int nBytes);

// go_CreateFromWkb helps us in handling an API/ABI break introduced by gdal 2.3.0.
OGRErr go_ImportFromWkb(OGRGeometryH hGeom, void* pabyData, int nSize);

// go_ExportToWkb helps us in handling an API/ABI break introduced by gdal 2.3.0.
OGRErr go_ExportToWkb(OGRGeometryH hGeom, OGRwkbByteOrder eOrder, unsigned char* pabyDstBuffer);

#endif  // GO_OGR_WKB_H_

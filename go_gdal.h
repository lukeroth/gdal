// Copyright 2011 go-gdal. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#ifndef GO_GDAL_H_
#define GO_GDAL_H_

#include <gdal.h>
#include <gdal_alg.h>
#include <gdal_utils.h>
#include <gdalwarper.h>
#include <cpl_conv.h>
#include <ogr_srs_api.h>

// transform GDALProgressFunc to go func
GDALProgressFunc goGDALProgressFuncProxyB();

#endif // GO_GDAL_H_



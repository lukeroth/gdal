// Copyright 2011 go-gdal. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "go_gdal.h"
#include "_cgo_export.h"

#include <cpl_conv.h>

static int goGDALProgressFuncProxyB_(
	double complete, 
	const char *message, 
	void *progressArg
) {
	GoInterface* args = (GoInterface*)progressArg;
	int returnVal = goGDALProgressFuncProxyA(complete, (char*)message, args);
	return (int)returnVal;
}

GDALProgressFunc goGDALProgressFuncProxyB() {
	return goGDALProgressFuncProxyB_;
}



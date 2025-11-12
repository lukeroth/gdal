package gdal

/*
#include "go_gdal.h"
*/
import "C"
import (
	"errors"
	"fmt"
)

func CPLGetLastErrorType() CPLErr {
	return CPLErr(C.CPLGetLastErrorType())
}

func CPLGetLastErrorMsg() error {
	cErrMsg := C.CPLGetLastErrorMsg()
	if cErrMsg == nil {
		return errors.New("unknown CPL error")
	}
	return errors.New(C.GoString(cErrMsg))
}

func CPLGetErr() error {
	if cplErr := CPLGetLastErrorType(); cplErr == CE_Failure || cplErr == CE_Fatal {
		return fmt.Errorf("%s: %s", ErrFailure.Error(), CPLGetLastErrorMsg())
	}
	return nil
}

func CPLGetWarn() error {
	if cplErr := CPLGetLastErrorType(); cplErr == CE_Warning {
		return fmt.Errorf("%s: %s", ErrWarning.Error(), CPLGetLastErrorMsg())
	}
	return nil
}

/* ==================================================================== */
/*      GDAL Driver                                                     */
/* ==================================================================== */

// Check if the driver returned an error
func (driver Driver) Err() error {
	if err := CPLGetErr(); err != nil {
		return err
	}
	return nil
}

// Check if the driver returned a warning
func (driver Driver) Warn() error {
	if err := CPLGetWarn(); err != nil {
		return err
	}
	return nil
}

/* ==================================================================== */
/*      GDAL Dataset                                                    */
/* ==================================================================== */

// Check if the dataset returned an error
func (dataset Dataset) Err() error {
	if err := CPLGetErr(); err != nil {
		return err
	}
	return nil
}

// Check if the dataset returned a warning
func (dataset Dataset) Warn() error {
	if err := CPLGetWarn(); err != nil {
		return err
	}
	return nil
}

/* ==================================================================== */
/*      GDAL Layer                                                      */
/* ==================================================================== */

// Check if the layer returned an error
func (layer Layer) Err() error {
	if err := CPLGetErr(); err != nil {
		return err
	}
	return nil
}

// Check if the layer returned a warning
func (layer Layer) Warn() error {
	if err := CPLGetWarn(); err != nil {
		return err
	}
	return nil
}

/* ==================================================================== */
/*      GDAL Spatial Reference                                          */
/* ==================================================================== */

// Check if the spatial reference returned an error
func (sr SpatialReference) Err() error {
	if err := CPLGetErr(); err != nil {
		return err
	}
	return nil
}

// Check if the spatial reference returned a warning
func (sr SpatialReference) Warn() error {
	if err := CPLGetWarn(); err != nil {
		return err
	}
	return nil
}

/* ==================================================================== */
/*      GDAL Feature                                                    */
/* ==================================================================== */

// Check if the feature returned an error
func (feature Feature) Err() error {
	if err := CPLGetErr(); err != nil {
		return err
	}
	return nil
}

// Check if the feature returned a warning
func (feature Feature) Warn() error {
	if err := CPLGetWarn(); err != nil {
		return err
	}
	return nil
}

/* ==================================================================== */
/*      GDAL Feature Definition                                        */
/* ==================================================================== */

// Check if the feature definition returned an error
func (fd FeatureDefinition) Err() error {
	if err := CPLGetErr(); err != nil {
		return err
	}
	return nil
}

// Check if the feature definition returned a warning
func (fd FeatureDefinition) Warn() error {
	if err := CPLGetWarn(); err != nil {
		return err
	}
	return nil
}

/* ==================================================================== */
/*      GDAL Geometry                                                   */
/* ==================================================================== */

// Check if the geometry returned an error
func (geometry Geometry) Err() error {
	if err := CPLGetErr(); err != nil {
		return err
	}
	return nil
}

// Check if the geometry returned a warning
func (geometry Geometry) Warn() error {
	if err := CPLGetWarn(); err != nil {
		return err
	}
	return nil
}

/* ==================================================================== */
/*      GDAL Geometry Field Definition                                 */
/* ==================================================================== */

// Check if the geometry field definition returned an error
func (gfd GeometryFieldDefinition) Err() error {
	if err := CPLGetErr(); err != nil {
		return err
	}
	return nil
}

// Check if the geometry field definition returned a warning
func (gfd GeometryFieldDefinition) Warn() error {
	if err := CPLGetWarn(); err != nil {
		return err
	}
	return nil
}

/* ==================================================================== */
/*      GDAL Coordinate Transform                                       */
/* ==================================================================== */

// Check if the coordinate transform returned an error
func (ct CoordinateTransform) Err() error {
	if err := CPLGetErr(); err != nil {
		return err
	}
	return nil
}

// Check if the coordinate transform returned a warning
func (ct CoordinateTransform) Warn() error {
	if err := CPLGetWarn(); err != nil {
		return err
	}
	return nil
}

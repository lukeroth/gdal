package gdal

// Check if the driver is null
func (driver Driver) IsNull() bool {
	return driver.cval == nil
}

// Check if the dataset is null
func (dataset Dataset) IsNull() bool {
	return dataset.cval == nil
}

// Check if the layer is null
func (layer Layer) IsNull() bool {
	return layer.cval == nil
}

// Check if the spatial reference is null
func (sr SpatialReference) IsNull() bool {
	return sr.cval == nil
}

// Check if the feature is null
func (feature Feature) IsNull() bool {
	return feature.cval == nil
}

// Check if the feature definition is null
func (fd FeatureDefinition) IsNull() bool {
	return fd.cval == nil
}

// Check if the geometry is null
func (geom Geometry) IsNull() bool {
	return geom.cval == nil
}

// Check if the geometry field definition is null
func (gfd GeometryFieldDefinition) IsNull() bool {
	return gfd.cval == nil
}

// Check if the coordinate transform is null
func (ct CoordinateTransform) IsNull() bool {
	return ct.cval == nil
}

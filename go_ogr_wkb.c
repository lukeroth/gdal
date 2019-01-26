#include "go_ogr_wkb.h"

#include <gdal_version.h>

#if GDAL_COMPUTE_VERSION(GDAL_VERSION_MAJOR, GDAL_VERSION_MINOR, GDAL_VERSION_PATCH) >= GDAL_COMPUTE_VERSION(2, 3, 0)

OGRErr go_CreateFromWkb(void* pabyData, OGRSpatialReferenceH hSRS, OGRGeometryH* phGeometry, int nBytes) {
    return OGR_G_CreateFromWkb(pabyData, hSRS, phGeometry, nBytes);
}

OGRErr go_ImportFromWkb(OGRGeometryH hGeom, void* pabyData, int nSize) {
    return OGR_G_ImportFromWkb(hGeom, pabyData, nSize);
}

OGRErr go_ExportToWkb(OGRGeometryH hGeom, OGRwkbByteOrder eOrder, unsigned char* pabyDstBuffer) {
    return OGR_G_ExportToWkb(hGeom, eOrder, pabyDstBuffer);
}

#elif GDAL_COMPUTE_VERSION(GDAL_VERSION_MAJOR, GDAL_VERSION_MINOR, GDAL_VERSION_PATCH) < GDAL_COMPUTE_VERSION(2, 3, 0)

OGRErr go_CreateFromWkb(void* pabyData, OGRSpatialReferenceH hSRS, OGRGeometryH* phGeometry, int nBytes) {
    return OGR_G_CreateFromWkb((unsigned char*)pabyData, hSRS, phGeometry, nBytes);
}

OGRErr go_ImportFromWkb(OGRGeometryH hGeom, void* pabyData, int nSize) {
    return OGR_G_ImportFromWkb(hGeom, (unsigned char*)pabyData, nSize);
}

OGRErr go_ExportToWkb(OGRGeometryH hGeom, OGRwkbByteOrder eOrder, unsigned char* pabyDstBuffer) {
    return OGR_G_ExportToWkb(hGeom, eOrder, pabyDstBuffer);
}

#endif // GDAL_COMPUTE_VERSION(GDAL_VERSION_MAJOR, GDAL_VERSION_MINOR, GDAL_VERSION_PATCH) >= GDAL_COMPUTE_VERSION(2, 3, 0)
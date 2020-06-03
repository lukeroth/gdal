// +build linux,amd64 darwin,amd64

package gdal

/*
#cgo LDFLAGS: -lgdal -leccodes -leccodes_memfs -lpng -laec -ljasper -lopenjp2 -lpthread -fopenmp -lz -lm
*/
import "C"

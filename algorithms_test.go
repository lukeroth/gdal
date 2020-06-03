package gdal

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
	"testing"
)

func readGridFile(filename string) (x, y, z []float64, err error) {
	var b []byte
	if b, err = ioutil.ReadFile(filename); err != nil {
		return
	}
	arr := strings.Split(string(b), "\n")
	x, y, z = make([]float64, len(arr)), make([]float64, len(arr)), make([]float64, len(arr))
	for i, el := range arr {
		xyz := strings.Split(el, ",")
		if len(xyz) != 3 {
			err = errors.New("wrong input file format, should be CSV with 3 columns: y,x,z")
		}

		if y[i], err = strconv.ParseFloat(xyz[0], 64); err != nil {
			return
		}
		if x[i], err = strconv.ParseFloat(xyz[1], 64); err != nil {
			return
		}
		if z[i], err = strconv.ParseFloat(xyz[2], 64); err != nil {
			return
		}
	}
	return
}

func TestGridCreate(t *testing.T) {
	x, y, z, err := readGridFile("testdata/grid.csv")
	if err != nil {
		t.Fatalf("failed to readGridFile: %v", err)
		return
	}
	var nX, nY uint = 420, 470

	// finding max and min values
	var xMin, xMax, yMin, yMax = math.MaxFloat64, -math.MaxFloat64, math.MaxFloat64, -math.MaxFloat64
	for i := range x {
		if x[i] < xMin {
			xMin = x[i]
		}
		if x[i] > xMax {
			xMax = x[i]
		}
		if y[i] < yMin {
			yMin = y[i]
		}
		if y[i] > yMax {
			yMax = y[i]
		}
	}

	fmt.Println("Calling gdal.GridCreate")
	data, err := GridCreate(
		GA_Linear,
		GridLinearOptions{Radius: -1, NoDataValue: 0},
		x, y, z,
		xMin, xMax, yMin, yMax,
		nX, nY,
		DummyProgress,
		nil,
	)
	if err != nil {
		t.Errorf("GridCreate: %v", err)
	}
	if expectedDataLen := int(nX * nY); len(data) != expectedDataLen {
		t.Errorf("expected length of data equal to %d", expectedDataLen)
	}
}

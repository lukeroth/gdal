package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"

	"github.com/lukeroth/gdal"
)

func readFile(filename string) (x, y, z []float64, err error) {
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

func step(xMin, xMax, yMin, yMax float64, nX, nY uint) (stepX float64, stepY float64) {
	return (xMax - xMin) / float64(nX), (yMax - yMin) / float64(nY)
}

func coordsForIndex(i int, xMin, xStep, yMin, yStep float64, nX uint) (x float64, y float64) {
	return xStep*float64(i%int(nX)) + xMin + xStep/2, yStep*float64(i/int(nX)) + yMin + yStep/2
}

func writeFile(filename string, z []float64, xMin, xMax, yMin, yMax float64, nX, nY uint) (err error) {
	var b bytes.Buffer
	xStep, yStep := step(xMin, xMax, yMin, yMax, nX, nY)
	for i := range z {
		x, y := coordsForIndex(i, xMin, xStep, yMin, yStep, nX)
		if _, err = b.WriteString(fmt.Sprintf("%f %f %f\n", x, y, z[i])); err != nil {
			return
		}
	}
	return ioutil.WriteFile(filename, b.Bytes(), 0644)
}

func main() {
	flag.Parse()
	filename := flag.Arg(0)
	if filename == "" {
		fmt.Println("Usage: grid [filename]")
		fmt.Println("Should be a CSV with columns y,x,z")
		return
	}
	fmt.Printf("Filename: %s\n", filename)

	fmt.Println("Reading file")
	x, y, z, err := readFile(filename)
	if err != nil {
		fmt.Println(err.Error())
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
	data, err := gdal.GridCreate(
		gdal.GA_InverseDistancetoAPower,
		gdal.GridInverseDistanceToAPowerOptions{Power: 2},
		x, y, z,
		xMin, xMax, yMin, yMax,
		nX, nY,
		gdal.DummyProgress,
		nil,
	)
	if err != nil {
		fmt.Println("gdal.GridCreate error: ", err.Error())
		return
	}

	fmt.Println("Writing file")
	if err := writeFile(filename+".out", data, xMin, xMax, yMin, yMax, nX, nY); err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("End program")
}

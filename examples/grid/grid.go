package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/mtfelian/gdal"
)

func readFile(filename string) (x, y, z []float64, err error) {
	var b []byte
	if b, err = ioutil.ReadFile(filename); err != nil {
		return
	}
	s := string(b)
	arr := strings.Split(s, "\n")
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

func writeFile(filename string, z []float64) error {
	var s string
	for i := range z {
		s += fmt.Sprintf("%.5f\n", z[i])
	}
	return ioutil.WriteFile(filename, []byte(s), 0644)
}

func main() {
	flag.Parse()
	filename := flag.Arg(0)
	if filename == "" {
		fmt.Printf("Usage: grid [filename]\n")
		return
	}
	fmt.Printf("Filename: %s\n", filename)

	fmt.Printf("Reading file\n")
	x, y, z, err := readFile(filename)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//var nX, nY uint = 420, 470
	var nX, nY uint = 420, 470

	fmt.Printf("Allocating buffer\n")
	buffer := make([]float64, nX*nY)
	fmt.Printf("Calling gdal.CreateGrid\n")
	if err := gdal.CreateGrid(
		gdal.GGA_Linear,
		[]string{"0", "0.0"},
		x, y, z,
		nX, nY,
		buffer,
		gdal.DummyProgress,
		nil,
	); err != nil {
		fmt.Println("gdal.CreateGrid error: ", err.Error())
		return
	}

	fmt.Printf("Writing file\n")
	if err := writeFile(filename+".out", buffer); err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("End program\n")
}

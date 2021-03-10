package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/lukeroth/gdal"
)

func main() {
	flag.Parse()
	inputFile := flag.Arg(0)
	outputFile := flag.Arg(1)
	if len(flag.Args()) < 3 {
		fmt.Printf("Usage: warp inputFile outputFile options\n")
		return
	}
	options := flag.Args()[2:]
	if inputFile == "" {
		fmt.Printf("Usage: warp inputFile outputFile options\n")
		return
	}
	fmt.Printf("Input filename: %s\n", inputFile)
	if outputFile == "" {
		fmt.Printf("Usage: warp inputFile outputFile options\n")
		return
	}
	fmt.Printf("Output filename: %s\n", outputFile)

	fmt.Printf("GDALWarp options: %s\n", strings.Join(options, " "))

	ds, err := gdal.Open(inputFile, gdal.ReadOnly)
	if err != nil {
		log.Fatal(err)
	}
	defer ds.Close()

	outputDs, err := gdal.Warp(outputFile, nil, []gdal.Dataset{ds}, options)
	if err != nil {
		log.Fatal(err)
	}
	defer outputDs.Close()

	fmt.Printf("End program\n")
}

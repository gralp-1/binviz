package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"
)

// global to "flatten out" data transfer
var modifier *float64
var outputFileName *string
var dimension *int

const version string = "0.0.1"

type Point struct {
	Coordinate []uint8 // X, Y, Z, W
	Value      uint8
}

func (p *Point) ToString() string {
	coordinateString := ""
	for _, v := range p.Coordinate {
		coordinateString += fmt.Sprintf("%d ", v)
	}
	return fmt.Sprintf("%s%d\n", coordinateString, p.Value)
}

func main() {
	// Command line arguments
	fileName := flag.String("file", "", "Input file name")
	outputFileName = flag.String("output", "output.viz", "Output file name")
	dimension = flag.Int("dimension", 2, "Dimensionality of output (2 or 3)")
	modifier = flag.Float64("brightness", 1.0, "Brightness modifier")
	// TODO check these arguments for validity
	flag.Parse()

	// log arguments
	fmt.Println("File:       ", *fileName)
	fmt.Println("Output:     ", *outputFileName)
	fmt.Println("Dimension:  ", *dimension)
	fmt.Println("Brightness: ", *modifier)
	file, err := os.OpenFile(*fileName, os.O_RDONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	// read file to byte array
	stat, err := file.Stat()
	fileBytes := make([]byte, stat.Size())
	_, err = file.Read(fileBytes)
	file.Close()

	if err != nil {
		log.Fatal(err)
	}

	// call appropriate function
	points := Binviz(fileBytes)
	err = Serialise(points)
	if err != nil {
		log.Fatal(err)
	}
}

func Binviz(fileBytes []byte) []Point {
	// similar to Binviz2D, but with variable dimensionality
	// 0xXXYYZZWW : frequency
	var freqencyMap map[uint32]int = make(map[uint32]int)
	// count number of occurrences of each byte pair
	for i := 0; i < len(fileBytes)-*dimension; i++ {
		var key uint32 = 0
		for j := 0; j < *dimension; j++ {
			key = key << 8
			key += uint32(fileBytes[i+j])
		}
		freqencyMap[key]++
	}
	// find max value
	max := 0
	for _, v := range freqencyMap {
		if v > max {
			max = v
		}
	}
	// normalise values and adjust brightness
	var points []Point = make([]Point, len(freqencyMap))
	i := 0
	for k, v := range freqencyMap {
		points[i].Coordinate = make([]uint8, *dimension)
		for j := 0; j < *dimension; j++ {
			points[i].Coordinate[*dimension-j-1] = uint8(k >> (8 * j))
		}
		points[i].Value = uint8(AdjustBrightness(float64(v)/float64(max)) * 255)
		i++
	}
	return points
}
func Serialise(points []Point) error {
	// Write header
	f, err := os.OpenFile(*outputFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer f.Close()
	f.WriteString(fmt.Sprintln("VERSION", version))
	f.WriteString(fmt.Sprintln("DIMENSION", *dimension))
	// Write points
	for _, p := range points {
		f.WriteString(p.ToString())
	}
	return nil
}

func AdjustBrightness(pixel float64) float64 {
	// some good 'ol dereferencing in the middle of an equation for ya
	return math.Min(1.0-math.Pow((1.0-(*modifier)*pixel), 5.0), 1)
}

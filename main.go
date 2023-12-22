package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

var modifier float64

func AdjustBrightness(pixel float64) float64 {
	return math.Min(1.0-math.Pow((1.0-modifier*pixel), 5.0), 1)
}

func main() {
	// Command line arguments
	fileName := os.Args[1]
	outputName := os.Args[2]
	dimension, err := strconv.Atoi(os.Args[3])
	modifier, err = strconv.ParseFloat(os.Args[4], 64)
	if err != nil {
		log.Print(err)
		log.Print("Defaulting to 2D")
	}
	// log arguments
	log.Printf("File: %s", fileName)
	log.Printf("Output: %s", outputName)
	log.Printf("Dimension: %d", dimension)
	log.Printf("Brightness Modifier: %f", modifier)

	file, err := os.OpenFile(fileName, os.O_RDONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	// read file to byte array
	stat, err := file.Stat()
	fileBytes := make([]byte, stat.Size())
	_, err = file.Read(fileBytes)
	if err != nil {
		log.Fatal(err)
	}

	// call appropriate function
	switch dimension {
	case 2:
		Binviz2D(fileBytes, outputName)
	case 3:
		Binviz3D(fileBytes, outputName)
	default:
		log.Fatal("Invalid dimensionality")
	}
	file.Close()
}
func Binviz3D(fileBytes []byte, outputName string) {
	// initialise 256x256 array with 0s
	var pixels [256][256][256]int = [256][256][256]int{}
	// count number of occurrences of each byte pair
	for i := 0; i < len(fileBytes)-2; i++ {
		b1 := fileBytes[i]
		b2 := fileBytes[i+1]
		b3 := fileBytes[i+2]
		pixels[b1][b2][b3]++
	}
	// find max value
	max := 0
	for i := 0; i < 256; i++ {
		for j := 0; j < 256; j++ {
			for k := 0; k < 256; k++ {
				if pixels[i][j][k] > max {
					max = pixels[i][j][k]
				}
			}
		}
	}
	// normalise values
	for i := 0; i < 256; i++ {
		for j := 0; j < 256; j++ {
			for k := 0; k < 256; k++ {
				pixels[i][j][k] = int(AdjustBrightness(float64(pixels[i][j][k])/float64(max)) * 255)
			}
		}
	}
	WritePCD(pixels, outputName)
}

func Binviz2D(fileBytes []byte, outputName string) {
	// initialise 256x256 array with 0s
	var pixels [256][256]int = [256][256]int{}
	// count number of occurrences of each byte pair
	for i := 0; i < len(fileBytes)-1; i++ {
		b1 := fileBytes[i]
		b2 := fileBytes[i+1]
		pixels[b1][b2]++
	}
	// find max value
	max := 0.0
	for i := 0; i < 256; i++ {
		for j := 0; j < 256; j++ {
			f := 0.0
			if pixels[i][j] > 0 {
				f = math.Log(float64(pixels[i][j]))
				if f > float64(max) {
					max = f
				}
			}
		}
	}
	// normalise values
	for i := 0; i < 256; i++ {
		for j := 0; j < 256; j++ {
			t := math.Log(float64(pixels[i][j])) / max
			pixels[i][j] = int(t * 255)
		}
	}
	// write to file
	WritePPM(pixels, outputName)
	// convert to png
}
func WritePCD(pixels [256][256][256]int, fName string) error {
	f, err := os.OpenFile(fName, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	if err := f.Truncate(0); err != nil {
		return err
	}

	f.WriteString("VERSION 0.7\n")
	f.WriteString("FIELDS x y z rgb\n")
	f.WriteString("SIZE 4 4 4 4\n")
	f.WriteString("TYPE I I I I\n")
	f.WriteString("COUNT 1 1 1 1\n")
	f.WriteString("WIDTH 256\n")
	f.WriteString("HEIGHT 256\n")
	f.WriteString("VIEWPOINT 0 0 0 1 0 0 0\n")
	f.WriteString("POINTS 65536\n")
	f.WriteString("DATA ascii\n")
	for x := 0; x < 256; x++ {
		for y := 0; y < 256; y++ {
			for z := 0; z < 256; z++ {
				val := pixels[x][y][z]
				if val != 0 {
					f.WriteString(fmt.Sprintf("%d %d %d %d %d %d\n", x, y, z, val, val, val))
				}
			}
		}
	}
	return nil
}

func WritePPM(pixels [256][256]int, fName string) error {
	f, err := os.OpenFile(fName, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	if err := f.Truncate(0); err != nil {
		return err
	}
	f.WriteString("P3\n256 256\n255\n")
	for i := 0; i < 256; i++ {
		for j := 0; j < 256; j++ {
			f.WriteString(fmt.Sprintf("%d %d %d\n", pixels[i][j], pixels[i][j], pixels[i][j]))
		}
	}
	return nil
}
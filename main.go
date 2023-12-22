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

func main() {
	// Command line arguments
	fileName := flag.String("file", "", "Input file name")
	outputName := flag.String("output", "output.ppm", "Output file name")
	dimension := flag.Int("dimension", 2, "Dimensionality of output (2 or 3)")
	modifier = flag.Float64("brightness", 1.0, "Brightness modifier")
	flag.Parse()

	// log arguments
	fmt.Println("File:       ", *fileName)
	fmt.Println("Output:     ", *outputName)
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
	switch *dimension {
	case 2:
		var pixels = Binviz2D(fileBytes)
		WritePPM(pixels, *outputName)
	case 3:
		var pixels = Binviz3D(fileBytes)
		WritePCD(pixels, *outputName)
	default:
		log.Fatal("Invalid dimensionality")
	}
}
func AdjustBrightness(pixel float64) float64 {
	// some good 'ol dereferencing in the middle of an equation for ya
	return math.Min(1.0-math.Pow((1.0-(*modifier)*pixel), 5.0), 1)
}
func AdjustBrightnessSigmoid(pixel float64) float64 {
	return math.Pow(1+math.Pow(10, -20*pixel), -10)
}

func Binviz3D(fileBytes []byte) [256][256][256]int {
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
	return pixels
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

func Binviz2D(fileBytes []byte) [256][256]int {
	// initialise 256x256 array with 0s
	var pixels [256][256]int = [256][256]int{}
	// count number of occurrences of each byte pair
	for i := 0; i < len(fileBytes)-1; i++ {
		b1 := fileBytes[i]
		b2 := fileBytes[i+1]
		pixels[b1][b2]++
	}
	// find max value
	max := 0
	for i := 0; i < 256; i++ {
		for j := 0; j < 256; j++ {
			if pixels[i][j] > max {
				max = pixels[i][j]
			}
		}
	}
	// normalise values and adjust brightness
	for i := 0; i < 256; i++ {
		for j := 0; j < 256; j++ {
			pixels[i][j] = int(AdjustBrightness(float64(pixels[i][j])/float64(max)) * 255)
		}
	}
	// write to file
	return pixels
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

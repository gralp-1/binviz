# Binviz
## What is it?
Binviz is a binary visualization tool. It takes a file and produces an image or point cloud that visualizes the binary data.

## How does it work?
### 2D image
Each consecutive pair of bytes is a coordinate to a 256x256 image. The image counts the occurance of each pair of bytes which is used to represent the brightness of the pixel at that coordinate.
Before the image is created, the data is normalized to the range 0-255 with an additional transformation function to amplify lower values.

### 3D point cloud
Each consecutive triplet of bytes is a coordinate to a 256x256x256 point cloud. The point cloud counts the occurance of each triplet of bytes which is used to represent the brightness of the pixel at that coordinate.

## How do I use it?
```
go build .
./binviz -file <input file> -output <output name> -dimension <dimension> -brightness <brightness> # output name, brightness and dimension are optional
```
This will produce either a PPM image file or a PCD point cloud file depending on the dimension argument.

## Examples
Input: <br>
![zezin](https://github.com/gralp-1/binviz/blob/main/examples/zezin.gif) <br>
Output 2D: <br>
![image](https://github.com/gralp-1/binviz/assets/62028969/a867a475-ef86-426f-be10-e274bd8ec6f6) <br>
Output 3D: <br>
![image](https://github.com/gralp-1/binviz/blob/main/examples/binviz-zezin-3D.png)

## File format
I can't find any file formats that support 4 dimensions, so time to make my own. <br>
I don't know what I'm going to call it yet and I also have no idea what I'm doing. <br>
This format is WIP and will have the following features:
- 2D, 3D and 4D support
- Colour and greyscale support

### Example
```
// Header (always in ascii)
VERSION <float>
DIMENSION <2 | 3 | 4>
COLOUR <true | false>
ENCODING <ASCII | BINARY>
// File start
x y r g b     // 2D ASCII colour
x y val       // 2D ASCII greyscale
x y z r g b   // 3D ASCII colour
x y z val     // 3D ASCII greyscale
x y z w r g b // 4D ASCII colour
x y z w val   // 4D ASCII greyscale
```
Each line is delimited by a newline character. <br>
Comments are added with `//` and are ignored by the parser. <br>
all fields must be from 0-255. <br>
If encoded in binary, each field is one byte and everything is just assumed to be in the correct order, if bytes are missing / there are extra bytes, you'll just get something strange in the visualzer <br>
For example a point at (34,35,69,255) with a colour of (20,40,10) would be encoded as `34 35 69 20 40 10` in ASCII and `22 23 45 14 28 0A` in binary.

# TODO
- [x] Refactor to make consistent
- [x] Add proper command line flags & arguments
- [ ] Look into 4D file formats and visualisation
- [ ] Make a custom visualiser and file format which supports 2, 3 and 4 dimensions
- [ ] Make a way to change the brightness adjustment function
- [ ] Optimise

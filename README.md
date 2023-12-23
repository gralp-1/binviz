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
```sh
go build .
./binviz -file <input file> -output <output name> -dimension <dimension> -brightness <brightness> # output name, brightness and dimension are optional
```

## Examples
Input: <br>
![zezin](https://github.com/gralp-1/binviz/blob/main/examples/zezin.gif) <br>
Output 2D: <br>
![image](https://github.com/gralp-1/binviz/assets/62028969/a867a475-ef86-426f-be10-e274bd8ec6f6) <br>
Output 3D: <br>
![image](https://github.com/gralp-1/binviz/blob/main/examples/binviz-zezin-3D.png)

## File format
I can't find any file formats that support 4 dimensions, so time to make my own. <br>
The file format is very simple, it's just a header followed by a list of points. It takes inspiration from the [.ppm](https://netpbm.sourceforge.net/doc/ppm.html) and [.pcd](https://pointclouds.org/documentation/tutorials/pcd_file_format.html) formats. <br>

### Example
```
VERSION <semver string>
DIMENSION <2 | 3 | 4>
x y val       // 2D
x y z val     // 3D
x y z w val   // 4D
```
All fields must be from 0-255. <br>

# TODO
- [x] Refactor to make consistent
- [x] Add proper command line flags & arguments
- [ ] Adjustable brightness curves
- [ ] Make vizualiser

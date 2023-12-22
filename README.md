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
./binviz <input file> <output name> <dimension>
```
This will produce either a PPM image file or a PCD point cloud file depending on the dimension argument.

## Examples
Input: <br>
![zezin](https://github.com/gralp-1/binviz/blob/main/examples/zezin.gif) <br>
Output 2D: <br>
![image](https://github.com/gralp-1/binviz/assets/62028969/a867a475-ef86-426f-be10-e274bd8ec6f6) <br>
Output 3D: <br>
![image](https://github.com/gralp-1/binviz/blob/main/examples/binviz-zezin-3D.png)

# TODO
- [x] Refactor to make consistent
- [ ] Add proper command line flags & arguments
- [ ] Look into 4D file formats and visualisation

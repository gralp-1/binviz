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
Input:     ![zezin](ADDME)
Output:    ![zezin 2D](ADDME)
Output 3D: ![zezin 3D](ADDME)

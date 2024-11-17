# Bitmap Command Line Tool

## Overview
The Bitmap Command Line Tool allows you to manipulate bitmap images using various commands. This tool provides functionality for displaying header information, applying filters, rotating, mirroring, and cropping bitmap images.

## Table of Contents
- [Usage](#usage)
- [Commands](#commands)
- [Flags](#flags)
- [Options](#options)
  - [Mirror Options](#mirror-options)
  - [Filter Options](#filter-options)
  - [Rotate Options](#rotate-options)
  - [Crop Options](#crop-options)
- [Header Command](#header-command)
- [License](#license)

## Usage
```bash
bitmap <command> [arguments]
```

## Commands
- **header**: Prints bitmap file header information.
- **apply**: Applies processing to the image and saves it to the file.

## Flags
- `-h`, `--help`: Prints program usage information.

## Options

### Mirror Options
- `--mirror`: Mirrors a bitmap image either horizontally or vertically.
  - `--mirror=horizontal`: Mirrors the image horizontally.
  - `--mirror=vertical`: Mirrors the image vertically.

### Filter Options
- `--filter`: Applies various filters to the image.
  - `--filter=blue`: Retains only the blue channel.
  - `--filter=red`: Retains only the red channel.
  - `--filter=green`: Retains only the green channel.
  - `--filter=grayscale`: Converts the image to grayscale.
  - `--filter=negative`: Applies a negative filter.
  - `--filter=pixelate`: Applies a pixelation effect (default block size: 20 pixels).
  - `--filter=blur`: Applies a blur effect.

### Rotate Options
- `--rotate`: Rotates a bitmap image by a specified angle.
  - `--rotate=90`: Rotates the image 90 degrees clockwise.
  - `--rotate=180`: Rotates the image 180 degrees.
  - `--rotate=270`: Rotates the image 270 degrees clockwise.
  - `--rotate=left`: Rotates the image counterclockwise.
  - `--rotate=-90`: Rotates the image 90 degrees counterclockwise.
  - `--rotate=-180`: Rotates the image 180 degrees.
  - `--rotate=-270`: Rotates the image 270 degrees counterclockwise.
  - `--rotate=right`: Rotates the image clockwise.

### Crop Options
- `--crop`: Trims a bitmap image according to specified parameters.
  - Accepts values in the format: `OffsetX-OffsetY-Width-Height`
  - Example: `--crop=10-20-100-200` will crop starting from (10, 20) with width 100 and height 200.

## Header Command
### Usage
```bash
bitmap header <source_file>
```

### Description
Prints bitmap file header information.

### BMP Header
- **FileType**: indicates that the file is a bitmap file.
- **FileSizeInBytes**: total size of the bitmap file in bytes.
- **HeaderSize**: size of the header in bytes.

### DIB Header
- **DibHeaderSize**: size of the DIB header in bytes.
- **WidthInPixels**: width of the image in pixels.
- **HeightInPixels**: height of the image in pixels.
- **PixelSizeInBits**: number of bits used for each pixel; 24 bits means true color.
- **ImageSizeInBytes**: size of the raw image data in bytes.

## License
This project is licensed under a free license. You can use, modify, and distribute it without restrictions.
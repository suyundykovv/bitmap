package general

import (
	"fmt"
	"os"
)

// header
func PrintHeaderHelp() {
	if len(os.Args) < 3 {
		fmt.Println("Use help-flag:\n   bitmap header <flag>\n\nflags:\n   -h, --help   prints header usage information")
		os.Exit(1)
	}
	for _, arg := range os.Args {
		if arg == "--help" || arg == "-h" || arg == "-help" || arg == "--h" {
			fmt.Println("Usage:\n   bitmap header <source_file>\n\nDescription:\n   prints bitmap file header information\n")
			fmt.Println("BMP Header:")
			fmt.Println("   - FileType:            indicates that the file is a bitmap file.")
			fmt.Println("   - FileSizeInBytes:     total size of the bitmap file in bytes.")
			fmt.Println("   - HeaderSize:          size of the header in bytes.")
			fmt.Println("\nDIB Header:")
			fmt.Println("   - DibHeaderSize:       size of the DIB header in bytes.")
			fmt.Println("   - WidthInPixels:       width of the image in pixels.")
			fmt.Println("   - HeightInPixels:      height of the image in pixels.")
			fmt.Println("   - PixelSizeInBits:     number of bits used for each pixel; 24 bits means true color.")
			fmt.Println("   - ImageSizeInBytes:    size of the raw image data in bytes.)")
			os.Exit(0)
		}
	}
}

// apply
func PrintApplyHelp() {
	if len(os.Args) < 3 {
		fmt.Println("Use help-flag:\n   bitmap header <flag>\n\nflags:\n   -h, --help   prints header usage information")
		os.Exit(1)
	}

	for _, arg := range os.Args {
		if arg == "--help" || arg == "-h" || arg == "-help" || arg == "--h" {
			fmt.Println("Usage:\n   bitmap apply [options] <source_file> <output_file>\n")
			fmt.Println("The options are:")
			fmt.Println("\n   --mirror: mirrors a bitmap image either horizontally or vertically.")
			fmt.Println("   [Options]:")
			fmt.Println("       --mirror=horizontal   mirrors the image horizontally.")
			fmt.Println("       --mirror=vertical     mirrors the image vertically.")
			fmt.Println("\n    --filter: apply various filters to the image.")
			fmt.Println("   [Options]:")
			fmt.Println("       --filter=blue         retains only the blue channel.")
			fmt.Println("       --filter=red          retains only the red channel.")
			fmt.Println("       --filter=green        retains only the green channel.")
			fmt.Println("       --filter=grayscale    converts the image to grayscale.")
			fmt.Println("       --filter=negative     applies a negative filter.")
			fmt.Println("       --filter=pixelate     applies a pixelation effect (default block size: 20 pixels).")
			fmt.Println("       --filter=blur         applies a blur effect.")
			fmt.Println("\n    --rotate: rotate a bitmap image by a specified angle.")
			fmt.Println("   [Options]:")
			fmt.Println("       --rotate=90           rotates the image 90 degrees clockwise.")
			fmt.Println("       --rotate=180          rotates the image 180 degrees.")
			fmt.Println("       --rotate=270          rotates the image 270 degrees clockwise.")
			fmt.Println("       --rotate=left         rotates the image counterclockwise.")
			fmt.Println("       --rotate=-90          rotates the image 90 degrees counterclockwise.")
			fmt.Println("       --rotate=-180         rotates the image 180 degrees.")
			fmt.Println("       --rotate=-270         rotates the image 270 degrees counterclockwise.")
			fmt.Println("       --rotate=right        rotates the image clockwise.")
			fmt.Println("\n    --crop: trims a bitmap image according to specified parameters.")
			fmt.Println("   [Important Note]:")
			fmt.Println("       Accepts values in the format: OffsetX-OffsetY-Width-Height")
			fmt.Println("       Example: --crop=10-20-100-200 will crop starting from (10, 20) with width 100 and height 200.")
			os.Exit(0)
		}
	}
}

// bitmap
func PrintProgrammHelp() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:\n   bitmap <command> [arguments]\n\nThe commands are:\n   header    prints bitmap file header information\n   apply     applies processing to the image and saves it to the file\n\nCommands help:\n   bitmap <command> [flag]\n")
		fmt.Println("   flags:")
		fmt.Println("       -h, --help   prints program usage information")
		fmt.Println("\n   --mirror: mirrors a bitmap image either horizontally or vertically.")
		fmt.Println("   [Options]:")
		fmt.Println("       --mirror=horizontal   mirrors the image horizontally.")
		fmt.Println("       --mirror=vertical     mirrors the image vertically.")
		fmt.Println("\n    --filter: apply various filters to the image.")
		fmt.Println("   [Options]:")
		fmt.Println("       --filter=blue         retains only the blue channel.")
		fmt.Println("       --filter=red          retains only the red channel.")
		fmt.Println("       --filter=green        retains only the green channel.")
		fmt.Println("       --filter=grayscale    converts the image to grayscale.")
		fmt.Println("       --filter=negative     applies a negative filter.")
		fmt.Println("       --filter=pixelate     applies a pixelation effect (default block size: 20 pixels).")
		fmt.Println("       --filter=blur         applies a blur effect.")
		fmt.Println("\n    --rotate: rotate a bitmap image by a specified angle.")
		fmt.Println("   [Options]:")
		fmt.Println("       --rotate=90           rotates the image 90 degrees clockwise.")
		fmt.Println("       --rotate=180          rotates the image 180 degrees.")
		fmt.Println("       --rotate=270          rotates the image 270 degrees clockwise.")
		fmt.Println("       --rotate=left         rotates the image counterclockwise.")
		fmt.Println("       --rotate=-90          rotates the image 90 degrees counterclockwise.")
		fmt.Println("       --rotate=-180         rotates the image 180 degrees.")
		fmt.Println("       --rotate=-270         rotates the image 270 degrees counterclockwise.")
		fmt.Println("       --rotate=right        rotates the image clockwise.")
		fmt.Println("\n    --crop: trims a bitmap image according to specified parameters.")
		fmt.Println("   [Important Note]:")
		fmt.Println("       Accepts values in the format: OffsetX-OffsetY-Width-Height")
		fmt.Println("       Example: --crop=10-20-100-200 will crop starting from (10, 20) with width 100 and height 200.")
		os.Exit(0)
	}
}

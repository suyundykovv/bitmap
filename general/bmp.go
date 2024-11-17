package general

import (
	"encoding/binary"
	"fmt"
	"os"
)

// Pixel представляет пиксель в формате RGB
type Pixel struct {
	R, G, B uint8
}

// BMPHeader представляет заголовок BMP
type BMPHeader struct {
	FileType         [2]byte // 1
	FileSize         uint32  // 2
	Reserved1        uint16  // 3
	Reserved2        uint16  // 4
	PixelArrayOffset uint32  // 5
	DIBHeaderSize    uint32  // 6
	Width            uint32  // 7
	Height           uint32  // 8
	Planes           uint16  // 9
	BitCount         uint16  // 10
	Compression      uint32  // 11
	ImageSize        uint32  // 12
	XPixelsPerMeter  uint32  // 13
	YPixelsPerMeter  uint32  // 14
	ColorsUsed       uint32  // 15
	ImportantColors  uint32  // 16
}

// LoadHeader загружает заголовок BMP файла
func LoadHeader(filename string) (BMPHeader, error) {
	file, err := os.Open(filename)
	if err != nil {
		return BMPHeader{}, err
	}
	defer file.Close()

	var header BMPHeader
	err = binary.Read(file, binary.LittleEndian, &header)
	if err != nil {
		return BMPHeader{}, err
	}

	if string(header.FileType[:]) != "BM" {
		return BMPHeader{}, fmt.Errorf("not a BMP file")
	}
	// fmt.Println(header.BitCount)

	if header.BitCount != 24 {
		return BMPHeader{}, fmt.Errorf("not a 24 file")
	}

	return header, nil
}

// LoadBMP загружает BMP файл и возвращает заголовок и пиксели
func LoadBMP(filename string) (BMPHeader, [][]Pixel, error) {
	file, err := os.Open(filename)
	if err != nil {
		return BMPHeader{}, nil, err
	}
	defer file.Close()

	var header BMPHeader
	err = binary.Read(file, binary.LittleEndian, &header)
	if header.Planes != 1 {
		fmt.Println("ERROR planes", header.Planes)
		os.Exit(1)
	} else if header.BitCount != 24 {
		fmt.Println("ERROR bits")
		os.Exit(1)
	}
	if err != nil {
		return BMPHeader{}, nil, err
	}
	if string(header.FileType[:]) != "BM" {
		return BMPHeader{}, nil, fmt.Errorf("not a BMP file")
	}

	image := make([][]Pixel, header.Height)
	for i := range image {
		image[i] = make([]Pixel, header.Width)
	}

	bytesPerRow := int(header.Width) * 3   // 3 байта на пиксель (R, G, B)
	padding := (4 - (bytesPerRow % 4)) % 4 // padding для выравнивания по 4 байта
	rowSize := bytesPerRow + padding
	// Считываем данные изображения
	data := make([]byte, rowSize*int(header.Height))
	_, err = file.Read(data)
	if err != nil {
		return BMPHeader{}, nil, err
	}

	for i := 0; i < int(header.Height); i++ {
		for j := 0; j < int(header.Width); j++ {
			k := (i*int(header.Width) + j) * 3
			image[i][j] = Pixel{R: data[k+2], G: data[k+1], B: data[k]}
		}
	}

	return header, image, nil
}

// SaveBMP сохраняет BMP файл
func SaveBMP(filename string, header BMPHeader, image [][]Pixel) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Update header dimensions after cropping
	if len(image) == 0 {
		fmt.Println("height or width equal to 0")
		os.Exit(1)
	}
	header.Width = uint32(len(image[0])) // New width after cropping
	header.Height = uint32(len(image))   // New height after cropping

	// Each row must be a multiple of 4 bytes, so we calculate the necessary padding
	rowSize := int(header.Width) * 3
	padding := (4 - (rowSize % 4)) % 4 // BMP rows are padded to multiples of 4 bytes

	// Update ImageSize in header (height * (rowSize + padding))
	header.ImageSize = uint32((rowSize + padding) * int(header.Height))

	// Write BMP header
	err = binary.Write(file, binary.LittleEndian, header)
	if err != nil {
		return err
	}

	// Write pixel data with padding
	for i := 0; i < int(header.Height); i++ {
		for j := 0; j < int(header.Width); j++ {
			// Write pixel data directly without using k
			_, err = file.Write([]byte{image[i][j].B, image[i][j].G, image[i][j].R})
			if err != nil {
				return err
			}
		}

		// Write padding bytes if necessary
		if padding > 0 {
			_, err = file.Write(make([]byte, padding))
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// printHeaderInfo выводит информацию о заголовке BMP
func PrintHeaderInfo(header BMPHeader) {
	fmt.Println("BMP Header:")
	fmt.Printf("- FileType: %s\n", string(header.FileType[:])) // 1
	fmt.Printf("- FileSizeInBytes: %d\n", header.FileSize)     // 2
	fmt.Printf("- HeaderSize: %d\n", 14+header.DIBHeaderSize)  // 3
	fmt.Println("DIB Header:")
	fmt.Printf("- DibHeaderSize: %d\n", header.DIBHeaderSize) // 4
	fmt.Printf("- WidthInPixels: %d\n", header.Width)         // 5
	fmt.Printf("- HeightInPixels: %d\n", header.Height)       // 6
	fmt.Printf("- PixelSizeInBits: %d\n", header.BitCount)    // 7
	fmt.Printf("- ImageSizeInBytes: %d\n", header.ImageSize)  // 8
}

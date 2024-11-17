package general

import (
	"fmt"
	"os"
)

// ApplyBlueFilter применяет синий фильтр к изображению
func ApplyBlueFilter(image [][]Pixel) [][]Pixel {
	for i := range image {
		for j := range image[i] {
			// Убираем красный и зеленый каналы, оставляя только синий
			image[i][j].R = 0
			image[i][j].G = 0
		}
	}
	return image
}

// ApplyRedFilter применяет красный фильтр к изображению
func ApplyRedFilter(image [][]Pixel) [][]Pixel {
	for i := range image {
		for j := range image[i] {
			// Убираем зеленый и синий каналы, оставляя только красный
			image[i][j].G = 0
			image[i][j].B = 0
		}
	}
	return image
}

// ApplyGreenFilter применяет зеленый фильтр к изображению
func ApplyGreenFilter(image [][]Pixel) [][]Pixel {
	for i := range image {
		for j := range image[i] {
			// Убираем красный и синий каналы, оставляя только зеленый
			image[i][j].R = 0
			image[i][j].B = 0
		}
	}
	return image
}

// ApplyNegativeFilter применяет негатив к изображению
func ApplyNegativeFilter(image [][]Pixel) [][]Pixel {
	for i := range image {
		for j := range image[i] {
			// Инвертируем цвета
			image[i][j].R = 255 - image[i][j].R
			image[i][j].G = 255 - image[i][j].G
			image[i][j].B = 255 - image[i][j].B
		}
	}
	return image
}

func ApplyGrayscaleFilter(image [][]Pixel) [][]Pixel {
	for i := range image {
		for j := range image[i] {
			// Calculate grayscale value
			gray := uint8(0.3*float64(image[i][j].R) + 0.59*float64(image[i][j].G) + 0.11*float64(image[i][j].B))
			image[i][j].R = gray
			image[i][j].G = gray
			image[i][j].B = gray
		}
	}
	return image
}

// ApplyPixelateFilter применяет пикселизацию к изображению
func ApplyPixelateFilter(image [][]Pixel) [][]Pixel {
	pixelationSize := 20 // размер пикселя
	for i := 0; i < len(image); i += pixelationSize {
		for j := 0; j < len(image[i]); j += pixelationSize {
			// Вычисляем средний цвет для блока пикселей
			avgR, avgG, avgB := uint32(0), uint32(0), uint32(0)
			count := 0

			for x := 0; x < pixelationSize && i+x < len(image); x++ {
				for y := 0; y < pixelationSize && j+y < len(image[i]); y++ {
					avgR += uint32(image[i+x][j+y].R)
					avgG += uint32(image[i+x][j+y].G)
					avgB += uint32(image[i+x][j+y].B)
					count++
				}
			}

			avgR /= uint32(count)
			avgG /= uint32(count)
			avgB /= uint32(count)

			// Устанавливаем средний цвет для блока пикселей
			for x := 0; x < pixelationSize && i+x < len(image); x++ {
				for y := 0; y < pixelationSize && j+y < len(image[i]); y++ {
					image[i+x][j+y].R = uint8(avgR)
					image[i+x][j+y].G = uint8(avgG)
					image[i+x][j+y].B = uint8(avgB)
				}
			}
		}
	}
	return image
}

// ApplyBlurFilter применяет размытие к изображению
func ApplyBlurFilter(image [][]Pixel) [][]Pixel {
	height := len(image)
	width := len(image[0])
	radius := 10

	// Создаем копию изображения для хранения результата
	newImage := make([][]Pixel, height)
	for i := range newImage {
		newImage[i] = make([]Pixel, width)
	}

	// Проходим по каждому пикселю изображения
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			var sumR, sumG, sumB, count uint32

			// Проходим по соседним пикселям в пределах радиуса
			for y := i - radius; y <= i+radius; y++ {
				for x := j - radius; x <= j+radius; x++ {
					// Проверяем, что координаты соседнего пикселя внутри изображения
					if y >= 0 && y < height && x >= 0 && x < width {
						sumR += uint32(image[y][x].R)
						sumG += uint32(image[y][x].G)
						sumB += uint32(image[y][x].B)
						count++
					}
				}
			}

			// Усредняем значения цвета и присваиваем пикселю
			newImage[i][j] = Pixel{
				R: uint8(sumR / count),
				G: uint8(sumG / count),
				B: uint8(sumB / count),
			}
		}
	}

	return newImage
}

// ApplyHorizontalMirrorFilter применяет горизонтальное зеркалирование
func ApplyHorizontalMirrorFilter(image [][]Pixel) [][]Pixel {
	for i := range image {
		for j := 0; j < len(image[i])/2; j++ {
			// Меняем местами пиксели с двух сторон
			image[i][j], image[i][len(image[i])-1-j] = image[i][len(image[i])-1-j], image[i][j]
		}
	}
	return image
}

// ApplyVerticalMirrorFilter применяет вертикальное зеркалирование
func ApplyVerticalMirrorFilter(image [][]Pixel) [][]Pixel {
	// Для вертикального зеркалирования меняем местами строки
	for i := 0; i < len(image)/2; i++ {
		image[i], image[len(image)-1-i] = image[len(image)-1-i], image[i]
	}
	return image
}

func ApplyRotateLeftFilter(image [][]Pixel) [][]Pixel {
	height := len(image)
	width := len(image[0])

	// Create a new image with dimensions swapped
	newImage := make([][]Pixel, width)
	for i := range newImage {
		newImage[i] = make([]Pixel, height)
	}

	// Rotate the image 90 degrees clockwise
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			newImage[j][height-1-i] = image[i][j]
		}
	}

	return newImage
}

func ApplyRotateRightFilter(image [][]Pixel) [][]Pixel {
	height := len(image)
	width := len(image[0])

	// Create a new image with dimensions swapped
	newImage := make([][]Pixel, width)
	for i := range newImage {
		newImage[i] = make([]Pixel, height)
	}

	// Rotate the image 90 degrees counterclockwise
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			newImage[width-1-j][i] = image[i][j]
		}
	}

	return newImage
}

func CropImage(img [][]Pixel, x, y, width, height int) [][]Pixel {
	if x == y && y == 0 {
		if height > len(img) || width > len(img[0]) {
			fmt.Println("ERROR crop")
			os.Exit(1)
		}
		height = len(img) - height
		img = img[:height]
		for i := range img {
			img[i] = img[i][width:]
		}
	} else {
		if x < 0 || y < 0 || x+width > len(img[0]) || y+height > len(img) {
			fmt.Println("ERROR croped")
			os.Exit(1)
		}
		img = img[len(img)-y-height : len(img)-y] // выбираем строки сверху

		for i := range img {
			// Обрезаем столбцы
			img[i] = img[i][x : x+width]
		}
	}
	return img
	// imgheight := len(img)
	// if imgheight == 0 {
	// 	return nil
	// }
	// imgwidth := len(img[0])
	// croppedImage := make([][]Pixel, height)
	// for i := range croppedImage {
	// 	croppedImage[i] = make([]Pixel, width)
	// }
	// for i := 0; i < height; i++ {
	// 	for j := 0; j < width; j++ {
	// 		croppedImage[i][j] = img[x+i][y+j]
	// 	}
	// }
	// return croppedImage
}

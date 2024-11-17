package main

import (
	"bitmap/general"
	"flag"
	"fmt"
	"os"
)

func main() {
	general.PrintProgrammHelp()

	switch os.Args[1] {
	case "apply":
		// Определение флагов для команды apply
		applyCmd := flag.NewFlagSet("apply", flag.ExitOnError)

		var filters []string
		applyCmd.Var((*flagList)(&filters), "filter", "Filters to apply (e.g. blue, red, green, negative, pixelate, blur) (can be used multiple times)")

		var mirrors []string
		applyCmd.Var((*flagList)(&mirrors), "mirror", "Mirror direction (horizontal, vertical) (can be used multiple times)")

		var rotations []string
		applyCmd.Var((*flagList)(&rotations), "rotate", "Rotate direction (right, 90, 180, 270, left, -90, -180, -270) (can be used multiple times)")

		var crops []string
		applyCmd.Var((*flagList)(&crops), "crop", "Crop parameters (e.g. 100-100 or 20-20-100-100)")

		applyHelp := applyCmd.Bool("help", false, "Show header Usage")
		applyCmd.BoolVar(applyHelp, "h", false, "\nShow header Usage")

		general.PrintApplyHelp()
		if err := applyCmd.Parse(os.Args[2:]); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		// Получение аргументов после флагов
		args := applyCmd.Args()

		if len(args) < 2 {
			fmt.Println("not enough arguments")
			os.Exit(1)
		}

		inputFile := args[0]
		outputFile := args[1]
		// Загрузка BMP файла
		if len(inputFile) > 4 && inputFile[len(inputFile)-4:] != ".bmp" {
			fmt.Println("ERROR: input file not bmp")
			os.Exit(1)
		}
		if len(outputFile) > 4 && outputFile[len(outputFile)-4:] != ".bmp" {
			fmt.Println("ERROR: output file not bmp")
			os.Exit(1)
		}
		header, image, err := general.LoadBMP(inputFile)
		if err != nil {
			fmt.Printf("Error loading BMP file: %v\n", err)
			os.Exit(1)
		}

		// Применение crop
		for _, crop := range crops {
			x, y, width, height, err := general.ParseCropFlag(crop)
			if err != nil {
				fmt.Println("Invalid crop parameters:", err)
				os.Exit(1)
			}
			image = general.CropImage(image, x, y, width, height)
			header.Width = uint32(width)
			header.Height = uint32(height)
		}

		// Применение фильтров и зеркалирования
		for _, filter := range filters {
			switch filter {
			case "grayscale":
				image = general.ApplyGrayscaleFilter(image)
			case "blue":
				image = general.ApplyBlueFilter(image)
			case "red":
				image = general.ApplyRedFilter(image)
			case "green":
				image = general.ApplyGreenFilter(image)
			case "negative":
				image = general.ApplyNegativeFilter(image)
			case "pixelate":
				image = general.ApplyPixelateFilter(image)
			case "blur":
				image = general.ApplyBlurFilter(image)
			default:
				fmt.Printf("Unknown filter: %s\n", filter)
				os.Exit(1)
			}
		}

		// Применение зеркалирования в цикле
		for _, mirror := range mirrors {
			switch mirror {
			case "horizontal":
				image = general.ApplyHorizontalMirrorFilter(image)
			case "hor":
				image = general.ApplyHorizontalMirrorFilter(image)
			case "h":
				image = general.ApplyHorizontalMirrorFilter(image)
			case "vertical":
				image = general.ApplyVerticalMirrorFilter(image)
			case "ver":
				image = general.ApplyVerticalMirrorFilter(image)
			case "v":
				image = general.ApplyVerticalMirrorFilter(image)
			default:
				fmt.Printf("Unknown mirror option: %s\n", mirror)
				os.Exit(1)
			}
		}

		// Need to update dimensions after rotation (header.Width, header.Height = header.Height, header.Width)
		for _, r := range rotations {
			switch r {
			case "right":
				image = general.ApplyRotateRightFilter(image)
				header.Width, header.Height = header.Height, header.Width
			case "left":
				image = general.ApplyRotateLeftFilter(image)
				header.Width, header.Height = header.Height, header.Width
			case "90":
				image = general.ApplyRotateRightFilter(image)
				header.Width, header.Height = header.Height, header.Width
			case "180":
				image = general.ApplyRotateRightFilter(image)
				image = general.ApplyRotateRightFilter(image)
			case "270":
				image = general.ApplyRotateRightFilter(image)
				image = general.ApplyRotateRightFilter(image)
				image = general.ApplyRotateRightFilter(image)
				header.Width, header.Height = header.Height, header.Width
			case "-90":
				image = general.ApplyRotateLeftFilter(image)
			case "-180":
				image = general.ApplyRotateLeftFilter(image)
				image = general.ApplyRotateLeftFilter(image)
			case "-270":
				image = general.ApplyRotateLeftFilter(image)
				image = general.ApplyRotateLeftFilter(image)
				image = general.ApplyRotateLeftFilter(image)
				header.Width, header.Height = header.Height, header.Width
			default:
				fmt.Printf("Unknown rotation option: %s\n", r)
				os.Exit(1)
			}
		}

		// Сохранение результата
		err = general.SaveBMP(outputFile, header, image)
		if err != nil {
			fmt.Printf("Error saving BMP file: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Successfully applied filters and saved the output.")

	case "header":
		headerCmd := flag.NewFlagSet("header", flag.ExitOnError)
		headerHelp := headerCmd.Bool("help", false, "Show header Usage")
		headerCmd.BoolVar(headerHelp, "h", false, "\nShow header Usage")

		if err := headerCmd.Parse(os.Args[2:]); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		general.PrintHeaderHelp()

		inputFile := os.Args[2]
		header, err := general.LoadHeader(inputFile)
		if err != nil {
			fmt.Printf("Error loading BMP file: %v\n", err)
			os.Exit(1)
		}
		general.PrintHeaderInfo(header)

	default:
		fmt.Println("unknown command\nUse:\n   ./bitmap   prints program usage information")
		os.Exit(1)
	}
}

// Мульти флаги(подсчет флагов)
type flagList []string

func (f *flagList) String() string {
	return fmt.Sprint(*f)
}

func (f *flagList) Set(value string) error {
	*f = append(*f, value)
	return nil
}

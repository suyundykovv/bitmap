package general

import (
	"errors"
	"flag"
	"fmt"
	"strconv"
	"strings"
)

// Определение глобальных флагов
var (
	FilterFlag = flag.String("filter", "", "Filter to apply: grayscale, red, green, blue, negative, pixelate, blur")
	MirrorFlag = flag.String("mirror", "", "Mirror direction: horizontal, vertical")
	RotateFlag = flag.String("rotate", "", "Rotate direction: right, left")
)

func ParseFlags() {
	flag.Parse()
}

func ParseCropFlag(cropArg string) (int, int, int, int, error) {
	parts := strings.Split(cropArg, "-")

	if len(parts) != 2 && len(parts) != 4 {
		return 0, 0, 0, 0, errors.New("invalid number of crop parameters, expected 2 (width-height) or 4 (x-y-width-height)")
	}

	values := make([]int, len(parts))
	for i, part := range parts {
		v, err := strconv.Atoi(part)
		if err != nil {
			return 0, 0, 0, 0, fmt.Errorf("invalid crop parameter: %s", part)
		}
		values[i] = v
	}

	if len(values) == 2 {
		return 0, 0, values[0], values[1], nil // x=0, y=0, width, height
	}
	return values[0], values[1], values[2], values[3], nil // x, y, width, height
}

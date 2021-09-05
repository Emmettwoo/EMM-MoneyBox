package util

import "strconv"

func String2Float32(number string) float32 {
	if numberInFloat32, err := strconv.ParseFloat(number, 32); err == nil {
		return float32(numberInFloat32)
	} else {
		return 0
	}
}

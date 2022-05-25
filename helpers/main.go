package helpers

import (
	"math"
	"strconv"
)

func GrayToDec(seq string, a float64, b float64) float64 {
	n := grayToInt(seq)
	m := float64(len(seq))
	return a + float64(n)*(b-a)/(math.Pow(2, m)-1)
}

func grayToInt(seq string) int {
	n, err := strconv.ParseInt(seq, 2, 64)
	if err != nil {
		panic(err)
	}
	mask := n
	for mask != 0 {
		mask >>= 1
		n ^= mask
	}
	return int(n)
}

func DecToGray(x float64, a float64, b float64) string {
	m := 10.0
	n := int((x - a) * (math.Pow(2, m) - 1) / (b - a))
	return binaryToGray(n)
}

func binaryToGray(n int) string {
	n ^= n >> 1
	return strconv.FormatInt(int64(n), 2)
}

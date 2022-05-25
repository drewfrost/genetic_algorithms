package healthfunctions

import (
	"genetic_algorithms/helpers"
	"math"
	"strings"
)

func FConst(seq string) float64 {
	if seq == strings.Repeat("0", len(seq)) || seq == strings.Repeat("1", len(seq)) {
		return float64(len(seq))
	}
	return 0
}

func Fh(seq string) float64 {
	count := 0.0
	for i := 0; i < len(seq); i++ {
		if seq[i] == '0' {
			count++
		}
	}
	return count
}

func fhd(seq string, q int) float64 {
	l := len(seq)
	var k int
	for i := 0; i < l; i++ {
		if seq[i] == '0' {
			k++
		}
	}
	return float64((l - k) + k*q)

}

func Fhd10(seq string) float64 {
	return fhd(seq, 10)
}

func Fhd50(seq string) float64 {
	return fhd(seq, 50)
}

func Fhd150(seq string) float64 {
	return fhd(seq, 150)
}

func GrayX(seq string) float64 {
	return helpers.GrayToDec(seq, 0.0, 10.23)
}

func GrayX2(seq string) float64 {
	x := helpers.GrayToDec(seq, 0.0, 10.23)
	return math.Pow(x, 2.0)
}

func GrayX4(seq string) float64 {
	x := GrayX(seq)
	return math.Pow(x, 4.0)
}

func Gray2X2(seq string) float64 {
	x := GrayX(seq)
	return 2 * math.Pow(x, 2.0)
}

func Gray512X2(seq string) float64 {
	x := helpers.GrayToDec(seq, -5.11, 5.12)
	return math.Pow(5.12, 2.0) - math.Pow(x, 2.0)
}

func Gray512X4(seq string) float64 {
	x := GrayX(seq)
	return math.Pow(5.12, 4.0) - math.Pow(x, 4.0)
}

func grayEXC(seq string, c float64) float64 {
	x := helpers.GrayToDec(seq, 0, 10.23)
	return math.Exp(c * x)
}

func GrayEX025(seq string) float64 {
	return grayEXC(seq, 0.25)
}

func GrayEX1(seq string) float64 {
	return grayEXC(seq, 1.0)
}

func GrayEX2(seq string) float64 {
	return grayEXC(seq, 2.0)
}

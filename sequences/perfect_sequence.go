package sequences

import (
	"genetic_algorithms/helpers"
	"strings"
)

func FConstPerfectSequence(seq string) bool {
	if seq == strings.Repeat("0", len(seq)) || seq == strings.Repeat("1", len(seq)) {
		return true
	}
	return false
}

func FhPerfectSequence(seq string) bool {
	return seq == strings.Repeat("0", len(seq))
}

func PerfectSequence0(seq string) bool {
	x := helpers.GrayToDec(seq, -5.11, 5.12)
	if x >= -0.02 && x <= 0.02 {
		return true
	}
	return false
}

func PerfectSequence1023(seq string) bool {
	x := helpers.GrayToDec(seq, 0.0, 10.23)
	if x >= 10.21 && x <= 10.25 {
		return true
	}
	return false
}

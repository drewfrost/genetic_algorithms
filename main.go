package main

import (
	"fmt"
	"genetic_algorithms/healthfunctions"
	"genetic_algorithms/helpers"
	"genetic_algorithms/sequences"
)

type param struct {
	populationGenerator sequences.Generator
	healthFunction      func(string) float64
	perfectItemFunc     func(string) bool
	calcNoise           bool
}

type Dictionary map[string]interface{}

func main() {
	filename := "report.csv"
	populationSize := 491
	noMutationTests(filename, populationSize)
	// healthfunctions.GrayX2("0001000000")
}

func noMutationTests(filename string, populationSize int) {
	params := []param{}
	params = append(params, param{
		populationGenerator: sequences.NewFConstPopulationGenerator(100),
		healthFunction:      healthfunctions.FConst,
		perfectItemFunc:     sequences.FConstPerfectSequence,
		calcNoise:           true,
	},
		param{
			populationGenerator: sequences.NewBinomialPopulationGenerator(100, ""),
			healthFunction:      healthfunctions.Fh,
			perfectItemFunc:     sequences.FhPerfectSequence,
			calcNoise:           false,
		},
		param{
			populationGenerator: sequences.NewBinomialPopulationGenerator(100, ""),
			healthFunction:      healthfunctions.Fhd10,
			perfectItemFunc:     sequences.FhPerfectSequence,
			calcNoise:           false,
		},
		param{
			populationGenerator: sequences.NewBinomialPopulationGenerator(100, ""),
			healthFunction:      healthfunctions.Fhd50,
			perfectItemFunc:     sequences.FhPerfectSequence,
			calcNoise:           false,
		},
		param{
			populationGenerator: sequences.NewBinomialPopulationGenerator(100, ""),
			healthFunction:      healthfunctions.Fhd150,
			perfectItemFunc:     sequences.FhPerfectSequence,
			calcNoise:           false,
		},
		param{
			populationGenerator: sequences.NewBinomialPopulationGenerator(10, helpers.DecToGray(10.23, 0, 10.23)),
			healthFunction:      healthfunctions.GrayX2,
			perfectItemFunc:     sequences.PerfectSequence1023,
			calcNoise:           false,
		},
		param{
			populationGenerator: sequences.NewBinomialPopulationGenerator(10, helpers.DecToGray(10.23, 0, 10.23)),
			healthFunction:      healthfunctions.GrayX,
			perfectItemFunc:     sequences.PerfectSequence1023,
			calcNoise:           false,
		},
		param{
			populationGenerator: sequences.NewBinomialPopulationGenerator(10, helpers.DecToGray(10.23, 0, 10.23)),
			healthFunction:      healthfunctions.GrayX4,
			perfectItemFunc:     sequences.PerfectSequence1023,
			calcNoise:           false,
		},
		param{
			populationGenerator: sequences.NewBinomialPopulationGenerator(10, helpers.DecToGray(10.23, 0, 10.23)),
			healthFunction:      healthfunctions.Gray2X2,
			perfectItemFunc:     sequences.PerfectSequence1023,
			calcNoise:           false,
		},
		param{
			populationGenerator: sequences.NewBinomialPopulationGenerator(10, helpers.DecToGray(0, -5.11, 5.12)),
			healthFunction:      healthfunctions.Gray512X2,
			perfectItemFunc:     sequences.PerfectSequence0,
			calcNoise:           false,
		},
		param{
			populationGenerator: sequences.NewBinomialPopulationGenerator(10, helpers.DecToGray(0, -5.11, 5.12)),
			healthFunction:      healthfunctions.Gray512X4,
			perfectItemFunc:     sequences.PerfectSequence0,
			calcNoise:           false,
		},
		param{
			populationGenerator: sequences.NewBinomialPopulationGenerator(10, helpers.DecToGray(10.23, 0, 10.23)),
			healthFunction:      healthfunctions.GrayEX025,
			perfectItemFunc:     sequences.PerfectSequence1023,
			calcNoise:           false,
		},
		param{
			populationGenerator: sequences.NewBinomialPopulationGenerator(10, helpers.DecToGray(10.23, 0, 10.23)),
			healthFunction:      healthfunctions.GrayEX1,
			perfectItemFunc:     sequences.PerfectSequence1023,
			calcNoise:           false,
		},
		param{
			populationGenerator: sequences.NewBinomialPopulationGenerator(10, helpers.DecToGray(10.23, 0, 10.23)),
			healthFunction:      healthfunctions.GrayEX2,
			perfectItemFunc:     sequences.PerfectSequence1023,
			calcNoise:           false,
		},
	)
	fmt.Println(params)

}

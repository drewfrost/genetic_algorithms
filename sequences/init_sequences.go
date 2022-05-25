package sequences

import (
	"math/rand"
	"strings"
)

type Generator interface {
	Generate(populationSize int, addPerfect bool) []string
}

type FConstPopulationGenerator struct {
	Generator
	SequenceLenght int
}

func (fcpg *FConstPopulationGenerator) GeneratePopulation(populationSize int) []string {
	return fcpg.Generate(populationSize, false)
}

func (fcpg *FConstPopulationGenerator) Generate(populationSize int, addPerfect bool) []string {
	var population []string
	for i := 0; i < populationSize/2; i++ {
		population = append(population, strings.Repeat("0", fcpg.SequenceLenght))
		population = append(population, strings.Repeat("1", fcpg.SequenceLenght))
	}
	if populationSize%2 == 1 {
		population = append(population, strings.Repeat("1", fcpg.SequenceLenght))
	}
	return population
}

func (fcpg *FConstPopulationGenerator) Init(sequenceLenght int) {
	fcpg.SequenceLenght = sequenceLenght
}

func NewFConstPopulationGenerator(sequenceLenght int) *FConstPopulationGenerator {
	fcpg := new(FConstPopulationGenerator)
	fcpg.Init(sequenceLenght)
	return fcpg
}

type BinomialPopulationGenerator struct {
	Generator
	SequenceLenght int
	perfectItem    string
}

func (bpg *BinomialPopulationGenerator) Init(sequenceLenght int, perfectItem string) {
	bpg.SequenceLenght = sequenceLenght
	if perfectItem == "" {
		bpg.perfectItem = bpg.generateOptiomalSequence()
	} else {
		bpg.perfectItem = perfectItem
	}
}

func NewBinomialPopulationGenerator(sequenceLenght int, perfectItem string) *BinomialPopulationGenerator {
	bpg := new(BinomialPopulationGenerator)
	bpg.Init(sequenceLenght, perfectItem)
	return bpg
}

func (bpg *BinomialPopulationGenerator) Generate(populationSize int, addPerfect bool) []string {
	var population []string
	if addPerfect {
		population = append(population, bpg.generateOptiomalSequence())
	}
	for len(population) < populationSize {
		population = append(population, bpg.generateDefaultSequence())
	}
	return population
}

func (bpg *BinomialPopulationGenerator) generateOptiomalSequence() string {
	return strings.Repeat("0", bpg.SequenceLenght)
}

func (bpg *BinomialPopulationGenerator) generateDefaultSequence() string {
	var sequence string
	for i := 0; i < bpg.SequenceLenght; i++ {
		if rand.Float64() < 0.5 {
			sequence += "0"
		} else {
			sequence += "1"
		}
	}
	return sequence
}

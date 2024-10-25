package models

type Chromosome []float64

type Individual struct {
	X Chromosome
	J float64
}

type Population []Individual

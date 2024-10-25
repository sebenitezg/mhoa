package models

type ObjFunc func(Chromosome) float64

type ProblemParams struct {
	Xmin    Chromosome
	Xmax    Chromosome
	NP      int
	ObjFunc ObjFunc
}

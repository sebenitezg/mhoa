package main

import (
	"github.com/sebenitezg/mhoa/internal/algorithms/differentialevolution"
	"github.com/sebenitezg/mhoa/internal/models"
	"github.com/sebenitezg/mhoa/internal/problems"
)

func main() {

	problemParams := models.ProblemParams{
		Xmin:    models.Chromosome{-10.0, -10.0},
		Xmax:    models.Chromosome{10.0, 10.0},
		NP:      100,
		ObjFunc: problems.BoothFunc,
	}
	deParams := models.DEParams{
		Gmax: 10000,
		Fmin: 0.3,
		Fmax: 0.9,
		CR:   0.5,
		Xi:   0.0001,
	}

	alg, err := differentialevolution.NewDifferentialEvolution(problemParams, deParams)
	if err != nil {
		panic(err)
	}
	alg.Execute()
}

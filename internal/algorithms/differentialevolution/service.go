package differentialevolution

import (
	"fmt"
	"math/rand"
	"reflect"

	"github.com/sebenitezg/mhoa/internal/models"
)

// TODO: Use Generics instead of interface
type DifferentialEvolution struct {
	Population models.Population
	Xmin       models.Chromosome
	Xmax       models.Chromosome
	NP         int
	DX         int
	ObjFunc    models.ObjFunc
	Gmax       int
	Fmin       float64
	Fmax       float64
	CR         float64
}

func NewDifferentialEvolution(problemParams models.ProblemParams, deParams models.DEParams) (*DifferentialEvolution, error) {
	if len(problemParams.Xmin) != len(problemParams.Xmax) {
		return nil, fmt.Errorf("the minimum and maximum chromosomes sizes are different")
	}
	for i := 0; i < len(problemParams.Xmin); i++ {
		if reflect.TypeOf(problemParams.Xmin[i]) != reflect.TypeOf(problemParams.Xmax[i]) {
			return nil, fmt.Errorf("the %dth gene of the chromosomes are different", i)
		}
	}
	return &DifferentialEvolution{
		Population: make(models.Population, problemParams.NP),
		Xmin:       problemParams.Xmin,
		Xmax:       problemParams.Xmax,
		NP:         problemParams.NP,
		DX:         len(problemParams.Xmin),
		ObjFunc:    problemParams.ObjFunc,
		Gmax:       deParams.Gmax,
		Fmin:       deParams.Fmin,
		Fmax:       deParams.Fmax,
		CR:         deParams.CR,
	}, nil
}

func (de *DifferentialEvolution) Execute() (*models.Population, error) {
	// Initialize the population
	err := de.InitializePopulation()
	if err != nil {
		return nil, err
	}

	minIndividual := de.GetMinIndividual()
	de.PrintIndividual(minIndividual)
	for i := 0; i < de.Gmax; i++ {
		de.Rand1Bin()
		fmt.Print("Generation: ", i)
		fmt.Println()
		minIndividual := de.GetMinIndividual()
		de.PrintIndividual(minIndividual)
	}

	return nil, nil
}

func (de *DifferentialEvolution) InitializePopulation() error {
	xMin := de.Xmin
	xMax := de.Xmax
	for i := 0; i < de.NP; i++ {
		x := make(models.Chromosome, de.DX)
		for j := 0; j < de.DX; j++ {
			x[j] = xMin[j] + rand.Float64()*(xMax[j]-xMin[j])
		}
		fitness := de.ObjFunc(x)
		de.Population[i] = models.Individual{
			X: x,
			J: fitness,
		}
	}
	return nil
}

func (de *DifferentialEvolution) Rand1Bin() error {
	newPop := make(models.Population, de.NP)
	for i, parent := range de.Population {
		// TODO: Research how the scale factor is used with integer values
		F := de.Fmin + rand.Float64()*(de.Fmax-de.Fmin)

		// Select individuals randomly for crossover and mutation
		rIndiv := de.Selection(de.Population, i)
		jRand := rand.Intn(de.DX)

		// Apply crossover and mutation
		child := models.Individual{
			X: make(models.Chromosome, de.DX),
		}
		for j := 0; j < de.DX; j++ {
			if rand.Float64() < de.CR || j == jRand {
				child.X[j] = rIndiv[0].X[j] + F*(rIndiv[1].X[j]-rIndiv[2].X[j])
				if child.X[j] < de.Xmin[j] || child.X[j] > de.Xmax[j] {
					child.X[j] = de.Xmin[j] + rand.Float64()*(de.Xmax[j]-de.Xmin[j])
				}
			} else {
				child.X[j] = parent.X[j]
			}
		}

		// Evaluate objective function
		child.J = de.ObjFunc(child.X)

		// Replacement
		newPop[i] = de.Replacement(parent, child)
	}

	de.Population = newPop

	return nil
}

func (de *DifferentialEvolution) Crossover(individuals []models.Individual) {

}

func (de *DifferentialEvolution) Mutation(individuals []models.Individual) {

}

func (de *DifferentialEvolution) Replacement(parent, child models.Individual) models.Individual {
	// Selection
	if child.J < parent.J {
		return child
	}
	return parent
}

func (de *DifferentialEvolution) Selection(pop models.Population, excludeIndex int) []models.Individual {
	r := rand.Perm(de.NP)[0:4]
	idx := 3
	for i, num := range r {
		if num == excludeIndex {
			idx = i
		}
	}
	r = append(r[:idx], r[idx+1:]...)

	return []models.Individual{pop[r[0]], pop[r[1]], pop[r[2]]}
}

func (de *DifferentialEvolution) GetMinIndividual() models.Individual {
	minIndividual := de.Population[0]
	for _, individual := range de.Population {
		if individual.J < minIndividual.J {
			minIndividual = individual
		}

	}
	return minIndividual
}

func (de *DifferentialEvolution) GetMaxIndividual() models.Individual {
	maxIndividual := de.Population[0]
	for _, individual := range de.Population {
		if individual.J > maxIndividual.J {
			maxIndividual = individual
		}

	}
	return maxIndividual
}

func (de *DifferentialEvolution) PrintIndividual(individual models.Individual) {
	fmt.Print(individual.X)
	fmt.Print("\t ", individual.J)
	fmt.Println()
}

func (de *DifferentialEvolution) PrintPopulation() {
	for _, individual := range de.Population {
		de.PrintIndividual(individual)
	}
}

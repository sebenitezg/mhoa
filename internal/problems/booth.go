package problems

import (
	"math"

	"github.com/sebenitezg/mhoa/internal/models"
)

func BoothFunc(c models.Chromosome) float64 {
	x := c[0]
	y := c[1]

	a := x + 2.0*y - 7.0
	b := 2.0*x + y - 5.0
	return math.Pow(a, 2.0) + math.Pow(b, 2.0)
}

package tsp

import (
	"github.com/ybping/metaheuristics/ga"
	"math"
	"math/rand"
)

// Species info
type Species struct {
	cityOrder []int
	cityCount int
	score     float64
	tsp       *TSP
}

func newSpecies(cityCount int) *Species {
	// 随机生成一个初始解
	var cityOrder []int
	for i := 0; i < cityCount; i++ {
		cityOrder = append(cityOrder, rand.Intn(cityCount))
	}
	return &Species{
		cityOrder: cityOrder,
		cityCount: cityCount,
		score:     -1,
	}

}

func (s *Species) cross(tsp Species) Species {
	return Species{}
}

func (s *Species) mutate() Species {
	first, second := rand.Intn(s.cityCount), rand.Intn(s.cityCount)
	s.cityOrder[first], s.cityOrder[second] = s.cityOrder[second], s.cityOrder[first]
	return s
}

func (s *Species) fitness() float64 {
	if s.score == -1 {
		s.score = s.tsp.fitness()
	}
	return s.score
}

// City info
type City struct {
	Name string
	Lng  float64
	Lat  float64
}

// TSP info
type TSP struct {
	cities []City
	ga     *GeneticAlgorithm
}

// NewTSP return tsp
func NewTSP(cities []City) TSP {
	tsp := &TSP{}

	// init specices
	species := newSpecies(len(cities))
	species.tsp = tsp

	// init tsp
	tsp.species = species
	tsp.ga = ga.NewGeneticAlgorithm(species, 100, 0.7, 0.02)
}

func (tsp TSP) distance() float64 {
	distance := 0.0
	cityCount := len(tsp.cities)
	calculate := func(from, to City) (distance float64) {
		dlng := from.Lng - to.Lng
		dlat := from.Lat - to.Lat
		return math.Sqrt(math.Pow(dlng, 2.0) + math.Pow(dlat, 2.0))
	}
	for k := range tsp.ga.cityOrder {
		from := tsp.cities[k]
		to := tsp.cities[(k+1)%cityCount]
		distance += calculate(from, to)
	}
	return distance
}

func (tsp TSP) fitness() float64 {
	return 1.0 / tsp.distance()
}

// Solve start to solve tsp
func (tsp TSP) Solve() {
	tsp.ga.Evoluation()
	log.Println("The distance is %f", tsp.distance())
}

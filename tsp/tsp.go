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

func newSpecies(cityCount int) Species {
	// 随机生成一个初始解
	var cityOrder []int
	for i := 0; i < cityCount; i++ {
		cityOrder = append(cityOrder, rand.Intn(cityCount))
	}
	return Species{
		cityOrder: cityOrder,
		cityCount: cityCount,
		score:     -1,
	}

}

func (s *Species) cross(t Species) Species {
	leftPos := rand.Intn(s.cityCount)
	rightPos := rand.Intn(s.cityCount)
	if leftPos > rightPos {
		leftPos, rightPos = rightPos, leftPos
	}
	cityOrder := s.cityOrder[0:leftPos]
	cityOrder = append(cityOrder, t.cityOrder[leftPos:rightPos]...)
	cityOrder = append(cityOrder, s.cityOrder[rightPos:s.cityCount]...)
	existed := make(map[int]bool)
	for _, v := range cityOrder[leftPos:rightPos] {
		existed[v] = true
	}
	for k, v := range cityOrder {
		if existed[v] == true {
			if k < leftPos || k >= rightPos {
				for _, v1 := range t.cityOrder {
					if existed[v1] == false {
						cityOrder[k] = v1
						existed[v1] = true
					}
				}
			}
		}
	}
	return Species{
		cityOrder: cityOrder,
		cityCount: s.cityCount,
		score:     -1,
		tsp:       s.tsp,
	}
}

func (s *Species) mutate() {
	first, second := rand.Intn(s.cityCount), rand.Intn(s.cityCount)
	s.cityOrder[first], s.cityOrder[second] = s.cityOrder[second], s.cityOrder[first]
}

func (s *Species) fitness() float64 {
	if s.score == -1 {
		s.score = s.tsp.fitness(*s)
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
	ga     *ga.GeneticAlgorithm
}

// NewTSP return tsp
func NewTSP(cities []City) *TSP {
	tsp := &TSP{}

	// init specices
	var population []ga.Species
	for i := 0; i < 100; i++ {
		species := newSpecies(len(cities))
		species.tsp = tsp
		population = append(population, species)
	}

	// init tsp
	tsp.ga = ga.NewGeneticAlgorithm(100, 0.7, 0.02, population)
	return tsp
}

func (tsp TSP) distance(s Species) float64 {
	distance := 0.0
	cityCount := len(tsp.cities)
	calculate := func(from, to City) (distance float64) {
		dlng := from.Lng - to.Lng
		dlat := from.Lat - to.Lat
		return math.Sqrt(math.Pow(dlng, 2.0) + math.Pow(dlat, 2.0))
	}
	for k := range s.cityOrder {
		from := tsp.cities[k]
		to := tsp.cities[(k+1)%cityCount]
		distance += calculate(from, to)
	}
	return distance
}

func (tsp TSP) fitness(s Species) float64 {
	return 1.0 / tsp.distance(s)
}

// Solve start to solve tsp
func (tsp TSP) Solve() {
	tsp.ga.Evolution()
}

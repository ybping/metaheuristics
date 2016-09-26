package tsp

import (
	"github.com/ybping/metaheuristics/ga"
	"log"
	"math"
	"math/rand"
)

// const arguments
const (
	INITSCORE = 0
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
	cityOrder := rand.Perm(cityCount)
	for i := 0; i < cityCount; i++ {
		idx := rand.Intn(cityCount)
		cityOrder[i], cityOrder[idx] = cityOrder[idx], cityOrder[i]
	}
	return Species{
		cityOrder: cityOrder,
		cityCount: cityCount,
		score:     INITSCORE,
	}

}

// Cross ...
func (s Species) Cross(ts ga.Species) ga.Species {
	leftPos := rand.Intn(s.cityCount)
	rightPos := rand.Intn(s.cityCount)
	if leftPos > rightPos {
		leftPos, rightPos = rightPos, leftPos
	}
	t, hash := ts.(*Species), make(map[int]bool)
	for _, v := range t.cityOrder[leftPos:rightPos] {
		hash[v] = true
	}
	var cityOrder []int
	var index = 0
	for _, v := range s.cityOrder {
		if index == leftPos {
			cityOrder = append(cityOrder, t.cityOrder[leftPos:rightPos]...)
			index++
		}
		if _, ok := hash[v]; !ok {
			cityOrder = append(cityOrder, v)
			index++
		}
	}
	return &Species{
		cityOrder: cityOrder,
		cityCount: s.cityCount,
		score:     INITSCORE,
		tsp:       s.tsp,
	}
}

// Mutate ...
func (s Species) Mutate() ga.Species {
	first, second := rand.Intn(s.cityCount), rand.Intn(s.cityCount)
	s.cityOrder[first], s.cityOrder[second] = s.cityOrder[second], s.cityOrder[first]
	return &s
}

// Fitness ...
func (s Species) Fitness() float64 {
	return s.tsp.fitness(s)
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
	tsp := &TSP{cities: cities}

	// init specices
	var population []ga.Species
	for i := 0; i < 100; i++ {
		species := newSpecies(len(cities))
		species.tsp = tsp
		population = append(population, &species)
		//	log.Println(species.cityOrder, species.Fitness())
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
	log.Println("Start TSP")
	bestSpecies := tsp.ga.Evolution()
	log.Println(bestSpecies.(*Species).cityOrder)
}

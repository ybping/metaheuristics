package ga

import (
	"math/rand"
)

// Species interace 封装了具体问题的通用接口
type Species interface {
	cross(species Species) Species
	mutate() Species
	fitness() float64
}

// GeneticAlgorithm 遗传算法
type GeneticAlgorithm struct {
	// 种群规模
	population []Species
	// 种群大小
	populationSize int
	// 交叉概率
	crossRate float32
	// 变异概率
	mutationRate float32
	// 进化代数
	generationSize int
	// 进化最佳方案
	bestSpecies Species
}

// NewGeneticAlgorithm 返回一个遗传算法实例
func NewGeneticAlgorithm(population []Species, generationSize int, crossRate, mutationRate float32) *GeneticAlgorithm {
	return &GeneticAlgorithm{
		population:     population,
		populationSize: len(population),
		generationSize: generationSize,
		crossRate:      crossRate,
		mutationRate:   mutationRate,
	}
}

func (ga GeneticAlgorithm) rounds() float64 {
	var rounds = 0.0
	ga.bestSpecies = ga.population[0]
	for index := range ga.population {
		rounds += ga.population[index].fitness()
		if ga.bestSpecies.fitness() < ga.population[index].fitness() {
			ga.bestSpecies = ga.population[index]
		}
	}
	return rounds
}

func (ga GeneticAlgorithm) getParent(rounds float64) Species {
	rate := rand.Float64() * rounds
	for _, species := range ga.population {
		rate -= species.fitness()
		if rate <= 0 {
			return species
		}
	}
	panic("Unexpected Error: getSpecies")
}

func (ga GeneticAlgorithm) getChild(father, mather Species) (child Species) {
	// 交叉
	rate := rand.Float32()
	if rate < ga.crossRate {
		child = father.cross(mather)
	}

	// 突变
	rate = rand.Float32()
	if rate < ga.mutationRate {
		child = child.mutate()
	}
	return child
}

// Evolution 进化演进
func (ga *GeneticAlgorithm) Evolution() {
	for i := 0; i < ga.generationSize; i++ {
		var newPopulation []Species
		// 对当前种群进行适应值计算
		rounds := ga.rounds()

		// 精英选择优化，最好的物种基因尽可能往下一代传递
		newPopulation = append(newPopulation, ga.bestSpecies)

		// 进行新一代进化
		for i := 0; i < len(ga.population); i++ {
			father := ga.getParent(rounds)
			mather := ga.getParent(rounds)
			newPopulation = append(newPopulation, ga.getChild(father, mather))
		}
		ga.population = newPopulation
	}
}

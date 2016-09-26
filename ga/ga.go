package ga

import (
	"log"
	"math/rand"
)

// Species interace 封装了具体问题的通用接口
type Species interface {
	Cross(Species) Species
	Mutate() Species
	Fitness() float64
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
func NewGeneticAlgorithm(generationSize int, crossRate, mutationRate float32, population []Species) *GeneticAlgorithm {
	return &GeneticAlgorithm{
		population:     population,
		populationSize: len(population),
		generationSize: generationSize,
		crossRate:      crossRate,
		mutationRate:   mutationRate,
	}
}

func (ga *GeneticAlgorithm) rounds() float64 {
	var rounds = 0.0
	ga.bestSpecies = ga.population[0]
	for index := range ga.population {
		rounds += ga.population[index].Fitness()
		if ga.bestSpecies.Fitness() < ga.population[index].Fitness() {
			ga.bestSpecies = ga.population[index]
		}
	}
	return rounds
}

func (ga GeneticAlgorithm) getParent(rounds float64) Species {
	rate := rand.Float64() * rounds
	for _, species := range ga.population {
		rate -= species.Fitness()
		if rate <= 0 {
			return species
		}
	}
	panic("Unexpected Error: getSpecies")
}

func (ga GeneticAlgorithm) getChild(father, mather Species) (child Species) {
	child = father
	// 交叉
	rate := rand.Float32()
	if rate < ga.crossRate {
		child = father.Cross(mather)
	}

	// 突变
	rate = rand.Float32()
	if rate < ga.mutationRate {
		child = child.Mutate()
	}
	return child
}

// Evolution 进化演进
func (ga *GeneticAlgorithm) Evolution() Species {
	for i := 0; i < ga.generationSize; i++ {
		var newPopulation []Species
		// 对当前种群进行适应值计算
		rounds := ga.rounds()

		// 精英选择优化，最好的物种基因尽可能往下一代传递
		newPopulation = append(newPopulation, ga.bestSpecies)
		// 进行新一代进化
		count := 0
		for count < len(ga.population) {
			father := ga.getParent(rounds)
			mather := ga.getParent(rounds)
			child := ga.getChild(father, mather)
			if child != nil {
				newPopulation = append(newPopulation, child)
				count++
			}
		}
		ga.population = newPopulation
		log.Println(ga.bestSpecies.Fitness())
	}
	return ga.bestSpecies
}

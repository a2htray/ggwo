package ggwo

import (
	"math"
	"math/rand"
)

type Option func(*GWO)

func WithT(t int) Option {
	return func(g *GWO) {
		g.t = t
	}
}

func WithObjectFunc(objectFunc func([]float64) float64) Option {
	return func(g *GWO) {
		g.objectFunc = objectFunc
	}
}

func WithLowerLimit(limit []float64) Option {
	return func(g *GWO) {
		g.lowerLimit = limit
	}
}

func WithUpperLimit(limit []float64) Option {
	return func(g *GWO) {
		g.upperLimit = limit
	}
}

type GWO struct {
	m            int
	n            int
	t            int
	objectFunc   func([]float64) float64
	lowerLimit   []float64
	upperLimit   []float64
	population   []*Wolf
	historyBests []*Wolf
	alpha        *Wolf
	beta         *Wolf
	gamma        *Wolf
}

func (g *GWO) Run() {
	var (
		a  float64
		a1 float64
		a2 float64
		a3 float64
		c1 float64
		c2 float64
		c3 float64
	)
	x1 := make([]float64, g.n)
	x2 := make([]float64, g.n)
	x3 := make([]float64, g.n)
	for t := 0; t < g.t; t++ {
		a = 2.0 * (1.0 - float64(t)/float64(g.t))
		for i := 0; i < g.m; i++ {
			a1, a2, a3 = a*(2.0*rand.Float64()-1), a*(2.0*rand.Float64()), a*(2.0*rand.Float64())
			c1, c2, c3 = a*rand.Float64(), a*rand.Float64(), a*rand.Float64()
			wolf := &Wolf{
				position: make([]float64, g.n),
				fitness:  0,
			}
			for j := 0; j < g.n; j++ {
				x1[j] = g.alpha.position[j] - a1*math.Abs(c1*g.alpha.position[j]-g.population[i].position[j])
				x2[j] = g.beta.position[j] - a2*math.Abs(c2*g.beta.position[j]-g.population[i].position[j])
				x3[j] = g.gamma.position[j] - a3*math.Abs(c3*g.gamma.position[j]-g.population[i].position[j])
				wolf.position[j] = (x1[j] + x2[j] + x3[j]) / 3.0
			}
			if wolf.calculate(g.objectFunc) < g.population[i].fitness {
				g.population[i] = wolf
			}
		}
		fitnessValues := make([]float64, g.m)
		for i := 0; i < g.m; i++ {
			fitnessValues[i] = g.population[i].fitness
		}
		sortedIndexes := argsortFunc(fitnessValues)
		alphaIndex, betaIndex, gammaIndex := sortedIndexes[0], sortedIndexes[1], sortedIndexes[2]
		g.alpha = g.population[alphaIndex].copy()
		g.beta = g.population[betaIndex].copy()
		g.gamma = g.population[gammaIndex].copy()

		g.historyBests[t] = g.alpha.copy()
	}
}

func (g *GWO) HistoryBests() []*Wolf {
	return g.historyBests
}

func New(m, n int, options ...Option) *GWO {
	gwo := &GWO{
		m: m,
		n: n,
		t: 100,
		objectFunc: func(fs []float64) float64 {
			var fitness float64
			for _, f := range fs {
				fitness = math.Pow(f, 2)
			}
			return fitness
		},
		lowerLimit: floats(-10, n),
		upperLimit: floats(10, n),
		population: make([]*Wolf, m),
	}

	for _, option := range options {
		option(gwo)
	}

	gwo.historyBests = make([]*Wolf, gwo.t)
	gwo.population = wolves(m, n, gwo.lowerLimit, gwo.upperLimit)
	fitnessValues := make([]float64, gwo.m)
	for i := 0; i < gwo.m; i++ {
		fitnessValues[i] = gwo.population[i].calculate(gwo.objectFunc)
	}
	sortedIndexes := argsortFunc(fitnessValues)
	alphaIndex, betaIndex, gammaIndex := sortedIndexes[0], sortedIndexes[1], sortedIndexes[2]
	gwo.alpha = gwo.population[alphaIndex].copy()
	gwo.beta = gwo.population[betaIndex].copy()
	gwo.gamma = gwo.population[gammaIndex].copy()
	return gwo
}

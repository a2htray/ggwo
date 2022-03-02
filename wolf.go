package ggwo

import "math/rand"

type Wolf struct {
	position []float64
	fitness float64
}

func (w *Wolf) calculate(f func([]float64)float64) float64 {
	w.fitness = f(w.position)
	return w.fitness
}

func (w Wolf) copy() *Wolf {
	newWolf := &Wolf{
		position: make([]float64, len(w.position)),
		fitness:  w.fitness,
	}
	copy(newWolf.position, w.position)
	return newWolf
}

func (w Wolf) Fitness() float64 {
	return w.fitness
}

func (w Wolf) Values() []float64 {
	return w.position
}

func NewWolf(n int, lowerLimit, upperLimit []float64) *Wolf {
	produce := func(lower, upper []float64) []float64 {
		diff := float64sSub(upper, lower)
		ret := make([]float64, n)
		for i := 0; i < n; i++ {
			ret[i] = lower[i] + rand.Float64() * diff[i]
		}
		return ret
	}
	p := &Wolf{
		position: produce(lowerLimit, upperLimit),
	}
	return p
}

func wolves(m, n int, lowerLimit, upperLimit []float64) []*Wolf {
	ps := make([]*Wolf, m)
	for i := 0; i < m; i++ {
		ps[i] = NewWolf(n, lowerLimit, upperLimit)
	}
	return ps
}

package ggwo

func floats(v float64, n int) []float64 {
	ret := make([]float64, n)
	for i, _ := range ret {
		ret[i] = v
	}
	return ret
}

func float64sSub(a, b []float64) []float64 {
	ret := make([]float64, len(a))
	for i, v := range a {
		ret[i] = v - b[i]
	}
	return ret
}
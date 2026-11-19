package domain

type WeightedSumStrategy struct{}

func (s WeightedSumStrategy) Calculate(
	weightedSum float64,
	totalWeight float64,
) float64 {
	return min(weightedSum*10, 10)
}

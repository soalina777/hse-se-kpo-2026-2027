package domain

type NormalizedWeightedStrategy struct{}

func (s NormalizedWeightedStrategy) Calculate(
	weightedSum float64,
	totalWeight float64,
) float64 {
	if totalWeight == 0 {
		return 0
	}

	return weightedSum / totalWeight * 10
}

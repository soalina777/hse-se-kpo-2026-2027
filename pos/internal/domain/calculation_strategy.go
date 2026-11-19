package domain

type CalculationStrategy interface {
	Calculate(weightedSum float64, totalWeight float64) float64
}

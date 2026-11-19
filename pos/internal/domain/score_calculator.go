package domain

import (
	"fmt"
	"math"
	"strings"
)

type ScoreCalculator struct {
	strategy CalculationStrategy
}

func NewScoreCalculator(strategy CalculationStrategy) ScoreCalculator {
	return ScoreCalculator{
		strategy: strategy,
	}
}

func (sc ScoreCalculator) Calculate(
	evaluations []Evaluation,
	criterions []Criterion,
) (float64, error) {
	if sc.strategy == nil {
		return 0, fmt.Errorf("calculation strategy is not configured")
	}

	criterionsByID := make(map[string]Criterion)

	for _, criterion := range criterions {
		if strings.TrimSpace(criterion.ID()) == "" {
			return 0, fmt.Errorf("criterion ID cannot be empty")
		}

		if criterion.MaxScore() <= 0 {
			return 0, fmt.Errorf(
				"criterion %s has invalid maximum score",
				criterion.ID(),
			)
		}

		weight := criterion.Weight()

		if math.IsNaN(weight) || math.IsInf(weight, 0) {
			return 0, fmt.Errorf(
				"criterion %s has non-finite weight",
				criterion.ID(),
			)
		}

		if weight < 0 || weight > 1 {
			return 0, fmt.Errorf(
				"criterion %s has invalid weight",
				criterion.ID(),
			)
		}

		if _, exists := criterionsByID[criterion.ID()]; exists {
			return 0, fmt.Errorf(
				"duplicate criterion ID: %s",
				criterion.ID(),
			)
		}

		criterionsByID[criterion.ID()] = criterion
	}

	var weightedSum float64
	var totalWeight float64

	for _, evaluation := range evaluations {
		if strings.TrimSpace(evaluation.StudentID()) == "" {
			return 0, fmt.Errorf("evaluation has empty student ID")
		}

		if strings.TrimSpace(evaluation.CriterionID()) == "" {
			return 0, fmt.Errorf("evaluation has empty criterion ID")
		}

		criterion, ok := criterionsByID[evaluation.CriterionID()]

		if !ok {
			return 0, fmt.Errorf(
				"criterion with ID %s not found",
				evaluation.CriterionID(),
			)
		}

		score := evaluation.Score()

		if math.IsNaN(score) || math.IsInf(score, 0) {
			return 0, fmt.Errorf(
				"evaluation for criterion %s has non-finite score",
				evaluation.CriterionID(),
			)
		}

		if score < 0 || score > float64(criterion.MaxScore()) {
			return 0, fmt.Errorf(
				"evaluation score for criterion %s is out of range",
				evaluation.CriterionID(),
			)
		}

		normalizedScore := score / float64(criterion.MaxScore())

		weightedSum += normalizedScore * criterion.Weight()
		totalWeight += criterion.Weight()
	}

	result := sc.strategy.Calculate(weightedSum, totalWeight)

	if math.IsNaN(result) || math.IsInf(result, 0) {
		return 0, fmt.Errorf("calculation strategy returned non-finite result")
	}

	if result < 0 || result > 10 {
		return 0, fmt.Errorf(
			"calculation strategy returned result outside [0, 10]: %f",
			result,
		)
	}

	return result, nil
}

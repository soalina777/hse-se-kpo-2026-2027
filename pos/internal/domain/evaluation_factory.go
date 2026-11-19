package domain

import (
	"errors"
	"math"
	"strings"
)

type EvaluationFactory struct{}

func (f EvaluationFactory) Create(
	studentID string,
	criterion Criterion,
	score float64,
) (Evaluation, error) {
	studentID = strings.TrimSpace(studentID)

	if studentID == "" {
		return Evaluation{}, errors.New("student ID cannot be empty")
	}

	if criterion.id == "" || criterion.maxScore <= 0 {
		return Evaluation{}, errors.New("invalid criterion")
	}

	if math.IsNaN(score) || math.IsInf(score, 0) {
		return Evaluation{}, errors.New("score must be a finite number")
	}

	if score < 0 || score > float64(criterion.maxScore) {
		return Evaluation{}, errors.New(
			"score must be between 0 and max score",
		)
	}

	return Evaluation{
		studentID:   studentID,
		criterionID: criterion.id,
		score:       score,
	}, nil
}

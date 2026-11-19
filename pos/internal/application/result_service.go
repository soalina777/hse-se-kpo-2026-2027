package application

import (
	"fmt"
	"pos/internal/domain"
)

type ResultService struct {
	studentRepo    domain.StudentRepository
	evaluationRepo domain.EvaluationRepository
	criterionRepo  domain.CriterionRepository
	calculator     domain.ScoreCalculator
}

func NewResultService(
	studentRepo domain.StudentRepository,
	evaluationRepo domain.EvaluationRepository,
	criterionRepo domain.CriterionRepository,
	calculator domain.ScoreCalculator,
) *ResultService {
	return &ResultService{
		studentRepo:    studentRepo,
		evaluationRepo: evaluationRepo,
		criterionRepo:  criterionRepo,
		calculator:     calculator,
	}
}

func (rs *ResultService) CalculateResult(studentID string) (float64, error) {
	_, err := rs.studentRepo.FindByID(studentID)

	if err != nil {
		return 0, fmt.Errorf("student not found: %w", err)
	}

	evaluations, err := rs.evaluationRepo.FindByStudentID(studentID)

	if err != nil {
		return 0, fmt.Errorf("failed to get student evaluations: %w", err)
	}

	criterions, err := rs.criterionRepo.FindAll()

	if err != nil {
		return 0, fmt.Errorf("failed to get criteria: %w", err)
	}

	totalScore, err := rs.calculator.Calculate(evaluations, criterions)

	if err != nil {
		return 0, fmt.Errorf("failed to calculate result: %w", err)
	}

	return totalScore, nil
}

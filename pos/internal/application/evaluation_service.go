package application

import (
	"fmt"
	"pos/internal/domain"
	"strings"
)

type EvaluationService struct {
	studentRepo       domain.StudentRepository
	criterionRepo     domain.CriterionRepository
	evaluationRepo    domain.EvaluationRepository
	evaluationFactory domain.EvaluationFactory
}

func NewEvaluationService(
	studentRepo domain.StudentRepository,
	criterionRepo domain.CriterionRepository,
	evaluationRepo domain.EvaluationRepository,
	evaluationFactory domain.EvaluationFactory,
) *EvaluationService {
	return &EvaluationService{
		studentRepo:       studentRepo,
		criterionRepo:     criterionRepo,
		evaluationRepo:    evaluationRepo,
		evaluationFactory: evaluationFactory,
	}
}

func (es *EvaluationService) Evaluate(
	studentID string,
	criterionID string,
	score float64,
) error {
	studentID = strings.TrimSpace(studentID)
	criterionID = strings.TrimSpace(criterionID)

	_, err := es.studentRepo.FindByID(studentID)

	if err != nil {
		return fmt.Errorf("failed to find student %s: %w", studentID, err)
	}

	criterion, err := es.criterionRepo.FindByID(criterionID)

	if err != nil {
		return fmt.Errorf("failed to find criterion %s: %w", criterionID, err)
	}

	evaluation, err := es.evaluationFactory.Create(studentID, criterion, score)

	if err != nil {
		return fmt.Errorf("failed to create evaluation: %w", err)
	}

	err = es.evaluationRepo.Save(evaluation)

	if err != nil {
		return fmt.Errorf("failed to save evaluation: %w", err)
	}

	return nil
}

func (es *EvaluationService) GetEvaluations(studentID string) ([]domain.Evaluation, error) {
	studentID = strings.TrimSpace(studentID)
	_, err := es.studentRepo.FindByID(studentID)
	if err != nil {
		return nil, fmt.Errorf("failed to find student: %w", err)
	}

	evaluations, err := es.evaluationRepo.FindByStudentID(studentID)
	if err != nil {
		return nil, fmt.Errorf("failed to get evaluations: %w", err)
	}

	return evaluations, nil
}

func (es *EvaluationService) DeleteEvaluation(
	studentID string,
	criterionID string,
) error {
	studentID = strings.TrimSpace(studentID)
	criterionID = strings.TrimSpace(criterionID)
	_, err := es.studentRepo.FindByID(studentID)
	if err != nil {
		return fmt.Errorf("failed to find student: %w", err)
	}

	_, err = es.criterionRepo.FindByID(criterionID)
	if err != nil {
		return fmt.Errorf("failed to find criterion: %w", err)
	}

	err = es.evaluationRepo.Delete(studentID, criterionID)
	if err != nil {
		return fmt.Errorf("failed to delete evaluation: %w", err)
	}

	return nil
}

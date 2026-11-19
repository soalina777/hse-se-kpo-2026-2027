package infrastructure

import (
	"fmt"
	"pos/internal/domain"
)

type evaluationKey struct {
	studentID   string
	criterionID string
}

type MemoryEvaluationRepository struct {
	evaluations map[evaluationKey]domain.Evaluation
}

func NewMemoryEvaluationRepository() *MemoryEvaluationRepository {
	return &MemoryEvaluationRepository{
		evaluations: make(map[evaluationKey]domain.Evaluation),
	}
}

func (er *MemoryEvaluationRepository) Save(evaluation domain.Evaluation) error {
	evalID := evaluationKey{
		studentID:   evaluation.StudentID(),
		criterionID: evaluation.CriterionID(),
	}

	er.evaluations[evalID] = evaluation
	return nil
}

func (er *MemoryEvaluationRepository) FindByStudentID(id string) ([]domain.Evaluation, error) {
	result := make([]domain.Evaluation, 0)

	for _, evaluation := range er.evaluations {
		if evaluation.StudentID() == id {
			result = append(result, evaluation)
		}
	}

	return result, nil
}

func (er *MemoryEvaluationRepository) Delete(studentID string, criterionID string) error {
	evalID := evaluationKey{
		studentID:   studentID,
		criterionID: criterionID,
	}

	_, ok := er.evaluations[evalID]

	if !ok {
		return fmt.Errorf(
			"student with ID %s does not have score for criterion with ID %s",
			evalID.studentID,
			evalID.criterionID,
		)
	}

	delete(er.evaluations, evalID)
	return nil
}

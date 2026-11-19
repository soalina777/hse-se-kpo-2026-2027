package application

import (
	"pos/internal/domain"
)

type POSFacade struct {
	management *ManagementService
	evaluation *EvaluationService
	result     *ResultService
	export     *ExportService
}

func NewPOSFacade(
	management *ManagementService,
	evaluation *EvaluationService,
	result *ResultService,
	export *ExportService,
) *POSFacade {
	return &POSFacade{
		management: management,
		evaluation: evaluation,
		result:     result,
		export:     export,
	}
}

func (facade *POSFacade) AddStudent(id string, name string) error {
	return facade.management.AddStudent(id, name)
}

func (facade *POSFacade) AddCriterion(
	id string,
	name string,
	maxScore int,
	weight float64,
) error {
	return facade.management.AddCriterion(id, name, maxScore, weight)
}

func (facade *POSFacade) GetStudents() ([]domain.Student, error) {
	return facade.management.GetStudents()
}

func (facade *POSFacade) GetCriterions() ([]domain.Criterion, error) {
	return facade.management.GetCriterions()
}

func (facade *POSFacade) Evaluate(
	studentID string,
	criterionID string,
	score float64,
) error {
	return facade.evaluation.Evaluate(studentID, criterionID, score)
}

func (facade *POSFacade) CalculateResult(studentID string) (float64, error) {
	return facade.result.CalculateResult(studentID)
}

func (facade *POSFacade) GetEvaluations(studentID string) ([]domain.Evaluation, error) {
	return facade.evaluation.GetEvaluations(studentID)
}

func (facade *POSFacade) DeleteEvaluation(
	studentID string,
	criterionID string,
) error {
	return facade.evaluation.DeleteEvaluation(studentID, criterionID)
}

func (facade *POSFacade) ExportResults(path string) error {
	return facade.export.ExportResults(path)
}

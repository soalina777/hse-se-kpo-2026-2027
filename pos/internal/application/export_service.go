package application

import (
	"fmt"
	"pos/internal/domain"
)

type ExportService struct {
	studentRepo    domain.StudentRepository
	criterionRepo  domain.CriterionRepository
	evaluationRepo domain.EvaluationRepository
	calculator     domain.ScoreCalculator
	exporter       Exporter
}

func NewExportService(
	studentRepo domain.StudentRepository,
	criterionRepo domain.CriterionRepository,
	evaluationRepo domain.EvaluationRepository,
	calculator domain.ScoreCalculator,
	exporter Exporter,
) *ExportService {
	return &ExportService{
		studentRepo:    studentRepo,
		criterionRepo:  criterionRepo,
		evaluationRepo: evaluationRepo,
		calculator:     calculator,
		exporter:       exporter,
	}
}

func (es *ExportService) ExportResults(path string) error {
	students, err := es.studentRepo.FindAll()
	if err != nil {
		return fmt.Errorf("failed to get students: %w", err)
	}

	criterions, err := es.criterionRepo.FindAll()
	if err != nil {
		return fmt.Errorf("failed to get criteria: %w", err)
	}

	criterionsByID := make(map[string]domain.Criterion)

	for _, criterion := range criterions {
		criterionsByID[criterion.ID()] = criterion
	}

	rows := make([]ExportRow, 0)

	for _, student := range students {
		evaluations, err := es.evaluationRepo.FindByStudentID(student.ID())
		if err != nil {
			return fmt.Errorf(
				"failed to get evaluations for student %s: %w",
				student.ID(),
				err,
			)
		}

		totalScore, err := es.calculator.Calculate(evaluations, criterions)
		if err != nil {
			return fmt.Errorf(
				"failed to calculate result for student %s: %w",
				student.ID(),
				err,
			)
		}

		if len(evaluations) == 0 {
			row := ExportRow{
				StudentID:   student.ID(),
				StudentName: student.Name(),
				Criterion:   "",
				Score:       0,
				FinalScore:  totalScore,
			}

			rows = append(rows, row)
			continue
		}

		for _, evaluation := range evaluations {
			criterion, ok := criterionsByID[evaluation.CriterionID()]
			if !ok {
				return fmt.Errorf(
					"criterion with ID %s not found",
					evaluation.CriterionID(),
				)
			}

			row := ExportRow{
				StudentID:   student.ID(),
				StudentName: student.Name(),
				Criterion:   criterion.Name(),
				Score:       evaluation.Score(),
				FinalScore:  totalScore,
			}

			rows = append(rows, row)
		}
	}

	return es.exporter.Export(path, rows)
}

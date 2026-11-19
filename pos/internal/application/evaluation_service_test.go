package application_test

import (
	"math"
	"testing"

	"pos/internal/application"
	"pos/internal/domain"
	"pos/internal/infrastructure"
)

func setupEvaluationService(t *testing.T) (
	*application.EvaluationService,
	*infrastructure.MemoryEvaluationRepository,
) {
	t.Helper()

	studentRepo := infrastructure.NewMemoryStudentRepository()
	criterionRepo := infrastructure.NewMemoryCriterionRepository()
	evaluationRepo := infrastructure.NewMemoryEvaluationRepository()

	student, err := domain.NewStudent("001", "Ivan")
	if err != nil {
		t.Fatal(err)
	}

	err = studentRepo.Save(student)
	if err != nil {
		t.Fatal(err)
	}

	criterion, err := domain.NewCriterion("exam", "Exam", 10, 1)
	if err != nil {
		t.Fatal(err)
	}

	err = criterionRepo.Save(criterion)
	if err != nil {
		t.Fatal(err)
	}

	service := application.NewEvaluationService(
		studentRepo,
		criterionRepo,
		evaluationRepo,
		domain.EvaluationFactory{},
	)

	return service, evaluationRepo
}

func TestEvaluationService_Evaluate(t *testing.T) {
	service, evaluationRepo := setupEvaluationService(t)

	err := service.Evaluate("001", "exam", 8.5)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	evaluations, err := evaluationRepo.FindByStudentID("001")
	if err != nil {
		t.Fatal(err)
	}

	if len(evaluations) != 1 {
		t.Fatalf("expected 1 evaluation, got %d", len(evaluations))
	}

	if evaluations[0].Score() != 8.5 {
		t.Errorf("expected score 8.5, got %f", evaluations[0].Score())
	}
}

func TestEvaluationService_Evaluate_UpdatesExistingEvaluation(t *testing.T) {
	service, evaluationRepo := setupEvaluationService(t)

	err := service.Evaluate("001", "exam", 6)
	if err != nil {
		t.Fatal(err)
	}

	err = service.Evaluate("001", "exam", 9)
	if err != nil {
		t.Fatal(err)
	}

	evaluations, err := evaluationRepo.FindByStudentID("001")
	if err != nil {
		t.Fatal(err)
	}

	if len(evaluations) != 1 {
		t.Fatalf("expected 1 evaluation after update, got %d", len(evaluations))
	}

	if evaluations[0].Score() != 9 {
		t.Errorf("expected updated score 9, got %f", evaluations[0].Score())
	}
}

func TestEvaluationService_Evaluate_StudentNotFound(t *testing.T) {
	service, _ := setupEvaluationService(t)

	err := service.Evaluate("999", "exam", 8)

	if err == nil {
		t.Fatal("expected error for nonexistent student")
	}
}

func TestEvaluationService_Evaluate_CriterionNotFound(t *testing.T) {
	service, _ := setupEvaluationService(t)

	err := service.Evaluate("001", "unknown", 8)

	if err == nil {
		t.Fatal("expected error for nonexistent criterion")
	}
}

func TestEvaluationService_Evaluate_InvalidScore(t *testing.T) {
	service, evaluationRepo := setupEvaluationService(t)

	err := service.Evaluate("001", "exam", 11)

	if err == nil {
		t.Fatal("expected error for score greater than maximum")
	}

	evaluations, err := evaluationRepo.FindByStudentID("001")
	if err != nil {
		t.Fatal(err)
	}

	if len(evaluations) != 0 {
		t.Error("evaluation should not be saved when score is invalid")
	}
}

func TestEvaluationService_GetEvaluations(t *testing.T) {
	service, _ := setupEvaluationService(t)

	err := service.Evaluate("001", "exam", 8)
	if err != nil {
		t.Fatal(err)
	}

	evaluations, err := service.GetEvaluations("001")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(evaluations) != 1 {
		t.Fatalf("expected 1 evaluation, got %d", len(evaluations))
	}

	if evaluations[0].Score() != 8 {
		t.Errorf("expected score 8, got %f", evaluations[0].Score())
	}
}

func TestEvaluationService_GetEvaluations_StudentNotFound(t *testing.T) {
	service, _ := setupEvaluationService(t)

	_, err := service.GetEvaluations("999")

	if err == nil {
		t.Fatal("expected error for nonexistent student")
	}
}

func TestEvaluationService_DeleteEvaluation(t *testing.T) {
	service, _ := setupEvaluationService(t)

	err := service.Evaluate("001", "exam", 8)
	if err != nil {
		t.Fatal(err)
	}

	err = service.DeleteEvaluation("001", "exam")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	evaluations, err := service.GetEvaluations("001")
	if err != nil {
		t.Fatal(err)
	}

	if len(evaluations) != 0 {
		t.Errorf("expected 0 evaluations after deletion, got %d", len(evaluations))
	}
}

func TestEvaluationService_DeleteEvaluation_NotFound(t *testing.T) {
	service, _ := setupEvaluationService(t)

	err := service.DeleteEvaluation("001", "exam")

	if err == nil {
		t.Fatal("expected error when deleting nonexistent evaluation")
	}
}

func TestEvaluationService_Evaluate_NonFiniteScore(t *testing.T) {
	tests := []struct {
		name  string
		score float64
	}{
		{
			name:  "NaN",
			score: math.NaN(),
		},
		{
			name:  "positive infinity",
			score: math.Inf(1),
		},
		{
			name:  "negative infinity",
			score: math.Inf(-1),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service, evaluationRepo := setupEvaluationService(t)

			err := service.Evaluate("001", "exam", tt.score)

			if err == nil {
				t.Fatal("expected error for non-finite score")
			}

			evaluations, err := evaluationRepo.FindByStudentID("001")
			if err != nil {
				t.Fatal(err)
			}

			if len(evaluations) != 0 {
				t.Errorf(
					"expected no saved evaluations, got %d",
					len(evaluations),
				)
			}
		})
	}
}

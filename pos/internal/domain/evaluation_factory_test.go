package domain

import (
	"math"
	"testing"
)

func TestEvaluationFactory_Create(t *testing.T) {
	factory := EvaluationFactory{}

	criterion, err := NewCriterion("exam", "Exam", 10, 1)
	if err != nil {
		t.Fatal(err)
	}

	evaluation, err := factory.Create("001", criterion, 8.5)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if evaluation.StudentID() != "001" {
		t.Errorf("expected student ID 001, got %s", evaluation.StudentID())
	}

	if evaluation.CriterionID() != "exam" {
		t.Errorf("expected criterion ID exam, got %s", evaluation.CriterionID())
	}

	if evaluation.Score() != 8.5 {
		t.Errorf("expected score 8.5, got %f", evaluation.Score())
	}
}

func TestEvaluationFactory_ValidBoundaryScores(t *testing.T) {
	factory := EvaluationFactory{}

	criterion, err := NewCriterion("exam", "Exam", 10, 1)
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name  string
		score float64
	}{
		{"zero score", 0},
		{"maximum score", 10},
		{"fractional score", 8.5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := factory.Create("001", criterion, tt.score)

			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
		})
	}
}

func TestEvaluationFactory_InvalidScores(t *testing.T) {
	factory := EvaluationFactory{}

	criterion, err := NewCriterion("exam", "Exam", 10, 1)
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name  string
		score float64
	}{
		{"negative score", -1},
		{"score greater than maximum", 11},
		{"NaN score", math.NaN()},
		{"positive infinity", math.Inf(1)},
		{"negative infinity", math.Inf(-1)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := factory.Create("001", criterion, tt.score)

			if err == nil {
				t.Errorf("expected error for score %v", tt.score)
			}
		})
	}
}

func TestEvaluationFactory_EmptyStudentID(t *testing.T) {
	factory := EvaluationFactory{}

	criterion, err := NewCriterion("exam", "Exam", 10, 1)
	if err != nil {
		t.Fatal(err)
	}

	tests := []string{"", "   ", "\t"}

	for _, studentID := range tests {
		_, err := factory.Create(studentID, criterion, 8)

		if err == nil {
			t.Errorf("expected error for student ID %q", studentID)
		}
	}
}

func TestEvaluationFactory_InvalidCriterion(t *testing.T) {
	factory := EvaluationFactory{}

	tests := []struct {
		name      string
		criterion Criterion
	}{
		{
			name:      "empty criterion",
			criterion: Criterion{},
		},
		{
			name: "criterion with zero maximum score",
			criterion: Criterion{
				id:       "exam",
				name:     "Exam",
				maxScore: 0,
				weight:   1,
			},
		},
		{
			name: "criterion with empty ID",
			criterion: Criterion{
				id:       "",
				name:     "Exam",
				maxScore: 10,
				weight:   1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := factory.Create("001", tt.criterion, 8)

			if err == nil {
				t.Error("expected error for invalid criterion")
			}
		})
	}
}

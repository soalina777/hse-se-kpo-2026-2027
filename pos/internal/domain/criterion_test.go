package domain

import (
	"math"
	"testing"
)

func TestNewCriterion(t *testing.T) {
	criterion, err := NewCriterion("1", "Exam", 10, 0.5)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if criterion.ID() != "1" {
		t.Errorf("expected ID 1, got %s", criterion.ID())
	}

	if criterion.Name() != "Exam" {
		t.Errorf("expected name Exam, got %s", criterion.Name())
	}

	if criterion.MaxScore() != 10 {
		t.Errorf("expected max score 10, got %d", criterion.MaxScore())
	}

	if criterion.Weight() != 0.5 {
		t.Errorf("expected weight 0.5, got %f", criterion.Weight())
	}
}

func TestNewCriterion_ZeroWeight(t *testing.T) {
	_, err := NewCriterion("1", "Exam", 10, 0)

	if err != nil {
		t.Fatalf("zero weight should be allowed: %v", err)
	}
}

func TestNewCriterion_MaximumWeight(t *testing.T) {
	_, err := NewCriterion("1", "Exam", 10, 1)

	if err != nil {
		t.Fatalf("weight 1 should be allowed: %v", err)
	}
}

func TestNewCriterion_EmptyID(t *testing.T) {
	_, err := NewCriterion("", "Exam", 10, 0.5)

	if err == nil {
		t.Fatal("expected error for empty ID")
	}
}

func TestNewCriterion_EmptyName(t *testing.T) {
	_, err := NewCriterion("1", "", 10, 0.5)

	if err == nil {
		t.Fatal("expected error for empty name")
	}
}

func TestNewCriterion_EmptyIDAndName(t *testing.T) {
	_, err := NewCriterion("      ", "     ", 10, 0.5)

	if err == nil {
		t.Fatal("expected error for empty ID and name")
	}
}

func TestNewCriterion_TrimsWhitespace(t *testing.T) {
	criterion, err := NewCriterion("  exam  ", "  Exam  ", 10, 0.5)

	if err != nil {
		t.Fatal(err)
	}

	if criterion.ID() != "exam" {
		t.Errorf("expected ID exam, got %q", criterion.ID())
	}

	if criterion.Name() != "Exam" {
		t.Errorf("expected name Exam, got %q", criterion.Name())
	}
}

func TestNewCriterion_InvalidMaxScore(t *testing.T) {
	_, err := NewCriterion("1", "Exam", 0, 0.5)

	if err == nil {
		t.Fatal("expected error for zero maximum score")
	}
}

func TestNewCriterion_NegativeMaxScore(t *testing.T) {
	_, err := NewCriterion("1", "Exam", -10, 0.5)

	if err == nil {
		t.Fatal("expected error for negative maximum score")
	}
}

func TestNewCriterion_NegativeWeight(t *testing.T) {
	_, err := NewCriterion("1", "Exam", 10, -0.1)

	if err == nil {
		t.Fatal("expected error for negative weight")
	}
}

func TestNewCriterion_WeightGreaterThanOne(t *testing.T) {
	_, err := NewCriterion("1", "Exam", 10, 1.1)

	if err == nil {
		t.Fatal("expected error for weight greater than 1")
	}
}

func TestNewCriterion_NaNWeight(t *testing.T) {
	_, err := NewCriterion("1", "Exam", 10, math.NaN())

	if err == nil {
		t.Fatal("expected error for NaN weight")
	}
}

func TestNewCriterion_PositiveInfinityWeight(t *testing.T) {
	_, err := NewCriterion("1", "Exam", 10, math.Inf(1))

	if err == nil {
		t.Fatal("expected error for positive infinity")
	}
}

func TestNewCriterion_NegativeInfinityWeight(t *testing.T) {
	_, err := NewCriterion("1", "Exam", 10, math.Inf(-1))

	if err == nil {
		t.Fatal("expected error for negative infinity")
	}
}

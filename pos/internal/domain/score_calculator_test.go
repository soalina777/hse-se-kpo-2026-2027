package domain

import (
	"math"
	"testing"
)

func TestScoreCalculator_SingleEvaluation(t *testing.T) {
	calculator := NewScoreCalculator(WeightedSumStrategy{})

	criterion, err := NewCriterion("exam", "Exam", 10, 1)
	if err != nil {
		t.Fatal(err)
	}

	evaluation := Evaluation{
		studentID:   "001",
		criterionID: "exam",
		score:       10,
	}

	result, err := calculator.Calculate(
		[]Evaluation{evaluation},
		[]Criterion{criterion},
	)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result != 10 {
		t.Errorf("expected 10, got %f", result)
	}
}

func TestScoreCalculator_FractionalScore(t *testing.T) {
	calculator := NewScoreCalculator(WeightedSumStrategy{})

	criterion, err := NewCriterion("exam", "Exam", 10, 1)
	if err != nil {
		t.Fatal(err)
	}

	evaluation := Evaluation{
		studentID:   "001",
		criterionID: "exam",
		score:       8.5,
	}

	result, err := calculator.Calculate(
		[]Evaluation{evaluation},
		[]Criterion{criterion},
	)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if math.Abs(result-8.5) > 1e-9 {
		t.Errorf("expected 8.5, got %f", result)
	}
}

func TestScoreCalculator_WeightedEvaluations(t *testing.T) {
	calculator := NewScoreCalculator(WeightedSumStrategy{})

	homework, err := NewCriterion("hw", "Homework", 10, 0.5)
	if err != nil {
		t.Fatal(err)
	}

	exam, err := NewCriterion("exam", "Exam", 10, 0.25)
	if err != nil {
		t.Fatal(err)
	}

	evaluations := []Evaluation{
		{
			studentID:   "001",
			criterionID: "hw",
			score:       8,
		},
		{
			studentID:   "001",
			criterionID: "exam",
			score:       6,
		},
	}

	result, err := calculator.Calculate(
		evaluations,
		[]Criterion{homework, exam},
	)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// (8/10 * 0.5 + 6/10 * 0.25) * 10 = 5.5
	if math.Abs(result-5.5) > 1e-9 {
		t.Errorf("expected 5.5, got %f", result)
	}
}

func TestScoreCalculator_WeightsSumGreaterThanOne(t *testing.T) {
	calculator := NewScoreCalculator(WeightedSumStrategy{})

	homework, err := NewCriterion("hw", "Homework", 10, 0.7)
	if err != nil {
		t.Fatal(err)
	}

	exam, err := NewCriterion("exam", "Exam", 10, 0.5)
	if err != nil {
		t.Fatal(err)
	}

	contest, err := NewCriterion("contest", "Contest", 10, 0.3)
	if err != nil {
		t.Fatal(err)
	}

	evaluations := []Evaluation{
		{studentID: "001", criterionID: "hw", score: 10},
		{studentID: "001", criterionID: "exam", score: 10},
		{studentID: "001", criterionID: "contest", score: 10},
	}

	result, err := calculator.Calculate(
		evaluations,
		[]Criterion{homework, exam, contest},
	)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result != 10 {
		t.Errorf("expected capped result 10, got %f", result)
	}
}

func TestScoreCalculator_EmptyEvaluations(t *testing.T) {
	calculator := NewScoreCalculator(WeightedSumStrategy{})

	criterion, err := NewCriterion("exam", "Exam", 10, 1)
	if err != nil {
		t.Fatal(err)
	}

	result, err := calculator.Calculate(
		[]Evaluation{},
		[]Criterion{criterion},
	)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result != 0 {
		t.Errorf("expected 0, got %f", result)
	}
}

func TestScoreCalculator_ZeroWeight(t *testing.T) {
	calculator := NewScoreCalculator(WeightedSumStrategy{})

	criterion, err := NewCriterion("exam", "Exam", 10, 0)
	if err != nil {
		t.Fatal(err)
	}

	evaluation := Evaluation{
		studentID:   "001",
		criterionID: "exam",
		score:       10,
	}

	result, err := calculator.Calculate(
		[]Evaluation{evaluation},
		[]Criterion{criterion},
	)

	if err != nil {
		t.Fatal(err)
	}

	if result != 0 {
		t.Errorf("expected 0, got %f", result)
	}
}

func TestScoreCalculator_CriterionNotFound(t *testing.T) {
	calculator := NewScoreCalculator(WeightedSumStrategy{})

	evaluation := Evaluation{
		studentID:   "001",
		criterionID: "unknown",
		score:       8,
	}

	_, err := calculator.Calculate(
		[]Evaluation{evaluation},
		[]Criterion{},
	)

	if err == nil {
		t.Fatal("expected error for missing criterion")
	}
}

func TestScoreCalculator_PartialScore(t *testing.T) {
	calculator := NewScoreCalculator(WeightedSumStrategy{})

	criterion, err := NewCriterion("exam", "Exam", 20, 0.5)
	if err != nil {
		t.Fatal(err)
	}

	evaluation := Evaluation{
		studentID:   "001",
		criterionID: "exam",
		score:       10,
	}

	result, err := calculator.Calculate(
		[]Evaluation{evaluation},
		[]Criterion{criterion},
	)

	if err != nil {
		t.Fatal(err)
	}

	// (10 / 20 * 0.5) * 10 = 2.5
	if math.Abs(result-2.5) > 1e-9 {
		t.Errorf("expected 2.5, got %f", result)
	}
}

func TestScoreCalculator_ResultNeverExceedsTen(t *testing.T) {
	calculator := NewScoreCalculator(WeightedSumStrategy{})

	criterion, err := NewCriterion("exam", "Exam", 10, 1)
	if err != nil {
		t.Fatal(err)
	}

	evaluation := Evaluation{
		studentID:   "001",
		criterionID: "exam",
		score:       10,
	}

	result, err := calculator.Calculate(
		[]Evaluation{evaluation},
		[]Criterion{criterion},
	)

	if err != nil {
		t.Fatal(err)
	}

	if result > 10 {
		t.Errorf("result must not exceed 10, got %f", result)
	}
}

func TestScoreCalculator_NormalizedStrategy(t *testing.T) {
	calculator := NewScoreCalculator(NormalizedWeightedStrategy{})

	criterion, err := NewCriterion("exam", "Exam", 10, 0.5)
	if err != nil {
		t.Fatal(err)
	}

	evaluation, err := (EvaluationFactory{}).Create(
		"001",
		criterion,
		8,
	)
	if err != nil {
		t.Fatal(err)
	}

	result, err := calculator.Calculate(
		[]Evaluation{evaluation},
		[]Criterion{criterion},
	)

	if err != nil {
		t.Fatal(err)
	}

	if math.Abs(result-8) > 1e-9 {
		t.Errorf("expected 8, got %f", result)
	}
}

func TestScoreCalculator_WeightedSumStrategy(t *testing.T) {
	calculator := NewScoreCalculator(WeightedSumStrategy{})

	criterion, err := NewCriterion("exam", "Exam", 10, 0.5)
	if err != nil {
		t.Fatal(err)
	}

	evaluation, err := (EvaluationFactory{}).Create(
		"001",
		criterion,
		8,
	)
	if err != nil {
		t.Fatal(err)
	}

	result, err := calculator.Calculate(
		[]Evaluation{evaluation},
		[]Criterion{criterion},
	)

	if err != nil {
		t.Fatal(err)
	}

	if math.Abs(result-4) > 1e-9 {
		t.Errorf("expected 4, got %f", result)
	}
}

func TestScoreCalculator_NormalizedStrategy_ZeroWeight(t *testing.T) {
	calculator := NewScoreCalculator(NormalizedWeightedStrategy{})

	criterion, err := NewCriterion("exam", "Exam", 10, 0)
	if err != nil {
		t.Fatal(err)
	}

	evaluation, err := (EvaluationFactory{}).Create(
		"001",
		criterion,
		10,
	)
	if err != nil {
		t.Fatal(err)
	}

	result, err := calculator.Calculate(
		[]Evaluation{evaluation},
		[]Criterion{criterion},
	)

	if err != nil {
		t.Fatal(err)
	}

	if result != 0 {
		t.Errorf("expected 0, got %f", result)
	}
}

func TestScoreCalculator_CompareStrategiesWithMultipleCriteria(t *testing.T) {
	homework, err := NewCriterion("hw", "Homework", 10, 0.5)
	if err != nil {
		t.Fatal(err)
	}

	exam, err := NewCriterion("exam", "Exam", 10, 0.25)
	if err != nil {
		t.Fatal(err)
	}

	evaluations := []Evaluation{
		{
			studentID:   "001",
			criterionID: "hw",
			score:       8,
		},
		{
			studentID:   "001",
			criterionID: "exam",
			score:       6,
		},
	}

	criteria := []Criterion{homework, exam}

	weightedCalculator := NewScoreCalculator(
		WeightedSumStrategy{},
	)

	normalizedCalculator := NewScoreCalculator(
		NormalizedWeightedStrategy{},
	)

	weightedResult, err := weightedCalculator.Calculate(
		evaluations,
		criteria,
	)
	if err != nil {
		t.Fatal(err)
	}

	normalizedResult, err := normalizedCalculator.Calculate(
		evaluations,
		criteria,
	)
	if err != nil {
		t.Fatal(err)
	}

	if math.Abs(weightedResult-5.5) > 1e-9 {
		t.Errorf("expected weighted result 5.5, got %f", weightedResult)
	}

	if math.Abs(normalizedResult-7.333333333333333) > 1e-9 {
		t.Errorf(
			"expected normalized result approximately 7.3333, got %f",
			normalizedResult,
		)
	}

	if weightedResult == normalizedResult {
		t.Error("expected strategies to produce different results")
	}
}

func TestScoreCalculator_IncorrectInput(t *testing.T) {
	validCriterion := Criterion{
		id:       "exam",
		name:     "Exam",
		maxScore: 10,
		weight:   1,
	}

	tests := []struct {
		name        string
		evaluations []Evaluation
		criteria    []Criterion
	}{
		{
			name: "NaN score",
			evaluations: []Evaluation{
				{studentID: "001", criterionID: "exam", score: math.NaN()},
			},
			criteria: []Criterion{validCriterion},
		},
		{
			name: "positive infinity score",
			evaluations: []Evaluation{
				{studentID: "001", criterionID: "exam", score: math.Inf(1)},
			},
			criteria: []Criterion{validCriterion},
		},
		{
			name: "negative infinity score",
			evaluations: []Evaluation{
				{studentID: "001", criterionID: "exam", score: math.Inf(-1)},
			},
			criteria: []Criterion{validCriterion},
		},
		{
			name: "negative score",
			evaluations: []Evaluation{
				{studentID: "001", criterionID: "exam", score: -1},
			},
			criteria: []Criterion{validCriterion},
		},
		{
			name: "score exceeds maximum",
			evaluations: []Evaluation{
				{studentID: "001", criterionID: "exam", score: 11},
			},
			criteria: []Criterion{validCriterion},
		},
		{
			name: "zero maximum score",
			evaluations: []Evaluation{
				{studentID: "001", criterionID: "exam", score: 5},
			},
			criteria: []Criterion{
				{id: "exam", name: "Exam", maxScore: 0, weight: 1},
			},
		},
		{
			name: "NaN criterion weight",
			evaluations: []Evaluation{
				{studentID: "001", criterionID: "exam", score: 5},
			},
			criteria: []Criterion{
				{id: "exam", name: "Exam", maxScore: 10, weight: math.NaN()},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calculator := NewScoreCalculator(WeightedSumStrategy{})

			_, err := calculator.Calculate(tt.evaluations, tt.criteria)

			if err == nil {
				t.Fatal("expected error for incorrect input")
			}
		})
	}
}

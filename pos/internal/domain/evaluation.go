package domain

type Evaluation struct {
	studentID   string
	criterionID string
	score       float64
}

func (e Evaluation) StudentID() string {
	return e.studentID
}

func (e Evaluation) CriterionID() string {
	return e.criterionID
}

func (e Evaluation) Score() float64 {
	return e.score
}

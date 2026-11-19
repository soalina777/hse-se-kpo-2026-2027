package domain

type EvaluationRepository interface {
	Save(evaluation Evaluation) error
	FindByStudentID(id string) ([]Evaluation, error)
	Delete(studentID string, criterionID string) error
}

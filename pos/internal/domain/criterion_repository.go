package domain

type CriterionRepository interface {
	Save(c Criterion) error
	FindByID(ID string) (Criterion, error)
	FindAll() ([]Criterion, error)
}

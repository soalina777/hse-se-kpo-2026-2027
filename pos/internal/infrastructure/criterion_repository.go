package infrastructure

import (
	"fmt"
	"pos/internal/domain"
)

type MemoryCriterionRepository struct {
	criterions map[string]domain.Criterion
}

func NewMemoryCriterionRepository() *MemoryCriterionRepository {
	return &MemoryCriterionRepository{
		criterions: make(map[string]domain.Criterion),
	}
}

func (cr *MemoryCriterionRepository) Save(criterion domain.Criterion) error {
	_, ok := cr.criterions[criterion.ID()]

	if ok {
		return fmt.Errorf(
			"criterion with ID %s already exists",
			criterion.ID(),
		)
	}
	cr.criterions[criterion.ID()] = criterion
	return nil
}

func (cr *MemoryCriterionRepository) FindByID(id string) (domain.Criterion, error) {
	criterion, ok := cr.criterions[id]

	if ok {
		return criterion, nil
	}

	return domain.Criterion{}, fmt.Errorf(
		"criterion with ID %s was not found",
		id,
	)
}

func (cr *MemoryCriterionRepository) FindAll() ([]domain.Criterion, error) {
	result := make([]domain.Criterion, 0, len(cr.criterions))

	for _, criterion := range cr.criterions {
		result = append(result, criterion)
	}

	return result, nil
}

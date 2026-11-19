package application

import (
	"fmt"
	"pos/internal/domain"
)

type ManagementService struct {
	studentRepo   domain.StudentRepository
	criterionRepo domain.CriterionRepository
}

func NewManagementService(
	studentRepo domain.StudentRepository,
	criterionRepo domain.CriterionRepository,
) *ManagementService {
	return &ManagementService{
		studentRepo:   studentRepo,
		criterionRepo: criterionRepo,
	}
}

func (ms *ManagementService) AddStudent(id string, name string) error {
	student, err := domain.NewStudent(id, name)

	if err != nil {
		return fmt.Errorf("failed to add student: %w", err)
	}

	err = ms.studentRepo.Save(student)

	if err != nil {
		return fmt.Errorf("failed to save student: %w", err)
	}

	return nil
}

func (ms *ManagementService) AddCriterion(
	id string,
	name string,
	maxScore int,
	weight float64,
) error {
	criterion, err := domain.NewCriterion(id, name, maxScore, weight)

	if err != nil {
		return fmt.Errorf("failed to add criterion: %w", err)
	}

	err = ms.criterionRepo.Save(criterion)

	if err != nil {
		return fmt.Errorf("failed to save criterion: %w", err)
	}

	return nil
}

func (ms *ManagementService) GetStudents() ([]domain.Student, error) {
	students, err := ms.studentRepo.FindAll()

	if err != nil {
		return nil, fmt.Errorf("failed to get all students: %w", err)
	}

	return students, nil
}

func (ms *ManagementService) GetCriterions() ([]domain.Criterion, error) {
	criterions, err := ms.criterionRepo.FindAll()

	if err != nil {
		return nil, fmt.Errorf("failed to get all criterions: %w", err)
	}

	return criterions, nil
}

package infrastructure

import (
	"fmt"
	"pos/internal/domain"
)

type MemoryStudentRepository struct {
	students map[string]domain.Student
}

func NewMemoryStudentRepository() *MemoryStudentRepository {
	return &MemoryStudentRepository{
		students: make(map[string]domain.Student),
	}
}

func (sr *MemoryStudentRepository) Save(student domain.Student) error {
	_, ok := sr.students[student.ID()]

	if ok {
		return fmt.Errorf(
			"student with ID %s already exists",
			student.ID(),
		)
	}

	sr.students[student.ID()] = student
	return nil
}

func (sr *MemoryStudentRepository) FindByID(id string) (domain.Student, error) {
	student, ok := sr.students[id]

	if ok {
		return student, nil
	}

	return domain.Student{}, fmt.Errorf(
		"student with ID %s was not found",
		id,
	)
}

func (sr *MemoryStudentRepository) FindAll() ([]domain.Student, error) {
	result := make([]domain.Student, 0, len(sr.students))

	for _, student := range sr.students {
		result = append(result, student)
	}

	return result, nil
}

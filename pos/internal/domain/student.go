package domain

import (
	"errors"
	"strings"
)

type Student struct {
	id   string
	name string
}

func NewStudent(id string, name string) (Student, error) {
	id = strings.TrimSpace(id)
	name = strings.TrimSpace(name)

	if id == "" {
		return Student{}, errors.New("ID cannot be empty")
	}

	if name == "" {
		return Student{}, errors.New("Name cannot be empty")
	}

	return Student{
		id:   id,
		name: name,
	}, nil
}

func (s Student) ID() string {
	return s.id
}

func (s Student) Name() string {
	return s.name
}

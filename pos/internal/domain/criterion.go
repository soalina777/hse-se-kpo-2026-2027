package domain

import (
	"errors"
	"math"
	"strings"
)

type Criterion struct {
	id       string
	name     string
	maxScore int
	weight   float64
}

func NewCriterion(
	id string,
	name string,
	maxScore int,
	weight float64,
) (Criterion, error) {
	id = strings.TrimSpace(id)
	name = strings.TrimSpace(name)

	if id == "" {
		return Criterion{}, errors.New("id cannot be empty")
	}

	if name == "" {
		return Criterion{}, errors.New("name cannot be empty")
	}

	if maxScore <= 0 {
		return Criterion{}, errors.New(
			"max Score must be positive",
		)
	}

	if math.IsNaN(weight) || math.IsInf(weight, 0) {
		return Criterion{}, errors.New("weight must be a finite number")
	}

	if weight < 0 || weight > 1 {
		return Criterion{}, errors.New(
			"weight must be between 0 and 1",
		)
	}

	return Criterion{
		id:       id,
		name:     name,
		maxScore: maxScore,
		weight:   weight,
	}, nil
}

func (c Criterion) ID() string {
	return c.id
}

func (c Criterion) Name() string {
	return c.name
}

func (c Criterion) MaxScore() int {
	return c.maxScore
}

func (c Criterion) Weight() float64 {
	return c.weight
}

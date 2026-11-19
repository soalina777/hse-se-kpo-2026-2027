package domain

type StudentRepository interface {
	Save(student Student) error
	FindByID(id string) (Student, error)
	FindAll() ([]Student, error)
}

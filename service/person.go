package service

import (
	"go-api/domain"
	"go-api/repository"
)

type PersonService interface {
	Save(p domain.Person) error
	Get() []domain.Person
}

type personService struct {
	r repository.PersonRepository
}

func NewPersonService(r repository.PersonRepository) PersonService {
	return personService{r: r}
}

func (s personService) Save(p domain.Person) error {
	return s.r.Save(p)
}

func (s personService) Get() []domain.Person {
	return s.r.Get()
}

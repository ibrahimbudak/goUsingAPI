package repository

import (
	"go-api/domain"
	"gorm.io/gorm"
)

type PersonRepository interface {
	Save(p domain.Person) error
	Get() []domain.Person
}

type personRepository struct {
	db *gorm.DB
}

func NewPersonRepository(db *gorm.DB) PersonRepository {
	err := db.AutoMigrate(&domain.Person{})
	if err != nil {
		return nil
	}

	return personRepository{db: db}
}

func (r personRepository) Save(p domain.Person) error {
	return r.db.Create(&p).Error
}

func (r personRepository) Get() []domain.Person {
	var p []domain.Person
	_ = r.db.Find(&p).Error
	return p
}

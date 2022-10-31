package repository

import (
	"fmt"
	"go-api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	crd *config.PostgresCredentials
}

func (p *Postgres) Connect() (*gorm.DB, error) {
	defaultDsn := "host=%s user=%s password=%s dbname=%s port=%d sslmode=disable"
	dsnSource := fmt.Sprintf(defaultDsn, p.crd.Host, p.crd.User, p.crd.Password, p.crd.DB, p.crd.Port)

	db, err := gorm.Open(postgres.Open(dsnSource), &gorm.Config{})

	return db, err
}

func NewPostgres(c *config.PostgresCredentials) *Postgres {
	return &Postgres{crd: c}
}

package config

import (
	"fmt"
	"go-api/resource"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Postgres *PostgresCredentials `yaml:"postgres"`
}

type PostgresCredentials struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DB       string `yaml:"db"`
	Port     int    `yaml:"port"`
}

func NewConfiguration() *Config {
	conf := &Config{}

	err := yaml.Unmarshal([]byte(resource.ApplicationConfigIn), conf)
	if err != nil {
		fmt.Println(err.Error())
	}

	return conf
}

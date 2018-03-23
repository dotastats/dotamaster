package config

import (
	"dotamaster/cmd"
	"sync"
)

var (
	conf Config
	once sync.Once
)

type Config struct {
	PostgreSQL PostgreSQL
}

type PostgreSQL struct {
	Username     string
	Password     string
	Host         string
	Port         int
	Db           string
	Debug        bool
	MaxIdleConns int
	MaxOpenConns int
}

func load() {
	once.Do(func() {
		if err := cmd.GetViper().Unmarshal(&conf); err != nil {
			panic(err)
		}
	})
}

func Load() Config {
	load()
	return conf
}

func Get() Config {
	load()
	return conf
}

func GetPostgreSQL() PostgreSQL { return conf.GetPostgreSQL() }
func (c Config) GetPostgreSQL() PostgreSQL {
	return c.PostgreSQL
}

package config

import (
	"master/cmd"
	"sync"
)

var (
	conf Config
	once sync.Once
)

type Config struct {
	App        App
	PostgreSQL PostgreSQL
	Log        Log
}

type App struct {
	Host      string
	Port      int
	Debug     bool
	Whitelist []string
}

type Log struct {
	Prefix     string
	Dir        string
	LevelDebug bool
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

func Load() {
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

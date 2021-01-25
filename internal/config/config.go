package config

import "github.com/BurntSushi/toml"

type Config struct {
	Logger  LoggerConf
	Storage StorageConf
}

type LoggerConf struct {
	Level string
	Path  string
}

type StorageConf struct {
	Dsn string
}

func NewConfig(path string) (*Config, error) {
	var conf Config
	if _, err := toml.DecodeFile(path, &conf); err != nil {
		return nil, err
	}

	return &conf, nil
}

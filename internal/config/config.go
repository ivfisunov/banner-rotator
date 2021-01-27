package config

import "github.com/BurntSushi/toml"

type Config struct {
	Env     string
	Logger  LoggerConf
	Storage StorageConf
	HTTP    HTTPConf
}

type LoggerConf struct {
	Level string
	Path  string
}

type StorageConf struct {
	Dsn string
}

type HTTPConf struct {
	Host string
	Port string
}

func NewConfig(path string) (*Config, error) {
	var conf Config
	if _, err := toml.DecodeFile(path, &conf); err != nil {
		return nil, err
	}

	return &conf, nil
}

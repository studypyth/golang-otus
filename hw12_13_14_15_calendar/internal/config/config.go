package config

// При желании конфигурацию можно вынести в internal/config.
// Организация конфига в main принуждает нас сужать API компонентов, использовать
// при их конструировании только необходимые параметры, а также уменьшает вероятность циклической зависимости.

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/rs/zerolog"
)

type Config struct {
	Logger  LoggerConf  `yaml:"logger"`
	Storage StorageConf `yaml:"storage"`
}

type LoggerConf struct {
	Level      zerolog.Level `yaml:"level"`
	Timeformat string        `yaml:"timeformat"`
	File       bool          `yaml:"file"`
	FilePath   string        `yaml:"file_path"`
}

type StorageConf struct {
	Type string        `yaml:"type"`
	Db   DbStorageConf `yaml:"db"`
}

type DbStorageConf struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	Db   string `yaml:"db"`
	User string `yaml:"user"`
	Pass string `yaml:"pass"`
}

func New(configFile string) Config {
	var cfg Config
	err := cleanenv.ReadConfig(configFile, &cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}

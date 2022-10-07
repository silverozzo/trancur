package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	errLog *log.Logger
}

func NewConfig(errLog *log.Logger) *Config {
	if err := godotenv.Load(); err != nil {
		errLog.Fatalln("не загружен файл окружения:", err)
	}

	cfg := &Config{
		errLog: errLog,
	}

	return cfg
}

func (cfg *Config) GetSelfHttpPort() string {
	return cfg.loadField("SELF_HTTP_PORT")
}

func (cfg *Config) loadField(fld string) string {
	val, ok := os.LookupEnv(fld)
	if !ok {
		cfg.errLog.Fatalln("в окружении нет поля:", fld)
	}

	return val
}

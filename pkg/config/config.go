package config

import (
	"os"
	"strconv"
)

type Config struct {
	RdbDsn  string
	WebAddr string
	Debug   bool
}

func NewEnvConfig() Config {
	cfg := Config{}

	cfg.RdbDsn = os.Getenv("RDB_DSN")
	cfg.WebAddr = os.Getenv("WEB_ADDR")
	cfg.Debug, _ = strconv.ParseBool(os.Getenv("DEBUG"))

	return cfg
}

package cfg

import (
	"errors"
	"os"

	"github.com/BurntSushi/toml"
)

type CanaryConfig struct {
	CanaryFileName string
	CanaryDocument string
	SmtpHost       string
	SmtpPort       int
	SmtpProto      string
	SmtpUser       string
	SmtpPass       string
	SmtpFrom       string
	SmtpSubject    string
}

func NewConfig(configPath string) (cfg *CanaryConfig, err error) {
	if _, err := os.Stat(configPath); err != nil {
		return nil, errors.New("config file does not exist")
	}

	cfg = &CanaryConfig{}
	_, err = toml.DecodeFile(configPath, cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

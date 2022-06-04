package cfg

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type CanaryConfig struct {
	ForceOverwrite bool
	CanaryFileName string
	CanaryDocument string
	SmtpHost       string
	SmtpPort       int
	SmtpProto      string
	SmtpUser       string
	SmtpPass       string
	SmtpFrom       string
	SmtpSubject    string
	SmtpTo         []string
}

func NewConfig(configPath string) (cfg *CanaryConfig, err error) {
	if _, err := os.Stat(configPath); err != nil {
		return nil, fmt.Errorf("config file does not exist at %s", configPath)
	}

	cfg = &CanaryConfig{}
	_, err = toml.DecodeFile(configPath, cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

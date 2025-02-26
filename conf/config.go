package conf

import (
	"fmt"

	"github.com/aranw/yamlcfg"
	"github.com/go-playground/validator/v10"
)

type WorkOs struct {
	ClientID   string `yaml:"client_id" validate:"required"`
	ApiKey     string `yaml:"api_key" validate:"required"`
	AuthkitURL string `yaml:"authkit_url" validate:"required"`
}

type HTTP struct {
	Address string `yaml:"address" validate:"required"`
}

type Config struct {
	WorkOs WorkOs `yaml:"workos" validate:"required"`
	HTTP   HTTP   `yaml:"http" validate:"required"`
}

func (c Config) Validate() error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	return validate.Struct(c)
}

func ParseConfig(path string) (*Config, error) {
	cfg, err := yamlcfg.Parse[Config](path)
	if err != nil {
		return nil, fmt.Errorf("parsing yaml config failed: %w", err)
	}

	return cfg, nil
}

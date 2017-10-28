package integrations

import (
	"errors"
)

type Config struct {
	Type    string
	Options struct {
		Url string
	}
}

func BuildIntegration(cfg Config) (intg Integration, e error) {
	switch cfg.Type {
	case "Slack":
		return Slack{cfg.Options.Url}, nil
	}

	return nil, errors.New("Invalid integration type " + cfg.Type)
}

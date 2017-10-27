package config

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Job struct {
	Name        string
	Source      string
	Destination string
}

type Config struct {
	Integrations struct {
		Slack struct {
			Url string
		}
	}
	Jobs  []Job
	Rsync struct {
		User   string `yaml:",omitempty"`
		Host   string `yaml:",omitempty"`
		Path   string `yaml:",omitempty"`
		Rsh    string `yaml:",omitempty"`
		Delete bool
	}
}

func ParseFile(path string) (c Config, e error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return Config{}, err
	}
	return Parse(content)
}

func Parse(source []byte) (c Config, err error) {
	result := Config{}

	err = yaml.UnmarshalStrict(source, &result)
	if err != nil {
		return result, err
	}

	for _, job := range result.Jobs {
		if job.Name == "" {
			return result, errors.New("all jobs must have names")
		}
		if job.Source == "" {
			return result, errors.New("all jobs must have sources")
		}
		if job.Destination == "" {
			return result, errors.New("all jobs must have destinations")
		}
	}

	return result, nil
}

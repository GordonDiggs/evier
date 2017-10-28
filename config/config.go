package config

import (
	"errors"
	"evier/integrations"
	"evier/job"
	"evier/rsync"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Slack integrations.Slack
	Jobs  []job.Job
	Rsync rsync.Options
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

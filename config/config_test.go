package config

import "testing"

func TestIntegrations(t *testing.T) {
	slack_url := "https://slack.com/webhooks"
	source := []byte(`integrations:
  - type: "slack"
    options:
      url: ` + slack_url)
	config, _ := Parse(source)
	if config.Integrations[0].Options.Url != slack_url {
		t.Error("Could not parse Slack URL")
	}
}

func TestJobs(t *testing.T) {
	source := []byte(`jobs:
  - name: "Photos"
    source: "/sup"
    destination: "/dest"
  `)
	config, _ := Parse(source)
	if len(config.Jobs) != 1 {
		t.Error("Did not parse the right number of jobs")
	}
}

func TestRsyncConfig(t *testing.T) {
	source := []byte(`rsync:
  user: arthur
  `)
	config, _ := Parse(source)
	if config.Rsync.User != "arthur" {
		t.Error("Did not parse correct user")
	}
	if config.Rsync.Delete != false {
		t.Error("Did not default delete properly")
	}
	if config.Rsync.Path != "" {
		t.Error("Did not default path properly")
	}
}

func TestJobError(t *testing.T) {
	source := []byte(`jobs:
  - name: "Photos"
  `)
	_, err := Parse(source)
	if err == nil {
		t.Error("Expected an error but got none")
	}
}

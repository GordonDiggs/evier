package integrations

import (
	"evier/job"
	"time"
)

type Slack struct {
	Url string
}

func (intg Slack) NotifyProcessStart(startTime time.Time) (e error) {
	return nil
}

func (intg Slack) NotifyJobStart(job job.Job, startTime time.Time) (e error) {
	return nil
}
func (intg Slack) NotifyJobSuccess(job job.Job, startTime time.Time, endTime time.Time) (e error) {
	return nil
}
func (intg Slack) NotifyJobFailure(job job.Job, startTime time.Time, endTime time.Time) (e error) {
	return nil
}
func (intg Slack) NotifyProcessSuccess(startTime time.Time, endTime time.Time) (e error) {
	return nil
}

// Verify that Slack implements the whole interface
var _ Integration = (*Slack)(nil)

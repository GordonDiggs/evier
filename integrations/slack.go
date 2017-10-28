package integrations

import (
	"evier/job"
	"fmt"
	"time"
)

type Slack struct {
	Url string
}

func (intg Slack) NotifyProcessStart(startTime time.Time) (e error) {
	fmt.Println("starting " + startTime.String())
	return nil
}

func (intg Slack) NotifyJobStart(job job.Job, startTime time.Time) (e error) {
	fmt.Println("starting " + job.Name + startTime.String())
	return nil
}
func (intg Slack) NotifyJobSuccess(job job.Job, startTime time.Time, endTime time.Time) (e error) {
	fmt.Println("finishing " + job.Name + startTime.String() + endTime.String())
	return nil
}
func (intg Slack) NotifyJobFailure(job job.Job, startTime time.Time, endTime time.Time) (e error) {
	fmt.Println("Failed " + job.Name + startTime.String() + endTime.String())
	return nil
}
func (intg Slack) NotifyProcessSuccess(startTime time.Time, endTime time.Time) (e error) {
	fmt.Println("finishing " + startTime.String() + endTime.String())
	return nil
}

// Verify that Slack implements the whole interface
var _ Integration = (*Slack)(nil)

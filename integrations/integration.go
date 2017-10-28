package integrations

import (
	"evier/job"
	"time"
)

type Integration interface {
	NotifyProcessStart(startTime time.Time) (e error)

	NotifyJobStart(job job.Job, startTime time.Time) (e error)

	NotifyJobSuccess(job job.Job, startTime time.Time, endTime time.Time) (e error)

	NotifyJobFailure(job job.Job, startTime time.Time, endTime time.Time) (e error)

	NotifyProcessSuccess(startTime time.Time, endTime time.Time) (e error)
}

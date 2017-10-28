package integrations

import (
	"evier/job"
	"time"
)

type IntegrationGroup struct {
	Integrations []Integration
}

func (ig IntegrationGroup) NotifyProcessStart(startTime time.Time) (e error) {
	for _, intg := range ig.Integrations {
		e = intg.NotifyProcessStart(startTime)
		if e != nil {
			return e
		}
	}
	return nil
}

func (ig IntegrationGroup) NotifyJobStart(job job.Job, startTime time.Time) (e error) {
	for _, intg := range ig.Integrations {
		e = intg.NotifyJobStart(job, startTime)
		if e != nil {
			return e
		}
	}
	return nil
}

func (ig IntegrationGroup) NotifyJobSuccess(job job.Job, startTime time.Time, endTime time.Time) (e error) {
	for _, intg := range ig.Integrations {
		e = intg.NotifyJobSuccess(job, startTime, endTime)
		if e != nil {
			return e
		}
	}
	return nil
}

func (ig IntegrationGroup) NotifyJobFailure(job job.Job, startTime time.Time, endTime time.Time) (e error) {
	for _, intg := range ig.Integrations {
		e = intg.NotifyJobFailure(job, startTime, endTime)
		if e != nil {
			return e
		}
	}
	return nil
}

func (ig IntegrationGroup) NotifyProcessSuccess(startTime time.Time, endTime time.Time) (e error) {
	for _, intg := range ig.Integrations {
		e = intg.NotifyProcessSuccess(startTime, endTime)
		if e != nil {
			return e
		}
	}
	return nil
}

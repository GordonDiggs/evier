package integrations

import (
	"evier/job"
	"time"
)

type IntegrationGroup struct {
	Configs []Config
}

func (ig IntegrationGroup) NotifyProcessStart(startTime time.Time) (e error) {
	for _, cfg := range ig.Configs {
		intg := IntegrationMap[cfg.Type]
		e = intg.NotifyProcessStart(startTime)
		if e != nil {
			return e
		}
	}
	return nil
}

func (ig IntegrationGroup) NotifyJobStart(job job.Job, startTime time.Time) (e error) {
	for _, cfg := range ig.Configs {
		intg := IntegrationMap[cfg.Type]
		e = intg.NotifyJobStart(job, startTime)
		if e != nil {
			return e
		}
	}
	return nil
}

func (ig IntegrationGroup) NotifyJobSuccess(job job.Job, startTime time.Time, endTime time.Time) (e error) {
	for _, cfg := range ig.Configs {
		intg := IntegrationMap[cfg.Type]
		e = intg.NotifyJobSuccess(job, startTime, endTime)
		if e != nil {
			return e
		}
	}
	return nil
}

func (ig IntegrationGroup) NotifyJobFailure(job job.Job, startTime time.Time, endTime time.Time) (e error) {
	for _, cfg := range ig.Configs {
		intg := IntegrationMap[cfg.Type]
		e = intg.NotifyJobFailure(job, startTime, endTime)
		if e != nil {
			return e
		}
	}
	return nil
}

func (ig IntegrationGroup) NotifyProcessSuccess(startTime time.Time, endTime time.Time) (e error) {
	for _, cfg := range ig.Configs {
		intg := IntegrationMap[cfg.Type]
		e = intg.NotifyProcessSuccess(startTime, endTime)
		if e != nil {
			return e
		}
	}
	return nil
}

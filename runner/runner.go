package runner

import (
	"evier/config"
	"evier/integrations"
	"time"
)

func Run(cfg config.Config) (e error) {
	rsyncOptions := cfg.Rsync
	intgs := integrations.IntegrationGroup{cfg.Integrations}

	startTime := time.Now()
	intgs.NotifyProcessStart(startTime)

	for _, job := range cfg.Jobs {
		jobStart := time.Now()

		e = intgs.NotifyJobStart(job, jobStart)
		if e != nil {
			return e
		}

		e = job.Perform(rsyncOptions)

		jobEnd := time.Now()

		if e == nil {
			e = intgs.NotifyJobSuccess(job, jobStart, jobEnd)
			if e != nil {
				return e
			}
		} else {
			err := intgs.NotifyJobFailure(job, jobStart, jobEnd)
			if err != nil {
				return err
			}

			return e
		}
	}

	intgs.NotifyProcessSuccess(startTime, time.Now())
	return nil
}

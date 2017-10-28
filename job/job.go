package job

import (
	"evier/rsync"
)

type Job struct {
	Name        string
	Source      string
	Destination string
}

func (j Job) Perform(rsyncOpts rsync.Options) (e error) {
	return nil
}

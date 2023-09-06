// This file contains the implementation of the job template type.

package awx

type JobTemplate struct {
	id               int
	name             string
	askLimitOnLaunch bool
	askVarsOnLaunch  bool
}

func (t *JobTemplate) Id() int {
	return t.id
}

func (t *JobTemplate) Name() string {
	return t.name
}

func (t *JobTemplate) AskLimitOnLaunch() bool {
	return t.askLimitOnLaunch
}

func (t *JobTemplate) AskVarsOnLaunch() bool {
	return t.askVarsOnLaunch
}

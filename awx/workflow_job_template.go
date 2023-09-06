// This file contains the implementation of the workflow job template type.

package awx

type WorkflowJobTemplate struct {
	id               int
	name             string
	askLimitOnLaunch bool
	askVarsOnLaunch  bool
}

func (t *WorkflowJobTemplate) Id() int {
	return t.id
}

func (t *WorkflowJobTemplate) Name() string {
	return t.name
}

func (t *WorkflowJobTemplate) AskLimitOnLaunch() bool {
	return t.askLimitOnLaunch
}

func (t *WorkflowJobTemplate) AskVarsOnLaunch() bool {
	return t.askVarsOnLaunch
}

// This file contains the data structures used for sending and receiving workflow job templates.

package data

type WorkflowJobTemplate struct {
	Id               int    `json:"id,omitempty"`
	Name             string `json:"name,omitempty"`
	AskLimitOnLaunch bool   `json:"ask_limit_on_launch,omitempty"`
	AskVarsOnLaunch  bool   `json:"ask_variables_on_launch,omitempty"`
}

type WorkflowJobTemplateGetResponse struct {
	WorkflowJobTemplate
}

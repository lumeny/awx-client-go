// This file contains the data structures used for sending and receiving job templates.

package data

type JobTemplate struct {
	Id               int    `json:"id,omitempty"`
	Name             string `json:"name,omitempty"`
	AskLimitOnLaunch bool   `json:"ask_limit_on_launch,omitempty"`
	AskVarsOnLaunch  bool   `json:"ask_variables_on_launch,omitempty"`
}

type JobTemplateGetResponse struct {
	JobTemplate
}

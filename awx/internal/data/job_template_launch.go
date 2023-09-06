// This file contains the data structures used to launch jobs from job templates.

package data

type JobTemplateLaunchGetResponse struct {
	JobTemplateData *JobTemplateGetResponse `json:"job_template_data,omitempty"`
}

type JobTemplateLaunchPostRequest struct {
	ExtraVars string `json:"extra_vars,omitempty"`
	Limit     string `json:"limit,omitempty"`
}

type JobTemplateLaunchPostResponse struct {
	Job int `json:"job,omitempty"`
}

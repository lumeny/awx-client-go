// This file contains the data structures used to launch jobs from workflow job templates.

package data

type WorkflowJobTemplateLaunchGetResponse struct {
	WorkflowJobTemplateData *WorkflowJobTemplateGetResponse `json:"workflow_job_template_data,omitempty"`
}

type WorkflowJobTemplateLaunchPostRequest struct {
	ExtraVars string `json:"extra_vars,omitempty"`
	Limit     string `json:"limit,omitempty"`
}

type WorkflowJobTemplateLaunchPostResponse struct {
	WorkflowJob int `json:"workflow_job,omitempty"`
}

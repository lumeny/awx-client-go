// This file contains the implementation of the workflow job template types.

package awx

type WorkflowJobTemplateLaunch struct {
	WorkflowJobTemplateData *WorkflowJobTemplate `json:"workflow_job_template_data,omitempty"`
}

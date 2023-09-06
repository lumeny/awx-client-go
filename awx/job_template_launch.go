// This file contains the implementation of the job template types.

package awx

type JobTemplateLaunch struct {
	JobTemplateData *JobTemplate `json:"job_template_data,omitempty"`
}

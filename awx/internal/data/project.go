// This file contains the data structures used for sending and receiving projects.

package data

type Project struct {
	Id        int    `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	SCMType   string `json:"scm_type,omitempty"`
	SCMURL    string `json:"scm_url,omitempty"`
	SCMBranch string `json:"scm_branch,omitempty"`
}

type ProjectGetResponse struct {
	Project
}

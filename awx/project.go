// This file contains the implementation of the project type.

package awx

// Project represents an AWX project.
type Project struct {
	id        int
	name      string
	scmType   string
	scmURL    string
	scmBranch string
}

// Id returns the unique identifier of the project.
func (p *Project) Id() int {
	return p.id
}

// Name returns the name of the project.
func (p *Project) Name() string {
	return p.name
}

// SCMType returns the source code management system type of the project.
func (p *Project) SCMType() string {
	return p.scmType
}

// SCMType returns the source code management system URL of the project.
func (p *Project) SCMURL() string {
	return p.scmURL
}

// SCMBranch returns the source code management system branch of the project.
func (p *Project) SCMBranch() string {
	return p.scmBranch
}

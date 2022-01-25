package openstack

// Project struct of an OpenStack project
type Project struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description"`
	// IsLocked is set to true when an ongoing process blocks the project from being modified
	IsLocked bool `json:"isLocked"`
	// IsBlockes is set to true when a project has been administratively blocked
	IsBlocked bool `json:"isBlocked"`
}

// projectWrapper struct contains a Project in it,
// this is solely used for unmarshalling/marshalling
type projectWrapper struct {
	Project Project `json:"project"`
}

// projectsWrapper struct contains a list of Projects in it,
// this is solely used for unmarshalling/marshalling
type projectsWrapper struct {
	Projects []Project `json:"projects"`
}

type handoverRequest struct {
	Action             string `json:"action"`
	TargetCustomerName string `json:"targetCustomerName"`
}

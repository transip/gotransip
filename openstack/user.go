package openstack

// User contains information about an OpenStack user
type User struct {
	ID          string `json:"id,omitempty"`
	Username    string `json:"username"`
	Description string `json:"description"`
	Email       string `json:"email"`
}

// CreateUserRequest is a struct used for creating OpenStack users
type CreateUserRequest struct {
	// ProjectID to grant the user access to
	ProjectID   string `json:"projectId"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Description string `json:"description"`
	Email       string `json:"email"`
}

type changePasswordRequest struct {
	NewPassword string `json:"newPassword"`
}

// userWrapper struct contains a User in it,
// this is solely used for unmarshalling/marshalling
type userWrapper struct {
	User User `json:"user"`
}

// usersWrapper struct contains a list of Projects in it,
// this is solely used for unmarshalling/marshalling
type usersWrapper struct {
	Users []User `json:"users"`
}

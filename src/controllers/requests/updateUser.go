package requests

// TODO: change password using sendgrid
type UpdateUser struct {
	ID   string `json:"-"`
	Name string `json:"name"`
}

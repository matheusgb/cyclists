package requests

type UpdateUser struct {
	ID   string `json:"-"`
	Name string `json:"name"`
}

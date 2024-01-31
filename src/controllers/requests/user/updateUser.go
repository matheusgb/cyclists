package requests

type UpdateUser struct {
	Name string `json:"name" validate:"min=3,max=255"`
}

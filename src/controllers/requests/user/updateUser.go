package requests

type UpdateUser struct {
	Name string `json:"name" validate:"required,min=3,max=255"`
}

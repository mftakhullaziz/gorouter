package payload

type CategoryRequest struct {
	Name string `validate:"required" json:"name"`
	Desc string `validate:"required" json:"desc"`
}

package payload

type CategoryUpdateRequest struct {
	Id   int    `validate:"required" json:"id"`
	Name string `validate:"required" json:"name"`
	Desc string `validate:"required" json:"desc"`
}

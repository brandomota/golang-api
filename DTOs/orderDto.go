package DTOs

type OrderDto struct {
	Status string `json:"status" xml:"status" form:"status" validate:""`
}

package DTOs

type ProductDto struct {
	Name string `json:"name" xml:"name" form:"name" validate:"min=1,max=100"`
	Description string `json:"description" xml:"description" form:"description" validate:"min=1,max=255"`
	ImageUrl string `json:"imageUrl" xml:"imageUrl" form:"imageUrl" validate:"min=1,max=100"`
	Price float64 `json:"price" xml:"price" form:"price" validate:"nonzero"`
}

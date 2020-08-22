package DTOs

type UserDto struct {
	Name     string `json:"name" xml:"name" form:"name" validate:"min=1,max=100"`
	Email    string `json:"email" xml:"email" form:"email" validate:"min=1,max=100"`
	Age      int64  `json:"age" xml:"age" form:"age" validate:"nonzero"`
	UserType int64  `json:"userType" xml:"userType" form:"userType" validate:"nonzero"`
}

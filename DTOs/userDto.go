package DTOs

type UserDto struct {
	Name     string `json:"name" xml:"name" form:"name"`
	Email    string `json:"email" xml:"email" form:"email"`
	Age      int64  `json:"age" xml:"age" form:"age"`
	UserType int64  `json:"userType" xml:"userType" form:"userType"`
}

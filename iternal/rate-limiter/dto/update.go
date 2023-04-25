package dto

type UpdateDTO struct {
	Request int `form:"req" binding:"omitempty,min=1,max=2147483647"`
	Second  int `form:"sec" binding:"omitempty,min=1,max=2147483647"`
}

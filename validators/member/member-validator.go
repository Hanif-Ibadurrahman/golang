package member

type UpdateMemberValidator struct {
	ID    uint64 `json:"id" form:"id"`
	Name  string `json:"name" form:"name" binding:"required"`
	Email string `json:"email" form:"email" binding:"required,email"`
}

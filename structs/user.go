package structs

type UserResponse struct {
	Id        uint    `json:"id"`
	Name      string  `json:"name"`
	Username  string  `json:"username"`
	Email     string  `json:"email"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
	Token     *string `json:"token,omiempty"`
}

type UserCreateRequest struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required" gorm:"unique;not null"`
	Email    string `json:"email" binding:"required" gorm:"unique:not null"`
	Password string `json:"password" binding:"required" gorm:"unique;not mull"`
}

type UserUpdateRequest struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required" gorm:"unique;not null"`
	Email    string `json:"email" binding:"required" gorm:"unique:not null"`
	Password string `json:"password,omiempty"`
}

type UserLoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

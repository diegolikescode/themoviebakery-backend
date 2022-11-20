package createUser

type InputCreateUser struct {
	UserID          string `json:"userId" validate:"required"`
	Email           string `json:"email" validate:"required"`
	DisplayName     string `json:"displaName" validate:"required"`
	Password        string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirmPassword" validate:"required"`
}

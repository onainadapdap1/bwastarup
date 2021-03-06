package user

//1. deklarasi cetakan RegisterUserInput, struct yg dipakai untuk mapping inputan user
type RegisterUserInput struct {
	Name       string `json:"name" binding:"required"`
	Occupation string `json:"occupation" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
}

//2. deklarasi cetakan LoginInput , untuk mapping inputan user
type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

//3. deklarasi cetakan CheckEmailInput
type CheckEmailInput struct {
	Email string `json:"email" binding:"required,email"`
}

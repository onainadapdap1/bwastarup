package user

//1. deklarasi cetakan UserFormatter
type UserFormatter struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Occupation string `json:"occupation"`
	Email      string `json:"email"`
	Token      string `json:"token"`
}

//2. fungsi untuk memformat kembalian json
func FormatUser(user User, token string) UserFormatter {
	//-> set objek cetakan UserFormatter
	formatter := UserFormatter{
		ID:         user.ID,
		Name:       user.Name,
		Occupation: user.Occupation,
		Email:      user.Email,
		Token:      token,
	}

	return formatter
}

package user

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

//1. mewakili bisnis logic/ fitur yang ada di aplikasi
//-- mapping struct input ke struct User
type Service interface {
	RegisterUser(input RegisterUserInput) (User, error) //method dgn param = struct RegisterUserInput
	LoginUser(input LoginInput) (User, error)
	IsEmailAvailable(input CheckEmailInput) (bool, error)
	SaveAvatar(ID int, fileLocation string) (User, error) //fileLocation = alamat file disimpan
	GetUserByID(ID int) (User, error)
}

//2. deklarasi cetakan service
type service struct {
	repository Repository //has method Save
}

//3. membuat objek struct service, butuh fungsi yg return objek service
func NewService(repository Repository) *service {
	return &service{repository: repository}
}

//4. set implementasi method RegisterUser(), milik struct service
func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	//5. mapping struct input ke struct User
	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	//6. cek jika ada error
	if err != nil {
		return user, err
	}
	user.PasswordHash = string(passwordHash)
	user.Role = "user"

	//7. memanggil repository, untuk menyimpan user ke database
	newUser, err := s.repository.Save(user)

	//8. cek jika ada error
	if err != nil {
		return user, err
	}

	//9. return newUser, nil
	return newUser, nil
}

//--- implementasi method Login()
func (s *service) LoginUser(input LoginInput) (User, error) {
	// mendapatkan email dan password user
	email := input.Email
	password := input.Password

	// mencari user dengan alamat email input
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("No user found on that email")
	}

	//matching password
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return user, err
	}
	return user, nil
}

//-- set implementasi method isEmailAvailable()
func (s *service) IsEmailAvailable(input CheckEmailInput) (bool, error) {
	//tangkap input user
	email := input.Email
	//kirim ke repository
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return false, err
	}
	if user.ID == 0 {
		return true, nil
	}

	return false, nil //kasih nilai default
}

//--implement SaveAvatar(,)
func (s *service) SaveAvatar(ID int, fileLocation string) (User, error) {
	//dapatkan user id
	user, err := s.repository.FindById(ID)
	if err != nil {
		return user, err
	}
	//update atribute avatar filename -> memberikan nilai baru
	user.AvatarFileName = fileLocation
	//simpan perubahan filename ke db
	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return updatedUser, err
	}
	return updatedUser, nil
}

//-- implement GetByID()
func (s *service) GetUserByID(ID int) (User, error) {
	//--panggil FIndByID() dari repository
	user, err := s.repository.FindById(ID)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("No user found on with that ID")
	}
	return user, nil
}

// service = mapping struct input ke struct user
// simpan struct User melalui repository

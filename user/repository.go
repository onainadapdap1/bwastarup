package user

import "gorm.io/gorm"

//1. set interface public -> objek lain mengacu ke repository ini
type Repository interface {
	Save(user User) (User, error) //2. return User struct
}

//3. deklarasi cetakan repository struct private
type repository struct {
	db *gorm.DB //nilai awal = kosong
}

//4. set function NewRepository(), untuk membuat sebuah objek cetakan repository dgn nilai db sudah berisi
func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

//5. set implementasi method save, milik struct repository
func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error
	//6. cek jika ada error
	if err != nil {
		return user, err
	}
	//7. jika tidak ada error
	return user, nil
}

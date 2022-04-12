package user

import "gorm.io/gorm"

//1. set interface public -> objek lain mengacu ke repository ini
type Repository interface {
	Save(user User) (User, error)           //2. return User struct
	FindByEmail(email string) (User, error) //6. membuat fungsi yg bisa mencari user yang memiliki alamat email X
	FindById(ID int) (User, error)          //8. FindById()
	Update(user User) (User, error)         //10
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

//7. set implementasi method FindByEmail()
func (r *repository) FindByEmail(email string) (User, error) {
	//set objek cetakan User
	var user User
	//-> find email, jika ada yg dikembalikan, simpan ke dalam objek user
	err := r.db.Where("email = ?", email).Find(&user).Error
	//-> cek jika ada error
	if err != nil {
		return user, err
	}

	return user, nil
}

//9. set implementasi method FindById()
func (r *repository) FindById(ID int) (User, error) {
	//-- set objek penampung cetakan User
	var user User
	//-- find ID, jika ada, simpan ke objek user
	err := r.db.Where("id = ?", ID).Find(&user).Error
	//--jika ada error, return error
	if err != nil {
		return user, err
	}
	return user, nil
}

//11. implement kontrak Update()
func (r *repository) Update(user User) (User, error) {
	//-- update data user di db
	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

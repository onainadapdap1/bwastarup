package main

import (
	"bwastarup/handler"
	"bwastarup/user"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	//1. koneksi database
	dsn := "root:root@tcp(127.0.0.1:3306)/bwastarup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	//2. cek jika error
	if err != nil {
		log.Fatal(err.Error())
	}

	//3. memanggil fungsi NewRepository dari repository.go
	userRepository := user.NewRepository(db) //userRepository = objek berisi field db
	//6. memanggil fungsi RegisterUser dari service.go
	userService := user.NewService(userRepository)
	//9. memanggil fungsi RegisterUser
	userHandler := handler.NewUserHandler(userService)

	//17. memanggil method IsEmailAvailable() dari service.go
	//input := user.CheckEmailInput{
	//	Email: "contoh@gmail.com",
	//}
	//isTrue, err := userService.IsEmailAvailable(input)
	//if err != nil {
	//	fmt.Println("Sudah ada email di db")
	//	fmt.Println(err.Error())
	//}
	//fmt.Println("is True :", isTrue)

	//15. memanggil method Login() dari service.go
	//input := user.LoginInput{
	//	Email:    "contoh@gmail.com",
	//	Password: "password",
	//}
	//user, err := userService.Login(input)
	//if err != nil {
	//	fmt.Println("Terjadi kesalahan")
	//	fmt.Println(err.Error())
	//}
	//fmt.Println(user.Email)
	//fmt.Println(user.Name)

	//14. memanggil method FindByEmail dari repository.go
	//userByEmail, err := userRepository.FindByEmail("pegasus@gmail.com")
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//if userByEmail.ID == 0 {
	//	fmt.Println("User tidak ditemukan")
	//} else {
	//	fmt.Println(userByEmail.Name)
	//}

	//10. set default router
	router := gin.Default()
	//11. set grouping versioning
	api := router.Group("/api/v1")

	//12. set handler register
	api.POST("/users", userHandler.RegisterUser)
	//16. set handler login
	api.POST("/sessions", userHandler.Login)
	//18. set handler checkEmailAvailability
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	//13. menjalankan server
	router.Run()
	//7. set objek cetakan RegisterUserInput
	//userInput := user.RegisterUserInput{
	//	Name:       "Tes simpan dari service",
	//	Email:      "contoh@gmail.com",
	//	Occupation: "anak band",
	//	Password:   "password",
	//}

	//8. menyimpan user dengan userService
	//userService.RegisterUser(userInput)

	//4. set objek cetakan User dari user.go
	//user := user.User{
	//	Name: "Test save user",
	//}

	//5. menyimpan data user ke database
	//userRepository.Save(user)
}

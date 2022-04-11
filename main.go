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

	//10. set default router
	router := gin.Default()
	//11. set grouping versioning
	api := router.Group("/api/v1")

	//12. set handler
	api.POST("/users", userHandler.RegisterUser)

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

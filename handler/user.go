package handler

import (
	"bwastarup/auth"
	"bwastarup/helper"
	"bwastarup/user"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//1. deklarasi cetakan userHandler, service menjadi dependency handle
type userHandler struct {
	userService user.Service
	authService auth.Service
}

//2. fungsi ini, return objek struct userHandler
func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService: userService, authService: authService}
}

//3. method handler RegisterUser()
func (h *userHandler) RegisterUser(c *gin.Context) {
	//tangkap input dari user
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input) //mendapatkan objek input dari user, simpan ke objek input
	if err != nil {
		errors := helper.FormatValidationError(err)
		//set response error message
		errorMessage := gin.H{"errors": errors}

		//set response json
		response := helper.APIResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	//memasukkan input ke dalam service
	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	//--generate token ketikan register
	token, err := h.authService.GenerateToken(newUser.ID)
	if err != nil {
		response := helper.APIResponse("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	//memanggil formatter dari formatter.go, seteleh mengirim data ke userService
	formatter := user.FormatUser(newUser, token)

	//-> memanggil fungsi APIResponse() dari helper.go
	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
	// map input dari user ke struct RegisterUserInput
	// struct di atas kita passing sebagai parameter service -> kemudian ke repository
}

//4. set method Login(), milik struct userHandler
func (h *userHandler) Login(c *gin.Context) {
	//-> user memasukkan input
	//-> input ditangkap handler
	//-> mapping dari input user ke input struct
	//-> di service mencari dgn bantuan repo user dengan email x
	//-> matching password

	//-> set objek cetakan LoginInput
	var input user.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		//set response json
		response := helper.APIResponse("Login failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	//memasukkan inputan ke dalam service
	loggedInUser, err := h.userService.LoginUser(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Login gagal", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	//--generate new token ketika berhasil login
	token, err := h.authService.GenerateToken(loggedInUser.ID)
	if err != nil {
		response := helper.APIResponse("Login gagal", http.StatusUnprocessableEntity, "error", nil)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	//-- jika berhasil validasi dan mengrim ke service, set response json
	formatter := user.FormatUser(loggedInUser, token)
	response := helper.APIResponse("Successfuly loggedin", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
}

//5. set fungsi CheckEmailAvailability(c *gin.Context)
func (h *userHandler) CheckEmailAvailability(c *gin.Context) {
	//-- ada input email dari user
	//--input email di mapping ke struct input
	//-- struct input di mapping ke service
	//-- service akan memanggil repository, untuk cek email sudah ada atau belum
	//-- repository mengirim ke database

	//menangkap inputan user
	var input user.CheckEmailInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		//set response json
		response := helper.APIResponse("Email checking failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	//memanggil service, pass nilai inputan user
	isEmailAvailable, err := h.userService.IsEmailAvailable(input)
	if err != nil {
		errorMessage := gin.H{"errors": "Server error"}
		//set response json
		response := helper.APIResponse("Email checking failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{
		"is_available": isEmailAvailable,
	}

	metaMessage := "Email has been registered"
	if isEmailAvailable {
		metaMessage = "Email is available"
	}

	response := helper.APIResponse(metaMessage, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}

//6. set method UploadAvatar()
func (h *userHandler) UploadAvatar(c *gin.Context) {
	//-- menangkap file inputan user
	file, err := c.FormFile("avatar") //json:avatar
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)
		//set response
		c.JSON(http.StatusBadRequest, response)
		return
	}
	//harusnya dari jwt
	//userID := 1 //anggap dari jwt
	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID
	// images/namafile.png ->images/ID-namafile.png
	path := "images/" + file.Filename
	path = fmt.Sprintf("images/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar image", http.StatusBadRequest, "error", data)
		//set response
		c.JSON(http.StatusBadRequest, response)
		return
	}

	//call service, untuk save avatar
	_, err = h.userService.SaveAvatar(userID, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed to upload avatar image broo0", http.StatusBadRequest, "error", data)
		//set response
		c.JSON(http.StatusBadRequest, response)
		return
	}
	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse("Avatar successfuly uploaded", http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
	//c.SaveUploadedFile(file, )
	//--input dari user -> berupa form data
	//--simpan gambar di folder "/images"
	//--di service memanggil repo
	//JWT (sementara di hardcode, seakan2 yg login ID = 1 )
	//--repo ambil data user id 1
	//--repo update data user, simpan lokasi file
}

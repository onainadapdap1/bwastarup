package handler

import (
	"bwastarup/helper"
	"bwastarup/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

//1. deklarasi cetakan userHandler, service menjadi dependency handle
type userHandler struct {
	userService user.Service
}

//2. fungsi ini, return objek struct userHandler
func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService: userService}
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

	//memanggil formatter dari formatter.go, seteleh mengirim data ke userService
	formatter := user.FormatUser(newUser, "tokentokentokentoken")

	//-> memanggil fungsi APIResponse() dari helper.go
	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
	// map input dari user ke struct RegisterUserInput
	// struct di atas kita passing sebagai parameter service -> kemudian ke repository
}

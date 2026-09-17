package controllers

import (
	"AuthInGo/dto"
	"AuthInGo/services"
	"AuthInGo/utils"
	"fmt"
	"net/http"
	
)

type UserController struct {
	UserService  services.UserService
}

func NewUserController(_userService services.UserService) *UserController{
	return &UserController{
		UserService: _userService,
	}
}

func (uc *UserController) GetUserById(w http.ResponseWriter, r *http.Request){
	fmt.Println("user controller")

	UserId := r.Context().Value("UserId").(int64)

	id := int64(UserId)


	uc.UserService.GetUserById(id)
	w.Write([]byte("User registration end point"))
}

func (uc *UserController) GetAllUser(w http.ResponseWriter, r *http.Request){
	fmt.Println("user controller")
	uc.UserService.GetAllUser()
	w.Write([]byte("User registration end point"))
}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request){

	payload := r.Context().Value("payload").(dto.CreateUserRequestDto)

	user,err := uc.UserService.CreateUser(&payload)

	if err != nil {
		utils.WriteJsonErrorsResponse(w,http.StatusInternalServerError,"something went wrong while Creating user in",err)
		return
	}

	utils.WriteJsonSuccessResponse(w,http.StatusOK,"User Signup  successfully",user)

}

func (uc *UserController) LoginUser(w http.ResponseWriter, r *http.Request){
	fmt.Println("LoginUser called in from controller")

	payload := r.Context().Value("payload").(dto.LoginUserRequestDto)

	jwtToken,err := uc.UserService.LoginUser(&payload)

	if err != nil {
		utils.WriteJsonErrorsResponse(w,http.StatusInternalServerError,"something went wrong while logging in",err)
		return
	}

	utils.WriteJsonSuccessResponse(w,http.StatusOK,"User logged in successfully",jwtToken)
	
}
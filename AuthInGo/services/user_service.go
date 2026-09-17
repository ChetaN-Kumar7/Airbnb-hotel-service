package services

import (
	env "AuthInGo/config/env"
	db "AuthInGo/db/repositories"
	"AuthInGo/dto"
	"AuthInGo/models"
	"AuthInGo/utils"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	
)

type UserService interface {
	GetUserById (id int64) (*models.User,error)
	CreateUser (payload *dto.CreateUserRequestDto) (*models.User,error)
	LoginUser (payload *dto.LoginUserRequestDto) (string,error)
	GetAllUser () error
}

type UserServiceImpl struct {
	userRepository db.UserRepository
}

func NewUserService(_userRepository db.UserRepository) UserService{

	return &UserServiceImpl{
		userRepository : _userRepository,
	}

}

func(u *UserServiceImpl) GetUserById(id int64) (*models.User,error){
	fmt.Println("Fetching user by id")
	user,err := u.userRepository.GetById(id)

	if err != nil{
		fmt.Println("Error fetching user:",err)
		return nil,err
	}

	return user,nil
}

func(u *UserServiceImpl) GetAllUser() error{
	fmt.Println("Fetching user by id")
	u.userRepository.GetAll()
	return  nil
}


func(u *UserServiceImpl) CreateUser(payload *dto.CreateUserRequestDto) (*models.User,error){

	hashPassword, err := utils.HashPassword(payload.Password)


	if err != nil {
		fmt.Println("Error in creating hash password",err)
		return nil,err
	}
	user,err:=u.userRepository.Create(payload.Username,payload.Email,hashPassword)

	if err != nil {
		fmt.Println("Error in creating User",err)
		return nil,err
	}	

	return  user,nil
}

func(u *UserServiceImpl) LoginUser(payload *dto.LoginUserRequestDto) (string,error){
	email := payload.Email
	password := payload.Password


	user,err := u.userRepository.GetByEmail(email)

	if err != nil{
		fmt.Println("Error fetching user by email",err)
		return "",err
	}

	if user == nil {
		fmt.Println("No user is found with the given email")
		return "",fmt.Errorf("no user found with the given email%s",email)
	}
	isPasswordValid := utils.CheckPasswordHash(user.Password,password)

	if !isPasswordValid {
		fmt.Println("Password does not match")
		return "",nil
	}

	Jwtpayload := jwt.MapClaims{
		"email" : user.Email,
		"id" :  user.Id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,Jwtpayload)

	tokenString, Jwterr := token.SignedString([]byte(env.GetString("JWT_SECRET","TOKEN")))

	if Jwterr != nil{
		fmt.Println("Error signing token",Jwterr)
		return "",Jwterr
	}

	fmt.Println("JWT Token:",tokenString)

	return  tokenString,nil
}
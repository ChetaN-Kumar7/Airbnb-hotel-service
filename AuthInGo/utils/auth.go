package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(plainPassword string) (string,error){
	hashPassword,err := bcrypt.GenerateFromPassword([]byte(plainPassword),bcrypt.DefaultCost)

	if err != nil {
		fmt.Println("Error creating hash password",err)
		return "",err
	}

	return string(hashPassword),nil
}

func CheckPasswordHash(hashedPassword,plainpassword string) bool{
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword),[]byte(plainpassword) )
	
	return err == nil
}
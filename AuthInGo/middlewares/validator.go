package middlewares

import (
	"AuthInGo/dto"
	"AuthInGo/utils"
	"context"
	"net/http"
)



func UserLoginRequestValidator(next http.Handler) http.Handler{

	return  http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload dto.LoginUserRequestDto

		if jsonErr := utils.ReadJsonBody(r,&payload);jsonErr != nil {
			utils.WriteJsonErrorsResponse(w,http.StatusInternalServerError,"something went wrong while logging in",jsonErr)
			return 
		}

		if validationErr := utils.Validator.Struct(payload); validationErr != nil{
			utils.WriteJsonErrorsResponse(w,http.StatusInternalServerError,"something went wrong while logging in",validationErr)
			return 
		}
		ctx := context.WithValue(r.Context(),"payload",payload)

		next.ServeHTTP(w,r.WithContext(ctx))
		})

}


func UserCreateRequestValidator(next http.Handler) http.Handler{

	return  http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload dto.CreateUserRequestDto

		if jsonErr := utils.ReadJsonBody(r,&payload);jsonErr != nil {
			utils.WriteJsonErrorsResponse(w,http.StatusInternalServerError,"something went wrong while logging in",jsonErr)
			return 
		}

		if validationErr := utils.Validator.Struct(payload); validationErr != nil{
			utils.WriteJsonErrorsResponse(w,http.StatusInternalServerError,"something went wrong while logging in",validationErr)
			return 
		}

		ctx := context.WithValue(r.Context(),"payload",payload)

		next.ServeHTTP(w,r.WithContext(ctx))
	})

}
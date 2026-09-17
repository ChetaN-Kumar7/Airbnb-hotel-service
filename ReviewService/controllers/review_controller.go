package controllers

import (
	"ReviewService/services"
	"fmt"
	"net/http"
)


type ReviewController struct{
	ReviewService services.ReviewService
}

func NewReviewController(_reviewService services.ReviewService) *ReviewController{
	return &ReviewController{
		ReviewService: _reviewService,
	}
}

func (rc *ReviewController) Create(w http.ResponseWriter, r *http.Request){
	fmt.Println("Calling from controller")
	rc.ReviewService.Create(1,12,123,"Good hotel",4)
	w.Write([]byte("User registration end point"))
}

func (rc *ReviewController) GetByID(w http.ResponseWriter, r *http.Request){
	fmt.Println("Calling from Controller")
	rc.ReviewService.GetByID(1)
	w.Write([]byte("User registration end point"))
}

func (rc *ReviewController) GetAll(w http.ResponseWriter, r *http.Request){
	fmt.Println("Calling from Controller")
	rc.ReviewService.GetAll()
	w.Write([]byte("User registration end point"))
}
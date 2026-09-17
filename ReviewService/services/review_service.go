package services

import (
	db "ReviewService/db/repositories"
	"fmt"
)

type ReviewService interface{
	Create (userId int64, bookingId int64, hotelId int64, comment string, rating int) error
	GetByID (id int64) error
	GetAll () error
}

type ReviewServiceImpl struct{
	reviewRepository db.ReviewRepository
}

func NewReviewService(_reviewRepository db.ReviewRepository) ReviewService{
	return &ReviewServiceImpl{
		reviewRepository: _reviewRepository,
	}
}

func (r *ReviewServiceImpl) Create(userId int64, bookingId int64, hotelId int64, comment string, rating int) error {
	fmt.Println("Creating user in Service Layer")
	r.reviewRepository.Create(userId,bookingId,hotelId,comment,rating)
	return nil
}

func (r *ReviewServiceImpl) GetByID(id int64) error{
	fmt.Println("Fetching review by id in Service Layer")
	r.reviewRepository.GetByID(id)
	return nil
}

func (r *ReviewServiceImpl) GetAll() error{
	fmt.Println("Fetching all review")
	r.reviewRepository.GetAll()
	return nil
}
package db

import (
	"ReviewService/models"
	"database/sql"
	"fmt"
)

type ReviewRepository interface{
	Create (userId int64, bookingId int64, hotelId int64, comment string, rating int) (*models.Review,error)
	GetByID(id int64) (*models.Review, error)
	GetAll() ([]*models.Review, error)
}

type ReviewRepositoryImpl struct{
	db *sql.DB
}

func NewReviewRepository(_db *sql.DB) ReviewRepository{
	return &ReviewRepositoryImpl{
		db:_db,
	}
}

func(r *ReviewRepositoryImpl) Create(userId int64, bookingId int64, hotelId int64, comment string, rating int) (*models.Review,error){

	query := "INSERT INTO reviews(user_id, booking_id, hotel_id, comment, rating) VALUES (?, ?, ?, ?, ?)"

	res,err:=r.db.Exec(query,userId,bookingId,hotelId,comment,rating)

	if err != nil{
		fmt.Println("Error inserting users:",err)
		return nil,err
	}

	lastInsertID, rowErr := res.LastInsertId()
	if rowErr != nil {
		fmt.Println("Error getting last insert ID:", rowErr)
		return nil, rowErr
	}

	review := &models.Review{
		Id:        lastInsertID,
		UserId:    userId,
		BookingId: bookingId,
		HotelId:   hotelId,
		Comment:   comment,
		Rating:    rating,
		IsSynced:  false,
	}

	fmt.Println("Review created successfully:", review)

	return review, nil

}


func(r *ReviewRepositoryImpl) GetByID(id int64) (*models.Review, error){
	query := "SELECT id, user_id, booking_id, hotel_id, comment, rating, created_at, updated_at, deleted_at, is_synced FROM reviews Where id=?"

	row := r.db.QueryRow(query, id)

	//step 3 process the result

	review := &models.Review{}

	err := row.Scan(&review.Id,&review.UserId, &review.BookingId, &review.HotelId, &review.Comment, &review.Rating, &review.CreatedAt, &review.UpdatedAt, &review.DeletedAt, &review.IsSynced)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No user found with the given ID")
			return nil, err
		} else {
			fmt.Println("Error Scanning user:", err)
			return nil, err
		}
	}

	// step 4 print the user details
	fmt.Println("User fetch successfully", review)

	return review, nil
}


func(r *ReviewRepositoryImpl) GetAll() ([]*models.Review,error){
	query := "SELECT id, user_id, booking_id, hotel_id, comment, rating, created_at, updated_at, deleted_at, is_synced FROM reviews"

	//step 2 execute the query

	rows,err := r.db.Query(query)

	if err != nil {
        return nil, err
    }
    defer rows.Close()

	//step 3 process the result

	var reviews [] *models.Review

	for rows.Next() {
		review := &models.Review{}
		if err := rows.Scan(&review.Id,&review.UserId, &review.BookingId, &review.HotelId, &review.Comment, &review.Rating, &review.CreatedAt, &review.UpdatedAt, &review.DeletedAt, &review.IsSynced); err != nil{
			return nil,err
		}
		reviews = append(reviews,review)

	}

	if err = rows.Err(); err != nil {
        return reviews, err
    }

	fmt.Println("All reviews:",reviews)

    return reviews, nil
}

func (r *ReviewRepositoryImpl) Delete(id int64) error {
	query := "UPDATE reviews SET deleted_at = CURRENT_TIMESTAMP WHERE id = ? AND deleted_at IS NULL"
	result, err := r.db.Exec(query, id)

	if err != nil {
		fmt.Println("Error deleting review:", err)
		return err
	}

	rowsAffected, rowErr := result.RowsAffected()
	if rowErr != nil {
		fmt.Println("Error getting rows affected:", rowErr)
		return rowErr
	}
	if rowsAffected == 0 {
		fmt.Println("No rows were affected, review not found or already deleted")
		return fmt.Errorf("review not found")
	}
	fmt.Println("Review deleted successfully, rows affected:", rowsAffected)
	return nil
}
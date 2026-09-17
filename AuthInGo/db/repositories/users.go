package db

import (
	"AuthInGo/models"
	"database/sql"
	"fmt"
)

type UserRepository interface {
	GetById(id int64) (*models.User, error)
	Create(username string,email string,password string)  (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	GetAll() ([]*models.User, error)
	DeleteById(id int64) error
}

type UserRepositoryImpl struct {
	db *sql.DB
}

func (u *UserRepositoryImpl) Create(username string,email string,password string) (*models.User, error){

	query := "INSERT INTO users (username,email,password) VALUES (?,?,?)"
	res,err := u.db.Exec(query,username,email,password)

	if err != nil {
		fmt.Println("Error inserting users:",err)
		return nil,err
	}

	lastInsertID, rowErr := res.LastInsertId()
	if rowErr != nil {
		fmt.Println("Error getting last insert ID:", rowErr)
		return nil, rowErr
	}

	user := &models.User{
		Id:       lastInsertID,
		Username: username,
		Email:    email,
	}

	fmt.Println("User created successfully:", user)

	return user, nil
}

func NewUserRepository(_db *sql.DB) UserRepository {
	return &UserRepositoryImpl{
		db: _db,
	}
}

func (u *UserRepositoryImpl) GetById(id int64) (*models.User, error) {
	fmt.Println("Creating user in user repository")
	//step 1 prepare the query
	query := "SELECT id,username,email,created_at,updated_at FROM users WHERE id =?"

	//step 2 execute the query

	row := u.db.QueryRow(query, id)

	//step 3 process the result

	user := &models.User{}

	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.Created_at, &user.Updated_at)

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
	fmt.Println("User fetch successfully", user)

	return user, nil
}

func (u *UserRepositoryImpl) GetByEmail(email string) (*models.User, error) {
	
	//step 1 prepare the query
	query := "SELECT id,username,email,password,created_at,updated_at FROM users WHERE email =?"

	//step 2 execute the query

	row := u.db.QueryRow(query, email)

	//step 3 process the result

	user := &models.User{}

	err := row.Scan(&user.Id, &user.Username, &user.Email,&user.Password,&user.Created_at, &user.Updated_at)

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
	fmt.Println("User fetch successfully", user)

	return user, nil
}

func (u *UserRepositoryImpl) DeleteById(id int64) error{
	query := "DELETE FROM USER WHERE id=?"

	row,err := u.db.Exec(query,id)
	
	if err != nil{
		fmt.Println("Failed to delete the user")
		return err
	}

	rowsAffected,rowErr := row.RowsAffected()
	if rowErr != nil {
		fmt.Println("Error getting affected row",rowErr)
		return rowErr
	}

	fmt.Println("User Deleted successfully:",rowsAffected)

	return nil
}

func (u *UserRepositoryImpl) GetAll() ([]*models.User, error) {
	//step 1 prepare the query
	query := "SELECT id,username,email,created_at,updated_at FROM users"

	//step 2 execute the query

	rows,err := u.db.Query(query)

	if err != nil {
        return nil, err
    }
    defer rows.Close()

	//step 3 process the result

	var users [] *models.User

	for rows.Next() {
		user := &models.User{}
		if err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.Created_at, &user.Updated_at); err != nil{
			return nil,err
		}
		users = append(users,user)

	}

	if err = rows.Err(); err != nil {
        return users, err
    }

    return users, nil
}
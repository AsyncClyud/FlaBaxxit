package userstorage

import (
	"context"
	"encoding/json"
	"flabaxxit/internal/models"
	"log"
	"time"
)

/* Check if user exist.
If user doens't exist - return 0, "", false.
Else returns User Id, User Password and true.
*/
func (ur *UserRepository) CheckIfUserExist(ctx context.Context, User models.User) (user_id int, hashed_password string, success bool) {
	var user models.User
	var users []models.User
	rows, err := ur.db.Query(ctx, "SELECT Id, Password FROM users WHERE Username = $1", User.Username)
	if err != nil {
		rows.Err()
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Password)
		if err != nil {
			log.Fatalln(err)
		}
		users = append(users, user)
	}
	if len(users) == 0 {
		return 0, "", false
	} else {
		return user.Id, user.Password, true
	}

}

// Get user info by his Id from database.
func (ur *UserRepository) GetUserInfo(ctx context.Context, user_id int) (user_info string, error error) {
	var user models.User
	err := ur.db.QueryRow(ctx, "SELECT Username, Bio, Created_at::text, Profile_pic FROM Users WHERE Id = $1", user_id).Scan(&user.Username, &user.Bio, &user.Created_at, &user.Profile_pic)
	if err != nil {
		log.Printf("Rows error: %v", err)
		return "", err
	}
	result, err := json.MarshalIndent(user, "", " ")
	if err != nil {
		log.Fatalln(err)
		return "", err
	}

	return string(result), nil
}

// Return user password from database.
func (ur *UserRepository) GetUserPassword(ctx context.Context, user_id int) (hashed_password string, error error) {
	var user models.User
	err := ur.db.QueryRow(ctx, "SELECT Password FROM users WHERE Id = $1", user_id).Scan(&user.Password)
	if err != nil {
		return "", err
	}
	return user.Password, nil
}

/* Create new user in database.
  Its checks if user with username already exist in database,
  then if doesn't exist: create new user and returns Id for JWT token.
*/
func (ur *UserRepository) CreateUser(ctx context.Context, new_user models.User) (user_id int, success bool) {
	_, _, UserExist := ur.CheckIfUserExist(ctx, new_user)
	var user models.User
	if !UserExist {
		_, err := ur.db.Exec(ctx, "INSERT INTO Users(Username, Password, Bio, Created_at) VALUES($1, $2, $3, $4)", new_user.Username, new_user.Password, "", time.Now())
		if err != nil {
			log.Printf("User Query error: %v", err)
			return 0, false
		}
		error := ur.db.QueryRow(ctx, "SELECT Id FROM Users WHERE Username = $1", new_user.Username).Scan(&user.Id)
		if error != nil {
			log.Printf("User Query error: %v", error)
			return 0, false
		}
	} else {
		return 0, false
	}
	log.Printf("New user has been created with id: %v", user.Id)
	return user.Id, true
}

// Delete account from database.
func (ur *UserRepository) DeleteAccount(ctx context.Context, user_id int) (success bool) {
	_, err := ur.db.Exec(ctx, "DELETE FROM Users WHERE Id = $1", user_id)
	if err != nil {
		log.Printf("Delete account query error: %v", err)
		return false
	}

	return true
}

package userstorage

import (
	"blog/internal/models"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

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

func (ur *UserRepository) GetUserInfo(ctx context.Context, user_id int) (user_info string, err error) {
	rows, err := ur.db.Query(ctx, "SELECT Username, Bio, Created_at::text, Profile_pic FROM Users WHERE Id = $1", user_id)
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()

	var user models.User
	for rows.Next() {
		err := rows.Scan(&user.Username, &user.Bio, &user.Created_at, &user.Profile_pic)
		if err != nil {
			log.Fatalln(err)
			rows.Err()
		}
	}
	result, err := json.MarshalIndent(user, "", " ")
	if err != nil {
		log.Fatalln(err)
		return "", fmt.Errorf("Error: %v", err)
	}

	return string(result), nil
}

func (ur *UserRepository) GetUserPassword(ctx context.Context, user_id int) (hashed_password string, err error) {
	var user models.User
	rows, err := ur.db.Query(ctx, "SELECT Password FROM users WHERE Id = $1", user_id)
	if err != nil {
		return "", rows.Err()
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&user.Password)
		if err != nil {
			log.Fatalln(err)
		}
	}
	return user.Password, nil
}

func (ur *UserRepository) CreateUser(ctx context.Context, new_user models.User) (user_id int, success bool) {
	_, _, UserExist := ur.CheckIfUserExist(ctx, new_user)
	var user models.User
	if !UserExist {
		_, err := ur.db.Exec(ctx, "INSERT INTO Users(Username, Password, Bio, Created_at) VALUES($1, $2, $3, $4)", new_user.Username, new_user.Password, "", time.Now())
		if err != nil {
			log.Printf("User Query error: %v", err)
			return 0, false
		}
		rows, err := ur.db.Query(ctx, "SELECT Id FROM Users WHERE Username = $1", new_user.Username)
		if err != nil {
			log.Printf("User Query error: %v", rows.Err())
			return 0, false
		}
		defer rows.Close()
		for rows.Next() {
			err := rows.Scan(&user.Id)
			if err != nil {
				log.Printf("Rows scan error: %v", err)
				return 0, false
			}
		}
	} else {
		return 0, false
	}
	log.Printf("New user has been created with id: %v", user.Id)
	return user.Id, true
}

func (ur *UserRepository) DeleteAccount(ctx context.Context, user_id int) (success bool) {
	_, err := ur.db.Exec(ctx, "DELETE FROM Users WHERE Id = $1", user_id)
	if err != nil {
		log.Printf("Delete account query error: %v", err)
		return false
	}

	return true
}

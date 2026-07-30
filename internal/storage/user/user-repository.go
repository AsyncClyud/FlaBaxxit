package userstorage

import (
	"blog/internal/models"
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) UpdateUsername(ctx context.Context, user models.User, user_id int) (success bool) {
	_, _, UsernameExist := ur.CheckIfUserExist(ctx, user)
	if !UsernameExist {
		_, err := ur.db.Exec(ctx, "UPDATE Users SET Username = $1 WHERE Id = $2", user.Username, user_id)
		if err != nil {
			log.Printf("User Query error: %v", err)
			return false
		}
		return true
	}
	return false
}

func (ur *UserRepository) UpdateBio(ctx context.Context, user models.User, user_id int) (success bool) {
	_, err := ur.db.Exec(ctx, "UPDATE Users SET Bio = $1 WHERE Id = $2", user.Bio, user_id)
	if err != nil {
		log.Printf("User Query error: %v", err)
		return false
	}

	return true
}

func (ur *UserRepository) UpdatePassword(ctx context.Context, new_password string, user_id int) (success bool) {
	_, err := ur.db.Exec(ctx, "UPDATE Users SET Password = $1 WHERE Id = $2", new_password, user_id)
	if err != nil {
		log.Printf("User Query error: %v", err)
		return false
	}

	return true
}

func (ur *UserRepository) ChangeAvatar(ctx context.Context, profile_pic models.User, user_id int) (success bool) {
	_, err := ur.db.Exec(ctx, "UPDATE Users SET profile_pic = $1 WHERE Id = $2", profile_pic.Profile_pic, user_id)
	if err != nil {
		log.Printf("Update avatar query error: %v", err)
		return false
	}

	return true
}

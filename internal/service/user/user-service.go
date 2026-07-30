package userservice

import (
	_ "blog/internal/config"
	"blog/internal/models"
	userstorage "blog/internal/storage/user"
	"context"
	"net/http"
)

type AuthService struct {
	repo       userstorage.UserRepository
	secret_key []byte
}

func NewAuthService(userDB userstorage.UserRepository, secret_key []byte) *AuthService {
	return &AuthService{repo: userDB, secret_key: secret_key}
}

func (ur *AuthService) FetchUser(ctx context.Context, user_id int) (user string, status_code int) {
	user, err := ur.repo.GetUserInfo(ctx, user_id)
	if err != nil {
		return "", http.StatusBadRequest
	}
	return user, http.StatusOK
}

func (ur *AuthService) ChangeUsername(ctx context.Context, user models.User, user_id int) (status_code int) {
	if len(user.Username) <= 2 {
		return http.StatusBadRequest
	}

	message := ur.repo.UpdateUsername(ctx, user, user_id)
	if !message {
		return http.StatusConflict
	}
	return http.StatusOK

}

func (ur *AuthService) ChangeBio(ctx context.Context, user models.User, user_id int) (status_code int) {
	message := ur.repo.UpdateBio(ctx, user, user_id)
	if !message {
		return http.StatusBadRequest
	}

	return http.StatusOK

}

func (ur *AuthService) ChangePassword(ctx context.Context, password models.NewPassword, user_id int) (status_code int) {
	hashed_password, _ := ur.repo.GetUserPassword(ctx, user_id)
	if ur.CheckPaswordHash(password.Old_Password, hashed_password) != nil {
		return http.StatusBadRequest
	}
	hashed_password, err := ur.HashPassword(password.New_Password)
	if err != nil {
		return http.StatusBadGateway
	}
	message := ur.repo.UpdatePassword(ctx, hashed_password, user_id)
	if !message {
		return http.StatusBadRequest
	}
	return http.StatusOK
}

func (ur *AuthService) ChangeAvatar(ctx context.Context, avatar_id models.User, user_id int) (status_code int) {
	if ok := ur.repo.ChangeAvatar(ctx, avatar_id, user_id); !ok {
		return http.StatusInternalServerError
	}
	return http.StatusOK

}

func (ur *AuthService) DeleteAccount(ctx context.Context, user_id int) (status_code int) {
	if success := ur.repo.DeleteAccount(ctx, user_id); !success {
		return http.StatusInternalServerError
	}
	return http.StatusOK
}

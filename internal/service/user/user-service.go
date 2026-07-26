package userservice

import (
	_ "blog/internal/config"
	"blog/internal/models"
	userstorage "blog/internal/storage/user"
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo       userstorage.UserRepository
	secret_key []byte
}

func NewAuthService(userDB userstorage.UserRepository, secret_key []byte) *AuthService {
	return &AuthService{repo: userDB, secret_key: secret_key}
}

func (ur *AuthService) HashPassword(password string) (hashed_password string, err error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (ur *AuthService) CheckPaswordHash(password string, hashedPassword string) (err error) {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (ur *AuthService) Generate_Token(id int) (token string, err error) {
	claims := jwt.MapClaims{
		"UserID": id,
		"exp":    time.Now().Add(24 * time.Hour).Unix(),
	}
	new_token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	return new_token.SignedString(ur.secret_key)

}

func (ur *AuthService) Validate_Token(tokenString string) (Id int, err error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
		return ur.secret_key, nil
	})
	if err != nil {
		return 0, fmt.Errorf("Parsing error: %v", err)
	}
	if !token.Valid {
		return 0, fmt.Errorf("Invalid token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, fmt.Errorf("Invalid claims")
	}
	userIDValue, exists := claims["UserID"]
	if !exists {
		return 0, fmt.Errorf("Invalid UserID")
	}
	userID := userIDValue.(float64)

	return int(userID), nil
}

func (ur *AuthService) SetTokenInCookie(c *gin.Context, id int) {
	jwt_token, err := ur.Generate_Token(id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.SetCookie("jwt-token", jwt_token, 3200*20, "/", "", true, true)
}

func (ur *AuthService) ValidateUserData(user models.User) (status_code int) {
	if strings.ContainsAny(user.Username, ` !@#$%^&*();:{}[]"'.,?/-_+=`) == true {
		return http.StatusNotAcceptable
	}
	if len(user.Username) <= 2 {
		return http.StatusBadRequest
	}
	if len(user.Password) <= 5 {
		return http.StatusUnprocessableEntity
	}
	return http.StatusOK
}

func (ur *AuthService) FetchUser(ctx context.Context, user_id int) (user string, status_code int) {
	user, err := ur.repo.GetUserInfo(ctx, user_id)
	if err != nil {
		return "", http.StatusBadRequest
	}
	return user, http.StatusOK
}

func (ur *AuthService) Register(ctx context.Context, user models.User) (status_code, id int) {
	if code := ur.ValidateUserData(user); code != http.StatusOK {
		return code, 0
	}
	hashed_password, err := ur.HashPassword(user.Password)
	if err != nil {
		return http.StatusBadGateway, 0
	}
	new_user := models.User{Username: user.Username, Password: hashed_password}

	id, message := ur.repo.CreateUser(ctx, new_user)
	if !message {
		return http.StatusConflict, 0
	}
	return http.StatusOK, id
}

func (ur *AuthService) Login(ctx context.Context, user models.User) (status_code, id int) {
	if code := ur.ValidateUserData(user); code != http.StatusOK {
		return code, 0
	}
	id, hashed_password, ok := ur.repo.CheckIfUserExist(ctx, user)
	if !ok {
		return http.StatusNotFound, 0
	}
	if ur.CheckPaswordHash(user.Password, hashed_password) != nil {
		return http.StatusNotFound, 0
	}
	return http.StatusOK, id
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

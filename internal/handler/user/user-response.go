package userhandler

import (
	"blog/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FormatIntoJson(c *gin.Context, status int, payload string) {
	c.Header("Content-Type", "application/json")

	var message models.Message = models.Message{Message: payload}

	c.JSON(status, message)

}

func ResponseRegistration(status_code int, c *gin.Context) {
	switch status_code {
	case http.StatusOK:
		FormatIntoJson(c, http.StatusOK, "Account has been created!")
	case http.StatusBadRequest:
		FormatIntoJson(c, http.StatusBadRequest, "Username must be at least 4 characters long!")
	case http.StatusForbidden:
		FormatIntoJson(c, http.StatusForbidden, "Captcha verification failed!")
	case http.StatusNotAcceptable:
		FormatIntoJson(c, http.StatusNotAcceptable, "Username can only contain letters, numbers!")
	case http.StatusConflict:
		FormatIntoJson(c, http.StatusConflict, "Account with this username already exist!")
	case http.StatusUnprocessableEntity:
		FormatIntoJson(c, http.StatusUnprocessableEntity, "Password must be at least 6 characters long!")
	case http.StatusBadGateway:
		FormatIntoJson(c, http.StatusInternalServerError, "Internal error!")
	}
}

func ResponseLogin(status_code int, c *gin.Context) {
	switch status_code {
	case http.StatusOK:
		FormatIntoJson(c, http.StatusOK, "You has been successfully logined!")
	case http.StatusBadRequest:
		FormatIntoJson(c, http.StatusBadRequest, "User with this username doesn't exist!")
	case http.StatusForbidden:
		FormatIntoJson(c, http.StatusForbidden, "Captcha verification failed!")
	case http.StatusNotFound:
		FormatIntoJson(c, http.StatusNotFound, "Invalid password!")
	case http.StatusNotAcceptable:
		FormatIntoJson(c, http.StatusNotAcceptable, "Username can only contain letters, numbers!")
	case http.StatusConflict:
		FormatIntoJson(c, http.StatusConflict, "Account with this username already exist!")
	case http.StatusUnprocessableEntity:
		FormatIntoJson(c, http.StatusUnprocessableEntity, "Password must be at least 6 characters long!")
	case http.StatusBadGateway:
		FormatIntoJson(c, http.StatusInternalServerError, "Internal error!")
	}
}

func ResponseUsernameChange(status_code int, c *gin.Context) {
	switch status_code {
	case http.StatusOK:
		FormatIntoJson(c, http.StatusOK, "Your username has been updated!")
	case http.StatusBadRequest:
		FormatIntoJson(c, http.StatusBadRequest, "Username is too short!")
	case http.StatusConflict:
		FormatIntoJson(c, http.StatusConflict, "Username already in use!")
	}
}

func ResponseBioChange(status_code int, c *gin.Context) {
	switch status_code {
	case http.StatusOK:
		FormatIntoJson(c, http.StatusOK, "Your bio has been updated!")
	case http.StatusBadRequest:
		FormatIntoJson(c, http.StatusBadRequest, "Bio is too long! 2000 chars max!")
	}
}

func ResponsePasswordChange(status_code int, c *gin.Context) {
	switch status_code {
	case http.StatusOK:
		FormatIntoJson(c, http.StatusOK, "Password has been updated!")
	case http.StatusBadRequest:
		FormatIntoJson(c, http.StatusBadRequest, "Incorrect password!")
	}
}

func ResponseAvatarChange(status_code int, c *gin.Context) {
	switch status_code {
	case http.StatusOK:
		FormatIntoJson(c, http.StatusOK, "Avatar has been changed!")
	case http.StatusInternalServerError:
		FormatIntoJson(c, http.StatusInternalServerError, "Internal error!")
	}
}

func ResponseAccountDelete(status_code int, c *gin.Context) {
	switch status_code {
	case http.StatusOK:
		FormatIntoJson(c, http.StatusOK, "Account has been successfuly deleted!")
	case http.StatusInternalServerError:
		FormatIntoJson(c, http.StatusInternalServerError, "Internal error!")
	}
}

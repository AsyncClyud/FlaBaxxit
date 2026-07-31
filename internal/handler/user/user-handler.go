package userhandler

import (
	"flabaxxit/internal/config"
	"flabaxxit/internal/models"
	userservice "flabaxxit/internal/service/user"
	captcha "flabaxxit/internal/turnstile"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	authService userservice.AuthService
	Turnslite   captcha.Verifier
	Config      config.Config
}

func NewUserHandler(service userservice.AuthService, config config.Config) *UserHandler {
	return &UserHandler{authService: service, Turnslite: *captcha.NewVerifier(config), Config: config}
}

func (ush *UserHandler) ProfileHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	id, err := strconv.Atoi(c.Param("Id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Message": "Invalid body"})
	}

	ctx := c.Request.Context()
	user, status_code := ush.authService.FetchUser(ctx, id)
	if status_code != http.StatusOK {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Message": "Cannot fetch user data"})
		return
	}

	c.JSON(http.StatusOK, user)

}

func (ush *UserHandler) ChangeUsernameHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	cookie, exist := c.Cookie("jwt-token")
	if exist != nil {
		c.AbortWithError(http.StatusUnauthorized, exist)
		return
	}
	claims, ok := ush.authService.Validate_Token(cookie)
	if ok != nil {
		c.AbortWithError(http.StatusUnauthorized, ok)
		return
	}
	ctx := c.Request.Context()
	status_code := ush.authService.ChangeUsername(ctx, user, claims)
	ResponseUsernameChange(status_code, c)
}

func (ush *UserHandler) ChangeBioHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	cookie, exist := c.Cookie("jwt-token")
	if exist != nil {
		c.AbortWithError(http.StatusUnauthorized, exist)
		return
	}
	claims, ok := ush.authService.Validate_Token(cookie)
	if ok != nil {
		c.AbortWithError(http.StatusUnauthorized, ok)
		return
	}
	ctx := c.Request.Context()
	status_code := ush.authService.ChangeBio(ctx, user, claims)
	ResponseBioChange(status_code, c)
}

func (ush *UserHandler) ChangePasswordHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	var passwords models.NewPassword
	err := c.ShouldBindJSON(&passwords)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	cookie, exist := c.Cookie("jwt-token")
	if exist != nil {
		c.AbortWithError(http.StatusUnauthorized, exist)
		return
	}
	claims, ok := ush.authService.Validate_Token(cookie)
	if ok != nil {
		c.AbortWithError(http.StatusUnauthorized, ok)
		return
	}
	ctx := c.Request.Context()
	status_code := ush.authService.ChangePassword(ctx, passwords, claims)
	ResponsePasswordChange(status_code, c)
}

func (ush *UserHandler) ChangeAvatarHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	cookie, exist := c.Cookie("jwt-token")
	if exist != nil {
		c.AbortWithError(http.StatusUnauthorized, exist)
		return
	}
	claims, ok := ush.authService.Validate_Token(cookie)
	if ok != nil {
		c.AbortWithError(http.StatusUnauthorized, ok)
		return
	}

	var avatar_id models.User
	err := c.ShouldBind(&avatar_id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	status_code := ush.authService.ChangeAvatar(c, avatar_id, claims)
	ResponseAvatarChange(status_code, c)
}

func (ush *UserHandler) GetArticleAuthorHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	cookie, exist := c.Cookie("jwt-token")
	if exist != nil {
		c.AbortWithError(http.StatusUnauthorized, exist)
		return
	}
	claims, ok := ush.authService.Validate_Token(cookie)
	if ok != nil {
		c.AbortWithError(http.StatusUnauthorized, ok)
		return
	}

	var author models.Article

	err := c.ShouldBindJSON(&author)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}
	if claims != author.Author {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}

func (ush *UserHandler) FetchUserProfile(c *gin.Context) {
	c.Header("Content-Type", "application/json")


}

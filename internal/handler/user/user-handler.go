package userhandler

import (
	"blog/internal/config"
	"blog/internal/contextutil"
	"blog/internal/models"
	userservice "blog/internal/service/user"
	captcha "blog/internal/turnstile"
	"net/http"

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

func (ush *UserHandler) IsAuth(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	cookie, err := c.Cookie("jwt-token")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"authorized": false})
		return
	}
	token, err := ush.authService.Validate_Token(cookie)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"authorized": false})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"authorized": true,
		"userID":     token,
	})
}

func (ush *UserHandler) RegisterHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		status_code := http.StatusBadRequest
		ResponseRegistration(status_code, c)
		return
	}
	cfToken := user.Turnstile_token
	remoteAddr := c.RemoteIP()

	ok, err := ush.Turnslite.Verify(c.Request.Context(), cfToken, remoteAddr)
	if err != nil || !ok {
		status_code := http.StatusForbidden
		ResponseRegistration(status_code, c)
		return
	}

	status_code, id := ush.authService.Register(user)
	if status_code == 200 {
		ush.authService.SetTokenInCookie(c, id)
		ResponseRegistration(status_code, c)
		return
	} else {
		ResponseRegistration(status_code, c)
	}
}

func (ush *UserHandler) LoginHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		status_code := http.StatusBadRequest
		ResponseRegistration(status_code, c)
		return
	}
	cfToken := user.Turnstile_token
	remoteAddr := c.RemoteIP()

	ok, err := ush.Turnslite.Verify(c.Request.Context(), cfToken, remoteAddr)
	if err != nil || !ok {
		status_code := http.StatusForbidden
		ResponseLogin(status_code, c)
		return
	}

	status_code, id := ush.authService.Login(user)
	if status_code == 200 {
		ush.authService.SetTokenInCookie(c, id)
		ResponseLogin(status_code, c)
		return
	} else {
		ResponseLogin(status_code, c)
		return
	}

}

func (ush *UserHandler) ProfileHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	userID, ok := c.Request.Context().Value(contextutil.UserIDKey).(int)
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	user, status_code := ush.authService.FetchUser(userID)
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

	status_code := ush.authService.ChangeUsername(user, claims)
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

	status_code := ush.authService.ChangeBio(user, claims)
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

	status_code := ush.authService.ChangePassword(passwords, claims)
	ResponsePasswordChange(status_code, c)
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

func (ush *UserHandler) LogoutHandler(c *gin.Context) {
	c.SetCookie("jwt-token", "", 0, "/", "", true, true)
}

package userhandler

import (
	"flabaxxit/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ush *UserHandler) IsAuth(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	cookie, err := c.Cookie("jwt-token")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"authorized": false,
			"userID": 0,
		})
		return
	}
	token, err := ush.authService.Validate_Token(cookie)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"authorized": false,
			"userID": 0,
		})
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
	ctx := c.Request.Context()
	status_code, id := ush.authService.Register(ctx, user)
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
	ctx := c.Request.Context()
	status_code, id := ush.authService.Login(ctx, user)
	if status_code == 200 {
		ush.authService.SetTokenInCookie(c, id)
		ResponseLogin(status_code, c)
		return
	} else {
		ResponseLogin(status_code, c)
		return
	}

}

func (ush *UserHandler) LogoutHandler(c *gin.Context) {
	c.SetCookie("jwt-token", "", 0, "/", "", true, true)
}

func (ush *UserHandler) DeleteAccountHandler(c *gin.Context) {
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
	status_code := ush.authService.DeleteAccount(c.Request.Context(), claims)
	ResponseAccountDelete(status_code, c)
}

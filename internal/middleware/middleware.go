package middleware

import (
	"context"
	"flabaxxit/internal/contextutil"
	userservice "flabaxxit/internal/service/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Middleware struct {
	Middleware userservice.AuthService
}

func NewAuthMiddleware(service userservice.AuthService) *Middleware {
	return &Middleware{Middleware: service}
}

func (md Middleware) SecureHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("X-Frame-Options", "DENY")
		c.Header("Referrer-Policy", "strict-origin-when-cross-origin")
		c.Header("Permissions-Policy", "geolocation=(), microphone=(), camera=()")
		c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
		c.Header("Cross-Origin-Opener-Policy", "same-origin")
		c.Header("Cross-Origin-Resource-Policy", "same-origin")

		c.Next()
	}
}

func (md Middleware) RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("X-Frame-Options", "DENY")
		c.Header("Referrer-Policy", "strict-origin-when-cross-origin")
		c.Header("Permissions-Policy", "geolocation=(), microphone=(), camera=()")
		c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
		c.Header("Cross-Origin-Opener-Policy", "same-origin")
		c.Header("Cross-Origin-Resource-Policy", "same-origin")

		cookie, err := c.Cookie("jwt-token")
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}
		claims, err := md.Middleware.Validate_Token(cookie)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		ctx := context.WithValue(c.Request.Context(), contextutil.UserIDKey, claims)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}

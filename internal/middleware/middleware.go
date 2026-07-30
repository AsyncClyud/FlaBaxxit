package middleware

import (
	"blog/internal/contextutil"
	userservice "blog/internal/service/user"
	"context"
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
	return func(ctx *gin.Context) {
		ctx.Header("X-Content-Type-Options", "nosniff")
		ctx.Header("Referrer-Policy", "strict-origin-when-cross-origin")
		ctx.Header("Permissions-Policy", "geolocation=(), microphone=(), camera=()")
		ctx.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains")

		ctx.Next()
	}
}

func (md Middleware) RequireAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("X-Content-Type-Options", "nosniff")
		ctx.Header("Referrer-Policy", "strict-origin-when-cross-origin")
		ctx.Header("Permissions-Policy", "geolocation=(), microphone=(), camera=()")
		ctx.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains")

		cookie, err := ctx.Cookie("jwt-token")
		if err != nil {
			ctx.AbortWithError(http.StatusUnauthorized, err)
			return
		}
		claims, err := md.Middleware.Validate_Token(cookie)
		if err != nil {
			ctx.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		c := context.WithValue(ctx.Request.Context(), contextutil.UserIDKey, claims)
		ctx.Request = ctx.Request.WithContext(c)

		ctx.Next()
	}
}

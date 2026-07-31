package router

import (
	posthandler "flabaxxit/internal/handler/post"
	userhandler "flabaxxit/internal/handler/user"
	"flabaxxit/internal/middleware"
	userservice "flabaxxit/internal/service/user"
	poststorage "flabaxxit/internal/storage/post"
	userstorage "flabaxxit/internal/storage/user"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func Router(postDB poststorage.PostRepository, userDB userstorage.UserRepository, postHandler posthandler.PostHandler, userHandler userhandler.UserHandler, authUser userservice.AuthService, middleware middleware.Middleware) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.SetTrustedProxies(nil)
	r.Use(gzip.Gzip(gzip.BestCompression))
	r.Use(gin.Logger(), gin.ErrorLogger())
	r.NoRoute(func(c *gin.Context) {
		c.File("./web/not_found.html")
	})

	r.GET("/", middleware.SecureHeaders(), postHandler.ServePage("./web/index.html"))
	r.GET("/terms", middleware.SecureHeaders(), postHandler.ServePage("./web/terms.html"))
	r.GET("/privacy", middleware.SecureHeaders(), postHandler.ServePage("./web/privacy.html"))
	r.GET("/not_found", middleware.SecureHeaders(), postHandler.ServePage("./web/not_found.html"))

	r.GET("/api/auth", middleware.SecureHeaders(), userHandler.IsAuth)
	r.GET("/auth/register", middleware.SecureHeaders(), postHandler.ServePage("./web/auth/register.html"))
	r.POST("/auth/register", middleware.SecureHeaders(), userHandler.RegisterHandler)
	r.GET("/auth/login", middleware.SecureHeaders(), postHandler.ServePage("./web/auth/login.html"))
	r.POST("/auth/login", middleware.SecureHeaders(), userHandler.LoginHandler)

	r.GET("/api/profile/:Id", middleware.SecureHeaders(), userHandler.ProfileHandler)
	r.PUT("/api/profile/username", middleware.RequireAuth(), userHandler.ChangeUsernameHandler)
	r.PUT("/api/profile/password", middleware.RequireAuth(), userHandler.ChangePasswordHandler)
	r.PUT("/api/profile/bio", middleware.RequireAuth(), userHandler.ChangeBioHandler)
	r.PUT("/api/profile/avatar", middleware.RequireAuth(), userHandler.ChangeAvatarHandler)
	r.GET("/profile/:Id", middleware.SecureHeaders(), postHandler.ServePage("./web/profile/main_profile.html"))
	r.GET("/profile/settings", middleware.RequireAuth(), postHandler.ServePage("./web/profile/settings.html"))
	r.POST("/api/logout", middleware.RequireAuth(), userHandler.LogoutHandler)
	r.DELETE("/api/users", middleware.RequireAuth(), userHandler.DeleteAccountHandler)

	r.GET("/api/articles", middleware.SecureHeaders(), postHandler.GetArticlesHandler)
	r.GET("/api/articles/:Id", middleware.SecureHeaders(), postHandler.GetArticleByIdHandler)
	r.POST("/api/articles", middleware.RequireAuth(), postHandler.InsertArticleHandler)
	r.PUT("/api/articles", middleware.RequireAuth(), postHandler.UpdateArticleHandler)
	r.DELETE("/api/articles", middleware.RequireAuth(), postHandler.DeleteArticleHandler)
	r.GET("/article/:Id", middleware.SecureHeaders(), postHandler.ServePage("./web/article/article.html"))
	r.GET("/article/create", middleware.RequireAuth(), postHandler.ServePage("./web/article/create_article.html"))
	r.GET("/article/update/:Id", middleware.RequireAuth(), postHandler.ServePage("./web/article/update_article.html"))

	r.GET("/api/comments/:Id", middleware.SecureHeaders(), postHandler.GetArticleComments)
	r.POST("/api/comments", middleware.RequireAuth(), postHandler.InsertCommentHandler)

	r.POST("/api/users", middleware.RequireAuth(), userHandler.GetArticleAuthorHandler)

	r.Static("/static", "./web/static")

	return r
}

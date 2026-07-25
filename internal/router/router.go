package router

import (
	posthandler "blog/internal/handler/post"
	userhandler "blog/internal/handler/user"
	"blog/internal/middleware"
	userservice "blog/internal/service/user"
	poststorage "blog/internal/storage/post"
	userstorage "blog/internal/storage/user"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func Router(postDB poststorage.PostRepository, userDB userstorage.UserRepository, postHandler posthandler.PostHandler, userHandler userhandler.UserHandler, authUser userservice.AuthService, middleware middleware.Middleware) *gin.Engine {
	r := gin.Default()
	r.SetTrustedProxies(nil)
	r.Use(gzip.Gzip(gzip.DefaultCompression))
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

	r.GET("/api/profile", middleware.RequireAuth(), userHandler.ProfileHandler)
	r.PUT("/api/profile/username", middleware.RequireAuth(), userHandler.ChangeUsernameHandler)
	r.PUT("/api/profile/password", middleware.RequireAuth(), userHandler.ChangePasswordHandler)
	r.PUT("/api/profile/bio", middleware.RequireAuth(), userHandler.ChangeBioHandler)
	r.GET("/profile", middleware.RequireAuth(), postHandler.ServePage("./web/profile/main_profile.html"))
	r.GET("/profile/settings", middleware.RequireAuth(), postHandler.ServePage("./web/profile/settings.html"))

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

	r.POST("/api/logout", middleware.RequireAuth(), userHandler.LogoutHandler)
	r.POST("/api/users", middleware.RequireAuth(), userHandler.GetArticleAuthorHandler)

	r.Static("/static", "./web/static")

	return r
}

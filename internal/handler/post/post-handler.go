package posthandler

import (
	"blog/internal/config"
	"blog/internal/models"
	postservice "blog/internal/service/post"
	userservice "blog/internal/service/user"
	captcha "blog/internal/turnstile"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	postService postservice.PostService
	authService userservice.AuthService
	Turnslite   captcha.Verifier
	Config      config.Config
}

func NewPostHandler(postservice postservice.PostService, auth userservice.AuthService, config config.Config) *PostHandler {
	return &PostHandler{postService: postservice, authService: auth, Turnslite: *captcha.NewVerifier(config), Config: config}
}

func (psh *PostHandler) ServePage(html_file string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.File(html_file)
	}
}

func (psh *PostHandler) GetArticlesHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	ctx := c.Request.Context()
	articles := psh.postService.GetArticles(ctx)

	c.JSON(http.StatusOK, articles)
}

func (psh *PostHandler) GetArticleByIdHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	ctx := c.Request.Context()
	Id, err := strconv.Atoi(c.Param("Id"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	article := psh.postService.GetArticleById(ctx, Id)

	c.JSON(http.StatusOK, article)
}

func (psh *PostHandler) InsertArticleHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	var article models.Article
	err := c.ShouldBindJSON(&article)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	cfToken := article.Turnstile_token
	remoteAddr := c.RemoteIP()

	ok, err := psh.Turnslite.Verify(c.Request.Context(), cfToken, remoteAddr)
	if err != nil || !ok {
		status_code := http.StatusForbidden
		ResponseArticle(status_code, c)
		return
	}

	cookie, err := c.Cookie("jwt-token")
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}
	userID, err := psh.authService.Validate_Token(cookie)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}
	ctx := c.Request.Context()
	status_code := psh.postService.InsertArticle(ctx, article, userID)
	ResponseArticle(status_code, c)
}

func (psh *PostHandler) UpdateArticleHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	var article models.Article
	err := c.ShouldBindJSON(&article)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	cfToken := article.Turnstile_token
	remoteAddr := c.RemoteIP()

	ok, err := psh.Turnslite.Verify(c.Request.Context(), cfToken, remoteAddr)
	if err != nil || !ok {
		status_code := http.StatusForbidden
		ResponseArticle(status_code, c)
		return
	}

	cookie, err := c.Cookie("jwt-token")
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}
	userID, err := psh.authService.Validate_Token(cookie)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}
	if article.Author != userID {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}
	ctx := c.Request.Context()
	status_code := psh.postService.UpdateArticle(ctx, article)
	ResponseArticle(status_code, c)
}

func (psh *PostHandler) DeleteArticleHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	var article models.Article
	err := c.ShouldBindJSON(&article)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	cookie, err := c.Cookie("jwt-token")
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}
	userID, err := psh.authService.Validate_Token(cookie)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}
	if article.Author != userID {
		ResponseArticle(http.StatusForbidden, c)
		return
	}
	ctx := c.Request.Context()
	psh.postService.DeleteArticle(ctx, article)
}

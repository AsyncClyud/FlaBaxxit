package posthandler

import (
	"flabaxxit/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (psh *PostHandler) GetArticleComments(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	ctx := c.Request.Context()
	id, ok := strconv.Atoi(c.Param("Id"))
	if ok != nil {
		c.AbortWithError(http.StatusBadRequest, ok)
	}

	comments := psh.postService.GetArticleCommentsById(ctx, id)

	c.JSON(http.StatusOK, comments)
}

func (psh *PostHandler) InsertCommentHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")

	var comment models.Comment
	err := c.ShouldBindJSON(&comment)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	cfToken := comment.Turnstile_token
	remoteAddr := c.RemoteIP()

	ok, err := psh.Turnslite.Verify(c.Request.Context(), cfToken, remoteAddr)
	if err != nil || !ok {
		status_code := http.StatusForbidden
		ResponseComment(status_code, c)
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
	status_code := psh.postService.InsertComment(ctx, comment, userID)
	ResponseComment(status_code, c)
}

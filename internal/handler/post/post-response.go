package posthandler

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

func ResponseArticle(status_code int, c *gin.Context) {
	switch status_code {
	case http.StatusOK:
		FormatIntoJson(c, http.StatusOK, "Success!")
	case http.StatusBadRequest:
		FormatIntoJson(c, http.StatusBadRequest, "Article title is too short!")
	case http.StatusUnprocessableEntity:
		FormatIntoJson(c, http.StatusBadRequest, "Article content is too short!")
	case http.StatusUnauthorized:
		FormatIntoJson(c, http.StatusUnauthorized, "You don't have permission to delete/edit this article!")
	case http.StatusForbidden:
		FormatIntoJson(c, http.StatusForbidden, "Captcha verification required!")
	}
}

func ResponseComment(status_code int, c *gin.Context) {
	switch status_code {
	case http.StatusOK:
		FormatIntoJson(c, http.StatusOK, "Success!")
	case http.StatusBadRequest:
		FormatIntoJson(c, http.StatusBadRequest, "Comment content cannot be null!")
	case http.StatusForbidden:
		FormatIntoJson(c, http.StatusForbidden, "Captcha verification required!")
	}
}

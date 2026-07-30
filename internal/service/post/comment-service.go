package postservice

import (
	"blog/internal/models"
	"blog/internal/moderation"
	"context"
	"net/http"
)


func (pr PostService) ValidateComment(comment models.Comment) (status_code int) {
	if len(comment.Comment_content) == 0 {
		return http.StatusBadRequest
	}
	return http.StatusOK
}

func (pr PostService) GetArticleCommentsById(ctx context.Context, id int) (comments string) {
	return pr.repo.GetArticleCommentsById(ctx, id)
}

func (pr PostService) InsertComment(ctx context.Context, comment models.Comment, author_id int) (status_code int) {
	status := pr.ValidateComment(comment)
	if status != http.StatusOK {
		return status
	}
	contains_bad_content := moderation.ModerateText(comment.Comment_content)
	if contains_bad_content == true {
		return http.StatusUnprocessableEntity
	}

	pr.repo.InsertComment(ctx, comment, author_id)
	return http.StatusOK
}

package postservice

import (
	"blog/internal/models"
	"blog/internal/moderation"
	poststorage "blog/internal/storage/post"
	"context"
	"net/http"
)

type PostService struct {
	repo poststorage.PostRepository
}

func NewPostService(repo poststorage.PostRepository) *PostService {
	return &PostService{repo: repo}
}

func (pr PostService) ValidateArticle(article models.Article) (status_code int) {
	if len(article.Title) <= 3 {
		return http.StatusBadRequest
	}
	if len(article.Content) <= 3 {
		return http.StatusUnprocessableEntity
	}
	return http.StatusOK
}

func (pr PostService) GetArticles(ctx context.Context) (articles string) {
	return pr.repo.GetAllArticles(ctx)
}

func (pr PostService) GetArticleById(ctx context.Context, id int) (article string) {
	return pr.repo.GetArticleById(ctx, id)
}

func (pr PostService) InsertArticle(ctx context.Context, article models.Article, Author_Id int) (status_code int) {
	status := pr.ValidateArticle(article)
	if status == 400 {
		return http.StatusBadRequest
	}
	contains_bad_title := moderation.ModerateText(article.Title)
	contains_bad_article := moderation.ModerateText(article.Content)
	if contains_bad_title == true || contains_bad_article == true {
		return http.StatusConflict
	}

	pr.repo.InsertArticle(ctx, article, Author_Id)
	return http.StatusOK
}

func (pr PostService) UpdateArticle(ctx context.Context, article models.Article) (status_code int) {
	status := pr.ValidateArticle(article)
	if status == 400 {
		return http.StatusBadRequest
	}
	contains_bad_title := moderation.ModerateText(article.Title)
	contains_bad_article := moderation.ModerateText(article.Content)

	if contains_bad_title == true || contains_bad_article == true {
		return http.StatusConflict
	}

	pr.repo.UpdateArticle(ctx, article)
	return http.StatusOK
}

func (pr PostService) DeleteArticle(ctx context.Context, article models.Article) {
	pr.repo.DeleteArticle(ctx, article)
}

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
	if contains_bad_content == true{
		return http.StatusUnprocessableEntity
	}

	pr.repo.InsertComment(ctx, comment, author_id)
	return http.StatusOK
}

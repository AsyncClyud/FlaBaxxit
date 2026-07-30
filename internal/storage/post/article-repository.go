package poststorage

import (
	"blog/internal/models"
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostRepository struct {
	db *pgxpool.Pool
}

func NewPostRepo(db *pgxpool.Pool) *PostRepository {
	return &PostRepository{db: db}
}

func (pr *PostRepository) GetAllArticles(ctx context.Context) (all_articles string) {
	query := `SELECT
        p.id, p.title,
        u.id, u.username, u.profile_pic
    FROM posts p
    JOIN users u ON u.id = p.author`

	rows, err := pr.db.Query(ctx, query)
	if err != nil {
		log.Println("Rows error:", err)
		rows.Err()
	}
	defer rows.Close()

	articles := []models.ArticleWithAuthor{}
	for rows.Next() {
		article := models.ArticleWithAuthor{}
		err := rows.Scan(&article.Id, &article.Title, &article.Author_Id, &article.Author_Username, &article.Author_Avatar)
		if err != nil {
			log.Println(err)
		}
		articles = append(articles, article)
	}

	result, err := json.MarshalIndent(articles, "", " ")
	if err != nil {
		log.Println(err)
	}

	return string(result)
}

func (pr *PostRepository) GetArticleById(ctx context.Context, Id int) (byid_article string) {
	query := `SELECT
        p.title, p.content, p.created_at::text,
        u.id, u.username, u.profile_pic
    FROM posts p
    JOIN users u ON u.id = p.author
    WHERE p.id = $1`

	rows, err := pr.db.Query(ctx, query, Id)
	if err != nil {
		log.Println("Rows error:", err)
		rows.Err()
	}
	defer rows.Close()

	var article models.ArticleWithAuthor
	for rows.Next() {
		err := rows.Scan(&article.Title, &article.Content, &article.Created_at, &article.Author_Id, &article.Author_Username, &article.Author_Avatar)
		if err != nil {
			log.Println(err)
		}
	}
	result, err := json.MarshalIndent(article, "", " ")
	if err != nil {
		log.Println(err)
	}

	return string(result)
}

func (pr *PostRepository) InsertArticle(ctx context.Context, article models.Article, author int) {
	_, err := pr.db.Exec(
		ctx, "INSERT INTO Posts(Title, Content, Created_at, Author) VALUES ($1, $2, $3, $4)", article.Title, article.Content, time.Now(), author)
	if err != nil {
		log.Println("Insert article query error:", err)
	}
	log.Printf("Inserted new article with title %v; Article author: %v", article.Title, author)
}

func (pr *PostRepository) UpdateArticle(ctx context.Context, article models.Article) {
	_, err := pr.db.Exec(
		ctx, "UPDATE Posts SET Title = $1, Content = $2 WHERE Id = $3", article.Title, article.Content, article.Id)
	if err != nil {
		log.Println("Update article query error:", err)
	}
	log.Printf("Updated article with title: %v", article.Title)
}

func (pr *PostRepository) DeleteArticle(ctx context.Context, article models.Article) {
	_, err := pr.db.Exec(ctx, "DELETE FROM Posts WHERE Id = $1", article.Id)
	if err != nil {
		log.Println("Delete article query error:", err)
	}
	log.Printf("Deleted article with id: %v", article.Id)
}

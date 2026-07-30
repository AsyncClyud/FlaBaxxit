package poststorage

import (
	"blog/internal/models"
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type PostRepository struct {
	db *pgxpool.Pool
	rdb *redis.Client
}

func NewPostRepo(db *pgxpool.Pool, rdb *redis.Client) *PostRepository {
	return &PostRepository{db: db, rdb: rdb}
}

func (pr *PostRepository) GetAllArticles(ctx context.Context) (all_articles []models.ArticleWithAuthor) {
	articlesCacheKey := "articles:list"

	val, err := pr.rdb.Get(ctx, articlesCacheKey).Result()
	if err == nil {
		var articles []models.ArticleWithAuthor
		err := json.Unmarshal([]byte(val), &articles)
		if err != nil {
			log.Println(err)
		}

		return articles

	}

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
	data, err := json.Marshal(articles)
	if err != nil {
		log.Printf("Json marshal error: %v", err)
	}
	pr.rdb.Set(ctx, articlesCacheKey, data, 5*time.Minute)

	return articles
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

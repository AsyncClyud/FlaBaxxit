package poststorage

import (
	"context"
	"encoding/json"
	"flabaxxit/internal/models"
	"log"
	"strconv"
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
		log.Printf("Json marshal error: %v\n", err)
	}
	pr.rdb.Set(ctx, articlesCacheKey, data, 15*time.Minute)

	return articles
}

func (pr *PostRepository) GetArticleById(ctx context.Context, Id int) (articles models.ArticleWithAuthor) {
	articleCacheKey := strconv.Itoa(Id)

	val, err := pr.rdb.Get(ctx, articleCacheKey).Result()
	if err == nil {
		var article models.ArticleWithAuthor
		err := json.Unmarshal([]byte(val), &article)
		if err != nil {
			log.Println(err)
		}
		return article
	}
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

	data, err := json.Marshal(article)
	if err != nil {
		log.Printf("Json Marshal error: %v\n", err)
	}
	pr.rdb.Set(ctx, articleCacheKey, data, 15*time.Minute)

	return article
}

func (pr *PostRepository) InsertArticle(ctx context.Context, article models.Article, author int) {
	articlesCacheKey := "articles:list"
	_, err := pr.db.Exec(
		ctx, "INSERT INTO Posts(Title, Content, Created_at, Author) VALUES ($1, $2, $3, $4)", article.Title, article.Content, time.Now(), author)
	if err != nil {
		log.Println("Insert article query error:", err)
	}
	log.Printf("Inserted new article with title %v; Article author: %v", article.Title, author)
	pr.rdb.Del(ctx, articlesCacheKey)
}

func (pr *PostRepository) UpdateArticle(ctx context.Context, article models.Article) {
	articleCacheKey := strconv.Itoa(article.Id)
	_, err := pr.db.Exec(
		ctx, "UPDATE Posts SET Title = $1, Content = $2 WHERE Id = $3", article.Title, article.Content, article.Id)
	if err != nil {
		log.Println("Update article query error:", err)
	}
	pr.rdb.Del(ctx, articleCacheKey)
	log.Printf("Updated article with title: %v", article.Title)
}

func (pr *PostRepository) DeleteArticle(ctx context.Context, article models.Article) {
	articlesCacheKey := "articles:list"
	_, err := pr.db.Exec(ctx, "DELETE FROM Posts WHERE Id = $1", article.Id)
	if err != nil {
		log.Println("Delete article query error:", err)
	}
	log.Printf("Deleted article with id: %v", article.Id)
	pr.rdb.Del(ctx, articlesCacheKey)
}

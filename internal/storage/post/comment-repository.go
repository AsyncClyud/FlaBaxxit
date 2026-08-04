package poststorage

import (
	"context"
	"encoding/json"
	"flabaxxit/internal/models"
	"log"
	"strconv"
	"time"
)

// Returns slice with all article comments from redis cache or from database (if cache is empty).
func (pr *PostRepository) GetArticleCommentsById(ctx context.Context, id int) (article_comments []models.Comment) {
	comments_id := strconv.Itoa(id) + ":list"

	val, err := pr.rdb.Get(ctx, comments_id).Result()
	if err == nil {
		var comments []models.Comment
		err := json.Unmarshal([]byte(val), &comments)
		if err != nil {
			log.Println(err)
		}
		return comments
	}

	query := `SELECT
        c.comment_content, c.created_at::text, c.author,
        u.username, u.profile_pic
    FROM comments c
    JOIN users u ON u.id = c.author
    WHERE c.post_id = $1`
	var comments []models.Comment

	rows, err := pr.db.Query(ctx, query, id)
	if err != nil {
		log.Println("Rows error:", err)
		rows.Err()
	}
	defer rows.Close()

	for rows.Next() {
		var comment models.Comment
		err := rows.Scan(&comment.Comment_content, &comment.Created_at, &comment.Author_Id, &comment.Author_Username, &comment.Author_Avatar)
		if err != nil {
			log.Println(err)
		}
		comments = append(comments, comment)
	}
	data, err := json.Marshal(comments)
	if err != nil {
		log.Printf("Json marshal error: %v", data)
	}
	pr.rdb.Set(ctx, comments_id, data, 15*time.Minute)

	return comments

}

// Insert article in database and delete {Id}:list from cache.
func (pr *PostRepository) InsertComment(ctx context.Context, comment models.Comment, author int) {
	comments_id := strconv.Itoa(comment.Post_id) + ":list"
	_, err := pr.db.Exec(
		ctx, "INSERT INTO Comments(Comment_content, Created_at, Post_id, Author) VALUES($1, $2, $3, $4)", comment.Comment_content, time.Now(), comment.Post_id, author)
	if err != nil {
		log.Println("Insert comment query error:", err)
	}
	log.Printf("Inserted comment in post with id: %v", comment.Post_id)
	pr.rdb.Del(ctx, comments_id)
}

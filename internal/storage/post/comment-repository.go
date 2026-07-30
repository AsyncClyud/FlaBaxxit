package poststorage

import (
	"blog/internal/models"
	"context"
	"encoding/json"
	"log"
	"time"
)

func (pr *PostRepository) GetArticleCommentsById(ctx context.Context, id int) (comment string) {
	var comments []models.Comment

	rows, err := pr.db.Query(ctx, "SELECT Comment_content, Created_at::text, Author FROM Comments WHERE Post_id = $1", id)
	if err != nil {
		log.Println("Rows error:", err)
		rows.Err()
	}
	defer rows.Close()

	for rows.Next() {
		var comment models.Comment
		err := rows.Scan(&comment.Comment_content, &comment.Created_at, &comment.Author)
		if err != nil {
			log.Println(err)
		}
		comments = append(comments, comment)
	}
	result, err := json.MarshalIndent(comments, "", " ")
	if err != nil {
		log.Fatalln(err)
	}

	return string(result)

}

func (pr *PostRepository) InsertComment(ctx context.Context, comment models.Comment, author int) {
	_, err := pr.db.Exec(
		ctx, "INSERT INTO Comments(Comment_content, Created_at, Post_id, Author) VALUES($1, $2, $3, $4)", comment.Comment_content, time.Now(), comment.Post_id, author)
	if err != nil {
		log.Println("Insert comment query error:", err)
	}
	log.Printf("Inserted comment in post with id: %v", comment.Post_id)
}

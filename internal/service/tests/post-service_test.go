package tests

import (
	"flabaxxit/internal/models"
	service "flabaxxit/internal/service/post"
	"net/http"
	"testing"
)

func TestArticleValidation(t *testing.T) {
	tests := []struct {
		test_name string
		article   models.Article
		want_code int
	}{{
		test_name: "valid article",
		article:   models.Article{Title: "Matadora linux", Content: "Matadora linux is the best!!!!!!!"},
		want_code: http.StatusOK,
	},
		{
			test_name: "short title",
			article:   models.Article{Title: "Ma", Content: "Matadora linux!!!"},
			want_code: http.StatusBadRequest,
		},
		{
			test_name: "short content",
			article:   models.Article{Title: "Simbios", Content: "11"},
			want_code: http.StatusUnprocessableEntity,
		},
	}

	ps := &service.PostService{}

	for _, tt := range tests {
		t.Run(tt.test_name, func(t *testing.T) {
			got := ps.ValidateArticle(tt.article)
			if got != tt.want_code {
				t.Errorf("ValidateUserData() = %d; want = %d", got, tt.want_code)
			}
		})
	}
}

func TestCommentValidation(t *testing.T) {
	tests := []struct {
		test_name string
		comment   models.Comment
		want_code int
	}{{
		test_name: "valid comment",
		comment:   models.Comment{Comment_content: "Cool article!!!!!"},
		want_code: http.StatusOK,
	},
		{
			test_name: "null comment content",
			comment:   models.Comment{Comment_content: ""},
			want_code: http.StatusBadRequest,
		},
	}

	ps := &service.PostService{}

	for _, tt := range tests {
		t.Run(tt.test_name, func(t *testing.T) {
			got := ps.ValidateComment(tt.comment)
			if got != tt.want_code {
				t.Errorf("ValidateComment() = %d; want = %d", got, tt.want_code)
			}
		})
	}
}

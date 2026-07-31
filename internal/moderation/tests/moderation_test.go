package tests

import (
	"flabaxxit/internal/moderation"
	"testing"
)

func TestArticleValidation(t *testing.T) {
	tests := []struct {
		test_name string
		article   string
		want      bool
	}{{
		test_name: "valid article RU",
		article:   "Привет!!",
		want:      false,
	},
		{
			test_name: "valid article EN",
			article:   "Hello!",
			want:      false,
		},
		{
			test_name: "Bad words RU",
			article:   "Сука блядь бля",
			want:      true,
		},
		{
			test_name: "Bad words EN",
			article:   "Slut cum fuck",
			want:      true,
		},
		{
			test_name: "Bad words EN with obfuscate",
			article:   "N1gger",
			want:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.test_name, func(t *testing.T) {
			got := moderation.ModerateText(tt.article)
			if got != tt.want {
				t.Errorf("ValidateUserData() = %v; want = %v", got, tt.want)
			}
		})
	}
}

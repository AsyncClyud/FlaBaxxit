package tests

import (
	"flabaxxit/internal/models"
	service "flabaxxit/internal/service/user"
	"net/http"
	"testing"
)

func TestUserValidation(t *testing.T) {
	tests := []struct {
		test_name string
		user      models.User
		want_code int
	}{{
		test_name: "valid username",
		user:      models.User{Username: "Maximus2121", Password: "ALibabab11"},
		want_code: http.StatusOK,
	},
		{
			test_name: "short username",
			user:      models.User{Username: "Ma", Password: "ALibabab11"},
			want_code: http.StatusBadRequest,
		},
		{
			test_name: "bad username",
			user:      models.User{Username: "Mama!@@!", Password: "ALibabab11"},
			want_code: http.StatusNotAcceptable,
		},
		{
			test_name: "short password",
			user:      models.User{Username: "Maximus2121", Password: "1234"},
			want_code: http.StatusUnprocessableEntity,
		}}

	ur := &service.AuthService{}

	for _, tt := range tests {
		t.Run(tt.test_name, func(t *testing.T) {
			got := ur.ValidateUserData(tt.user)
			if got != tt.want_code {
				t.Errorf("ValidateUserData() = %d; want = %d", got, tt.want_code)
			}
		})
	}
}

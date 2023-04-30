package auth

import (
	"ekinoksgcc/model"
	"ekinoksgcc/repository"
	"testing"
)

func TestUserSignup(t *testing.T) {
	repository.InitDB()
	defer repository.DisposeDB()
	req := model.RequestUserRegister{
		User: model.User{
			UserFullName: "test",
			UserEmail:    "test@test.com",
		},
		Password:             "test",
		PasswordVerification: "test",
	}
	_, _, err := UserSignup(req)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}

func TestUserLogin(t *testing.T) {
	repository.InitDB()
	defer repository.DisposeDB()
	req := model.RequestUserLogin{
		UserEmail: "test@test.com",
		Password:  "test",
	}
	_, _, err := UserLogin(req)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}

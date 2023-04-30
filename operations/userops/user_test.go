package userops

import (
	"ekinoksgcc/model"
	"testing"
)

func TestGetUserFromDB(t *testing.T) {
	req := 1
	_, err := GetUserFromDB(req)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}

func TestCreateUserInDB(t *testing.T) {
	req := model.User{
		UserFullName: "test",
		UserEmail:    "test@test.com",
	}

	_, err := CreateUserInDB(req)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}

func TestGetRoleByUserIDFromDB(t *testing.T) {

	req := 1
	_, err := GetRoleByUserIDFromDB(req)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}

func TestGetUserRoleFromDB(t *testing.T) {

	req := 1
	_, err := GetUserRoleFromDB(req)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}

func TestAssignUserRoleInDB(t *testing.T) {

	req := model.UserRole{
		UserID: 1,
		RoleID: 1,
	}
	_, err := AssignUserRoleInDB(req)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}

func TestGetRoleFromDB(t *testing.T) {
	req := 1
	_, err := GetRoleFromDB(req)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}

func TestCreateRoleInDB(t *testing.T) {
	req := model.Role{
		RoleName: "test",
	}
	_, err := CreateRoleInDB(req)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}

func TestAddPasswordToDB(t *testing.T) {
	req := model.Password{
		PasswordID:    1,
		EncryptedData: "test",
	}
	_, err := AddPasswordToDB(req)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

}

func TestAddUserPasswordToDB(t *testing.T) {
	req := model.UserPassword{
		UserID:     1,
		PasswordID: 1,
	}
	_, err := AddUserPasswordToDB(req)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}

func TestGetPasswordFromDB(t *testing.T) {
	req := 1
	_, err := GetPasswordFromDB(req)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}

func TestGetUserPasswordFromDB(t *testing.T) {
	req := 1
	_, err := GetUserPasswordFromDB(req)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}

func TestGetPasswordByUserIDFromDB(t *testing.T) {
	req := 1
	_, err := GetPasswordByUserIDFromDB(req)
	if err != nil {
		t.Errorf("Error: %v", err)
	}

}

func TestGetUserByUserEmailFromDB(t *testing.T) {
	req := "test@test.com"
	_, err := GetUserByUserEmailFromDB(req)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}

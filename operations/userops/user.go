package userops

import (
	"ekinoksgcc/model"
	"ekinoksgcc/repository"
	"log"
)

//Database operations for user and role

func GetUserFromDB(userID int) (model.User, error) {
	var user model.User
	err := repository.DB.QueryRow("SELECT userID, userFullName, userEmail FROM users WHERE userID=$1", userID).Scan(&user.UserID, &user.UserFullName, &user.UserEmail)
	if err != nil {
		log.Print(err)
	}
	return user, err
}

func CreateUserInDB(user model.User) (int, error) {
	var id int
	err := repository.DB.QueryRow("INSERT INTO users(userFullName, userEmail) VALUES($1, $2) RETURNING userID", user.UserFullName, user.UserEmail).Scan(&id)
	if err != nil {
		log.Print(err)
	}
	return id, err
}

func UpdateUserInDB(user model.User) (int, error) {
	_, err := repository.DB.Exec("UPDATE users SET userFullName=$1, userEmail=$2 WHERE userID=$3", user.UserFullName, user.UserEmail, user.UserID)
	if err != nil {
		log.Print(err)
	}
	return int(user.UserID), err
}

func DeleteUserFromDB(userID int) (int, error) {
	_, err := repository.DB.Exec("DELETE FROM users WHERE userID=$1", userID)
	if err != nil {
		log.Print(err)
	}
	return userID, err
}

func GetRoleByUserIDFromDB(userID int) (model.Role, error) {
	var userRole model.UserRole
	err := repository.DB.QueryRow("SELECT userRoleID, userID, roleID FROM userRole WHERE userID=$1", userID).Scan(&userRole.UserRoleID, &userRole.UserID, &userRole.RoleID)
	if err != nil {
		log.Print(err)
	}
	var role model.Role
	err = repository.DB.QueryRow("SELECT roleID, roleName FROM roles WHERE roleID=$1", userRole.RoleID).Scan(&role.RoleID, &role.RoleName)
	if err != nil {
		log.Print(err)
	}
	return role, err
}

func GetUserRoleFromDB(userRoleID int) (model.UserRole, error) {
	var userRole model.UserRole
	err := repository.DB.QueryRow("SELECT userRoleID, userID, roleID FROM userRole WHERE userID=$1", userRoleID).Scan(&userRole.UserRoleID, &userRole.UserID, &userRole.RoleID)
	if err != nil {
		log.Print(err)
	}
	return userRole, err
}

func AssignUserRoleInDB(userRole model.UserRole) (int, error) {
	var id int
	err := repository.DB.QueryRow("INSERT INTO userRole(userID, roleID) VALUES($1, $2) RETURNING userRoleID", userRole.UserID, userRole.RoleID).Scan(&id)
	if err != nil {
		log.Print(err)
	}
	return id, err
}

func UpdateUserRoleInDB(userRole model.UserRole) (int, error) {
	_, err := repository.DB.Exec("UPDATE userRole SET roleID=$1 WHERE userID=$2", userRole.RoleID, userRole.UserID)
	if err != nil {
		log.Print(err)
	}
	return int(userRole.UserRoleID), err
}

func GetRoleFromDB(roleID int) (model.Role, error) {
	var role model.Role
	err := repository.DB.QueryRow("SELECT roleID, roleName FROM roles WHERE roleID=$1", roleID).Scan(&role.RoleID, &role.RoleName)
	if err != nil {
		log.Print(err)
	}
	return role, err
}

func CreateRoleInDB(role model.Role) (int, error) {
	var id int
	err := repository.DB.QueryRow("INSERT INTO roles(roleName) VALUES($1) RETURNING roleID", role.RoleName).Scan(&id)
	if err != nil {
		log.Print(err)
	}
	return id, err
}

func UpdateRoleInDB(role model.Role) (int, error) {
	_, err := repository.DB.Exec("UPDATE roles SET roleName=$1 WHERE roleID=$2", role.RoleName, role.RoleID)
	if err != nil {
		log.Print(err)
	}
	return int(role.RoleID), err
}

func DeleteRoleFromDB(roleID int) (int, error) {
	_, err := repository.DB.Exec("DELETE FROM roles WHERE roleID=$1", roleID)
	if err != nil {
		log.Print(err)
	}
	return roleID, err
}

func AddPasswordToDB(password model.Password) (int, error) {
	var id int
	err := repository.DB.QueryRow("INSERT INTO password(encryptedData) VALUES($1) RETURNING passwordID", password.EncryptedData).Scan(&id)
	if err != nil {
		log.Print(err)
	}
	return id, err
}

func AddUserPasswordToDB(userPassword model.UserPassword) (int, error) {
	var id int
	err := repository.DB.QueryRow("INSERT INTO userPassword(userID, passwordID) VALUES($1, $2) RETURNING userPasswordID", userPassword.UserID, userPassword.PasswordID).Scan(&id)
	if err != nil {
		log.Print(err)
	}
	return id, err
}

func GetPasswordFromDB(passwordID int) (model.Password, error) {
	var password model.Password
	err := repository.DB.QueryRow("SELECT passwordID, encryptedData FROM password WHERE passwordID=$1", passwordID).Scan(&password.PasswordID, &password.EncryptedData)
	if err != nil {
		log.Print(err)
	}
	return password, err
}

func GetUserPasswordFromDB(userPasswordID int) (model.UserPassword, error) {
	var userPassword model.UserPassword
	err := repository.DB.QueryRow("SELECT userPasswordID, userID, passwordID FROM userPassword WHERE userPasswordID=$1", userPasswordID).Scan(&userPassword.UserPasswordID, &userPassword.UserID, &userPassword.PasswordID)
	if err != nil {
		log.Print(err)
	}
	return userPassword, err
}

func GetPasswordByUserIDFromDB(userID int) (model.Password, error) {
	var userPassword model.UserPassword
	err := repository.DB.QueryRow("SELECT userPasswordID, userID, passwordID FROM userPassword WHERE userID=$1", userID).Scan(&userPassword.UserPasswordID, &userPassword.UserID, &userPassword.PasswordID)
	if err != nil {
		log.Print(err)
	}
	var password model.Password
	err = repository.DB.QueryRow("SELECT passwordID, encryptedData FROM password WHERE passwordID=$1", userPassword.PasswordID).Scan(&password.PasswordID, &password.EncryptedData)
	if err != nil {
		log.Print(err)
	}
	return password, err
}

func UpdatePasswordInDB(password model.Password) (int, error) {
	_, err := repository.DB.Exec("UPDATE password SET encryptedData=$1 WHERE passwordID=$2", password.EncryptedData, password.PasswordID)
	if err != nil {
		log.Print(err)
	}
	return int(password.PasswordID), err
}

func DeletePasswordFromDB(passwordID int) (int, error) {
	_, err := repository.DB.Exec("DELETE FROM password WHERE passwordID=$1", passwordID)
	if err != nil {
		log.Print(err)
	}
	return passwordID, err
}

func DeleteUserPasswordFromDB(userPasswordID int) (int, error) {
	_, err := repository.DB.Exec("DELETE FROM userPassword WHERE userPasswordID=$1", userPasswordID)
	if err != nil {
		log.Print(err)
	}
	return userPasswordID, err
}

func GetPasswordByEncryptedData(encryptedData string) (model.Password, error) {
	var password model.Password
	err := repository.DB.QueryRow("SELECT passwordID, encryptedData FROM password WHERE encryptedData=$1", encryptedData).Scan(&password.PasswordID, &password.EncryptedData)
	if err != nil {
		log.Print(err)
	}
	return password, err
}

func GetUserByUserEmailFromDB(userEmail string) (model.User, error) {
	var user model.User
	err := repository.DB.QueryRow("SELECT userID, userEmail, userFullName, FROM users WHERE userEmail=$1", userEmail).Scan(&user.UserID, &user.UserEmail, &user.UserFullName)
	if err != nil {
		log.Print(err)
	}
	return user, err
}

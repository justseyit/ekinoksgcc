package auth

import (
	"database/sql"
	"ekinoksgcc/model"
	"ekinoksgcc/operations/userops"
	"errors"
	"log"
	"net/http"
	"net/mail"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var secretKey = []byte("ekinoksgcc")

func getHash(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func CheckJWTValidity(tokenString string, user model.User) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		log.Println(err)
		return err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims["user_id"] != user.UserID {
			return errors.New("Invalid user id")
		}
		if claims["user_full_name"] != user.UserFullName {
			return errors.New("Invalid user full name")
		}
		if claims["user_email"] != user.UserEmail {
			return errors.New("Invalid user email")
		}
	} else {
		return errors.New("Invalid token")
	}
	return nil
}

func UserSignup(req model.RequestUserRegister) (string, time.Time, error) {
	if req.Password != req.PasswordVerification {
		return "", time.Time{}, errors.New("Passwords don't match")
	}

	if req.User.UserEmail == "" || req.User.UserFullName == "" {
		return "", time.Time{}, errors.New("User email or full name cannot be empty")
	}

	_, emailErr := mail.ParseAddress(req.User.UserEmail)
	if emailErr != nil {
		return "", time.Time{}, errors.New("Invalid email address")
	}

	hash := getHash([]byte(req.Password))
	var user model.User
	user.UserEmail = req.User.UserEmail
	user.UserFullName = req.User.UserFullName

	pwdID, err := userops.AddPasswordToDB(model.Password{EncryptedData: hash})
	if err != nil {
		return "", time.Time{}, err
	}
	usrID, err1 := userops.CreateUserInDB(user)
	if err1 != nil {
		return "", time.Time{}, err1
	}

	_, err2 := userops.AddUserPasswordToDB(model.UserPassword{UserID: usrID, PasswordID: pwdID})
	if err2 != nil {
		return "", time.Time{}, err2
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := jwt.MapClaims{
		"user_id":         user.UserID,
		"user_full_name":  user.UserFullName,
		"user_email":      user.UserEmail,
		"user_role":       "role.RoleName",
		"expiration_time": expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)

	//TODO: Assign customer role to user
	return tokenString, expirationTime, nil
}

func UserLogin(req model.RequestUserLogin) (string, time.Time, error) {
	_, emailErr := mail.ParseAddress(req.UserEmail)

	if emailErr != nil {
		return "", time.Time{}, errors.New("Invalid email address")
	}
	user, err := userops.GetUserByUserEmailFromDB(req.UserEmail)
	if err == nil {
		if err == sql.ErrNoRows {
			return "", time.Time{}, errors.New("User not found")
		}
		return "", time.Time{}, err
	}
	password, err := userops.GetPasswordByUserIDFromDB(user.UserID)
	if err == nil {
		if err == sql.ErrNoRows {
			return "", time.Time{}, errors.New("Password not found")
		}
		return "", time.Time{}, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(password.EncryptedData), []byte(req.Password))
	if err != nil {
		return "", time.Time{}, errors.New("Wrong password")
	}

	role, err := userops.GetRoleByUserIDFromDB(user.UserID)
	if err == nil {
		if err == sql.ErrNoRows {
			return "", time.Time{}, errors.New("Role not found")
		}
		return "", time.Time{}, err
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := jwt.MapClaims{
		"user_id":         user.UserID,
		"user_full_name":  user.UserFullName,
		"user_email":      user.UserEmail,
		"user_role":       role.RoleName,
		"expiration_time": expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err1 := token.SignedString(secretKey)

	return tokenString, expirationTime, err1
}

func UserLogout(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Expires: time.Now(),
	})
}

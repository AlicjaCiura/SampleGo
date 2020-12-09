package model

import (
	"crypto/sha512"
	"database/sql"
	"encoding/base64"
	"fmt"

	"os"

	"github.com/withmandala/go-log"

	"time"
)

const passwordSalt = "a99VVoWzmd1C9ujcitK0fIVNE0I5I61AC47C852RoLTsHDyLCltvP+ZHEkIl/2hkzTOW90c3ZEjtYRkdfTWJ1Q=="

//User is object with info about user
type User struct {
	ID        int
	Email     string
	Password  string
	FirstName string
	LastName  string
	LastLogin *time.Time
}

//Login fucntion to login by emaila and password, return User or error
func Login(email, password string) (*User, error) {
	log := log.New(os.Stdout)
	result := &User{}
	hasher := sha512.New()
	hasher.Write([]byte(passwordSalt))
	hasher.Write([]byte(email))
	hasher.Write([]byte(password))
	pwd := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	log.Infof("Login of users: %v\n", email)
	row := db.QueryRow(`
		SELECT id, email, firstname, lastname, last_login_date
		FROM public.user
		WHERE email = $1
		  AND password = $2`, email, pwd)
	err := row.Scan(&result.ID, &result.Email, &result.FirstName, &result.LastName, &result.LastLogin)
	switch {
	case err == sql.ErrNoRows:
		log.Debugf("User not found, %v", err)
		return nil, fmt.Errorf("User not found")
	case err != nil:
		return nil, err
	}
	t := time.Now()
	_, err = db.Exec(`
		UPDATE public.user
		SET last_login_date = $1
		WHERE id = $2`, t, result.ID)
	if err != nil {
		log.Infof("Failed to update login time for user %v to %v: %v", result.Email, t, err)
	}
	return result, nil
}

//AddNewUser function is a function to add new user to db
func AddNewUser(email, firstName, lastName, password string) (*User, error) {
	log := log.New(os.Stdout)
	result := &User{}
	hasher := sha512.New()
	hasher.Write([]byte(passwordSalt))
	hasher.Write([]byte(email))
	hasher.Write([]byte(password))
	pwd := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	log.Infof("Add new user: %v\n", email)
	row := db.QueryRow(`
	INSERT INTO public.user (email, firstname, lastname, password)
	VALUES ($1, $2, $3, $4)`, email, firstName, lastName, pwd)
	err := row.Scan(&result.ID, &result.Email, &result.FirstName, &result.LastName)
	switch {
	case err == sql.ErrNoRows:
		log.Errorf("User not found, %v", err)
		return nil, fmt.Errorf("User exists already, %v", err)
	case err != nil:
		return nil, err
	}
	return result, nil
}

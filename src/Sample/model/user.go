package model

import (
	"crypto/sha512"
	"database/sql"
	"encoding/base64"
	"fmt"
	"log"

	"time"
)

const passwordSalt = "a99VVoWzmd1C9ujcitK0fIVNE0I5I61AC47C852RoLTsHDyLCltvP+ZHEkIl/2hkzTOW90c3ZEjtYRkdfTWJ1Q=="

type User struct {
	ID        int
	Email     string
	Password  string
	FirstName string
	LastName  string
	LastLogin *time.Time
}

func (t *User) String() string {
	return t.LastLogin.Format("2020-12-12 12:00")
}

func Login(email, password string) (*User, error) {
	result := &User{}
	hasher := sha512.New()
	hasher.Write([]byte(passwordSalt))
	hasher.Write([]byte(email))
	hasher.Write([]byte(password))
	pwd := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	log.Printf("Data of users: %v\n", email)
	row := db.QueryRow(`
		SELECT id, email, firstname, lastname, last_login_date
		FROM public.user
		WHERE email = $1
		  AND password = $2`, email, pwd)
	err := row.Scan(&result.ID, &result.Email, &result.FirstName, &result.LastName, &result.LastLogin)
	switch {
	case err == sql.ErrNoRows:
		return nil, fmt.Errorf("User not found, %v", err)
	case err != nil:
		return nil, err
	}
	t := time.Now()
	_, err = db.Exec(`
		UPDATE public.user
		SET last_login_date = $1
		WHERE id = $2`, t, result.ID)
	if err != nil {
		log.Printf("Failed to update login time for user %v to %v: %v", result.Email, t, err)
	}
	return result, nil
}
func AddNewUser(email, firstName, lastName, password string) (*User, error) {
	result := &User{}
	hasher := sha512.New()
	hasher.Write([]byte(passwordSalt))
	hasher.Write([]byte(email))
	hasher.Write([]byte(password))
	pwd := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	log.Printf("Add new user: %v\n", email)
	row := db.QueryRow(`
	INSERT INTO public.user (email, firstname, lastname, password)
	VALUES ($1, $2, $3, $4)`, email, firstName, lastName, pwd)
	err := row.Scan(&result.ID, &result.Email, &result.FirstName, &result.LastName)
	switch {
	case err == sql.ErrNoRows:
		return nil, fmt.Errorf("User exists already, %v", err)
	case err != nil:
		return nil, err
	}
	return result, nil
}

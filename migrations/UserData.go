package main

import (
	"database/sql"
	"log"
	"time"
)

type UserData struct {
	ID             int       `json:"id"`
	UserName       string    `json:"user_name"`
	Description    string    `json:"description"`
	Email          string    `json:"email"`
	Phone          string    `json:"phone"`
	AvatarUrl      string    `json:"avatar_url"`
	Status         string    `json:"status"`
	Role           string    `json:"role"`
	PasswordHash   string    `json:"password_hash"`
	DateOfBirth    time.Time `json:"date_of_birth"`
	PrivacySettings string   `json:"privacy_settings"`
	IsActive       bool      `json:"is_active"`
	LastLogin      time.Time `json:"last_login"`
	ConfirmationToken string `json:"confirmation_token"`
	SocialProfiles string    `json:"social_profiles"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// CreateUserData creates a new user_data record in the database
func CreateUserData(db *sql.DB, userData *UserData) error {
	const query = `
		INSERT INTO user_data (username, description, email, phone, avatar, status, role, password_hash, date_of_birth, privacy_settings, is_active, last_login, confirmation_token, social_profiles)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
		RETURNING id
	`
	err := db.QueryRow(query, userData.UserName, userData.Description, userData.Email, userData.Phone, userData.AvatarUrl, userData.Status, userData.Role, userData.PasswordHash, userData.DateOfBirth, userData.PrivacySettings, userData.IsActive, userData.LastLogin, userData.ConfirmationToken, userData.SocialProfiles).Scan(&userData.ID)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// ReadUserData retrieves a user_data record from the database
func ReadUserData(db *sql.DB, id int) (*UserData, error) {
	userData := &UserData{}
	err := db.QueryRow(`SELECT * FROM user_data WHERE id = $1`, id).Scan(
		&userData.ID,
	)
	if err != nil {
		return nil, err
	}
	return userData, nil
}

// UpdateUserData updates a user_data record in the database
func UpdateUserData(db *sql.DB, userData *UserData) error {
	const query = `
		UPDATE user_data SET
			username = $2,
			description = $3,
			email = $4,
			phone = $5,
			avatar = $6,
			status = $7,
			role = $8,
			password_hash = $9,
			date_of_birth = $10,
			privacy_settings = $11,
			is_active = $12,
			last_login = $13,
			confirmation_token = $14,
			social_profiles = $15
		WHERE id = $1
	`
	_, err := db.Exec(query, userData.ID, userData.UserName, userData.Description, userData.Email, userData.Phone, userData.AvatarUrl, userData.Status, userData.Role, userData.PasswordHash, userData.DateOfBirth, userData.PrivacySettings, userData.IsActive, userData.LastLogin, userData.ConfirmationToken, userData.SocialProfiles)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// DeleteUserData deletes a user_data record from the database
func DeleteUserData(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM user_data WHERE id = $1", id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

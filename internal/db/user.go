package db

import (
	"database/sql"
	// "errors"
	"fmt"
	"log"
	"time"

	"github.com/VladislavSCV/GoCode/pkg"
)

type UserData struct {
	ID               int       `json:"id"`
	UserName         string    `json:"user_name"`
	Description      string    `json:"description"`
	Email            string    `json:"email"`
	Phone            string    `json:"phone"`
	AvatarUrl        string    `json:"avatar_url"`
	Status           string    `json:"status"`
	Role             string    `json:"role"`
	PasswordHash     string    `json:"password_hash"`
	DateOfBirth      time.Time `json:"date_of_birth"`
	PrivacySettings  string    `json:"privacy_settings"`
	IsActive         bool      `json:"is_active"`
	LastLogin        time.Time `json:"last_login"`
	ConfirmationToken string   `json:"confirmation_token"`
	SocialProfiles   string    `json:"social_profiles"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type UserDataLoginResp struct {
	UserName         string    `json:"user_name"`
	Description      string    `json:"description"`
	Email            string    `json:"email"`
	Phone            string    `json:"phone"`
	AvatarUrl        string    `json:"avatar_url"`
	Status           string    `json:"status"`
	Role             string    `json:"role"`
	DateOfBirth      time.Time `json:"date_of_birth"`
	PrivacySettings  string    `json:"privacy_settings"`
	IsActive         bool      `json:"is_active"`
	LastLogin        time.Time `json:"last_login"`
	ConfirmationToken string   `json:"confirmation_token"`
	SocialProfiles   string    `json:"social_profiles"`
}

// Insert into users table
func SignUp(db *sql.DB, UserName, Email, Phone, AvatarUrl string, Role string, PasswordHash string, DateOfBirth string) error {
	const query = `
	INSERT INTO user_data (username, email, phone, avatar, role, password_hash, date_of_birth)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
`
	_, err := db.Exec(query, UserName, Email, Phone, AvatarUrl, Role, PasswordHash, DateOfBirth)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func Login(db *sql.DB, email, password string) (*UserDataLoginResp, error) {
	if db == nil {
		return nil, fmt.Errorf("db cannot be nil")
	}

	checkedPassword, err := pkg.CheckAndHashPassword(password)
	if err != nil {
		return nil, fmt.Errorf("error checking or hashing password: %v", err)
	}

	const query = `SELECT  COALESCE(username, ''), COALESCE(description, ''), email, phone, avatar, COALESCE(status, ''), 
	role, COALESCE(date_of_birth::timestamp,'1970-01-01 00:00:00'), COALESCE(privacy_settings::text, ''), is_active, 
	COALESCE(last_login::timestamp, '1970-01-01 00:00:00'), COALESCE(confirmation_token::text, ''), COALESCE(social_profiles::text, '') FROM user_data WHERE email = $1 AND password_hash = $2`
	var userDataLogin UserDataLoginResp
	err = db.QueryRow(query, email, checkedPassword).Scan(
		&userDataLogin.UserName, &userDataLogin.Description, &userDataLogin.Email, &userDataLogin.Phone, &userDataLogin.AvatarUrl, &userDataLogin.Status,
		&userDataLogin.Role, &userDataLogin.DateOfBirth, &userDataLogin.PrivacySettings, &userDataLogin.IsActive,
		&userDataLogin.LastLogin, &userDataLogin.ConfirmationToken, &userDataLogin.SocialProfiles,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error during login: %v", err)
	}

	return &userDataLogin, nil
}

func UpdateLastActiveAt(db *sql.DB, userID int) error {
    _, err := db.Exec("UPDATE user_data SET last_active_at = $1 WHERE id = $2", time.Now(), userID)
    if err != nil {
        return fmt.Errorf("failed to update last_active_at: %v", err)
    }
    return nil
}

// GetUserData retrieves a user_data record from the database
func GetUserData(db *sql.DB, id int) (*UserData, error) {
	userData := &UserData{}
	err := db.QueryRow(`SELECT id, username, description, email, phone, avatar, status, role, password_hash, date_of_birth, privacy_settings, is_active, last_login, confirmation_token, social_profiles, created_at, updated_at FROM user_data WHERE id = $1`, id).Scan(
		&userData.ID,
		&userData.UserName,
		&userData.Description,
		&userData.Email,
		&userData.Phone,
		&userData.AvatarUrl,
		&userData.Status,
		&userData.Role,
		&userData.PasswordHash,
		&userData.DateOfBirth,
		&userData.PrivacySettings,
		&userData.IsActive,
		&userData.LastLogin,
		&userData.ConfirmationToken,
		&userData.SocialProfiles,
		&userData.CreatedAt,
		&userData.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return userData, nil
}

func UpdateUserData(db *sql.DB, userData *UserData) error {
	const query = `
		UPDATE user_data SET
			user_name = $2,
			description = $3,
			email = $4,
			phone = $5,
			avatar_url = $6,
			status = $7,
			role = $8,
			password_hash = $9,
			date_of_birth = $10,
			privacy_settings = $11,
			is_active = $12,
			last_login = $13,
			confirmation_token = $14,
			social_profiles = $15,
			updated_at = $16
		WHERE id = $1
	`
	_, err := db.Exec(query, userData.ID, userData.UserName, userData.Description, userData.Email, userData.Phone, userData.AvatarUrl, userData.Status, userData.Role, userData.PasswordHash, userData.DateOfBirth, userData.PrivacySettings, userData.IsActive, userData.LastLogin, userData.ConfirmationToken, userData.SocialProfiles, time.Now())
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// Functions for updating individual fields

func UpdateUserPassword(db *sql.DB, id int, passwordHash string) error {
	_, err := db.Exec("UPDATE user_data SET password_hash = $1, updated_at = $2 WHERE id = $3", passwordHash, time.Now(), id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func UpdateUserEmail(db *sql.DB, id int, email string) error {
	_, err := db.Exec("UPDATE user_data SET email = $1, updated_at = $2 WHERE id = $3", email, time.Now(), id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func UpdateUserPhone(db *sql.DB, id int, phone string) error {
	_, err := db.Exec("UPDATE user_data SET phone = $1, updated_at = $2 WHERE id = $3", phone, time.Now(), id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func UpdateUserStatus(db *sql.DB, id int, status string) error {
	_, err := db.Exec("UPDATE user_data SET status = $1, updated_at = $2 WHERE id = $3", status, time.Now(), id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func UpdateUserAvatar(db *sql.DB, id int, avatarUrl string) error {
	_, err := db.Exec("UPDATE user_data SET avatar_url = $1, updated_at = $2 WHERE id = $3", avatarUrl, time.Now(), id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func UpdateUserName(db *sql.DB, id int, userName string) error {
	_, err := db.Exec("UPDATE user_data SET user_name = $1, updated_at = $2 WHERE id = $3", userName, time.Now(), id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func UpdateUserDescription(db *sql.DB, id int, description string) error {
	_, err := db.Exec("UPDATE user_data SET description = $1, updated_at = $2 WHERE id = $3", description, time.Now(), id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func UpdateUserSkills(db *sql.DB, id int, skills string) error {
	_, err := db.Exec("UPDATE user_data SET skills = $1 WHERE id = $2", skills, id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

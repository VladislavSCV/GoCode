package migrations

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
func GetUserData(db *sql.DB, id int) (*UserData, error) {
	userData := &UserData{}
	err := db.QueryRow(`SELECT * FROM user_data WHERE id = $1`, id).Scan(
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

func UpdateUserPassword(db *sql.DB, id int, passwordHash string) error {
	_, err := db.Exec("UPDATE user_data SET password_hash = $1 WHERE id = $2", passwordHash, id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func UpdateUserEmail(db *sql.DB, id int, email string) error {
	_, err := db.Exec("UPDATE user_data SET email = $1 WHERE id = $2", email, id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func UpdateUserPhone(db *sql.DB, id int, phone string) error {
	_, err := db.Exec("UPDATE user_data SET phone = $1 WHERE id = $2", phone, id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func UpdateUserStatus(db *sql.DB, id int, status string) error {
	_, err := db.Exec("UPDATE user_data SET status = $1 WHERE id = $2", status, id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func UpdateUserAvatar(db *sql.DB, id int, avatarUrl string) error {
	_, err := db.Exec("UPDATE user_data SET avatar = $1 WHERE id = $2", avatarUrl, id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func UpdateUserName(db *sql.DB, id int, userName string) error {
	_, err := db.Exec("UPDATE user_data SET username = $1 WHERE id = $2", userName, id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func UpdateUserDescription(db *sql.DB, id int, description string) error {
	_, err := db.Exec("UPDATE user_data SET description = $1 WHERE id = $2", description, id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}


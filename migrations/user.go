package migrations

import (
	"database/sql"
	"log"
	"time"
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

// Insert into users table
func SignUp(db *sql.DB, userData *UserData) error {
	const query = `
		INSERT INTO user_data VALUES
			user_name = $1,
			description = $2,
			email = $3,
			phone = $4,
			avatar_url = $5,
			status = $6,
			role = $7,
			password_hash = $8,
			date_of_birth = $9,
			privacy_settings = $10,
			is_active = $11,
			last_login = $12,
			confirmation_token = $13,
			social_profiles = $14,
			updated_at = $15

	`
	_, err := db.Exec(query, userData.UserName, userData.Description, userData.Email, userData.Phone, userData.AvatarUrl, userData.Status, userData.Role, userData.PasswordHash, userData.DateOfBirth, userData.PrivacySettings, userData.IsActive, userData.LastLogin, userData.ConfirmationToken, userData.SocialProfiles, time.Now())
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func Login(db *sql.DB, email, password string) (bool, error) {
	isAccepted := false
	const query = `SELECT 1 FROM user_data WHERE email = $1 AND password_hash = $2`
	err := db.QueryRow(query, email, password).Scan(&isAccepted)
	if err != nil {
		log.Println("ERROR in LOGINING:", err)
		return false, err
	}
	return isAccepted, nil
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

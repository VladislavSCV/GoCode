package main

import (
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

// func (ud &UserData) getName
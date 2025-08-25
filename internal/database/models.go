/*
 * GORM Models
 * https://gorm.io/docs/models.html
 */
package database

import (
	"time"

	"gorm.io/gorm"
)

type GuildConfig struct {
	gorm.Model

	ID            string `gorm:"primaryKey"`
	Name          string
	OwnerID       string
	OwnerUsername string

	// if configured over /setup
	WelcomeChannelID string
	AutoVCCategoryID string
	AutoVCChannelID  string
}

type User struct {
	gorm.Model

	ID            string `gorm:"primaryKey"`
	Username      string
	Discriminator string
	GlobalName    string

	// OAuth2 Variables
	IsVerifiedOverOAuth2 bool
	Email                string
	Locale               string
	AccessToken          string
	RefreshToken         string
	RefreshTokenExpiry   time.Time
}

type Message struct {
	gorm.Model

	ID             string `gorm:"primaryKey"`
	ChannelID      string
	GuildID        string
	Content        string
	Timestamp      *time.Time
	AuthorID       string
	AuthorUsername string
}

package admin

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Admin represents an admin user
type Admin struct {
	ID                            uuid.UUID       `gorm:"type:uuid;primaryKey" json:"_id"`
	FirstName                     string          `gorm:"not null" json:"firstName"`
	LastName                      string          `gorm:"not null" json:"lastName"`
	UserName                      string          `json:"userName,omitempty"`
	Mobile                        string          `json:"mobile,omitempty"`
	EmailID                       string          `gorm:"not null;unique" json:"emailId"` // Ensure unique email
	Password                      string          `gorm:"not null" json:"-"`              // Exclude password from JSON
	Photo                         string          `json:"photo,omitempty"`
	EmailVerificationStatus       bool            `gorm:"default:true" json:"emailVerificationStatus"`
	VerificationToken             string          `json:"verificationToken,omitempty"`
	VerificationTokenCreationTime time.Time       `json:"verificationTokenCreationTime,omitempty"`
	DateOfBirth                   time.Time       `json:"dateOfBirth,omitempty"`
	Gender                        string          `json:"gender,omitempty"`
	Website                       string          `json:"website,omitempty"`
	Address                       string          `json:"address,omitempty"`
	FbId                          string          `json:"fbId,omitempty"`
	TwitterId                     string          `json:"twitterId,omitempty"`
	InstagramId                   string          `json:"instagramId,omitempty"`
	GithubId                      string          `json:"githubId,omitempty"`
	Codepen                       string          `json:"codepen,omitempty"`
	Slack                         string          `json:"slack,omitempty"`
	SendOTPToken                  string          `json:"sendOTPToken,omitempty"`
	ForgotToken                   string          `json:"forgotToken,omitempty"`
	ForgotTokenCreationTime       time.Time       `json:"forgotTokenCreationTime,omitempty"`
	DeviceToken                   string          `json:"deviceToken,omitempty"`
	Device                        string          `json:"device,omitempty"`
	IsThemeDark                   bool            `json:"isThemeDark,omitempty"`
	AddedBy                       uuid.UUID       `gorm:"type:uuid" json:"addedBy,omitempty"` // Use UUID type for consistency
	CountryCode                   string          `json:"countryCode,omitempty"`
	TimeZone                      string          `json:"timeZone,omitempty"`
	DateFormat                    string          `json:"dateFormat,omitempty"`
	Currency                      string          `json:"currency,omitempty"`
	TableColumnSettings           json.RawMessage `json:"tableColumnSettings,omitempty"`
	Status                        bool            `gorm:"default:true" json:"status"`
	IsDeleted                     bool            `gorm:"default:false" json:"isDeleted"`
	CreatedAt                     time.Time       `gorm:"autoCreateTime" json:"createdAt,omitempty"`
	UpdatedAt                     time.Time       `gorm:"autoUpdateTime" json:"updatedAt,omitempty"`
}

// AdminRelation represents the self-referential relationship for Admin
type AdminRelation struct {
	ID      uuid.UUID `gorm:"type:uuid;primaryKey" json:"_id"`
	AddedBy uuid.UUID `gorm:"type:uuid" json:"addedBy"`
	Admin   Admin     `gorm:"foreignKey:AddedBy;references:ID"`
}

// BeforeCreate hook to set UUID if not provided
func (a *Admin) BeforeCreate(tx *gorm.DB) (err error) {
	if a.ID == uuid.Nil {
		a.ID = uuid.New()
	}
	return
}

package admin

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Admin represents an admin user
type Admin struct {
	ID                            uuid.UUID `gorm:"type:uuid;primaryKey" json:"_id"`
	FirstName                     string    `gorm:"not null" json:"firstName"`
	LastName                      string    `gorm:"not null" json:"lastName"`
	UserName                      string    `json:"userName,omitempty"`
	Mobile                        string    `json:"mobile,omitempty"`
	EmailID                       string    `gorm:"not null" json:"emailId"`
	Password                      string    `json:"password,omitempty"`
	Photo                         string    `json:"photo,omitempty"`
	EmailVerificationStatus       bool      `gorm:"default:true" json:"emailVerificationStatus"`
	VerificationToken             string    `json:"verificationToken,omitempty"`
	VerificationTokenCreationTime time.Time `json:"verificationTokenCreationTime,omitempty"`
	DateOfBirth                   time.Time `json:"dateOfBirth,omitempty"`
	Gender                        string    `json:"gender,omitempty"`
	Website                       string    `json:"website,omitempty"`
	Address                       string    `json:"address,omitempty"`
	Status                        bool      `gorm:"default:true" json:"status"`
	IsDeleted                     bool      `gorm:"default:false" json:"isDeleted"`
	CreatedAt                     time.Time `gorm:"autoCreateTime" json:"createdAt,omitempty"`
	UpdatedAt                     time.Time `gorm:"autoUpdateTime" json:"updatedAt,omitempty"`
}

// BeforeCreate hook to set UUID if not provided
func (a *Admin) BeforeCreate(tx *gorm.DB) (err error) {
	if a.ID == uuid.Nil {
		a.ID = uuid.New()
	}
	return
}

package models

import (
	"time"

	"gorm.io/gorm"
)

type Patient struct {
	ID               uint            `json:"id" gorm:"primaryKey"`
	FirstName        string          `json:"first_name" gorm:"not null"`
	LastName         string          `json:"last_name" gorm:"not null"`
	DateOfBirth      time.Time       `json:"date_of_birth"`
	SSN              string          `json:"ssn,omitempty" gorm:"unique;column:ssn"`
	Phone            string          `json:"phone"`
	Address          string          `json:"address"`
	EmergencyContact string          `json:"emergency_contact"`
	MedicalRecords   []MedicalRecord `json:"medical_records,omitempty"`
	CreatedAt        time.Time       `json:"created_at"`
	UpdatedAt        time.Time       `json:"updated_at"`
}

func (p *Patient) BeforeCreate(tx *gorm.DB) (err error) {
	if p.CreatedAt.IsZero() {
		p.CreatedAt = time.Now()
	}
	if p.UpdatedAt.IsZero() {
		p.UpdatedAt = time.Now()
	}
	return
}

func (p *Patient) BeforeUpdate(tx *gorm.DB) (err error) {
	p.UpdatedAt = time.Now()
	return
}

func (p *Patient) GetFullName() string {
	return p.FirstName + " " + p.LastName
}

func (p *Patient) GetAge() int {
	now := time.Now()
	years := now.Year() - p.DateOfBirth.Year()

	if now.Month() < p.DateOfBirth.Month() ||
		(now.Month() == p.DateOfBirth.Month() && now.Day() < p.DateOfBirth.Day()) {
		years--
	}

	return years
}

func (p *Patient) SanitizeForRole(role UserRole) *Patient {
	sanitized := *p

	switch role {
	case RoleNurse:
		sanitized.SSN = ""
	case RoleAdmin:
		return nil
	}

	return &sanitized
}

func (p *Patient) TableName() string {
	return "patients"
}
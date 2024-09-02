package entity

import "time"

type Sector struct {
	ID            string    `json:"id" db:"id"`
	Name          string    `json:"name" db:"name"`
	Localization  string    `json:"localization" db:"localization"`
	ResponsibleID string    `json:"responsible_id" db:"responsible_id"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
}

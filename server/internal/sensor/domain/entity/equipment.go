package entity

import "time"

type Equipment struct {
	ID           string    `json:"id" db:"id"`
	Name         string    `json:"name" db:"name"`
	Model        string    `json:"model" db:"model"`
	Manufacturer string    `json:"manufacturer" db:"manufacturer"`
	SectorID     string    `json:"sector_id" db:"sector_id"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}

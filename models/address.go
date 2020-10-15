package models

import (
	"database/sql"
	"time"
)

type AddressForm struct {
	IP       string        `json:"ip_address" gorm:"type:varchar(15);not null"`
	Mac      string        `json:"mac" gorm:"type:varchar(17);not null"`
	Hostname string        `json:"hostname" gorm:"type:varchar(255)"`
	Reserved bool          `json:"reserved" gorm:"type:bool"`
	PoolID   sql.NullInt32 `json:"pool_id" gorm:"type:int(11)"`
}

type Address struct {
	ID int `json:"id" gorm:"primary_key"`

	AddressForm

	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

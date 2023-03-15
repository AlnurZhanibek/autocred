package storage

import "gorm.io/gorm"

type ResidentialAddress struct {
	gorm.Model
	Country     string `gorm:"size:100;"`
	Address     string `gorm:"size:100;"`
	Region      string `gorm:"size:100;"`
	Area        string `gorm:"size:100;"`
	District    string `gorm:"size:100;"`
	City        string `gorm:"size:100;"`
	Street      string `gorm:"size:100;"`
	House       string `gorm:"size:100;"`
	Flat        string `gorm:"size:100;"`
	Kato        string `gorm:"size:100;"`
	Phone       string `gorm:"size:100;"`
	SecondPhone string `gorm:"size:100;"`
	WorkPhone   string `gorm:"size:100;"`
	Email       string `gorm:"size:100;"`
	ClientID    uint
}

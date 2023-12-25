package model

import "gorm.io/gorm"

type Doktor struct {
	gorm.Model
	Isim       string
	Soyisim    string
	Hastane    string
	Uzmanlik   string
	Randevular []DoktorRandevu `gorm:"foreignKey:DoktorID"`
}

type DoktorRandevu struct {
	gorm.Model
	DoktorID            uint
	RandevuTarihi       int
	RandevuHastaIsmi    string
	RandevuHastaSoyismi string
}

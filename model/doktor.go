/*
--------------------------------------------------------------------------

	Denizli Meslek Yüksek Okulu Bilgisayar Programcılığı
	2. Sınıf öğrencileri
	Umutcan Biler ve Muhammet Yasin Seden'nin
	Sistem Analizi ve Tasarımı dönem sonu projesi

--------------------------------------------------------------------------
*/
package model

import (
	"gorm.io/gorm"
)

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
	RandevuTarihi       string
	RandevuHastaIsmi    string
	RandevuHastaSoyismi string
}

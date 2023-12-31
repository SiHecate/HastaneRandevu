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

type Randevu struct {
	gorm.Model
	Tarih            string // Yıl gün ay saat
	HastaIsim        string
	HastaSoyisim     string
	DoktorIsim       string // Isim ve soyisim çünkü id alınacak
	DoktorSoyisim    string // Isim ve soyisim çünkü id alınacak
	RandevuBölüm     string
	HastaRahatsizlik string
}

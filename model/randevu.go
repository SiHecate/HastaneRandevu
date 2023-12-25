/*
--------------------------------------------------------------------------

	Denizli Meslek Yüksek Okulu Bilgisayar Programcılığı
	2. Sınıf öğrencisi Umutcan Biler'in Sistem Analizi dönem sonu projesi

--------------------------------------------------------------------------
*/
package model

import "gorm.io/gorm"

/*
	Randevu:
	Hastanın ismi soyismi alınacak hastanın rahatsızlığı öğrenilecek
	hangi doktordan randevu alındığı öğrenilecek
	randevu tarihi alınacak bu sayede doktorun hangi saatlerde dolu hangi saatlerde boş olduğu anlaşılacak
	doktora saat başına toplam 3 (üç) tane randevu verilecek yani 12:00 - 12:20 - 12:40 şeklinde bu saatler dışında herhangi bir şekilde randevu alınamayacak
	bunun kontrolünü nasıl yapacağım açıkcası hiç bir fikrim yok fakat bir şekilde halledeceğim

	ToDo:
		Uygulama sırasında büyük ihtimal eklemeler ve çıkarmalar olacak onu sonra düşüneceğim

*/

type Randevu struct {
	gorm.Model
	Tarih            int // Yıl gün ay saat
	HastaIsim        string
	HastaSoyisim     string
	DoktorIsim       string // Isim ve soyisim çünkü id alınacak
	DoktorSoyisim    string // Isim ve soyisim çünkü id alınacak
	RandevuBölüm     string
	HastaRahatsizlik string
}

/*
--------------------------------------------------------------------------

	Denizli Meslek Yüksek Okulu Bilgisayar Programcılığı
	2. Sınıf öğrencileri
	Umutcan Biler ve Muhammet Yasin Seden'nin
	Sistem Analizi ve Tasarımı dönem sonu projesi

--------------------------------------------------------------------------
*/
package database

import (
	"fmt"
	model "hastane-uyg/Model"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Conn *gorm.DB

func Connect() {
	Database()
	MigrateTables()
}

func Database() {
	dsn := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=disable TimeZone=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_TIMEZONE"),
	)

	var err error
	Conn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Print("Database hatasi")
		panic(err)
	}
}

func MigrateTables() error {
	err := Conn.AutoMigrate(
		&model.Doktor{},
		&model.Randevu{},
		&model.DoktorRandevu{},
	)
	if err != nil {
		return err
	}
	return nil
}

/*
--------------------------------------------------------------------------

	Denizli Meslek Yüksek Okulu Bilgisayar Programcılığı
	2. Sınıf öğrencisi Umutcan Biler'in Sistem Analizi dönem sonu projesi

--------------------------------------------------------------------------
*/
package main

import (
	"fmt"
	"hastane-uyg/core/database"
	"hastane-uyg/router"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

/*
	/
	ToDo:
	Doktor randevu alma BİTTİ
	Randevu bilgileri (tarih - doktor adı - muayene türü vb...) BİTTİ
	Hasta ekleme BİTTİ
	Hastaların önceden aldığı randevuların geçmişi BİTTİ

	Proje bittiğinde
	Postman üzerinden front tarafına document ekle BİTTİ
	Git'e at ve api'yı nasıl live'a çekebileceği araştır BİTTİ

*/

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Printf("Hello world")

	database.Connect()

	app := fiber.New()
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))
	router.Router(app)
	app.Listen(":8080")
}

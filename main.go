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
	Doktor randevu alma
	Randevu bilgileri (tarih - doktor adı - muayene türü vb...)
	Hasta ekleme (randevu sistem içerisinde olacak ekstra düşünmeye gerek olduğunu düşünmüyorum okul projesi olduğu için)
	Hastaların önceden aldığı randevuların geçmişi

	Mümkünse bu gece bitir yarın bu konu hakkında konuş
	Proje bittiğinde
	Postman üzerinden front tarafına document ekle
	Git'e at ve api'yı nasıl live'a çekebileceği araştır

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

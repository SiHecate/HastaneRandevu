/*
--------------------------------------------------------------------------

	Denizli Meslek Yüksek Okulu Bilgisayar Programcılığı
	2. Sınıf öğrencileri
	Umutcan Biler ve Muhammet Yasin Seden'nin
	Sistem Analizi ve Tasarımı dönem sonu projesi

--------------------------------------------------------------------------
*/
package router

import (
	"hastane-uyg/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Router(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, PATCH, DELETE",
		AllowCredentials: true,
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Sistem analizi okul projesi")
	})

	doktor := app.Group("/doktor")
	doktor.Post("ekle", controller.DoktorEkle)
	doktor.Post("guncelle", controller.DoktorGüncelle)
	doktor.Get("liste", controller.DoktorListe)
	doktor.Delete("sil", controller.DoktorSil)

	randevu := app.Group("/randevu")
	randevu.Post("ekle", controller.RandevuOluştur)
	randevu.Get("hasta_listesi", controller.HastaListesi)
	randevu.Get("doktor_listesi", controller.DoktorListesi)
	randevu.Get("randevular_doktor", controller.DoktorRandevuListesi)
	randevu.Post("randevular_hasta", controller.HastaRandevuListesi)

}

/*
--------------------------------------------------------------------------

	Denizli Meslek Yüksek Okulu Bilgisayar Programcılığı
	2. Sınıf öğrencileri
	Umutcan Biler ve Muhammet Yasin Seden'nin
	Sistem Analizi ve Tasarımı dönem sonu projesi

--------------------------------------------------------------------------
*/
package controller

import (
	"fmt"

	model "hastane-uyg/Model"

	"hastane-uyg/core/database"

	"github.com/gofiber/fiber/v2"
)

func RandevuOluştur(c *fiber.Ctx) error {
	var RandevuBilgi struct {
		Tarih            string `json:"tarih"`
		HastaIsim        string `json:"hasta_isim"`
		HastaSoyisim     string `json:"hasta_soyisim"`
		DoktorIsim       string `json:"doktor_isim"`
		DoktorSoyisim    string `json:"doktor_soyisim"`
		HastaRahatsizlik string `json:"hasta_rahatsizlik"`
	}

	if err := c.BodyParser(&RandevuBilgi); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request data: " + err.Error()})
	}

	existingDoktorlar, err := DoktorKontrol(RandevuBilgi.DoktorIsim, RandevuBilgi.DoktorSoyisim)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"doktor_error": "Doktor mevcut değil."})
	}

	var doktor model.Doktor
	DoktorID(&doktor, RandevuBilgi.DoktorIsim, RandevuBilgi.DoktorSoyisim)

	doktorMüsait := false
	for _, doktor := range *existingDoktorlar {
		if DoktorRandevulari(&doktor, RandevuBilgi.Tarih) {
			doktorMüsait = true
			break
		}
	}

	if doktorMüsait {
		randevuResponse := model.Randevu{
			Tarih:            RandevuBilgi.Tarih,
			HastaIsim:        RandevuBilgi.HastaIsim,
			HastaSoyisim:     RandevuBilgi.HastaSoyisim,
			DoktorIsim:       RandevuBilgi.DoktorIsim,
			DoktorSoyisim:    RandevuBilgi.DoktorSoyisim,
			HastaRahatsizlik: RandevuBilgi.HastaRahatsizlik,
		}
		if err := database.Conn.Create(&randevuResponse).Error; err != nil {
			return c.Status(400).JSON(fiber.Map{"randevu_error": "Randevu oluşturulamadi."})
		}

		doktorRandevuResponse := model.DoktorRandevu{
			DoktorID:            doktor.ID,
			RandevuTarihi:       RandevuBilgi.Tarih,
			RandevuHastaIsmi:    RandevuBilgi.HastaIsim,
			RandevuHastaSoyismi: RandevuBilgi.HastaSoyisim,
		}
		if err := database.Conn.Create(&doktorRandevuResponse).Error; err != nil {
			return c.Status(400).JSON(fiber.Map{"randevu_error": "Randevu oluşturulamadi."})
		}

		return c.Status(200).JSON(fiber.Map{"success": "Randevu başariyla oluşturuldu."})
	} else {
		return c.Status(400).JSON(fiber.Map{"error": "Doktor müsait değil."})
	}
}

func DoktorKontrol(DoktorIsim string, DoktorSoyisim string) (*[]model.Doktor, error) {
	var existingDoktors []model.Doktor
	if err := database.Conn.Where("isim = ? AND soyisim = ?", DoktorIsim, DoktorSoyisim).Find(&existingDoktors).Error; err != nil {
		return nil, err
	}
	return &existingDoktors, nil
}

func DoktorRandevulari(doktor *model.Doktor, tarih string) bool {
	müsaitlik := true

	if err := database.Conn.Preload("Randevular").Find(&doktor).Error; err != nil {
		fmt.Println("Doktor randevulari alinamadi:", err)
		return false
	}

	for _, randevu := range doktor.Randevular {
		if randevu.RandevuTarihi == tarih {
			müsaitlik = false
			break
		}
	}

	return müsaitlik
}

func DoktorID(doktor *model.Doktor, doktorIsim string, doktorSoyisim string) int {
	if err := database.Conn.
		Where("isim = ? AND soyisim = ?", doktorIsim, doktorSoyisim).
		First(doktor).Error; err != nil {
		fmt.Println("Doktor randevuları alınamadı:", err)
		return 0
	}
	return int(doktor.ID)
}

func HastaListesi(c *fiber.Ctx) error {
	var randevular []model.Randevu
	database.Conn.Find(&randevular)
	return c.JSON(randevular)
}

func DoktorListesi(c *fiber.Ctx) error {
	var doktorlar []model.Doktor
	database.Conn.Find(&doktorlar)
	return c.JSON(&doktorlar)
}

func HastaRandevuListesi(c *fiber.Ctx) error {
	var RandevuKontrol struct {
		HastaIsim    string `json:"hasta_isim"`
		HastaSoyisim string `json:"hasta_soyisim"`
	}

	if err := c.BodyParser(&RandevuKontrol); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request data: " + err.Error()})
	}

	var Randevular []model.Randevu
	if err := database.Conn.Where("hasta_isim = ? AND hasta_soyisim = ?", RandevuKontrol.HastaIsim, RandevuKontrol.HastaSoyisim).
		Find(&Randevular).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Database error: " + err.Error()})
	}

	var response []fiber.Map
	for _, randevu := range Randevular {
		item := fiber.Map{
			"hasta_isim":     RandevuKontrol.HastaIsim,
			"hasta_soyisim":  RandevuKontrol.HastaSoyisim,
			"tarih":          randevu.Tarih,
			"randevu_bolum":  randevu.RandevuBölüm,
			"doktor_isim":    randevu.DoktorIsim,
			"doktor_soyisim": randevu.DoktorSoyisim,
		}
		response = append(response, item)
	}

	return c.JSON(response)
}

func DoktorRandevuListesi(c *fiber.Ctx) error {
	var Doktor []model.Doktor
	if err := database.Conn.Preload("Randevular").Find(&Doktor).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Database error: " + err.Error()})
	}

	var response []fiber.Map
	for _, doktor := range Doktor {
		for _, randevu := range doktor.Randevular {
			item := fiber.Map{
				"doktor_isim":    doktor.Isim,
				"doktor_soyisim": doktor.Soyisim,
				"randevular": fiber.Map{
					"hasta_ismi":    randevu.RandevuHastaIsmi,
					"hasta_soyismi": randevu.RandevuHastaSoyismi,
					"tarih":         randevu.RandevuTarihi,
				},
			}
			response = append(response, item)
		}
	}

	return c.JSON(response)
}

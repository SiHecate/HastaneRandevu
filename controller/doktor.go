package controller

import (
	model "hastane-uyg/Model"

	"hastane-uyg/core/database"

	"github.com/gofiber/fiber/v2"
)

func DoktorEkle(c *fiber.Ctx) error {
	var DoktorBilgi struct {
		Isim     string `json:"doktor_isim"`
		Soyisim  string `json:"doktor_soyisim"`
		Hastane  string `json:"doktor_hastane"`
		Uzmanlik string `json:"doktor_uzmanlık"`
	}

	if err := c.BodyParser(&DoktorBilgi); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request data: " + err.Error()})
	}

	var existingDoktor model.Doktor
	if err := database.Conn.Where("isim = ? AND soyisim = ?", DoktorBilgi.Isim, DoktorBilgi.Soyisim).First(&existingDoktor).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{"doktor_error": "Doktor mevcut"})
	}

	response := model.Doktor{
		Isim:     DoktorBilgi.Isim,
		Soyisim:  DoktorBilgi.Soyisim,
		Hastane:  DoktorBilgi.Hastane,
		Uzmanlik: DoktorBilgi.Uzmanlik,
	}

	if err := database.Conn.Create(response).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"doktor_error": "Doktor eklemede bir problem yaşandi"})
	}

	return c.Status(200).JSON(fiber.Map{"message": "Doktor başariyla eklendi"})
}

func DoktorListe(c *fiber.Ctx) error {
	var existingDoktors []model.Doktor
	if err := database.Conn.Find(&existingDoktors).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"doktor_error": "Doktor listeleme işleminde bir hata oluştu"})
	}

	return c.Status(200).JSON(fiber.Map{"message": "Listedeki bütün doktorlar", "data": existingDoktors})
}

func DoktorSil(c *fiber.Ctx) error {
	var DoktorBilgi struct {
		Isim     string `json:"doktor_isim"`
		Soyisim  string `json:"doktor_soyisim"`
		Hastane  string `json:"doktor_hastane"`
		Uzmanlik string `json:"doktor_uzmanlık"`
	}

	if err := c.BodyParser(&DoktorBilgi); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request data: " + err.Error()})
	}

	var existingDoktor model.Doktor
	if err := database.Conn.Where("isim = ? AND soyisim = ?", DoktorBilgi.Isim, DoktorBilgi.Soyisim).First(&existingDoktor).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{"doktor_error": "Doktor mevcut değil, silme işlemi gerçekleştirilemez."})
	}

	if err := database.Conn.Delete(&existingDoktor).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"doktor_error": "Doktor silme işleminde bir problem yaşandi"})
	}

	return c.Status(200).JSON(fiber.Map{"message": "Doktor başariyla silindi"})
}

func DoktorGüncelle(c *fiber.Ctx) error {
	var DoktorBilgi struct {
		Isim     string `json:"doktor_isim"`
		Soyisim  string `json:"doktor_soyisim"`
		Hastane  string `json:"doktor_hastane"`
		Uzmanlik string `json:"doktor_uzmanlık"`
	}

	if err := c.BodyParser(&DoktorBilgi); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request data: " + err.Error()})
	}

	var existingDoktor model.Doktor
	if err := database.Conn.Where("isim = ? AND soyisim = ?", DoktorBilgi.Isim, DoktorBilgi.Soyisim).First(&existingDoktor).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{"doktor_error": "Doktor mevcut değil, güncelleme işlemi gerçekleştirilemez."})
	}

	// Güncellenecek alanları belirle
	updateFields := map[string]interface{}{
		"Isim":     DoktorBilgi.Isim,
		"Soyisim":  DoktorBilgi.Soyisim,
		"Hastane":  DoktorBilgi.Hastane,
		"Uzmanlik": DoktorBilgi.Uzmanlik,
	}

	if err := database.Conn.Model(&existingDoktor).Updates(updateFields).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"doktor_error": "Doktor güncelleme işleminde bir hata oluştu"})
	}

	return c.Status(200).JSON(fiber.Map{"message": "Doktor başarıyla güncellendi"})
}

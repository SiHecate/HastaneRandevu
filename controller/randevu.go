package controller

import (
	"fmt"

	model "hastane-uyg/Model"

	"hastane-uyg/core/database"

	"github.com/gofiber/fiber/v2"
)

/*
	ToDo:
		İlk olarak randevu alma sistemi geliştirilecek:
			- Randevuyu alacak hastanın adı ve soyadı alınacak
			- Randevu alınacak hastanenin adı alınacak
			- Randevu alınacak doktorun ismi alınacak ve girilen hastanenin içerisinde bu doktorun olup olmadığı kontrol edilecek ve buna göre listelenecek doktorlar
			- Randevu alınan doktor'un dolu olduğu saatler ayarlanacak. Bu ayarlamayı SQL üzerinden halletmeyi planıyorum
				-- Bu saat kontrolünü galiba her doktorun altına bir tablo daha açarak müsait olduğu saatleri tutarak gidereceğim.
			- Eğer randevu alma işleminde herhangi bir sıkıntı olmazsa her doktorun altında bu randevu görülecek ve işlemleri olacak
			- Hastanın adına soy adına göre de bir fonksiyon yardımıyla randevularını gösterilecek
			- Randevu iptal (delete), düzenleme(update) işlemleri eklenecek
				-- Bu sayede CRUD işlemlerinin hepsi hallolucak
					--- Create = Randevu oluşturma
					--- Read   = Randevu listeleme
					--- Update = Randevu düzenleme
					--- Delete = Randevu iptali
		ANA BAŞLIKLAR BU KADAR.
*/

func RandevuOluştur(c *fiber.Ctx) error {
	var RandevuBilgi struct {
		Tarih            int    `json:"tarih"`
		HastaIsim        string `json:"hasta_isim"`
		HastaSoyisim     string `json:"hasta_soyisim"`
		DoktorIsim       string `json:"doktor_isim"`
		DoktorSoyisim    string `json:"doktor_soyisim"`
		RandevuBölüm     string `json:"randevu_bölüm"`
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
			RandevuBölüm:     RandevuBilgi.RandevuBölüm,
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

func DoktorRandevulari(doktor *model.Doktor, tarih int) bool {
	var müsaitlik bool
	if err := database.Conn.Preload("Randevular").Find(&doktor).Error; err != nil {
		fmt.Println("Doktor randevulari alinamadi:", err)
		return false
	}
	for _, randevu := range doktor.Randevular {
		if randevu.RandevuTarihi == tarih {
			müsaitlik = false
			break
		} else {
			müsaitlik = true
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

/*
	Randevu oluşturma tamamlandı şu an hem randevu tablosuna hemde doktorlara ait olan randevu tablosuna verileri gönderiliyor
	ToDo:
		Bütün hastaları lisleyecek endpoint BİTTİ
		Bütün doktorları lisleyecek endpoint
		Hastaların sahip olduğu randevuları listleyecek endpoint
		Doktorların sahip olduğu randevuları listleyecek endpoint
*/

// Admin paneli kısmında gözükecek olan kısım Get methodu
func HastaListesi(c *fiber.Ctx) error {
	var randevular []model.Randevu

	hastaIsim := c.Query("hasta_isim")
	hastaSoyisim := c.Query("hasta_soyisim")
	randevuBolum := c.Query("randevu_bolum")

	database.Conn.Where("hasta_isim = ? AND hasta_soyisim = ? AND randevu_bolum = ?",
		hastaIsim, hastaSoyisim, randevuBolum).
		Find(&randevular)

	return c.JSON(randevular)
}

// Admin paneli kısmında gözükecek olan kısım Get methodu
func DoktorListesi(c *fiber.Ctx) error {
	var doktorlar []model.Doktor

	doktorIsim := c.Query("isim")
	doktorSoyisim := c.Query("soyisim")
	doktorUzmanlik := c.Query("uzmanlik")
	doktorHastane := c.Query("hastane")

	database.Conn.Where("isim = ? AND soyisim = ? AND uzmanlik = ? AND hastane = ?",
		doktorIsim, doktorSoyisim, doktorUzmanlik, doktorHastane).
		Find(&doktorlar)

	return c.JSON(&doktorlar)
}

func HastaRandevuListesi(c *fiber.Ctx) error {
	var RandevuKontrol struct {
		HastaIsim    string `json:"hasta_isim"`
		HastaSoyisim string `json"hasta_soyisim"`
	}
}

func DoktorRandevuListesi(c *fiber.Ctx) error {
	var RandevuKontrol struct {
		DoktorIsim    string `json:"doktor_isim"`
		DoktorSoyisim string `json:"doktor_soyisim"`
	}
}

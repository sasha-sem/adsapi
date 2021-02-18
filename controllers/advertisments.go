package controllers

import (
	"sashasem/adsapi/models"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

//PostAd creates an ad
func PostAd(c *gin.Context) {
	db := models.InitDb()

	var ad models.Advertisement
	c.Bind(&ad)
	if ad.Name != "" && len(ad.Name) <= 200 && ad.Description != "" && len(ad.Description) <= 1000 && ad.Price > 0 && ad.Pictures != "" && len(strings.Split(ad.Pictures, ",")) <= 3 {
		db.Create(&ad)
		c.JSON(201, gin.H{"id": ad.ID})
	} else {
		c.JSON(422, gin.H{"error": "Fields are incorrect"})
	}
}

//GetAds gets all ads
func GetAds(c *gin.Context) {
	db := models.InitDb()
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	orderBy := c.DefaultQuery("order_by", "time_asc")
	if err != nil || page < 1 {
		c.JSON(422, gin.H{"error": "page is incorrect"})
		return
	}
	var ads []models.AdShort

	switch orderBy {
	case "time_asc":
		db.Table("advertisements").Select("id, name, (CASE WHEN INSTR(pictures,',') = 0 THEN pictures ELSE substr(pictures, 0, INSTR(pictures,',')) END) as main_picture, price, created_at").Order("created_at asc").Limit(10).Offset(10 * (page - 1)).Find(&ads)
	case "time_desc":
		db.Table("advertisements").Select("id, name, (CASE WHEN INSTR(pictures,',') = 0 THEN pictures ELSE substr(pictures, 0, INSTR(pictures,',')) END) as main_picture, price, created_at").Order("created_at desc").Limit(10).Offset(10 * (page - 1)).Find(&ads)
	case "price_asc":
		db.Table("advertisements").Select("id, name, (CASE WHEN INSTR(pictures,',') = 0 THEN pictures ELSE substr(pictures, 0, INSTR(pictures,',')) END) as main_picture, price, created_at").Order("price asc").Limit(10).Offset(10 * (page - 1)).Find(&ads)
	case "price_desc":
		db.Table("advertisements").Select("id, name, (CASE WHEN INSTR(pictures,',') = 0 THEN pictures ELSE substr(pictures, 0, INSTR(pictures,',')) END) as main_picture, price, created_at").Order("price desc").Limit(10).Offset(10 * (page - 1)).Find(&ads)
	default:
		c.JSON(422, gin.H{"error": "order_by is incorrect"})
		return
	}
	c.JSON(200, ads)

}

//GetAd gets single ad by id
func GetAd(c *gin.Context) {
	db := models.InitDb()

	id := c.Params.ByName("id")
	unitedFields := c.Query("fields")
	fields := strings.Split(unitedFields, ",")

	if len(fields) == 1 && fields[0] == "" {
		var ad models.AdShort
		db.Table("advertisements").Select("id, name, (CASE WHEN INSTR(pictures,',') = 0 THEN pictures ELSE substr(pictures, 0, INSTR(pictures,',')) END) as main_picture, price, created_at").First(&ad, id)
		if ad.ID != 0 {
			c.JSON(200, ad)
		} else {
			c.JSON(404, gin.H{"error": "ad not found"})
		}
	} else if len(fields) == 1 && fields[0] == "description" {
		var ad models.AdShortDescr
		db.Table("advertisements").Select("id, name, (CASE WHEN INSTR(pictures,',') = 0 THEN pictures ELSE substr(pictures, 0, INSTR(pictures,',')) END) as main_picture, price, description").First(&ad, id)
		if ad.ID != 0 {
			c.JSON(200, ad)
		} else {
			c.JSON(404, gin.H{"error": "ad not found"})
		}
	} else if len(fields) == 1 && fields[0] == "pictures" {
		var ad models.AdShortPictures
		db.Table("advertisements").Select("id, name, pictures, price").First(&ad, id)
		if ad.ID != 0 {
			c.JSON(200, ad)
		} else {
			c.JSON(404, gin.H{"error": "ad not found"})
		}
	} else if len(fields) == 2 && ((fields[0] == "pictures" && fields[1] == "description") || (fields[0] == "description" && fields[1] == "pictures")) {
		var ad models.Advertisement
		db.First(&ad, id)
		if ad.ID != 0 {
			c.JSON(200, ad)
		} else {
			c.JSON(404, gin.H{"error": "ad not found"})
		}
	} else {
		c.JSON(422, gin.H{"error": "fields are incorrect"})
	}

}

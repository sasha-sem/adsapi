package models

import (
	"time"
)

//Advertisement is main structure of ad
type Advertisement struct {
	ID          int       `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	Name        string    `gorm:"type:varchar(200);not null" form:"name" json:"name"`
	Description string    `gorm:"type:text;size:1000;not null" form:"description" json:"description"`
	Price       int       `gorm:"not null" form:"price" json:"price"`
	CreatedAt   time.Time `gorm:"not null" form:"created_at" json:"created_at"`
	Pictures    string    `gorm:"type:text" form:"pictures" json:"pictures"`
}

//AdShort is structure for short info about ads
type AdShort struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Price       int       `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	MainPicture string    `json:"main_picture"`
}

//AdShortDescr is structure for short info about ads with description
type AdShortDescr struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	MainPicture string `json:"main_picture"`
}

//AdShortPictures is structure for short info about ads with pictures
type AdShortPictures struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Pictures string `json:"pictures"`
}

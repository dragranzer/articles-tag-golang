package request

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Status  bool
	Title   string `gorm:"column:article_title"`
	Content string `gorm:"column:raw_content"`
}

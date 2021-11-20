package migrate

import (
	"my-module/config"
	_article_data "my-module/features/articles/data"
)

func AutoMigrate() {
	if err := config.DB.Exec("DROP TABLE IF EXISTS tags").Error; err != nil {
		panic(err)
	}

	if err := config.DB.Exec("DROP TABLE IF EXISTS articles").Error; err != nil {
		panic(err)
	}

	config.DB.AutoMigrate(&_article_data.Article{}, &_article_data.Tag{})

	article1 := _article_data.Article{
		Title:   "title1",
		Content: "article keren",
		Tags: []_article_data.Tag{
			{Name: "anime"},
			{Name: "spongebob"},
		},
	}

	article2 := _article_data.Article{
		Title:   "title2",
		Content: "article keren2",
		Tags: []_article_data.Tag{
			{Name: "anime"},
		},
	}

	// setelah dibuat, insert
	if err := config.DB.Create(&article1).Error; err != nil {
		panic(err)
	}
	if err := config.DB.Create(&article2).Error; err != nil {
		panic(err)
	}

}

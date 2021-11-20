package data

import (
	"fmt"
	"my-module/features/articles"

	"gorm.io/gorm"
)

// article model
type Article struct {
	gorm.Model

	Status  bool
	Title   string `gorm:"column:article_title"`
	Content string `gorm:"column:raw_content"`
	Tags    []Tag  `json:"tag" form:"tag" gorm:"foreigKey:ArticleID;references:ID"`
}

// tag model
type Tag struct {
	ID        int
	Name      string
	ArticleID int
}

//DTO

func toCoreRecord(a articles.Core) Article {
	fmt.Println("a =======", a)
	tags := []Tag{}

	for _, tag := range a.Tags {
		tags = append(tags, Tag{ID: tag.ID, Name: tag.Name})
	}

	return Article{
		Title:   a.Title,
		Status:  a.Status,
		Content: a.Content,
		Tags:    tags,
	}
}

func (a *Article) toCore() articles.Core {
	fmt.Println("Masuk to Core")
	return articles.Core{
		ID:        int(a.ID),
		Title:     a.Title,
		Status:    a.Status,
		Content:   a.Content,
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
	}
}

func toCoreList(resp []Article) []articles.Core {
	fmt.Println("Masuk to Core List")
	a := []articles.Core{}
	for _, article := range resp {
		a = append(a, article.toCore())
	}
	return a
}

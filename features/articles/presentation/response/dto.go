package response

import (
	"fmt"
	"time"

	"my-module/features/articles"
)

type Article struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Status    bool      `json:"status"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Tags      []TagCore `json:"tag" form:"tag"`
}

type TagCore struct {
	ID   int
	Name string
}

func FromCore(core articles.Core) Article {
	fmt.Println("Masuk DTO from core")
	return Article{
		ID:        core.ID,
		CreatedAt: core.CreatedAt,
		UpdatedAt: core.UpdatedAt,
		Status:    core.Status,
		Title:     core.Title,
		Content:   core.Content,
	}
}

func FromCoreSlice(core []articles.Core) []Article {
	fmt.Println("Masuk DTO from core slice")
	var artArray []Article
	for key := range core {
		artArray = append(artArray, FromCore(core[key]))
	}
	return artArray
}

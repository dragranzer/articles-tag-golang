package data

import (
	"fmt"
	"time"

	"my-module/features/tags"
)

type Tag struct {
	ID        int       `json:"id" form:"id"`
	Name      string    `json:"name" form:"name"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
}

//DTO

func toCoreRecord(t tags.Core) Tag {
	fmt.Println("t =======", t)
	return Tag{
		ID:   t.ID,
		Name: t.Name,
	}
}

func (t *Tag) toCore() tags.Core {
	fmt.Println("Masuk to Core")
	return tags.Core{
		ID:   int(t.ID),
		Name: t.Name,
	}
}

func toCoreList(resp []Tag) []tags.Core {
	fmt.Println("Masuk to Core List")
	a := []tags.Core{}
	for _, tag := range resp {
		a = append(a, tag.toCore())
	}
	return a
}

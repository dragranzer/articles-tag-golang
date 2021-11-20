package response

import (
	"fmt"
	"my-module/features/tags"
	"time"
)

type Tag struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
}

func FromCore(core tags.Core) Tag {
	fmt.Println("Masuk DTO from core")
	return Tag{
		ID:        core.ID,
		CreatedAt: core.CreatedAt,
		UpdatedAt: core.UpdatedAt,
		Name:      core.Name,
	}
}

func FromCoreSlice(core []tags.Core) []Tag {
	fmt.Println("Masuk DTO from core slice")
	var artArray []Tag
	for key := range core {
		artArray = append(artArray, FromCore(core[key]))
	}
	return artArray
}

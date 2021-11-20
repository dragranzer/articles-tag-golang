package tags

import "time"

type Core struct {
	ID        int       `json:"id" form:"id"`
	Name      string    `json:"name" form:"name"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
}

type Bussiness interface {
	CreateData(data Core) (err error)
	GetAllData(search string) (resp []Core)
	// continue the function abtraction
}

type Data interface {
	InsertData(data Core) (err error)
	SelectData(title string) (resp []Core)
}

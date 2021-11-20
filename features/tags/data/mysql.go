package data

import (
	"fmt"
	"my-module/features/tags"

	"gorm.io/gorm"
)

type mysqlTagRepository struct {
	Conn *gorm.DB
}

func NewTagRepository(conn *gorm.DB) tags.Data {
	// fmt.Println("Masuk New Tags Repo")
	return &mysqlTagRepository{
		Conn: conn,
	}
}

func (ar *mysqlTagRepository) InsertData(data tags.Core) error {
	// article := data{}
	fmt.Println("Data ========", data)
	recordData := toCoreRecord(data)
	fmt.Println("Record data ======== ", recordData)
	err := ar.Conn.Create(&recordData)
	if err != nil {
		return err.Error
	}
	return nil
}

func (ar *mysqlTagRepository) SelectData(title string) (resp []tags.Core) {
	fmt.Println("Masuk Select Data")
	record := []Tag{}
	if err := ar.Conn.Find(&record).Error; err != nil {
		return []tags.Core{}
	}
	return toCoreList(record)
}

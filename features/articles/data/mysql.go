package data

import (
	"fmt"
	"my-module/features/articles"

	"gorm.io/gorm"
)

type mysqlArticleRepository struct {
	Conn *gorm.DB
}

func NewArticleRepository(conn *gorm.DB) articles.Data {
	// fmt.Println("Masuk New Articles Repo")
	return &mysqlArticleRepository{
		Conn: conn,
	}
}

func (ar *mysqlArticleRepository) InsertData(data articles.Core) error {
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

func (ar *mysqlArticleRepository) SelectData(title string) (resp []articles.Core) {
	fmt.Println("Masuk Select Data")
	record := []Article{}

	if err := ar.Conn.Preload("Tags").Find(&record).Error; err != nil {
		return []articles.Core{}
	}

	// jangan lupa ditranlasiin ke core
	return toCoreList(record)
}

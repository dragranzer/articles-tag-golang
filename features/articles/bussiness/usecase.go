package bussiness

import (
	"fmt"
	"my-module/features/articles"
)

type articlesUsecase struct {
	articleData articles.Data
}

func NewAricleBussiness(artData articles.Data) articles.Bussiness {
	fmt.Println("Usecase Bussiness")
	fmt.Println(artData)
	return &articlesUsecase{
		articleData: artData,
	}
}

func (au *articlesUsecase) CreateData(data articles.Core) error {
	// if err := au.validate.Struct(data); err != nil {
	// 	return articles.Core{}, err
	// }
	fmt.Println("bussiness", data)
	err := au.articleData.InsertData(data)
	if err != nil {
		return err
	}
	return nil
}

func (au *articlesUsecase) GetAllData(search string) (resp []articles.Core) {
	resp = au.articleData.SelectData(search)
	return
}

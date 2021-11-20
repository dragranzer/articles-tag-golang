package bussiness

import (
	"fmt"
	"my-module/features/tags"
)

type tagsUsecase struct {
	tagData tags.Data
}

func NewAricleBussiness(artData tags.Data) tags.Bussiness {
	fmt.Println("Usecase Bussiness")
	fmt.Println(artData)
	return &tagsUsecase{
		tagData: artData,
	}
}

func (au *tagsUsecase) CreateData(data tags.Core) error {
	// if err := au.validate.Struct(data); err != nil {
	// 	return tags.Core{}, err
	// }
	fmt.Println("bussiness", data)
	err := au.tagData.InsertData(data)
	if err != nil {
		return err
	}
	return nil
}

func (au *tagsUsecase) GetAllData(search string) (resp []tags.Core) {
	resp = au.tagData.SelectData(search)
	return
}

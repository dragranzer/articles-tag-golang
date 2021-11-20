package factory

import (
	"my-module/config"
	_tag_bussiness "my-module/features/tags/bussiness"
	_tag_data "my-module/features/tags/data"
	"my-module/features/tags/presentation"
)

type TagPresenter struct {
	TagPresentation presentation.TagsHandler
}

func InitTag() TagPresenter {

	tagData := _tag_data.NewTagRepository(config.DB)
	tagBussiness := _tag_bussiness.NewAricleBussiness(tagData)

	return TagPresenter{
		TagPresentation: *presentation.NewTagHandler(tagBussiness),
	}
}

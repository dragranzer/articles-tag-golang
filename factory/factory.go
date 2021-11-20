package factory

import (
	"my-module/config"
	_article_bussiness "my-module/features/articles/bussiness"
	_article_data "my-module/features/articles/data"
	"my-module/features/articles/presentation"
)

type Presenter struct {
	ArticlePresentation presentation.ArticlesHandler
}

func Init() Presenter {

	articleData := _article_data.NewArticleRepository(config.DB)
	articleBussiness := _article_bussiness.NewAricleBussiness(articleData)

	return Presenter{
		ArticlePresentation: *presentation.NewArticleHandler(articleBussiness),
	}
}

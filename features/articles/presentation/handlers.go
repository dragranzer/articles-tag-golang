package presentation

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"my-module/features/articles"
	presentation_response "my-module/features/articles/presentation/response"
)

type ArticlesHandler struct {
	articleBussiness articles.Bussiness
}

func NewArticleHandler(ab articles.Bussiness) *ArticlesHandler {
	fmt.Println("Masuk Handlers")
	return &ArticlesHandler{
		articleBussiness: ab,
	}
}

func (ah *ArticlesHandler) GetAllArticle(c echo.Context) error {
	fmt.Println("Masuk Handlers F2")
	result := ah.articleBussiness.GetAllData("")

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "hope all feeling well",
		"data":    presentation_response.FromCoreSlice(result),
	})
}

func (ah *ArticlesHandler) InsertArticle(c echo.Context) error {
	// fmt.Println("Masuk Handlers F2")
	article := articles.Core{}
	c.Bind(&article)
	fmt.Println("article presentation ========== ", article)
	err := ah.articleBussiness.CreateData(article)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"article": presentation_response.FromCore(article),
	})
}

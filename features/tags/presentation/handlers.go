package presentation

import (
	"fmt"
	"my-module/features/tags"
	"net/http"

	"my-module/features/tags/presentation/response"

	"github.com/labstack/echo/v4"
)

type TagsHandler struct {
	tagBussiness tags.Bussiness
}

func NewTagHandler(tb tags.Bussiness) *TagsHandler {
	fmt.Println("Masuk Handlers")
	return &TagsHandler{
		tagBussiness: tb,
	}
}

func (th *TagsHandler) GetAllTag(c echo.Context) error {
	fmt.Println("Masuk Handlers F2")
	result := th.tagBussiness.GetAllData("")

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "hope all feeling well",
		"data":    response.FromCoreSlice(result),
	})
}

func (ah *TagsHandler) InsertTag(c echo.Context) error {
	// fmt.Println("Masuk Handlers F2")
	tag := tags.Core{}
	c.Bind(&tag)
	fmt.Println("tag presentation ========== ", tag)
	err := ah.tagBussiness.CreateData(tag)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"tag":     tag,
	})
}

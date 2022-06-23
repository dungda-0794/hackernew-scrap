package delivery

import (
	"hackernew-scrap/domain/scrap"
	"net/http"

	"github.com/labstack/echo"
)

// NewsHandler  represent the httphandler for article
type NewsHandler struct {
	AUsecase scrap.NewsUsecase
}

// NewNewsHandler will initialize the articles/ resources endpoint
func NewNewsHandler(e *echo.Echo, us scrap.NewsUsecase) {
	handler := &NewsHandler{
		AUsecase: us,
	}
	e.GET("/articles", handler.SendNews)
}

// SendNews will send article to slack
func (a *NewsHandler) SendNews(c echo.Context) error {
	res, err := a.AUsecase.FetchData()
	if err != nil || !res {
		return c.JSON(http.StatusBadRequest, map[string]string{"status": "fail"})
	}
	return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
}

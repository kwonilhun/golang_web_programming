package internal

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Controller struct {
	service Service
}

func NewController(service Service) *Controller {
	return &Controller{service: service}
}

func (controller *Controller) Create(c echo.Context) error {
	var req CreateRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "요청 값이 유효하지 않습니다.")
	}
	res, err := controller.service.Create(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "이미 값이 있습니다.")
	}
	return c.JSON(http.StatusCreated, res)
}

func (controller *Controller) GetByID(c echo.Context) error {
	// 요청을 읽어오고
	id := c.Param("id")
	res, err := controller.service.GetByID(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "요청값이 유효하지 않습니다.")
	}
	return c.JSON(http.StatusOK, res)
}

func (controller *Controller) updateId(c echo.Context) error {
	var req UpdateRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "요청값이 유효하지 않습니다.")
	}
	res, err := controller.service.updateById(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "요청값이 유효하지 않습니다.")
	}
	return c.JSON(http.StatusOK, res)
}

func (controller *Controller) deleteById(c echo.Context) error {
	id := c.Param("id")
	fmt.Println("삭제 id :", id)
	err := controller.service.deleteById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "요청값이 유효하지 않습니다.")
	}
	return c.JSON(http.StatusOK, "delete success")
}

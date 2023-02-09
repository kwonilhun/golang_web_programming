package internal

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

const _defaultPort = 8082

type Server struct {
	controller Controller
}

func NewDefaultServer() *Server {
	data := map[string]Membership{}
	service := NewService(*NewRepository(data))
	controller := NewController(*service)
	return &Server{
		controller: *controller,
	}
}

func (s *Server) Run() {
	e := echo.New()

	s.Routes(e)
	log.Fatal(e.Start(fmt.Sprintf(":%d", _defaultPort)))
}

func (s *Server) Routes(e *echo.Echo) {
	g := e.Group("/v1")
	RouteMemberships(g, s.controller)
}

func RouteMemberships(e *echo.Group, c Controller) {
	//e.GET("/memberships", c.GetByID, settingMiddleWare)
	e.GET("/memberships/:id", c.GetByID, middleware.RequestIDWithConfig(middleware.RequestIDConfig{
		TargetHeader: "X-My-Request-ID",
	}))
	e.POST("/memberships", c.Create, settingMiddleWare)
	e.DELETE("/memberships/:id", c.deleteById, settingMiddleWare)
	e.PATCH("/memberships", c.updateId, settingMiddleWare)
}

func settingMiddleWare(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Request().Header.Set("X-My-Request-Header", uuid.New().String())
		fmt.Println("request URL : ", c.Request().URL)
		fmt.Println("request http Method : ", c.Request().Method)

		fmt.Println("request body : ", c.Request().Body)
		fmt.Println("http response status code : ", c.Response().Status)
		//fmt.Println("http response body : ", c.Response().)
		return next(c)
	}
}

func bodyString() {

}

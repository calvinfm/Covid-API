package routes

import (
	"Covid-GO-API/controller"
	"github.com/labstack/echo"
)

type Routing struct {
	routes controller.Controller
}

type RoutingInterface interface {
	GetRoutes() *echo.Echo
}

func (Routing Routing) GetRoutes() *echo.Echo {
	e := echo.New()
	e.POST("/covid/", Routing.routes.PostsCovidDB)
	e.GET("/covid/", Routing.routes.GetCovidDB)
	e.GET("/info/covid/", Routing.routes.GetCovidInfoDB)
	e.GET("/covid/:id", Routing.routes.GetCovidIdDB)
	e.PUT("/covid/:id", Routing.routes.UpdateCovidDB)
	e.DELETE("/covid/:id", Routing.routes.DeleteCovidDb)

	return e
}

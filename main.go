package main

import "Covid-GO-API/routes"

// Starting server
func main() {
	echo := routes.Routing.GetRoutes(routes.Routing{})
	_ = echo.Start(":1103")
}

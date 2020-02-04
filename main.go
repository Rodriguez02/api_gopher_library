package main

import "api_gopher_library/router"

func main() {
	router.MapRoutes()
	router.Run()
}

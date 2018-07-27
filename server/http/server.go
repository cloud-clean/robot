package http

import (
	"net/http"
	"strconv"
)

func Start(port int){
	router := &Router{}
	router.RegRoutes(InitRouter())
	server := http.Server{
		Addr:":"+strconv.Itoa(port),
		Handler:router,
	}
	server.ListenAndServe()
}

func InitRouter()[]*Route{
	var routes []*Route
	routes = append(routes,NewRoute("/api/switch/on","get",PinHandler,[]string{"pos","status"}))

	return routes
}
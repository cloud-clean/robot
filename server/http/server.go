package http

import (
	"net/http"
	"strconv"
	"robot/common/logger"
)

var log = logger.NewLog()
func Start(port int){
	router := &Router{}
	router.RegRoutes(InitRouter())
	server := http.Server{
		Addr:":"+strconv.Itoa(port),
		Handler:router,
	}
	log.Infof("http server listen on %d",port)
	server.ListenAndServe()
}

func InitRouter()[]*Route{
	var routes []*Route
	routes = append(routes,NewRoute("/api/switch","get",PinHandler,[]string{"pos","status"}))
	routes = append(routes,NewRoute("/api/job/add","get",CronJob,[]string{"name","pos","status","time"}))
	routes = append(routes,NewRoute("/api/status","get",GetFromDb,[]string{"key"}))
	routes = append(routes,NewRoute("/api/status","post",SetToDb,[]string{"key","value"}))
	return routes
}

package http

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
	"strings"
	"fmt"
	"errors"
)

type HttpHandler func(params *Params)*WebResult

type Route struct{
	Handler HttpHandler
	Method string
	Path string
	Params []string
}

type Router struct{
	RouteMap map[string]*Route
}

func (router *Router) ServeHTTP(w http.ResponseWriter,r *http.Request){
	log.Infof("receive %s request for %s",r.Method,r.URL.Path)
	if route,ok := router.getRouter(r.URL.Path,r.Method);ok{
		params,err := getParams(r,route.Params)
		if err != nil{
			http.Error(w,err.Error(),400)
		}
		res := route.Handler(params)
		writeResp(w,res)
	}else{
		http.NotFound(w,r)
	}
}

func(router *Router) RegRoutes(routes []*Route){
	if router.RouteMap == nil {
		router.RouteMap = make(map[string]*Route)
	}
	for _,route := range routes{
		method := strings.ToUpper(route.Method)
		router.RouteMap[method+"_"+route.Path] = route
	}
}

func NewRoute(path,method string,handler HttpHandler,params []string)*Route{
	return &Route{Path:path,Method:method,Handler:handler,Params:params}
}


func (router *Router)getRouter(url,method string)(*Route,bool){
	if router == nil {
		return nil, false
	}
	route := router.RouteMap[method+"_"+url]
	if route == nil{
		return nil,false
	}
	return route,true
}

type Params struct {
	data map[string]interface{}
}

func (p *Params) Get(key string)string{
	if v,ok := p.data[key];ok{
		return fmt.Sprintf("%v",v)
	}else{
		return ""
	}
}

func getParams(r *http.Request,keys []string) (*Params,error){
	if keys == nil{
		return nil,nil
	}
	if len(keys) == 0{
		return nil,nil
	}
	parm := &Params{}
	params := make(map[string]interface{})
	parm.data = params
	if r.Method == "GET"{
		r.ParseForm()
		for _,key := range keys{
			value := r.Form.Get(key)
			if value == ""{
				continue
			}
			params[key] = value
		}
		return parm,nil
	}else{
		result,_ := ioutil.ReadAll(r.Body)
		err := json.Unmarshal(result,&params)
		if err != nil{
			return nil,err
		}
		if params == nil{
			return nil,errors.New("params is nil")
		}
		parm.data = params
		return parm,nil
	}
}

func writeResp(w http.ResponseWriter,res *WebResult){
	w.Header().Set("Content-Type","application/json")
	b,_ := json.Marshal(res)
	w.Write(b)
}

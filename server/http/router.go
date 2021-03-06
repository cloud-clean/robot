package http

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
	"strings"
	"fmt"
	"errors"
	"time"
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
	start:=time.Now();
	log.Infof("receive %s request for %s",r.Method,r.URL.Path)
	if route,ok := router.getRouter(r.URL.Path,r.Method);ok{
		params,err := getParams(r,route.Params)
		if err != nil{
			http.Error(w,err.Error(),400)
		}
		if route.Params != nil || len(route.Params) > 0{
			if(params == nil){
				http.Error(w,"params is not format",400)
			}
		}
		res := route.Handler(params)
		writeResp(w,res)
	}else{
		http.NotFound(w,r)
	}
	dur:=time.Now().Sub(start);
	log.Infof("http parse time:%d",dur.Nanoseconds())
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
	switch r.Method{
	case "GET":
		r.ParseForm()
		for _,key := range keys{
			value := r.Form.Get(key)
			if value == ""{
				continue
			}
			params[key] = value
		}
		return parm,nil
	case "POST":
		ct := r.Header.Get("Content-Type")
		switch(ct){
		case "application/x-www-form-urlencoded":
			r.ParseForm()
			for _,k := range keys{
				value := r.PostFormValue(k)

			}
			return
		case "application/json":
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
		break
	case "PUT":
		break
	default:
		break
		
	}
	if r.Method == "GET"{
		
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

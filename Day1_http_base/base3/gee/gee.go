package gee

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(w http.ResponseWriter, req *http.Request)

type Engine struct {
	router map[string]HandlerFunc
}

func New() *Engine {
	return &Engine{make(map[string]HandlerFunc)}
}

func (engine *Engine) addRouter(method string, pattern string, Handler HandlerFunc){
	key := method + "-" + pattern
	engine.router[key] = Handler
}

func (engine *Engine) Get(pattern string, Handler HandlerFunc){
	engine.addRouter("GET", pattern, Handler)
}

func (engine *Engine) Post(pattern string, Handler HandlerFunc){
	engine.addRouter("POST", pattern, Handler)
}

func (engine *Engine) Run(addr string) (err error){
	return http.ListenAndServe(addr, engine)
}
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request){
	key := req.Method + "-" + req.URL.Path
	if handler, ok := engine.router[key]; ok{
		handler(w, req)
	}else{
		fmt.Fprintf(w, "404 Not Found: %s\n", req.URL)
	}
}
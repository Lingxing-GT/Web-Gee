package gee

import (
	"net/http"
)

type HandlerFunc func(c *Context)

type Engine struct {
	router *Router
}

func New() *Engine {
	return &Engine{router: newRouter()}
}

func (engine *Engine) addRouter(method string, pattern string, Handler HandlerFunc){
	engine.router.addRouter(method, pattern, Handler)
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
	c := newContext(w, req)
	engine.router.handle(c)
}
package gee

import (
	"log"
	"net/http"
)

type Router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *Router {
	return &Router{make(map[string]HandlerFunc)}
}

func (r *Router) addRouter(method string, pattern string, Handler HandlerFunc)  {
	key := method + "-" + pattern
	r.handlers[key] = Handler
	log.Printf("Route %4s - %s", method, pattern)
}

func (r *Router) handle(c *Context){
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok{
		handler(c)
	}else{
		c.String(http.StatusNotFound, "404 NOT FOUND %s\n", c.Path)
	}
}
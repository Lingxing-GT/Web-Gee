package gee

import (
	"log"
	"net/http"
)

type HandlerFunc func(c *Context)

type (
	RouterGroup struct {
		prefix     string
		middleWare []HandlerFunc
		engine     *Engine
	}
	Engine struct {
		*RouterGroup
		router *Router
		groups []*RouterGroup
	}
)

// New is the constructor of gee.Engine8848
func New() *Engine {
	engine := &Engine{}
	engine.router = newRouter()
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

// Group is defined to create a new RouterGroup
func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (group *RouterGroup) addRouter(method string, comp string, Handler HandlerFunc) {
	pattern := group.prefix + comp
	group.engine.router.addRouter(method, pattern, Handler)
	log.Printf("Route %4s - %s", method, pattern)
}

func (group *RouterGroup) Get(pattern string, Handler HandlerFunc) {
	group.addRouter("GET", pattern, Handler)
}

func (group *RouterGroup) Post(pattern string, Handler HandlerFunc) {
	group.addRouter("POST", pattern, Handler)
}

func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.handle(c)
}

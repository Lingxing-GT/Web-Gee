package gee

import (
	"net/http"
	"strings"
)

type Router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

func newRouter() *Router {
	return &Router{
		make(map[string]*node),
		make(map[string]HandlerFunc),
	}
}

func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")
	parts := make([]string, 0)
	for _, part := range vs {
		if part != "" {
			parts = append(parts, part)
			if part[0] == '*' {
				break
			}
		}
	}
	return parts
}

func (r *Router) addRouter(method string, pattern string, Handler HandlerFunc) {
	key := method + "-" + pattern
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}
	parts := parsePattern(pattern)
	r.roots[method].insert(pattern, parts, 0)
	r.handlers[key] = Handler
}

func (r *Router) getRoute(method string, path string) (*node, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string)
	root, ok := r.roots[method]
	if !ok {
		return nil, nil
	}

	n := root.search(searchParts, 0)
	if n != nil {
		parts := parsePattern(n.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params
	}
	return nil, nil
}

/*func (r Router) getRoutes(method string) []*node {
	return nil
}*/

func (r *Router) handle(c *Context) {
	n, params := r.getRoute(c.Method, c.Path)
	if n != nil {
		c.Params = params
		key := c.Method + "-" + n.pattern
		if handler, ok := r.handlers[key]; ok {
			handler(c)
		}
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND %s\n", c.Path)
	}
}

package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	// http objects
	Writer http.ResponseWriter
	Req *http.Request
	// request info
	Path string
	Method string
	// response info
	StatusCode int
}

func newContext(w http.ResponseWriter, req *http.Request) *Context{
	return &Context{
		Writer: w,
		Req: req,
		Path: req.URL.Path,
		Method: req.Method,
	}
}

func (c *Context) PostForm(key string) string{
	return c.Req.FormValue(key)
}

func (c *Context) Query(key string) string{
	return c.Req.URL.Query().Get(key)
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

func (c *Context) SetHeader(key string, value string){
	c.Writer.Header().Set(key, value)
}

func (c *Context) String(code int, format string, values ...interface{}){
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}
// Right sequence: SetHeader -> Status -> Write, Hader().Set -> WriteHeader -> Write
// Can't : WriteHeader -> Header().Set

func (c *Context) JSON(code int, obj interface{})(err error){
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err = encoder.Encode(obj); err != nil{
		panic(err)
	}
	return
}

func (c *Context) Data(code int, data []byte){
	c.Status(code)
	c.Writer.Write(data)
}

func (c *Context) HTML(code int, html string){
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}
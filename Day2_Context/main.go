package main

import (
	"net/http"

	"gee"
)

func main(){
	r := gee.New()
	r.Get("/", func(c *gee.Context){
		c.HTML(http.StatusOK, "<h1>Hello Gee!</h1>")
	})

	r.Get("/hello", func(c *gee.Context){
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.Post("/login", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
}


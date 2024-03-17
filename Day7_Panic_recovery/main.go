package main

import (
	"gee"
	"net/http"
)

func main() {
	r := gee.Default()
	r.Use(gee.Logger())

	r.Get("/", func(c *gee.Context) {
		c.String(http.StatusOK, "Hello Lingxing\n")
	})

	// index out of range for testing Recovery()
	r.Get("/panic", func(c *gee.Context) {
		names := []string{"Lingxing"}
		c.String(http.StatusOK, names[100])
	})

	r.Run(":9999")
}

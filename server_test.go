package proxy

import (
	"fmt"
	"html"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestServer(t *testing.T) {
	sc, err := Default("route.json")
	if err != nil {
		t.Error(err)
	}

	r := gin.Default()
	r.Use(func(ctx *gin.Context) {
		ctx.Request.Header.Add("X-Reverse-Poxy", "yes")
	})
	for _, server := range *sc {
		switch server.Method {
		case "GET":
			r.GET(server.Path, server.ReverseProxy())
		}
	}

	go createAPIServer("/api/Users", "122")

	go createAPIServer("/Users", "125")

	go createReverseProxyServer("123")

	r.Run(":1234")
}

func createAPIServer(path, port string) {
	r := gin.Default()
	r.GET(path, func(ctx *gin.Context) {
		fmt.Fprintf(ctx.Writer, "Hello, %q", html.EscapeString(ctx.Request.URL.Path))
		fmt.Printf("port %s=%v\n", port, ctx.Request.Header)
	})
	r.Run(":" + port)
}

func createReverseProxyServer(port string) {
	sc, err := Default("route.json")
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	for _, server := range *sc {
		switch server.Method {
		case "GET":
			r.GET(server.Path, func(ctx *gin.Context) {
				fmt.Printf("port %s=%v\n", port, ctx.Request.Header)
			}, server.ReverseProxy())
		}
	}

	r.Run(":" + port)
}

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
	for _, server := range *sc {
		switch server.Method {
		case "GET":
			r.GET(server.Path, server.ReverseProxy())
		}
	}

	go ServeAPI()

	r.Run(":1234")
}

func ServeAPI() {
	r := gin.Default()
	r.GET("/api/Users", func(ctx *gin.Context) {
		fmt.Fprintf(ctx.Writer, "Hello, %q", html.EscapeString(ctx.Request.URL.Path))
	})
	r.Run(":123")
}

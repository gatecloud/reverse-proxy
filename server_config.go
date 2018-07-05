package proxy

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

// ServerConfig representing an array of Server
type ServerConfig []Server

// Default creates a ServerConfig handler
func Default(fileName string) (*ServerConfig, error) {
	sc := &ServerConfig{}
	if err := sc.loadFile(fileName); err != nil {
		return nil, err
	}
	return sc, nil
}

// LoadFile loads routing and forwarding files
func (sc *ServerConfig) loadFile(fileName string) error {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, sc)
	if err != nil {
		return err
	}
	return nil
}

// RunProxy runs routing and forwarding configuration
func (sc *ServerConfig) RunProxy(r *gin.Engine) {
	for _, server := range *sc {
		switch server.Method {
		case "GET":
			r.GET(server.Path, server.ReverseProxy())
			r.GET(server.Path+"/:id", server.ReverseProxy())
		case "POST":
			r.POST(server.Path, server.ReverseProxy())
		case "PUT":
			r.PUT(server.Path, server.ReverseProxy())
		case "DELETE":
			r.DELETE(server.Path, server.ReverseProxy())
		case "PATCH":
			r.PATCH(server.Path, server.ReverseProxy())
		case "OPTIONS":
			r.OPTIONS(server.Path, server.ReverseProxy())
		case "HEAD":
			r.HEAD(server.Path, server.ReverseProxy())
		}
	}
}

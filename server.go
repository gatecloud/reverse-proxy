package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Method        string
	Path          string
	ProxyScheme   string
	ProxyPass     string
	ProxyPassPath string
	APIVersion    string
	CustomConfigs CustomConfig
}

func ReverseProxy(s Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		suffixPath := strings.Replace(ctx.Request.RequestURI, s.Path, "", 1)
		director := func(req *http.Request) {
			req.Host = s.ProxyPass
			req.URL.Host = s.ProxyPass
			req.URL.Scheme = s.ProxyScheme
			req.URL.Path = s.ProxyPassPath + suffixPath
		}

		proxy := httputil.NewSingleHostReverseProxy(&url.URL{})
		proxy.Director = director
		proxy.ServeHTTP(ctx.Writer, ctx.Request)
	}
}

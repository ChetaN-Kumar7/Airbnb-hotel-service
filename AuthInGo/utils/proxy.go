package utils

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"path"
)

func ProxyToService(targetBaseUrl string,pathPrefix string) http.HandlerFunc{
	target,err := url.Parse(targetBaseUrl)

	if err != nil {
		fmt.Println("error in Parsing the url")
		return nil
	}

	proxy := &httputil.ReverseProxy{
		Rewrite: func(pr *httputil.ProxyRequest) {

		originalPath := pr.In.URL.Path
		strippedPath := strings.TrimPrefix(originalPath,pathPrefix)

		pr.SetURL(target)

		pr.SetXForwarded()

		pr.Out.URL.Path = path.Join(target.Path, strippedPath)

		if userId, ok := pr.In.Context().Value("userId").(string); ok {
			pr.Out.Header.Set("X-User-Id", userId)
		}
	},
	}

	return proxy.ServeHTTP
}
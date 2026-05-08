package utils

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func ReverseProxyToService(targetBaseUrl string) http.HandlerFunc {
	target, err := url.Parse(targetBaseUrl)
	if err != nil {
		panic(err)
	}

	proxy := &httputil.ReverseProxy{}

	proxy.Rewrite = func(pr *httputil.ProxyRequest) {
		pr.SetURL(target)
		pr.Out.Host = target.Host

		pr.Out.Header.Set("X-Forwarded-Host", pr.In.Host)
		pr.Out.Header.Set("X-Forwarded-Proto", pr.In.URL.Scheme)

		// Append client IP
		pr.Out.Header.Set("X-Forwarded-For", pr.In.RemoteAddr)
	}

	return proxy.ServeHTTP
}

package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

const pwdURL = "https://play-with-docker.org"

func main() {
	target, err := url.Parse(pwdURL)
	if err != nil {
		panic(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(target)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/session" {
			q := r.URL.Query()
			q.Set("session_duration", "2h")
			r.URL.RawQuery = q.Encode()
		}
		proxy.ServeHTTP(w, r)
	})

	http.ListenAndServe(":8080", nil)
}

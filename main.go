package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

type ProxyConfig struct {
	Path   string `json:"path"`
	Target string `json:"target"`
}
type Config struct {
	Host  string        `json:"host"`
	Proxy []ProxyConfig `json:"proxy"`
}

func main() {
	cfg := Config{}
	// open file and read config
	cfgFileName := "config.json"
	if len(os.Args) > 1 {
		cfgFileName = os.Args[1]
	}
	cfgFile, err := os.Open(cfgFileName)
	if err != nil {
		panic(err)
	}
	if err := json.NewDecoder(cfgFile).Decode(&cfg); err != nil {
		panic(err)
	}

	router := mux.NewRouter()

	for _, proxy := range cfg.Proxy {
		proxy := proxy
		proxyURL, err := url.Parse(proxy.Target)
		if err != nil {
			panic(err)
		}
		proxyHandler := httputil.NewSingleHostReverseProxy(proxyURL)

		router.PathPrefix(proxy.Path).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("%s %s\n", r.Method, r.URL)
			r.URL.Path = strings.TrimPrefix(r.URL.Path, proxy.Path)
			proxyHandler.ServeHTTP(w, r)
		})

		fmt.Println("Proxying", proxy.Path, "to", proxy.Target)

	}

	fmt.Printf("Starting proxy server on %s\n", cfg.Host)
	if err := http.ListenAndServe(cfg.Host, router); err != nil {
		panic(err)
	}
}

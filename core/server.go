package core

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/morpheuszero/mavis/config"
)

type Server struct {
	config *config.Config
}

func NewServer() *Server {
	var appConfig = config.NewConfig()
	err := appConfig.LoadConfigFile("mavis.json")
	if err != nil {
		panic(err)
	}
	return &Server{
		config: appConfig,
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Mistakes are not shackles that stops a person's progress. They are the fuel that raises the heart. -Mavis Vermillion")
}

func (s *Server) Run() error {
	http.HandleFunc("/mavis", healthHandler)
	err := s.BuildReverseProxyHandlers()
	if err != nil {
		fmt.Println("Error building reverse proxies:", err)
		return err
	}
	if err = http.ListenAndServe(":"+s.config.GetServerPort(), nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
	return nil
}

func (s *Server) BuildReverseProxyHandlers() error {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		host := strings.Split(r.Host, ":")[0]
		fmt.Println(r.URL.String())

		for _, proxy := range s.config.File.ProxyHosts {
			if proxy.Domain == host {
				target := fmt.Sprintf("%s://%s:%s", proxy.Protocol, proxy.Host, proxy.Port)
				url, err := url.Parse(target)
				if err != nil {
					http.Error(w, "Bad Gateway", http.StatusBadGateway)
					return
				}
				proxy := httputil.NewSingleHostReverseProxy(url)
				proxy.ServeHTTP(w, r)
				return
			} else {
				http.Error(w, "Not Found", http.StatusNotFound)
			}
		}
	})

	return nil
}

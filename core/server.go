package core

import (
	"fmt"
	"github.com/morpheuszero/mavis/config"
	"net/http"
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

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World proxy!")
}

func (s *Server) Run() error {
	println("SERVER PORT: " + s.config.GetServerPort())
	http.HandleFunc("/", handler)
	fmt.Println("Starting server on port :" + s.config.GetServerPort())
	if err := http.ListenAndServe(":"+s.config.GetServerPort(), nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
	return nil
}

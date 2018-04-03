package main

import (
	"github.com/jasonrsmith/crafty/http"
	"github.com/jasonrsmith/crafty/memstore"
)

func main() {
	service := memstore.NewCraftyService()
	craftyHandler := http.NewCraftyHandler(service)
	handler := http.Handler{
		CraftyHandler:      craftyHandler,
		HealthCheckHandler: http.NewHealthCheckHandler(),
	}

	server := http.NewServer(&handler)
	server.ListenAndServe()
}

package core

import (
	"github.com/bobyhw39/go-strapper/stringutils"
	"github.com/go-chi/chi/v5"
	"log"
	"net"
	"net/http"
)

func (a *Module) startWebServer() {
	log.Printf("Listen HTTP at :%s", a.options.HttpAddress)
	router := chi.NewRouter()
	a.Router = router
	http.ListenAndServe(a.options.HttpAddress, a.Router)
}

func (a *Module) startGRPCServer() {
	if !stringutils.IsPointerBlank(a.options.GrpcAddress) {
		lis, err := net.Listen("tcp", *a.options.GrpcAddress)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		if err := a.GrpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %s", err)
		}
	}
}

package main

import (
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/api"
	"github.com/micro/micro/v3/service/logger"
	"github.com/m3o/m3o/services/space/handler"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("space"),
		service.Version("latest"),
	)

	// Register handler
	//pb.RegisterSpaceHandler(srv.Server(), handler.NewSpace(srv))

	srv.Server().Handle(
		srv.Server().NewHandler(
			handler.NewSpace(srv),
			api.WithEndpoint(
				&api.Endpoint{
					Name:    "Space.Download",
					Handler: "api",
					Method:  []string{"POST", "GET"},
					Path:    []string{"/space/download"},
				}),
		))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}

package handlers

import (
	"github.com/LittleMikle/gRPC_tinyurl/pkg/service"
	"github.com/LittleMikle/gRPC_tinyurl/proto"
)

type Implementation struct {
	proto.UnimplementedURLserviseServer

	urlService service.URL
}

func NewURL(urlService service.URL) *Implementation {
	return &Implementation{
		urlService: urlService,
	}
}

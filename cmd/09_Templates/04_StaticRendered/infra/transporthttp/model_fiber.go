package httpserve

// serving static with fiber

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Cfg struct {
	StaticAssetsFolder string

	fiber.Config

	ListenPort int
}

type HTTPSServer struct {
	cfg *Cfg

	server *fiber.App
}

func NewHTTPServer(cfg *Cfg) (*HTTPSServer, error) {
	result := HTTPSServer{
		cfg:    cfg,
		server: fiber.New(cfg.Config),
	}

	return &result, nil
}

func (s *HTTPSServer) Start() error {
	s.server.Static("/", s.cfg.StaticAssetsFolder)

	return s.server.Listen(":" + strconv.Itoa(s.cfg.ListenPort))
}

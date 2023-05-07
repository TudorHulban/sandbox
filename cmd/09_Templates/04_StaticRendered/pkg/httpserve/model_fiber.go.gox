package httpserve

// serving static with fiber

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Cfg struct {
	ListenPort         int
	StaticAssetsFolder string

	fiber.Config
}

type HTTPSServer struct {
	cfg Cfg

	server *fiber.App
}

func NewHTTPServer(cfg Cfg) (*HTTPSServer, error) {
	// TODO: validate configuration

	result := HTTPSServer{
		cfg:    cfg,
		server: fiber.New(cfg.Config),
	}

	return &result, nil
}

func (s *HTTPSServer) Start() {
	s.server.Static("/", s.cfg.StaticAssetsFolder)
	s.server.Listen(":" + strconv.Itoa(s.cfg.ListenPort))
}

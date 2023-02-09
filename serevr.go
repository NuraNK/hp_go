package golangproject

import "github.com/gofiber/fiber/v2"

type Server struct {
	httpserver *fiber.App
}

func (s *Server) Run(port string) error {
	return s.httpserver.Listen(":" + port)
}

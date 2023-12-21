package servers

import (
	monitorHandlers "lolyshop/modules/monitor/handlers"

	"github.com/gofiber/fiber/v2"
)

type IModuleFactory interface {
	MonitorModule()
}

type moduleFactory struct {
	router fiber.Router
	server *server
}

func InitModule(r fiber.Router, s *server) IModuleFactory {
	return &moduleFactory{
		router: r,
		server: s,
	}
}

func (m *moduleFactory) MonitorModule() {
	handler := monitorHandlers.MonitorHandler(m.server.cfg)

	m.router.Get("/health-check", handler.HealthCheck)
}

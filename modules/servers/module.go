package servers

import (
	middlewareshandlers "lolyshop/modules/middlewares/middlewaresHandlers"
	middlewaresrepositories "lolyshop/modules/middlewares/middlewaresRepositories"
	middlewaresusecases "lolyshop/modules/middlewares/middlewaresUsecases"
	monitorHandlers "lolyshop/modules/monitor/handlers"

	"github.com/gofiber/fiber/v2"
)

type IModuleFactory interface {
	MonitorModule()
}

type moduleFactory struct {
	router fiber.Router
	server *server
	mid    middlewareshandlers.IMiddlewaresHandler
}

func InitModule(r fiber.Router, s *server, mid middlewareshandlers.IMiddlewaresHandler) IModuleFactory {
	return &moduleFactory{
		router: r,
		server: s,
		mid:    mid,
	}
}

func InitMiddlewares(s *server) middlewareshandlers.IMiddlewaresHandler {
	repository := middlewaresrepositories.MiddlewaresRepository(s.db)
	usecase := middlewaresusecases.MiddlewaresUsecase(repository)
	handler := middlewareshandlers.MiddlewaresHandler(s.cfg, usecase)
	return handler
}

func (m *moduleFactory) MonitorModule() {
	handler := monitorHandlers.MonitorHandler(m.server.cfg)

	m.router.Get("/health-check", handler.HealthCheck)
}

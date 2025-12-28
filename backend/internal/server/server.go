package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/ilyes-rhdi/buildit-Gql/internal/gql"
	restMiddlewares "github.com/ilyes-rhdi/buildit-Gql/internal/middlewares/rest"
	"github.com/ilyes-rhdi/buildit-Gql/internal/router"
	"github.com/ilyes-rhdi/buildit-Gql/pkg/logger"
)

type Server struct {
	PORT string
}

func NewServer(port string) *Server {
	return &Server{PORT: port}
}

func (s *Server) Setup(e *echo.Echo) {
	e.Static("/public", "public")
	router.SetRoutes(e)

	// CORS configuration
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"}, // TODO: change this
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	// logging middleware
	e.Use(restMiddlewares.LoggingMiddleware)
}

func (s *Server) Run() {
	e := echo.New()
	s.Setup(e)
	gql.Execute(e)
	logger.LogInfo().Msg("graphql server running on /graphql")
	logger.LogInfo().Msg(e.Start(s.PORT).Error())
}

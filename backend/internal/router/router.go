package router

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/ilyes-rhdi/buildit-Gql/config"
	"github.com/ilyes-rhdi/buildit-Gql/internal/sse"
	"github.com/ilyes-rhdi/buildit-Gql/pkg/types"
)

var (
	jwtMiddelware echo.MiddlewareFunc
)

func init() {
	//Initialize the middlware
	config.Load()
	jwtMiddelware = echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWT_SECRET),
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(types.Claims)
		},
	})

}
func SetRoutes(e *echo.Echo) {
	notifier := sse.NewNotifier()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Server Working. GraphQL: /graphql")
	})

	e.GET("/notifications", notifier.NotificationHandler)

	v1 := e.Group("/api/v1")
	AuthRoutes(v1)
	profileRoutes(v1)

	// optionnels
	// UploadRoutes(v1)
	// HealthRoutes(v1)
}


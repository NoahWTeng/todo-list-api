package injectors

import (
	"fmt"
	"github.com/NoahWTeng/todo-api-go/config"
	"github.com/go-chi/chi"
	"net/http"
)

func HttpServerInjector(c *config.GlobalConfig, router *chi.Mux) *http.Server {
	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", c.BaseVariables.Port),
		Handler: router,
	}
	return httpServer
}

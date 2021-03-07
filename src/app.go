package app

import (
	"context"
	"github.com/NoahWTeng/todo-api-go/config"
	"github.com/NoahWTeng/todo-api-go/src/app/tasks"
	"github.com/NoahWTeng/todo-api-go/src/app/users"
	"github.com/NoahWTeng/todo-api-go/src/infra/db/mongodb"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type Container struct {
	config     *config.GlobalConfig
	httpServer *http.Server
	router     *chi.Mux
	mongodb    *mongodb.Handler

	// APP
	usersServices    users.Services
	usersControllers users.Controllers

	tasksServices    tasks.Services
	tasksControllers tasks.Controllers
}

type Application interface {
	Init() error
}

func (c *Container) Init() error {
	c.routes()
	c.startGracefulShutdown()
	return nil
}

func (c *Container) startGracefulShutdown() {
	log.Printf("Server started at port with graceful shutdown: %d \n", c.config.BaseVariables.Port)
	connClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		<-sigint

		// We received an interrupt signal, shut down.
		if err := c.httpServer.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("Failed HTTP server Shutdown: %v", err)
		}
		close(connClosed)
	}()

	if err := c.httpServer.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		log.Fatalf("Failed to start http server: %v", err)
	}

	<-connClosed
}

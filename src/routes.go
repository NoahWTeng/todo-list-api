package app

import (
	"github.com/NoahWTeng/todo-api-go/src/infra/middleware"
	"github.com/go-chi/chi"
)

func (c *Container) routes() {
	routes := c.router

	routes.Route("/api/v1", func(r chi.Router) {
		r.Post("/users/login", c.usersControllers.Login)
		r.Route("/users", func(r chi.Router) {
			r.Use(mdw.Authorization)
			r.With(mdw.Pagination).Get("/", c.usersControllers.Search)
			r.Get("/{id}", c.usersControllers.GetById)
			r.Post("/", c.usersControllers.Create)
			r.Put("/{id}", c.usersControllers.UpdateOne)
			r.Delete("/{id}", c.usersControllers.DeleteOne)
		})

		r.Route("/tasks", func(r chi.Router) {
			r.With(mdw.Pagination).Get("/", c.tasksControllers.Search)
			r.Get("/{id}", c.tasksControllers.GetById)
			r.Post("/", c.tasksControllers.Create)
			r.Put("/{id}", c.tasksControllers.UpdateOne)
			r.Delete("/{id}", c.tasksControllers.DeleteOne)
		})

	})

}

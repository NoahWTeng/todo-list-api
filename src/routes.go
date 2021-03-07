package app

import (
	"github.com/NoahWTeng/todo-api-go/src/infra/middleware"
	"github.com/go-chi/chi"
)

func (c *Container) routes() {
	routes := c.router

	routes.Route("/api/v1", func(r chi.Router) {

		r.Route("/users", func(r chi.Router) {
			r.With(mdw.Pagination).Get("/", c.usersControllers.Search)
			r.Get("/{id}", c.usersControllers.GetById)
			r.Post("/", c.usersControllers.Create)
			r.Put("/{id}", c.usersControllers.UpdateOne)
			r.Delete("/{id}", c.usersControllers.DeleteOne)

		})

	})

}

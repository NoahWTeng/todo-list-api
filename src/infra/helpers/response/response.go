package response

import (
	"github.com/go-chi/render"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

func Error (w http.ResponseWriter, r *http.Request, status int, msg string) {

	render.Status(r, status)
	render.JSON(w, r, bson.M{"status": status, "message": msg})
}

func Json (w http.ResponseWriter, r *http.Request, status int, data interface{}){

	render.Status(r, status)
	render.JSON(w, r, data)
}
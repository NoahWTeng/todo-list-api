package users

import (
	"context"
	"github.com/NoahWTeng/todo-api-go/src/infra/db/mongodb"
	"github.com/NoahWTeng/todo-api-go/src/infra/helpers/pagination"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

type Model struct {
	RawID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updateAt" bson:"updateAt"`
	Name      string             `json:"name" bson:"name"`
	Email     string             `json:"email" bson:"email"`
	Password  string             `json:"password" bson:"password"`
}

type Database struct {
	*mongodb.Handler
	Collection string
}

type Repository struct {
	UsersServices Services
}

type Services interface {
	Create(ctx context.Context, user *Model) (*Model, error)
	FindOne(ctx context.Context, id string, email string) (*Model,error)
	FindAll(ctx context.Context) *pagination.Pages
	Update(ctx context.Context, user *Model, id string) (*Model, error)
	Delete(ctx context.Context, id string) (*Model, error)
}

type Controllers interface {
	Create(writer http.ResponseWriter, request *http.Request)
	Search(writer http.ResponseWriter, request *http.Request)
	GetById(writer http.ResponseWriter, request *http.Request)
	UpdateOne(writer http.ResponseWriter, request *http.Request)
	DeleteOne(writer http.ResponseWriter, request *http.Request)
}

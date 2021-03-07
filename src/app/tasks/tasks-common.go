package tasks

import (
	"context"
	"github.com/NoahWTeng/todo-api-go/src/infra/db/mongodb"
	"github.com/NoahWTeng/todo-api-go/src/infra/helpers/pagination"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

type Model struct {
	RawID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updateAt" bson:"updateAt"`
	Title     string             `json:"title" bson:"title"`
	Status    string             `json:"status" bson:"status"`
	Comment   string             `json:"comment" bson:"comment"`
}

type Database struct {
	*mongodb.Handler
	Collection string
}

type Repository struct {
	TasksRepository Services
}

type Services interface {
	FindAll(ctx context.Context) *pagination.Pages
	Create(ctx context.Context, task *Model) (*Model, error)
	FindOne(ctx context.Context, id string) (*Model, error)
	Update(ctx context.Context, task *Model, id string) (*Model, error)
	Delete(ctx context.Context, id string) (int64, error)
}

type Controllers interface {
	Search(writer http.ResponseWriter, request *http.Request)
	Create(writer http.ResponseWriter, request *http.Request)
	GetById(writer http.ResponseWriter, request *http.Request)
	UpdateOne(writer http.ResponseWriter, request *http.Request)
	DeleteOne(writer http.ResponseWriter, request *http.Request)
}

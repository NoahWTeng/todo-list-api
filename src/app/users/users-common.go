package users

import (
	"context"
	"github.com/NoahWTeng/todo-api-go/src/infra/db/mongodb"
	"github.com/NoahWTeng/todo-api-go/src/infra/helpers/pagination"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

type Model struct {
	RawID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time          `json:"updateAt" bson:"updateAt"`
	Name      string             `json:"name" bson:"name"`
	Email     string             `json:"email" bson:"email"`
	Password  string             `json:"password" bson:"password"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.StandardClaims
}

// Create the JWT key used to create the signature
var JwtKey = []byte("my_secret_key")

type Database struct {
	*mongodb.Handler
	Collection string
}

type Repository struct {
	UsersServices Services
}

type Services interface {
	Create(ctx context.Context, user *Model) (*Model, error)
	FindOne(ctx context.Context, id string, email string) (*Model, error)
	FindAll(ctx context.Context) *pagination.Pages
	Update(ctx context.Context, user *Model, id string) (*Model, error)
	Delete(ctx context.Context, id string) (int64, error)
	SignIn(ctx context.Context, login *Login) (string, error)
}

type Controllers interface {
	Create(writer http.ResponseWriter, request *http.Request)
	Search(writer http.ResponseWriter, request *http.Request)
	GetById(writer http.ResponseWriter, request *http.Request)
	UpdateOne(writer http.ResponseWriter, request *http.Request)
	DeleteOne(writer http.ResponseWriter, request *http.Request)
	Login(writer http.ResponseWriter, request *http.Request)
}

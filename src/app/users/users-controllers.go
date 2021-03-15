package users

import (
	"encoding/json"
	"fmt"
	"github.com/NoahWTeng/todo-api-go/src/infra/helpers/errors"
	"github.com/NoahWTeng/todo-api-go/src/infra/helpers/response"
	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"time"
)

func (repo *Repository) Search(writer http.ResponseWriter, request *http.Request) {

	ctx := request.Context()
	user := ctx.Value("user").(*Claims)

	fmt.Println(user)
	results := repo.UsersServices.FindAll(request.Context())

	response.Json(writer, request, http.StatusOK, results)
}

func (repo *Repository) GetById(writer http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "id")

	result, err := repo.UsersServices.FindOne(request.Context(), id, "")

	if err != nil {
		err := errors.BadRequest(err.Error())
		response.Error(writer, request, err.Status, err.Message)
		return
	}

	response.Json(writer, request, http.StatusOK, result)
}

func (repo *Repository) Create(writer http.ResponseWriter, request *http.Request) {

	var user Model

	_ = json.NewDecoder(request.Body).Decode(&user)

	getUser, _ := repo.UsersServices.FindOne(request.Context(), "", user.Email)

	if getUser.Email != "" {
		err := errors.Conflict("")
		response.Error(writer, request, err.Status, fmt.Sprintf("This %s email is already exists!", user.Email))
		return
	}

	newUser := Model{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
	}

	result, err := repo.UsersServices.Create(request.Context(), &newUser)

	if err != nil {
		err := errors.BadRequest(err.Error())

		response.Error(writer, request, err.Status, err.Message)
		return
	}

	response.Json(writer, request, http.StatusCreated, result)
}

func (repo *Repository) UpdateOne(writer http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "id")
	var user Model

	_ = json.NewDecoder(request.Body).Decode(&user)

	if user.Password == "" && user.Email == "" && user.Name == "" {
		err := errors.BadRequest("")
		response.Error(writer, request, err.Status, err.Message)
		return
	}

	result, err := repo.UsersServices.Update(request.Context(), &user, id)

	if err != nil {
		err := errors.BadRequest(err.Error())

		response.Error(writer, request, err.Status, err.Message)
		return
	}

	response.Json(writer, request, http.StatusOK, result)
}

func (repo *Repository) DeleteOne(writer http.ResponseWriter, request *http.Request) {

	id := chi.URLParam(request, "id")

	type Deleted struct {
		ID      string `json:"id"`
		Deleted bool   `json:"deleted"`
	}

	result, err := repo.UsersServices.Delete(request.Context(), id)

	if err != nil {
		err := errors.BadRequest(err.Error())
		response.Error(writer, request, err.Status, err.Message)
		return
	}

	if result == 0 {
		response.Json(writer, request, http.StatusOK, Deleted{ID: id, Deleted: false})
		return
	}

	response.Json(writer, request, http.StatusOK, Deleted{ID: id, Deleted: true})
}

// Login
func (repo Repository) Login(writer http.ResponseWriter, request *http.Request) {
	var login Login

	_ = json.NewDecoder(request.Body).Decode(&login)

	result, err := repo.UsersServices.SignIn(request.Context(), &login)

	if err != nil {
		err := errors.BadRequest(err.Error())
		response.Error(writer, request, err.Status, err.Message)
		return
	}
	response.Json(writer, request, http.StatusOK, bson.M{"token": result})
}

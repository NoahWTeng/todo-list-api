package tasks

import (
	"encoding/json"
	"github.com/NoahWTeng/todo-api-go/src/infra/helpers/errors"
	"github.com/NoahWTeng/todo-api-go/src/infra/helpers/response"
	"github.com/go-chi/chi"
	"net/http"
	"time"
)

func (repo *Repository) Search(writer http.ResponseWriter, request *http.Request)  {

	results := repo.TasksRepository.FindAll(request.Context())

	response.Json(writer, request, http.StatusOK, results)
}

func (repo *Repository) Create(writer http.ResponseWriter, request *http.Request)  {
	var task Model

	_ = json.NewDecoder(request.Body).Decode(&task)

	newTask := Model{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Title:      task.Title,
		Status:     task.Status,
		Comment:  task.Comment,
	}

	if newTask.Status == "" {
		newTask.Status = "pending"
	}

	result, err := repo.TasksRepository.Create(request.Context(), &newTask)

	if err != nil {
		err := errors.BadRequest(err.Error())

		response.Error(writer, request,err.Status,  err.Message)
		return
	}

	response.Json(writer, request, http.StatusCreated, result)
}

func (repo *Repository) GetById(writer http.ResponseWriter, request *http.Request)  {
	id := chi.URLParam(request, "id")

	result, err := repo.TasksRepository.FindOne(request.Context(), id)

	if err != nil {
		err := errors.BadRequest(err.Error())
		response.Error(writer, request,err.Status,  err.Message)
		return
	}

	response.Json(writer, request, http.StatusOK, result)
}

func (repo *Repository) UpdateOne(writer http.ResponseWriter, request *http.Request)  {
	id := chi.URLParam(request, "id")
	var task Model

	_ = json.NewDecoder(request.Body).Decode(&task)

	result , err:= repo.TasksRepository.Update(request.Context(),&task, id)

	if err != nil {
		err := errors.BadRequest(err.Error())

		response.Error(writer, request,err.Status,  err.Message)
		return
	}

	response.Json(writer, request, http.StatusOK, result)
}

func (repo *Repository) DeleteOne(writer http.ResponseWriter, request *http.Request)  {

	id := chi.URLParam(request, "id")

	type Deleted struct {
		ID  string `json:"id"`
		Deleted bool `json:"deleted"`
	}

	result , err:= repo.TasksRepository.Delete(request.Context(), id)

	if err != nil {
		err := errors.BadRequest(err.Error())
		response.Error(writer, request,err.Status,  err.Message)
		return
	}

	if result == 0 {
		response.Json(writer, request, http.StatusOK, Deleted{ID: id, Deleted: false})
		return
	}

	response.Json(writer, request, http.StatusOK, Deleted{ID: id, Deleted: true})
}
package handler

import (
	"encoding/json"
	"go_todo_app/entity"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type AddTask struct {
	Service   AddTaskService
	Validator *validator.Validate

	// DB   *sqlx.DB
	// Repo *store.Repository
}

func (at *AddTask) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var b struct {
		Title string `json:"title" validate:"required"`
	}

	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}

	err := at.Validator.Struct(b)
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusBadRequest)
		return
	}

	// t := &entity.Task{
	// 	Title:  b.Title,
	// 	Status: entity.TaskStatusTodo,
	// }
	// err = at.Repo.AddTask(ctx, at.DB, t)
	// if err != nil {
	// 	RespondJSON(ctx, w, &ErrResponse{
	// 		Message: err.Error(),
	// 	}, http.StatusInternalServerError)
	// 	return
	// }

	// rsp := struct {
	// 	ID entity.TaskID `json:"id"`
	// }{ID: t.ID}
	// RespondJSON(ctx, w, rsp, http.StatusOK)

	t, err := at.Service.AddTask(ctx, b.Title)
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
	rsp := struct {
		ID entity.TaskID `json:"id"`
	}{ID: t.ID}
	RespondJSON(ctx, w, rsp, http.StatusOK)
}

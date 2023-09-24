package handler

import (
	"go_todo_app/entity"
	"go_todo_app/store"
	"net/http"
)

type ListTask struct {
	Store *store.TaskStore
}

type task struct {
	ID     entity.TaskID     `json:"id"`
	Title  string            `json:"title"`
	Status entity.TaskStatus `json:"status"`
}

func (lt *ListTask) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	tasks := lt.Store.All()
	// rsp := []task{}
	// for _, t := range tasks {
	// 	rsp = append(rsp, task{
	// 		ID:     t.ID,
	// 		Title:  t.Title,
	// 		Status: t.Status,
	// 	})
	// }
	n := len(tasks)
	rsp := make([]task, n)
	for i, t := range tasks {
		rsp[i] = task{
			ID:     t.ID,
			Title:  t.Title,
			Status: t.Status,
		}
	}
	RespondJSON(ctx, w, rsp, http.StatusOK)
}

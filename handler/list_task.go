package handler

import (
	"go_todo_app/entity"
	"net/http"
)

type ListTask struct {
	Service ListTasksService
}

type task struct {
	ID     entity.TaskID     `json:"id"`
	Title  string            `json:"title"`
	Status entity.TaskStatus `json:"status"`
}

func (lt *ListTask) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	tasks, err := lt.Service.ListTasks(ctx)
	if err != nil {
		RespondJSON(ctx, w, &ErrResponse{
			Message: err.Error(),
		}, http.StatusInternalServerError)
		return
	}
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

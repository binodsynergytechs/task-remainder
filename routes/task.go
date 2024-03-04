package routes

import (
	"binod210/task_remainder/controller"

	"github.com/go-chi/chi/v5"
)

func InitRoutes() *chi.Mux {
	r := chi.NewRouter()

	taskController := controller.NewTaskController()
	remainderController := controller.NewRemainderController()
	r.Route("/v1", func(r chi.Router) {
		r.Route("/task", func(r chi.Router) {
			r.Post("/", taskController.Create)
			r.Patch("/{id}", taskController.Update)
			r.Delete("/{id}", taskController.Delete)
			r.Get("/{id}", taskController.Get)
			r.Get("/", taskController.GetAll)
		})
		r.Route("/reminder", func(r chi.Router) {
			r.Post("/", remainderController.Create)
			r.Patch("/{id}", remainderController.Update)
			r.Get("/{id}", remainderController.Get)
			r.Get("/", remainderController.GetAll)
			r.Delete("/{id}", remainderController.Delete)
		})
	})
	return r
}

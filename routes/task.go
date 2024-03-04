package routes

import (
	"binod210/task_remainder/controller"

	"github.com/go-chi/chi/v5"
)

func InitRoutes() *chi.Mux {
	r := chi.NewRouter()

	taskController := controller.GetTaskController()
	r.Route("/v1", func(r chi.Router) {
		r.Route("/task", func(r chi.Router) {
			r.Post("/", taskController.Create)
			r.Patch("/:id", taskController.Update)
			r.Delete("/:id", taskController.Delete)
			r.Get("/:id", taskController.Get)
			r.Get("/", taskController.GetAll)
		})
		r.Route("/remainder", func(r chi.Router) {
		})
	})
	return r
}

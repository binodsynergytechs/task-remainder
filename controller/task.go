package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"binod210/task_remainder/db"
	"binod210/task_remainder/models"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

type taskController struct{}

func NewTaskController() *taskController {
	return &taskController{}
}

func (tc *taskController) Create(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	task.CreatedAt = time.Now()
	result := db.DB.Create(&task)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func (tc *taskController) Get(w http.ResponseWriter, r *http.Request) {
	taskIdStr := chi.URLParam(r, "id")
	taskId, err := strconv.Atoi(taskIdStr)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	var task models.Task
	result := db.DB.First(&task, taskId)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			http.Error(w, "Task not found", http.StatusNotFound)
		} else {
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func (tc *taskController) Update(w http.ResponseWriter, r *http.Request) {
	taskIdStr := chi.URLParam(r, "id")
	taskId, err := strconv.Atoi(taskIdStr)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	var updatedTask models.Task
	if err := json.NewDecoder(r.Body).Decode(&updatedTask); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var existingTask models.Task
	result := db.DB.First(&existingTask, taskId)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			http.Error(w, "Task not found", http.StatusNotFound)
			return
		}
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	existingTask.Title = updatedTask.Title
	existingTask.Description = updatedTask.Description
	existingTask.Priority = updatedTask.Priority
	existingTask.DueAt = updatedTask.DueAt
	existingTask.UpdatedAt = time.Now()

	result = db.DB.Save(&existingTask)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(existingTask)
}

func (tc *taskController) Delete(w http.ResponseWriter, r *http.Request) {
	taskIdStr := chi.URLParam(r, "id")
	taskId, err := strconv.Atoi(taskIdStr)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	result := db.DB.Delete(&models.Task{}, taskId)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	if result.RowsAffected == 0 {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (tc *taskController) GetAll(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task
	result := db.DB.Find(&tasks)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

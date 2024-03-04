package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"binod210/task_remainder/db"
	"binod210/task_remainder/models"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

type remainderController struct{}

func NewRemainderController() *remainderController {
	return &remainderController{}
}

func (rc *remainderController) Create(w http.ResponseWriter, r *http.Request) {
	var reminder models.Reminder
	if err := json.NewDecoder(r.Body).Decode(&reminder); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := db.DB.Create(&reminder).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(reminder)
}

func (rc *remainderController) Get(w http.ResponseWriter, r *http.Request) {
	reminderIDStr := chi.URLParam(r, "id")
	reminderID, err := strconv.Atoi(reminderIDStr)
	if err != nil {
		http.Error(w, "Invalid reminder ID", http.StatusBadRequest)
		return
	}

	var reminder models.Reminder
	if err := db.DB.First(&reminder, reminderID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "Reminder not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reminder)
}

func (rc *remainderController) Update(w http.ResponseWriter, r *http.Request) {
	reminderIDStr := chi.URLParam(r, "id")
	reminderID, err := strconv.Atoi(reminderIDStr)
	if err != nil {
		http.Error(w, "Invalid reminder ID", http.StatusBadRequest)
		return
	}

	var updatedReminder models.Reminder
	if err := json.NewDecoder(r.Body).Decode(&updatedReminder); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var reminder models.Reminder
	if err := db.DB.First(&reminder, reminderID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "Reminder not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	reminder.TaskID = updatedReminder.TaskID
	reminder.Contact = updatedReminder.Contact
	reminder.ReminderAt = updatedReminder.ReminderAt

	if err := db.DB.Save(&reminder).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reminder)
}

func (rc *remainderController) Delete(w http.ResponseWriter, r *http.Request) {
	reminderIDStr := chi.URLParam(r, "id")
	reminderID, err := strconv.Atoi(reminderIDStr)
	if err != nil {
		http.Error(w, "Invalid reminder ID", http.StatusBadRequest)
		return
	}

	if err := db.DB.Delete(&models.Reminder{}, reminderID).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (rc *remainderController) GetAll(w http.ResponseWriter, r *http.Request) {
	var reminders []models.Reminder
	if err := db.DB.Find(&reminders).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reminders)
}

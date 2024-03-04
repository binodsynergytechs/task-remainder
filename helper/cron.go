package helper

import (
	"log"
	"time"

	"binod210/task_remainder/db"
	"binod210/task_remainder/models"

	"github.com/robfig/cron/v3"
)

func StartCron() {
	c := cron.New()

	// Schedule the job to run every minute
	_, err := c.AddFunc("* * * * *", sendReminders)
	if err != nil {
		log.Fatal("Error scheduling cron job: ", err)
	}

	// Start the cron scheduler
	c.Start()
}

func sendReminders() {
	now := time.Now()
	twoMinutesAgo := now.Add(-2 * time.Minute)

	// Retrieve reminders where reminder_at is within the last 2 minutes
	var reminders []models.Reminder
	if err := db.DB.Where("reminder_at BETWEEN ? AND ?", twoMinutesAgo, now).Find(&reminders).Error; err != nil {
		log.Println(err)
	} else {
		for _, remainder := range reminders {
			log.Println("sending remainder with appropriate method for ", remainder.Contact)
		}
	}
}

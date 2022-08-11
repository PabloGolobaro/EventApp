package core

import (
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_bot/models"
	"time"
)

func CheckTodayBirthdays(birthdays []models.Birthday) (result []models.Birthday) {
	result = make([]models.Birthday, 0)
	for _, birthday := range birthdays {
		if birthday.BirthDate.Day() == time.Now().Day() {
			result = append(result, birthday)
		}
	}
	return
}

func CheckTomorrowBirthdays(birthdays []models.Birthday) (result []models.Birthday) {
	result = make([]models.Birthday, 0)
	for _, birthday := range birthdays {
		if birthday.BirthDate.Day() == time.Now().Add(time.Hour*24).Day() {
			result = append(result, birthday)
		}
	}
	return

}

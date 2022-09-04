package core

import (
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/models"
	"time"
)

func CheckTodayBirthdays(birthdays []models.Birthday) (result []models.Birthday) {
	result = make([]models.Birthday, 0)
	for _, birthday := range birthdays {
		if birthday.BirthDate.Day() == time.Now().Day() && birthday.BirthDate.Month() == time.Now().Month() {
			result = append(result, birthday)
		}
	}
	return
}

func CheckTomorrowBirthdays(birthdays []models.Birthday) (result []models.Birthday) {
	result = make([]models.Birthday, 0)
	tomorrow := time.Now().Add(time.Hour * 24)
	for _, birthday := range birthdays {
		if birthday.BirthDate.Day() == tomorrow.Day() && birthday.BirthDate.Month() == tomorrow.Month() {
			result = append(result, birthday)
		}
	}
	return

}
func CheckMonthBirthdays(birthdays []models.Birthday) (result []models.Birthday) {
	result = make([]models.Birthday, 0)
	for _, birthday := range birthdays {
		if birthday.BirthDate.Month() == time.Now().Month() {
			result = append(result, birthday)
		}
	}
	return

}

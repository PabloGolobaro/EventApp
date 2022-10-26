package helpers

import (
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/models"
	"sort"
	"time"
)

func Shorten_date(time time.Time) string {
	return time.Format("02.01.2006")
}

func Expired_date(t time.Time) bool {

	if t.Month() < time.Now().Month() {
		return true
	} else if t.Month() == time.Now().Month() && t.Day() <= time.Now().Day() {
		return true
	} else {
		return false
	}

}
func Sort_birthdays(birthdays []models.Birthday) []models.Birthday {
	sort.Slice(birthdays, func(i, j int) bool {
		month_i := birthdays[i].BirthDate.Month()
		month_j := birthdays[j].BirthDate.Month()
		if month_j == month_i {
			day_i := birthdays[i].BirthDate.Day()
			day_j := birthdays[j].BirthDate.Day()
			return day_i < day_j
		} else {
			return month_i < month_j
		}

	})
	return birthdays
}

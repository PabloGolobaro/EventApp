package helpers

import (
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/daos"
	"log"
	"strings"
)

func CheckUserPass(username, password string) bool {
	user, err := daos.NewUserDAO().ReadByUsername(username)
	if err != nil {
		return false
	}
	log.Println("checkUserPass", username, password, user.Password)

	if password == user.Password {
		return true
	} else {
		return false
	}

}

func EmptyUserPass(username, password string) bool {
	return strings.Trim(username, " ") == "" || strings.Trim(password, " ") == ""
}

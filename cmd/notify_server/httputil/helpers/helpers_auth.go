package helpers

import (
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/daos"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/localconf"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

func CheckUserPass(username, password string) bool {
	user, err := daos.NewUserDAO(localconf.Config.DB).ReadByUsername(username)
	if err != nil {
		return false
	}
	if CheckPasswordHash(password, user.PasswordHash) {
		return true
	} else {
		return false
	}

}

func EmptyUserPass(username, password string) bool {
	return strings.Trim(username, " ") == "" || strings.Trim(password, " ") == ""
}
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

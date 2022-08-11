package test_data

import (
	"fmt"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_bot/config"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_bot/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"io/ioutil"
	"strings"
)

func init() {
	// the test may be started from the home directory or a subdirectory
	err := config.LoadConfig("C:\\Users\\Professional\\GolandProjects\\go-notify-project\\config") // on host use absolute path
	if err != nil {
		panic(err)
	}
	config.Config.DB, config.Config.DBErr = gorm.Open(sqlite.Open("tests.db"), &gorm.Config{})
	config.Config.DB.Exec("PRAGMA foreign_keys = ON") // SQLite defaults to `foreign_keys = off'`
	if config.Config.DBErr != nil {
		panic(config.Config.DBErr)
	}

	config.Config.DB.AutoMigrate(&models.Birthday{})
}

// Resets testing database - deletes all tables, creates new ones using GORM migration and populates them using `db.sql` file
func ResetDB() *gorm.DB {
	config.Config.DB.Migrator().DropTable(&models.Birthday{}) // Note: Order matters
	config.Config.DB.AutoMigrate(&models.Birthday{})
	if err := runSQLFile(config.Config.DB, getSQLFile()); err != nil {
		panic(fmt.Errorf("error while initializing test database: %s", err))
	}
	return config.Config.DB
}

func getSQLFile() string {
	return "C:\\Users\\Professional\\GolandProjects\\go-notify-project\\cmd\\notify_bot\\test_data\\db.sql" // on host use absolute path
}

func GetTestCaseFolder() string {
	return "C:\\Users\\Professional\\GolandProjects\\go-notify-project\\cmd\\notify_bot\\test_data\\test_case_data" // on host use absolute path
}

// Executes SQL file specified by file argument
func runSQLFile(db *gorm.DB, file string) error {
	s, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	lines := strings.Split(string(s), ";")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if result := db.Exec(line); result.Error != nil {
			fmt.Println(line)
			return result.Error
		}
	}
	return nil
}

package config

import (
	"fmt"
	"github.com/PabloGolobaro/go-notify-project/cmd/notify_server/models"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"log"
	"sync"
)

var Config botConfig

type botConfig struct {
	PageMap MutexMap
	Cache   BirthdaysCache
	// the data source name (DSN) for connecting to the database. required.
	DSN string `mapstructure:"dsn"`
	//admins id
	Admins []string `mapstructure:"admins"`
	// bot token
	Token string `mapstructure:"token"`
	//site domain
	Domain string `mapstructure:"domain"`
	// the shared DB ORM object
	DB *gorm.DB
	// the error thrown be GORM when using DB ORM object
	DBErr error
}

type MutexMap struct {
	Map map[int64]int
	sync.RWMutex
}
type BirthdaysCache struct {
	M map[uint][]models.Birthday
	sync.RWMutex
}

func LoadConfig() error {
	log.Println("Loading config...")
	v := viper.New()
	v.SetConfigName("bot_config")
	v.SetConfigType("yaml")
	v.AddConfigPath("./config")

	//for _, path := range configPaths {
	//	v.AddConfigPath(path)
	//}
	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read the configuration file: %s", err)
	}

	log.Println("Loaded config...")
	Config.PageMap.Map = make(map[int64]int)
	Config.Cache.M = make(map[uint][]models.Birthday)
	return v.Unmarshal(&Config)
}

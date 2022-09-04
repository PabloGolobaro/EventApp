package config

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"log"
	"strconv"
)

// Config is global object that holds all application level variables.
var Config appConfig

type appConfig struct {
	// the shared DB ORM object
	DB *gorm.DB
	// the error thrown be GORM when using DB ORM object
	DBErr error
	// the server port. Defaults to 8080
	ServerPort int `mapstructure:"server_port"`
	// the data source name (DSN) for connecting to the database. required.
	DSN string `mapstructure:"dsn"`
	// the API key needed to authorize to API. required.
	ApiKey string `mapstructure:"api_key"`
	// Certificate file for HTTPS
	CertFile string `mapstructure:"cert_file"`
	// Private key file for HTTPS
	KeyFile string `mapstructure:"key_file"`
}

// LoadConfig loads config from files
func LoadConfig(configPaths ...string) error {
	log.Println("Loading config...")
	v := viper.New()
	v.SetEnvPrefix("notify")
	v.AutomaticEnv()

	Config.DSN = v.Get("dsn").(string)
	log.Println("Got DSN...")
	Config.ApiKey = v.Get("api_key").(string)
	log.Println("Got API_KEY...")
	port := v.Get("server_port").(string)
	Config.ServerPort, _ = strconv.Atoi(port)
	log.Println("Got SERVER_PORT...")
	log.Println("Loaded config...")
	return v.Unmarshal(&Config)
}

package localconf

import (
	api "github.com/PabloGolobaro/go-notify-project/cmd/notify_server/api/smtp_api"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"log"
	"strconv"
	"time"
)

// Config is global object that holds all application level variables.
var Config appConfig

type appConfig struct {
	API *api.APIClient
	// the shared DB ORM object
	DB *gorm.DB
	// the error thrown be GORM when using DB ORM object
	DBErr error
	// the server port. Defaults to 8080
	ServerPort int `mapstructure:"SERVER_PORT"`
	//site domain
	Domain string `mapstructure:"DOMAIN"`
	// the API key needed to authorize to API. required.
	ApiKey string `mapstructure:"API_KEY"`

	DBHost         string `mapstructure:"POSTGRES_HOST"`
	DBUserName     string `mapstructure:"POSTGRES_USER"`
	DBUserPassword string `mapstructure:"POSTGRES_PASSWORD"`
	DBName         string `mapstructure:"POSTGRES_DB"`
	DBPort         string `mapstructure:"POSTGRES_PORT"`

	ClientOrigin string `mapstructure:"CLIENT_ORIGIN"`

	TokenSecret    string        `mapstructure:"TOKEN_SECRET"`
	TokenExpiresIn time.Duration `mapstructure:"TOKEN_EXPIRED_IN"`
	TokenMaxAge    int           `mapstructure:"TOKEN_MAXAGE"`

	EmailFrom string `mapstructure:"EMAIL_FROM"`
	SMTPHost  string `mapstructure:"SMTP_HOST"`
	SMTPPass  string `mapstructure:"SMTP_PASS"`
	SMTPPort  int    `mapstructure:"SMTP_PORT"`
	SMTPUser  string `mapstructure:"SMTP_USER"`

	SMTPBZ_API_KEY string `mapstructure:"SMTPBZ_API_KEY"`
}

// LoadConfig loads config from files
func LoadConfig(path string) error {
	log.Println("Loading config...")
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	log.Println("Loaded config...")
	return viper.Unmarshal(&Config)
}

func loadConfigFromEnvVariables(configPaths ...string) error {
	log.Println("Loading config...")
	v := viper.New()
	v.SetEnvPrefix("notify")
	v.AutomaticEnv()

	Config.ApiKey = v.Get("api_key").(string)
	log.Println("Got API_KEY...")
	port := v.Get("server_port").(string)
	Config.ServerPort, _ = strconv.Atoi(port)
	log.Println("Got SERVER_PORT...")
	Config.Domain = v.Get("domain").(string)
	log.Println("Got DOMAIN...")
	log.Println("Loaded config...")
	return v.Unmarshal(&Config)
}

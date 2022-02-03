package config

import (
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/kkyr/fig"
)

// Config structure for all settings of application
type Config struct {
	App struct {
		ServerAddress  string        `fig:"serverAddress" envconfig:"ADDRESS"` //default:"localhost"`
		ServerPort     int           `fig:"serverPort" envconfig:"PORT" default:"9000"`
		StoragePath    string        `fig:"storagePath" envconfig:"STORAGE" default:"../internal/storage/storage.db" `
		TimeoutRequest time.Duration `fig:"timeoutRequest" envconfig:"TIMEOUT" default:"10"`
		LogLevel       string        `envconfig:"GIN_MODE" default:"debug"`
	} `fig:"app"`

	DBConfig struct {
		Host     string `fig:"host" envconfig:"DB_HOST"`
		User     string `fig:"user" envconfig:"DB_USER"`
		Password string `fig:"password" envconfig:"DB_PASSWORD"`
		DBName   string `fig:"dbName" envconfig:"DB_NAME"`
		Port     int    `fig:"port" envconfig:"DB_PORT"`
	} `fig:"dbConfig"`
}

// InitConfig function for initialize Config structure
func InitConfig(useConfig *string) (*Config, error) {
	var cfg = Config{}
	err := fig.Load(&cfg, fig.File(*useConfig))
	if err != nil {
		err = fig.Load(&cfg, fig.File("config.yml"))
		if err != nil {
			log.Printf("can't load configuration file: %s", err)
			err = envconfig.Process("", &cfg)
			if err != nil {
				log.Printf("failed to read config from env: %v", err)
				return nil, err
			}
		}
	}

	return &cfg, err
}

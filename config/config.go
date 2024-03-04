package config

import (
	"errors"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Environment string

const (
	Dev  Environment = "dev"
	Prod Environment = "prod"
	Stag Environment = "stag"
)

// load .env file for creating environment variable
func LoadDotEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
}

var (
	instance *Config
	once     sync.Once
)

// loading GOENV and get configs
func GetConfig() *Config {
	env := os.Getenv("GOENV")
	once.Do(func() {
		instance = &Config{}
		if err := initConfig(Environment(env)); err != nil {
			log.Fatalf("error initializing config: %v", err)
		}
	})
	return instance
}

// config structure to parse yml file
type Config struct {
	Database struct {
		Dsn string `mapstructure:"dsn"`
	} `mapstructure:"database"`

	App struct {
		Port  string `mapstructure:"port"`
		Debug bool   `mapstructure:"debug"`
	}

	JWT struct {
		Secret string `mapstructure:"secret"`
		Expiry int    `mapstructure:"expiry"`
	} `mapstructure:"jwt"`

	Twilio struct {
		SenderName     string `mapstructure:"sender_name"`
		SenderEmailId  string `mapstructure:"sender_email_id"`
		SendgridApiKey string `mapstructure:"sendgrid_api_key"`
	} `mapstructure:"twilio"`
}

// read .yml file from config with env selection
func initConfig(env Environment) error {
	if env == "" {
		return errors.New("environment not set")
	}
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigName(string(env))
	v.AddConfigPath("./config/")
	v.AddConfigPath("config")

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return errors.New("configuration file not found")
		}
		return errors.New("error on parsing configuration file")
	}

	if err := v.Unmarshal(instance); err != nil {
		return errors.New("error on parsing configuration file to struct")
	}

	v.WatchConfig()

	return nil
}

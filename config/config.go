package config

import (
	"os"
	"strconv"

	"go.uber.org/zap"
)

type Credentials struct {
	Database string
	Host     string
	Port     string
	User     string
	Password string
}

type Mqtt struct {
	HiveMQ   string
	Username string
	Password string
	ClientID string
	Topic    string
}

type ServiceConfig struct {
	Host string
	Port int
}

type Config struct {
	Creden  Credentials
	Service ServiceConfig
	MQTT    Mqtt
}

func NewConfig(log *zap.Logger) *Config {
	var config Config
	log.Info("ENV")
	if config.Creden.Database = os.Getenv("DATABASE"); config.Creden.Database == "" {
		log.Fatal("Database environment variable not set")
	}

	if config.Creden.Host = os.Getenv("HOST"); config.Creden.Host == "" {
		log.Fatal("HOST environment variable not set")
	}

	if config.Creden.Port = os.Getenv("PORT"); config.Creden.Port == "" {
		log.Fatal("PORT environment variable not set")
	}

	if config.Creden.User = os.Getenv("USER"); config.Creden.User == "" {
		log.Fatal("USER environment variable not set")
	}

	if config.Creden.Password = os.Getenv("PASSWORD"); config.Creden.Password == "" {
		log.Fatal("PASSWORD environment variable not set")
	}

	if config.MQTT.HiveMQ = os.Getenv("HIVEMQ"); config.MQTT.HiveMQ == "" {
		log.Fatal("HIVEMQ environment variable not set")
	}

	if config.MQTT.Username = os.Getenv("HUSERNAME"); config.MQTT.Username == "" {
		log.Fatal("HUSERNAME environment variable not set")
	}

	if config.MQTT.Password = os.Getenv("HPASSWORD"); config.MQTT.Password == "" {
		log.Fatal("HPASSWORD environment variable not set")
	}

	if config.MQTT.ClientID = os.Getenv("HCLIENT"); config.MQTT.ClientID == "" {
		log.Fatal("HCLIENT environment variable not set")
	}

	if config.MQTT.Topic = os.Getenv("HTOPIC"); config.MQTT.Topic == "" {
		log.Fatal("HTOPIC environment variable not set")
	}

	if config.Service.Host = os.Getenv("SERVICEHOST"); config.Service.Host == "" {
		log.Fatal("SERVICEHOST environment variable not set")
	}

	if config.Service.Port, _ = strconv.Atoi(os.Getenv("SERVICEPORT")); config.Service.Port == 0 {
		log.Fatal("SERVICEPORT environment variable not set")
	}

	return &config
}

func (config *Config) GetCredentials() Credentials {
	return config.Creden
}

func (config *Config) GetServiceConfig() ServiceConfig {
	return config.Service
}

func (config *Config) GetMqtt() Mqtt {
	return config.MQTT
}

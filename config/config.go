package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"
)

type Config struct {
	AppName            string   `yaml:"app_name" envconfig:"APP_NAME"    required:"true"`
	AppVersion         string   `yaml:"app_version" envconfig:"APP_VERSION" required:"true"`
	ClickhouseAddr     []string `yaml:"clickhouse_addr" envconfig:"CLICKHOUSE_ADDR" required:"true"`         // хост:порт clickhouse
	ClickhouseDatabase string   `yaml:"clickhouse_database" envconfig:"CLICKHOUSE_DATABASE" required:"true"` // название базы clickhouse
	ClickhouseUsername string   `yaml:"clickhouse_username" envconfig:"CLICKHOUSE_USERNAME" required:"true"` // имя пользователя базы clickhouse
	ClickhousePassword string   `yaml:"clickhouse_password" envconfig:"CLICKHOUSE_PASSWORD" required:"true"` // пароль пользователя базы clickhouse
	GrpcPort           string   `yaml:"grpc_port" envconfig:"GRPC_PORT" required:"true"`
	JWTServiceName     string   `yaml:"jwt_service_name" envconfig:"JWT_SERVICE_NAME" required:"true"`     // Название сервиса (для сверки с JWTValidServices при авторизаии)
	JWTSecret          string   `yaml:"jwt_secret" envconfig:"JWT_SECRET" required:"true"`                 // JWT секрет
	JWTValidServices   []string `yaml:"jwt_valid_services" envconfig:"JWT_VALID_SERVICES" required:"true"` // список микросервисов (через запятую), которым разрешен доступ
}

func New(filePath string, envFile string) (Config, error) {
	var config Config
	var err error

	// 1. Читаем из config.yaml. Самый низкий приоритет
	file, err := os.Open(filePath)
	if err == nil {
		defer file.Close()
		decoder := yaml.NewDecoder(file)
		if decodeErr := decoder.Decode(&config); decodeErr != nil {
			log.Fatalf("Ошибка при чтении config.yaml: %v", decodeErr)
		}
	} else {
		log.Printf("config.yaml не найден, используются значения по умолчанию: %v", err)
	}

	// 2.1 Загрузка переменных окружения из .env
	err = godotenv.Load(envFile)
	if err != nil {
		return config, fmt.Errorf("godotenv.Load: %w", err)
	}

	// 2.2 Переопределяем переменные, полученные из конфиг файла
	err = envconfig.Process("", &config)
	if err != nil {
		return config, fmt.Errorf("envconfig.Process: %w", err)
	}

	// 3. Чтение параметров командной строки
	// ... TODO добавить по необходимости

	return config, nil
}

package main

import (
	"form-to-mail-service/models"
	"form-to-mail-service/pkg/handler"
	"form-to-mail-service/pkg/repository"
	"form-to-mail-service/pkg/service"
	"github.com/joho/godotenv"
	"github.com/siruspen/logrus"
	"github.com/spf13/viper"
	"os"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	if err := initConfig(); err != nil {
		logrus.Fatal("Error while initializing config: ", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatal("Error while loading .env file: ", err.Error())
	}

	repos := repository.NewRepository(&repository.SMTPConfig{
		EmailAddress: os.Getenv("SMTP_EMAIL"),
		Password:     os.Getenv("SMTP_PASSWORD"),
		Host:         "smtp.yandex.ru",
		Port:         viper.GetString("smtp.port"),
	})

	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(models.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("Error while running http server: ", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}

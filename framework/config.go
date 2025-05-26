package framework

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Prefix        string
	BotName       string
	Version       string
	BotToken      string
	DefaultStatus string
	ServerIP      string
	RconPort      string
	RconPassword  string
	HttpPort      string
	HttpHost      string
}

func LoadConfig() *Config {
	var isProd bool = os.Getenv("BOT_ENV") == "prod"
	if !isProd {
		log.Println("Running in local development mode")
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file", err)
		}
	}

	var httpPort string = os.Getenv("PORT")
	if !isProd {
		httpPort = os.Getenv("HTTP_PORT")
	}

	conf := new(Config)
	conf.BotToken = os.Getenv("BOT_TOKEN")
	conf.BotName = os.Getenv("NAME")
	conf.Version = os.Getenv("VERSION")
	conf.Prefix = os.Getenv("PREFIX")
	conf.ServerIP = os.Getenv("RCON_HOST")
	conf.RconPassword = os.Getenv("RCON_PASSWORD")
	conf.RconPort = os.Getenv("RCON_PORT")
	conf.DefaultStatus = os.Getenv("DEFAULT_STATUS")
	conf.HttpPort = httpPort
	conf.HttpHost = os.Getenv("HTTP_HOST")

	log.Println("Bot env (prod): ", isProd)
	log.Println("Bot token: ", conf.BotToken)
	log.Println("Bot name: ", conf.BotName)
	log.Println("Bot version: ", conf.Version)
	log.Println("Bot prefix: ", conf.Prefix)
	log.Println("Bot default status: ", conf.DefaultStatus)
	log.Println("MC server: ", conf.ServerIP)
	log.Println("MC Server (RCON) port: ", conf.RconPort)
	log.Println("Http port: ", conf.HttpPort)
	log.Println("Http host: ", conf.HttpHost)

	return conf
}

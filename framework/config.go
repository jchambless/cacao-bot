package framework

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Prefix        string
	BotName       string
	Version       string
	BotToken      string
	DefaultStatus string
	ServerIP      string
	RconPort      int
	RconPassword  string
	HttpPort      string
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

	conf := new(Config)
	conf.BotToken = os.Getenv("BOT_TOKEN")
	conf.BotName = os.Getenv("BOT_NAME")
	conf.Version = os.Getenv("BOT_VERSION")
	conf.Prefix = os.Getenv("BOT_PREFIX")
	conf.ServerIP = os.Getenv("MC_SERVER")
	conf.RconPassword = os.Getenv("MC_RCON_PASSWORD")
	conf.DefaultStatus = os.Getenv("BOT_DEFAULT_STATUS")
	conf.HttpPort = os.Getenv("PORT")

	log.Println("Bot env (prod): ", isProd)
	log.Println("Bot token: ", conf.BotToken)
	log.Println("Bot name: ", conf.BotName)
	log.Println("Bot version: ", conf.Version)
	log.Println("Bot prefix: ", conf.Prefix)
	log.Println("Bot default status: ", conf.DefaultStatus)
	log.Println("MC server: ", conf.ServerIP)
	log.Println("Http port: ", conf.HttpPort)

	rcon, err := strconv.Atoi(os.Getenv("MC_RCON_PORT"))
	if err != nil {
		log.Fatal("MC Rcon port is not set")
	}
	conf.RconPort = rcon

	return conf
}

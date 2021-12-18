package framework

import (
	"log"
	"os"
	"strconv"
	//"github.com/joho/godotenv"
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
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file", err)
	// }

	conf := new(Config)
	conf.BotToken = os.Getenv("BOT_TOKEN")
	conf.BotName = os.Getenv("BOT_NAME")
	conf.Version = os.Getenv("BOT_VERSION")
	conf.Prefix = os.Getenv("BOT_PREFIX")
	conf.ServerIP = os.Getenv("MC_SERVER")
	conf.RconPassword = os.Getenv("MC_RCON_PASSWORD")
	conf.DefaultStatus = os.Getenv("BOT_DEFAULT_STATUS")
	conf.HttpPort = os.Getenv("PORT")

	rcon, err := strconv.Atoi(os.Getenv("MC_RCON_PORT"))
	if err != nil {
		log.Fatal("MC Rcon port is not set")
	}
	conf.RconPort = rcon

	return conf
}

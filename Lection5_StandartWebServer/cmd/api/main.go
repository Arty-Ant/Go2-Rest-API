package main

import (
	"flag"
	"log"

	"github.com/Arty-Ant/Go2-Rest-API/tree/74676d691751daff73c0e61c9c054ad9dbc88bc1/Lection5_StandartWebServer/internal/app/api"
	"github.com/BurntSushi/toml"
)

var (
	configPath string
)

func init() {
	//Скажем, что наше приложение будет на этапе запуска получать путь до конфиг файла из внешнего мира
	flag.StringVar(&configPath, "path", "configs/api.toml", "path to config file in .toml format")
}

func main() {
	//В этот момент происходит инициализация переменной configPath значением
	flag.Parse()
	log.Println("It works")
	//server instance initialization
	config := api.NewConfig()
	_, err := toml.DecodeFile(configPath, config) // Десериализиуете содержимое .toml файла
	if err != nil {
		log.Println("can not find configs file. using default values:", err)
	}

	server := api.New(config)

	//api server start
	log.Fatal(server.Start())
}

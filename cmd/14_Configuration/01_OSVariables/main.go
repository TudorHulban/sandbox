package main

import (
	"fmt"

	"github.com/caarlos0/env/v7"
	"github.com/rs/zerolog/log"
)

type ConfigPorts struct {
	Port1 string `env:"PORT1,required" envDefault:"9001"`
	Port2 string `env:"PORT2,required" envDefault:"9002"`
}

type Endpoints struct {
	Endpoint1 string `env:"ENDPOINT1,required" envDefault:"abcd"`
	Endpoint2 string `env:"ENDPOINT2,required" envDefault:"efgh"`
}

type Configuration struct {
	ConfigPorts
	Endpoints

	DevelopmentMode bool `env:"DEV_MODE,required" envDefault:"true"`
}

func main() {
	var config Configuration

	if errParse := env.Parse(&config); errParse != nil {
		log.Error().Msg(errParse.Error())
	}

	fmt.Printf("%+v\n", config)
	// should print:
	// {ConfigPorts:{Port1:9001 Port2:9002} Endpoints:{Endpoint1:abcd Endpoint2:efgh} DevelopmentMode:true}
}

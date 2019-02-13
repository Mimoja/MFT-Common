package MFTCommon

import (
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

type AppBundle struct {
	MessageQueue MessageBundle
	Storage      Storage
	DB           DataBase
	Log          *logrus.Logger
	Config       AppConfiguration
}

func Init(serviceName string) AppBundle {

	args := os.Args
	if(len(args) < 2){
		log.Fatalf("Path to config file missing. Please provide via '%s pathToConfig.yml'", os.Args[0])
	}

	config := ConfigInit(os.Args[1])
	log := LoggerConnect(config, serviceName)

	return AppBundle{
		MessageQueue: mqConnect(config, log),
		Storage:      storageConnect(config, log),
		DB:           dbConnect(config, log),
		Log:          log,
		Config:       config,
	}

}

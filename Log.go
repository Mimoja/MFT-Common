package MFTCommon

import (
	"github.com/olivere/elastic"
	"github.com/sirupsen/logrus"
	"gopkg.in/sohlich/elogrus.v3"
)

func loggerConnect(config *AppConfiguration, service string) *logrus.Logger {

	log := logrus.New()

	log.Level = logrus.DebugLevel

	client, err := elastic.NewSimpleClient(elastic.SetURL(config.Logger.Protocol + "://" + config.Logger.URI))
	if err != nil {
		log.WithError(err).Error("Could not connect to Elastic for logging")
		return log
	}

	hook, err := elogrus.NewElasticHook(client, service, logrus.DebugLevel, "logging")
	if err != nil {
		log.WithError(err).Error("Could not create logging hook")
		return log
	}

	log.Hooks.Add(hook)
	log.Info("Logger up and running")
	return log
}

package MFTCommon

import (
	"github.com/spf13/viper"
	"log"
	"path/filepath"
	"strings"
)

type AppConfiguration struct {
	Logger LoggerConfiguration
	MQ MQConfiguration
	Storage StorageConfiguration
	DB DBConfiguration
}

func ConfigInit(file string) AppConfiguration{
	viper.SetConfigType("yaml")
	base := filepath.Base(file)
	base = strings.TrimSuffix(base, filepath.Ext(base))
	viper.SetConfigName(base)
	viper.AddConfigPath(filepath.Dir(file))
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	var appConfig AppConfiguration

	err := viper.Unmarshal(&appConfig)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	return appConfig
}
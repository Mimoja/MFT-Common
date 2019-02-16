package MFTCommon

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"path/filepath"
	"strings"
)

type LoggerConfiguration struct {
	Protocol string
	URI      string
}

type MQConfiguration struct {
	User     string
	Password string
	URI      string
}

type StorageConfiguration struct {
	URI             string
	AccessKeyID     string
	SecretAccessKey string
	UseSSL          bool
}

type DBConfiguration struct {
	URI      string
	Protocol string
}

type AppRunConfiguration struct {
	Frontend struct {
		DownloadEnabled bool
		ReportResults   int
	}
}

type AppConfiguration struct {
	Logger  LoggerConfiguration
	MQ      MQConfiguration
	Storage StorageConfiguration
	DB      DBConfiguration
	App     AppRunConfiguration
}

func ConfigInit(file string) *AppConfiguration {
	var appConfig AppConfiguration

	viper.SetConfigType("yaml")
	base := filepath.Base(file)
	base = strings.TrimSuffix(base, filepath.Ext(base))
	viper.SetConfigName(base)
	viper.AddConfigPath(filepath.Dir(file))
	viper.AddConfigPath(".")
	viper.WatchConfig()

	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Println("Config file changed:", e.Name)
		err := viper.Unmarshal(&appConfig)
		if err != nil {
			log.Fatalf("unable to decode into struct, %v", err)
		}
	})

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&appConfig)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	return &appConfig
}

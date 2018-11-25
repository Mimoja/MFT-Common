package main

import (
	"github.com/Mimoja/MimojaFirmwareToolkit/pkg/Common"
)

var MQ MFTCommon.MessageBundle

func main() {

	Log := MFTCommon.ConnectLogger("QueueListener")
	MQ := MFTCommon.MQConnect(Log)

	MQ.URLQueue.RegisterCallback("Downloader", func(payload string) error {
		Log.WithField("queue", "url").Infof("[URL]: '%s'\n", payload)
		return nil
	})
/*
	MQ.DownloadedQueue.RegisterCallback("Downloader", func(payload string) error {
		Log.WithField("queue", "downloaded").Infof("[Downloaded]: '%s'\n", payload)
		return nil
	})

	MQ.DeleteQueue.RegisterCallback("DownloaderTestConsumer", func(payload string) error {
		Log.WithField("queue", "Delete").Infof("[Delete]: '%s'\n", payload)
		return nil
	})

	MQ.MEImagesQueue.RegisterCallback("DownloaderTestConsumer", func(payload string) error {
		Log.WithField("queue", "MEImages").Infof("[MEImages]: '%s'\n", payload)
		return nil
	})

	MQ.FlashImagesQueue.RegisterCallback("DownloaderTestConsumer", func(payload string) error {
		Log.WithField("queue", "FlashImages").Infof("[FlashImages]: '%s'\n", payload)
		return nil
	})

	MQ.BiosImagesQueue.RegisterCallback("DownloaderTestConsumer", func(payload string) error {
		Log.WithField("queue", "bios").Infof("[Bios]: '%s'\n", payload)
		return nil
	})
	MQ.TestQueue.RegisterCallback("DownloaderTestConsumer", func(payload string) error {
		Log.WithField("queue", "test").Infof("[TEST]: '%s'\n", payload)
		return nil
	})
	*/
	select {}
}

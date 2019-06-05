package main

import (
	"MimojaFirmwareToolkit/pkg/Common"
)

var Bundle MFTCommon.AppBundle

func main() {
	Bundle = MFTCommon.Init("QueueTest")

	Bundle.MessageQueue.URLQueue.RegisterCallback("Downloader", func(payload string) error {
		Bundle.Log.WithField("queue", "url").Infof("[URL]: '%s'\n", payload)
		return nil
	})

	Bundle.MessageQueue.DownloadedQueue.RegisterCallback("Importer", func(payload string) error {
		Bundle.Log.WithField("queue", "downloaded").Infof("[Downloaded]: '%s'\n", payload)
		return nil
	})

	Bundle.MessageQueue.DeleteQueue.RegisterCallback("DownloaderTestConsumer", func(payload string) error {
		Bundle.Log.WithField("queue", "Delete").Infof("[Delete]: '%s'\n", payload)
		return nil
	})

	Bundle.MessageQueue.MEImagesQueue.RegisterCallback("DownloaderTestConsumer", func(payload string) error {
		Bundle.Log.WithField("queue", "MEImages").Infof("[MEImages]: '%s'\n", payload)
		return nil
	})

	Bundle.MessageQueue.FlashImagesQueue.RegisterCallback("DownloaderTestConsumer", func(payload string) error {
		Bundle.Log.WithField("queue", "FlashImages").Infof("[FlashImages]: '%s'\n", payload)
		return nil
	})

	Bundle.MessageQueue.BiosImagesQueue.RegisterCallback("DownloaderTestConsumer", func(payload string) error {
		Bundle.Log.WithField("queue", "bios").Infof("[Bios]: '%s'\n", payload)
		return nil
	})
	Bundle.MessageQueue.TestQueue.RegisterCallback("DownloaderTestConsumer", func(payload string) error {
		Bundle.Log.WithField("queue", "test").Infof("[TEST]: '%s'\n", payload)
		return nil
	})

	select {}
}

package main

import (
	"fmt"
	"github.com/Mimoja/MimojaFirmwareToolkit/pkg/Common"
	"time"
)

var Bundle MFTCommon.AppBundle

func main() {
	Bundle = MFTCommon.Init("MQTest")

	cha := make(chan bool)

	Bundle.MessageQueue.TestQueue.RegisterCallback("DownloaderTestConsumer", func(payload string) error {
		Bundle.Log.Infof("Payload: '%s'", payload)
		panic("1237")
		//cha <- true
		err := fmt.Errorf("penis")
		return err
	})

	Bundle.MessageQueue.TestQueue.Send("Test")

	<-cha

	time.Sleep(1 * time.Second)

}

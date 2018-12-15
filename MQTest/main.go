package main

import (
	"fmt"
	"github.com/Mimoja/MimojaFirmwareToolkit/pkg/Common"
	"time"
)

var MQ MFTCommon.MessageBundle

func main() {

	Log := MFTCommon.ConnectLogger("TestService")
	MQ := MFTCommon.MQConnect(Log)

	cha := make(chan bool)

	MQ.TestQueue.RegisterCallback("DownloaderTestConsumer", func(payload string) error {
		Log.Infof("Payload: '%s'", payload)
		panic("1237")
		//cha <- true
		err :=  fmt.Errorf("penis")
		return err
	})

	MQ.TestQueue.Send("Test")

	<- cha

	time.Sleep(1*time.Second)

}

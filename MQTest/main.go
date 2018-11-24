package main

import (
	"fmt"
	"github.com/Mimoja/MimojaFirmwareToolkit/pkg/Common"
	"math/rand"
	"time"
)

var MQ MFTCommon.MessageBundle

func main() {

	fmt.Printf("%f", rand.Float32())
	Log := MFTCommon.ConnectLogger("TestService")
	MQ := MFTCommon.MQConnect(Log, "testService")

	cha := make(chan bool)

	MQ.TestQueue.RegisterCallback("DownloaderTestConsumer", func(payload string) error {
		Log.Infof("Payload: '%s'", payload)
		//cha <- true
		err :=  fmt.Errorf("penis")
		return err
	})

	MQ.TestQueue.Send("Test")

	<- cha

	time.Sleep(1*time.Second)

}

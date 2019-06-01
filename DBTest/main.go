package main

import (
	"MimojaFirmwareToolkit/pkg/Common"
)

var Bundle MFTCommon.AppBundle

func main() {
	Bundle = MFTCommon.Init("DBTest")

	data := []byte{0x1, 0x2, 0x3, 0x4}

	id := "2345"
	idData := MFTCommon.GenerateID(data)
	Bundle.DB.StoreElement("testing", nil, idData, &id)

}

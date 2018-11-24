package main

import (
	"github.com/Mimoja/MimojaFirmwareToolkit/pkg/Common"
)

func main() {
	Log := MFTCommon.ConnectLogger("DBTest")
	db := MFTCommon.DBConnect(Log)

	data := []byte{0x1, 0x2, 0x3, 0x4}

	id := "2345"
	idData := MFTCommon.GenerateID(data)
	db.StoreElement("testing", nil, idData, &id)

}

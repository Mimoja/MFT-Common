package MFTCommon

import (
	"github.com/mimoja/amdfw"
)

type (
	AMDFirmware struct {
		AGESA []AMDAGESA
		Firmware *AMDImage
	}

	AMDBinaryEntry struct {
		Header    map[string]string `json:Header`
		Signature string
		Comment   []string
		TypeInfo  *amdfw.TypeInfo
		Version   string
		Size      string
		Type      string
	}

	AMDTypeInfo struct {
		Type    string
		Name    string
		Comment string
	}

	DBEntry struct {
		AMDBinaryEntry
		ID IDEntry
	}

	AMDAGESA struct {
		Header string
		Raw    string
		Offset uint32
	}

	AMDImage struct {
		FET          *AMDFirmwareEntryTable
		FlashMapping string
		Roms         []AMDRom
	}

	AMDFirmwareEntryTable struct {
		Location string
		Signature     string
		ImcRomBase    string
		GecRomBase    string
		XHCRomBase    string
		PSPDirBase    string
		NewPSPDirBase string
		BHDDirBase    string
		NewBHDDirBase string
	}

	AMDRom struct {
		Type        amdfw.RomType
		Directories []AMDDirectory
	}

	AMDDirectory struct {
		Header   AMDDirectoryHeader
		Entries  []AMDEntry
		Location string
	}

	AMDDirectoryHeader struct {
		Cookie        string
		Checksum      string
		ChecksumValid bool
		TotalEntries  string
		Reserved      string
	}

	AMDDirectoryEntry struct {
		Type     string
		Size     string
		Location string
		Reserved string
		Unknown  string
	}

	AMDEntry struct {
		DBEntry
		DirectoryEntry AMDDirectoryEntry
	}
)
package MFTCommon

import "github.com/Mimoja/PSP-Entry-Types"

const AMDSignature = uint32(0x55AA55AA)
const AMDFLASHMAPPING = uint32(0xFF000000)

type AMDFirmware struct {
	FirmwareEntryTable *AMDFirmwareEntryTable
	FlashMapping       uint32

	AGESA []AMDAGESA

	PSPDir    []AMDPSPDirectory
	NewPSPDir []AMDPSPDirectory
}

type AMDFirmwareEntryTable struct {
	Signature     uint32
	ImcRomBase    *uint32
	GecRomBase    *uint32
	XHCRomBase    *uint32
	PSPDirBase    *uint32
	NewPSPDirBase *uint32
	BDHDirBase    *uint32
	Unknown1      *uint32
}

type AMDAGESA struct {
	Header string
	Raw    string
	Offset uint32
}

/**
 * PSP related constants
 */

const AMDPSPCOOCKIE = "$PSP"
const AMDDUALPSPCOOCKIE = "2PSP"
const AMDBHDCOOCKIE = "$BHD"
const AMDSECONDPSPCOOCKIE = "$PL2"
const AMDSECONDBHDCOOCKIE = "$BL2"

/**
 * PSP related data structures
 */

type AMDPSPDirectory struct {
	Header  *AMDPSPDirectoryHeader
	Entries []AMDPSPDirectoryEntry
}

type AMDPSPDirectoryHeader struct {
	PspCookie    [4]byte
	Checksum     uint32
	TotalEntries uint32
	Reserved     uint32
}

type AMDPSPDirectoryEntry struct {
	ID       IDEntry
	Raw      *AMDPSPDirectoryBinaryEntry
	Header   *AMDPSPEntryHeader
	TypeInfo *pspentrytype.AMDPSPDirectoryEntryType
}

type AMDPSPDirectoryBinaryEntry struct {
	Type     uint32
	Size     uint32
	Location uint32
	Reserved uint32
}

type AMDPSPEntryBinaryHeader struct {
	Unknown1       [16]byte
	ID             uint32
	SizeSigned     uint32
	Unknown2       [0x18]byte
	UnknownType    uint32
	Unknown3       [4]byte
	SigFingerprint [16]byte
	IsCompressed   uint32
	Unknown4       uint32
	FullSize       uint32
	Unknown5       [12]byte
	Version        [4]byte
	Unknown6       [4]byte
	Unknown7       [4]byte
	SizePacked     uint32
}

type AMDPSPEntryHeader struct {
	ID             IDEntry
	ContentID      IDEntry
	Raw            AMDPSPEntryBinaryHeader
	Ident          uint32
	SizeSigned     uint32
	SigFingerprint string
	IsCompressed   bool
	FullSize       uint32
	Version        string
	SizePacked     uint32
}

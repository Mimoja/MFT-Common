package MFTCommon

const AMDSignature = uint32(0x55AA55AA)
const AMDFLASHMAPPING = uint32(0xFF000000)
const AMDPSPCOOCKIE = uint32(0x50535024)

type AMDFirmware struct {
	AMDFlashDescriptor AMDFirmwareEntryTable
	AMDPSP *AMDPSPDirectory
}

const  (
	PSPDirectoryEntryAMD_PUBLIC_KEY                  = iota
	PSPDirectoryEntryPSP_FW_BOOT_LOADER
	PSPDirectoryEntryPSP_FW_TRUSTED_OS
	PSPDirectoryEntryPSP_FW_RECOVERY_BOOT_LOADER
	PSPDirectoryEntryPSP_NV_DATA
	PSPDirectoryEntryBIOS_PUBLIC_KEY
	PSPDirectoryEntryBIOS_RTM_FIRMWARE
	PSPDirectoryEntryBIOS_RTM_SIGNATURE
	AMDPSPDirectoryEntrySMU_OFFCHIP_FW
)

type AMDFirmwareEntryTable struct {
	Signature uint32
	ImcRomBase uint32
	GecRomBase uint32
	XHCRomBase uint32
	PspDirBase uint32
}

type AMDPSPDirectory struct {
	Header AMDPSPDirectoryHeader
	Entries []AMDPSPDirectoryEntry
}

type AMDPSPDirectoryHeader struct {
  PspCookie uint32
  Checksum uint32
  TotalEntries uint32
  Reserved uint32
}

type AMDPSPDirectoryEntry struct {
	Type uint32
	Size uint32
	Location uint64
}


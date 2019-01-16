package MFTCommon

const AMDSignature = uint32(0x55AA55AA)
const AMDFLASHMAPPING = uint32(0xFF000000)

type AMDFirmware struct {
	FirmwareEntryTable AMDFirmwareEntryTable
	FlashMapping    uint32

	PSPDir             *AMDPSPDirectory
	NewPSPDir             *AMDPSPDirectory
}

type AMDFirmwareEntryTable struct {
	Signature     uint32
	ImcRomBase    uint32
	GecRomBase    uint32
	XHCRomBase    uint32
	PSPDirBase    uint32
	NewPSPDirBase uint32
	BDHDirBase    uint32
	Unknown1      uint32
}

/**
 * PSP related constants
 */

const AMDPSPCOOCKIE = "$PSP"
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
	ID  IDEntry
	Raw *AMDPSPDirectoryBinaryEntry
	Header *AMDPSPEntryHeader
	TypeInfo *AMDPSPDirectoryEntryType
}

type AMDPSPDirectoryBinaryEntry struct {
	Type     uint32
	Size     uint32
	Location uint64
}


type AMDPSPDirectoryEntryType struct {
	Type    uint32
	Name    string
	Comment string
}

type AMDPSPEntryBinaryHeader struct {
	Unknown1 [16]byte
	ID uint32
	SizeSigned uint32
	Unknown2 [0x18]byte
	AlwaysOne uint32 // Could be a type
	Unknown3 [4]byte
	SigFingerprint [16]byte
	IsCompressed uint32
	Unknown4 uint32
	FullSize uint32
	Unknown5 [12]byte
	Version [4] byte
	Unknown6 [4]byte
	Unknown7 [4]byte
	SizePacked uint32
}

type AMDPSPEntryHeader struct {
	ID IDEntry
	ContentID IDEntry
	Raw AMDPSPEntryBinaryHeader
	Ident uint32
	SizeSigned uint32
	SigFingerprint string
	IsCompressed bool
	FullSize uint32
	Version string
	SizePacked uint32
}

var AMDPSPDirectoryEntries = []AMDPSPDirectoryEntryType{

	{0x00, "AMD_PUBLIC_KEY", "AMD public key"},
	{0x01, "PSP_FW_BOOT_LOADER", "PSP boot loader in SPI space"},
	{0x02, "PSP_FW_TRUSTED_OS", "PSP Firmware region in SPI space"},
	{0x03, "PSP_FW_RECOVERY_BOOT_LOADER", "PSP recovery region"},
	{0x04, "PSP_NV_DATA", "PSP data region in SPI space"},
	{0x05, "BIOS_PUBLIC_KEY", "BIOS public key stored in SPI space"},
	{0x06, "BIOS_RTM_FIRMWARE", "BIOS RTM code (PEI volume) in SPI space"},
	{0x07, "BIOS_RTM_SIGNATURE", "Signed BIOS RTM hash stored  in SPI space"},
	{0x08, "SMU_OFFCHIP_FW", "SMU image"},
	{0x09, "AMD_SEC_DBG_PUBLIC_KEY", "Secure Unlock Public key"},
	{0x0A, "OEM_PSP_FW_PUBLIC_KEY", "Optional public part of the OEM PSP Firmware"},
	{0x0B, "AMD_SOFT_FUSE_CHAIN_01", "64bit PSP Soft Fuse Chain"},
	{0x0C, "PSP_BOOT_TIME_TRUSTLETS", "Boot-loaded trustlet binaries"},
	{0x0D, "PSP_BOOT_TIME_TRUSTLETS_KEY", "Key of the boot-loaded trustlet binaries"},
	{0x10, "PSP_AGESA_RESUME_FW", "PSP Agsa-Resume-Firmware"},
	{0x12, "SMU_OFF_CHIP_FW_2", "Secondary SMU image"},
	{0x14, "!PSP_MCLF_TRUSTLETS", "very similiar to ~PspTrustlets.bin~ in coreboot blobs"},
	{0x1A, "PSP_S3_NV_DATA", "S3 Data Blob"},
	{0x31, "{0x31~ABL_ARM_CODE~", "a _lot_ of strings and also some ARM code"},
	{0x38, "!PSP_ENCRYPTED_NV_DATA", ""},
	{0x39, "!SEV_APP", ""},
	{0x40, "!PL2_SECONDARY_DIRECTORY", ""},
	{0x5f, "FW_PSP_SMUSCS", "Software Configuration Settings Data Block"},
	{0x60, "FW_IMC", ""},
	{0x61, "FW_GEC", ""},
	{0x62, "FW_XHCI", ""},
	{0x63, "FW_INVALID", ""},
	{0x70, "!BL2_SECONDARY_DIRECTORY", ""},
	{0x108, "PSP_SMU_FN_FIRMWARE", ""},
	{0x112, "!SMU_OFF_CHIP_FW_3", "seems to tbe a tertiary SMU image"},
	{0x118, "PSP_SMU_FN_FIRMWARE2", ""},
	{0x15f, "!FW_PSP_SMUSCS_2", "seems to be a secondary FW_PSP_SMUSCS"},
	{0x30062, "{0x30062~UEFI-IMAGE~", ""},
}


package MFTCommon

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"github.com/glaslos/ssdeep"
	"golang.org/x/crypto/sha3"
	"io"
	"log"
)

type DownloadEntry struct {
	Vendor           string  `json:",omitempty"`
	Product          string  `json:",omitempty"`
	Version          string  `json:",omitempty"`
	Title            string  `json:",omitempty"`
	Description      string  `json:",omitempty"`
	ReleaseDate      string  `json:",omitempty"`
	DownloadFileSize string  `json:",omitempty"`
	DownloadURL      string  `json:",omitempty"`
	DownloadPath     string  `json:",omitempty"`
	DownloadTime     string  `json:",omitempty" hash:"-"`
	PackageID        IDEntry `json:",omitempty" hash:"-"`
}

type ImportEntry struct {
	ImportDataDefinition string `json:",omitempty"`
	MetaData             DownloadEntry   `json:",omitempty"`
	Contents             []StorageEntry  `json:",omitempty"`
	ImportTime           string          `json:",omitempty"`
	Success              bool            `json:",omitempty"`
}
type StorageEntry struct {
	ID        IDEntry  `json:",omitempty"`
	PackageID IDEntry  `json:",omitempty"`
	Path      string   `json:",omitempty"`
	Tags      []string `json:",omitempty"`
}

type FlashImage struct {
	FlashimageDataDefinition string `json:",omitempty"`
	MetaData       DownloadEntry  `json:",omitempty"`
	ID             IDEntry        `json:",omitempty"`
	Tags           []string       `json:",omitempty"`
	FirmwareOffset int64          `json:",omitempty"`
	AMD            *AMDFirmware   `json:"AMD"`
	INTEL          *IntelFirmware `json:"INTEL""`
	Certificates   []string       `json:"Certificates"`
}

type IntelFirmware struct {
	IFD *IntelFlashDescriptor `json:"IFD"`
}

type IDEntry struct {
	SSDEEP    string `json:",omitempty"`
	SHA3_512  string `json:",omitempty"`
	SHA512    string `json:",omitempty"`
	SHA256    string `json:",omitempty"`
	SHA1      string `json:",omitempty"`
	MD5       string `json:",omitempty"`
	Algorithm string `json:",omitempty"`
}

func (d IDEntry) GetID() string {
	return d.SHA256
}

func GenerateID(data []byte) IDEntry {
	ssdeepHash := ssdeep.NewSSDEEP()
	sha3Hash := sha3.New512()
	sha512Hash := sha512.New()
	sha256Hash := sha256.New()
	sha1Hash := sha1.New()
	md5Hash := md5.New()

	if _, err := io.Copy(sha3Hash, bytes.NewReader(data)); err != nil {
		log.Fatal(err)
	}

	if _, err := io.Copy(sha512Hash, bytes.NewReader(data)); err != nil {
		log.Fatal(err)
	}

	if _, err := io.Copy(sha256Hash, bytes.NewReader(data)); err != nil {
		log.Fatal(err)
	}

	if _, err := io.Copy(sha1Hash, bytes.NewReader(data)); err != nil {
		log.Fatal(err)
	}

	if _, err := io.Copy(md5Hash, bytes.NewReader(data)); err != nil {
		log.Fatal(err)
	}

	ssdString := "_UNKNOWN_"
	ssd, err := ssdeepHash.FuzzyByte(data)
	if err == nil {
		ssdString = ssd.String()
	}

	return IDEntry{
		SSDEEP:    ssdString,
		SHA3_512:  fmt.Sprintf("%X", sha3Hash.Sum(nil)),
		SHA512:    fmt.Sprintf("%X", sha512Hash.Sum(nil)),
		SHA256:    fmt.Sprintf("%X", sha256Hash.Sum(nil)),
		SHA1:      fmt.Sprintf("%X", sha1Hash.Sum(nil)),
		MD5:       fmt.Sprintf("%X", md5Hash.Sum(nil)),
		Algorithm: "sha256",
	}
}

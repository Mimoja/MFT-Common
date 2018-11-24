package MFTCommon

import (
	"io"
)

func ReadNextBytes(reader io.Reader, number int64) ([]byte, error) {
	bytes := make([]byte, number)

	_, err := reader.Read(bytes)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

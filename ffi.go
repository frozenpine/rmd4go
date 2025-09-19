package rmd4go

import (
	"bytes"
	"io"
	"log/slog"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func SetCString(buff []uint8, value string) {
	copy(buff[:len(buff)-1], []uint8(value))
}

func IsASCII(buff []byte) bool {
	for _, b := range buff {
		if b > 127 {
			return false
		}
	}

	return true
}

var decoder = simplifiedchinese.GB18030.NewDecoder()

func ReadCString(buff []uint8) string {
	idx := bytes.IndexByte(buff, 0)
	if idx <= 0 {
		return ""
	}

	if IsASCII(buff[:idx]) {
		return string(buff[:idx])
	}

	reader := transform.NewReader(bytes.NewReader(buff[:idx]), decoder)

	if decoded, err := io.ReadAll(reader); err != nil {
		slog.Error(
			"decode GB18030 failed",
			slog.Any("error", err),
			slog.Any("buff", buff),
		)
		return ""
	} else {
		return string(decoded)
	}
}

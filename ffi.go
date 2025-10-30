package rmd4go

/*
#include "helper.h"
*/
import "C"
import (
	"bytes"
	"io"
	"log/slog"
	"unsafe"

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
		idx = len(buff)
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

func WriteCString(dst unsafe.Pointer, value string, len C.size_t) {
	if value == "" {
		return
	}

	buffer := []byte(value)

	ptr := unsafe.Pointer(&buffer[0])

	C.memset(dst, 0, len)
	C.memcpy(dst, ptr, len-1)
}

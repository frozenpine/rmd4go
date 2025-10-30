NAME="testdemo"
VERSION=1.0.1
BUILD=`date +%FT%T%z`

CFLAGS=-I../../ -I../../dependencies/includes -ldl
EXTRA_FLAGS=

ifeq ($(OS), Windows_NT)
	BINARY=${NAME}.exe
	EXTRA_FLAGS=-D_WIN32
else
	BINARY=${NAME}
endif

default:
	@mkdir -p bin
	@gcc main.c ../../helper.c -o bin/${BINARY} ${CFLAGS} ${EXTRA_FLAGS}

clean:
	@if [ -f bin/${BINARY} ] ; then rm -f bin/${BINARY} ; fi

.PHONY: default clean
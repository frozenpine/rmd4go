package rmd4go

/*
#cgo CFLAGS: -I${SRCDIR} -I${SRCDIR}/dependencies/includes
#cgo LDFLAGS: -ldl

#include <stdbool.h>
#include <stdint.h>
#include <stdlib.h>
#include <strings.h>
#include <dlfcn.h>

#include "FtdcApiStruct.h"

typedef void (*Release)(void* this);
typedef void (*Init)(void* this);
typedef int (*Join)(void* this);
typedef bool (*ChkCurrentThread)(void* this);
typedef void (*RegisterNameServer)(void* this, const char *pszNsAddress);
typedef void (*RegisterSpi)(void* this, void *pSpi);
typedef int (*ReqUserLogin)(
	void* this,
	struct CRsaFtdcReqUserLoginField *pReqUserLoginField,
	int nRequestID);
typedef int (*ReqSubMarketData)(
	void* this,
	char *ppInstrumentID[],
	int nCount);
typedef int (*ReqUnSubMarketData)(
	void* this,
	char *ppInstrumentID[],
	int nCount);
typedef int (*ReqQryMarketData)(
	void* this,
	struct CRsaFtdcQryMarketDataField *pQryMarketData,
	int nRequestID);
typedef int (*ReqSendDepthMarketDataRead)(
	void* this,
	struct CRsaFtdcDepthMarketDataField *pDepthMarketData,
	int nRequestID);
typedef int (*ReqSendBarMarketDataRead)(
	void* this,
	struct CRsaFtdcBarMarketDataField *pBarMarketData,
	int nRequestID);
typedef int (*ReqBtSubMarketData)(
	void* this,
	struct CRsaFtdcBtSubMarketDataField *ppFiels[],
	int nCount);
typedef int (*ReqBtUnSubMarketData)(
	void* this,
	struct CRsaFtdcBtSubMarketDataField *ppFiels[],
	int nCount);
typedef int (*ReqSubMarketDataCompleted)(void* this);
typedef int (*ReqSubCombMarketData)(
	void* this,
	struct CRsaFtdcSubCombMarketDataField* ppFiels[],
	int nCount);
typedef int (*ReqQryBarMarketData)(
	void* this,
	struct CRsaFtdcBtSubMarketDataField* subMd,
	struct CRsaFtdcBarMarketDataField* bars[],
	int* count);
typedef int (*ReqQryDepthMarketData)(
	void* this,
	struct CRsaFtdcQryMarketDataField* pQryMarketData,
	struct CRsaFtdcDepthMarketDataField* depthMd[],
	int* count);

typedef struct {
	Release CFtdcMdApiVtable_Release;
	Init CFtdcMdApiVtable_Init;
	Join CFtdcMdApiVtable_Join;
	ChkCurrentThread CFtdcMdApiVtable_ChkCurrentThread;
	RegisterNameServer CFtdcMdApiVtable_RegisterNameServer;
	RegisterSpi CFtdcMdApiVtable_RegisterSpi;
	ReqUserLogin CFtdcMdApiVtable_ReqUserLogin;
	ReqSubMarketData CFtdcMdApiVtable_ReqSubMarketData;
	ReqUnSubMarketData CFtdcMdApiVtable_ReqUnSubMarketData;
	ReqQryMarketData CFtdcMdApiVtable_ReqQryMarketData;
	ReqSendDepthMarketDataRead CFtdcMdApiVtable_ReqSendDepthMarketDataRead;
	ReqSendBarMarketDataRead CFtdcMdApiVtable_ReqSendBarMarketDataRead;
	ReqBtSubMarketData CFtdcMdApiVtable_ReqBtSubMarketData;
	ReqBtUnSubMarketData CFtdcMdApiVtable_ReqBtUnSubMarketData;
	ReqSubMarketDataCompleted CFtdcMdApiVtable_ReqSubMarketDataCompleted;
	ReqSubCombMarketData CFtdcMdApiVtable_ReqSubCombMarketData;
	ReqQryBarMarketData CFtdcMdApiVtable_ReqQryBarMarketData;
	ReqQryDepthMarketData CFtdcMdApiVtable_ReqQryDepthMarketData;
} CFtdcMdApiVtable;

typedef struct {
	CFtdcMdApiVtable vtable;
} CFtdcMdApiExt;

typedef void* (*CreateFtdcMdApi)(
	const char *pszFrontAddress,
	const char *pszFlowPath,
	const char *pszUserApiType,
	bool needHistoryTick);

void* CallCreateFtdcMdApi(
	CreateFtdcMdApi fn,
	const char *pszFrontAddress,
	const char *pszFlowPath,
	const char *pszUserApiType,
	bool needHistoryTick
) {
	return fn(pszFrontAddress, pszFlowPath, pszUserApiType, needHistoryTick);
}

void CallRelease(Release fn, void* this) { return fn(this); }

void CallInit(Init fn, void* this) { return fn(this); }

int CallJoin(Join fn, void* this) { return fn(this); }

bool CallChkCurrentThread(ChkCurrentThread fn, void* this) { return fn(this); }

void CallRegisterNameServer(
	RegisterNameServer fn, void* this,
	const char* pszNsAddress
) { return fn(this, pszNsAddress); }

void CallRegisterSpi(
	RegisterSpi fn, void* this,
	void* pSpi
) { return fn(this, pSpi); }

int CallReqUserLogin(
	ReqUserLogin fn, void* this,
	struct CRsaFtdcReqUserLoginField *pReqUserLoginField,
	int nRequestID
) { return fn(this, pReqUserLoginField, nRequestID); }

int CallReqSubMarketData(
	ReqSubMarketData fn, void* this,
	char *ppInstrumentID[],
	int nCount
) { return fn(this, ppInstrumentID, nCount); }

int CallReqUnSubMarketData(
	ReqUnSubMarketData fn, void* this,
	char *ppInstrumentID[],
	int nCount
) { return fn(this, ppInstrumentID, nCount); }

int CallReqQryMarketData(
	ReqQryMarketData fn, void* this,
	struct CRsaFtdcQryMarketDataField *pQryMarketData,
	int nRequestID
) { return fn(this, pQryMarketData, nRequestID); }

int CallReqSendDepthMarketDataRead(
	ReqSendDepthMarketDataRead fn, void* this,
	struct CRsaFtdcDepthMarketDataField *pDepthMarketData,
	int nRequestID
) { return fn(this, pDepthMarketData, nRequestID); }

int CallReqSendBarMarketDataRead(
	ReqSendBarMarketDataRead fn, void* this,
	struct CRsaFtdcBarMarketDataField *pBarMarketData,
	int nRequestID
) { return fn(this, pBarMarketData, nRequestID); }

int CallReqBtSubMarketData(
	ReqBtSubMarketData fn, void* this,
	struct CRsaFtdcBtSubMarketDataField *ppFiels[],
	int nCount
) { return fn(this, ppFiels, nCount); }

int CallReqBtUnSubMarketData(
	ReqBtUnSubMarketData fn, void* this,
	struct CRsaFtdcBtSubMarketDataField *ppFiels[],
	int nCount
) { return fn(this, ppFiels, nCount); }

int CallReqSubMarketDataCompleted(
	ReqSubMarketDataCompleted fn, void* this
) { return fn(this); }

int CallReqSubCombMarketData(
	ReqSubCombMarketData fn, void* this,
	struct CRsaFtdcSubCombMarketDataField* ppFiels[],
	int nCount
) { return fn(this, ppFiels, nCount); }

int CallReqQryBarMarketData(
	ReqQryBarMarketData fn, void* this,
	struct CRsaFtdcBtSubMarketDataField* subMd,
	struct CRsaFtdcBarMarketDataField* bars[],
	int* count
) { return fn(this, subMd, bars, count); }

int CallReqQryDepthMarketData(
	ReqQryDepthMarketData fn, void* this,
	struct CRsaFtdcQryMarketDataField* pQryMarketData,
	struct CRsaFtdcDepthMarketDataField* depthMd[],
	int* count
) { return fn(this, pQryMarketData, depthMd, count); }
*/
import "C"

import (
	"errors"
	"fmt"
	"log/slog"
	"runtime"
	"sync"
	"sync/atomic"
	"unsafe"
)

type reqRtn int

func (rtn reqRtn) Error() string {
	return fmt.Sprintf("req execution failed: %d", rtn)
}

func Rtn(rtn C.int) error {
	if rtn == 0 {
		return nil
	}

	return reqRtn(rtn)
}

type CFtdcMdApi struct {
	apiPtr *C.CFtdcMdApiExt

	spiPtr       cFtdcMdSpi
	registerOnce sync.Once

	reqID atomic.Int32
}

const (
	CREATE_RMD_API_WIN   = "?CreateFtdcMdApi@CFtdcMdApi@@SAPEAV1@PEBD00_N@Z"
	CREATE_RMD_API_LINUX = "_ZN10CFtdcMdApi15CreateFtdcMdApiEPKcS1_S1_b"
)

var (
	errLibOpenFialed     = errors.New("lib open failed")
	errLibSymbolNotFound = errors.New("symbol not found")

	ErrInvalidReqData = errors.New("invalid req data")
)

func CreateFtdcMdApi(
	libPath string,
	frontAddr, flowPath, userApiType string,
	needHisTick bool,
) (api *CFtdcMdApi, err error) {
	lib_path := C.CString(libPath)
	defer C.free(unsafe.Pointer(&lib_path))

	lib := C.dlopen(lib_path, C.RTLD_LAZY)
	if lib == nil {
		msg := C.dlerror()

		err = fmt.Errorf(
			"%w: %s", errLibOpenFialed, C.GoString(msg),
		)
		return
	}

	var creatorName *C.char

	switch runtime.GOOS {
	case "linux":
		creatorName = C.CString(CREATE_RMD_API_LINUX)
	case "windows":
		creatorName = C.CString(CREATE_RMD_API_WIN)
	default:
		return nil, errors.New("unsupported platform")
	}
	defer C.free(unsafe.Pointer(creatorName))

	if creator := C.dlsym(
		lib, creatorName,
	); creator == nil {
		msg := C.dlerror()

		err = fmt.Errorf(
			"%w: %s", errLibSymbolNotFound, C.GoString(msg),
		)
		return
	} else {
		pszFrontAddr := C.CString(frontAddr)
		pszFlowPath := C.CString(flowPath)
		pszUserApiType := C.CString(userApiType)
		defer func() {
			C.free(unsafe.Pointer(pszFrontAddr))
			C.free(unsafe.Pointer(pszFlowPath))
			C.free(unsafe.Pointer(pszUserApiType))
		}()

		if instance := C.CallCreateFtdcMdApi(
			C.CreateFtdcMdApi(creator),
			pszFrontAddr, pszFlowPath, pszUserApiType,
			C.bool(needHisTick),
		); instance == nil {
			return nil, errors.New("create rmd api failed")
		} else {
			api = &CFtdcMdApi{
				apiPtr: (*C.CFtdcMdApiExt)(instance),
			}

			runtime.SetFinalizer(api, func(ins *CFtdcMdApi) {
				ins.Release()
				ins.spiPtr.callback = nil
			})
		}
	}

	return
}

func (api *CFtdcMdApi) Release() {
	C.CallRelease(
		api.apiPtr.vtable.CFtdcMdApiVtable_Release,
		unsafe.Pointer(api.apiPtr),
	)
}

func (api *CFtdcMdApi) Init() {
	C.CallInit(
		api.apiPtr.vtable.CFtdcMdApiVtable_Init,
		unsafe.Pointer(api.apiPtr),
	)
}

func (api *CFtdcMdApi) Join() error {
	return Rtn(C.CallJoin(
		api.apiPtr.vtable.CFtdcMdApiVtable_Join,
		unsafe.Pointer(api.apiPtr),
	))
}

func (api *CFtdcMdApi) ChkCurrentThread() bool {
	return bool(C.CallChkCurrentThread(
		api.apiPtr.vtable.CFtdcMdApiVtable_ChkCurrentThread,
		unsafe.Pointer(api.apiPtr),
	))
}

func (api *CFtdcMdApi) RegisterNameServer(nsAddress string) {
	pszNsAddress := C.CString(nsAddress)
	defer C.free(unsafe.Pointer(pszNsAddress))

	C.CallRegisterNameServer(
		api.apiPtr.vtable.CFtdcMdApiVtable_RegisterNameServer,
		unsafe.Pointer(api.apiPtr),
		pszNsAddress,
	)
}

func (api *CFtdcMdApi) RegisterSpi(spi RmdSpi) {
	if spi == nil {
		slog.Warn("spi interface is nil")
		return
	}

	api.registerOnce.Do(func() {
		api.spiPtr.vtable = callbackVTable
		api.spiPtr.callback = spi

		C.CallRegisterSpi(
			api.apiPtr.vtable.CFtdcMdApiVtable_RegisterSpi,
			unsafe.Pointer(api.apiPtr),
			unsafe.Pointer(&api.spiPtr),
		)
	})
}

func (api *CFtdcMdApi) ReqUserLogin(login *CRsaFtdcReqUserLoginField) error {
	if login == nil {
		return fmt.Errorf(
			"%w: login info is nil", ErrInvalidReqData,
		)
	}

	return Rtn(C.CallReqUserLogin(
		api.apiPtr.vtable.CFtdcMdApiVtable_ReqUserLogin,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CRsaFtdcReqUserLoginField)(unsafe.Pointer(login)),
		C.int(api.reqID.Add(1)),
	))
}

func (api *CFtdcMdApi) ReqSubMarketData(instruments ...string) error {
	if len(instruments) <= 0 {
		slog.Warn("sub instruments is empty")

		return nil
	}

	ppInstruments := make([]*C.char, len(instruments))
	for idx, ins := range instruments {
		ppInstruments[idx] = C.CString(ins)
	}
	defer func() {
		for _, ptr := range ppInstruments {
			C.free(unsafe.Pointer(ptr))
		}
	}()

	return Rtn(C.CallReqSubMarketData(
		api.apiPtr.vtable.CFtdcMdApiVtable_ReqSubMarketData,
		unsafe.Pointer(api.apiPtr),
		(**C.char)(unsafe.Pointer(ppInstruments[0])),
		C.int(len(ppInstruments)),
	))
}

func (api *CFtdcMdApi) ReqUnSubMarketData(instruments ...string) error {
	if len(instruments) <= 0 {
		slog.Warn("un-sub instruments is empty")

		return nil
	}

	ppInstruments := make([]*C.char, len(instruments))
	for idx, ins := range instruments {
		ppInstruments[idx] = C.CString(ins)
	}
	defer func() {
		for _, ptr := range ppInstruments {
			C.free(unsafe.Pointer(ptr))
		}
	}()

	return Rtn(C.CallReqUnSubMarketData(
		api.apiPtr.vtable.CFtdcMdApiVtable_ReqUnSubMarketData,
		unsafe.Pointer(api.apiPtr),
		(**C.char)(unsafe.Pointer(ppInstruments[0])),
		C.int(len(ppInstruments)),
	))
}

func (api *CFtdcMdApi) ReqQryMarketData(qry *CRsaFtdcQryMarketDataField) error {
	if qry == nil {
		return fmt.Errorf(
			"%w: query market data info is nil", ErrInvalidReqData,
		)
	}

	return Rtn(C.CallReqQryMarketData(
		api.apiPtr.vtable.CFtdcMdApiVtable_ReqQryMarketData,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CRsaFtdcQryMarketDataField)(unsafe.Pointer(qry)),
		C.int(api.reqID.Add(1)),
	))
}

func (api *CFtdcMdApi) ReqSendDepthMarketDataRead(
	depth *CRsaFtdcDepthMarketDataField,
) error {
	if depth == nil {
		return fmt.Errorf(
			"%w: requst send depth market data is nil", ErrInvalidReqData,
		)
	}

	return Rtn(C.CallReqSendDepthMarketDataRead(
		api.apiPtr.vtable.CFtdcMdApiVtable_ReqSendDepthMarketDataRead,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CRsaFtdcDepthMarketDataField)(unsafe.Pointer(depth)),
		C.int(api.reqID.Add(1)),
	))
}

func (api *CFtdcMdApi) ReqSendBarMarketDataRead(
	bar *CRsaFtdcBarMarketDataField,
) error {
	return Rtn(C.CallReqSendBarMarketDataRead(
		api.apiPtr.vtable.CFtdcMdApiVtable_ReqSendBarMarketDataRead,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CRsaFtdcBarMarketDataField)(unsafe.Pointer(bar)),
		C.int(api.reqID.Add(1)),
	))
}

func (api *CFtdcMdApi) ReqBtSubMarketData(
	subs ...*CRsaFtdcBtSubMarketDataField,
) error {
	if len(subs) <= 0 {
		return nil
	}

	ppFields := make([]*C.struct_CRsaFtdcBtSubMarketDataField, len(subs))
	for idx, sub := range subs {
		ppFields[idx] = (*C.struct_CRsaFtdcBtSubMarketDataField)(
			unsafe.Pointer(sub))
	}

	return Rtn(C.CallReqBtSubMarketData(
		api.apiPtr.vtable.CFtdcMdApiVtable_ReqBtSubMarketData,
		unsafe.Pointer(api.apiPtr),
		(**C.struct_CRsaFtdcBtSubMarketDataField)(unsafe.Pointer(&ppFields[0])),
		C.int(len(subs)),
	))
}

func (api *CFtdcMdApi) ReqBtUnSubMarketData(
	unSubs ...*CRsaFtdcBtSubMarketDataField,
) error {
	if len(unSubs) <= 0 {
		return nil
	}

	ppFields := make([]*C.struct_CRsaFtdcBtSubMarketDataField, len(unSubs))
	for idx, sub := range unSubs {
		ppFields[idx] = (*C.struct_CRsaFtdcBtSubMarketDataField)(
			unsafe.Pointer(sub))
	}

	return Rtn(C.CallReqBtUnSubMarketData(
		api.apiPtr.vtable.CFtdcMdApiVtable_ReqBtUnSubMarketData,
		unsafe.Pointer(api.apiPtr),
		(**C.struct_CRsaFtdcBtSubMarketDataField)(unsafe.Pointer(&ppFields[0])),
		C.int(len(unSubs)),
	))
}

func (api *CFtdcMdApi) ReqSubMarketDataCompleted() error {
	return Rtn(C.CallReqSubMarketDataCompleted(
		api.apiPtr.vtable.CFtdcMdApiVtable_ReqSubMarketDataCompleted,
		unsafe.Pointer(api.apiPtr),
	))
}

func (api *CFtdcMdApi) ReqSubCombMarketData(
	subs ...*CRsaFtdcSubCombMarketDataField,
) error {
	if len(subs) <= 0 {
		return nil
	}

	ppFields := make([]*C.struct_CRsaFtdcSubCombMarketDataField, len(subs))
	for idx, sub := range subs {
		ppFields[idx] = (*C.struct_CRsaFtdcSubCombMarketDataField)(
			unsafe.Pointer(sub))
	}

	return Rtn(C.CallReqSubCombMarketData(
		api.apiPtr.vtable.CFtdcMdApiVtable_ReqSubCombMarketData,
		unsafe.Pointer(api.apiPtr),
		(**C.struct_CRsaFtdcSubCombMarketDataField)(
			unsafe.Pointer(&ppFields[0])),
		C.int(len(subs)),
	))
}

func (api *CFtdcMdApi) ReqQryBarMarketData(
	sub *CRsaFtdcBtSubMarketDataField, count int,
) ([]*CRsaFtdcBarMarketDataField, error) {
	if sub == nil {
		return nil, fmt.Errorf(
			"%w: query bar data is nil", ErrInvalidReqData,
		)
	}

	if count <= 0 {
		return nil, fmt.Errorf(
			"%w: query bar count is empty", ErrInvalidReqData,
		)
	}

	data := make([]*C.struct_CRsaFtdcBarMarketDataField, count)
	outCount := C.int(count)

	if err := Rtn(C.CallReqQryBarMarketData(
		api.apiPtr.vtable.CFtdcMdApiVtable_ReqQryBarMarketData,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CRsaFtdcBtSubMarketDataField)(unsafe.Pointer(sub)),
		(**C.struct_CRsaFtdcBarMarketDataField)(unsafe.Pointer(&data[0])),
		&outCount,
	)); err != nil {
		return nil, err
	} else {
		bars := make([]*CRsaFtdcBarMarketDataField, outCount)

		for idx, bar := range data[:outCount] {
			bars[idx] = (*CRsaFtdcBarMarketDataField)(unsafe.Pointer(bar))
		}

		return bars, nil
	}
}

func (api *CFtdcMdApi) ReqQryDepthMarketData(
	sub *CRsaFtdcQryMarketDataField, count int,
) ([]*CRsaFtdcDepthMarketDataField, error) {
	if sub == nil {
		return nil, fmt.Errorf(
			"%w: query bar data is nil", ErrInvalidReqData,
		)
	}

	if count <= 0 {
		return nil, fmt.Errorf(
			"%w: query bar count is empty", ErrInvalidReqData,
		)
	}

	data := make([]*C.struct_CRsaFtdcDepthMarketDataField, count)
	outCount := C.int(count)

	if err := Rtn(C.CallReqQryDepthMarketData(
		api.apiPtr.vtable.CFtdcMdApiVtable_ReqQryDepthMarketData,
		unsafe.Pointer(api.apiPtr),
		(*C.struct_CRsaFtdcQryMarketDataField)(unsafe.Pointer(sub)),
		(**C.struct_CRsaFtdcDepthMarketDataField)(unsafe.Pointer(&data[0])),
		&outCount,
	)); err != nil {
		return nil, err
	} else {
		ticks := make([]*CRsaFtdcDepthMarketDataField, outCount)

		for idx, tick := range data[:outCount] {
			ticks[idx] = (*CRsaFtdcDepthMarketDataField)(unsafe.Pointer(tick))
		}

		return ticks, nil
	}
}

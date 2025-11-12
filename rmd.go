package rmd4go

/*
#cgo CFLAGS: -I${SRCDIR} -I${SRCDIR}/dependencies/includes
#cgo LDFLAGS: -ldl

#include "helper.h"
*/
import "C"

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"path/filepath"
	"runtime"
	"sync"
	"sync/atomic"
	"unsafe"
)

var (
	spiCVtablePtr *C.CFtdcMdSpiVtable
)

type RmdSpi interface {
	OnMdFrontConnected()
	OnMdFrontDisconnected(reason int)
	OnMdRspUserLogin(
		login *CRsaFtdcRspUserLoginField,
		info *CRsaFtdcRspInfoField,
		requestID int, isLast bool,
	)
	OnRtnDepthMarketData(tick *CRsaFtdcDepthMarketDataField)
	OnRspSubMarketData(
		instrument *CRsaFtdcSpecificInstrumentField,
		info *CRsaFtdcRspInfoField,
		requestID int, isLast bool,
	)
	OnRspUnSubMarketData(
		instrument *CRsaFtdcSpecificInstrumentField,
		info *CRsaFtdcRspInfoField,
		requestID int, isLast bool,
	)
	OnRtnBarMarketData(bar *CRsaFtdcBarMarketDataField)
	OnRspQryMarketData(
		md *CRsaFtdcRspMarketDataField,
		info *CRsaFtdcRspInfoField,
		requestID int, isLast bool,
	)
	OnRtnMarketDataEnd(dt *CRsaFtdcNtfMarketDataEndField)
}

type goFtdcMdSpi struct {
	pinner   runtime.Pinner
	callback RmdSpi
}

//export CgoOnMdFrontConnected
func CgoOnMdFrontConnected(this unsafe.Pointer) {
	slog.Log(
		context.Background(), slog.LevelDebug-1,
		"rmd CgoOnMdFrontConnected called",
		slog.Any("spi", this),
	)

	(*goFtdcMdSpi)(
		(*C.CFtdcMdSpiExt)(this).spi,
	).callback.OnMdFrontConnected()
}

//export CgoOnMdFrontDisconnected
func CgoOnMdFrontDisconnected(this unsafe.Pointer, nReason C.int) {
	slog.Log(
		context.Background(), slog.LevelDebug-1,
		"rmd CgoOnMdFrontDisconnected called",
		slog.Any("spi", this),
		slog.Any("reason", nReason),
	)

	(*goFtdcMdSpi)(
		(*C.CFtdcMdSpiExt)(this).spi,
	).callback.OnMdFrontDisconnected(int(nReason))
}

//export CgoOnMdRspUserLogin
func CgoOnMdRspUserLogin(
	this unsafe.Pointer,
	pRspUserLogin *C.struct_CRsaFtdcRspUserLoginField,
	pRspInfo *C.struct_CRsaFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-1,
		"rmd CgoOnMdRspUserLogin called",
		slog.Any("spi", this),
		slog.Any("login", pRspUserLogin),
		slog.Any("info", pRspInfo),
		slog.Int("req", int(nRequestID)),
		slog.Bool("last", bool(bIsLast)),
	)
	(*goFtdcMdSpi)(
		(*C.CFtdcMdSpiExt)(this).spi,
	).callback.OnMdRspUserLogin(
		(&CRsaFtdcRspUserLoginField{}).FromCStruct(pRspUserLogin),
		(&CRsaFtdcRspInfoField{}).FromCStruct(pRspInfo),
		int(nRequestID), bool(bIsLast),
	)
}

var (
	tickPool = sync.Pool{New: func() any { return new(CRsaFtdcDepthMarketDataField) }}
)

func GetTick() *CRsaFtdcDepthMarketDataField {
	count := 0
	for {
		if data := tickPool.Get(); data != nil {
			runtime.SetFinalizer(data, func(v any) {
				runtime.SetFinalizer(v, nil)
				tickPool.Put(v)
			})

			return data.(*CRsaFtdcDepthMarketDataField)
		} else {
			count++

			if count <= 10 {
				slog.Warn(
					"get tick from pool failed",
					slog.Int("fail_count", count),
				)
			} else {
				slog.Error(
					"get tick attemps exceeded",
					slog.Int("fail_count", count),
				)

				panic("get tick attemps exceeded")
			}
		}
	}
}

//export CgoOnRtnDepthMarketData
func CgoOnRtnDepthMarketData(
	this unsafe.Pointer,
	pDepthMarketData *C.struct_CRsaFtdcDepthMarketDataField,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"rmd CgoOnRtnDepthMarketData called",
		slog.Any("spi", this),
		slog.Any("depth", pDepthMarketData),
	)
	(*goFtdcMdSpi)(
		(*C.CFtdcMdSpiExt)(this).spi,
	).callback.OnRtnDepthMarketData(
		GetTick().FromCStruct(pDepthMarketData),
	)
}

//export CgoOnRspSubMarketData
func CgoOnRspSubMarketData(
	this unsafe.Pointer,
	pSpecificInstrument *C.struct_CRsaFtdcSpecificInstrumentField,
	pRspInfo *C.struct_CRsaFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-1,
		"rmd CgoOnRspSubMarketData called",
		slog.Any("spi", this),
		slog.Any("ins", pSpecificInstrument),
		slog.Any("info", pRspInfo),
		slog.Int("req", int(nRequestID)),
		slog.Bool("last", bool(bIsLast)),
	)
	(*goFtdcMdSpi)(
		(*C.CFtdcMdSpiExt)(this).spi,
	).callback.OnRspSubMarketData(
		(&CRsaFtdcSpecificInstrumentField{}).FromCStruct(pSpecificInstrument),
		(&CRsaFtdcRspInfoField{}).FromCStruct(pRspInfo),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRspUnSubMarketData
func CgoOnRspUnSubMarketData(
	this unsafe.Pointer,
	pSpecificInstrument *C.struct_CRsaFtdcSpecificInstrumentField,
	pRspInfo *C.struct_CRsaFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-1,
		"rmd CgoOnRspUnSubMarketData called",
		slog.Any("spi", this),
		slog.Any("ins", pSpecificInstrument),
		slog.Any("info", pRspInfo),
		slog.Int("req", int(nRequestID)),
		slog.Bool("last", bool(bIsLast)),
	)
	(*goFtdcMdSpi)(
		(*C.CFtdcMdSpiExt)(this).spi,
	).callback.OnRspUnSubMarketData(
		(&CRsaFtdcSpecificInstrumentField{}).FromCStruct(pSpecificInstrument),
		(&CRsaFtdcRspInfoField{}).FromCStruct(pRspInfo),
		int(nRequestID), bool(bIsLast),
	)
}

var (
	barPool = sync.Pool{New: func() any {
		return new(CRsaFtdcBarMarketDataField)
	}}
)

func GetBar() *CRsaFtdcBarMarketDataField {
	count := 0
	for {
		if data := barPool.Get(); data != nil {
			runtime.SetFinalizer(data, func(v any) {
				runtime.SetFinalizer(v, nil)
				barPool.Put(v)
			})

			return data.(*CRsaFtdcBarMarketDataField)
		} else {
			count++

			if count <= 10 {
				slog.Warn(
					"get bar from pool failed",
					slog.Int("fail_count", count),
				)
			} else {
				slog.Error(
					"get bar attemps exceeded",
					slog.Int("fail_count", count),
				)

				panic("get bar attemps exceeded")
			}
		}
	}
}

//export CgoOnRtnBarMarketData
func CgoOnRtnBarMarketData(
	this unsafe.Pointer,
	pBarMarketData *C.struct_CRsaFtdcBarMarketDataField,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-2,
		"rmd CgoOnRtnBarMarketData called",
		slog.Any("spi", this),
		slog.Any("bar", pBarMarketData),
	)
	(*goFtdcMdSpi)(
		(*C.CFtdcMdSpiExt)(this).spi,
	).callback.OnRtnBarMarketData(
		GetBar().FromCStruct(pBarMarketData),
	)
}

//export CgoOnRspQryMarketData
func CgoOnRspQryMarketData(
	this unsafe.Pointer,
	pRspMarketData *C.struct_CRsaFtdcRspMarketDataField,
	pRspInfo *C.struct_CRsaFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-1,
		"rmd CgoOnRspQryMarketData called",
		slog.Any("spi", this),
		slog.Any("md", pRspMarketData),
		slog.Any("info", pRspInfo),
		slog.Int("req", int(nRequestID)),
		slog.Bool("last", bool(bIsLast)),
	)
	(*goFtdcMdSpi)(
		(*C.CFtdcMdSpiExt)(this).spi,
	).callback.OnRspQryMarketData(
		(&CRsaFtdcRspMarketDataField{}).FromCStruct(pRspMarketData),
		(&CRsaFtdcRspInfoField{}).FromCStruct(pRspInfo),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRtnMarketDataEnd
func CgoOnRtnMarketDataEnd(
	this unsafe.Pointer,
	pMarketDataEnd *C.struct_CRsaFtdcNtfMarketDataEndField,
) {
	slog.Log(
		context.Background(), slog.LevelDebug-1,
		"rmd CgoOnRtnMarketDataEnd called",
		slog.Any("spi", this),
		slog.Any("end", pMarketDataEnd),
	)
	(*goFtdcMdSpi)(
		(*C.CFtdcMdSpiExt)(this).spi,
	).callback.OnRtnMarketDataEnd(
		(&CRsaFtdcNtfMarketDataEndField{}).FromCStruct(pMarketDataEnd),
	)
}

type reqRtn int

func (rtn reqRtn) Error() string {
	return fmt.Sprintf("req execution failed: %d", rtn)
}

func CheckRtn(rtn C.int) error {
	if rtn == 0 {
		return nil
	}

	return reqRtn(rtn)
}

type CFtdcMdApi struct {
	apiPtr *C.CFtdcMdApiExt

	lib unsafe.Pointer

	spiPtr       *goFtdcMdSpi
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
	libPath, err = filepath.Abs(libPath)
	if err != nil {
		return nil, err
	}

	lib_path := C.CString(libPath)
	defer C.free(unsafe.Pointer(lib_path))

	slog.Info(
		"try to open rmd api lib",
		slog.String("lib", libPath),
	)

	lib := C.dlopen(lib_path, C.RTLD_LAZY)
	if lib == nil {
		msg := C.dlerror()

		err = fmt.Errorf(
			"%w: %s", errLibOpenFialed,
			ReadCString(([]byte)(C.GoString(msg))),
		)

		return
	}

	slog.Info("rmd api lib opened", slog.String("lib", libPath))

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
		slog.Info(
			"rmd api instance creator found, try to create api instance",
		)

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
				lib:    lib,
			}

			runtime.SetFinalizer(api, func(ins *CFtdcMdApi) {
				ins.Release()
				ins.spiPtr.pinner.Unpin()
				ins.spiPtr = nil

				C.dlclose(ins.lib)
			})
		}
	}

	return
}

func (api *CFtdcMdApi) Release() {
	slog.Log(
		context.Background(), slog.LevelDebug-1,
		"executing rmd CallRelease",
	)

	C.CallRelease(
		api.apiPtr.vtable.CFtdcMdApiVtable_Release,
		unsafe.Pointer(api.apiPtr),
	)
	slog.Log(
		context.Background(), slog.LevelDebug-1,
		"rmd CallRelease executed",
	)
}

func (api *CFtdcMdApi) Init() {
	slog.Log(
		context.Background(), slog.LevelDebug-1,
		"executing rmd CallInit",
	)
	C.CallInit(
		api.apiPtr.vtable.CFtdcMdApiVtable_Init,
		unsafe.Pointer(api.apiPtr),
	)
	slog.Log(
		context.Background(), slog.LevelDebug-1,
		"rmd CallInit executed",
	)
}

func (api *CFtdcMdApi) Join() (rtn error) {
	slog.Log(
		context.Background(), slog.LevelDebug-1,
		"executing rmd CallJoin",
	)
	rtn = CheckRtn(C.CallJoin(
		api.apiPtr.vtable.CFtdcMdApiVtable_Join,
		unsafe.Pointer(api.apiPtr),
	))
	slog.Log(
		context.Background(), slog.LevelDebug-1,
		"rmd CallJoin executed",
		slog.Any("rtn", rtn),
	)

	return
}

func (api *CFtdcMdApi) ChkCurrentThread() (rtn bool) {
	slog.Log(
		context.Background(), slog.LevelDebug-1,
		"executing rmd CallChkCurrentThread",
	)
	rtn = bool(C.CallChkCurrentThread(
		api.apiPtr.vtable.CFtdcMdApiVtable_ChkCurrentThread,
		unsafe.Pointer(api.apiPtr),
	))

	slog.Log(
		context.Background(), slog.LevelDebug-1,
		"rmd CallChkCurrentThread executed",
		slog.Any("rtn", rtn),
	)

	return
}

func (api *CFtdcMdApi) RegisterNameServer(nsAddress string) {
	pszNsAddress := C.CString(nsAddress)
	defer C.free(unsafe.Pointer(pszNsAddress))

	slog.Log(
		context.Background(), slog.LevelDebug-1,
		"executing rmd CallRegisterNameServer",
		slog.String("addr", nsAddress),
	)
	C.CallRegisterNameServer(
		api.apiPtr.vtable.CFtdcMdApiVtable_RegisterNameServer,
		unsafe.Pointer(api.apiPtr),
		pszNsAddress,
	)

	slog.Log(
		context.Background(), slog.LevelDebug-1,
		"rmd CallRegisterNameServer executed",
		slog.String("addr", nsAddress),
	)
}

func (api *CFtdcMdApi) RegisterSpi(spi RmdSpi) {
	if spi == nil {
		slog.Warn("spi interface is nil")
		return
	}

	api.registerOnce.Do(func() {
		api.spiPtr = new(goFtdcMdSpi)
		api.spiPtr.callback = spi

		// 将go的spi实例pin在内存中，以免GC或其他情况移动内存导致指针失效
		api.spiPtr.pinner.Pin(unsafe.Pointer(api.spiPtr))
		slog.Debug(
			"spi wrapper pinned",
			slog.Any("spi", spi),
			slog.Any("pinned", api.spiPtr),
		)

		cSpi := C.malloc(C.sizeof_CFtdcMdSpiExt)
		(*C.CFtdcMdSpiExt)(cSpi).vtable = spiCVtablePtr
		(*C.CFtdcMdSpiExt)(cSpi).spi = unsafe.Pointer(api.spiPtr)

		slog.Log(
			context.Background(), slog.LevelDebug-1,
			"executing rmd CallRegisterSpi",
			slog.Any("spi", spi),
		)

		C.CallRegisterSpi(
			api.apiPtr.vtable.CFtdcMdApiVtable_RegisterSpi,
			unsafe.Pointer(api.apiPtr),
			cSpi,
		)

		slog.Log(
			context.Background(), slog.LevelDebug-1,
			"rmd CallRegisterSpi executed",
			slog.Any("spi", spi),
		)
	})
}

func (api *CFtdcMdApi) ReqUserLogin(
	login *CRsaFtdcReqUserLoginField,
) (reqID int32, rtn error) {
	if login == nil {
		return -1, fmt.Errorf(
			"%w: login info is nil", ErrInvalidReqData,
		)
	}

	reqID = api.reqID.Add(1)

	slog.Log(
		context.Background(), slog.LevelDebug-1,
		"executing rmd CallReqUserLogin",
		slog.Any("login", login),
		slog.Int("req_id", int(reqID)),
	)
	rtn = CheckRtn(C.CallReqUserLogin(
		api.apiPtr.vtable.CFtdcMdApiVtable_ReqUserLogin,
		unsafe.Pointer(api.apiPtr),
		login.ToCStruct(),
		C.int(reqID),
	))

	slog.Log(
		context.Background(), slog.LevelDebug-1,
		"rmd CallReqUserLogin executed",
		slog.Any("login", login),
		slog.Int("req_id", int(reqID)),
		slog.Any("rtn", rtn),
	)

	return
}

func (api *CFtdcMdApi) ReqSubMarketData(
	instruments ...string,
) (reqID int32, rtn error) {
	if len(instruments) <= 0 {
		return -1, fmt.Errorf(
			"%w: no sub instruments specified", ErrInvalidReqData,
		)
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

	reqID = api.reqID.Add(1)
	slog.Log(
		context.Background(), slog.LevelDebug-1,
		"executing rmd CallReqSubMarketData",
		slog.Any("subs", instruments),
		slog.Int("req_id", int(reqID)),
	)

	rtn = CheckRtn(C.CallReqSubMarketData(
		api.apiPtr.vtable.CFtdcMdApiVtable_ReqSubMarketData,
		unsafe.Pointer(api.apiPtr),
		(**C.char)(unsafe.Pointer(&ppInstruments[0])),
		C.int(len(ppInstruments)),
		C.int(reqID),
	))

	slog.Log(
		context.Background(), slog.LevelDebug-1,
		"rmd CallReqSubMarketData executed",
		slog.Any("subs", instruments),
		slog.Int("req_id", int(reqID)),
		slog.Any("rtn", rtn),
	)

	return
}

func (api *CFtdcMdApi) ReqUnSubMarketData(
	instruments ...string,
) (reqID int32, rtn error) {
	if len(instruments) <= 0 {
		return -1, fmt.Errorf(
			"%w: no un-sub instruments specified", ErrInvalidReqData,
		)
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

	reqID = api.reqID.Add(1)

	slog.Log(
		context.Background(), slog.LevelDebug-1,
		"executing rmd CallReqUnSubMarketData",
		slog.Any("un-subs", instruments),
		slog.Int("req_id", int(reqID)),
	)
	rtn = CheckRtn(C.CallReqUnSubMarketData(
		api.apiPtr.vtable.CFtdcMdApiVtable_ReqUnSubMarketData,
		unsafe.Pointer(api.apiPtr),
		(**C.char)(unsafe.Pointer(&ppInstruments[0])),
		C.int(len(ppInstruments)),
		C.int(reqID),
	))

	slog.Log(
		context.Background(), slog.LevelDebug-1,
		"rmd CallReqUnSubMarketData executed",
		slog.Any("un-subs", instruments),
		slog.Int("req_id", int(reqID)),
		slog.Any("rtn", rtn),
	)

	return
}

func (api *CFtdcMdApi) ReqQryMarketData(
	qry *CRsaFtdcQryMarketDataField,
) (reqID int32, rtn error) {
	if qry == nil {
		return -1, fmt.Errorf(
			"%w: query market data info is nil", ErrInvalidReqData,
		)
	}

	reqID = api.reqID.Add(1)
	slog.Log(
		context.Background(), slog.LevelDebug-1,
		"executing rmd CallReqQryMarketData",
		slog.Any("qry", qry),
		slog.Int("req_id", int(reqID)),
	)
	rtn = CheckRtn(C.CallReqQryMarketData(
		api.apiPtr.vtable.CFtdcMdApiVtable_ReqQryMarketData,
		unsafe.Pointer(api.apiPtr),
		qry.ToCStruct(),
		C.int(reqID),
	))

	slog.Log(
		context.Background(), slog.LevelDebug-1,
		"rmd CallReqQryMarketData executed",
		slog.Any("qry", qry),
		slog.Int("req_id", int(reqID)),
		slog.Any("rtn", rtn),
	)

	return
}

func (api *CFtdcMdApi) ReqSendDepthMarketDataRead(
	depth *CRsaFtdcDepthMarketDataField,
) (reqID int32, rtn error) {
	if depth == nil {
		return -1, fmt.Errorf(
			"%w: requst send depth market data is nil", ErrInvalidReqData,
		)
	}

	reqID = api.reqID.Add(1)
	slog.Log(
		context.Background(), slog.LevelDebug-1,
		"executing rmd CallReqSendDepthMarketDataRead",
		slog.Any("depth", depth),
		slog.Int("req_id", int(reqID)),
	)
	rtn = CheckRtn(C.CallReqSendDepthMarketDataRead(
		api.apiPtr.vtable.CFtdcMdApiVtable_ReqSendDepthMarketDataRead,
		unsafe.Pointer(api.apiPtr),
		depth.ToCStruct(),
		C.int(reqID),
	))

	slog.Log(
		context.Background(), slog.LevelDebug-1,
		"rmd CallReqSendDepthMarketDataRead executed",
		slog.Any("depth", depth),
		slog.Int("req_id", int(reqID)),
		slog.Any("rtn", rtn),
	)

	return
}

func (api *CFtdcMdApi) ReqSendBarMarketDataRead(
	bar *CRsaFtdcBarMarketDataField,
) (reqID int32, rtn error) {
	if bar == nil {
		return -1, fmt.Errorf(
			"%w: requst send bar market data is nil", ErrInvalidReqData,
		)
	}

	reqID = api.reqID.Add(1)
	slog.Log(
		context.Background(), slog.LevelDebug-1,
		"executing rmd CallReqSendBarMarketDataRead",
		slog.Any("bar", bar),
		slog.Int("req_id", int(reqID)),
	)

	rtn = CheckRtn(C.CallReqSendBarMarketDataRead(
		api.apiPtr.vtable.CFtdcMdApiVtable_ReqSendBarMarketDataRead,
		unsafe.Pointer(api.apiPtr),
		bar.ToCStruct(),
		C.int(reqID),
	))

	slog.Log(
		context.Background(), slog.LevelDebug-1,
		"rmd CallReqSendBarMarketDataRead executed",
		slog.Any("bar", bar),
		slog.Any("rtn", rtn),
	)

	return
}

func (api *CFtdcMdApi) ReqBtSubMarketData(
	subs ...*CRsaFtdcBtSubMarketDataField,
) (reqID int32, rtn error) {
	if len(subs) <= 0 {
		return -1, fmt.Errorf(
			"%w: no sub list specified", ErrInvalidReqData,
		)
	}

	ppFields := make([]*C.struct_CRsaFtdcBtSubMarketDataField, len(subs))
	for idx, sub := range subs {
		ppFields[idx] = sub.ToCStruct()
	}
	defer func() {
		for _, ptr := range ppFields {
			C.free(unsafe.Pointer(ptr))
		}
	}()

	reqID = api.reqID.Add(1)
	slog.Log(
		context.Background(), slog.LevelDebug-1,
		"executing rmd CallReqBtSubMarketData",
		slog.Any("subs", subs),
		slog.Int("req_id", int(reqID)),
	)
	rtn = CheckRtn(C.CallReqBtSubMarketData(
		api.apiPtr.vtable.CFtdcMdApiVtable_ReqBtSubMarketData,
		unsafe.Pointer(api.apiPtr),
		(**C.struct_CRsaFtdcBtSubMarketDataField)(unsafe.Pointer(&ppFields[0])),
		C.int(len(subs)),
		C.int(reqID),
	))

	slog.Log(
		context.Background(), slog.LevelDebug-1,
		"rmd CallReqBtSubMarketData executed",
		slog.Any("subs", subs),
		slog.Int("req_id", int(reqID)),
		slog.Any("rtn", rtn),
	)

	return
}

func (api *CFtdcMdApi) ReqBtUnSubMarketData(
	unSubs ...*CRsaFtdcBtSubMarketDataField,
) (reqID int32, rtn error) {
	if len(unSubs) <= 0 {
		return -1, fmt.Errorf(
			"%w: no un-sub list specified", ErrInvalidReqData,
		)
	}

	ppFields := make([]*C.struct_CRsaFtdcBtSubMarketDataField, len(unSubs))
	for idx, sub := range unSubs {
		ppFields[idx] = sub.ToCStruct()
	}
	defer func() {
		for _, ptr := range ppFields {
			C.free(unsafe.Pointer(ptr))
		}
	}()

	reqID = api.reqID.Add(1)
	slog.Log(
		context.Background(), slog.LevelDebug-1,
		"executing rmd CallReqBtUnSubMarketData",
		slog.Any("un-subs", unSubs),
		slog.Int("req_id", int(reqID)),
	)
	rtn = CheckRtn(C.CallReqBtUnSubMarketData(
		api.apiPtr.vtable.CFtdcMdApiVtable_ReqBtUnSubMarketData,
		unsafe.Pointer(api.apiPtr),
		(**C.struct_CRsaFtdcBtSubMarketDataField)(unsafe.Pointer(&ppFields[0])),
		C.int(len(unSubs)),
		C.int(reqID),
	))

	slog.Log(
		context.Background(), slog.LevelDebug-1,
		"rmd CallReqBtUnSubMarketData executed",
		slog.Any("un-subs", unSubs),
		slog.Int("req_id", int(reqID)),
		slog.Any("rtn", rtn),
	)

	return
}

func (api *CFtdcMdApi) ReqSubMarketDataCompleted() (rtn error) {
	slog.Log(
		context.Background(), slog.LevelDebug-1,
		"executing rmd CallReqSubMarketDataCompleted",
	)

	rtn = CheckRtn(C.CallReqSubMarketDataCompleted(
		api.apiPtr.vtable.CFtdcMdApiVtable_ReqSubMarketDataCompleted,
		unsafe.Pointer(api.apiPtr),
	))

	slog.Log(
		context.Background(), slog.LevelDebug-1,
		"rmd CallReqSubMarketDataCompleted executed",
	)

	return
}

func (api *CFtdcMdApi) ReqSubCombMarketData(
	subs ...*CRsaFtdcSubCombMarketDataField,
) (rtn error) {
	if len(subs) <= 0 {
		return fmt.Errorf(
			"%w: no combo sub list specified",
			ErrInvalidReqData,
		)
	}

	ppFields := make([]*C.struct_CRsaFtdcSubCombMarketDataField, len(subs))
	for idx, sub := range subs {
		ppFields[idx] = sub.ToCStruct()
	}

	slog.Log(
		context.Background(), slog.LevelDebug-1,
		"executing rmd CallReqSubCombMarketData",
		slog.Any("subs", subs),
	)

	rtn = CheckRtn(C.CallReqSubCombMarketData(
		api.apiPtr.vtable.CFtdcMdApiVtable_ReqSubCombMarketData,
		unsafe.Pointer(api.apiPtr),
		(**C.struct_CRsaFtdcSubCombMarketDataField)(
			unsafe.Pointer(&ppFields[0])),
		C.int(len(subs)),
	))

	slog.Log(
		context.Background(), slog.LevelDebug-1,
		"rmd CallReqSubCombMarketData executed",
		slog.Any("subs", subs),
		slog.Any("rtn", rtn),
	)

	return
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

	slog.Log(
		context.Background(), slog.LevelDebug-1,
		"executing rmd CallReqQryBarMarketData",
		slog.Any("sub", sub),
		slog.Int("count", count),
	)
	rtn := CheckRtn(C.CallReqQryBarMarketData(
		api.apiPtr.vtable.CFtdcMdApiVtable_ReqQryBarMarketData,
		unsafe.Pointer(api.apiPtr),
		sub.ToCStruct(),
		(**C.struct_CRsaFtdcBarMarketDataField)(unsafe.Pointer(&data[0])),
		&outCount,
	))

	buffCount := int(outCount)

	slog.Log(
		context.Background(), slog.LevelDebug-1,
		"rmd CallReqQryBarMarketData executed",
		slog.Any("sub", sub),
		slog.Int("out_count", buffCount),
		slog.Any("rtn", rtn),
	)

	if rtn != nil {
		return nil, rtn
	} else if buffCount > count {
		return nil, errors.New("invalid queried bar buffer len")
	} else {
		bars := make([]*CRsaFtdcBarMarketDataField, buffCount)

		for idx, bar := range data[:buffCount] {
			if bar == nil {
				slog.Error(
					"invalid queried bar pointer",
					slog.Int("idx", idx),
					slog.Any("pointer", bar),
				)
				return nil, errors.New("invalid bar pointer")
			}
			bars[idx] = GetBar().FromCStruct(bar)
		}

		return bars, nil
	}
}

func (api *CFtdcMdApi) ReqQryDepthMarketData(
	sub *CRsaFtdcQryMarketDataField, count int,
) ([]*CRsaFtdcDepthMarketDataField, error) {
	if sub == nil {
		return nil, fmt.Errorf(
			"%w: query depth data is nil", ErrInvalidReqData,
		)
	}

	if count <= 0 {
		return nil, fmt.Errorf(
			"%w: query bar count is empty", ErrInvalidReqData,
		)
	}

	data := make([]*C.struct_CRsaFtdcDepthMarketDataField, count)
	outCount := C.int(count)

	slog.Log(
		context.Background(), slog.LevelDebug-1,
		"executing rmd CallReqQryDepthMarketData",
		slog.Any("sub", sub),
		slog.Int("count", count),
	)

	rtn := CheckRtn(C.CallReqQryDepthMarketData(
		api.apiPtr.vtable.CFtdcMdApiVtable_ReqQryDepthMarketData,
		unsafe.Pointer(api.apiPtr),
		sub.ToCStruct(),
		(**C.struct_CRsaFtdcDepthMarketDataField)(unsafe.Pointer(&data[0])),
		&outCount,
	))

	buffCount := int(outCount)

	slog.Log(
		context.Background(), slog.LevelDebug-1,
		"rmd CallReqQryDepthMarketData executed",
		slog.Any("sub", sub),
		slog.Int("out_count", buffCount),
		slog.Any("rtn", rtn),
	)

	if rtn != nil {
		return nil, rtn
	} else if buffCount > count {
		return nil, errors.New("invalid queried tick buff len")
	} else {
		ticks := make([]*CRsaFtdcDepthMarketDataField, buffCount)

		for idx, tick := range data[:buffCount] {
			if tick == nil {
				slog.Error(
					"invalid queried tick pointer",
					slog.Int("idx", idx),
					slog.Any("pointer", tick),
				)
				return nil, errors.New("invalid tick pointer")
			}
			ticks[idx] = GetTick().FromCStruct(tick)
		}

		return ticks, nil
	}
}

func init() {
	spiCVtablePtr = (*C.CFtdcMdSpiVtable)(C.malloc(C.sizeof_CFtdcMdSpiVtable))

	spiCVtablePtr.CFtdcMdSpi_OnMdFrontConnected = (C.OnMdFrontConnected)(
		unsafe.Pointer(C.COnMdFrontConnected))
	spiCVtablePtr.CFtdcMdSpi_OnMdFrontDisconnected = (C.OnMdFrontDisconnected)(
		unsafe.Pointer(C.COnMdFrontDisconnected))
	spiCVtablePtr.CFtdcMdSpi_OnMdRspUserLogin = (C.OnMdRspUserLogin)(
		unsafe.Pointer(C.COnMdRspUserLogin))
	spiCVtablePtr.CFtdcMdSpi_OnRtnDepthMarketData = (C.OnRtnDepthMarketData)(
		unsafe.Pointer(C.COnRtnDepthMarketData))
	spiCVtablePtr.CFtdcMdSpi_OnRspSubMarketData = (C.OnRspSubMarketData)(
		unsafe.Pointer(C.COnRspSubMarketData))
	spiCVtablePtr.CFtdcMdSpi_OnRspUnSubMarketData = (C.OnRspUnSubMarketData)(
		unsafe.Pointer(C.COnRspUnSubMarketData))
	spiCVtablePtr.CFtdcMdSpi_OnRtnBarMarketData = (C.OnRtnBarMarketData)(
		unsafe.Pointer(C.COnRtnBarMarketData))
	spiCVtablePtr.CFtdcMdSpi_OnRspQryMarketData = (C.OnRspQryMarketData)(
		unsafe.Pointer(C.COnRspQryMarketData))
	spiCVtablePtr.CFtdcMdSpi_OnRtnMarketDataEnd = (C.OnRtnMarketDataEnd)(
		unsafe.Pointer(C.COnRtnMarketDataEnd))
}

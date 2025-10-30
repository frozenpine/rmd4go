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
	OnRtnMarketDataEnd(td *CRsaFtdcNtfMarketDataEndField)
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
	rsp := CRsaFtdcRspUserLoginField{}
	rsp.TradingDay = C.GoString(&pRspUserLogin.TradingDay[0])
	info := CRsaFtdcRspInfoField{}
	info.ErrorID = int(pRspInfo.ErrorID)
	info.ErrorMsg = ReadCString([]byte(C.GoString(&pRspInfo.ErrorMsg[0])))

	(*goFtdcMdSpi)(
		(*C.CFtdcMdSpiExt)(this).spi,
	).callback.OnMdRspUserLogin(
		&rsp, &info, int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRtnDepthMarketData
func CgoOnRtnDepthMarketData(
	this unsafe.Pointer,
	pDepthMarketData *C.struct_CRsaFtdcDepthMarketDataField,
) {
	(*goFtdcMdSpi)(
		(*C.CFtdcMdSpiExt)(this).spi,
	).callback.OnRtnDepthMarketData(
		(*CRsaFtdcDepthMarketDataField)(unsafe.Pointer(pDepthMarketData)))
}

//export CgoOnRspSubMarketData
func CgoOnRspSubMarketData(
	this unsafe.Pointer,
	pSpecificInstrument *C.struct_CRsaFtdcSpecificInstrumentField,
	pRspInfo *C.struct_CRsaFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	(*goFtdcMdSpi)(
		(*C.CFtdcMdSpiExt)(this).spi,
	).callback.OnRspSubMarketData(
		(*CRsaFtdcSpecificInstrumentField)(unsafe.Pointer(pSpecificInstrument)),
		(*CRsaFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
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
	(*goFtdcMdSpi)(
		(*C.CFtdcMdSpiExt)(this).spi,
	).callback.OnRspUnSubMarketData(
		(*CRsaFtdcSpecificInstrumentField)(unsafe.Pointer(pSpecificInstrument)),
		(*CRsaFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRtnBarMarketData
func CgoOnRtnBarMarketData(
	this unsafe.Pointer,
	pBarMarketData *C.struct_CRsaFtdcBarMarketDataField,
) {
	(*goFtdcMdSpi)(
		(*C.CFtdcMdSpiExt)(this).spi,
	).callback.OnRtnBarMarketData(
		(*CRsaFtdcBarMarketDataField)(unsafe.Pointer(pBarMarketData)),
	)
}

//export CgoOnRspQryMarketData
func CgoOnRspQryMarketData(
	this unsafe.Pointer,
	pRspMarketData *C.struct_CRsaFtdcRspMarketDataField,
	pRspInfo *C.struct_CRsaFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	(*goFtdcMdSpi)(
		(*C.CFtdcMdSpiExt)(this).spi,
	).callback.OnRspQryMarketData(
		(*CRsaFtdcRspMarketDataField)(unsafe.Pointer(pRspMarketData)),
		(*CRsaFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRtnMarketDataEnd
func CgoOnRtnMarketDataEnd(
	this unsafe.Pointer,
	pMarketDataEnd *C.struct_CRsaFtdcNtfMarketDataEndField,
) {
	(*goFtdcMdSpi)(
		(*C.CFtdcMdSpiExt)(this).spi,
	).callback.OnRtnMarketDataEnd(
		(*CRsaFtdcNtfMarketDataEndField)(unsafe.Pointer(pMarketDataEnd)),
	)
}

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
				ins.spiPtr = nil
				ins.spiPtr.pinner.Unpin()

				C.dlclose(ins.lib)
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
		api.spiPtr = new(goFtdcMdSpi)
		api.spiPtr.callback = spi

		api.spiPtr.pinner.Pin(unsafe.Pointer(api.spiPtr))
		slog.Debug(
			"spi wrapper pinned",
			slog.Any("spi", spi),
			slog.Any("pinned", api.spiPtr),
		)

		cSpi := C.malloc(C.sizeof_CFtdcMdSpiExt)
		(*C.CFtdcMdSpiExt)(cSpi).vtable = spiCVtablePtr
		(*C.CFtdcMdSpiExt)(cSpi).spi = unsafe.Pointer(api.spiPtr)

		C.CallRegisterSpi(
			api.apiPtr.vtable.CFtdcMdApiVtable_RegisterSpi,
			unsafe.Pointer(api.apiPtr),
			cSpi,
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

func init() {
	slog.Info("initializing RMD SPI vtable")

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

	slog.Info("RMD SPI vtable initialized")
}

package rmd4go

/*
#cgo CFLAGS: -I${SRCDIR} -I${SRCDIR}/dependencies/includes

#include <stdbool.h>
#include <stdint.h>

#include "FtdcApiStruct.h"

// spi definitions
typedef void (*OnMdFrontConnected)(void* this);
typedef void (*OnMdFrontDisconnected)(void* this, int nReason);
typedef void (*OnMdRspUserLogin)(
	void* this,
	struct CRsaFtdcRspUserLoginField *pRspUserLogin,
	struct CRsaFtdcRspInfoField *pRspInfo,
	int nRequestID, bool bIsLast);
typedef void (*OnRtnDepthMarketData)(
	void* this,
	struct CRsaFtdcDepthMarketDataField *pDepthMarketData);
typedef void (*OnRspSubMarketData)(
	void* this,
	struct CRsaFtdcSpecificInstrumentField *pSpecificInstrument,
	struct CRsaFtdcRspInfoField *pRspInfo,
	int nRequestID, bool bIsLast);
typedef void (*OnRspUnSubMarketData)(
	void* this,
	struct CRsaFtdcSpecificInstrumentField *pSpecificInstrument,
	struct CRsaFtdcRspInfoField *pRspInfo,
	int nRequestID, bool bIsLast);
typedef void (*OnRtnBarMarketData)(
	void* this, struct CRsaFtdcBarMarketDataField *pBarMarketData);
typedef void (*OnRspQryMarketData)(
	void* this,
	struct CRsaFtdcRspMarketDataField *pRspMarketData,
	struct CRsaFtdcRspInfoField *pRspInfo,
	int nRequestID, bool bIsLast);
typedef void (*OnRtnMarketDataEnd)(
	void* this, struct CRsaFtdcNtfMarketDataEndField *pMarketDataEnd);

extern void CgoOnMdFrontConnected(void* this);

extern void CgoOnMdFrontDisconnected(void* this, int nReason);

extern void CgoOnMdRspUserLogin(
	void* this,
	struct CRsaFtdcRspUserLoginField *pRspUserLogin,
	struct CRsaFtdcRspInfoField *pRspInfo,
	int nRequestID, bool bIsLast);

extern void CgoOnRtnDepthMarketData(
	void* this, struct CRsaFtdcDepthMarketDataField *pDepthMarketData);

extern void CgoOnRspSubMarketData(
	void* this,
	struct CRsaFtdcSpecificInstrumentField *pSpecificInstrument,
	struct CRsaFtdcRspInfoField *pRspInfo,
	int nRequestID, bool bIsLast);

extern void CgoOnRspUnSubMarketData(
	void* this,
	struct CRsaFtdcSpecificInstrumentField *pSpecificInstrument,
	struct CRsaFtdcRspInfoField *pRspInfo,
	int nRequestID, bool bIsLast);

extern void CgoOnRtnBarMarketData(
	void* this, struct CRsaFtdcBarMarketDataField *pBarMarketData);

extern void CgoOnRspQryMarketData(
	void* this,
	struct CRsaFtdcRspMarketDataField *pRspMarketData,
	struct CRsaFtdcRspInfoField *pRspInfo,
	int nRequestID, bool bIsLast);

extern void CgoOnRtnMarketDataEnd(
	void* this, struct CRsaFtdcNtfMarketDataEndField *pMarketDataEnd);

typedef struct {
	OnMdFrontConnected CFtdcMdSpi_OnMdFrontConnected;
	OnMdFrontDisconnected CFtdcMdSpi_OnMdFrontDisconnected;
	OnMdRspUserLogin CFtdcMdSpi_OnMdRspUserLogin;
	OnRtnDepthMarketData CFtdcMdSpi_OnRtnDepthMarketData;
	OnRspSubMarketData CFtdcMdSpi_OnRspSubMarketData;
	OnRspUnSubMarketData CFtdcMdSpi_OnRspUnSubMarketData;
	OnRtnBarMarketData CFtdcMdSpi_OnRtnBarMarketData;
	OnRspQryMarketData CFtdcMdSpi_OnRspQryMarketData;
	OnRtnMarketDataEnd CFtdcMdSpi_OnRtnMarketDataEnd;
} CFtdcMdSpiVtable;
*/
import "C"
import "unsafe"

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

var callbackVTable = C.CFtdcMdSpiVtable{
	CFtdcMdSpi_OnMdFrontConnected: (C.OnMdFrontConnected)(
		unsafe.Pointer(C.CgoOnMdFrontConnected)),
	CFtdcMdSpi_OnMdFrontDisconnected: (C.OnMdFrontDisconnected)(
		unsafe.Pointer(C.CgoOnMdFrontDisconnected)),
	CFtdcMdSpi_OnMdRspUserLogin: (C.OnMdRspUserLogin)(
		unsafe.Pointer(C.CgoOnMdRspUserLogin)),
	CFtdcMdSpi_OnRtnDepthMarketData: (C.OnRtnDepthMarketData)(
		unsafe.Pointer(C.CgoOnRtnDepthMarketData)),
	CFtdcMdSpi_OnRspSubMarketData: (C.OnRspSubMarketData)(
		unsafe.Pointer(C.CgoOnRspSubMarketData)),
	CFtdcMdSpi_OnRspUnSubMarketData: (C.OnRspUnSubMarketData)(
		unsafe.Pointer(C.CgoOnRspUnSubMarketData)),
	CFtdcMdSpi_OnRtnBarMarketData: (C.OnRtnBarMarketData)(
		unsafe.Pointer(C.CgoOnRtnBarMarketData)),
	CFtdcMdSpi_OnRspQryMarketData: (C.OnRspQryMarketData)(
		unsafe.Pointer(C.CgoOnRspQryMarketData)),
	CFtdcMdSpi_OnRtnMarketDataEnd: (C.OnRtnMarketDataEnd)(
		unsafe.Pointer(C.CgoOnRtnMarketDataEnd)),
}

type cFtdcMdSpi struct {
	vtable C.CFtdcMdSpiVtable

	callback RmdSpi
}

//export CgoOnMdFrontConnected
func CgoOnMdFrontConnected(this unsafe.Pointer) {
	(*cFtdcMdSpi)(this).callback.OnMdFrontConnected()
}

//export CgoOnMdFrontDisconnected
func CgoOnMdFrontDisconnected(this unsafe.Pointer, nReason C.int) {
	(*cFtdcMdSpi)(this).callback.OnMdFrontDisconnected(int(nReason))
}

//export CgoOnMdRspUserLogin
func CgoOnMdRspUserLogin(
	this unsafe.Pointer,
	pRspUserLogin *C.struct_CRsaFtdcRspUserLoginField,
	pRspInfo *C.struct_CRsaFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	(*cFtdcMdSpi)(this).callback.OnMdRspUserLogin(
		(*CRsaFtdcRspUserLoginField)(unsafe.Pointer(pRspUserLogin)),
		(*CRsaFtdcRspInfoField)(unsafe.Pointer(pRspInfo)),
		int(nRequestID), bool(bIsLast),
	)
}

//export CgoOnRtnDepthMarketData
func CgoOnRtnDepthMarketData(
	this unsafe.Pointer,
	pDepthMarketData *C.struct_CRsaFtdcDepthMarketDataField,
) {
	(*cFtdcMdSpi)(this).callback.OnRtnDepthMarketData(
		(*CRsaFtdcDepthMarketDataField)(unsafe.Pointer(pDepthMarketData)))
}

//export CgoOnRspSubMarketData
func CgoOnRspSubMarketData(
	this unsafe.Pointer,
	pSpecificInstrument *C.struct_CRsaFtdcSpecificInstrumentField,
	pRspInfo *C.struct_CRsaFtdcRspInfoField,
	nRequestID C.int, bIsLast C.bool,
) {
	(*cFtdcMdSpi)(this).callback.OnRspSubMarketData(
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
	(*cFtdcMdSpi)(this).callback.OnRspUnSubMarketData(
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
	(*cFtdcMdSpi)(this).callback.OnRtnBarMarketData(
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
	(*cFtdcMdSpi)(this).callback.OnRspQryMarketData(
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
	(*cFtdcMdSpi)(this).callback.OnRtnMarketDataEnd(
		(*CRsaFtdcNtfMarketDataEndField)(unsafe.Pointer(pMarketDataEnd)),
	)
}

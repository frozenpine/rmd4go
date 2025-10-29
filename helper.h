#pragma once

#ifndef RMD4GO_HELPER_H
#define RMD4GO_HELPER_H

#ifdef __cplusplus
extern "C" {
#endif

#include <stdbool.h>
#include <stdint.h>
#include <stdlib.h>
#include <strings.h>
#include <dlfcn.h>

#include "FtdcApiStruct.h"

/// api definetions
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
);

void CallRelease(Release fn, void* this);

void CallInit(Init fn, void* this);

int CallJoin(Join fn, void* this);

bool CallChkCurrentThread(ChkCurrentThread fn, void* this);

void CallRegisterNameServer(
	RegisterNameServer fn, void* this,
	const char* pszNsAddress
);

void CallRegisterSpi(
	RegisterSpi fn, void* this,
	void* pSpi
);

int CallReqUserLogin(
	ReqUserLogin fn, void* this,
	struct CRsaFtdcReqUserLoginField *pReqUserLoginField,
	int nRequestID
);

int CallReqSubMarketData(
	ReqSubMarketData fn, void* this,
	char *ppInstrumentID[],
	int nCount
);

int CallReqUnSubMarketData(
	ReqUnSubMarketData fn, void* this,
	char *ppInstrumentID[],
	int nCount
);

int CallReqQryMarketData(
	ReqQryMarketData fn, void* this,
	struct CRsaFtdcQryMarketDataField *pQryMarketData,
	int nRequestID
);

int CallReqSendDepthMarketDataRead(
	ReqSendDepthMarketDataRead fn, void* this,
	struct CRsaFtdcDepthMarketDataField *pDepthMarketData,
	int nRequestID
);

int CallReqSendBarMarketDataRead(
	ReqSendBarMarketDataRead fn, void* this,
	struct CRsaFtdcBarMarketDataField *pBarMarketData,
	int nRequestID
);

int CallReqBtSubMarketData(
	ReqBtSubMarketData fn, void* this,
	struct CRsaFtdcBtSubMarketDataField *ppFiels[],
	int nCount
);

int CallReqBtUnSubMarketData(
	ReqBtUnSubMarketData fn, void* this,
	struct CRsaFtdcBtSubMarketDataField *ppFiels[],
	int nCount
);

int CallReqSubMarketDataCompleted(
	ReqSubMarketDataCompleted fn, void* this
);

int CallReqSubCombMarketData(
	ReqSubCombMarketData fn, void* this,
	struct CRsaFtdcSubCombMarketDataField* ppFiels[],
	int nCount
);

int CallReqQryBarMarketData(
	ReqQryBarMarketData fn, void* this,
	struct CRsaFtdcBtSubMarketDataField* subMd,
	struct CRsaFtdcBarMarketDataField* bars[],
	int* count
);

int CallReqQryDepthMarketData(
	ReqQryDepthMarketData fn, void* this,
	struct CRsaFtdcQryMarketDataField* pQryMarketData,
	struct CRsaFtdcDepthMarketDataField* depthMd[],
	int* count
);

// spi definetions
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

void COnMdFrontConnected(void* this);

void COnMdFrontDisconnected(void* this, int nReason);

void COnMdRspUserLogin(
	void* this,
	struct CRsaFtdcRspUserLoginField *pRspUserLogin,
	struct CRsaFtdcRspInfoField *pRspInfo,
	int nRequestID, bool bIsLast);

void COnRtnDepthMarketData(
	void* this, struct CRsaFtdcDepthMarketDataField *pDepthMarketData);

void COnRspSubMarketData(
	void* this,
	struct CRsaFtdcSpecificInstrumentField *pSpecificInstrument,
	struct CRsaFtdcRspInfoField *pRspInfo,
	int nRequestID, bool bIsLast);

void COnRspUnSubMarketData(
	void* this,
	struct CRsaFtdcSpecificInstrumentField *pSpecificInstrument,
	struct CRsaFtdcRspInfoField *pRspInfo,
	int nRequestID, bool bIsLast);

void COnRtnBarMarketData(
	void* this, struct CRsaFtdcBarMarketDataField *pBarMarketData);

void COnRspQryMarketData(
	void* this,
	struct CRsaFtdcRspMarketDataField *pRspMarketData,
	struct CRsaFtdcRspInfoField *pRspInfo,
	int nRequestID, bool bIsLast);

void COnRtnMarketDataEnd(
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

#ifdef __cplusplus
}
#endif
#endif
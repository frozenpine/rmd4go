#include "helper.h"

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

void COnMdFrontConnected(void* this) 
{ 
    CgoOnMdFrontConnected(this); 
}

void COnMdFrontDisconnected(void* this, int nReason) 
{
    CgoOnMdFrontDisconnected(this, nReason); 
}

void COnMdRspUserLogin(
	void* this,
	struct CRsaFtdcRspUserLoginField *pRspUserLogin,
	struct CRsaFtdcRspInfoField *pRspInfo,
	int nRequestID, bool bIsLast)
{
    CgoOnMdRspUserLogin(this, pRspUserLogin, pRspInfo, nRequestID, bIsLast);
}

void COnRtnDepthMarketData(
	void* this, struct CRsaFtdcDepthMarketDataField *pDepthMarketData)
{
    CgoOnRtnDepthMarketData(this, pDepthMarketData);
}

void COnRspSubMarketData(
	void* this,
	struct CRsaFtdcSpecificInstrumentField *pSpecificInstrument,
	struct CRsaFtdcRspInfoField *pRspInfo,
	int nRequestID, bool bIsLast)
{
    CgoOnRspSubMarketData(this, pSpecificInstrument, pRspInfo, nRequestID, bIsLast);
}

void COnRspUnSubMarketData(
	void* this,
	struct CRsaFtdcSpecificInstrumentField *pSpecificInstrument,
	struct CRsaFtdcRspInfoField *pRspInfo,
	int nRequestID, bool bIsLast)
{
    CgoOnRspUnSubMarketData(this, pSpecificInstrument, pRspInfo, nRequestID, bIsLast);
}

void COnRtnBarMarketData(
	void* this, struct CRsaFtdcBarMarketDataField *pBarMarketData)
{
    CgoOnRtnBarMarketData(this, pBarMarketData);
}

void COnRspQryMarketData(
	void* this,
	struct CRsaFtdcRspMarketDataField *pRspMarketData,
	struct CRsaFtdcRspInfoField *pRspInfo,
	int nRequestID, bool bIsLast)
{
    CgoOnRspQryMarketData(this, pRspMarketData, pRspInfo, nRequestID, bIsLast);
}

void COnRtnMarketDataEnd(
	void* this, struct CRsaFtdcNtfMarketDataEndField *pMarketDataEnd)
{
    CgoOnRtnMarketDataEnd(this, pMarketDataEnd);
}

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
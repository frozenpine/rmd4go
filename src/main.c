
#include <unistd.h>

#include "helper.h"

typedef struct {
	CFtdcMdSpiVtable* vtable;
} CFtdcMdSpiExt;


const char* CREATE_RMD_API_WIN = "?CreateFtdcMdApi@CFtdcMdApi@@SAPEAV1@PEBD00_N@Z";
const char* CREATE_RMD_API_LIN = "_ZN10CFtdcMdApi15CreateFtdcMdApiEPKcS1_S1_b";
// const char* LIB_PATH = "../dependencies/libs/rmdapi.dll";
const char* LIB_PATH = "../dependencies/libs/rmdapi.so";
const char* FRONT_ADDR = "tcp://172.16.200.105:30010";
const char* FLOW_PATH = "./flow/";

void CgoOnMdFrontConnected(void* this) {
    fprintf(stdout, "md front connected: %x\n", this);
}

void CgoOnMdFrontDisconnected(void* this, int nReason) {}

void CgoOnMdRspUserLogin(
	void* this,
	struct CRsaFtdcRspUserLoginField *pRspUserLogin,
	struct CRsaFtdcRspInfoField *pRspInfo,
	int nRequestID, bool bIsLast) {}

void CgoOnRtnDepthMarketData(
	void* this, struct CRsaFtdcDepthMarketDataField *pDepthMarketData) {}

void CgoOnRspSubMarketData(
	void* this,
	struct CRsaFtdcSpecificInstrumentField *pSpecificInstrument,
	struct CRsaFtdcRspInfoField *pRspInfo,
	int nRequestID, bool bIsLast) {}

void CgoOnRspUnSubMarketData(
	void* this,
	struct CRsaFtdcSpecificInstrumentField *pSpecificInstrument,
	struct CRsaFtdcRspInfoField *pRspInfo,
	int nRequestID, bool bIsLast) {}

void CgoOnRtnBarMarketData(
	void* this, struct CRsaFtdcBarMarketDataField *pBarMarketData){}

void CgoOnRspQryMarketData(
	void* this,
	struct CRsaFtdcRspMarketDataField *pRspMarketData,
	struct CRsaFtdcRspInfoField *pRspInfo,
	int nRequestID, bool bIsLast) {}

void CgoOnRtnMarketDataEnd(
	void* this, struct CRsaFtdcNtfMarketDataEndField *pMarketDataEnd) {}

int main(int argc, char** argv) 
{
    CFtdcMdSpiVtable* SPI_VTABLE = malloc(sizeof(CFtdcMdSpiVtable));
    SPI_VTABLE->CFtdcMdSpi_OnMdFrontConnected = COnMdFrontConnected;
    SPI_VTABLE->CFtdcMdSpi_OnMdFrontDisconnected = COnMdFrontDisconnected;
	SPI_VTABLE->CFtdcMdSpi_OnMdRspUserLogin = COnMdRspUserLogin;
	SPI_VTABLE->CFtdcMdSpi_OnRtnDepthMarketData = COnRtnDepthMarketData;
	SPI_VTABLE->CFtdcMdSpi_OnRspSubMarketData = COnRspSubMarketData;
	SPI_VTABLE->CFtdcMdSpi_OnRspUnSubMarketData = COnRspUnSubMarketData;
	SPI_VTABLE->CFtdcMdSpi_OnRtnBarMarketData = COnRtnBarMarketData;
	SPI_VTABLE->CFtdcMdSpi_OnRspQryMarketData = COnRspQryMarketData;
	SPI_VTABLE->CFtdcMdSpi_OnRtnMarketDataEnd = COnRtnMarketDataEnd;

    void* lib = dlopen(LIB_PATH, RTLD_LAZY);
    if (NULL == lib) {
        fprintf(stderr, "open lib[%s] failed: %s\n", LIB_PATH, dlerror());
        return -1;
    } else {
        fprintf(stdout, "lib[%s] opened\n", LIB_PATH);
    }

    // CreateFtdcMdApi creator = dlsym(lib, CREATE_RMD_API_WIN);
    CreateFtdcMdApi creator = dlsym(lib, CREATE_RMD_API_LIN);
    if (NULL == creator) {
        fprintf(stderr, "api creator[%s] not found: %s\n", CREATE_RMD_API_WIN, dlerror());
        dlclose(lib);
        return -1;
    } else {
        fprintf(stdout, "creator[%s] found in lib\n", CREATE_RMD_API_WIN);
    }

    CFtdcMdApiExt* api = creator(FRONT_ADDR, FLOW_PATH, "c demo", false);
    if (NULL == api) {
        fprintf(stderr, "create api failed\n");
        dlclose(lib);
        return 1;
    } else {
        fprintf(stdout, "api instance created: %x\n", api);
    }

    CFtdcMdApiExt* spi = malloc(sizeof(CFtdcMdSpiExt));
    spi->vtable = SPI_VTABLE;

    api->vtable->CFtdcMdApiVtable_RegisterSpi(api, spi);
    fprintf(stdout, "spi registered: %x\n", spi);
    
    api->vtable->CFtdcMdApiVtable_Init(api);
    fprintf(stdout, "waiting 5s for connect\n");
    
    sleep(5);

    api->vtable->CFtdcMdApiVtable_Release(api);
}
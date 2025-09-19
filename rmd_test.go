package rmd4go_test

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/frozenpine/rmd4go"
)

type TestClient struct {
	*rmd4go.CFtdcMdApi

	t    *testing.T
	wait chan bool
}

func (c *TestClient) OnFrontConnected() {
	c.t.Log("front connected")
	c.wait <- true
}

func (c *TestClient) OnMdFrontDisconnected(reason int) {
	c.t.Logf("front disconnected: %d", reason)
	c.wait <- false
}

func (c *TestClient) OnMdRspUserLogin(
	login *rmd4go.CRsaFtdcRspUserLoginField,
	info *rmd4go.CRsaFtdcRspInfoField,
	requestID int, isLast bool,
) {
	if info.ErrorID != 0 {
		c.t.Logf(
			"user login failed[%d]: %s",
			info.ErrorID, rmd4go.ReadCString(info.ErrorMsg[:]),
		)
		c.wait <- false
		return
	}

	if login != nil {
		c.t.Logf(
			"user login: %s.%s @ %s",
			rmd4go.ReadCString(login.BrokerID[:]),
			rmd4go.ReadCString(login.UserID[:]),
			rmd4go.ReadCString(login.LoginTime[:]),
		)
	}

	c.wait <- true
}

func (c *TestClient) OnRtnDepthMarketData(tick *rmd4go.CRsaFtdcDepthMarketDataField) {
	// TODO: print tick data
	c.t.Log("on rtn tick: ")
}

func (c *TestClient) OnRspSubMarketData(
	instrument *rmd4go.CRsaFtdcSpecificInstrumentField,
	info *rmd4go.CRsaFtdcRspInfoField,
	requestID int, isLast bool,
) {
	if info.ErrorID != 0 {
		c.t.Logf(
			"sub market data failed[%d]: %s",
			info.ErrorID, rmd4go.ReadCString(info.ErrorMsg[:]),
		)
		c.wait <- false
		return
	}

	if instrument != nil {
		c.t.Logf(
			"market data subscribed: %s.%s",
			rmd4go.ReadCString(instrument.ExchangeID[:]),
			rmd4go.ReadCString(instrument.InstrumentID[:]),
		)
	}

	if isLast {
		c.wait <- true
	}
}

func (c *TestClient) OnRspUnSubMarketData(
	instrument *rmd4go.CRsaFtdcSpecificInstrumentField,
	info *rmd4go.CRsaFtdcRspInfoField,
	requestID int, isLast bool,
) {
	if info.ErrorID != 0 {
		c.t.Logf(
			"un-sub market data failed[%d]: %s",
			info.ErrorID, rmd4go.ReadCString(info.ErrorMsg[:]),
		)
		c.wait <- false
		return
	}

	if instrument != nil {
		c.t.Logf(
			"market data un-subscribed: %s.%s",
			rmd4go.ReadCString(instrument.ExchangeID[:]),
			rmd4go.ReadCString(instrument.InstrumentID[:]),
		)
	}

	if isLast {
		c.wait <- true
	}
}

func (c *TestClient) OnRtnBarMarketData(bar *rmd4go.CRsaFtdcBarMarketDataField) {
	// TODO: bar data print
	c.t.Log("on rtn bar: ")
}

func (c *TestClient) OnRspQryMarketData(
	md *rmd4go.CRsaFtdcRspMarketDataField,
	info *rmd4go.CRsaFtdcRspInfoField,
	requestID int, isLast bool,
) {
	if info.ErrorID != 0 {
		c.t.Logf(
			"query market data failed[%d]: %s",
			info.ErrorID, rmd4go.ReadCString(info.ErrorMsg[:]),
		)
		c.wait <- false
		return
	}

	if md != nil {
		c.t.Logf(
			"market data un-subscribed: %s.%s",
			rmd4go.ReadCString(md.ExchangeID[:]),
			rmd4go.ReadCString(md.InstrumentID[:]),
		)
	}

	if isLast {
		c.wait <- true
	}
}

func (c *TestClient) OnRtnMarketDataEnd(td *rmd4go.CRsaFtdcNtfMarketDataEndField) {
	c.t.Log("on rtn md data end: ")
}

func TestApiSpi(t *testing.T) {
	var (
		err       error
		libPath   = "../dependencies/libs"
		frontAddr = "tcp://212.64.26.127:31121"
		flowPath  = "./flow/"
	)

	switch runtime.GOOS {
	case "linux":
		libPath = filepath.Join(libPath, "rmdapi.so")
	case "windows":
		libPath = filepath.Join(libPath, "rmdapi.dll")
	}

	if info, err := os.Stat(flowPath); err != nil {
		t.Fatal(err)
	} else if !info.IsDir() {
		t.Fatalf("flow path not dir: %s", flowPath)
	}

	api := TestClient{
		t: t,
	}

	if api.CFtdcMdApi, err = rmd4go.CreateFtdcMdApi(
		libPath, frontAddr, flowPath, "rmd-sys lib test", false,
	); err != nil {
		t.Fatal(err)
	}

	api.Init()

	select {
	case <-time.After(time.Second * 10):
		t.Fatal("connect to front timeout")
	case ok := <-api.wait:
		if !ok {
			t.Fatal("connect to front failed")
		}
	}

	if err := api.ReqUserLogin(&rmd4go.CRsaFtdcReqUserLoginField{}); err != nil {
		t.Fatalf("login failed: %+v", err)
	}
}

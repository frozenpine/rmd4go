package rmd4go_test

import (
	"log/slog"
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

func (c *TestClient) OnMdFrontConnected() {
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
			info.ErrorID, info.ErrorMsg,
		)
		c.wait <- false
		return
	}

	if login != nil {
		c.t.Logf(
			"user login: %+v", login,
		)
	}

	c.wait <- true
}

func (c *TestClient) OnRtnDepthMarketData(tick *rmd4go.CRsaFtdcDepthMarketDataField) {
	c.t.Logf("on rtn tick: %+v", tick)
}

func (c *TestClient) OnRspSubMarketData(
	instrument *rmd4go.CRsaFtdcSpecificInstrumentField,
	info *rmd4go.CRsaFtdcRspInfoField,
	requestID int, isLast bool,
) {
	if info.ErrorID != 0 {
		c.t.Logf(
			"sub market data failed[%d]: %s",
			info.ErrorID, info.ErrorMsg,
		)
		c.wait <- false
		return
	}

	if instrument != nil {
		c.t.Logf(
			"market data subscribed: %+v", instrument,
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
			info.ErrorID, info.ErrorMsg,
		)
		c.wait <- false
		return
	}

	if instrument != nil {
		c.t.Logf(
			"market data un-subscribed: %+v", instrument,
		)
	}

	if isLast {
		c.wait <- true
	}
}

func (c *TestClient) OnRtnBarMarketData(bar *rmd4go.CRsaFtdcBarMarketDataField) {
	c.t.Logf("on rtn bar: %+v", bar)
}

func (c *TestClient) OnRspQryMarketData(
	md *rmd4go.CRsaFtdcRspMarketDataField,
	info *rmd4go.CRsaFtdcRspInfoField,
	requestID int, isLast bool,
) {
	if info.ErrorID != 0 {
		c.t.Logf(
			"query market data failed[%d]: %s",
			info.ErrorID, info.ErrorMsg,
		)
		c.wait <- false
		return
	}

	if md != nil {
		c.t.Logf(
			"market data un-subscribed: %+v", md,
		)
	}

	if isLast {
		c.wait <- true
	}
}

func (c *TestClient) OnRtnMarketDataEnd(td *rmd4go.CRsaFtdcNtfMarketDataEndField) {
	c.t.Logf("on rtn md data end: %+v", td)
}

func TestApiSpi(t *testing.T) {
	slog.SetLogLoggerLevel(slog.LevelDebug)
	var (
		err       error
		libPath   = "./dependencies/libs"
		frontAddr = "tcp://172.16.200.105:30010"
		flowPath  = "./flow/"
	)

	switch runtime.GOOS {
	case "linux":
		libPath = filepath.Join(libPath, "rmdapi.so")
	case "windows":
		libPath = filepath.Join(libPath, "rmdapi.dll")
	}

	t.Log(libPath)

	if info, err := os.Stat(flowPath); err != nil {
		if err = os.MkdirAll(flowPath, os.ModePerm); err != nil {
			t.Fatal("make flow path failed")
		}
	} else if !info.IsDir() {
		t.Fatalf("flow path not dir: %s", flowPath)
	} else {
		t.Log("flow path ok")
	}

	api := TestClient{
		t:    t,
		wait: make(chan bool),
	}

	t.Log("creating rmd api")

	if api.CFtdcMdApi, err = rmd4go.CreateFtdcMdApi(
		libPath, frontAddr, flowPath, "rmd-sys lib test", false,
	); err != nil {
		t.Fatal(err)
	}

	api.RegisterSpi(&api)

	t.Log("rmd spi registered")

	api.Init()

	t.Log("rmd api initiated")

	select {
	case <-time.After(time.Second * 10):
		t.Fatal("connect to front timeout")
	case ok := <-api.wait:
		if !ok {
			t.Fatal("connect to front failed")
		}
	}

	if err := api.ReqUserLogin(&rmd4go.CRsaFtdcReqUserLoginField{
		BrokerID: "rdrk",
		UserID:   "test",
		Password: "test",
	}); err != nil {
		t.Fatalf("req login failed: %+v", err)
	} else {
		<-api.wait
	}

	api.ReqBtSubMarketData(&rmd4go.CRsaFtdcBtSubMarketDataField{
		ExchangeID:   "CFFEX",
		InstrumentID: "IC2512",
	}, &rmd4go.CRsaFtdcBtSubMarketDataField{
		ExchangeID:   "CFFEX",
		InstrumentID: "IM2512",
	})
	<-api.wait

	<-time.After(time.Second * 20)
}

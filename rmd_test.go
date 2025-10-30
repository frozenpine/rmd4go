package rmd4go_test

import (
	"context"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/frozenpine/rmd4go"
)

type TestSpi struct{}

func (c TestSpi) OnMdFrontConnected() {
	slog.Info("front connected")
}

func (c TestSpi) OnMdFrontDisconnected(reason int) {
	slog.Info("front disconnected", slog.Int("reason", reason))
}

func (c TestSpi) OnMdRspUserLogin(
	login *rmd4go.CRsaFtdcRspUserLoginField,
	info *rmd4go.CRsaFtdcRspInfoField,
	requestID int, isLast bool,
) {
	if info.ErrorID != 0 {
		slog.Error(
			"user login failed",
			slog.Any("info", info),
		)

		return
	}

	if login != nil {
		slog.Info(
			"user login",
			slog.Any("login", login),
			slog.Int("req", requestID),
			slog.Bool("last", isLast),
		)
	}
}

func (c TestSpi) OnRtnDepthMarketData(tick *rmd4go.CRsaFtdcDepthMarketDataField) {
	slog.Info("on rtn tick", slog.Any("tick", tick))
}

func (c TestSpi) OnRspSubMarketData(
	instrument *rmd4go.CRsaFtdcSpecificInstrumentField,
	info *rmd4go.CRsaFtdcRspInfoField,
	requestID int, isLast bool,
) {
	if info.ErrorID != 0 {
		slog.Error(
			"sub market data failed",
			slog.Any("info", info),
		)

		return
	}

	if instrument != nil {
		slog.Info(
			"market data subscribed",
			slog.Any("ins", instrument),
			slog.Int("req", requestID),
			slog.Bool("last", isLast),
		)
	}
}

func (c TestSpi) OnRspUnSubMarketData(
	instrument *rmd4go.CRsaFtdcSpecificInstrumentField,
	info *rmd4go.CRsaFtdcRspInfoField,
	requestID int, isLast bool,
) {
	if info.ErrorID != 0 {
		slog.Error(
			"un-sub market data failed",
			slog.Any("info", info),
		)

		return
	}

	if instrument != nil {
		slog.Info(
			"market data un-subscribed:",
			slog.Any("ins", instrument),
			slog.Int("req", requestID),
		)
	}
}

func (c TestSpi) OnRtnBarMarketData(bar *rmd4go.CRsaFtdcBarMarketDataField) {
	slog.Info("on rtn bar", slog.Any("bar", bar))
}

func (c TestSpi) OnRspQryMarketData(
	md *rmd4go.CRsaFtdcRspMarketDataField,
	info *rmd4go.CRsaFtdcRspInfoField,
	requestID int, isLast bool,
) {
	if info.ErrorID != 0 {
		slog.Error(
			"query market data failed",
			slog.Any("info", info),
		)

		return
	}

	if md != nil {
		slog.Info(
			"market data un-subscribed",
			slog.Any("md", md),
			slog.Int("req", requestID),
			slog.Bool("last", isLast),
		)
	}
}

func (c TestSpi) OnRtnMarketDataEnd(dt *rmd4go.CRsaFtdcNtfMarketDataEndField) {
	slog.Info("on rtn md data end", slog.Any("date", dt))
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

	spi := TestSpi{}

	t.Log("creating rmd api")

	api, err := rmd4go.CreateFtdcMdApi(
		libPath, frontAddr, flowPath, "rmd-sys lib test", false,
	)
	if err != nil {
		t.Fatal(err)
	}

	api.RegisterSpi(spi)

	t.Log("rmd spi registered")

	api.Init()

	t.Log("rmd api initiated")

	<-time.After(time.Second)

	if err := api.ReqUserLogin(&rmd4go.CRsaFtdcReqUserLoginField{
		BrokerID: "rdrk",
		UserID:   "test",
		Password: "test",
	}); err != nil {
		t.Fatalf("req login failed: %+v", err)
	}

	<-time.After(time.Second)

	if err := api.ReqBtSubMarketData(&rmd4go.CRsaFtdcBtSubMarketDataField{
		ExchangeID:   "CFFEX",
		InstrumentID: "IC2512",
		BarPreces:    rmd4go.BarPrecesMinute,
		BarPeriod:    1,
	}, &rmd4go.CRsaFtdcBtSubMarketDataField{
		ExchangeID:   "CFFEX",
		InstrumentID: "IM2512",
	}); err != nil {
		t.Fatal(err)
	}

	<-time.After(time.Second)

	if data, err := api.ReqQryBarMarketData(&rmd4go.CRsaFtdcBtSubMarketDataField{
		ExchangeID:   "CFFEX",
		InstrumentID: "IC2512",
		BarPreces:    rmd4go.BarPrecesMinute,
		BarPeriod:    1,
	}, 1000); err != nil {
		t.Fatal(err)
	} else {
		for _, b := range data {
			t.Logf("%+v", b)
		}

		t.Log("count", len(data))
	}
}

func TestChannelSpi(t *testing.T) {
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

	api, err := rmd4go.CreateFtdcMdApi(
		libPath, frontAddr, flowPath, "rmd-sys lib test", false,
	)

	if err != nil {
		t.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(t.Context(), time.Second*20)
	defer cancel()

	spi := rmd4go.NewChannelSpi[TestSpi](ctx, 100)

	api.RegisterSpi(spi)

	api.Init()

	<-time.After(time.Second)

	if err := api.ReqUserLogin(
		&rmd4go.CRsaFtdcReqUserLoginField{},
	); err != nil {
		t.Fatal(err)
	}

	<-time.After(time.Second)

	if err := api.ReqBtSubMarketData(
		&rmd4go.CRsaFtdcBtSubMarketDataField{
			ExchangeID:   "CFFEX",
			InstrumentID: "IC2512",
		},
	); err != nil {
		t.Fatal(err)
	}

	<-time.After(time.Second)

	if data, err := api.ReqQryBarMarketData(
		&rmd4go.CRsaFtdcBtSubMarketDataField{
			ExchangeID:   "CFFEX",
			InstrumentID: "IC2512",
			BarPreces:    rmd4go.BarPrecesMinute,
			BarPeriod:    1,
		}, 1000,
	); err != nil {
		t.Fatal(err)
	} else {
		for _, v := range data {
			t.Logf("%+v", v)
		}

		t.Log(len(data))
	}

	spi.Close()

	spi.Join()
}

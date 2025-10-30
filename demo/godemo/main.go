package main

import (
	"context"
	"flag"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/frozenpine/rmd4go"
)

type TestSpi struct {
	wait chan bool
}

func (c *TestSpi) OnMdFrontConnected() {
	slog.Info("front connected")
	if c.wait != nil {
		c.wait <- true
	}
}

func (c *TestSpi) OnMdFrontDisconnected(reason int) {
	slog.Info("front disconnected", slog.Int("reason", reason))
	if c.wait != nil {
		c.wait <- false
	}
}

func (c *TestSpi) OnMdRspUserLogin(
	login *rmd4go.CRsaFtdcRspUserLoginField,
	info *rmd4go.CRsaFtdcRspInfoField,
	requestID int, isLast bool,
) {
	if info != nil && info.ErrorID != 0 {
		slog.Error(
			"user login failed",
			slog.Any("info", info),
		)

		if c.wait != nil {
			c.wait <- false
		}

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

	if isLast && c.wait != nil {
		c.wait <- true
	}
}

func (c *TestSpi) OnRtnDepthMarketData(tick *rmd4go.CRsaFtdcDepthMarketDataField) {
	slog.Info("on rtn tick", slog.Any("tick", tick))
}

func (c *TestSpi) OnRspSubMarketData(
	instrument *rmd4go.CRsaFtdcSpecificInstrumentField,
	info *rmd4go.CRsaFtdcRspInfoField,
	requestID int, isLast bool,
) {
	if info != nil && info.ErrorID != 0 {
		slog.Error(
			"sub market data failed",
			slog.Any("info", info),
		)

		if c.wait != nil {
			c.wait <- false
		}

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

	if isLast && c.wait != nil {
		c.wait <- true
	}
}

func (c *TestSpi) OnRspUnSubMarketData(
	instrument *rmd4go.CRsaFtdcSpecificInstrumentField,
	info *rmd4go.CRsaFtdcRspInfoField,
	requestID int, isLast bool,
) {
	if info != nil && info.ErrorID != 0 {
		slog.Error(
			"un-sub market data failed",
			slog.Any("info", info),
		)

		if c.wait != nil {
			c.wait <- false
		}

		return
	}

	if instrument != nil {
		slog.Info(
			"market data un-subscribed:",
			slog.Any("ins", instrument),
			slog.Int("req", requestID),
		)
	}

	if isLast && c.wait != nil {
		c.wait <- true
	}
}

func (c *TestSpi) OnRtnBarMarketData(bar *rmd4go.CRsaFtdcBarMarketDataField) {
	slog.Info("on rtn bar", slog.Any("bar", bar))
}

func (c *TestSpi) OnRspQryMarketData(
	md *rmd4go.CRsaFtdcRspMarketDataField,
	info *rmd4go.CRsaFtdcRspInfoField,
	requestID int, isLast bool,
) {
	if info.ErrorID != 0 {
		slog.Error(
			"query market data failed",
			slog.Any("info", info),
		)

		if c.wait != nil {
			c.wait <- false
		}

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

	if isLast && c.wait != nil {
		c.wait <- true
	}
}

func (c *TestSpi) OnRtnMarketDataEnd(dt *rmd4go.CRsaFtdcNtfMarketDataEndField) {
	slog.Info("on rtn md data end", slog.Any("date", dt))
}

var (
	libPath, flowPath, remoteConn string
)

func init() {
	flag.StringVar(&libPath, "lib", "", "RMD lib DIR path")
	flag.StringVar(&flowPath, "flow", "./flow/", "RMD lib flow path")
	flag.StringVar(&remoteConn, "remote", "", "RMD remote conn string")
}

func main() {
	if !flag.Parsed() {
		flag.Parse()
	}

	switch runtime.GOOS {
	case "linux":
		libPath = filepath.Join(libPath, "rmdapi.so")
	case "windows":
		libPath = filepath.Join(libPath, "rmdapi.dll")
	}

	if info, err := os.Stat(flowPath); err != nil {
		if err = os.MkdirAll(flowPath, os.ModePerm); err != nil {
			slog.Error("make flow path failed", slog.String("flow", flowPath))
			os.Exit(1)
		}
	} else if !info.IsDir() {
		slog.Error("flow path not dir", slog.String("flow", flowPath))
		os.Exit(1)
	} else {
		slog.Info("flow path check ok")
	}

	api, err := rmd4go.CreateFtdcMdApi(
		libPath, remoteConn, flowPath, "rmd-sys lib test", false,
	)
	if err != nil {
		slog.Error("create api instance failed", slog.Any("error", err))
		os.Exit(1)
	}
	defer api.Release()

	ctx := context.Background()

	wait := make(chan bool)

	spi, err := rmd4go.NewChannelSpi[*TestSpi](
		ctx, 100,
		func(ts **TestSpi) error {
			*ts = &TestSpi{wait: wait}
			return nil
		},
	)
	if err != nil {
		slog.Error("create spi instance failed", slog.Any("error", err))
		os.Exit(1)
	}
	defer spi.Close()

	api.RegisterSpi(spi)

	api.Init()

	connStart := time.Now()
	select {
	case ts := <-time.After(time.Minute):
		slog.Error(
			"connect to front timeouted",
			slog.Duration("timeout", ts.Sub(connStart)),
		)
		os.Exit(1)
	case flag := <-wait:
		if !flag {
			slog.Error("connect")
		}
	case <-ctx.Done():
		return
	}
}

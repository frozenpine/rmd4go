package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/frozenpine/rmd4go"
	"github.com/valyala/bytebufferpool"
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

type subList []*rmd4go.CRsaFtdcBtSubMarketDataField

func (s subList) String() string {
	buff := bytebufferpool.Get()
	defer bytebufferpool.Put(buff)

	buff.WriteString("Subs[")
	for idx, v := range s {
		if idx > 0 {
			buff.WriteString(", ")
		}
		fmt.Fprintf(
			buff, "%s.%s@%d%s",
			v.ExchangeID, v.InstrumentID, v.BarPeriod, v.BarPreces,
		)
	}
	buff.WriteByte(']')

	return buff.String()
}

func (s *subList) Set(v string) error {
	identities := strings.SplitN(v, "@", 2)
	if len(identities) != 2 {
		return errors.New("invalid sub identity")
	}
	symbols := strings.SplitN(identities[0], ".", 2)
	if len(symbols) != 2 {
		return errors.New("invalid symbol")
	}

	period, err := strconv.Atoi(string(identities[1][0]))
	if err != nil {
		return errors.Join(err, errors.New("invalid period"))
	}
	idx := strings.Index("smd", string(identities[1][1]))
	if idx < 0 {
		return errors.New("invalid preces")
	}
	preces := rmd4go.TRsaFtdcBarPrecesType(idx) + 48

	*s = append(*s, &rmd4go.CRsaFtdcBtSubMarketDataField{
		ExchangeID:   symbols[0],
		InstrumentID: symbols[1],
		BarPreces:    preces,
		BarPeriod:    period,
	})

	return nil
}

var (
	version, goVersion, gitVersion, buildTime string

	libPath, flowPath, remoteConn string
	subs                          subList
	debug                         int = -1
)

func init() {
	flag.StringVar(&libPath, "lib", "", "RMD lib DIR path")
	flag.StringVar(&flowPath, "flow", "./flow/", "RMD lib flow path")
	flag.StringVar(&remoteConn, "remote", "", "RMD remote conn string")

	flag.BoolFunc("v", "Set verbose level", func(_ string) error {
		debug += 1
		return nil
	})

	flag.Var(&subs, "sub", "Subscribe list")

	flag.BoolFunc("ver", "Show version info", func(_ string) error {
		fmt.Printf(
			"rmd4go version: %s build by %s\nGit version: %s, Build time: %s",
			version, goVersion, gitVersion, buildTime,
		)

		os.Exit(0)

		return nil
	})
}

func waitRsp(ctx context.Context, wait <-chan bool, timeout time.Duration) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case v, ok := <-wait:
		if !ok || !v {
			return errors.New("wait failed")
		}
	case <-time.After(timeout):
		return errors.New("wait rsp timeouted")
	}

	return nil
}

func main() {
	if !flag.Parsed() {
		flag.Parse()
	}

	if libPath == "" || flowPath == "" || remoteConn == "" {
		slog.Error("invalid args for connection")
		os.Exit(1)
	}

	if len(subs) <= 0 {
		slog.Error("no subs specified")
		os.Exit(1)
	}

	if debug > 0 {
		slog.SetLogLoggerLevel(slog.LevelDebug - slog.Level(debug))
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

	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGABRT, syscall.SIGINT, syscall.SIGHUP, syscall.SIGKILL,
	)
	defer cancel()

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

	if err := waitRsp(ctx, wait, time.Second*10); err != nil {
		slog.Error(
			"wait front connected failed",
			slog.Any("error", err),
		)
		os.Exit(1)
	}

	if err := api.ReqUserLogin(
		&rmd4go.CRsaFtdcReqUserLoginField{},
	); err != nil {
		slog.Error(
			"req user login failed",
			slog.Any("error", err),
		)
		os.Exit(1)
	} else if err = waitRsp(ctx, wait, time.Second*10); err != nil {
		slog.Error(
			"wait rsp login failed",
			slog.Any("error", err),
		)
	}

	if err := api.ReqBtSubMarketData(subs...); err != nil {
		slog.Error(
			"req sub market data failed",
			slog.Any("error", err),
		)
	} else if err = waitRsp(ctx, wait, time.Minute); err != nil {
		slog.Error(
			"wait rsp sub market data failed",
			slog.Any("error", err),
		)
	}

	for _, sub := range subs {
		if data, err := api.ReqQryBarMarketData(sub, 1000); err != nil {
			slog.Error(
				"req query market data failed",
				slog.Any("error", err),
				slog.Any("sub", sub),
			)

			os.Exit(1)
		} else {
			for _, v := range data {
				slog.Info(
					"market data queried for sub",
					slog.Any("bar", v),
				)
			}
		}
	}

	spi.Join()
}

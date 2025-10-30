package rmd4go

import (
	"errors"
	"log/slog"
	"runtime"
	"sync"
)

type EventType int

const (
	EventOnMdFrontConnected EventType = iota
	EventOnMdFrontDisconnected
	EventOnMdRspUserLogin
	EventOnRtnDepthMarketData
	EventOnRspSubMarketData
	EventOnRspUnSubMarketData
	EventOnRtnBarMarketData
	EventOnRspQryMarketData
	EventOnRtnMarketDataEnd
)

type Response struct {
	Info   *CRsaFtdcRspInfoField
	ReqId  int
	IsLast bool
}

type Event struct {
	evtType  EventType
	data     any
	response *Response
}

var (
	eventPool = sync.Pool{New: func() any {
		return new(Event)
	}}
)

func TryMakeEvent(attemps int) (*Event, error) {
	count := 0

	for {
		if evt := eventPool.Get(); evt != nil {
			runtime.SetFinalizer(evt, func(v any) {
				runtime.SetFinalizer(v, nil)
				eventPool.Put(v)
			})

			if rtn, ok := evt.(*Event); ok {
				return rtn, nil
			}
		} else {
			count++

			if count >= attemps {
				return nil, errors.New("max attemps execeeded")
			} else {
				slog.Warn(
					"get event from pool failed",
					slog.Int("fail_count", count),
				)
			}
		}
	}
}

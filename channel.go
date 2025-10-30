package rmd4go

import (
	"context"
	"log/slog"
	"sync/atomic"
)

// ChannelSpi 解耦Cpp回调和Go接口事件处理，以避免慢处理阻塞Cpp回调
type ChannelSpi[T RmdSpi] struct {
	callback   T
	channel    chan *Event
	chanClosed atomic.Bool
	eventDone  chan struct{}

	ctx    context.Context
	cancel context.CancelFunc
}

func (spi *ChannelSpi[T]) OnMdFrontConnected() {
	if evt, err := TryMakeEvent(10); err != nil {
		slog.Error(
			"can not get event from pool OnMdFrontConnected",
			slog.Any("error", err),
		)
	} else {
		evt.evtType = EventOnMdFrontConnected

		spi.sendEvent(evt)
	}
}

func (spi *ChannelSpi[T]) OnMdFrontDisconnected(reason int) {
	if evt, err := TryMakeEvent(10); err != nil {
		slog.Error(
			"can not get event from pool OnMdFrontDisconnected",
			slog.Any("error", err),
			slog.Int("reason", reason),
		)
	} else {
		evt.evtType = EventOnMdFrontDisconnected
		evt.data = reason

		spi.sendEvent(evt)
	}
}

func (spi *ChannelSpi[T]) OnMdRspUserLogin(
	login *CRsaFtdcRspUserLoginField,
	info *CRsaFtdcRspInfoField,
	requestID int, isLast bool,
) {
	if evt, err := TryMakeEvent(10); err != nil {
		slog.Error(
			"can not get event from pool OnMdRspUserLogin",
			slog.Any("error", err),
			slog.Any("login", login),
			slog.Any("info", info),
			slog.Int("req", requestID),
			slog.Bool("is_last", isLast),
		)
	} else {
		evt.evtType = EventOnMdRspUserLogin
		evt.data = login

		if evt.response == nil {
			evt.response = &Response{}
		}
		evt.response.Info = info
		evt.response.ReqId = requestID
		evt.response.IsLast = isLast

		spi.sendEvent(evt)
	}
}

func (spi *ChannelSpi[T]) OnRtnDepthMarketData(
	tick *CRsaFtdcDepthMarketDataField,
) {
	if evt, err := TryMakeEvent(10); err != nil {
		slog.Error(
			"can not get event from pool OnRtnDepthMarketData",
			slog.Any("error", err),
			slog.Any("tick", tick),
		)
	} else {
		evt.evtType = EventOnRtnDepthMarketData
		evt.data = tick

		spi.sendEvent(evt)
	}
}

func (spi *ChannelSpi[T]) OnRspSubMarketData(
	instrument *CRsaFtdcSpecificInstrumentField,
	info *CRsaFtdcRspInfoField,
	requestID int, isLast bool,
) {
	if evt, err := TryMakeEvent(10); err != nil {
		slog.Error(
			"can not get event from pool OnRspSubMarketData",
			slog.Any("error", err),
			slog.Any("instrument", instrument),
			slog.Any("info", info),
			slog.Int("req", requestID),
			slog.Bool("is_last", isLast),
		)
	} else {
		evt.evtType = EventOnRspSubMarketData
		evt.data = instrument

		if evt.response == nil {
			evt.response = &Response{}
		}
		evt.response.Info = info
		evt.response.ReqId = requestID
		evt.response.IsLast = isLast

		spi.sendEvent(evt)
	}
}

func (spi *ChannelSpi[T]) OnRspUnSubMarketData(
	instrument *CRsaFtdcSpecificInstrumentField,
	info *CRsaFtdcRspInfoField,
	requestID int, isLast bool,
) {
	if evt, err := TryMakeEvent(10); err != nil {
		slog.Error(
			"can not get event from pool OnRspUnSubMarketData",
			slog.Any("error", err),
			slog.Any("instrument", instrument),
			slog.Any("info", info),
			slog.Int("req", requestID),
			slog.Bool("is_last", isLast),
		)
	} else {
		evt.evtType = EventOnRspSubMarketData
		evt.data = instrument

		if evt.response == nil {
			evt.response = &Response{}
		}
		evt.response.Info = info
		evt.response.ReqId = requestID
		evt.response.IsLast = isLast

		spi.sendEvent(evt)
	}
}

func (spi *ChannelSpi[T]) OnRtnBarMarketData(
	bar *CRsaFtdcBarMarketDataField,
) {
	if evt, err := TryMakeEvent(10); err != nil {
		slog.Error(
			"can not get event from pool OnRtnBarMarketData",
			slog.Any("error", err),
			slog.Any("bar", bar),
		)
	} else {
		evt.evtType = EventOnRspSubMarketData
		evt.data = bar

		spi.sendEvent(evt)
	}
}

func (spi *ChannelSpi[T]) OnRspQryMarketData(
	md *CRsaFtdcRspMarketDataField,
	info *CRsaFtdcRspInfoField,
	requestID int, isLast bool,
) {
	if evt, err := TryMakeEvent(10); err != nil {
		slog.Error(
			"can not get event from pool OnRspUnSubMarketData",
			slog.Any("error", err),
			slog.Any("md", md),
			slog.Any("info", info),
			slog.Int("req", requestID),
			slog.Bool("is_last", isLast),
		)
	} else {
		evt.evtType = EventOnRspSubMarketData
		evt.data = md

		if evt.response == nil {
			evt.response = &Response{}
		}
		evt.response.Info = info
		evt.response.ReqId = requestID
		evt.response.IsLast = isLast

		spi.sendEvent(evt)
	}
}

func (spi *ChannelSpi[T]) OnRtnMarketDataEnd(
	dt *CRsaFtdcNtfMarketDataEndField,
) {
	if evt, err := TryMakeEvent(10); err != nil {
		slog.Error(
			"can not get event from pool OnRtnMarketDataEnd",
			slog.Any("error", err),
			slog.Any("date", dt),
		)
	} else {
		evt.evtType = EventOnRspSubMarketData
		evt.data = dt

		spi.sendEvent(evt)
	}
}

func (spi *ChannelSpi[T]) sendEvent(evt *Event) {
	if evt == nil {
		slog.Log(
			context.Background(), slog.LevelDebug-1,
			"channel spi send nil event",
		)
		return
	}

	if spi.chanClosed.Load() {
		slog.Log(
			context.Background(), slog.LevelDebug-1,
			"channel spi send event when channel closed",
			slog.Any("event", evt),
		)
		return
	}

	spi.channel <- evt
}

func (spi *ChannelSpi[T]) dispatcher() {
	slog.Info(
		"starting channel spi event dispatcher",
	)

	defer close(spi.eventDone)

	for {
		select {
		case <-spi.ctx.Done():
			if !spi.chanClosed.Load() {
				slog.Info(
					"channel spi context exited",
				)

				spi.chanClosed.Store(true)
				close(spi.channel)
			}
		case evt, ok := <-spi.channel:
			if !ok {
				slog.Info(
					"all channeled events done",
				)
				return
			}

			switch evt.evtType {
			case EventOnMdFrontConnected:
				spi.callback.OnMdFrontConnected()
			case EventOnMdFrontDisconnected:
				spi.callback.OnMdFrontDisconnected(evt.data.(int))
			case EventOnMdRspUserLogin:
				spi.callback.OnMdRspUserLogin(
					evt.data.(*CRsaFtdcRspUserLoginField),
					evt.response.Info,
					evt.response.ReqId,
					evt.response.IsLast,
				)
			case EventOnRtnDepthMarketData:
				spi.callback.OnRtnDepthMarketData(
					evt.data.(*CRsaFtdcDepthMarketDataField),
				)
			case EventOnRspSubMarketData:
				spi.callback.OnRspSubMarketData(
					evt.data.(*CRsaFtdcSpecificInstrumentField),
					evt.response.Info,
					evt.response.ReqId,
					evt.response.IsLast,
				)
			case EventOnRspUnSubMarketData:
				spi.callback.OnRspUnSubMarketData(
					evt.data.(*CRsaFtdcSpecificInstrumentField),
					evt.response.Info,
					evt.response.ReqId,
					evt.response.IsLast,
				)
			case EventOnRtnBarMarketData:
				spi.callback.OnRtnBarMarketData(
					evt.data.(*CRsaFtdcBarMarketDataField),
				)
			case EventOnRspQryMarketData:
				spi.callback.OnRspQryMarketData(
					evt.data.(*CRsaFtdcRspMarketDataField),
					evt.response.Info,
					evt.response.ReqId,
					evt.response.IsLast,
				)
			case EventOnRtnMarketDataEnd:
				spi.callback.OnRtnMarketDataEnd(
					evt.data.(*CRsaFtdcNtfMarketDataEndField),
				)
			default:
				slog.Error(
					"unknown event when dispatching",
					slog.Any("event", evt),
				)
			}
		}
	}
}

func (spi *ChannelSpi[T]) Close() {
	spi.cancel()
}

func (spi *ChannelSpi[T]) Join() {
	<-spi.ctx.Done()

	slog.Info("waiting all events dispatched")
	<-spi.eventDone
	slog.Info("all events dispached, exit join")
}

func NewChannelSpi[T RmdSpi](ctx context.Context, buffSize int) *ChannelSpi[T] {
	if ctx == nil {
		ctx = context.Background()
	}

	if buffSize <= 0 {
		buffSize = 100
	}

	var inner T

	spi := ChannelSpi[T]{
		callback:  inner,
		channel:   make(chan *Event, buffSize),
		eventDone: make(chan struct{}),
	}

	spi.ctx, spi.cancel = context.WithCancel(ctx)

	go spi.dispatcher()

	return &spi
}

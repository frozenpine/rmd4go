package rmd4go

// 响应信息
type CRsaFtdcRspInfoField struct {
	// 错误代码
	ErrorID int
	// 错误信息
	ErrorMsg string
}

// 系统用户登录请求
type CRsaFtdcReqUserLoginField struct {
	// 交易日
	TradingDay string
	// 交易用户代码
	UserID string
	// 经纪公司编号
	BrokerID string
	// 密码
	Password string
	// 用户端产品信息
	UserProductInfo string
	// 接口端产品信息
	InterfaceProductInfo string
	// 协议信息
	ProtocolInfo string
	// IP地址
	IPAddress string
	// Mac地址
	MacAddress string
	// 数据中心代码
	DataCenterID int
	// 动态密码
	OneTimePassword string
	// 终端IP地址
	ClientIPAddress string
	// 授权编码
	AuthCode string
	// 自定义字段
	FieldContent string
	// 客户端信息采集
	ClientSystemInfo string
}

// 系统用户登录应答
type CRsaFtdcRspUserLoginField struct {
	// 交易日
	TradingDay string
	// 经纪公司编号
	BrokerID string
	// 交易用户代码
	UserID string
	// 登录成功时间
	LoginTime string
	// 用户最大本地报单号
	MaxOrderLocalID string
	// 交易系统名称
	TradingSystemName string
	// 数据中心代码
	DataCenterID int
	// 会员私有流当前长度
	PrivateFlowSize int
	// 交易员私有流当前长度
	UserFlowSize int
	// 登录附加信息
	LoginInfo string
	// 会话编号
	SessionID int
	// 前置编号
	FrontID int
}

// 用户登出请求
type CRsaFtdcReqUserLogoutField struct {
	// 经纪公司编号
	BrokerID string
	// 交易用户代码
	UserID string
}

// 用户登出请求
type CRsaFtdcRspUserLogoutField struct {
	// 经纪公司编号
	BrokerID string
	// 交易用户代码
	UserID string
}

// 强制用户退出
type CRsaFtdcForceUserExitField struct {
	// 经纪公司编号
	BrokerID string
	// 交易用户代码
	UserID string
}

// 用户口令修改
type CRsaFtdcUserPasswordUpdateField struct {
	// 经纪公司编号
	BrokerID string
	// 交易用户代码
	UserID string
	// 旧密码
	OldPassword string
	// 新密码
	NewPassword string
}

// 输入报单
type CRsaFtdcInputOrderField struct {
	// 经纪公司编号
	BrokerID string
	// 投资者编号
	InvestorID string
	// 用户代码
	UserID string
	// 用户本地报单号
	UserOrderLocalID string
	// 用户自定义域
	UserCustom string
	// qsc全局编号
	QscId int
	// 插件标识
	PluginTaskId int
	// 外部标识
	OthMarkId string
	// 交易所代码
	ExchangeID string
	// 系统报单编号
	OrderSysID string
	// 合约代码
	InstrumentID string
	// 报单类型
	OrderPriceType TRsaFtdcOrderPriceTypeType
	// 买卖方向
	Direction TRsaFtdcDirectionType
	// 开平标志
	OffsetFlag TRsaFtdcOffsetFlagType
	// 投机套保标志
	HedgeFlag TRsaFtdcHedgeFlagType
	// 价格
	LimitPrice float64
	// 数量
	Volume float64
	// 有效期类型
	TimeCondition TRsaFtdcTimeConditionType
	// GTD日期
	GTDDate string
	// 成交量类型
	VolumeCondition TRsaFtdcVolumeConditionType
	// 最小成交量
	MinVolume float64
	// 止损价
	StopPrice float64
	// 强平原因
	ForceCloseReason TRsaFtdcForceCloseReasonType
	// 自动挂起标志
	IsAutoSuspend bool
	// 业务单元
	BusinessUnit string
	// 业务类别
	BusinessType TRsaFtdcBusinessTypeType
	// 数量（小数）
	VolumeP float64
	// 策略限制价格
	AlgoLimitPrice float64
	// 策略限制tick数量
	AlgoLimitPriceTick int
	// 策略限制时间(ms)
	AlgoLimitTime int
}

// 报单操作
type CRsaFtdcOrderActionField struct {
	// 经纪公司编号
	BrokerID string
	// 投资者编号
	InvestorID string
	// 用户代码
	UserID string
	// 用户本地报单号
	UserOrderLocalID string
	// 用户自定义域
	UserCustom string
	// qsc全局编号
	QscId int
	// 插件标识
	PluginTaskId int
	// 外部标识
	OthMarkId string
	// 交易所代码
	ExchangeID string
	// 报单编号
	OrderSysID string
	// 合约代码
	InstrumentID string
	// 本次撤单操作的本地编号
	UserOrderActionLocalID string
	// 报单操作标志
	ActionFlag TRsaFtdcActionFlagType
	// 价格
	LimitPrice float64
	// 数量变化
	VolumeChange float64
	// 前置编号
	FrontID int
	// 会话编号
	SessionID int
}

// 内存表导出
type CRsaFtdcMemDBField struct {
	// 内存表名
	MemTableName TRsaFtdcMemTableNameType
}

// 投资者资金帐户出入金请求
type CRsaFtdcReqAccountDepositField struct {
	// 经纪公司编号
	BrokerID string
	// 资金流水号
	AccountSeqNo TRsaFtdcAccountSeqNoType
	// 资金帐号
	AccountID string
	// 投资者编号
	InvestorID string
	// 金额
	Amount float64
	// 出入金方向
	AmountDirection TRsaFtdcAccountDirectionType
	// 用户代码
	UserID string
}

// 投资者资金帐户出入金应答
type CRsaFtdcRspAccountDepositField struct {
	// 经纪公司编号
	BrokerID string
	// 投资者编号
	InvestorID string
	// 资金帐号
	AccountID string
	// 资金流水号
	AccountSeqNo TRsaFtdcAccountSeqNoType
	// 金额
	Amount float64
	// 出入金方向
	AmountDirection TRsaFtdcAccountDirectionType
	// 可用资金
	Available float64
	// 结算准备金
	Balance float64
	// 用户代码
	UserID string
}

// 实时上场
type CRsaFtdcTBCommandField struct {
	// DB命令序号
	CommandNo TRsaFtdcCommandNoType
	// 操作日期
	CommandDate string
	// 操作时间
	CommandTime string
	// DB命令类型
	CommandType TRsaFtdcCommandTypeType
	// 字段名
	FieldName TRsaFtdcFieldNameType
	// 字段内容
	FieldContent string
}

// 出入金结果
type CRsaFtdcInvestorAccountDepositResField struct {
	// 经纪公司编号
	BrokerID string
	// 投资者编号
	InvestorID string
	// 资金帐号
	AccountID string
	// 资金流水号
	AccountSeqNo TRsaFtdcAccountSeqNoType
	// 金额
	Amount float64
	// 出入金方向
	AmountDirection TRsaFtdcAccountDirectionType
	// 可用资金
	Available float64
	// 结算准备金
	Balance float64
	// 用户代码
	UserID string
}

// 合约状态
type CRsaFtdcInstrumentStatusField struct {
	// 交易所代码
	ExchangeID string
	// 品种代码
	ProductID TRsaFtdcProductIDType
	// 品种名称
	ProductName TRsaFtdcProductNameType
	// 合约代码
	InstrumentID string
	// 合约名称
	InstrumentName TRsaFtdcInstrumentNameType
	// 交割年份
	DeliveryYear TRsaFtdcYearType
	// 交割月
	DeliveryMonth TRsaFtdcMonthType
	// 限价单最大下单量
	MaxLimitOrderVolume float64
	// 限价单最小下单量
	MinLimitOrderVolume float64
	// 市价单最大下单量
	MaxMarketOrderVolume float64
	// 市价单最小下单量
	MinMarketOrderVolume float64
	// 数量乘数
	VolumeMultiple TRsaFtdcVolumeMultipleType
	// 报价单位
	PriceTick TRsaFtdcPriceTickType
	// 币种
	Currency string
	// 多头限仓
	LongPosLimit float64
	// 空头限仓
	ShortPosLimit float64
	// 跌停板价
	LowerLimitPrice float64
	// 涨停板价
	UpperLimitPrice float64
	// 昨结算
	PreSettlementPrice float64
	// 合约交易状态
	InstrumentStatus TRsaFtdcInstrumentStatusType
	// 创建日
	CreateDate string
	// 上市日
	OpenDate string
	// 到期日
	ExpireDate string
	// 开始交割日
	StartDelivDate string
	// 最后交割日
	EndDelivDate string
	// 挂牌基准价
	BasisPrice float64
	// 当前是否交易
	IsTrading bool
	// 基础商品代码
	UnderlyingInstrID string
	// 基础商品乘数
	UnderlyingMultiple float64
	// 持仓类型
	PositionType TRsaFtdcPositionTypeType
	// 执行价
	StrikePrice float64
	// 期权类型
	OptionsType TRsaFtdcOptionsTypeType
	// 汇率
	ExchangeRate float64
	// 产品类型
	ProductClass TRsaFtdcProductClassType
	// 期权行权方式
	OptionsMode TRsaFtdcOptionsModeType
	// 保证金算法类型
	MarginCombType TRsaFtdcMarginCombTypeType
	// 多头保证金率
	LongMarginRate float64
	// 空头保证金率
	ShortMarginRate float64
	// 多头保证金按手数
	LongMarginAmt float64
	// 空头保证金按手数
	ShortMarginAmt float64
}

// 审批模式设置请求
type CRsaFtdcAdviceRunModeField struct {
	// 经纪公司编号
	BrokerID string
	// 投资者编号
	InvestorID string
	// 用户代码
	UserID string
	// 自动或者手动
	RunMode TRsaFtdcAdviceRunModeType
}

// 订阅合约状态
type CRsaFtdcSubInstrumentStatusField struct {
	// 交易所代码
	ExchangeID string
	// 合约代码
	InstrumentID string
	// 合约交易状态
	InstrumentStatus TRsaFtdcInstrumentStatusType
	// 进入本状态时间
	EnterTime string
	// 进入本状态原因
	EnterReason TRsaFtdcEnterReasonType
}

// 执行宣告输入
type CRsaFtdcInputExecOrderField struct {
	// 经纪公司编号
	BrokerID string
	// 投资者编号
	InvestorID string
	// 用户代码
	UserID string
	// 用户本地报单号
	UserOrderLocalID string
	// 用户自定义域
	UserCustom string
	// qsc全局编号
	QscId int
	// 插件标识
	PluginTaskId int
	// 外部标识
	OthMarkId string
	// 交易所代码
	ExchangeID string
	// 系统报单编号
	OrderSysID string
	// 合约代码
	InstrumentID string
	// 开平标志
	OffsetFlag TRsaFtdcOffsetFlagType
	// 投机套保标志
	HedgeFlag TRsaFtdcHedgeFlagType
	// 数量
	Volume float64
	// 业务单元
	BusinessUnit string
	// 执行类型
	ActionType TRsaFtdcActionTypeType
	// 保留头寸申请的持仓方向
	PosiDirection TRsaFtdcPosiDirectionType
	// 期权行权后是否保留期货头寸的标记
	ReservePositionFlag TRsaFtdcReservePositionFlagTypeType
	// 期权行权后生成的头寸是否自动平仓
	CloseFlag TRsaFtdcCloseFlagTypeType
	// 数量(小数)
	VolumeP float64
}

// 执行宣告操作
type CRsaFtdcExecOrderActionField struct {
	// 经纪公司编号
	BrokerID string
	// 投资者编号
	InvestorID string
	// 用户代码
	UserID string
	// 用户本地报单号
	UserOrderLocalID string
	// 用户自定义域
	UserCustom string
	// qsc全局编号
	QscId int
	// 插件标识
	PluginTaskId int
	// 外部标识
	OthMarkId string
	// 交易所代码
	ExchangeID string
	// 报单编号
	OrderSysID string
	// 合约代码
	InstrumentID string
	// 本次撤单操作的本地编号
	UserOrderActionLocalID string
	// 报单操作标志
	ActionFlag TRsaFtdcActionFlagType
	// 价格
	LimitPrice float64
	// 数量变化
	VolumeChange float64
	// 前置编号
	FrontID int
	// 会话编号
	SessionID int
}

// 询价录入
type CRsaFtdcInputForQuoteField struct {
	// 经纪公司编号
	BrokerID string
	// 投资者编号
	InvestorID string
	// 用户代码
	UserID string
	// 用户本地报单号
	UserOrderLocalID string
	// 用户自定义域
	UserCustom string
	// qsc全局编号
	QscId int
	// 插件标识
	PluginTaskId int
	// 外部标识
	OthMarkId string
	// 交易所代码
	ExchangeID string
	// 合约代码
	InstrumentID string
	// 本地报单编号
	OrderLocalID string
}

// 报价录入
type CRsaFtdcInputQuoteField struct {
	// 经纪公司编号
	BrokerID string
	// 投资者编号
	InvestorID string
	// 用户代码
	UserID string
	// 用户本地报单号
	UserOrderLocalID string
	// 用户自定义域
	UserCustom string
	// qsc全局编号
	QscId int
	// 插件标识
	PluginTaskId int
	// 外部标识
	OthMarkId string
	// 交易所代码
	ExchangeID string
	// 合约代码
	InstrumentID string
	// 系统报单编号
	OrderSysID string
	// 应价编号
	ForQuoteSysID string
	// 卖价格
	AskPrice float64
	// 卖数量
	AskVolume float64
	// 卖开平标志
	AskOffsetFlag TRsaFtdcOffsetFlagType
	// 卖投机套保标志
	AskHedgeFlag TRsaFtdcHedgeFlagType
	// 买价格
	BidPrice float64
	// 买数量
	BidVolume float64
	// 买开平标志
	BidOffsetFlag TRsaFtdcOffsetFlagType
	// 买投机套保标志
	BidHedgeFlag TRsaFtdcHedgeFlagType
	// 业务单元
	BusinessUnit string
	// 买方本地报单编号
	BidOrderRef string
	// 卖方本地报单编号
	AskOrderRef string
}

// 报价操作
type CRsaFtdcQuoteActionField struct {
	// 经纪公司编号
	BrokerID string
	// 投资者编号
	InvestorID string
	// 用户代码
	UserID string
	// 用户本地报单号
	UserOrderLocalID string
	// 用户自定义域
	UserCustom string
	// qsc全局编号
	QscId int
	// 插件标识
	PluginTaskId int
	// 外部标识
	OthMarkId string
	// 交易所代码
	ExchangeID string
	// 报单编号
	OrderSysID string
	// 合约代码
	InstrumentID string
	// 本次撤单操作的本地编号
	UserOrderActionLocalID string
	// 报单操作标志
	ActionFlag TRsaFtdcActionFlagType
	// 价格
	LimitPrice float64
	// 数量变化
	VolumeChange float64
	// 前置编号
	FrontID int
	// 会话编号
	SessionID int
}

// 期权自对冲录入
type CRsaFtdcInputOptionSelfCloseField struct {
	// 经纪公司编号
	BrokerID string
	// 投资者编码
	InvestorID string
	// 合约代码
	InstrumentID string
	// 期权自对冲引用
	OptionSelfCloseRef string
	// 用户代码
	UserID string
	// 数量
	Volume float64
	// 请求编号
	RequestID int
	// 业务单元
	BusinessUnit string
	// 投机套保标志
	HedgeFlag TRsaFtdcHedgeFlagType
	// 期权行权的头寸是否自对冲
	OptSelfCloseFlag TRsaFtdcOptSelfCloseFlagType
	// 交易所代码
	ExchangeID string
	// 投资单元代码
	InvestUnitID string
	// 资金帐号
	AccountID string
	// 币种
	Currency string
	// 客户代码
	ClientID string
	// 下单IP地址
	IPAddress string
	// 下单MAC地址
	MacAddress string
}

// 期权自对冲操作
type CRsaFtdcOptionSelfCloseActionField struct {
	// 经纪公司编号
	BrokerID string
	// 投资者编码
	InvestorID string
	// 期权自对冲操作引用
	OptionSelfCloseActionRef TRsaFtdcOrderActionRefType
	// 期权自对冲引用
	OptionSelfCloseRef string
	// 请求编号
	RequestID int
	// 前置编号
	FrontID int
	// 会话编号
	SessionID int
	// 交易所代码
	ExchangeID string
	// 期权自对冲操作编号
	OptionSelfCloseSysID string
	// 报单操作标志
	ActionFlag TRsaFtdcActionFlagType
	// 用户代码
	UserID string
	// 合约代码
	InstrumentID string
	// 投资单元代码
	InvestUnitID string
	// 下单IP地址
	IPAddress string
	// 下单MAC地址
	MacAddress string
}

// 席位状态切换
type CRsaFtdcChangeSeatStatusField struct {
	// 使用哪个通道下单
	APIID int
	// 下单席位号
	SeatID TRsaFtdcSeatIDType
	// 席位状态方向
	SeatDirection TRsaFtdcSeatDirectionType
}

// 锁定录入
type CRsaFtdcInputLockField struct {
	// 经纪公司编号
	BrokerID string
	// 投资者编号
	InvestorID string
	// 用户代码
	UserID string
	// 用户本地报单号
	UserOrderLocalID string
	// 用户自定义域
	UserCustom string
	// qsc全局编号
	QscId int
	// 插件标识
	PluginTaskId int
	// 外部标识
	OthMarkId string
	// 交易所代码
	ExchangeID string
	// 合约代码
	InstrumentID string
	// 数量
	Volume float64
	// 锁定类型
	LockType TRsaFtdcLockTypeType
	// 业务单元
	BusinessUnit string
}

// 用户信息采集
type CRsaFtdcSubmitUserSystemInfoField struct {
	// 经纪公司编号
	BrokerID string
	// 用户代码
	UserID string
	// 信息长度
	ClientSystemInfoLen int
	// 信息内容
	ClientSystemInfo string
	// 客户公网ip
	ClientPublicIP string
	// 客户端口
	ClientIPPort int
	// 登录日期
	ClientLoginDate string
	// 登录时间
	ClientLoginTime string
	// 客户认证appid
	ClientAppID string
}

// 行情基础属性
type CRsaFtdcMarketDataBaseField struct {
	// 交易日
	TradingDay string
	// 昨结算
	PreSettlementPrice float64
	// 昨收盘
	PreClosePrice float64
	// 昨持仓量
	PreOpenInterest float64
	// 昨虚实度
	PreDelta float64
}

// 行情静态属性
type CRsaFtdcMarketDataStaticField struct {
	// 今开盘
	OpenPrice float64
	// 最高价
	HighestPrice float64
	// 最低价
	LowestPrice float64
	// 今收盘
	ClosePrice float64
	// 涨停板价
	UpperLimitPrice float64
	// 跌停板价
	LowerLimitPrice float64
	// 今结算
	SettlementPrice float64
	// 今虚实度
	CurrDelta float64
}

// 行情最新成交属性
type CRsaFtdcMarketDataLastMatchField struct {
	// 最新价
	LastPrice float64
	// 数量
	Volume float64
	// 成交金额
	Turnover float64
	// 持仓量
	OpenInterest float64
}

// 行情最优价属性
type CRsaFtdcMarketDataBestPriceField struct {
	// 申买价一
	BidPrice1 float64
	// 申买量一
	BidVolume1 float64
	// 申卖价一
	AskPrice1 float64
	// 申卖量一
	AskVolume1 float64
}

// 行情申买二、三属性
type CRsaFtdcMarketDataBid23Field struct {
	// 申买价二
	BidPrice2 float64
	// 申买量二
	BidVolume2 float64
	// 申买价三
	BidPrice3 float64
	// 申买量三
	BidVolume3 float64
}

// 行情申卖二、三属性
type CRsaFtdcMarketDataAsk23Field struct {
	// 申卖价二
	AskPrice2 float64
	// 申卖量二
	AskVolume2 float64
	// 申卖价三
	AskPrice3 float64
	// 申卖量三
	AskVolume3 float64
}

// 行情申买四、五属性
type CRsaFtdcMarketDataBid45Field struct {
	// 申买价四
	BidPrice4 float64
	// 申买量四
	BidVolume4 float64
	// 申买价五
	BidPrice5 float64
	// 申买量五
	BidVolume5 float64
}

// 行情申卖四、五属性
type CRsaFtdcMarketDataAsk45Field struct {
	// 申卖价四
	AskPrice4 float64
	// 申卖量四
	AskVolume4 float64
	// 申卖价五
	AskPrice5 float64
	// 申卖量五
	AskVolume5 float64
}

// 行情更新时间属性
type CRsaFtdcMarketDataUpdateTimeField struct {
	// 合约代码
	InstrumentID string
	// 最后修改时间
	UpdateTime string
	// 最后修改毫秒
	UpdateMillisec int
}

// 深度行情
type CRsaFtdcDepthMarketDataField struct {
	// 交易日
	TradingDay string
	// 昨结算
	PreSettlementPrice float64
	// 昨收盘
	PreClosePrice float64
	// 昨持仓量
	PreOpenInterest float64
	// 昨虚实度
	PreDelta float64
	// 今开盘
	OpenPrice float64
	// 最高价
	HighestPrice float64
	// 最低价
	LowestPrice float64
	// 今收盘
	ClosePrice float64
	// 涨停板价
	UpperLimitPrice float64
	// 跌停板价
	LowerLimitPrice float64
	// 今结算
	SettlementPrice float64
	// 今虚实度
	CurrDelta float64
	// 最新价
	LastPrice float64
	// 数量
	Volume float64
	// 成交金额
	Turnover float64
	// 持仓量
	OpenInterest float64
	// 申买价一
	BidPrice1 float64
	// 申买量一
	BidVolume1 float64
	// 申卖价一
	AskPrice1 float64
	// 申卖量一
	AskVolume1 float64
	// 申买价二
	BidPrice2 float64
	// 申买量二
	BidVolume2 float64
	// 申买价三
	BidPrice3 float64
	// 申买量三
	BidVolume3 float64
	// 申卖价二
	AskPrice2 float64
	// 申卖量二
	AskVolume2 float64
	// 申卖价三
	AskPrice3 float64
	// 申卖量三
	AskVolume3 float64
	// 申买价四
	BidPrice4 float64
	// 申买量四
	BidVolume4 float64
	// 申买价五
	BidPrice5 float64
	// 申买量五
	BidVolume5 float64
	// 申卖价四
	AskPrice4 float64
	// 申卖量四
	AskVolume4 float64
	// 申卖价五
	AskPrice5 float64
	// 申卖量五
	AskVolume5 float64
	// 申卖价六
	AskPrice6 float64
	// 申卖量六
	AskVolume6 float64
	// 申卖价七
	AskPrice7 float64
	// 申卖量七
	AskVolume7 float64
	// 申卖价八
	AskPrice8 float64
	// 申卖量八
	AskVolume8 float64
	// 申卖价九
	AskPrice9 float64
	// 申卖量九
	AskVolume9 float64
	// 申买价六
	BidPrice6 float64
	// 申买量六
	BidVolume6 float64
	// 申买价七
	BidPrice7 float64
	// 申买量七
	BidVolume7 float64
	// 申买价八
	BidPrice8 float64
	// 申买量八
	BidVolume8 float64
	// 申买价九
	BidPrice9 float64
	// 申买量九
	BidVolume9 float64
	// 申卖价十
	AskPrice10 float64
	// 申卖量十
	AskVolume10 float64
	// 申买价十
	BidPrice10 float64
	// 申买量十
	BidVolume10 float64
	// 合约代码
	InstrumentID string
	// 最后修改时间
	UpdateTime string
	// 最后修改毫秒
	UpdateMillisec int
	// 交易所代码
	ExchangeID string
	// 总卖出数量
	VolumeAskLot float64
	// 总买入数量
	VolumeBidLot float64
	// 合约状态
	InstrumentStatus TRsaFtdcInstrumentStatusType
	// 当前日期
	CalendarDate string
	// 组播增量序号
	PacketNo int
	// 合约序号
	InstrumentNo int
	// 数据中心代码
	DataCenterID int
	// 更新时间
	UpdateTs int64
	// 最新成交量
	LastTraded float64
	// 最新成交额
	LastTurnover float64
}

// 订阅合约的相关信息
type CRsaFtdcSpecificInstrumentField struct {
	// 交易所代码
	ExchangeID string
	// 合约代码
	InstrumentID string
}

// 投资者权限
type CRsaFtdcInvestorRightField struct {
	// 投资者编号
	InvestorID string
	// 是否使用净持仓报单,投顾账户有效
	PositionType TRsaFtdcPositionTypeType
	// 自成交处理方式,资金账户有效
	SelfTradeAvoidType TRsaFtdcSelfTradeAvoidTypeType
	// 交易权限
	TradingRight TRsaFtdcTradingRightType
	// 是否检查资金
	CheckAccount bool
	// 是否检查持仓
	CheckPosition bool
	// 报单是否报警
	OrderWarning bool
	// 资金持仓是否上查
	QryUpAccPosi bool
	// 投顾选择资金账户类型
	DistributeType TRsaFtdcDistributeTypeType
}

// 交易所代码
type CRsaFtdcMarketDataExchangeIDField struct {
	// 交易所代码
	ExchangeID string
}

// 共享内存行情查询
type CRsaFtdcShmDepthMarketDataField struct {
	// 合约代码
	InstrumentID string
	// IP地址
	IPAddress string
	// 订阅号
	TopicID TRsaFtdcTopicIDType
}

// TOPIC查询
type CRsaFtdcTopicSearchField struct {
	// 订阅号
	TopicID TRsaFtdcTopicIDType
}

// 分价行情
type CRsaFtdcMBLMarketDataField struct {
	// 合约代码
	InstrumentID string
	// 买卖方向
	Direction TRsaFtdcDirectionType
	// 价格
	Price float64
	// 数量
	Volume float64
	// 最后修改时间
	UpdateTime string
	// 最后修改毫秒
	UpdateMillisec int
}

// 套利合约买卖总量
type CRsaFtdcMarketDataCMBVolumeField struct {
	// 总卖出数量
	VolumeAskLot float64
	// 总买入数量
	VolumeBidLot float64
}

// 合约状态
type CRsaFtdcQmdInstrumentStateField struct {
	// 交易所代码
	ExchangeID string
	// 合约代码
	InstrumentID string
	// 合约交易状态
	InstrumentStatus TRsaFtdcInstrumentStatusType
}

// 订阅合约的相关信息
type CRsaFtdcSubSpecificInstrumentField struct {
	// 交易所代码
	ExchangeID string
	// 合约代码
	InstrumentID string
}

// 行情申卖六七属性
type CRsaFtdcMarketDataAsk67Field struct {
	// 申卖价六
	AskPrice6 float64
	// 申卖量六
	AskVolume6 float64
	// 申卖价七
	AskPrice7 float64
	// 申卖量七
	AskVolume7 float64
}

// 行情申卖八九属性
type CRsaFtdcMarketDataAsk89Field struct {
	// 申卖价八
	AskPrice8 float64
	// 申卖量八
	AskVolume8 float64
	// 申卖价九
	AskPrice9 float64
	// 申卖量九
	AskVolume9 float64
}

// 行情申买六七属性
type CRsaFtdcMarketDataBid67Field struct {
	// 申买价六
	BidPrice6 float64
	// 申买量六
	BidVolume6 float64
	// 申买价七
	BidPrice7 float64
	// 申买量七
	BidVolume7 float64
}

// 行情申买八九属性
type CRsaFtdcMarketDataBid89Field struct {
	// 申买价八
	BidPrice8 float64
	// 申买量八
	BidVolume8 float64
	// 申买价九
	BidPrice9 float64
	// 申买量九
	BidVolume9 float64
}

// 行情申买申卖十属性
type CRsaFtdcMarketDataAskBid10Field struct {
	// 申卖价十
	AskPrice10 float64
	// 申卖量十
	AskVolume10 float64
	// 申买价十
	BidPrice10 float64
	// 申买量十
	BidVolume10 float64
}

// 行情合约状态
type CRsaFtdcMarketDataInstrumentStatusField struct {
	// 合约状态
	InstrumentStatus TRsaFtdcInstrumentStatusType
}

// Bar行情
type CRsaFtdcBarMarketDataField struct {
	// 交易日
	TradingDay string
	// 交易所代码
	ExchangeID string
	// 合约代码
	InstrumentID string
	// Bar精度
	BarPreces TRsaFtdcBarPrecesType
	// Bar周期
	BarPeriod int
	// Bar时间
	BarTime int64
	// 更新时间
	UpdateTs int64
	// 最后修改时间
	UpdateTime string
	// 最后修改毫秒
	UpdateMillisec int
	// 最高价
	HighestPrice float64
	// 最低价
	LowestPrice float64
	// 开盘价
	Open float64
	// 最高价
	High float64
	// 最低价
	Low float64
	// 收盘价
	Close float64
	// 总成交量
	Volume float64
	// 最新成交量
	LastTraded float64
	// 持仓量
	OpenInterest float64
	// 成交额
	Turnover float64
	// 最新成交额
	LastTurnover float64
}

// 查询Bar行情请求
type CRsaFtdcReqQryBarFromRmdField struct {
	// 交易所代码
	ExchangeID string
	// 合约代码
	InstrumentID string
}

// L2Tick行情
type CRsaFtdcL2TickMarketDataField struct {
	// 交易日
	TradingDay string
	// 交易所代码
	ExchangeID string
	// 合约代码
	InstrumentID string
	// Tick序号
	TickIndex int64
	// Tick时间
	TickTime int64
	// Tick类型
	TickType TRsaFtdcTickTypeType
	// 买方报单编号
	BuyOrderSysID string
	// 卖方报单编号
	SellOrderSysID string
	// 价格
	Price float64
	// 数量
	Volume float64
	// Tick买卖标志
	TickBSFlag TRsaFtdcTickBSFlagType
	// 通道号
	ChannelID int
}

// 成交
type CRsaFtdcTradeField struct {
	// 交易日
	TradingDay string
	// 经纪公司编号
	BrokerID string
	// 投资者编号
	InvestorID string
	// 用户代码
	UserID string
	// 用户本地报单号
	UserOrderLocalID string
	// 用户自定义域
	UserCustom string
	// qsc全局编号
	QscId int
	// 插件标识
	PluginTaskId int
	// 外部标识
	OthMarkId string
	// 交易所代码
	ExchangeID string
	// 成交编号
	TradeID TRsaFtdcTradeIDType
	// 报单编号
	OrderSysID string
	// 合约代码
	InstrumentID string
	// 买卖方向
	Direction TRsaFtdcDirectionType
	// 开平标志
	OffsetFlag TRsaFtdcOffsetFlagType
	// 投机套保标志
	HedgeFlag TRsaFtdcHedgeFlagType
	// 成交价格
	TradePrice float64
	// 成交数量
	TradeVolume float64
	// 成交时间
	TradeTime string
	// 清算会员编号
	ClearingPartID TRsaFtdcParticipantIDType
	// 成交金额
	TradeAmnt float64
	// 交易类型
	TradeType TRsaFtdcTradeTypeType
	// 成交数量（小数）
	TradeVolumeP float64
	// 营业部代码
	BranchID TRsaFtdcBranchIDType
	// 营业部名称
	BranchName TRsaFtdcBranchNameType
}

// 报单
type CRsaFtdcOrderField struct {
	// 经纪公司编号
	BrokerID string
	// 投资者编号
	InvestorID string
	// 用户代码
	UserID string
	// 用户本地报单号
	UserOrderLocalID string
	// 用户自定义域
	UserCustom string
	// qsc全局编号
	QscId int
	// 插件标识
	PluginTaskId int
	// 外部标识
	OthMarkId string
	// 交易所代码
	ExchangeID string
	// 系统报单编号
	OrderSysID string
	// 合约代码
	InstrumentID string
	// 报单类型
	OrderPriceType TRsaFtdcOrderPriceTypeType
	// 买卖方向
	Direction TRsaFtdcDirectionType
	// 开平标志
	OffsetFlag TRsaFtdcOffsetFlagType
	// 投机套保标志
	HedgeFlag TRsaFtdcHedgeFlagType
	// 价格
	LimitPrice float64
	// 数量
	Volume float64
	// 有效期类型
	TimeCondition TRsaFtdcTimeConditionType
	// GTD日期
	GTDDate string
	// 成交量类型
	VolumeCondition TRsaFtdcVolumeConditionType
	// 最小成交量
	MinVolume float64
	// 止损价
	StopPrice float64
	// 强平原因
	ForceCloseReason TRsaFtdcForceCloseReasonType
	// 自动挂起标志
	IsAutoSuspend bool
	// 业务单元
	BusinessUnit string
	// 业务类别
	BusinessType TRsaFtdcBusinessTypeType
	// 数量（小数）
	VolumeP float64
	// 策略限制价格
	AlgoLimitPrice float64
	// 策略限制tick数量
	AlgoLimitPriceTick int
	// 策略限制时间(ms)
	AlgoLimitTime int
	// 交易日
	TradingDay string
	// 插入时间
	InsertTime string
	// 报单来源
	OrderSource TRsaFtdcOrderSourceType
	// 报单状态
	OrderStatus TRsaFtdcOrderStatusType
	// 撤销时间
	CancelTime string
	// 撤单用户编号
	CancelUserID string
	// 已经成交数量
	VolumeTraded float64
	// 队列剩余数量
	VolumeRemain float64
	// 已经撤单数量
	VolumeCancled float64
	// 错误代码
	ErrorID int
	// 错误信息
	ErrorMsg string
	// 比上次增加的撤单数量
	VolumeIncCancled float64
	// 比上次增加的成交数量
	VolumeIncTraded float64
	// 资金帐号
	AccountID string
	// 前置编号
	FrontID int
	// 会话编号
	SessionID int
	// 请求编号
	RequestID int
	// 每手手续费
	EachFee float64
	// 每手保证金
	EachMargin float64
	// 每手期权权利金收支
	EachPremium float64
	// 下单IP地址
	IPAddress string
	// 下单MAC地址
	MacAddress string
	// 营业部代码
	BranchID TRsaFtdcBranchIDType
	// 记录编号
	RecNum int
	// 硬盘序列号
	HDSerialID TRsaFtdcHDSerialIDType
	// 营业部名称
	BranchName TRsaFtdcBranchNameType
	// 报单提交状态
	OrderSubmitStatus TRsaFtdcOrderSubmitStatusType
}

// 数据流回退
type CRsaFtdcFlowMessageCancelField struct {
	// 序列系列号
	SequenceSeries TRsaFtdcSequenceSeriesType
	// 交易日
	TradingDay string
	// 数据中心代码
	DataCenterID int
	// 回退起始序列号
	StartSequenceNo int
	// 回退结束序列号
	EndSequenceNo int
}

// 信息分发
type CRsaFtdcDisseminationField struct {
	// 序列系列号
	SequenceSeries TRsaFtdcSequenceSeriesType
	// 序列号
	SequenceNo int
}

// 警告消息通知
type CRsaFtdcMessageNotifyField struct {
	// 经纪公司编号
	BrokerID string
	// 修改用户编号
	UserID string
	// 投资者编号
	InvestorID string
	// 累加次数
	Nums TRsaFtdcTotalNumsType
	// 警告级别
	WarnLevel TRsaFtdcErrorLevelType
	// 警告编号
	WarnID int
	// 警告信息
	WarnMsg TRsaFtdcLargeErrorMsgType
	// 警告时间
	WarnTime string
}

// 工作流指令通知
type CRsaFtdcAdviceOrderField struct {
	// 投顾用户代码
	InvestorAdvicer string
	// 投资经理
	InvestorManager string
	// 交易员
	UserTrader string
	// 冻结数量
	VolumeFrozen float64
	// 经纪公司编号
	BrokerID string
	// 投资者编号
	InvestorID string
	// 用户代码
	UserID string
	// 用户本地报单号
	UserOrderLocalID string
	// 用户自定义域
	UserCustom string
	// qsc全局编号
	QscId int
	// 插件标识
	PluginTaskId int
	// 外部标识
	OthMarkId string
	// 交易所代码
	ExchangeID string
	// 系统报单编号
	OrderSysID string
	// 合约代码
	InstrumentID string
	// 报单类型
	OrderPriceType TRsaFtdcOrderPriceTypeType
	// 买卖方向
	Direction TRsaFtdcDirectionType
	// 开平标志
	OffsetFlag TRsaFtdcOffsetFlagType
	// 投机套保标志
	HedgeFlag TRsaFtdcHedgeFlagType
	// 价格
	LimitPrice float64
	// 数量
	Volume float64
	// 有效期类型
	TimeCondition TRsaFtdcTimeConditionType
	// GTD日期
	GTDDate string
	// 成交量类型
	VolumeCondition TRsaFtdcVolumeConditionType
	// 最小成交量
	MinVolume float64
	// 止损价
	StopPrice float64
	// 强平原因
	ForceCloseReason TRsaFtdcForceCloseReasonType
	// 自动挂起标志
	IsAutoSuspend bool
	// 业务单元
	BusinessUnit string
	// 业务类别
	BusinessType TRsaFtdcBusinessTypeType
	// 数量（小数）
	VolumeP float64
	// 策略限制价格
	AlgoLimitPrice float64
	// 策略限制tick数量
	AlgoLimitPriceTick int
	// 策略限制时间(ms)
	AlgoLimitTime int
	// 交易日
	TradingDay string
	// 插入时间
	InsertTime string
	// 报单来源
	OrderSource TRsaFtdcOrderSourceType
	// 报单状态
	OrderStatus TRsaFtdcOrderStatusType
	// 撤销时间
	CancelTime string
	// 撤单用户编号
	CancelUserID string
	// 已经成交数量
	VolumeTraded float64
	// 队列剩余数量
	VolumeRemain float64
	// 已经撤单数量
	VolumeCancled float64
	// 错误代码
	ErrorID int
	// 错误信息
	ErrorMsg string
	// 比上次增加的撤单数量
	VolumeIncCancled float64
	// 比上次增加的成交数量
	VolumeIncTraded float64
	// 资金帐号
	AccountID string
	// 前置编号
	FrontID int
	// 会话编号
	SessionID int
	// 请求编号
	RequestID int
	// 每手手续费
	EachFee float64
	// 每手保证金
	EachMargin float64
	// 每手期权权利金收支
	EachPremium float64
	// 下单IP地址
	IPAddress string
	// 下单MAC地址
	MacAddress string
	// 营业部代码
	BranchID TRsaFtdcBranchIDType
	// 记录编号
	RecNum int
	// 指令单用户信息
	AdviceUserInfo TRsaFtdcAdviceUserInfoType
	// 硬盘序列号
	HDSerialID TRsaFtdcHDSerialIDType
	// 营业部名称
	BranchName TRsaFtdcBranchNameType
	// 报单提交状态
	OrderSubmitStatus TRsaFtdcOrderSubmitStatusType
	// 父投资者编号
	ParentInvestorID string
}

// 执行宣告通知
type CRsaFtdcExecOrderField struct {
	// 经纪公司编号
	BrokerID string
	// 投资者编号
	InvestorID string
	// 用户代码
	UserID string
	// 用户本地报单号
	UserOrderLocalID string
	// 用户自定义域
	UserCustom string
	// qsc全局编号
	QscId int
	// 插件标识
	PluginTaskId int
	// 外部标识
	OthMarkId string
	// 交易所代码
	ExchangeID string
	// 系统报单编号
	OrderSysID string
	// 合约代码
	InstrumentID string
	// 开平标志
	OffsetFlag TRsaFtdcOffsetFlagType
	// 投机套保标志
	HedgeFlag TRsaFtdcHedgeFlagType
	// 数量
	Volume float64
	// 业务单元
	BusinessUnit string
	// 执行类型
	ActionType TRsaFtdcActionTypeType
	// 保留头寸申请的持仓方向
	PosiDirection TRsaFtdcPosiDirectionType
	// 期权行权后是否保留期货头寸的标记
	ReservePositionFlag TRsaFtdcReservePositionFlagTypeType
	// 期权行权后生成的头寸是否自动平仓
	CloseFlag TRsaFtdcCloseFlagTypeType
	// 数量(小数)
	VolumeP float64
	// 执行结果
	ExecResultType TRsaFtdcExecResultTypeType
	// 交易日
	TradingDay string
	// 插入时间
	InsertTime string
	// 报单来源
	OrderSource TRsaFtdcOrderSourceType
	// 报单状态
	OrderStatus TRsaFtdcOrderStatusType
	// 撤销时间
	CancelTime string
	// 撤单用户编号
	CancelUserID string
	// 已经成交数量
	VolumeTraded float64
	// 队列剩余数量
	VolumeRemain float64
	// 已经撤单数量
	VolumeCancled float64
	// 错误代码
	ErrorID int
	// 错误信息
	ErrorMsg string
	// 比上次增加的撤单数量
	VolumeIncCancled float64
	// 比上次增加的成交数量
	VolumeIncTraded float64
	// 资金帐号
	AccountID string
	// 前置编号
	FrontID int
	// 会话编号
	SessionID int
	// 请求编号
	RequestID int
	// 每手手续费
	EachFee float64
	// 每手保证金
	EachMargin float64
	// 每手期权权利金收支
	EachPremium float64
	// 下单IP地址
	IPAddress string
	// 下单MAC地址
	MacAddress string
	// 营业部代码
	BranchID TRsaFtdcBranchIDType
	// 记录编号
	RecNum int
	// 报单提交状态
	OrderSubmitStatus TRsaFtdcOrderSubmitStatusType
}

// 询价流通知
type CRsaFtdcForQuoteRspField struct {
	// 交易日
	TradingDay string
	// 合约代码
	InstrumentID string
	// 询价编号
	OrderSysID string
	// 交易所代码
	ExchangeID string
	// 插入时间
	InsertTime string
}

// 报价通知
type CRsaFtdcQuoteField struct {
	// 经纪公司编号
	BrokerID string
	// 投资者编号
	InvestorID string
	// 用户代码
	UserID string
	// 用户本地报单号
	UserOrderLocalID string
	// 用户自定义域
	UserCustom string
	// qsc全局编号
	QscId int
	// 插件标识
	PluginTaskId int
	// 外部标识
	OthMarkId string
	// 交易所代码
	ExchangeID string
	// 合约代码
	InstrumentID string
	// 系统报单编号
	OrderSysID string
	// 应价编号
	ForQuoteSysID string
	// 卖价格
	AskPrice float64
	// 卖数量
	AskVolume float64
	// 卖开平标志
	AskOffsetFlag TRsaFtdcOffsetFlagType
	// 卖投机套保标志
	AskHedgeFlag TRsaFtdcHedgeFlagType
	// 买价格
	BidPrice float64
	// 买数量
	BidVolume float64
	// 买开平标志
	BidOffsetFlag TRsaFtdcOffsetFlagType
	// 买投机套保标志
	BidHedgeFlag TRsaFtdcHedgeFlagType
	// 业务单元
	BusinessUnit string
	// 买方本地报单编号
	BidOrderRef string
	// 卖方本地报单编号
	AskOrderRef string
	// 交易日
	TradingDay string
	// 插入时间
	InsertTime string
	// 报单来源
	OrderSource TRsaFtdcOrderSourceType
	// 报单状态
	OrderStatus TRsaFtdcOrderStatusType
	// 撤销时间
	CancelTime string
	// 撤单用户编号
	CancelUserID string
	// 已经成交数量
	VolumeTraded float64
	// 队列剩余数量
	VolumeRemain float64
	// 已经撤单数量
	VolumeCancled float64
	// 错误代码
	ErrorID int
	// 错误信息
	ErrorMsg string
	// 比上次增加的撤单数量
	VolumeIncCancled float64
	// 比上次增加的成交数量
	VolumeIncTraded float64
	// 资金帐号
	AccountID string
	// 前置编号
	FrontID int
	// 会话编号
	SessionID int
	// 请求编号
	RequestID int
	// 每手手续费
	EachFee float64
	// 每手保证金
	EachMargin float64
	// 每手期权权利金收支
	EachPremium float64
	// 下单IP地址
	IPAddress string
	// 下单MAC地址
	MacAddress string
	// 营业部代码
	BranchID TRsaFtdcBranchIDType
	// 记录编号
	RecNum int
	// 买方系统报单编号
	BidOrderSysID string
	// 卖方系统报单编号
	AskOrderSysID string
	// 报单提交状态
	OrderSubmitStatus TRsaFtdcOrderSubmitStatusType
}

// 期权自对冲
type CRsaFtdcNtfOptionSelfCloseField struct {
	// 经纪公司编号
	BrokerID string
	// 投资者编码
	InvestorID string
	// 合约代码
	InstrumentID string
	// 期权自对冲引用
	OptionSelfCloseRef string
	// 用户代码
	UserID string
	// 数量
	Volume float64
	// 请求编号
	RequestID int
	// 业务单元
	BusinessUnit string
	// 投机套保标志
	HedgeFlag TRsaFtdcHedgeFlagType
	// 期权行权的头寸是否自对冲
	OptSelfCloseFlag TRsaFtdcOptSelfCloseFlagType
	// 本地期权自对冲编号
	OptionSelfCloseLocalID string
	// 交易所代码
	ExchangeID string
	// 会员编号
	ParticipantID TRsaFtdcParticipantIDType
	// 客户代码
	ClientID string
	// 合约在交易所的代码
	ExchangeInstID TRsaFtdcExchangeInstIDType
	// 交易所交易员代码
	TraderID TRsaFtdcTraderIDType
	// 安装编号
	InstallID TRsaFtdcInstallIDType
	// 期权自对冲提交状态
	OrderSubmitStatus TRsaFtdcOrderSubmitStatusType
	// 报单提示序号
	NotifySequence int
	// 交易日
	TradingDay string
	// 结算编号
	SettlementID TRsaFtdcSettlementIDType
	// 期权自对冲操作编号
	OptionSelfCloseSysID string
	// 报单日期
	InsertDate string
	// 插入时间
	InsertTime string
	// 撤销时间
	CancelTime string
	// 自对冲结果
	ExecResultType TRsaFtdcExecResultTypeType
	// 清算会员编号
	ClearingPartID TRsaFtdcParticipantIDType
	// 序号
	SequenceNo int
	// 前置编号
	FrontID int
	// 会话编号
	SessionID int
	// 用户端产品信息
	UserProductInfo string
	// 状态信息
	StatusMsg TRsaFtdcStatusMsgType
	// 操作用户代码
	ActiveUserID string
	// 经纪公司报单编号
	BrokerOptionSelfCloseSeq int
	// 营业部编号
	BranchID TRsaFtdcBranchIDType
	// 投资单元代码
	InvestUnitID string
	// 资金帐号
	AccountID string
	// 币种
	Currency string
	// 下单IP地址
	IPAddress string
	// 下单MAC地址
	MacAddress string
}

// 锁定
type CRsaFtdcLockField struct {
	// 经纪公司编号
	BrokerID string
	// 投资者编号
	InvestorID string
	// 用户代码
	UserID string
	// 用户本地报单号
	UserOrderLocalID string
	// 用户自定义域
	UserCustom string
	// qsc全局编号
	QscId int
	// 插件标识
	PluginTaskId int
	// 外部标识
	OthMarkId string
	// 交易所代码
	ExchangeID string
	// 合约代码
	InstrumentID string
	// 数量
	Volume float64
	// 锁定类型
	LockType TRsaFtdcLockTypeType
	// 业务单元
	BusinessUnit string
	// 交易日
	TradingDay string
	// 锁定编号
	LockSysID string
	// 插入时间
	InsertTime string
	// 锁定状态
	LockStatus TRsaFtdcOrderActionStatusType
	// 撤销时间
	CancelTime string
	// 资金帐号
	AccountID string
	// 前置编号
	FrontID int
	// 会话编号
	SessionID int
	// 请求编号
	RequestID int
	// 下单IP地址
	IPAddress string
	// 下单MAC地址
	MacAddress string
	// 营业部代码
	BranchID TRsaFtdcBranchIDType
	// 错误代码
	ErrorID int
	// 错误信息
	ErrorMsg string
	// 记录编号
	RecNum int
}

// 报单查询
type CRsaFtdcQryOrderField struct {
	// 经纪公司编号
	BrokerID string
	// 用户代码
	UserID string
	// 交易所代码
	ExchangeID string
	// 投资者编号
	InvestorID string
	// 报单编号
	OrderSysID string
	// 合约代码
	InstrumentID string
	// 开始时间
	InsertTimeStart string
	// 结束时间
	InsertTimeEnd string
}

// 成交查询
type CRsaFtdcQryTradeField struct {
	// 经纪公司编号
	BrokerID string
	// 用户代码
	UserID string
	// 交易所代码
	ExchangeID string
	// 投资者编号
	InvestorID string
	// 成交编号
	TradeID TRsaFtdcTradeIDType
	// 合约代码
	InstrumentID string
}

// 合约查询
type CRsaFtdcQryInstrumentField struct {
	// 交易所代码
	ExchangeID string
	// 产品代码
	ProductID TRsaFtdcProductIDType
	// 合约代码
	InstrumentID string
}

// 合约查询应答
type CRsaFtdcRspInstrumentField struct {
	// 交易所代码
	ExchangeID string
	// 品种代码
	ProductID TRsaFtdcProductIDType
	// 品种名称
	ProductName TRsaFtdcProductNameType
	// 合约代码
	InstrumentID string
	// 合约名称
	InstrumentName TRsaFtdcInstrumentNameType
	// 交割年份
	DeliveryYear TRsaFtdcYearType
	// 交割月
	DeliveryMonth TRsaFtdcMonthType
	// 限价单最大下单量
	MaxLimitOrderVolume float64
	// 限价单最小下单量
	MinLimitOrderVolume float64
	// 市价单最大下单量
	MaxMarketOrderVolume float64
	// 市价单最小下单量
	MinMarketOrderVolume float64
	// 数量乘数
	VolumeMultiple TRsaFtdcVolumeMultipleType
	// 报价单位
	PriceTick TRsaFtdcPriceTickType
	// 币种
	Currency string
	// 多头限仓
	LongPosLimit float64
	// 空头限仓
	ShortPosLimit float64
	// 跌停板价
	LowerLimitPrice float64
	// 涨停板价
	UpperLimitPrice float64
	// 昨结算
	PreSettlementPrice float64
	// 合约交易状态
	InstrumentStatus TRsaFtdcInstrumentStatusType
	// 创建日
	CreateDate string
	// 上市日
	OpenDate string
	// 到期日
	ExpireDate string
	// 开始交割日
	StartDelivDate string
	// 最后交割日
	EndDelivDate string
	// 挂牌基准价
	BasisPrice float64
	// 当前是否交易
	IsTrading bool
	// 基础商品代码
	UnderlyingInstrID string
	// 基础商品乘数
	UnderlyingMultiple float64
	// 持仓类型
	PositionType TRsaFtdcPositionTypeType
	// 执行价
	StrikePrice float64
	// 期权类型
	OptionsType TRsaFtdcOptionsTypeType
	// 汇率
	ExchangeRate float64
	// 产品类型
	ProductClass TRsaFtdcProductClassType
	// 期权行权方式
	OptionsMode TRsaFtdcOptionsModeType
	// 保证金算法类型
	MarginCombType TRsaFtdcMarginCombTypeType
	// 多头保证金率
	LongMarginRate float64
	// 空头保证金率
	ShortMarginRate float64
	// 多头保证金按手数
	LongMarginAmt float64
	// 空头保证金按手数
	ShortMarginAmt float64
}

// 投资者资金查询
type CRsaFtdcQryInvestorAccountField struct {
	// 经纪公司编号
	BrokerID string
	// 用户代码
	UserID string
	// 投资者编号
	InvestorID string
}

// 投资者资金应答
type CRsaFtdcRspInvestorAccountField struct {
	// 经纪公司编号
	BrokerID string
	// 投资者编号
	InvestorID string
	// 资金帐号
	AccountID string
	// 上次结算准备金
	PreBalance float64
	// 上日可用资金
	PreAvailable float64
	// 入金金额
	Deposit float64
	// 出金金额
	Withdraw float64
	// 占用保证金
	Margin float64
	// 期权权利金收支
	Premium float64
	// 手续费
	Fee float64
	// 冻结的保证金
	FrozenMargin float64
	// 冻结权利金
	FrozenPremium float64
	// 冻结手续费
	FrozenFee float64
	// 平仓盈亏
	CloseProfit float64
	// 持仓盈亏
	PositionProfit float64
	// 可用资金
	Available float64
	// 结算准备金
	Balance float64
	// 多头占用保证金
	LongMargin float64
	// 空头占用保证金
	ShortMargin float64
	// 多头冻结的保证金
	LongFrozenMargin float64
	// 空头冻结的保证金
	ShortFrozenMargin float64
	// 动态权益
	DynamicRights float64
	// 风险度
	Risk float64
	// 其他费用
	OtherFee float64
	// 质押金额
	Mortgage float64
	// 币种
	Currency string
	// 可取资金
	WithdrawQuota float64
	// 冻结的资金
	FrozenCash float64
	// 账户余额
	CurrentBalance float64
	// 总资产
	AssetBalance float64
	// 总市值
	MarketValue float64
	// 优先资金
	Prioprity float64
	// 现金资产
	FundAsset float64
	// 上海限购额度
	PurQuotaSh float64
	// 上海可用限购额度
	EnPurQuotaSh float64
	// 深圳限购额度
	PurQuotaSz float64
	// 深圳可用限购额度
	EnPurQuotaSz float64
}

// 可用投资者查询
type CRsaFtdcQryUserInvestorField struct {
	// 经纪公司编号
	BrokerID string
	// 用户代码
	UserID string
}

// 可用投资者应答
type CRsaFtdcRspUserInvestorField struct {
	// 经纪公司编号
	BrokerID string
	// 交易用户代码
	UserID string
	// 投资者编号
	InvestorID string
	// 默认用户代码
	NextUserID string
	// 是否可下单
	OrderMode TRsaFtdcAdviceUserOrderModeType
	// 自动或者手动
	RunMode TRsaFtdcAdviceRunModeType
	// 工作流断线处理标示符
	AdviceSwitchFlag TRsaFtdcAdviceSwitchFlagType
	// 合约交易权限
	TradingRight TRsaFtdcTradingRightType
}

// 投资者关系查询
type CRsaFtdcQryInvestorRelationField struct {
	// 经纪公司编号
	BrokerID string
	// 用户代码
	UserID string
	// 投资者编号
	InvestorID string
}

// 投资者关系应答
type CRsaFtdcRspInvestorRelationField struct {
	// 经纪公司编号
	BrokerID string
	// 投资者编号
	InvestorID string
	// 父投资者编号
	ParentInvestorID string
	// 顺序编号
	SpecifyNo int
}

// 交易所查询请求
type CRsaFtdcQryExchangeField struct {
	// 交易所代码
	ExchangeID string
}

// 交易所查询应答
type CRsaFtdcRspExchangeField struct {
	// 交易日
	TradingDay string
	// 交易所代码
	ExchangeID string
	// 交易所名称
	ExchangeName TRsaFtdcExchangeNameType
}

// 投资者持仓查询请求
type CRsaFtdcQryInvestorPositionField struct {
	// 经纪公司编号
	BrokerID string
	// 用户代码
	UserID string
	// 交易所代码
	ExchangeID string
	// 投资者编号
	InvestorID string
	// 合约代码
	InstrumentID string
}

// 投资者持仓查询应答
type CRsaFtdcRspInvestorPositionField struct {
	// 经纪公司编号
	BrokerID string
	// 交易所代码
	ExchangeID string
	// 投资者编号
	InvestorID string
	// 合约代码
	InstrumentID string
	// 买卖方向
	Direction TRsaFtdcDirectionType
	// 投机套保标志
	HedgeFlag TRsaFtdcHedgeFlagType
	// 占用保证金
	UsedMargin float64
	// 总持仓量
	Position float64
	// 总持仓成本
	PositionCost float64
	// 初始昨持仓量(当日不变)
	YdPosition float64
	// 初始昨日持仓成本(当日不变)
	YdPositionCost float64
	// 冻结的保证金
	FrozenMargin float64
	// 开仓冻结持仓
	FrozenPosition float64
	// 平仓冻结持仓
	FrozenClosing float64
	// 持仓盈亏
	PositionProfit float64
	// 冻结的权利金
	FrozenPremium float64
	// 最后一笔成交编号
	LastTradeID TRsaFtdcTradeIDType
	// 最后一笔本地报单编号
	LastOrderLocalID string
	// 总持仓可平仓数量(包括平仓冻结持仓)
	PositionClose float64
	// 昨持仓可平仓数量(包括平仓冻结持仓)
	YdPositionClose float64
	// 昨持仓平仓冻结持仓
	YdFrozenClosing float64
	// 今日开仓数量(不包括冻结)
	OpenVolume float64
	// 今日平仓数量(包括昨持仓的平仓,不包括冻结)
	CloseVolume float64
	// 平仓盈亏
	CloseProfit float64
	// 执行冻结持仓
	StrikeFrozenPosition float64
	// 放弃执行冻结持仓
	AbandonFrozenPosition float64
	// 成本价
	CostPrice float64
	// 开仓成本
	OpenCost float64
	// 今持仓(包含冻结)
	TodayPosition float64
	// 今持仓冻结持仓
	TodayFrozenClosing float64
	// 总持仓量(小数)
	PositionP float64
}

// 用户查询
type CRsaFtdcQryUserField struct {
	// 交易用户代码
	StartUserID string
	// 交易用户代码
	EndUserID string
}

// 用户
type CRsaFtdcRspUserField struct {
	// 经纪公司编号
	BrokerID string
	// 用户代码
	UserID string
	// 用户登录密码
	Password string
	// 是否活跃
	IsActive TRsaFtdcIsActiveType
	// 用户名称
	UserName TRsaFtdcUserNameType
	// 用户类型
	UserType TRsaFtdcUserTypeType
	// 营业部
	Department TRsaFtdcDepartmentType
	// 授权功能集
	GrantFuncSet TRsaFtdcGrantFuncSetType
	// 最大在线数
	MaxOnline int
	// IP地址
	IPAddress string
	// 是否检查连接产品
	CheckProductInfo bool
	// Mac地址
	MacAddress string
	// 是否检查IP和MAC
	CheckIPMacAddr bool
	// 错误登陆次数限制
	LoginErrorLimit int
	// 目前错误登陆次数
	LoginErrorCount int
	// 硬盘序列号
	HDSerialID TRsaFtdcHDSerialIDType
	// 密码是否更新
	IsUpdate TRsaFtdcIsUpdateType
}

// 投资者手续费率查询
type CRsaFtdcQryInvestorFeeField struct {
	// 经纪公司编号
	BrokerID string
	// 用户代码
	UserID string
	// 投资者编号
	InvestorID string
	// 交易所代码
	ExchangeID string
	// 合约代码
	InstrumentID string
}

// 投资者手续费率
type CRsaFtdcRspInvestorFeeField struct {
	// 经纪公司编号
	BrokerID string
	// 交易所代码
	ExchangeID string
	// 投资者编码
	InvestorID string
	// 合约代码
	InstrumentID string
	// 开仓手续费按比例
	OpenFeeRate float64
	// 开仓手续费按手数
	OpenFeeAmt float64
	// 平仓手续费按比例
	OffsetFeeRate float64
	// 平仓手续费按手数
	OffsetFeeAmt float64
	// 平今仓手续费按比例
	OTFeeRate float64
	// 平今仓手续费按手数
	OTFeeAmt float64
}

// 投资者保证金率查询
type CRsaFtdcQryInvestorMarginField struct {
	// 经纪公司编号
	BrokerID string
	// 用户代码
	UserID string
	// 投资者编号
	InvestorID string
	// 交易所代码
	ExchangeID string
	// 合约代码
	InstrumentID string
	// 投机套保标志
	HedgeFlag TRsaFtdcHedgeFlagType
}

// 投资者保证金率
type CRsaFtdcRspInvestorMarginField struct {
	// 经纪公司编号
	BrokerID string
	// 交易所代码
	ExchangeID string
	// 投资者编码
	InvestorID string
	// 合约代码
	InstrumentID string
	// 投机套保标志
	HedgeFlag TRsaFtdcHedgeFlagType
	// 多头占用保证金按比例
	LongMarginRate float64
	// 多头保证金按手数
	LongMarginAmt float64
	// 空头占用保证金按比例
	ShortMarginRate float64
	// 空头保证金按手数
	ShortMarginAmt float64
}

// 投资者查询
type CRsaFtdcQryInvestorField struct {
	// 经纪公司编号
	BrokerID string
	// 用户代码
	UserID string
	// 投资者编号
	InvestorID string
}

// 投资者应答
type CRsaFtdcRspInvestorField struct {
	// 经纪公司编号
	BrokerID string
	// 投资者编号
	InvestorID string
	// 投资者名称
	InvestorName TRsaFtdcInvestorNameType
	// 投资者类型
	InvestorType TRsaFtdcInvestorTypeType
	// 投资者组编号
	GroupID string
	// 投资者组名
	GroupName TRsaFtdcInvestorNameType
	// 别名
	Alias TRsaFtdcAliasType
	// 连接类型
	LinkType TRsaFtdcLinkTypeType
	// 资金账户编号
	ParentInvestorID string
	// 登录次数
	LoginStatus int
	// 是否使用净持仓报单,投顾账户有效
	PositionType TRsaFtdcPositionTypeType
	// 自成交处理方式,资金账户有效
	SelfTradeAvoidType TRsaFtdcSelfTradeAvoidTypeType
	// 交易权限
	TradingRight TRsaFtdcTradingRightType
	// 是否检查资金
	CheckAccount bool
	// 是否检查持仓
	CheckPosition bool
	// 报单是否报警
	OrderWarning bool
	// 资金持仓是否上查
	QryUpAccPosi bool
	// 投顾选择资金账户类型
	DistributeType TRsaFtdcDistributeTypeType
	// 跟随资金账户编号
	FollowParentInvestorID string
	// 倍率
	FollowRate float64
}

// 账号风险状况查询
type CRsaFtdcQryAccountRiskField struct {
	// 经纪公司编号
	BrokerID string
	// 用户代码
	UserID string
	// 投资者编号
	InvestorID string
}

// 账号风险状况应答
type CRsaFtdcRspAccountRiskField struct {
	// 经纪公司编号
	BrokerID string
	// 投资者编号
	InvestorID string
	// 资金帐号
	AccountID string
	// 设置准备金
	SetBalance float64
	// 开仓红线
	OpenAvailable float64
	// 平仓红线
	OffsetAvailable float64
	// 账户状态
	AccountStatus TRsaFtdcAccountStatusType
	// 动态权益
	DynamicRights float64
	// 账户权限
	TradingRight TRsaFtdcTradingRightType
}

// 交易用户会话查询
type CRsaFtdcQryUserSessionField struct {
	// 交易用户代码
	StartUserID string
	// 交易用户代码
	EndUserID string
}

// 交易用户会话应答
type CRsaFtdcRspUserSessionField struct {
	// 经纪公司代码
	BrokerID string
	// 交易用户代码
	UserID string
	// 交易用户类型
	UserType TRsaFtdcUserTypeType
	// 会话编号
	SessionID int
	// 前置编号
	FrontID int
	// 登录时间
	LoginTime string
	// IP地址
	IPAddress string
	// Mac地址
	MacAddress string
	// 用户端产品信息
	UserProductInfo string
	// 接口端产品信息
	InterfaceProductInfo string
	// 协议信息
	ProtocolInfo string
	// 硬盘序列号
	HDSerialID TRsaFtdcHDSerialIDType
}

// 行情查询
type CRsaFtdcQryMarketDataField struct {
	// 交易所代码
	ExchangeID string
	// 合约代码
	InstrumentID string
}

// 行情应答
type CRsaFtdcRspMarketDataField struct {
	// 交易所代码
	ExchangeID string
	// 交易日
	TradingDay string
	// 昨结算
	PreSettlementPrice float64
	// 昨收盘
	PreClosePrice float64
	// 昨持仓量
	PreOpenInterest float64
	// 昨虚实度
	PreDelta float64
	// 今开盘
	OpenPrice float64
	// 最高价
	HighestPrice float64
	// 最低价
	LowestPrice float64
	// 今收盘
	ClosePrice float64
	// 涨停板价
	UpperLimitPrice float64
	// 跌停板价
	LowerLimitPrice float64
	// 今结算
	SettlementPrice float64
	// 今虚实度
	CurrDelta float64
	// 最新价
	LastPrice float64
	// 数量
	Volume float64
	// 成交金额
	Turnover float64
	// 持仓量
	OpenInterest float64
	// 申买价一
	BidPrice1 float64
	// 申买量一
	BidVolume1 float64
	// 申卖价一
	AskPrice1 float64
	// 申卖量一
	AskVolume1 float64
	// 合约代码
	InstrumentID string
	// 最后修改时间
	UpdateTime string
	// 最后修改毫秒
	UpdateMillisec int
	// 合约状态
	InstrumentStatus TRsaFtdcInstrumentStatusType
}

// 产品查询
type CRsaFtdcQryProductField struct {
	// 交易所代码
	ExchangeID string
	// 产品代码
	ProductID TRsaFtdcProductIDType
}

// 产品应答
type CRsaFtdcRspProductField struct {
	// 产品代码
	ProductID TRsaFtdcProductIDType
	// 交易所代码
	ExchangeID string
	// 投资者发行的产品名称
	ProductName TRsaFtdcProductNameType
	// 产品类型
	ProductClass TRsaFtdcProductClassType
	// 数量乘数
	VolumeMultiple TRsaFtdcVolumeMultipleType
	// 报价单位
	PriceTick TRsaFtdcPriceTickType
	// 市价单最大下单量
	MaxMarketOrderVolume float64
	// 市价单最小下单量
	MinMarketOrderVolume float64
	// 限价单最大下单量
	MaxLimitOrderVolume float64
	// 限价单最小下单量
	MinLimitOrderVolume float64
	// 交易节代码
	SessionName TRsaFtdcSessionNameType
}

// 内部来回消息
type CRsaFtdcPingPongField struct {
	// 经纪公司编号
	BrokerID string
	// 修改用户编号
	UserID string
	// 投资者编号
	InvestorID string
	// 编号
	SequenceNo int
	// 用户自定义域
	UserCustom string
}

// 最大可交易数量查询请求
type CRsaFtdcQryMaxVolumeField struct {
	// 经纪公司编号
	BrokerID string
	// 用户代码
	UserID string
	// 投资者编号
	InvestorID string
	// 交易所代码
	ExchangeID string
	// 合约代码
	InstrumentID string
	// 买卖方向
	Direction TRsaFtdcDirectionType
	// 开平标志
	OffsetFlag TRsaFtdcOffsetFlagType
	// 投机套保标志
	HedgeFlag TRsaFtdcHedgeFlagType
	// 价格
	LimitPrice float64
}

// 最大可交易数量查询应答
type CRsaFtdcRspMaxVolumeField struct {
	// 经纪公司编号
	BrokerID string
	// 用户代码
	UserID string
	// 投资者编号
	InvestorID string
	// 交易所代码
	ExchangeID string
	// 合约代码
	InstrumentID string
	// 买卖方向
	Direction TRsaFtdcDirectionType
	// 开平标志
	OffsetFlag TRsaFtdcOffsetFlagType
	// 投机套保标志
	HedgeFlag TRsaFtdcHedgeFlagType
	// 价格
	LimitPrice float64
	// 数量
	Volume float64
}

// 工作流交易员查询请求
type CRsaFtdcQryAdviceTraderField struct {
	// 经纪公司编号
	BrokerID string
	// 用户代码
	UserID string
	// 投资者编号
	InvestorID string
	// 登录经纪公司编号
	LogBrokerID string
	// 登录用户代码
	LogUserID string
	// 用户端接口版本
	ApiVersion string
	// 用户端程序版本
	UserVersion string
}

// 工作流交易员查询应答
type CRsaFtdcRspAdviceTraderField struct {
	// 经纪公司编号
	BrokerID string
	// 投资者编号
	InvestorID string
	// 用户代码
	UserID string
	// 关联用户代码
	RelationUserID string
	// 关联用户名称
	RelationUserName TRsaFtdcUserNameType
}

// 组合持仓明细查询请求
type CRsaFtdcQryCmbPositionDetailField struct {
	// 投资者编号
	InvestorID string
}

// 组合持仓明细查询应答
type CRsaFtdcRspCmbPositionDetailField struct {
	// 经纪公司编号
	BrokerID string
	// 交易所代码
	ExchangeID string
	// 投资者编号
	InvestorID string
	// 单腿合约代码
	InstrumentID string
	// 单腿合约代码
	CmbInstrumentID string
	// 成交编号
	TradeID TRsaFtdcTradeIDType
	// 报单编号
	OrderSysID string
	// 买卖方向
	Direction TRsaFtdcDirectionType
	// 投机套保标志
	HedgeFlag TRsaFtdcHedgeFlagType
	// 持仓量
	Position float64
	// 单腿编号
	LegQty int
	// 交易日期
	TradingDay string
}

// 投资者结算结果查询请求
type CRsaFtdcQrySettlementInfoField struct {
	// 经纪公司编号
	BrokerID string
	// 投资者编号
	InvestorID string
	// 交易日期
	TradingDay string
}

// 投资者结算结果查询应答
type CRsaFtdcRspSettlementInfoField struct {
	// 交易日期
	TradingDay string
	// 经纪公司编号
	BrokerID string
	// 投资者编号
	InvestorID string
	// 序号
	SequenceNo int
	// 消息正文
	Content TRsaFtdcContentType
}

// 查询结果集分页
type CRsaFtdcQryResultPageField struct {
	// 结果集编号
	ResultID int
	// 总记录数
	Count int
	// 总页数
	Pages int
	// 剩余页数
	Remain int
}

// 查询持仓明细请求
type CRsaFtdcQryPositionDetailField struct {
	// 经纪公司编号
	BrokerID string
	// 交易所代码
	ExchangeID string
	// 投资者编号
	InvestorID string
	// 单腿合约代码
	InstrumentID string
}

// 查询持仓明细应答
type CRsaFtdcRspPositionDetailField struct {
	// 经纪公司编号
	BrokerID string
	// 交易所代码
	ExchangeID string
	// 投资者编号
	InvestorID string
	// 合约代码
	InstrumentID string
	// 投机套保标志
	HedgeFlag TRsaFtdcHedgeFlagType
	// 买卖方向
	Direction TRsaFtdcDirectionType
	// 开仓日期
	OpenDate string
	// 成交编号
	TradeID TRsaFtdcTradeIDType
	// 数量
	Volume float64
	// 开仓价
	OpenPrice float64
	// 交易日
	TradingDay string
	// 交易类型
	TradeType TRsaFtdcTradeTypeType
	// 组合合约代码
	CombInstrumentID string
	// 逐日盯市平仓盈亏
	CloseProfitByDate float64
	// 逐笔对冲平仓盈亏
	CloseProfitByTrade float64
	// 逐日盯市持仓盈亏
	PositionProfitByDate float64
	// 逐笔对冲持仓盈亏
	PositionProfitByTrade float64
	// 投资者保证金
	Margin float64
	// 交易所保证金
	ExchMargin float64
	// 保证金率
	MarginRateByMoney float64
	// 保证金率(按手数)
	MarginRateByVolume float64
	// 昨结算价
	LastSettlementPrice float64
	// 结算价
	SettlementPrice float64
	// 平仓量
	CloseVolume float64
	// 平仓金额
	CloseAmount float64
	// 结算编号
	SettlementID TRsaFtdcSettlementIDType
	// 投资单元代码
	InvestUnitID string
	// 持仓明细编号
	SeqNum int
	// 用户自定义域
	UserCustom string
	// 持仓明细冻结数量
	FrozenVolume float64
	// 数量(小数)
	VolumeP float64
	// 持仓明细冻结数量(小数)
	FrozenVolumeP float64
	// 平仓量(小数)
	CloseVolumeP float64
}

// 询价查询
type CRsaFtdcQryForQuoteField struct {
	// 经纪公司编号
	BrokerID string
	// 投资者编号
	InvestorID string
	// 合约代码
	InstrumentID string
	// 交易所代码
	ExchangeID string
}

// 询价查询应答
type CRsaFtdcForQuoteField struct {
	// 经纪公司编号
	BrokerID string
	// 投资者编号
	InvestorID string
	// 用户代码
	UserID string
	// 用户本地报单号
	UserOrderLocalID string
	// 用户自定义域
	UserCustom string
	// qsc全局编号
	QscId int
	// 插件标识
	PluginTaskId int
	// 外部标识
	OthMarkId string
	// 交易所代码
	ExchangeID string
	// 合约代码
	InstrumentID string
	// 本地报单编号
	OrderLocalID string
	// 询价状态
	ForQuoteStatus TRsaFtdcForQuoteStatusTypeType
	// 交易日
	TradingDay string
	// 插入时间
	InsertTime string
}

// 报价查询
type CRsaFtdcQryQuoteField struct {
	// 经纪公司编号
	BrokerID string
	// 用户代码
	UserID string
	// 投资者编号
	InvestorID string
	// 合约代码
	InstrumentID string
	// 交易所代码
	ExchangeID string
	// 系统报单编号
	OrderSysID string
	// 开始时间
	InsertTimeStart string
	// 结束时间
	InsertTimeEnd string
}

// 投资者期权手续费率查询
type CRsaFtdcQryInvestorOptionFeeField struct {
	// 经纪公司编号
	BrokerID string
	// 用户代码
	UserID string
	// 投资者编号
	InvestorID string
	// 交易所代码
	ExchangeID string
	// 合约代码
	InstrumentID string
}

// 投资者期权手续费率
type CRsaFtdcRspInvestorOptionFeeField struct {
	// 经纪公司编号
	BrokerID string
	// 交易所代码
	ExchangeID string
	// 投资者编码
	InvestorID string
	// 合约代码
	InstrumentID string
	// 开仓手续费按比例
	OpenFeeRate float64
	// 开仓手续费按手数
	OpenFeeAmt float64
	// 平仓手续费按比例
	OffsetFeeRate float64
	// 平仓手续费按手数
	OffsetFeeAmt float64
	// 平今仓手续费按比例
	OTFeeRate float64
	// 平今仓手续费按手数
	OTFeeAmt float64
	// 执行手续费按比例
	StrikeFeeRate float64
	// 执行手续费按手数
	StrikeFeeAmt float64
}

// 执行宣告查询
type CRsaFtdcQryExecOrderField struct {
	// 经纪公司编号
	BrokerID string
	// 用户代码
	UserID string
	// 交易所代码
	ExchangeID string
	// 投资者编号
	InvestorID string
	// 报单编号
	OrderSysID string
	// 合约代码
	InstrumentID string
	// 开始时间
	InsertTimeStart string
	// 结束时间
	InsertTimeEnd string
}

// 投资者通知
type CRsaFtdcRtnInvestorField struct {
	// 经纪公司编号
	BrokerID string
	// 投资者编号
	InvestorID string
	// 投资者名称
	InvestorName TRsaFtdcInvestorNameType
	// 投资者类型
	InvestorType TRsaFtdcInvestorTypeType
	// 投资者组编号
	GroupID string
	// 投资者组名
	GroupName TRsaFtdcInvestorNameType
	// 别名
	Alias TRsaFtdcAliasType
	// 连接类型
	LinkType TRsaFtdcLinkTypeType
	// 资金账户编号
	ParentInvestorID string
	// 登录次数
	LoginStatus int
	// 是否使用净持仓报单,投顾账户有效
	PositionType TRsaFtdcPositionTypeType
	// 自成交处理方式,资金账户有效
	SelfTradeAvoidType TRsaFtdcSelfTradeAvoidTypeType
	// 交易权限
	TradingRight TRsaFtdcTradingRightType
	// 是否检查资金
	CheckAccount bool
	// 是否检查持仓
	CheckPosition bool
	// 报单是否报警
	OrderWarning bool
	// 资金持仓是否上查
	QryUpAccPosi bool
	// 投顾选择资金账户类型
	DistributeType TRsaFtdcDistributeTypeType
	// 跟随资金账户编号
	FollowParentInvestorID string
	// 倍率
	FollowRate float64
}

// 投资者资金通知
type CRsaFtdcRtnInvestorAccountField struct {
	// 经纪公司编号
	BrokerID string
	// 投资者编号
	InvestorID string
	// 资金帐号
	AccountID string
	// 上次结算准备金
	PreBalance float64
	// 上日可用资金
	PreAvailable float64
	// 入金金额
	Deposit float64
	// 出金金额
	Withdraw float64
	// 占用保证金
	Margin float64
	// 期权权利金收支
	Premium float64
	// 手续费
	Fee float64
	// 冻结的保证金
	FrozenMargin float64
	// 冻结权利金
	FrozenPremium float64
	// 冻结手续费
	FrozenFee float64
	// 平仓盈亏
	CloseProfit float64
	// 持仓盈亏
	PositionProfit float64
	// 可用资金
	Available float64
	// 结算准备金
	Balance float64
	// 多头占用保证金
	LongMargin float64
	// 空头占用保证金
	ShortMargin float64
	// 多头冻结的保证金
	LongFrozenMargin float64
	// 空头冻结的保证金
	ShortFrozenMargin float64
	// 动态权益
	DynamicRights float64
	// 风险度
	Risk float64
	// 其他费用
	OtherFee float64
	// 质押金额
	Mortgage float64
	// 币种
	Currency string
	// 可取资金
	WithdrawQuota float64
	// 冻结的资金
	FrozenCash float64
	// 账户余额
	CurrentBalance float64
	// 总资产
	AssetBalance float64
	// 总市值
	MarketValue float64
	// 优先资金
	Prioprity float64
	// 现金资产
	FundAsset float64
	// 上海限购额度
	PurQuotaSh float64
	// 上海可用限购额度
	EnPurQuotaSh float64
	// 深圳限购额度
	PurQuotaSz float64
	// 深圳可用限购额度
	EnPurQuotaSz float64
}

// 投资者风控权限通知
type CRsaFtdcRtnInvestorTradingRightRiskField struct {
	// 经纪公司编号
	BrokerID string
	// 投资者编号
	InvestorID string
	// 交易所代码
	ExchangeID string
	// 合约代码
	InstrumentID string
	// 买开权限
	BuyOpen bool
	// 买平权限
	BuyClose bool
	// 卖开权限
	SellOpen bool
	// 卖平权限
	SellClose bool
	// 产品类型
	ProductClass TRsaFtdcProductClassType
}

// 投资者权限通知
type CRsaFtdcRtnInvestorTradingRightField struct {
	// 经纪公司编号
	BrokerID string
	// 投资者编号
	InvestorID string
	// 交易所代码
	ExchangeID string
	// 合约代码
	InstrumentID string
	// 合约交易权限
	TradingRight TRsaFtdcTradingRightType
	// 产品类型
	ProductClass TRsaFtdcProductClassType
}

// 交易汇率查询
type CRsaFtdcQryExchangeRateField struct {
	// 币种
	InstCurrency string
	// 币种
	AccCurrency string
}

// 交易汇率
type CRsaFtdcRspExchangeRateField struct {
	// 币种
	InstCurrency string
	// 币种
	AccCurrency string
	// 汇率
	ExchangeRate float64
}

// 交易汇率通知
type CRsaFtdcRtnExchangeRateField struct {
	// 币种
	InstCurrency string
	// 币种
	AccCurrency string
	// 汇率
	ExchangeRate float64
}

// 新组合持仓明细查询请求
type CRsaFtdcQryCombPositionDetailField struct {
	// 投资者编号
	InvestorID string
}

// 新组合持仓明细查询应答
type CRsaFtdcRspCombPositionDetailField struct {
	// 交易日期
	TradingDay string
	// 开仓日期
	OpenDate string
	// 交易所代码
	ExchangeID string
	// 结算编号
	SettlementID TRsaFtdcSettlementIDType
	// 经纪公司编号
	BrokerID string
	// 投资者编号
	InvestorID string
	// 组合编号
	ComTradeID TRsaFtdcTradeIDType
	// 成交编号
	TradeID TRsaFtdcTradeIDType
	// 合约代码
	InstrumentID string
	// 投机套保标志
	HedgeFlag TRsaFtdcHedgeFlagType
	// 买卖方向
	Direction TRsaFtdcDirectionType
	// 持仓量
	Position float64
	// 投资者保证金
	Margin float64
	// 交易所保证金
	ExchMargin float64
	// 保证金率
	MarginRateByMoney float64
	// 保证金率(按手数)
	MarginRateByVolume float64
	// 单腿编号
	LegID TRsaFtdcLegIDType
	// 单腿乘数
	LegMultiple TRsaFtdcLegMultipleType
	// 组合持仓合约代码
	CmbInstrumentID string
	// 成交组号
	TradeGroupID TRsaFtdcTradeGroupIDType
}

// 投资者持仓通知
type CRsaFtdcRtnInvestorPositionField struct {
	// 经纪公司编号
	BrokerID string
	// 交易所代码
	ExchangeID string
	// 投资者编号
	InvestorID string
	// 合约代码
	InstrumentID string
	// 买卖方向
	Direction TRsaFtdcDirectionType
	// 投机套保标志
	HedgeFlag TRsaFtdcHedgeFlagType
	// 占用保证金
	UsedMargin float64
	// 总持仓量
	Position float64
	// 总持仓成本
	PositionCost float64
	// 初始昨持仓量(当日不变)
	YdPosition float64
	// 初始昨日持仓成本(当日不变)
	YdPositionCost float64
	// 冻结的保证金
	FrozenMargin float64
	// 开仓冻结持仓
	FrozenPosition float64
	// 平仓冻结持仓
	FrozenClosing float64
	// 持仓盈亏
	PositionProfit float64
	// 冻结的权利金
	FrozenPremium float64
	// 最后一笔成交编号
	LastTradeID TRsaFtdcTradeIDType
	// 最后一笔本地报单编号
	LastOrderLocalID string
	// 总持仓可平仓数量(包括平仓冻结持仓)
	PositionClose float64
	// 昨持仓可平仓数量(包括平仓冻结持仓)
	YdPositionClose float64
	// 昨持仓平仓冻结持仓
	YdFrozenClosing float64
	// 今日开仓数量(不包括冻结)
	OpenVolume float64
	// 今日平仓数量(包括昨持仓的平仓,不包括冻结)
	CloseVolume float64
	// 平仓盈亏
	CloseProfit float64
	// 执行冻结持仓
	StrikeFrozenPosition float64
	// 放弃执行冻结持仓
	AbandonFrozenPosition float64
	// 成本价
	CostPrice float64
	// 开仓成本
	OpenCost float64
	// 今持仓(包含冻结)
	TodayPosition float64
	// 今持仓冻结持仓
	TodayFrozenClosing float64
	// 总持仓量(小数)
	PositionP float64
}

// 期权自对冲查询
type CRsaFtdcQryOptionSelfCloseField struct {
	// 经纪公司编号
	BrokerID string
	// 投资者编码
	InvestorID string
	// 合约代码
	InstrumentID string
	// 交易所代码
	ExchangeID string
	// 期权自对冲操作编号
	OptionSelfCloseSysID string
	// 开始时间
	InsertTimeStart string
	// 结束时间
	InsertTimeEnd string
}

// 锁定证券仓位查询
type CRsaFtdcQryLockPositionField struct {
	// 经纪公司编号
	BrokerID string
	// 投资者编号
	InvestorID string
	// 合约代码
	InstrumentID string
	// 交易所代码
	ExchangeID string
}

// 锁定证券仓位
type CRsaFtdcLockPositionField struct {
	// 经纪公司编号
	BrokerID string
	// 投资者编号
	InvestorID string
	// 会员编号
	ParticipantID TRsaFtdcParticipantIDType
	// 客户代码
	ClientID string
	// 合约代码
	InstrumentID string
	// 交易所代码
	ExchangeID string
	// 锁仓数量
	Volume float64
	// 冻结数量
	FrozenVolume float64
	// 今日锁定数量
	TodayVolume float64
}

// 投资者其他费率查询
type CRsaFtdcQryInvestorOtherFeeField struct {
	// 经纪公司编号
	BrokerID string
	// 用户代码
	UserID string
	// 投资者编号
	InvestorID string
	// 交易所代码
	ExchangeID string
	// 合约代码
	InstrumentID string
	// 费率类型
	FeeType TRsaFtdcFeeTypeType
}

// 投资者其他费率
type CRsaFtdcRspInvestorOtherFeeField struct {
	// 经纪公司编号
	BrokerID string
	// 交易所代码
	ExchangeID string
	// 投资者编码
	InvestorID string
	// 合约代码
	InstrumentID string
	// 开仓手续费按比例
	OpenFeeRate float64
	// 开仓手续费按手数
	OpenFeeAmt float64
	// 平仓手续费按比例
	OffsetFeeRate float64
	// 平仓手续费按手数
	OffsetFeeAmt float64
	// 平今仓手续费按比例
	OTFeeRate float64
	// 平今仓手续费按手数
	OTFeeAmt float64
	// 费率类型
	FeeType TRsaFtdcFeeTypeType
}

// 用户对应的交易员查询请求
type CRsaFtdcQryUserTraderField struct {
	// 经纪公司编号
	BrokerID string
	// 用户代码
	UserID string
	// 投资者编号
	InvestorID string
}

// 用户对应的交易员查询应答
type CRsaFtdcRspUserTraderField struct {
	// 经纪公司编号
	BrokerID string
	// 用户代码
	UserID string
	// 投资者编号
	InvestorID string
	// 父投资者编号
	ParentInvestorID string
	// 是否默认
	IsDefault bool
}

// 合约组
type CRsaFtdcInstrumentGroupField struct {
	// 经纪公司编号
	BrokerID string
	// 交易所代码
	ExchangeID string
	// 合约代码
	InstrumentID string
	// 合约组代码
	InstrumentGroupID TRsaFtdcInstrumentGroupIDType
}

// 投资者组合保证金类型
type CRsaFtdcInvestorMarginCombTypeField struct {
	// 经纪公司编号
	BrokerID string
	// 投资者编号
	InvestorID string
	// 合约组代码
	InstrumentGroupID TRsaFtdcInstrumentGroupIDType
	// 保证金组合类型
	MarginCombType TRsaFtdcMarginCombTypeType
}

// 数据同步状态
type CRsaFtdcDataSyncStatusField struct {
	// 系统名称
	SystemName TRsaFtdcSObjectType
	// 数据同步状态
	DataSyncStatus TRsaFtdcDataSyncStatusType
}

// 交易节查询请求
type CRsaFtdcQrySessionInfoField struct {
	// 交易日
	TradingDay string
}

// 交易节查询应答
type CRsaFtdcRspSessionInfoField struct {
	// 交易节Json数据
	SessionJsonData TRsaFtdcLongJsonDataType
}

// 插件交互数据
type CRsaFtdcQscPluginDataField struct {
	// qsc全局编号
	QscId int
	// 插件标识
	PluginTaskId int
	// 外部标识
	OthMarkId string
	// 插件字符串标识
	PluginId TRsaFtdcPluginIdType
	// 指令类型
	PluginOrderType TRsaFtdcPluginOrderTypeType
	// 指令编号
	FunctionId TRsaFtdcFunctionIdType
	// 经纪公司编号
	BrokerID string
	// 用户代码
	UserID string
	// 投资者编号
	InvestorID string
	// 会话编号
	SessionID int
	// 数据体
	Data TRsaFtdcJsonDataType
}

// 插件服务创建请求
type CRsaFtdcQscRegisterPluginServiceField struct {
	// qsc全局编号
	QscId int
	// 插件标识
	PluginId TRsaFtdcPluginIdType
	// 服务json描述
	ServiceDesc TRsaFtdcJsonDataType
	// 指令类型
	PluginOrderType TRsaFtdcPluginOrderTypeType
}

// 插件服务查询请求
type CRsaFtdcQryPluginServiceField struct {
	// qsc全局编号
	QscId int
	// 插件标识
	PluginId TRsaFtdcPluginIdType
	// 插件状态
	PluginStatusType TRsaFtdcPluginStatusTypeType
}

// 插件服务
type CRsaFtdcPluginServiceField struct {
	// 插件状态
	PluginStatusType TRsaFtdcPluginStatusTypeType
	// qsc全局编号
	QscId int
	// 插件标识
	PluginId TRsaFtdcPluginIdType
	// 服务json描述
	ServiceDesc TRsaFtdcJsonDataType
}

// 插件任务查询请求
type CRsaFtdcQryPluginTaskField struct {
	// qsc全局编号
	QscId int
	// 用户代码
	UserID string
}

// 插件任务
type CRsaFtdcPluginTaskField struct {
	// 插件字符串标识
	PluginId TRsaFtdcPluginIdType
	// qsc全局编号
	QscId int
	// 插件标识
	PluginTaskId int
	// 外部标识
	OthMarkId string
	// 执行功能编号
	ActionId TRsaFtdcFunctionIdType
	// 用户代码
	UserID string
	// 经纪公司编号
	BrokerID string
	// 投资者编号
	InvestorID string
	// 插件指令类型
	PluginOrderType TRsaFtdcPluginOrderTypeType
	// 插件运行状态
	PluginStatus TRsaFtdcPluginStatusType
	// 最新提示
	LastMsg TRsaFtdcMessageType
	// 数据体
	Data TRsaFtdcJsonDataType
}

// 插件任务子状态查询请求
type CRsaFtdcQryPluginTaskSubStatusField struct {
	// qsc全局编号
	QscId int
	// 插件标识
	PluginTaskId int
	// 外部标识
	OthMarkId string
}

// 插件子任务状态
type CRsaFtdcPluginTaskSubStatusField struct {
	// 插件字符串标识
	PluginId TRsaFtdcPluginIdType
	// qsc全局编号
	QscId int
	// 插件标识
	PluginTaskId int
	// 外部标识
	OthMarkId string
	// 通知功能编号
	NotifyId TRsaFtdcFunctionIdType
	// 数据关键字
	DataKey TRsaFtdcDataKeyType
	// 用户代码
	UserID string
	// 经纪公司编号
	BrokerID string
	// 投资者编号
	InvestorID string
	// 数据体
	Data TRsaFtdcJsonDataType
}

// 持仓
type CRsaFtdcInvestorPositionField struct {
	// 经纪公司编号
	BrokerID string
	// 交易所代码
	ExchangeID string
	// 投资者编号
	InvestorID string
	// 合约代码
	InstrumentID string
	// 买卖方向
	Direction TRsaFtdcDirectionType
	// 投机套保标志
	HedgeFlag TRsaFtdcHedgeFlagType
	// 占用保证金
	UsedMargin float64
	// 总持仓量
	Position float64
	// 总持仓成本
	PositionCost float64
	// 初始昨持仓量(当日不变)
	YdPosition float64
	// 初始昨日持仓成本(当日不变)
	YdPositionCost float64
	// 冻结的保证金
	FrozenMargin float64
	// 开仓冻结持仓
	FrozenPosition float64
	// 平仓冻结持仓
	FrozenClosing float64
	// 持仓盈亏
	PositionProfit float64
	// 冻结的权利金
	FrozenPremium float64
	// 最后一笔成交编号
	LastTradeID TRsaFtdcTradeIDType
	// 最后一笔本地报单编号
	LastOrderLocalID string
	// 总持仓可平仓数量(包括平仓冻结持仓)
	PositionClose float64
	// 昨持仓可平仓数量(包括平仓冻结持仓)
	YdPositionClose float64
	// 昨持仓平仓冻结持仓
	YdFrozenClosing float64
	// 今日开仓数量(不包括冻结)
	OpenVolume float64
	// 今日平仓数量(包括昨持仓的平仓,不包括冻结)
	CloseVolume float64
	// 平仓盈亏
	CloseProfit float64
	// 执行冻结持仓
	StrikeFrozenPosition float64
	// 放弃执行冻结持仓
	AbandonFrozenPosition float64
	// 成本价
	CostPrice float64
	// 开仓成本
	OpenCost float64
	// 今持仓(包含冻结)
	TodayPosition float64
	// 今持仓冻结持仓
	TodayFrozenClosing float64
	// 总持仓量(小数)
	PositionP float64
	// 资金帐号
	AccountID string
	// 会员编号
	ParticipantID TRsaFtdcParticipantIDType
	// 客户代码
	ClientID string
	// 计算浮动盈亏最新价
	LastPrice float64
	// 产品代码
	ProductID TRsaFtdcProductIDType
	// 期权类型
	OptionType TRsaFtdcOptionsTypeType
	// 行权价
	StrikePrice float64
	// 期权对应的标的合约代码
	OptionInstrumentID string
	// 短线可平仓量
	ShortClose float64
}

// 资金
type CRsaFtdcInvestorAccountField struct {
	// 经纪公司编号
	BrokerID string
	// 投资者编号
	InvestorID string
	// 资金帐号
	AccountID string
	// 上次结算准备金
	PreBalance float64
	// 上日可用资金
	PreAvailable float64
	// 入金金额
	Deposit float64
	// 出金金额
	Withdraw float64
	// 占用保证金
	Margin float64
	// 期权权利金收支
	Premium float64
	// 手续费
	Fee float64
	// 冻结的保证金
	FrozenMargin float64
	// 冻结权利金
	FrozenPremium float64
	// 冻结手续费
	FrozenFee float64
	// 平仓盈亏
	CloseProfit float64
	// 持仓盈亏
	PositionProfit float64
	// 可用资金
	Available float64
	// 结算准备金
	Balance float64
	// 多头占用保证金
	LongMargin float64
	// 空头占用保证金
	ShortMargin float64
	// 多头冻结的保证金
	LongFrozenMargin float64
	// 空头冻结的保证金
	ShortFrozenMargin float64
	// 动态权益
	DynamicRights float64
	// 风险度
	Risk float64
	// 其他费用
	OtherFee float64
	// 质押金额
	Mortgage float64
	// 币种
	Currency string
	// 可取资金
	WithdrawQuota float64
	// 冻结的资金
	FrozenCash float64
	// 账户余额
	CurrentBalance float64
	// 总资产
	AssetBalance float64
	// 总市值
	MarketValue float64
	// 优先资金
	Prioprity float64
	// 现金资产
	FundAsset float64
	// 上海限购额度
	PurQuotaSh float64
	// 上海可用限购额度
	EnPurQuotaSh float64
	// 深圳限购额度
	PurQuotaSz float64
	// 深圳可用限购额度
	EnPurQuotaSz float64
	// 总冻结持仓
	TotalFrozenPos float64
	// 今日出入金
	TodayInOut float64
	// 当日释放保证金
	ReleaseMargin float64
	// 浮盈是否可用
	IsProfitCanUse bool
	// 临时冻结资金
	ManualFrozen float64
	// 期权权利金收入
	PremiumIn float64
	// 期权权利金支出
	PremiumOut float64
	// 现货买入冻结资金
	SpotFrozen float64
	// 现货买入成交资金冻结
	SpotTradeFrozen float64
	// 现货卖出成交回笼资金
	SpotTradeRelease float64
}

// 系统状态通知
type CRsaFtdcSystemStatusField struct {
	// 系统类型
	SysType TRsaFtdcSysTypeType
	// 系统标识
	SysTag TRsaFtdcSysTagType
	// 系统状态
	SysStatus TRsaFtdcSysStatusType
}

// 插件任务最新状态信息
type CRsaFtdcPluginTaskInfoField struct {
	// 插件字符串标识
	PluginId TRsaFtdcPluginIdType
	// qsc全局编号
	QscId int
	// 插件标识
	PluginTaskId int
	// 外部标识
	OthMarkId string
	// 插件运行状态
	PluginStatus TRsaFtdcPluginStatusType
	// 最新提示
	LastMsg TRsaFtdcMessageType
	// 数据体
	Data TRsaFtdcJsonDataType
}

// 账户新股申购额度查询
type CRsaFtdcQryInvestorAccountIPOField struct {
	// 投资者编号
	InvestorID string
	// 交易所代码
	ExchangeID string
}

// 账户新股申购额度
type CRsaFtdcInvestorAccountIPOField struct {
	// 经纪公司编号
	BrokerID string
	// 投资者编号
	InvestorID string
	// 交易所代码
	ExchangeID string
	// 资金帐号
	AccountID string
	// 产品编号
	CPBH TRsaFtdcCPBHType
	// 股东代码
	GDDM TRsaFtdcGDDMType
	// 主板申购额度
	ZBSGED TRsaFtdcZBSGEDType
	// 科创板申购额度
	KCBSGED TRsaFtdcKCBSGEDType
}

// 客户询价请求
type CRsaFtdcClientReqQuoteField struct {
	// 交易所代码
	ExchangeID string
	// 合约代码
	InstrumentID string
	// 投资者编号
	InvestorID string
	// 经纪公司编号
	BrokerID string
	// 用户代码
	UserID string
	// 系统询价编号
	QuoteID TRsaFtdcQuoteIDType
	// 用户本地报单号
	UserOrderLocalID string
	// 询价请求时间
	ReqTime string
	// 用户自定义域
	UserCustom string
}

// 客户询价请求应答
type CRsaFtdcRspClientReqQuoteField struct {
	// 交易所代码
	ExchangeID string
	// 合约代码
	InstrumentID string
	// 投资者编号
	InvestorID string
	// 经纪公司编号
	BrokerID string
	// 用户代码
	UserID string
	// 系统询价编号
	QuoteID TRsaFtdcQuoteIDType
	// 用户本地报单号
	UserOrderLocalID string
	// 询价请求时间
	ReqTime string
	// 用户自定义域
	UserCustom string
	// 价格
	LimitPrice float64
	// 库存
	Position float64
}

// 客户询价请求查询请求
type CRsaFtdcReqQryClientReqQuoteField struct {
	// 交易所代码
	ExchangeID string
	// 合约代码
	InstrumentID string
	// 投资者编号
	InvestorID string
	// 系统询价编号
	QuoteID TRsaFtdcQuoteIDType
}

// 客户询价请求应答
type CRsaFtdcRspQryClientReqQuoteField struct {
	// 交易所代码
	ExchangeID string
	// 合约代码
	InstrumentID string
	// 投资者编号
	InvestorID string
	// 经纪公司编号
	BrokerID string
	// 用户代码
	UserID string
	// 系统询价编号
	QuoteID TRsaFtdcQuoteIDType
	// 用户本地报单号
	UserOrderLocalID string
	// 询价请求时间
	ReqTime string
	// 用户自定义域
	UserCustom string
	// 价格
	LimitPrice float64
	// 库存
	Position float64
	// 询价应答时间
	RspTime string
	// 客户询价状态
	ClientReqQuoteStatus TRsaFtdcClientReqQuoteStatusType
	// 错误代码
	ErrorID int
	// 错误信息
	ErrorMsg string
}

// 发送通用消息请求
type CRsaFtdcComMessageField struct {
	// 用户代码
	UserID string
	// 投资者编号
	InvestorID string
	// 消息类型
	MessageType TRsaFtdcMessageTypeType
	// 信息内容
	MessageContent TRsaFtdcMessageType
}

// 共享内存登录消息
type CRsaFtdcReqShmLoginField struct {
	// 交易用户代码
	UserID string
	// 经纪公司编号
	BrokerID string
	// 密码
	Password string
}

// 切日通知
type CRsaFtdcUpdTradingDayField struct {
	// 当前交易日（切日前）
	CurrentTradingDay string
	// 下一交易日（切日后）
	NextTradingDay string
	// 更新时间
	UpdTime string
	// 是否为交易日结束
	IsTradingDayEnd bool
}

// 行情推送结束通知
type CRsaFtdcNtfMarketDataEndField struct {
	// 行情开始时间
	StartTradingDay string
	// 行情结束时间
	EndTradingDay string
	// 时间
	Time string
}

// 回测行情订阅
type CRsaFtdcBtSubMarketDataField struct {
	// 交易所代码
	ExchangeID string
	// 合约代码
	InstrumentID string
	// Bar精度
	BarPreces TRsaFtdcBarPrecesType
	// Bar周期
	BarPeriod int
}

// 主力合约信息查询请求
type CRsaFtdcQryHotInstrumentField struct {
	// 交易日
	TradingDay string
	// 交易所代码
	ExchangeID string
	// 基础品种代码
	UnderlyingID string
}

// 主力合约信息查询应答
type CRsaFtdcRspHotInstrumentField struct {
	// 交易日
	TradingDay string
	// 交易所代码
	ExchangeID string
	// 基础品种代码
	UnderlyingID string
	// 合约代码
	InstrumentID string
	// 产品类型
	ProductClass TRsaFtdcProductClassType
	// 成交量
	Volume float64
	// 品种最大成交量
	MaxVolume float64
	// 成交金额
	Turnover float64
	// 品种最大成交金额
	MaxTurnover float64
	// 持仓量
	OpenInterest float64
	// 品种最大持仓量
	MaxOpenInterest float64
	// 主力合约等级
	Rank float64
}

// Cta委托
type CRsaFtdcCtaOrderField struct {
	// 投资者编号
	InvestorID string
	// 交易所代码
	ExchangeID string
	// 合约代码
	InstrumentID string
	// 数量
	Volume float64
	// 报单类型
	OrderPriceType TRsaFtdcOrderPriceTypeType
	// 开平方式
	OpenCloseMethod TRsaFtdcOpenCloseMethodType
	// 价格
	LimitPrice float64
	// 业务单元
	BusinessUnit string
	// 业务类别
	BusinessType TRsaFtdcBusinessTypeType
	// 策略限制价格
	AlgoLimitPrice float64
	// 策略限制tick数量
	AlgoLimitPriceTick int
	// 策略限制时间(ms)
	AlgoLimitTime int
}

// 订阅组合行情
type CRsaFtdcSubCombMarketDataField struct {
	// 组合合约代码
	CombInstID string
	// 公式
	Formula TRsaFtdcFormulaType
	// 小数点位数
	Precision float64
	// 腿数
	LegNum float64
	// Leg1交易所代码
	Leg1ExchangeID string
	// Leg1合约代码
	Leg1InstrumentID string
	// Leg1买卖方向
	Leg1Direction TRsaFtdcDirectionType
	// Leg2易所代码
	Leg2ExchangeID string
	// Leg2合约代码
	Leg2InstrumentID string
	// Leg2买卖方向
	Leg2Direction TRsaFtdcDirectionType
	// Leg3交易所代码
	Leg3ExchangeID string
	// Leg3合约代码
	Leg3InstrumentID string
	// Leg3买卖方向
	Leg3Direction TRsaFtdcDirectionType
	// Leg4交易所代码
	Leg4ExchangeID string
	// Leg4合约代码
	Leg4InstrumentID string
	// Leg4买卖方向
	Leg4Direction TRsaFtdcDirectionType
	// Leg5交易所代码
	Leg5ExchangeID string
	// Leg5合约代码
	Leg5InstrumentID string
	// Leg5买卖方向
	Leg5Direction TRsaFtdcDirectionType
	// Leg6交易所代码
	Leg6ExchangeID string
	// Leg6合约代码
	Leg6InstrumentID string
	// Leg6买卖方向
	Leg6Direction TRsaFtdcDirectionType
	// Bar精度
	BarPreces TRsaFtdcBarPrecesType
	// Bar周期
	BarPeriod int
}

// 流水状态
type CRsaFtdcFlowStatusField struct {
	// 序列系列号
	SequenceSeries TRsaFtdcSequenceSeriesType
	// 是否就绪
	Ready bool
}

// 行情订阅完成请求
type CRsaFtdcReqSubMarketDataCompletedField struct {
	// 当前时间
	CurrTime int64
}

// 策略消息
type CRsaFtdcStgMsgField struct {
	// 经纪公司编号
	BrokerID string
	// 用户代码
	UserID string
	// 投资者编号
	InvestorID string
	// 系统报单编号
	OrderSysID string
	// 会话编号
	SessionID int
	// 策略消息类型
	StgMsgType TRsaFtdcStgMsgTypeType
	// 包数量
	Total int
	// 包序号
	Index int
	// Json数据串
	Data TRsaFtdcJsonDataType
}

// 策略委托消息
type CRsaFtdcStgOrderField struct {
	// 经纪公司编号
	BrokerID string
	// 用户代码
	UserID string
	// 投资者编号
	InvestorID string
	// 系统报单编号
	OrderSysID string
	// 会话编号
	SessionID int
	// 交易所代码
	ExchangeID string
	// 合约代码
	InstrumentID string
	// 买卖方向
	Direction TRsaFtdcDirectionType
	// 开平标志
	OffsetFlag TRsaFtdcOffsetFlagType
	// 数量
	Volume float64
	// 已经成交数量
	VolumeTraded float64
	// 报单状态
	OrderStatus TRsaFtdcOrderStatusType
	// 错误代码
	ErrorID int
	// 错误信息
	ErrorMsg string
}

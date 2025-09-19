package rmd4go

// 响应信息
type CRsaFtdcRspInfoField struct {
	// 错误代码
	ErrorID TRsaFtdcErrorIDType
	// 错误信息
	ErrorMsg TRsaFtdcErrorMsgType
}

// 系统用户登录请求
type CRsaFtdcReqUserLoginField struct {
	// 交易日
	TradingDay TRsaFtdcDateType
	// 交易用户代码
	UserID TRsaFtdcUserIDType
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 密码
	Password TRsaFtdcPasswordType
	// 用户端产品信息
	UserProductInfo TRsaFtdcProductInfoType
	// 接口端产品信息
	InterfaceProductInfo TRsaFtdcProductInfoType
	// 协议信息
	ProtocolInfo TRsaFtdcProtocolInfoType
	// IP地址
	IPAddress TRsaFtdcIPAddressType
	// Mac地址
	MacAddress TRsaFtdcMacAddressType
	// 数据中心代码
	DataCenterID TRsaFtdcDataCenterIDType
	// 动态密码
	OneTimePassword TRsaFtdcPasswordType
	// 终端IP地址
	ClientIPAddress TRsaFtdcIPAddressType
	// 授权编码
	AuthCode TRsaFtdcAuthCodeType
	// 自定义字段
	FieldContent TRsaFtdcFieldContentType
	// 客户端信息采集
	ClientSystemInfo TRsaFtdcClientSystemInfoType
}

// 系统用户登录应答
type CRsaFtdcRspUserLoginField struct {
	// 交易日
	TradingDay TRsaFtdcDateType
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 交易用户代码
	UserID TRsaFtdcUserIDType
	// 登录成功时间
	LoginTime TRsaFtdcTimeType
	// 用户最大本地报单号
	MaxOrderLocalID TRsaFtdcUserOrderLocalIDType
	// 交易系统名称
	TradingSystemName TRsaFtdcTradingSystemNameType
	// 数据中心代码
	DataCenterID TRsaFtdcDataCenterIDType
	// 会员私有流当前长度
	PrivateFlowSize TRsaFtdcSequenceNoType
	// 交易员私有流当前长度
	UserFlowSize TRsaFtdcSequenceNoType
	// 登录附加信息
	LoginInfo TRsaFtdcLoginInfoType
	// 会话编号
	SessionID TRsaFtdcSessionIDType
	// 前置编号
	FrontID TRsaFtdcFrontIDType
}

// 用户登出请求
type CRsaFtdcReqUserLogoutField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 交易用户代码
	UserID TRsaFtdcUserIDType
}

// 用户登出请求
type CRsaFtdcRspUserLogoutField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 交易用户代码
	UserID TRsaFtdcUserIDType
}

// 强制用户退出
type CRsaFtdcForceUserExitField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 交易用户代码
	UserID TRsaFtdcUserIDType
}

// 用户口令修改
type CRsaFtdcUserPasswordUpdateField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 交易用户代码
	UserID TRsaFtdcUserIDType
	// 旧密码
	OldPassword TRsaFtdcPasswordType
	// 新密码
	NewPassword TRsaFtdcPasswordType
}

// 输入报单
type CRsaFtdcInputOrderField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
	// 用户代码
	UserID TRsaFtdcUserIDType
	// 用户本地报单号
	UserOrderLocalID TRsaFtdcUserOrderLocalIDType
	// 用户自定义域
	UserCustom TRsaFtdcCustomType
	// qsc全局编号
	QscId TRsaFtdcNumberType
	// 插件标识
	PluginTaskId TRsaFtdcNumberType
	// 外部标识
	OthMarkId TRsaFtdcOthMarkIdType
	// 交易所代码
	ExchangeID TRsaFtdcExchangeIDType
	// 系统报单编号
	OrderSysID TRsaFtdcOrderSysIDType
	// 合约代码
	InstrumentID TRsaFtdcInstrumentIDType
	// 报单类型
	OrderPriceType TRsaFtdcOrderPriceTypeType
	// 买卖方向
	Direction TRsaFtdcDirectionType
	// 开平标志
	OffsetFlag TRsaFtdcOffsetFlagType
	// 投机套保标志
	HedgeFlag TRsaFtdcHedgeFlagType
	// 价格
	LimitPrice TRsaFtdcPriceType
	// 数量
	Volume TRsaFtdcVolumeType
	// 有效期类型
	TimeCondition TRsaFtdcTimeConditionType
	// GTD日期
	GTDDate TRsaFtdcDateType
	// 成交量类型
	VolumeCondition TRsaFtdcVolumeConditionType
	// 最小成交量
	MinVolume TRsaFtdcVolumeType
	// 止损价
	StopPrice TRsaFtdcPriceType
	// 强平原因
	ForceCloseReason TRsaFtdcForceCloseReasonType
	// 自动挂起标志
	IsAutoSuspend TRsaFtdcBoolType
	// 业务单元
	BusinessUnit TRsaFtdcBusinessUnitType
	// 业务类别
	BusinessType TRsaFtdcBusinessTypeType
	// 数量（小数）
	VolumeP TRsaFtdcPointVolumeType
	// 策略限制价格
	AlgoLimitPrice TRsaFtdcPriceType
	// 策略限制tick数量
	AlgoLimitPriceTick TRsaFtdcNumberType
	// 策略限制时间(ms)
	AlgoLimitTime TRsaFtdcNumberType
}

// 报单操作
type CRsaFtdcOrderActionField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
	// 用户代码
	UserID TRsaFtdcUserIDType
	// 用户本地报单号
	UserOrderLocalID TRsaFtdcUserOrderLocalIDType
	// 用户自定义域
	UserCustom TRsaFtdcCustomType
	// qsc全局编号
	QscId TRsaFtdcNumberType
	// 插件标识
	PluginTaskId TRsaFtdcNumberType
	// 外部标识
	OthMarkId TRsaFtdcOthMarkIdType
	// 交易所代码
	ExchangeID TRsaFtdcExchangeIDType
	// 报单编号
	OrderSysID TRsaFtdcOrderSysIDType
	// 合约代码
	InstrumentID TRsaFtdcInstrumentIDType
	// 本次撤单操作的本地编号
	UserOrderActionLocalID TRsaFtdcUserOrderLocalIDType
	// 报单操作标志
	ActionFlag TRsaFtdcActionFlagType
	// 价格
	LimitPrice TRsaFtdcPriceType
	// 数量变化
	VolumeChange TRsaFtdcVolumeType
	// 前置编号
	FrontID TRsaFtdcFrontIDType
	// 会话编号
	SessionID TRsaFtdcSessionIDType
}

// 内存表导出
type CRsaFtdcMemDBField struct {
	// 内存表名
	MemTableName TRsaFtdcMemTableNameType
}

// 投资者资金帐户出入金请求
type CRsaFtdcReqAccountDepositField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 资金流水号
	AccountSeqNo TRsaFtdcAccountSeqNoType
	// 资金帐号
	AccountID TRsaFtdcAccountIDType
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
	// 金额
	Amount TRsaFtdcMoneyType
	// 出入金方向
	AmountDirection TRsaFtdcAccountDirectionType
	// 用户代码
	UserID TRsaFtdcUserIDType
}

// 投资者资金帐户出入金应答
type CRsaFtdcRspAccountDepositField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
	// 资金帐号
	AccountID TRsaFtdcAccountIDType
	// 资金流水号
	AccountSeqNo TRsaFtdcAccountSeqNoType
	// 金额
	Amount TRsaFtdcMoneyType
	// 出入金方向
	AmountDirection TRsaFtdcAccountDirectionType
	// 可用资金
	Available TRsaFtdcMoneyType
	// 结算准备金
	Balance TRsaFtdcMoneyType
	// 用户代码
	UserID TRsaFtdcUserIDType
}

// 实时上场
type CRsaFtdcTBCommandField struct {
	// DB命令序号
	CommandNo TRsaFtdcCommandNoType
	// 操作日期
	CommandDate TRsaFtdcDateType
	// 操作时间
	CommandTime TRsaFtdcTimeType
	// DB命令类型
	CommandType TRsaFtdcCommandTypeType
	// 字段名
	FieldName TRsaFtdcFieldNameType
	// 字段内容
	FieldContent TRsaFtdcFieldContentType
}

// 出入金结果
type CRsaFtdcInvestorAccountDepositResField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
	// 资金帐号
	AccountID TRsaFtdcAccountIDType
	// 资金流水号
	AccountSeqNo TRsaFtdcAccountSeqNoType
	// 金额
	Amount TRsaFtdcMoneyType
	// 出入金方向
	AmountDirection TRsaFtdcAccountDirectionType
	// 可用资金
	Available TRsaFtdcMoneyType
	// 结算准备金
	Balance TRsaFtdcMoneyType
	// 用户代码
	UserID TRsaFtdcUserIDType
}

// 合约状态
type CRsaFtdcInstrumentStatusField struct {
	// 交易所代码
	ExchangeID TRsaFtdcExchangeIDType
	// 品种代码
	ProductID TRsaFtdcProductIDType
	// 品种名称
	ProductName TRsaFtdcProductNameType
	// 合约代码
	InstrumentID TRsaFtdcInstrumentIDType
	// 合约名称
	InstrumentName TRsaFtdcInstrumentNameType
	// 交割年份
	DeliveryYear TRsaFtdcYearType
	// 交割月
	DeliveryMonth TRsaFtdcMonthType
	// 限价单最大下单量
	MaxLimitOrderVolume TRsaFtdcVolumeType
	// 限价单最小下单量
	MinLimitOrderVolume TRsaFtdcVolumeType
	// 市价单最大下单量
	MaxMarketOrderVolume TRsaFtdcVolumeType
	// 市价单最小下单量
	MinMarketOrderVolume TRsaFtdcVolumeType
	// 数量乘数
	VolumeMultiple TRsaFtdcVolumeMultipleType
	// 报价单位
	PriceTick TRsaFtdcPriceTickType
	// 币种
	Currency TRsaFtdcCurrencyType
	// 多头限仓
	LongPosLimit TRsaFtdcVolumeType
	// 空头限仓
	ShortPosLimit TRsaFtdcVolumeType
	// 跌停板价
	LowerLimitPrice TRsaFtdcPriceType
	// 涨停板价
	UpperLimitPrice TRsaFtdcPriceType
	// 昨结算
	PreSettlementPrice TRsaFtdcPriceType
	// 合约交易状态
	InstrumentStatus TRsaFtdcInstrumentStatusType
	// 创建日
	CreateDate TRsaFtdcDateType
	// 上市日
	OpenDate TRsaFtdcDateType
	// 到期日
	ExpireDate TRsaFtdcDateType
	// 开始交割日
	StartDelivDate TRsaFtdcDateType
	// 最后交割日
	EndDelivDate TRsaFtdcDateType
	// 挂牌基准价
	BasisPrice TRsaFtdcPriceType
	// 当前是否交易
	IsTrading TRsaFtdcBoolType
	// 基础商品代码
	UnderlyingInstrID TRsaFtdcInstrumentIDType
	// 基础商品乘数
	UnderlyingMultiple TRsaFtdcUnderlyingMultipleType
	// 持仓类型
	PositionType TRsaFtdcPositionTypeType
	// 执行价
	StrikePrice TRsaFtdcPriceType
	// 期权类型
	OptionsType TRsaFtdcOptionsTypeType
	// 汇率
	ExchangeRate TRsaFtdcRatioType
	// 产品类型
	ProductClass TRsaFtdcProductClassType
	// 期权行权方式
	OptionsMode TRsaFtdcOptionsModeType
	// 保证金算法类型
	MarginCombType TRsaFtdcMarginCombTypeType
	// 多头保证金率
	LongMarginRate TRsaFtdcRatioType
	// 空头保证金率
	ShortMarginRate TRsaFtdcRatioType
	// 多头保证金按手数
	LongMarginAmt TRsaFtdcRatioType
	// 空头保证金按手数
	ShortMarginAmt TRsaFtdcRatioType
}

// 审批模式设置请求
type CRsaFtdcAdviceRunModeField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
	// 用户代码
	UserID TRsaFtdcUserIDType
	// 自动或者手动
	RunMode TRsaFtdcAdviceRunModeType
}

// 订阅合约状态
type CRsaFtdcSubInstrumentStatusField struct {
	// 交易所代码
	ExchangeID TRsaFtdcExchangeIDType
	// 合约代码
	InstrumentID TRsaFtdcInstrumentIDType
	// 合约交易状态
	InstrumentStatus TRsaFtdcInstrumentStatusType
	// 进入本状态时间
	EnterTime TRsaFtdcTimeType
	// 进入本状态原因
	EnterReason TRsaFtdcEnterReasonType
}

// 执行宣告输入
type CRsaFtdcInputExecOrderField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
	// 用户代码
	UserID TRsaFtdcUserIDType
	// 用户本地报单号
	UserOrderLocalID TRsaFtdcUserOrderLocalIDType
	// 用户自定义域
	UserCustom TRsaFtdcCustomType
	// qsc全局编号
	QscId TRsaFtdcNumberType
	// 插件标识
	PluginTaskId TRsaFtdcNumberType
	// 外部标识
	OthMarkId TRsaFtdcOthMarkIdType
	// 交易所代码
	ExchangeID TRsaFtdcExchangeIDType
	// 系统报单编号
	OrderSysID TRsaFtdcOrderSysIDType
	// 合约代码
	InstrumentID TRsaFtdcInstrumentIDType
	// 开平标志
	OffsetFlag TRsaFtdcOffsetFlagType
	// 投机套保标志
	HedgeFlag TRsaFtdcHedgeFlagType
	// 数量
	Volume TRsaFtdcVolumeType
	// 业务单元
	BusinessUnit TRsaFtdcBusinessUnitType
	// 执行类型
	ActionType TRsaFtdcActionTypeType
	// 保留头寸申请的持仓方向
	PosiDirection TRsaFtdcPosiDirectionType
	// 期权行权后是否保留期货头寸的标记
	ReservePositionFlag TRsaFtdcReservePositionFlagTypeType
	// 期权行权后生成的头寸是否自动平仓
	CloseFlag TRsaFtdcCloseFlagTypeType
	// 数量(小数)
	VolumeP TRsaFtdcPointVolumeType
}

// 执行宣告操作
type CRsaFtdcExecOrderActionField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
	// 用户代码
	UserID TRsaFtdcUserIDType
	// 用户本地报单号
	UserOrderLocalID TRsaFtdcUserOrderLocalIDType
	// 用户自定义域
	UserCustom TRsaFtdcCustomType
	// qsc全局编号
	QscId TRsaFtdcNumberType
	// 插件标识
	PluginTaskId TRsaFtdcNumberType
	// 外部标识
	OthMarkId TRsaFtdcOthMarkIdType
	// 交易所代码
	ExchangeID TRsaFtdcExchangeIDType
	// 报单编号
	OrderSysID TRsaFtdcOrderSysIDType
	// 合约代码
	InstrumentID TRsaFtdcInstrumentIDType
	// 本次撤单操作的本地编号
	UserOrderActionLocalID TRsaFtdcUserOrderLocalIDType
	// 报单操作标志
	ActionFlag TRsaFtdcActionFlagType
	// 价格
	LimitPrice TRsaFtdcPriceType
	// 数量变化
	VolumeChange TRsaFtdcVolumeType
	// 前置编号
	FrontID TRsaFtdcFrontIDType
	// 会话编号
	SessionID TRsaFtdcSessionIDType
}

// 询价录入
type CRsaFtdcInputForQuoteField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
	// 用户代码
	UserID TRsaFtdcUserIDType
	// 用户本地报单号
	UserOrderLocalID TRsaFtdcUserOrderLocalIDType
	// 用户自定义域
	UserCustom TRsaFtdcCustomType
	// qsc全局编号
	QscId TRsaFtdcNumberType
	// 插件标识
	PluginTaskId TRsaFtdcNumberType
	// 外部标识
	OthMarkId TRsaFtdcOthMarkIdType
	// 交易所代码
	ExchangeID TRsaFtdcExchangeIDType
	// 合约代码
	InstrumentID TRsaFtdcInstrumentIDType
	// 本地报单编号
	OrderLocalID TRsaFtdcOrderLocalIDType
}

// 报价录入
type CRsaFtdcInputQuoteField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
	// 用户代码
	UserID TRsaFtdcUserIDType
	// 用户本地报单号
	UserOrderLocalID TRsaFtdcUserOrderLocalIDType
	// 用户自定义域
	UserCustom TRsaFtdcCustomType
	// qsc全局编号
	QscId TRsaFtdcNumberType
	// 插件标识
	PluginTaskId TRsaFtdcNumberType
	// 外部标识
	OthMarkId TRsaFtdcOthMarkIdType
	// 交易所代码
	ExchangeID TRsaFtdcExchangeIDType
	// 合约代码
	InstrumentID TRsaFtdcInstrumentIDType
	// 系统报单编号
	OrderSysID TRsaFtdcOrderSysIDType
	// 应价编号
	ForQuoteSysID TRsaFtdcOrderSysIDType
	// 卖价格
	AskPrice TRsaFtdcPriceType
	// 卖数量
	AskVolume TRsaFtdcVolumeType
	// 卖开平标志
	AskOffsetFlag TRsaFtdcOffsetFlagType
	// 卖投机套保标志
	AskHedgeFlag TRsaFtdcHedgeFlagType
	// 买价格
	BidPrice TRsaFtdcPriceType
	// 买数量
	BidVolume TRsaFtdcVolumeType
	// 买开平标志
	BidOffsetFlag TRsaFtdcOffsetFlagType
	// 买投机套保标志
	BidHedgeFlag TRsaFtdcHedgeFlagType
	// 业务单元
	BusinessUnit TRsaFtdcBusinessUnitType
	// 买方本地报单编号
	BidOrderRef TRsaFtdcUserOrderLocalIDType
	// 卖方本地报单编号
	AskOrderRef TRsaFtdcUserOrderLocalIDType
}

// 报价操作
type CRsaFtdcQuoteActionField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
	// 用户代码
	UserID TRsaFtdcUserIDType
	// 用户本地报单号
	UserOrderLocalID TRsaFtdcUserOrderLocalIDType
	// 用户自定义域
	UserCustom TRsaFtdcCustomType
	// qsc全局编号
	QscId TRsaFtdcNumberType
	// 插件标识
	PluginTaskId TRsaFtdcNumberType
	// 外部标识
	OthMarkId TRsaFtdcOthMarkIdType
	// 交易所代码
	ExchangeID TRsaFtdcExchangeIDType
	// 报单编号
	OrderSysID TRsaFtdcOrderSysIDType
	// 合约代码
	InstrumentID TRsaFtdcInstrumentIDType
	// 本次撤单操作的本地编号
	UserOrderActionLocalID TRsaFtdcUserOrderLocalIDType
	// 报单操作标志
	ActionFlag TRsaFtdcActionFlagType
	// 价格
	LimitPrice TRsaFtdcPriceType
	// 数量变化
	VolumeChange TRsaFtdcVolumeType
	// 前置编号
	FrontID TRsaFtdcFrontIDType
	// 会话编号
	SessionID TRsaFtdcSessionIDType
}

// 期权自对冲录入
type CRsaFtdcInputOptionSelfCloseField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 投资者编码
	InvestorID TRsaFtdcInvestorIDType
	// 合约代码
	InstrumentID TRsaFtdcInstrumentIDType
	// 期权自对冲引用
	OptionSelfCloseRef TRsaFtdcOrderRefType
	// 用户代码
	UserID TRsaFtdcUserIDType
	// 数量
	Volume TRsaFtdcVolumeType
	// 请求编号
	RequestID TRsaFtdcRequestIDType
	// 业务单元
	BusinessUnit TRsaFtdcBusinessUnitType
	// 投机套保标志
	HedgeFlag TRsaFtdcHedgeFlagType
	// 期权行权的头寸是否自对冲
	OptSelfCloseFlag TRsaFtdcOptSelfCloseFlagType
	// 交易所代码
	ExchangeID TRsaFtdcExchangeIDType
	// 投资单元代码
	InvestUnitID TRsaFtdcInvestUnitIDType
	// 资金帐号
	AccountID TRsaFtdcAccountIDType
	// 币种
	Currency TRsaFtdcCurrencyType
	// 客户代码
	ClientID TRsaFtdcClientIDType
	// 下单IP地址
	IPAddress TRsaFtdcIPAddressType
	// 下单MAC地址
	MacAddress TRsaFtdcMacAddressType
}

// 期权自对冲操作
type CRsaFtdcOptionSelfCloseActionField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 投资者编码
	InvestorID TRsaFtdcInvestorIDType
	// 期权自对冲操作引用
	OptionSelfCloseActionRef TRsaFtdcOrderActionRefType
	// 期权自对冲引用
	OptionSelfCloseRef TRsaFtdcOrderRefType
	// 请求编号
	RequestID TRsaFtdcRequestIDType
	// 前置编号
	FrontID TRsaFtdcFrontIDType
	// 会话编号
	SessionID TRsaFtdcSessionIDType
	// 交易所代码
	ExchangeID TRsaFtdcExchangeIDType
	// 期权自对冲操作编号
	OptionSelfCloseSysID TRsaFtdcOrderSysIDType
	// 报单操作标志
	ActionFlag TRsaFtdcActionFlagType
	// 用户代码
	UserID TRsaFtdcUserIDType
	// 合约代码
	InstrumentID TRsaFtdcInstrumentIDType
	// 投资单元代码
	InvestUnitID TRsaFtdcInvestUnitIDType
	// 下单IP地址
	IPAddress TRsaFtdcIPAddressType
	// 下单MAC地址
	MacAddress TRsaFtdcMacAddressType
}

// 席位状态切换
type CRsaFtdcChangeSeatStatusField struct {
	// 使用哪个通道下单
	APIID TRsaFtdcNumberType
	// 下单席位号
	SeatID TRsaFtdcSeatIDType
	// 席位状态方向
	SeatDirection TRsaFtdcSeatDirectionType
}

// 锁定录入
type CRsaFtdcInputLockField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
	// 用户代码
	UserID TRsaFtdcUserIDType
	// 用户本地报单号
	UserOrderLocalID TRsaFtdcUserOrderLocalIDType
	// 用户自定义域
	UserCustom TRsaFtdcCustomType
	// qsc全局编号
	QscId TRsaFtdcNumberType
	// 插件标识
	PluginTaskId TRsaFtdcNumberType
	// 外部标识
	OthMarkId TRsaFtdcOthMarkIdType
	// 交易所代码
	ExchangeID TRsaFtdcExchangeIDType
	// 合约代码
	InstrumentID TRsaFtdcInstrumentIDType
	// 数量
	Volume TRsaFtdcVolumeType
	// 锁定类型
	LockType TRsaFtdcLockTypeType
	// 业务单元
	BusinessUnit TRsaFtdcBusinessUnitType
}

// 用户信息采集
type CRsaFtdcSubmitUserSystemInfoField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 用户代码
	UserID TRsaFtdcUserIDType
	// 信息长度
	ClientSystemInfoLen TRsaFtdcContentLengthType
	// 信息内容
	ClientSystemInfo TRsaFtdcClientSystemInfoType
	// 客户公网ip
	ClientPublicIP TRsaFtdcIPAddressType
	// 客户端口
	ClientIPPort TRsaFtdcPortType
	// 登录日期
	ClientLoginDate TRsaFtdcDateType
	// 登录时间
	ClientLoginTime TRsaFtdcTimeType
	// 客户认证appid
	ClientAppID TRsaFtdcAppIDType
}

// 行情基础属性
type CRsaFtdcMarketDataBaseField struct {
	// 交易日
	TradingDay TRsaFtdcTradingDayType
	// 昨结算
	PreSettlementPrice TRsaFtdcPriceType
	// 昨收盘
	PreClosePrice TRsaFtdcPriceType
	// 昨持仓量
	PreOpenInterest TRsaFtdcLargeVolumeType
	// 昨虚实度
	PreDelta TRsaFtdcRatioType
}

// 行情静态属性
type CRsaFtdcMarketDataStaticField struct {
	// 今开盘
	OpenPrice TRsaFtdcPriceType
	// 最高价
	HighestPrice TRsaFtdcPriceType
	// 最低价
	LowestPrice TRsaFtdcPriceType
	// 今收盘
	ClosePrice TRsaFtdcPriceType
	// 涨停板价
	UpperLimitPrice TRsaFtdcPriceType
	// 跌停板价
	LowerLimitPrice TRsaFtdcPriceType
	// 今结算
	SettlementPrice TRsaFtdcPriceType
	// 今虚实度
	CurrDelta TRsaFtdcRatioType
}

// 行情最新成交属性
type CRsaFtdcMarketDataLastMatchField struct {
	// 最新价
	LastPrice TRsaFtdcPriceType
	// 数量
	Volume TRsaFtdcVolumeType
	// 成交金额
	Turnover TRsaFtdcMoneyType
	// 持仓量
	OpenInterest TRsaFtdcLargeVolumeType
}

// 行情最优价属性
type CRsaFtdcMarketDataBestPriceField struct {
	// 申买价一
	BidPrice1 TRsaFtdcPriceType
	// 申买量一
	BidVolume1 TRsaFtdcVolumeType
	// 申卖价一
	AskPrice1 TRsaFtdcPriceType
	// 申卖量一
	AskVolume1 TRsaFtdcVolumeType
}

// 行情申买二、三属性
type CRsaFtdcMarketDataBid23Field struct {
	// 申买价二
	BidPrice2 TRsaFtdcPriceType
	// 申买量二
	BidVolume2 TRsaFtdcVolumeType
	// 申买价三
	BidPrice3 TRsaFtdcPriceType
	// 申买量三
	BidVolume3 TRsaFtdcVolumeType
}

// 行情申卖二、三属性
type CRsaFtdcMarketDataAsk23Field struct {
	// 申卖价二
	AskPrice2 TRsaFtdcPriceType
	// 申卖量二
	AskVolume2 TRsaFtdcVolumeType
	// 申卖价三
	AskPrice3 TRsaFtdcPriceType
	// 申卖量三
	AskVolume3 TRsaFtdcVolumeType
}

// 行情申买四、五属性
type CRsaFtdcMarketDataBid45Field struct {
	// 申买价四
	BidPrice4 TRsaFtdcPriceType
	// 申买量四
	BidVolume4 TRsaFtdcVolumeType
	// 申买价五
	BidPrice5 TRsaFtdcPriceType
	// 申买量五
	BidVolume5 TRsaFtdcVolumeType
}

// 行情申卖四、五属性
type CRsaFtdcMarketDataAsk45Field struct {
	// 申卖价四
	AskPrice4 TRsaFtdcPriceType
	// 申卖量四
	AskVolume4 TRsaFtdcVolumeType
	// 申卖价五
	AskPrice5 TRsaFtdcPriceType
	// 申卖量五
	AskVolume5 TRsaFtdcVolumeType
}

// 行情更新时间属性
type CRsaFtdcMarketDataUpdateTimeField struct {
	// 合约代码
	InstrumentID TRsaFtdcInstrumentIDType
	// 最后修改时间
	UpdateTime TRsaFtdcTimeType
	// 最后修改毫秒
	UpdateMillisec TRsaFtdcMillisecType
}

// 深度行情
type CRsaFtdcDepthMarketDataField struct {
	// 交易日
	TradingDay TRsaFtdcTradingDayType
	// 昨结算
	PreSettlementPrice TRsaFtdcPriceType
	// 昨收盘
	PreClosePrice TRsaFtdcPriceType
	// 昨持仓量
	PreOpenInterest TRsaFtdcLargeVolumeType
	// 昨虚实度
	PreDelta TRsaFtdcRatioType
	// 今开盘
	OpenPrice TRsaFtdcPriceType
	// 最高价
	HighestPrice TRsaFtdcPriceType
	// 最低价
	LowestPrice TRsaFtdcPriceType
	// 今收盘
	ClosePrice TRsaFtdcPriceType
	// 涨停板价
	UpperLimitPrice TRsaFtdcPriceType
	// 跌停板价
	LowerLimitPrice TRsaFtdcPriceType
	// 今结算
	SettlementPrice TRsaFtdcPriceType
	// 今虚实度
	CurrDelta TRsaFtdcRatioType
	// 最新价
	LastPrice TRsaFtdcPriceType
	// 数量
	Volume TRsaFtdcVolumeType
	// 成交金额
	Turnover TRsaFtdcMoneyType
	// 持仓量
	OpenInterest TRsaFtdcLargeVolumeType
	// 申买价一
	BidPrice1 TRsaFtdcPriceType
	// 申买量一
	BidVolume1 TRsaFtdcVolumeType
	// 申卖价一
	AskPrice1 TRsaFtdcPriceType
	// 申卖量一
	AskVolume1 TRsaFtdcVolumeType
	// 申买价二
	BidPrice2 TRsaFtdcPriceType
	// 申买量二
	BidVolume2 TRsaFtdcVolumeType
	// 申买价三
	BidPrice3 TRsaFtdcPriceType
	// 申买量三
	BidVolume3 TRsaFtdcVolumeType
	// 申卖价二
	AskPrice2 TRsaFtdcPriceType
	// 申卖量二
	AskVolume2 TRsaFtdcVolumeType
	// 申卖价三
	AskPrice3 TRsaFtdcPriceType
	// 申卖量三
	AskVolume3 TRsaFtdcVolumeType
	// 申买价四
	BidPrice4 TRsaFtdcPriceType
	// 申买量四
	BidVolume4 TRsaFtdcVolumeType
	// 申买价五
	BidPrice5 TRsaFtdcPriceType
	// 申买量五
	BidVolume5 TRsaFtdcVolumeType
	// 申卖价四
	AskPrice4 TRsaFtdcPriceType
	// 申卖量四
	AskVolume4 TRsaFtdcVolumeType
	// 申卖价五
	AskPrice5 TRsaFtdcPriceType
	// 申卖量五
	AskVolume5 TRsaFtdcVolumeType
	// 申卖价六
	AskPrice6 TRsaFtdcPriceType
	// 申卖量六
	AskVolume6 TRsaFtdcVolumeType
	// 申卖价七
	AskPrice7 TRsaFtdcPriceType
	// 申卖量七
	AskVolume7 TRsaFtdcVolumeType
	// 申卖价八
	AskPrice8 TRsaFtdcPriceType
	// 申卖量八
	AskVolume8 TRsaFtdcVolumeType
	// 申卖价九
	AskPrice9 TRsaFtdcPriceType
	// 申卖量九
	AskVolume9 TRsaFtdcVolumeType
	// 申买价六
	BidPrice6 TRsaFtdcPriceType
	// 申买量六
	BidVolume6 TRsaFtdcVolumeType
	// 申买价七
	BidPrice7 TRsaFtdcPriceType
	// 申买量七
	BidVolume7 TRsaFtdcVolumeType
	// 申买价八
	BidPrice8 TRsaFtdcPriceType
	// 申买量八
	BidVolume8 TRsaFtdcVolumeType
	// 申买价九
	BidPrice9 TRsaFtdcPriceType
	// 申买量九
	BidVolume9 TRsaFtdcVolumeType
	// 申卖价十
	AskPrice10 TRsaFtdcPriceType
	// 申卖量十
	AskVolume10 TRsaFtdcVolumeType
	// 申买价十
	BidPrice10 TRsaFtdcPriceType
	// 申买量十
	BidVolume10 TRsaFtdcVolumeType
	// 合约代码
	InstrumentID TRsaFtdcInstrumentIDType
	// 最后修改时间
	UpdateTime TRsaFtdcTimeType
	// 最后修改毫秒
	UpdateMillisec TRsaFtdcMillisecType
	// 交易所代码
	ExchangeID TRsaFtdcExchangeIDType
	// 总卖出数量
	VolumeAskLot TRsaFtdcVolumeType
	// 总买入数量
	VolumeBidLot TRsaFtdcVolumeType
	// 合约状态
	InstrumentStatus TRsaFtdcInstrumentStatusType
	// 当前日期
	CalendarDate TRsaFtdcTradingDayType
	// 组播增量序号
	PacketNo TRsaFtdcPacketNoType
	// 合约序号
	InstrumentNo TRsaFtdcNumberType
	// 数据中心代码
	DataCenterID TRsaFtdcDataCenterIDType
	// 更新时间
	UpdateTs TRsaFtdcInt64Type
	// 最新成交量
	LastTraded TRsaFtdcVolumeType
	// 最新成交额
	LastTurnover TRsaFtdcMoneyType
}

// 订阅合约的相关信息
type CRsaFtdcSpecificInstrumentField struct {
	// 交易所代码
	ExchangeID TRsaFtdcExchangeIDType
	// 合约代码
	InstrumentID TRsaFtdcInstrumentIDType
}

// 投资者权限
type CRsaFtdcInvestorRightField struct {
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
	// 是否使用净持仓报单,投顾账户有效
	PositionType TRsaFtdcPositionTypeType
	// 自成交处理方式,资金账户有效
	SelfTradeAvoidType TRsaFtdcSelfTradeAvoidTypeType
	// 交易权限
	TradingRight TRsaFtdcTradingRightType
	// 是否检查资金
	CheckAccount TRsaFtdcBoolType
	// 是否检查持仓
	CheckPosition TRsaFtdcBoolType
	// 报单是否报警
	OrderWarning TRsaFtdcBoolType
	// 资金持仓是否上查
	QryUpAccPosi TRsaFtdcBoolType
	// 投顾选择资金账户类型
	DistributeType TRsaFtdcDistributeTypeType
}

// 交易所代码
type CRsaFtdcMarketDataExchangeIDField struct {
	// 交易所代码
	ExchangeID TRsaFtdcExchangeIDType
}

// 共享内存行情查询
type CRsaFtdcShmDepthMarketDataField struct {
	// 合约代码
	InstrumentID TRsaFtdcInstrumentIDType
	// IP地址
	IPAddress TRsaFtdcIPAddressType
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
	InstrumentID TRsaFtdcInstrumentIDType
	// 买卖方向
	Direction TRsaFtdcDirectionType
	// 价格
	Price TRsaFtdcPriceType
	// 数量
	Volume TRsaFtdcVolumeType
	// 最后修改时间
	UpdateTime TRsaFtdcTimeType
	// 最后修改毫秒
	UpdateMillisec TRsaFtdcMillisecType
}

// 套利合约买卖总量
type CRsaFtdcMarketDataCMBVolumeField struct {
	// 总卖出数量
	VolumeAskLot TRsaFtdcVolumeType
	// 总买入数量
	VolumeBidLot TRsaFtdcVolumeType
}

// 合约状态
type CRsaFtdcQmdInstrumentStateField struct {
	// 交易所代码
	ExchangeID TRsaFtdcExchangeIDType
	// 合约代码
	InstrumentID TRsaFtdcInstrumentIDType
	// 合约交易状态
	InstrumentStatus TRsaFtdcInstrumentStatusType
}

// 订阅合约的相关信息
type CRsaFtdcSubSpecificInstrumentField struct {
	// 交易所代码
	ExchangeID TRsaFtdcExchangeIDType
	// 合约代码
	InstrumentID TRsaFtdcSubInstrumentIDType
}

// 行情申卖六七属性
type CRsaFtdcMarketDataAsk67Field struct {
	// 申卖价六
	AskPrice6 TRsaFtdcPriceType
	// 申卖量六
	AskVolume6 TRsaFtdcVolumeType
	// 申卖价七
	AskPrice7 TRsaFtdcPriceType
	// 申卖量七
	AskVolume7 TRsaFtdcVolumeType
}

// 行情申卖八九属性
type CRsaFtdcMarketDataAsk89Field struct {
	// 申卖价八
	AskPrice8 TRsaFtdcPriceType
	// 申卖量八
	AskVolume8 TRsaFtdcVolumeType
	// 申卖价九
	AskPrice9 TRsaFtdcPriceType
	// 申卖量九
	AskVolume9 TRsaFtdcVolumeType
}

// 行情申买六七属性
type CRsaFtdcMarketDataBid67Field struct {
	// 申买价六
	BidPrice6 TRsaFtdcPriceType
	// 申买量六
	BidVolume6 TRsaFtdcVolumeType
	// 申买价七
	BidPrice7 TRsaFtdcPriceType
	// 申买量七
	BidVolume7 TRsaFtdcVolumeType
}

// 行情申买八九属性
type CRsaFtdcMarketDataBid89Field struct {
	// 申买价八
	BidPrice8 TRsaFtdcPriceType
	// 申买量八
	BidVolume8 TRsaFtdcVolumeType
	// 申买价九
	BidPrice9 TRsaFtdcPriceType
	// 申买量九
	BidVolume9 TRsaFtdcVolumeType
}

// 行情申买申卖十属性
type CRsaFtdcMarketDataAskBid10Field struct {
	// 申卖价十
	AskPrice10 TRsaFtdcPriceType
	// 申卖量十
	AskVolume10 TRsaFtdcVolumeType
	// 申买价十
	BidPrice10 TRsaFtdcPriceType
	// 申买量十
	BidVolume10 TRsaFtdcVolumeType
}

// 行情合约状态
type CRsaFtdcMarketDataInstrumentStatusField struct {
	// 合约状态
	InstrumentStatus TRsaFtdcInstrumentStatusType
}

// Bar行情
type CRsaFtdcBarMarketDataField struct {
	// 交易日
	TradingDay TRsaFtdcTradingDayType
	// 交易所代码
	ExchangeID TRsaFtdcExchangeIDType
	// 合约代码
	InstrumentID TRsaFtdcInstrumentIDType
	// Bar精度
	BarPreces TRsaFtdcBarPrecesType
	// Bar周期
	BarPeriod TRsaFtdcNumberType
	// Bar时间
	BarTime TRsaFtdcInt64Type
	// 更新时间
	UpdateTs TRsaFtdcInt64Type
	// 最后修改时间
	UpdateTime TRsaFtdcTimeType
	// 最后修改毫秒
	UpdateMillisec TRsaFtdcMillisecType
	// 最高价
	HighestPrice TRsaFtdcPriceType
	// 最低价
	LowestPrice TRsaFtdcPriceType
	// 开盘价
	Open TRsaFtdcPriceType
	// 最高价
	High TRsaFtdcPriceType
	// 最低价
	Low TRsaFtdcPriceType
	// 收盘价
	Close TRsaFtdcPriceType
	// 总成交量
	Volume TRsaFtdcVolumeType
	// 最新成交量
	LastTraded TRsaFtdcVolumeType
	// 持仓量
	OpenInterest TRsaFtdcLargeVolumeType
	// 成交额
	Turnover TRsaFtdcMoneyType
	// 最新成交额
	LastTurnover TRsaFtdcMoneyType
}

// 查询Bar行情请求
type CRsaFtdcReqQryBarFromRmdField struct {
	// 交易所代码
	ExchangeID TRsaFtdcExchangeIDType
	// 合约代码
	InstrumentID TRsaFtdcInstrumentIDType
}

// L2Tick行情
type CRsaFtdcL2TickMarketDataField struct {
	// 交易日
	TradingDay TRsaFtdcTradingDayType
	// 交易所代码
	ExchangeID TRsaFtdcExchangeIDType
	// 合约代码
	InstrumentID TRsaFtdcInstrumentIDType
	// Tick序号
	TickIndex TRsaFtdcInt64Type
	// Tick时间
	TickTime TRsaFtdcInt64Type
	// Tick类型
	TickType TRsaFtdcTickTypeType
	// 买方报单编号
	BuyOrderSysID TRsaFtdcOrderSysIDType
	// 卖方报单编号
	SellOrderSysID TRsaFtdcOrderSysIDType
	// 价格
	Price TRsaFtdcPriceType
	// 数量
	Volume TRsaFtdcVolumeType
	// Tick买卖标志
	TickBSFlag TRsaFtdcTickBSFlagType
	// 通道号
	ChannelID TRsaFtdcNumberType
}

// 成交
type CRsaFtdcTradeField struct {
	// 交易日
	TradingDay TRsaFtdcTradingDayType
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
	// 用户代码
	UserID TRsaFtdcUserIDType
	// 用户本地报单号
	UserOrderLocalID TRsaFtdcUserOrderLocalIDType
	// 用户自定义域
	UserCustom TRsaFtdcCustomType
	// qsc全局编号
	QscId TRsaFtdcNumberType
	// 插件标识
	PluginTaskId TRsaFtdcNumberType
	// 外部标识
	OthMarkId TRsaFtdcOthMarkIdType
	// 交易所代码
	ExchangeID TRsaFtdcExchangeIDType
	// 成交编号
	TradeID TRsaFtdcTradeIDType
	// 报单编号
	OrderSysID TRsaFtdcOrderSysIDType
	// 合约代码
	InstrumentID TRsaFtdcInstrumentIDType
	// 买卖方向
	Direction TRsaFtdcDirectionType
	// 开平标志
	OffsetFlag TRsaFtdcOffsetFlagType
	// 投机套保标志
	HedgeFlag TRsaFtdcHedgeFlagType
	// 成交价格
	TradePrice TRsaFtdcPriceType
	// 成交数量
	TradeVolume TRsaFtdcVolumeType
	// 成交时间
	TradeTime TRsaFtdcTimeType
	// 清算会员编号
	ClearingPartID TRsaFtdcParticipantIDType
	// 成交金额
	TradeAmnt TRsaFtdcMoneyType
	// 交易类型
	TradeType TRsaFtdcTradeTypeType
	// 成交数量（小数）
	TradeVolumeP TRsaFtdcPointVolumeType
	// 营业部代码
	BranchID TRsaFtdcBranchIDType
	// 营业部名称
	BranchName TRsaFtdcBranchNameType
}

// 报单
type CRsaFtdcOrderField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
	// 用户代码
	UserID TRsaFtdcUserIDType
	// 用户本地报单号
	UserOrderLocalID TRsaFtdcUserOrderLocalIDType
	// 用户自定义域
	UserCustom TRsaFtdcCustomType
	// qsc全局编号
	QscId TRsaFtdcNumberType
	// 插件标识
	PluginTaskId TRsaFtdcNumberType
	// 外部标识
	OthMarkId TRsaFtdcOthMarkIdType
	// 交易所代码
	ExchangeID TRsaFtdcExchangeIDType
	// 系统报单编号
	OrderSysID TRsaFtdcOrderSysIDType
	// 合约代码
	InstrumentID TRsaFtdcInstrumentIDType
	// 报单类型
	OrderPriceType TRsaFtdcOrderPriceTypeType
	// 买卖方向
	Direction TRsaFtdcDirectionType
	// 开平标志
	OffsetFlag TRsaFtdcOffsetFlagType
	// 投机套保标志
	HedgeFlag TRsaFtdcHedgeFlagType
	// 价格
	LimitPrice TRsaFtdcPriceType
	// 数量
	Volume TRsaFtdcVolumeType
	// 有效期类型
	TimeCondition TRsaFtdcTimeConditionType
	// GTD日期
	GTDDate TRsaFtdcDateType
	// 成交量类型
	VolumeCondition TRsaFtdcVolumeConditionType
	// 最小成交量
	MinVolume TRsaFtdcVolumeType
	// 止损价
	StopPrice TRsaFtdcPriceType
	// 强平原因
	ForceCloseReason TRsaFtdcForceCloseReasonType
	// 自动挂起标志
	IsAutoSuspend TRsaFtdcBoolType
	// 业务单元
	BusinessUnit TRsaFtdcBusinessUnitType
	// 业务类别
	BusinessType TRsaFtdcBusinessTypeType
	// 数量（小数）
	VolumeP TRsaFtdcPointVolumeType
	// 策略限制价格
	AlgoLimitPrice TRsaFtdcPriceType
	// 策略限制tick数量
	AlgoLimitPriceTick TRsaFtdcNumberType
	// 策略限制时间(ms)
	AlgoLimitTime TRsaFtdcNumberType
	// 交易日
	TradingDay TRsaFtdcTradingDayType
	// 插入时间
	InsertTime TRsaFtdcTimeType
	// 报单来源
	OrderSource TRsaFtdcOrderSourceType
	// 报单状态
	OrderStatus TRsaFtdcOrderStatusType
	// 撤销时间
	CancelTime TRsaFtdcTimeType
	// 撤单用户编号
	CancelUserID TRsaFtdcUserIDType
	// 已经成交数量
	VolumeTraded TRsaFtdcVolumeType
	// 队列剩余数量
	VolumeRemain TRsaFtdcVolumeType
	// 已经撤单数量
	VolumeCancled TRsaFtdcVolumeType
	// 错误代码
	ErrorID TRsaFtdcErrorIDType
	// 错误信息
	ErrorMsg TRsaFtdcErrorMsgType
	// 比上次增加的撤单数量
	VolumeIncCancled TRsaFtdcVolumeType
	// 比上次增加的成交数量
	VolumeIncTraded TRsaFtdcVolumeType
	// 资金帐号
	AccountID TRsaFtdcAccountIDType
	// 前置编号
	FrontID TRsaFtdcFrontIDType
	// 会话编号
	SessionID TRsaFtdcSessionIDType
	// 请求编号
	RequestID TRsaFtdcRequestIDType
	// 每手手续费
	EachFee TRsaFtdcMoneyType
	// 每手保证金
	EachMargin TRsaFtdcMoneyType
	// 每手期权权利金收支
	EachPremium TRsaFtdcMoneyType
	// 下单IP地址
	IPAddress TRsaFtdcIPAddressType
	// 下单MAC地址
	MacAddress TRsaFtdcMacAddressType
	// 营业部代码
	BranchID TRsaFtdcBranchIDType
	// 记录编号
	RecNum TRsaFtdcSequenceNoType
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
	TradingDay TRsaFtdcDateType
	// 数据中心代码
	DataCenterID TRsaFtdcDataCenterIDType
	// 回退起始序列号
	StartSequenceNo TRsaFtdcSequenceNoType
	// 回退结束序列号
	EndSequenceNo TRsaFtdcSequenceNoType
}

// 信息分发
type CRsaFtdcDisseminationField struct {
	// 序列系列号
	SequenceSeries TRsaFtdcSequenceSeriesType
	// 序列号
	SequenceNo TRsaFtdcSequenceNoType
}

// 警告消息通知
type CRsaFtdcMessageNotifyField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 修改用户编号
	UserID TRsaFtdcUserIDType
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
	// 累加次数
	Nums TRsaFtdcTotalNumsType
	// 警告级别
	WarnLevel TRsaFtdcErrorLevelType
	// 警告编号
	WarnID TRsaFtdcErrorIDType
	// 警告信息
	WarnMsg TRsaFtdcLargeErrorMsgType
	// 警告时间
	WarnTime TRsaFtdcTimeType
}

// 工作流指令通知
type CRsaFtdcAdviceOrderField struct {
	// 投顾用户代码
	InvestorAdvicer TRsaFtdcUserIDType
	// 投资经理
	InvestorManager TRsaFtdcUserIDType
	// 交易员
	UserTrader TRsaFtdcUserIDType
	// 冻结数量
	VolumeFrozen TRsaFtdcVolumeType
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
	// 用户代码
	UserID TRsaFtdcUserIDType
	// 用户本地报单号
	UserOrderLocalID TRsaFtdcUserOrderLocalIDType
	// 用户自定义域
	UserCustom TRsaFtdcCustomType
	// qsc全局编号
	QscId TRsaFtdcNumberType
	// 插件标识
	PluginTaskId TRsaFtdcNumberType
	// 外部标识
	OthMarkId TRsaFtdcOthMarkIdType
	// 交易所代码
	ExchangeID TRsaFtdcExchangeIDType
	// 系统报单编号
	OrderSysID TRsaFtdcOrderSysIDType
	// 合约代码
	InstrumentID TRsaFtdcInstrumentIDType
	// 报单类型
	OrderPriceType TRsaFtdcOrderPriceTypeType
	// 买卖方向
	Direction TRsaFtdcDirectionType
	// 开平标志
	OffsetFlag TRsaFtdcOffsetFlagType
	// 投机套保标志
	HedgeFlag TRsaFtdcHedgeFlagType
	// 价格
	LimitPrice TRsaFtdcPriceType
	// 数量
	Volume TRsaFtdcVolumeType
	// 有效期类型
	TimeCondition TRsaFtdcTimeConditionType
	// GTD日期
	GTDDate TRsaFtdcDateType
	// 成交量类型
	VolumeCondition TRsaFtdcVolumeConditionType
	// 最小成交量
	MinVolume TRsaFtdcVolumeType
	// 止损价
	StopPrice TRsaFtdcPriceType
	// 强平原因
	ForceCloseReason TRsaFtdcForceCloseReasonType
	// 自动挂起标志
	IsAutoSuspend TRsaFtdcBoolType
	// 业务单元
	BusinessUnit TRsaFtdcBusinessUnitType
	// 业务类别
	BusinessType TRsaFtdcBusinessTypeType
	// 数量（小数）
	VolumeP TRsaFtdcPointVolumeType
	// 策略限制价格
	AlgoLimitPrice TRsaFtdcPriceType
	// 策略限制tick数量
	AlgoLimitPriceTick TRsaFtdcNumberType
	// 策略限制时间(ms)
	AlgoLimitTime TRsaFtdcNumberType
	// 交易日
	TradingDay TRsaFtdcTradingDayType
	// 插入时间
	InsertTime TRsaFtdcTimeType
	// 报单来源
	OrderSource TRsaFtdcOrderSourceType
	// 报单状态
	OrderStatus TRsaFtdcOrderStatusType
	// 撤销时间
	CancelTime TRsaFtdcTimeType
	// 撤单用户编号
	CancelUserID TRsaFtdcUserIDType
	// 已经成交数量
	VolumeTraded TRsaFtdcVolumeType
	// 队列剩余数量
	VolumeRemain TRsaFtdcVolumeType
	// 已经撤单数量
	VolumeCancled TRsaFtdcVolumeType
	// 错误代码
	ErrorID TRsaFtdcErrorIDType
	// 错误信息
	ErrorMsg TRsaFtdcErrorMsgType
	// 比上次增加的撤单数量
	VolumeIncCancled TRsaFtdcVolumeType
	// 比上次增加的成交数量
	VolumeIncTraded TRsaFtdcVolumeType
	// 资金帐号
	AccountID TRsaFtdcAccountIDType
	// 前置编号
	FrontID TRsaFtdcFrontIDType
	// 会话编号
	SessionID TRsaFtdcSessionIDType
	// 请求编号
	RequestID TRsaFtdcRequestIDType
	// 每手手续费
	EachFee TRsaFtdcMoneyType
	// 每手保证金
	EachMargin TRsaFtdcMoneyType
	// 每手期权权利金收支
	EachPremium TRsaFtdcMoneyType
	// 下单IP地址
	IPAddress TRsaFtdcIPAddressType
	// 下单MAC地址
	MacAddress TRsaFtdcMacAddressType
	// 营业部代码
	BranchID TRsaFtdcBranchIDType
	// 记录编号
	RecNum TRsaFtdcSequenceNoType
	// 指令单用户信息
	AdviceUserInfo TRsaFtdcAdviceUserInfoType
	// 硬盘序列号
	HDSerialID TRsaFtdcHDSerialIDType
	// 营业部名称
	BranchName TRsaFtdcBranchNameType
	// 报单提交状态
	OrderSubmitStatus TRsaFtdcOrderSubmitStatusType
	// 父投资者编号
	ParentInvestorID TRsaFtdcInvestorIDType
}

// 执行宣告通知
type CRsaFtdcExecOrderField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
	// 用户代码
	UserID TRsaFtdcUserIDType
	// 用户本地报单号
	UserOrderLocalID TRsaFtdcUserOrderLocalIDType
	// 用户自定义域
	UserCustom TRsaFtdcCustomType
	// qsc全局编号
	QscId TRsaFtdcNumberType
	// 插件标识
	PluginTaskId TRsaFtdcNumberType
	// 外部标识
	OthMarkId TRsaFtdcOthMarkIdType
	// 交易所代码
	ExchangeID TRsaFtdcExchangeIDType
	// 系统报单编号
	OrderSysID TRsaFtdcOrderSysIDType
	// 合约代码
	InstrumentID TRsaFtdcInstrumentIDType
	// 开平标志
	OffsetFlag TRsaFtdcOffsetFlagType
	// 投机套保标志
	HedgeFlag TRsaFtdcHedgeFlagType
	// 数量
	Volume TRsaFtdcVolumeType
	// 业务单元
	BusinessUnit TRsaFtdcBusinessUnitType
	// 执行类型
	ActionType TRsaFtdcActionTypeType
	// 保留头寸申请的持仓方向
	PosiDirection TRsaFtdcPosiDirectionType
	// 期权行权后是否保留期货头寸的标记
	ReservePositionFlag TRsaFtdcReservePositionFlagTypeType
	// 期权行权后生成的头寸是否自动平仓
	CloseFlag TRsaFtdcCloseFlagTypeType
	// 数量(小数)
	VolumeP TRsaFtdcPointVolumeType
	// 执行结果
	ExecResultType TRsaFtdcExecResultTypeType
	// 交易日
	TradingDay TRsaFtdcTradingDayType
	// 插入时间
	InsertTime TRsaFtdcTimeType
	// 报单来源
	OrderSource TRsaFtdcOrderSourceType
	// 报单状态
	OrderStatus TRsaFtdcOrderStatusType
	// 撤销时间
	CancelTime TRsaFtdcTimeType
	// 撤单用户编号
	CancelUserID TRsaFtdcUserIDType
	// 已经成交数量
	VolumeTraded TRsaFtdcVolumeType
	// 队列剩余数量
	VolumeRemain TRsaFtdcVolumeType
	// 已经撤单数量
	VolumeCancled TRsaFtdcVolumeType
	// 错误代码
	ErrorID TRsaFtdcErrorIDType
	// 错误信息
	ErrorMsg TRsaFtdcErrorMsgType
	// 比上次增加的撤单数量
	VolumeIncCancled TRsaFtdcVolumeType
	// 比上次增加的成交数量
	VolumeIncTraded TRsaFtdcVolumeType
	// 资金帐号
	AccountID TRsaFtdcAccountIDType
	// 前置编号
	FrontID TRsaFtdcFrontIDType
	// 会话编号
	SessionID TRsaFtdcSessionIDType
	// 请求编号
	RequestID TRsaFtdcRequestIDType
	// 每手手续费
	EachFee TRsaFtdcMoneyType
	// 每手保证金
	EachMargin TRsaFtdcMoneyType
	// 每手期权权利金收支
	EachPremium TRsaFtdcMoneyType
	// 下单IP地址
	IPAddress TRsaFtdcIPAddressType
	// 下单MAC地址
	MacAddress TRsaFtdcMacAddressType
	// 营业部代码
	BranchID TRsaFtdcBranchIDType
	// 记录编号
	RecNum TRsaFtdcSequenceNoType
	// 报单提交状态
	OrderSubmitStatus TRsaFtdcOrderSubmitStatusType
}

// 询价流通知
type CRsaFtdcForQuoteRspField struct {
	// 交易日
	TradingDay TRsaFtdcDateType
	// 合约代码
	InstrumentID TRsaFtdcInstrumentIDType
	// 询价编号
	OrderSysID TRsaFtdcOrderSysIDType
	// 交易所代码
	ExchangeID TRsaFtdcExchangeIDType
	// 插入时间
	InsertTime TRsaFtdcTimeType
}

// 报价通知
type CRsaFtdcQuoteField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
	// 用户代码
	UserID TRsaFtdcUserIDType
	// 用户本地报单号
	UserOrderLocalID TRsaFtdcUserOrderLocalIDType
	// 用户自定义域
	UserCustom TRsaFtdcCustomType
	// qsc全局编号
	QscId TRsaFtdcNumberType
	// 插件标识
	PluginTaskId TRsaFtdcNumberType
	// 外部标识
	OthMarkId TRsaFtdcOthMarkIdType
	// 交易所代码
	ExchangeID TRsaFtdcExchangeIDType
	// 合约代码
	InstrumentID TRsaFtdcInstrumentIDType
	// 系统报单编号
	OrderSysID TRsaFtdcOrderSysIDType
	// 应价编号
	ForQuoteSysID TRsaFtdcOrderSysIDType
	// 卖价格
	AskPrice TRsaFtdcPriceType
	// 卖数量
	AskVolume TRsaFtdcVolumeType
	// 卖开平标志
	AskOffsetFlag TRsaFtdcOffsetFlagType
	// 卖投机套保标志
	AskHedgeFlag TRsaFtdcHedgeFlagType
	// 买价格
	BidPrice TRsaFtdcPriceType
	// 买数量
	BidVolume TRsaFtdcVolumeType
	// 买开平标志
	BidOffsetFlag TRsaFtdcOffsetFlagType
	// 买投机套保标志
	BidHedgeFlag TRsaFtdcHedgeFlagType
	// 业务单元
	BusinessUnit TRsaFtdcBusinessUnitType
	// 买方本地报单编号
	BidOrderRef TRsaFtdcUserOrderLocalIDType
	// 卖方本地报单编号
	AskOrderRef TRsaFtdcUserOrderLocalIDType
	// 交易日
	TradingDay TRsaFtdcTradingDayType
	// 插入时间
	InsertTime TRsaFtdcTimeType
	// 报单来源
	OrderSource TRsaFtdcOrderSourceType
	// 报单状态
	OrderStatus TRsaFtdcOrderStatusType
	// 撤销时间
	CancelTime TRsaFtdcTimeType
	// 撤单用户编号
	CancelUserID TRsaFtdcUserIDType
	// 已经成交数量
	VolumeTraded TRsaFtdcVolumeType
	// 队列剩余数量
	VolumeRemain TRsaFtdcVolumeType
	// 已经撤单数量
	VolumeCancled TRsaFtdcVolumeType
	// 错误代码
	ErrorID TRsaFtdcErrorIDType
	// 错误信息
	ErrorMsg TRsaFtdcErrorMsgType
	// 比上次增加的撤单数量
	VolumeIncCancled TRsaFtdcVolumeType
	// 比上次增加的成交数量
	VolumeIncTraded TRsaFtdcVolumeType
	// 资金帐号
	AccountID TRsaFtdcAccountIDType
	// 前置编号
	FrontID TRsaFtdcFrontIDType
	// 会话编号
	SessionID TRsaFtdcSessionIDType
	// 请求编号
	RequestID TRsaFtdcRequestIDType
	// 每手手续费
	EachFee TRsaFtdcMoneyType
	// 每手保证金
	EachMargin TRsaFtdcMoneyType
	// 每手期权权利金收支
	EachPremium TRsaFtdcMoneyType
	// 下单IP地址
	IPAddress TRsaFtdcIPAddressType
	// 下单MAC地址
	MacAddress TRsaFtdcMacAddressType
	// 营业部代码
	BranchID TRsaFtdcBranchIDType
	// 记录编号
	RecNum TRsaFtdcSequenceNoType
	// 买方系统报单编号
	BidOrderSysID TRsaFtdcOrderSysIDType
	// 卖方系统报单编号
	AskOrderSysID TRsaFtdcOrderSysIDType
	// 报单提交状态
	OrderSubmitStatus TRsaFtdcOrderSubmitStatusType
}

// 期权自对冲
type CRsaFtdcNtfOptionSelfCloseField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 投资者编码
	InvestorID TRsaFtdcInvestorIDType
	// 合约代码
	InstrumentID TRsaFtdcInstrumentIDType
	// 期权自对冲引用
	OptionSelfCloseRef TRsaFtdcOrderRefType
	// 用户代码
	UserID TRsaFtdcUserIDType
	// 数量
	Volume TRsaFtdcVolumeType
	// 请求编号
	RequestID TRsaFtdcRequestIDType
	// 业务单元
	BusinessUnit TRsaFtdcBusinessUnitType
	// 投机套保标志
	HedgeFlag TRsaFtdcHedgeFlagType
	// 期权行权的头寸是否自对冲
	OptSelfCloseFlag TRsaFtdcOptSelfCloseFlagType
	// 本地期权自对冲编号
	OptionSelfCloseLocalID TRsaFtdcOrderLocalIDType
	// 交易所代码
	ExchangeID TRsaFtdcExchangeIDType
	// 会员编号
	ParticipantID TRsaFtdcParticipantIDType
	// 客户代码
	ClientID TRsaFtdcClientIDType
	// 合约在交易所的代码
	ExchangeInstID TRsaFtdcExchangeInstIDType
	// 交易所交易员代码
	TraderID TRsaFtdcTraderIDType
	// 安装编号
	InstallID TRsaFtdcInstallIDType
	// 期权自对冲提交状态
	OrderSubmitStatus TRsaFtdcOrderSubmitStatusType
	// 报单提示序号
	NotifySequence TRsaFtdcSequenceNoType
	// 交易日
	TradingDay TRsaFtdcDateType
	// 结算编号
	SettlementID TRsaFtdcSettlementIDType
	// 期权自对冲操作编号
	OptionSelfCloseSysID TRsaFtdcOrderSysIDType
	// 报单日期
	InsertDate TRsaFtdcDateType
	// 插入时间
	InsertTime TRsaFtdcTimeType
	// 撤销时间
	CancelTime TRsaFtdcTimeType
	// 自对冲结果
	ExecResultType TRsaFtdcExecResultTypeType
	// 清算会员编号
	ClearingPartID TRsaFtdcParticipantIDType
	// 序号
	SequenceNo TRsaFtdcSequenceNoType
	// 前置编号
	FrontID TRsaFtdcFrontIDType
	// 会话编号
	SessionID TRsaFtdcSessionIDType
	// 用户端产品信息
	UserProductInfo TRsaFtdcProductInfoType
	// 状态信息
	StatusMsg TRsaFtdcStatusMsgType
	// 操作用户代码
	ActiveUserID TRsaFtdcUserIDType
	// 经纪公司报单编号
	BrokerOptionSelfCloseSeq TRsaFtdcSequenceNoType
	// 营业部编号
	BranchID TRsaFtdcBranchIDType
	// 投资单元代码
	InvestUnitID TRsaFtdcInvestUnitIDType
	// 资金帐号
	AccountID TRsaFtdcAccountIDType
	// 币种
	Currency TRsaFtdcCurrencyType
	// 下单IP地址
	IPAddress TRsaFtdcIPAddressType
	// 下单MAC地址
	MacAddress TRsaFtdcMacAddressType
}

// 锁定
type CRsaFtdcLockField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
	// 用户代码
	UserID TRsaFtdcUserIDType
	// 用户本地报单号
	UserOrderLocalID TRsaFtdcUserOrderLocalIDType
	// 用户自定义域
	UserCustom TRsaFtdcCustomType
	// qsc全局编号
	QscId TRsaFtdcNumberType
	// 插件标识
	PluginTaskId TRsaFtdcNumberType
	// 外部标识
	OthMarkId TRsaFtdcOthMarkIdType
	// 交易所代码
	ExchangeID TRsaFtdcExchangeIDType
	// 合约代码
	InstrumentID TRsaFtdcInstrumentIDType
	// 数量
	Volume TRsaFtdcVolumeType
	// 锁定类型
	LockType TRsaFtdcLockTypeType
	// 业务单元
	BusinessUnit TRsaFtdcBusinessUnitType
	// 交易日
	TradingDay TRsaFtdcTradingDayType
	// 锁定编号
	LockSysID TRsaFtdcOrderSysIDType
	// 插入时间
	InsertTime TRsaFtdcTimeType
	// 锁定状态
	LockStatus TRsaFtdcOrderActionStatusType
	// 撤销时间
	CancelTime TRsaFtdcTimeType
	// 资金帐号
	AccountID TRsaFtdcAccountIDType
	// 前置编号
	FrontID TRsaFtdcFrontIDType
	// 会话编号
	SessionID TRsaFtdcSessionIDType
	// 请求编号
	RequestID TRsaFtdcRequestIDType
	// 下单IP地址
	IPAddress TRsaFtdcIPAddressType
	// 下单MAC地址
	MacAddress TRsaFtdcMacAddressType
	// 营业部代码
	BranchID TRsaFtdcBranchIDType
	// 错误代码
	ErrorID TRsaFtdcErrorIDType
	// 错误信息
	ErrorMsg TRsaFtdcErrorMsgType
	// 记录编号
	RecNum TRsaFtdcSequenceNoType
}

// 报单查询
type CRsaFtdcQryOrderField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 用户代码
	UserID TRsaFtdcUserIDType
	// 交易所代码
	ExchangeID TRsaFtdcExchangeIDType
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
	// 报单编号
	OrderSysID TRsaFtdcOrderSysIDType
	// 合约代码
	InstrumentID TRsaFtdcInstrumentIDType
	// 开始时间
	InsertTimeStart TRsaFtdcTimeType
	// 结束时间
	InsertTimeEnd TRsaFtdcTimeType
}

// 成交查询
type CRsaFtdcQryTradeField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 用户代码
	UserID TRsaFtdcUserIDType
	// 交易所代码
	ExchangeID TRsaFtdcExchangeIDType
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
	// 成交编号
	TradeID TRsaFtdcTradeIDType
	// 合约代码
	InstrumentID TRsaFtdcInstrumentIDType
}

// 合约查询
type CRsaFtdcQryInstrumentField struct {
	// 交易所代码
	ExchangeID TRsaFtdcExchangeIDType
	// 产品代码
	ProductID TRsaFtdcProductIDType
	// 合约代码
	InstrumentID TRsaFtdcInstrumentIDType
}

// 合约查询应答
type CRsaFtdcRspInstrumentField struct {
	// 交易所代码
	ExchangeID TRsaFtdcExchangeIDType
	// 品种代码
	ProductID TRsaFtdcProductIDType
	// 品种名称
	ProductName TRsaFtdcProductNameType
	// 合约代码
	InstrumentID TRsaFtdcInstrumentIDType
	// 合约名称
	InstrumentName TRsaFtdcInstrumentNameType
	// 交割年份
	DeliveryYear TRsaFtdcYearType
	// 交割月
	DeliveryMonth TRsaFtdcMonthType
	// 限价单最大下单量
	MaxLimitOrderVolume TRsaFtdcVolumeType
	// 限价单最小下单量
	MinLimitOrderVolume TRsaFtdcVolumeType
	// 市价单最大下单量
	MaxMarketOrderVolume TRsaFtdcVolumeType
	// 市价单最小下单量
	MinMarketOrderVolume TRsaFtdcVolumeType
	// 数量乘数
	VolumeMultiple TRsaFtdcVolumeMultipleType
	// 报价单位
	PriceTick TRsaFtdcPriceTickType
	// 币种
	Currency TRsaFtdcCurrencyType
	// 多头限仓
	LongPosLimit TRsaFtdcVolumeType
	// 空头限仓
	ShortPosLimit TRsaFtdcVolumeType
	// 跌停板价
	LowerLimitPrice TRsaFtdcPriceType
	// 涨停板价
	UpperLimitPrice TRsaFtdcPriceType
	// 昨结算
	PreSettlementPrice TRsaFtdcPriceType
	// 合约交易状态
	InstrumentStatus TRsaFtdcInstrumentStatusType
	// 创建日
	CreateDate TRsaFtdcDateType
	// 上市日
	OpenDate TRsaFtdcDateType
	// 到期日
	ExpireDate TRsaFtdcDateType
	// 开始交割日
	StartDelivDate TRsaFtdcDateType
	// 最后交割日
	EndDelivDate TRsaFtdcDateType
	// 挂牌基准价
	BasisPrice TRsaFtdcPriceType
	// 当前是否交易
	IsTrading TRsaFtdcBoolType
	// 基础商品代码
	UnderlyingInstrID TRsaFtdcInstrumentIDType
	// 基础商品乘数
	UnderlyingMultiple TRsaFtdcUnderlyingMultipleType
	// 持仓类型
	PositionType TRsaFtdcPositionTypeType
	// 执行价
	StrikePrice TRsaFtdcPriceType
	// 期权类型
	OptionsType TRsaFtdcOptionsTypeType
	// 汇率
	ExchangeRate TRsaFtdcRatioType
	// 产品类型
	ProductClass TRsaFtdcProductClassType
	// 期权行权方式
	OptionsMode TRsaFtdcOptionsModeType
	// 保证金算法类型
	MarginCombType TRsaFtdcMarginCombTypeType
	// 多头保证金率
	LongMarginRate TRsaFtdcRatioType
	// 空头保证金率
	ShortMarginRate TRsaFtdcRatioType
	// 多头保证金按手数
	LongMarginAmt TRsaFtdcRatioType
	// 空头保证金按手数
	ShortMarginAmt TRsaFtdcRatioType
}

// 投资者资金查询
type CRsaFtdcQryInvestorAccountField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 用户代码
	UserID TRsaFtdcUserIDType
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
}

// 投资者资金应答
type CRsaFtdcRspInvestorAccountField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
	// 资金帐号
	AccountID TRsaFtdcAccountIDType
	// 上次结算准备金
	PreBalance TRsaFtdcMoneyType
	// 上日可用资金
	PreAvailable TRsaFtdcMoneyType
	// 入金金额
	Deposit TRsaFtdcMoneyType
	// 出金金额
	Withdraw TRsaFtdcMoneyType
	// 占用保证金
	Margin TRsaFtdcMoneyType
	// 期权权利金收支
	Premium TRsaFtdcMoneyType
	// 手续费
	Fee TRsaFtdcMoneyType
	// 冻结的保证金
	FrozenMargin TRsaFtdcMoneyType
	// 冻结权利金
	FrozenPremium TRsaFtdcMoneyType
	// 冻结手续费
	FrozenFee TRsaFtdcMoneyType
	// 平仓盈亏
	CloseProfit TRsaFtdcMoneyType
	// 持仓盈亏
	PositionProfit TRsaFtdcMoneyType
	// 可用资金
	Available TRsaFtdcMoneyType
	// 结算准备金
	Balance TRsaFtdcMoneyType
	// 多头占用保证金
	LongMargin TRsaFtdcMoneyType
	// 空头占用保证金
	ShortMargin TRsaFtdcMoneyType
	// 多头冻结的保证金
	LongFrozenMargin TRsaFtdcMoneyType
	// 空头冻结的保证金
	ShortFrozenMargin TRsaFtdcMoneyType
	// 动态权益
	DynamicRights TRsaFtdcMoneyType
	// 风险度
	Risk TRsaFtdcRatioType
	// 其他费用
	OtherFee TRsaFtdcMoneyType
	// 质押金额
	Mortgage TRsaFtdcMoneyType
	// 币种
	Currency TRsaFtdcCurrencyType
	// 可取资金
	WithdrawQuota TRsaFtdcMoneyType
	// 冻结的资金
	FrozenCash TRsaFtdcMoneyType
	// 账户余额
	CurrentBalance TRsaFtdcMoneyType
	// 总资产
	AssetBalance TRsaFtdcMoneyType
	// 总市值
	MarketValue TRsaFtdcMoneyType
	// 优先资金
	Prioprity TRsaFtdcMoneyType
	// 现金资产
	FundAsset TRsaFtdcMoneyType
	// 上海限购额度
	PurQuotaSh TRsaFtdcMoneyType
	// 上海可用限购额度
	EnPurQuotaSh TRsaFtdcMoneyType
	// 深圳限购额度
	PurQuotaSz TRsaFtdcMoneyType
	// 深圳可用限购额度
	EnPurQuotaSz TRsaFtdcMoneyType
}

// 可用投资者查询
type CRsaFtdcQryUserInvestorField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 用户代码
	UserID TRsaFtdcUserIDType
}

// 可用投资者应答
type CRsaFtdcRspUserInvestorField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 交易用户代码
	UserID TRsaFtdcUserIDType
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
	// 默认用户代码
	NextUserID TRsaFtdcUserIDType
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
	BrokerID TRsaFtdcBrokerIDType
	// 用户代码
	UserID TRsaFtdcUserIDType
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
}

// 投资者关系应答
type CRsaFtdcRspInvestorRelationField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
	// 父投资者编号
	ParentInvestorID TRsaFtdcInvestorIDType
	// 顺序编号
	SpecifyNo TRsaFtdcSequenceNoType
}

// 交易所查询请求
type CRsaFtdcQryExchangeField struct {
	// 交易所代码
	ExchangeID TRsaFtdcExchangeIDType
}

// 交易所查询应答
type CRsaFtdcRspExchangeField struct {
	// 交易日
	TradingDay TRsaFtdcTradingDayType
	// 交易所代码
	ExchangeID TRsaFtdcExchangeIDType
	// 交易所名称
	ExchangeName TRsaFtdcExchangeNameType
}

// 投资者持仓查询请求
type CRsaFtdcQryInvestorPositionField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 用户代码
	UserID TRsaFtdcUserIDType
	// 交易所代码
	ExchangeID TRsaFtdcExchangeIDType
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
	// 合约代码
	InstrumentID TRsaFtdcInstrumentIDType
}

// 投资者持仓查询应答
type CRsaFtdcRspInvestorPositionField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 交易所代码
	ExchangeID TRsaFtdcExchangeIDType
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
	// 合约代码
	InstrumentID TRsaFtdcInstrumentIDType
	// 买卖方向
	Direction TRsaFtdcDirectionType
	// 投机套保标志
	HedgeFlag TRsaFtdcHedgeFlagType
	// 占用保证金
	UsedMargin TRsaFtdcMoneyType
	// 总持仓量
	Position TRsaFtdcVolumeType
	// 总持仓成本
	PositionCost TRsaFtdcPriceType
	// 初始昨持仓量(当日不变)
	YdPosition TRsaFtdcVolumeType
	// 初始昨日持仓成本(当日不变)
	YdPositionCost TRsaFtdcMoneyType
	// 冻结的保证金
	FrozenMargin TRsaFtdcMoneyType
	// 开仓冻结持仓
	FrozenPosition TRsaFtdcVolumeType
	// 平仓冻结持仓
	FrozenClosing TRsaFtdcVolumeType
	// 持仓盈亏
	PositionProfit TRsaFtdcMoneyType
	// 冻结的权利金
	FrozenPremium TRsaFtdcMoneyType
	// 最后一笔成交编号
	LastTradeID TRsaFtdcTradeIDType
	// 最后一笔本地报单编号
	LastOrderLocalID TRsaFtdcOrderLocalIDType
	// 总持仓可平仓数量(包括平仓冻结持仓)
	PositionClose TRsaFtdcVolumeType
	// 昨持仓可平仓数量(包括平仓冻结持仓)
	YdPositionClose TRsaFtdcVolumeType
	// 昨持仓平仓冻结持仓
	YdFrozenClosing TRsaFtdcVolumeType
	// 今日开仓数量(不包括冻结)
	OpenVolume TRsaFtdcVolumeType
	// 今日平仓数量(包括昨持仓的平仓,不包括冻结)
	CloseVolume TRsaFtdcVolumeType
	// 平仓盈亏
	CloseProfit TRsaFtdcMoneyType
	// 执行冻结持仓
	StrikeFrozenPosition TRsaFtdcVolumeType
	// 放弃执行冻结持仓
	AbandonFrozenPosition TRsaFtdcVolumeType
	// 成本价
	CostPrice TRsaFtdcPriceType
	// 开仓成本
	OpenCost TRsaFtdcMoneyType
	// 今持仓(包含冻结)
	TodayPosition TRsaFtdcVolumeType
	// 今持仓冻结持仓
	TodayFrozenClosing TRsaFtdcVolumeType
	// 总持仓量(小数)
	PositionP TRsaFtdcPointVolumeType
}

// 用户查询
type CRsaFtdcQryUserField struct {
	// 交易用户代码
	StartUserID TRsaFtdcUserIDType
	// 交易用户代码
	EndUserID TRsaFtdcUserIDType
}

// 用户
type CRsaFtdcRspUserField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 用户代码
	UserID TRsaFtdcUserIDType
	// 用户登录密码
	Password TRsaFtdcPasswordType
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
	MaxOnline TRsaFtdcNumberType
	// IP地址
	IPAddress TRsaFtdcIPAddressType
	// 是否检查连接产品
	CheckProductInfo TRsaFtdcBoolType
	// Mac地址
	MacAddress TRsaFtdcMacAddressType
	// 是否检查IP和MAC
	CheckIPMacAddr TRsaFtdcBoolType
	// 错误登陆次数限制
	LoginErrorLimit TRsaFtdcNumberType
	// 目前错误登陆次数
	LoginErrorCount TRsaFtdcNumberType
	// 硬盘序列号
	HDSerialID TRsaFtdcHDSerialIDType
	// 密码是否更新
	IsUpdate TRsaFtdcIsUpdateType
}

// 投资者手续费率查询
type CRsaFtdcQryInvestorFeeField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 用户代码
	UserID TRsaFtdcUserIDType
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
	// 交易所代码
	ExchangeID TRsaFtdcExchangeIDType
	// 合约代码
	InstrumentID TRsaFtdcInstrumentIDType
}

// 投资者手续费率
type CRsaFtdcRspInvestorFeeField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 交易所代码
	ExchangeID TRsaFtdcExchangeIDType
	// 投资者编码
	InvestorID TRsaFtdcInvestorIDType
	// 合约代码
	InstrumentID TRsaFtdcInstrumentIDType
	// 开仓手续费按比例
	OpenFeeRate TRsaFtdcRatioType
	// 开仓手续费按手数
	OpenFeeAmt TRsaFtdcRatioType
	// 平仓手续费按比例
	OffsetFeeRate TRsaFtdcRatioType
	// 平仓手续费按手数
	OffsetFeeAmt TRsaFtdcRatioType
	// 平今仓手续费按比例
	OTFeeRate TRsaFtdcRatioType
	// 平今仓手续费按手数
	OTFeeAmt TRsaFtdcRatioType
}

// 投资者保证金率查询
type CRsaFtdcQryInvestorMarginField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 用户代码
	UserID TRsaFtdcUserIDType
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
	// 交易所代码
	ExchangeID TRsaFtdcExchangeIDType
	// 合约代码
	InstrumentID TRsaFtdcInstrumentIDType
	// 投机套保标志
	HedgeFlag TRsaFtdcHedgeFlagType
}

// 投资者保证金率
type CRsaFtdcRspInvestorMarginField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 交易所代码
	ExchangeID TRsaFtdcExchangeIDType
	// 投资者编码
	InvestorID TRsaFtdcInvestorIDType
	// 合约代码
	InstrumentID TRsaFtdcInstrumentIDType
	// 投机套保标志
	HedgeFlag TRsaFtdcHedgeFlagType
	// 多头占用保证金按比例
	LongMarginRate TRsaFtdcRatioType
	// 多头保证金按手数
	LongMarginAmt TRsaFtdcRatioType
	// 空头占用保证金按比例
	ShortMarginRate TRsaFtdcRatioType
	// 空头保证金按手数
	ShortMarginAmt TRsaFtdcRatioType
}

// 投资者查询
type CRsaFtdcQryInvestorField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 用户代码
	UserID TRsaFtdcUserIDType
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
}

// 投资者应答
type CRsaFtdcRspInvestorField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
	// 投资者名称
	InvestorName TRsaFtdcInvestorNameType
	// 投资者类型
	InvestorType TRsaFtdcInvestorTypeType
	// 投资者组编号
	GroupID TRsaFtdcInvestorIDType
	// 投资者组名
	GroupName TRsaFtdcInvestorNameType
	// 别名
	Alias TRsaFtdcAliasType
	// 连接类型
	LinkType TRsaFtdcLinkTypeType
	// 资金账户编号
	ParentInvestorID TRsaFtdcInvestorIDType
	// 登录次数
	LoginStatus TRsaFtdcNumberType
	// 是否使用净持仓报单,投顾账户有效
	PositionType TRsaFtdcPositionTypeType
	// 自成交处理方式,资金账户有效
	SelfTradeAvoidType TRsaFtdcSelfTradeAvoidTypeType
	// 交易权限
	TradingRight TRsaFtdcTradingRightType
	// 是否检查资金
	CheckAccount TRsaFtdcBoolType
	// 是否检查持仓
	CheckPosition TRsaFtdcBoolType
	// 报单是否报警
	OrderWarning TRsaFtdcBoolType
	// 资金持仓是否上查
	QryUpAccPosi TRsaFtdcBoolType
	// 投顾选择资金账户类型
	DistributeType TRsaFtdcDistributeTypeType
	// 跟随资金账户编号
	FollowParentInvestorID TRsaFtdcInvestorIDType
	// 倍率
	FollowRate TRsaFtdcRatioType
}

// 账号风险状况查询
type CRsaFtdcQryAccountRiskField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 用户代码
	UserID TRsaFtdcUserIDType
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
}

// 账号风险状况应答
type CRsaFtdcRspAccountRiskField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
	// 资金帐号
	AccountID TRsaFtdcAccountIDType
	// 设置准备金
	SetBalance TRsaFtdcMoneyType
	// 开仓红线
	OpenAvailable TRsaFtdcMoneyType
	// 平仓红线
	OffsetAvailable TRsaFtdcMoneyType
	// 账户状态
	AccountStatus TRsaFtdcAccountStatusType
	// 动态权益
	DynamicRights TRsaFtdcMoneyType
	// 账户权限
	TradingRight TRsaFtdcTradingRightType
}

// 交易用户会话查询
type CRsaFtdcQryUserSessionField struct {
	// 交易用户代码
	StartUserID TRsaFtdcUserIDType
	// 交易用户代码
	EndUserID TRsaFtdcUserIDType
}

// 交易用户会话应答
type CRsaFtdcRspUserSessionField struct {
	// 经纪公司代码
	BrokerID TRsaFtdcBrokerIDType
	// 交易用户代码
	UserID TRsaFtdcUserIDType
	// 交易用户类型
	UserType TRsaFtdcUserTypeType
	// 会话编号
	SessionID TRsaFtdcSessionIDType
	// 前置编号
	FrontID TRsaFtdcFrontIDType
	// 登录时间
	LoginTime TRsaFtdcTimeType
	// IP地址
	IPAddress TRsaFtdcIPAddressType
	// Mac地址
	MacAddress TRsaFtdcMacAddressType
	// 用户端产品信息
	UserProductInfo TRsaFtdcProductInfoType
	// 接口端产品信息
	InterfaceProductInfo TRsaFtdcProductInfoType
	// 协议信息
	ProtocolInfo TRsaFtdcProtocolInfoType
	// 硬盘序列号
	HDSerialID TRsaFtdcHDSerialIDType
}

// 行情查询
type CRsaFtdcQryMarketDataField struct {
	// 交易所代码
	ExchangeID TRsaFtdcExchangeIDType
	// 合约代码
	InstrumentID TRsaFtdcInstrumentIDType
}

// 行情应答
type CRsaFtdcRspMarketDataField struct {
	// 交易所代码
	ExchangeID TRsaFtdcExchangeIDType
	// 交易日
	TradingDay TRsaFtdcTradingDayType
	// 昨结算
	PreSettlementPrice TRsaFtdcPriceType
	// 昨收盘
	PreClosePrice TRsaFtdcPriceType
	// 昨持仓量
	PreOpenInterest TRsaFtdcLargeVolumeType
	// 昨虚实度
	PreDelta TRsaFtdcRatioType
	// 今开盘
	OpenPrice TRsaFtdcPriceType
	// 最高价
	HighestPrice TRsaFtdcPriceType
	// 最低价
	LowestPrice TRsaFtdcPriceType
	// 今收盘
	ClosePrice TRsaFtdcPriceType
	// 涨停板价
	UpperLimitPrice TRsaFtdcPriceType
	// 跌停板价
	LowerLimitPrice TRsaFtdcPriceType
	// 今结算
	SettlementPrice TRsaFtdcPriceType
	// 今虚实度
	CurrDelta TRsaFtdcRatioType
	// 最新价
	LastPrice TRsaFtdcPriceType
	// 数量
	Volume TRsaFtdcVolumeType
	// 成交金额
	Turnover TRsaFtdcMoneyType
	// 持仓量
	OpenInterest TRsaFtdcLargeVolumeType
	// 申买价一
	BidPrice1 TRsaFtdcPriceType
	// 申买量一
	BidVolume1 TRsaFtdcVolumeType
	// 申卖价一
	AskPrice1 TRsaFtdcPriceType
	// 申卖量一
	AskVolume1 TRsaFtdcVolumeType
	// 合约代码
	InstrumentID TRsaFtdcInstrumentIDType
	// 最后修改时间
	UpdateTime TRsaFtdcTimeType
	// 最后修改毫秒
	UpdateMillisec TRsaFtdcMillisecType
	// 合约状态
	InstrumentStatus TRsaFtdcInstrumentStatusType
}

// 产品查询
type CRsaFtdcQryProductField struct {
	// 交易所代码
	ExchangeID TRsaFtdcExchangeIDType
	// 产品代码
	ProductID TRsaFtdcProductIDType
}

// 产品应答
type CRsaFtdcRspProductField struct {
	// 产品代码
	ProductID TRsaFtdcProductIDType
	// 交易所代码
	ExchangeID TRsaFtdcExchangeIDType
	// 投资者发行的产品名称
	ProductName TRsaFtdcProductNameType
	// 产品类型
	ProductClass TRsaFtdcProductClassType
	// 数量乘数
	VolumeMultiple TRsaFtdcVolumeMultipleType
	// 报价单位
	PriceTick TRsaFtdcPriceTickType
	// 市价单最大下单量
	MaxMarketOrderVolume TRsaFtdcVolumeType
	// 市价单最小下单量
	MinMarketOrderVolume TRsaFtdcVolumeType
	// 限价单最大下单量
	MaxLimitOrderVolume TRsaFtdcVolumeType
	// 限价单最小下单量
	MinLimitOrderVolume TRsaFtdcVolumeType
	// 交易节代码
	SessionName TRsaFtdcSessionNameType
}

// 内部来回消息
type CRsaFtdcPingPongField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 修改用户编号
	UserID TRsaFtdcUserIDType
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
	// 编号
	SequenceNo TRsaFtdcSequenceNoType
	// 用户自定义域
	UserCustom TRsaFtdcCustomType
}

// 最大可交易数量查询请求
type CRsaFtdcQryMaxVolumeField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 用户代码
	UserID TRsaFtdcUserIDType
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
	// 交易所代码
	ExchangeID TRsaFtdcExchangeIDType
	// 合约代码
	InstrumentID TRsaFtdcInstrumentIDType
	// 买卖方向
	Direction TRsaFtdcDirectionType
	// 开平标志
	OffsetFlag TRsaFtdcOffsetFlagType
	// 投机套保标志
	HedgeFlag TRsaFtdcHedgeFlagType
	// 价格
	LimitPrice TRsaFtdcPriceType
}

// 最大可交易数量查询应答
type CRsaFtdcRspMaxVolumeField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 用户代码
	UserID TRsaFtdcUserIDType
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
	// 交易所代码
	ExchangeID TRsaFtdcExchangeIDType
	// 合约代码
	InstrumentID TRsaFtdcInstrumentIDType
	// 买卖方向
	Direction TRsaFtdcDirectionType
	// 开平标志
	OffsetFlag TRsaFtdcOffsetFlagType
	// 投机套保标志
	HedgeFlag TRsaFtdcHedgeFlagType
	// 价格
	LimitPrice TRsaFtdcPriceType
	// 数量
	Volume TRsaFtdcVolumeType
}

// 工作流交易员查询请求
type CRsaFtdcQryAdviceTraderField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 用户代码
	UserID TRsaFtdcUserIDType
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
	// 登录经纪公司编号
	LogBrokerID TRsaFtdcBrokerIDType
	// 登录用户代码
	LogUserID TRsaFtdcUserIDType
	// 用户端接口版本
	ApiVersion TRsaFtdcProductInfoType
	// 用户端程序版本
	UserVersion TRsaFtdcProductInfoType
}

// 工作流交易员查询应答
type CRsaFtdcRspAdviceTraderField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
	// 用户代码
	UserID TRsaFtdcUserIDType
	// 关联用户代码
	RelationUserID TRsaFtdcUserIDType
	// 关联用户名称
	RelationUserName TRsaFtdcUserNameType
}

// 组合持仓明细查询请求
type CRsaFtdcQryCmbPositionDetailField struct {
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
}

// 组合持仓明细查询应答
type CRsaFtdcRspCmbPositionDetailField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 交易所代码
	ExchangeID TRsaFtdcExchangeIDType
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
	// 单腿合约代码
	InstrumentID TRsaFtdcInstrumentIDType
	// 单腿合约代码
	CmbInstrumentID TRsaFtdcInstrumentIDType
	// 成交编号
	TradeID TRsaFtdcTradeIDType
	// 报单编号
	OrderSysID TRsaFtdcOrderSysIDType
	// 买卖方向
	Direction TRsaFtdcDirectionType
	// 投机套保标志
	HedgeFlag TRsaFtdcHedgeFlagType
	// 持仓量
	Position TRsaFtdcVolumeType
	// 单腿编号
	LegQty TRsaFtdcNumberType
	// 交易日期
	TradingDay TRsaFtdcTradingDayType
}

// 投资者结算结果查询请求
type CRsaFtdcQrySettlementInfoField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
	// 交易日期
	TradingDay TRsaFtdcTradingDayType
}

// 投资者结算结果查询应答
type CRsaFtdcRspSettlementInfoField struct {
	// 交易日期
	TradingDay TRsaFtdcTradingDayType
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
	// 序号
	SequenceNo TRsaFtdcSequenceNoType
	// 消息正文
	Content TRsaFtdcContentType
}

// 查询结果集分页
type CRsaFtdcQryResultPageField struct {
	// 结果集编号
	ResultID TRsaFtdcSessionIDType
	// 总记录数
	Count TRsaFtdcNumberType
	// 总页数
	Pages TRsaFtdcNumberType
	// 剩余页数
	Remain TRsaFtdcNumberType
}

// 查询持仓明细请求
type CRsaFtdcQryPositionDetailField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 交易所代码
	ExchangeID TRsaFtdcExchangeIDType
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
	// 单腿合约代码
	InstrumentID TRsaFtdcInstrumentIDType
}

// 查询持仓明细应答
type CRsaFtdcRspPositionDetailField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 交易所代码
	ExchangeID TRsaFtdcExchangeIDType
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
	// 合约代码
	InstrumentID TRsaFtdcInstrumentIDType
	// 投机套保标志
	HedgeFlag TRsaFtdcHedgeFlagType
	// 买卖方向
	Direction TRsaFtdcDirectionType
	// 开仓日期
	OpenDate TRsaFtdcDateType
	// 成交编号
	TradeID TRsaFtdcTradeIDType
	// 数量
	Volume TRsaFtdcVolumeType
	// 开仓价
	OpenPrice TRsaFtdcMoneyType
	// 交易日
	TradingDay TRsaFtdcDateType
	// 交易类型
	TradeType TRsaFtdcTradeTypeType
	// 组合合约代码
	CombInstrumentID TRsaFtdcInstrumentIDType
	// 逐日盯市平仓盈亏
	CloseProfitByDate TRsaFtdcMoneyType
	// 逐笔对冲平仓盈亏
	CloseProfitByTrade TRsaFtdcMoneyType
	// 逐日盯市持仓盈亏
	PositionProfitByDate TRsaFtdcMoneyType
	// 逐笔对冲持仓盈亏
	PositionProfitByTrade TRsaFtdcMoneyType
	// 投资者保证金
	Margin TRsaFtdcMoneyType
	// 交易所保证金
	ExchMargin TRsaFtdcMoneyType
	// 保证金率
	MarginRateByMoney TRsaFtdcMoneyType
	// 保证金率(按手数)
	MarginRateByVolume TRsaFtdcMoneyType
	// 昨结算价
	LastSettlementPrice TRsaFtdcPriceType
	// 结算价
	SettlementPrice TRsaFtdcPriceType
	// 平仓量
	CloseVolume TRsaFtdcVolumeType
	// 平仓金额
	CloseAmount TRsaFtdcMoneyType
	// 结算编号
	SettlementID TRsaFtdcSettlementIDType
	// 投资单元代码
	InvestUnitID TRsaFtdcInvestUnitIDType
	// 持仓明细编号
	SeqNum TRsaFtdcNumberType
	// 用户自定义域
	UserCustom TRsaFtdcCustomType
	// 持仓明细冻结数量
	FrozenVolume TRsaFtdcVolumeType
	// 数量(小数)
	VolumeP TRsaFtdcPointVolumeType
	// 持仓明细冻结数量(小数)
	FrozenVolumeP TRsaFtdcPointVolumeType
	// 平仓量(小数)
	CloseVolumeP TRsaFtdcPointVolumeType
}

// 询价查询
type CRsaFtdcQryForQuoteField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
	// 合约代码
	InstrumentID TRsaFtdcInstrumentIDType
	// 交易所代码
	ExchangeID TRsaFtdcExchangeIDType
}

// 询价查询应答
type CRsaFtdcForQuoteField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
	// 用户代码
	UserID TRsaFtdcUserIDType
	// 用户本地报单号
	UserOrderLocalID TRsaFtdcUserOrderLocalIDType
	// 用户自定义域
	UserCustom TRsaFtdcCustomType
	// qsc全局编号
	QscId TRsaFtdcNumberType
	// 插件标识
	PluginTaskId TRsaFtdcNumberType
	// 外部标识
	OthMarkId TRsaFtdcOthMarkIdType
	// 交易所代码
	ExchangeID TRsaFtdcExchangeIDType
	// 合约代码
	InstrumentID TRsaFtdcInstrumentIDType
	// 本地报单编号
	OrderLocalID TRsaFtdcOrderLocalIDType
	// 询价状态
	ForQuoteStatus TRsaFtdcForQuoteStatusTypeType
	// 交易日
	TradingDay TRsaFtdcDateType
	// 插入时间
	InsertTime TRsaFtdcTimeType
}

// 报价查询
type CRsaFtdcQryQuoteField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 用户代码
	UserID TRsaFtdcUserIDType
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
	// 合约代码
	InstrumentID TRsaFtdcInstrumentIDType
	// 交易所代码
	ExchangeID TRsaFtdcExchangeIDType
	// 系统报单编号
	OrderSysID TRsaFtdcOrderSysIDType
	// 开始时间
	InsertTimeStart TRsaFtdcTimeType
	// 结束时间
	InsertTimeEnd TRsaFtdcTimeType
}

// 投资者期权手续费率查询
type CRsaFtdcQryInvestorOptionFeeField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 用户代码
	UserID TRsaFtdcUserIDType
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
	// 交易所代码
	ExchangeID TRsaFtdcExchangeIDType
	// 合约代码
	InstrumentID TRsaFtdcInstrumentIDType
}

// 投资者期权手续费率
type CRsaFtdcRspInvestorOptionFeeField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 交易所代码
	ExchangeID TRsaFtdcExchangeIDType
	// 投资者编码
	InvestorID TRsaFtdcInvestorIDType
	// 合约代码
	InstrumentID TRsaFtdcInstrumentIDType
	// 开仓手续费按比例
	OpenFeeRate TRsaFtdcRatioType
	// 开仓手续费按手数
	OpenFeeAmt TRsaFtdcRatioType
	// 平仓手续费按比例
	OffsetFeeRate TRsaFtdcRatioType
	// 平仓手续费按手数
	OffsetFeeAmt TRsaFtdcRatioType
	// 平今仓手续费按比例
	OTFeeRate TRsaFtdcRatioType
	// 平今仓手续费按手数
	OTFeeAmt TRsaFtdcRatioType
	// 执行手续费按比例
	StrikeFeeRate TRsaFtdcRatioType
	// 执行手续费按手数
	StrikeFeeAmt TRsaFtdcRatioType
}

// 执行宣告查询
type CRsaFtdcQryExecOrderField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 用户代码
	UserID TRsaFtdcUserIDType
	// 交易所代码
	ExchangeID TRsaFtdcExchangeIDType
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
	// 报单编号
	OrderSysID TRsaFtdcOrderSysIDType
	// 合约代码
	InstrumentID TRsaFtdcInstrumentIDType
	// 开始时间
	InsertTimeStart TRsaFtdcTimeType
	// 结束时间
	InsertTimeEnd TRsaFtdcTimeType
}

// 投资者通知
type CRsaFtdcRtnInvestorField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
	// 投资者名称
	InvestorName TRsaFtdcInvestorNameType
	// 投资者类型
	InvestorType TRsaFtdcInvestorTypeType
	// 投资者组编号
	GroupID TRsaFtdcInvestorIDType
	// 投资者组名
	GroupName TRsaFtdcInvestorNameType
	// 别名
	Alias TRsaFtdcAliasType
	// 连接类型
	LinkType TRsaFtdcLinkTypeType
	// 资金账户编号
	ParentInvestorID TRsaFtdcInvestorIDType
	// 登录次数
	LoginStatus TRsaFtdcNumberType
	// 是否使用净持仓报单,投顾账户有效
	PositionType TRsaFtdcPositionTypeType
	// 自成交处理方式,资金账户有效
	SelfTradeAvoidType TRsaFtdcSelfTradeAvoidTypeType
	// 交易权限
	TradingRight TRsaFtdcTradingRightType
	// 是否检查资金
	CheckAccount TRsaFtdcBoolType
	// 是否检查持仓
	CheckPosition TRsaFtdcBoolType
	// 报单是否报警
	OrderWarning TRsaFtdcBoolType
	// 资金持仓是否上查
	QryUpAccPosi TRsaFtdcBoolType
	// 投顾选择资金账户类型
	DistributeType TRsaFtdcDistributeTypeType
	// 跟随资金账户编号
	FollowParentInvestorID TRsaFtdcInvestorIDType
	// 倍率
	FollowRate TRsaFtdcRatioType
}

// 投资者资金通知
type CRsaFtdcRtnInvestorAccountField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
	// 资金帐号
	AccountID TRsaFtdcAccountIDType
	// 上次结算准备金
	PreBalance TRsaFtdcMoneyType
	// 上日可用资金
	PreAvailable TRsaFtdcMoneyType
	// 入金金额
	Deposit TRsaFtdcMoneyType
	// 出金金额
	Withdraw TRsaFtdcMoneyType
	// 占用保证金
	Margin TRsaFtdcMoneyType
	// 期权权利金收支
	Premium TRsaFtdcMoneyType
	// 手续费
	Fee TRsaFtdcMoneyType
	// 冻结的保证金
	FrozenMargin TRsaFtdcMoneyType
	// 冻结权利金
	FrozenPremium TRsaFtdcMoneyType
	// 冻结手续费
	FrozenFee TRsaFtdcMoneyType
	// 平仓盈亏
	CloseProfit TRsaFtdcMoneyType
	// 持仓盈亏
	PositionProfit TRsaFtdcMoneyType
	// 可用资金
	Available TRsaFtdcMoneyType
	// 结算准备金
	Balance TRsaFtdcMoneyType
	// 多头占用保证金
	LongMargin TRsaFtdcMoneyType
	// 空头占用保证金
	ShortMargin TRsaFtdcMoneyType
	// 多头冻结的保证金
	LongFrozenMargin TRsaFtdcMoneyType
	// 空头冻结的保证金
	ShortFrozenMargin TRsaFtdcMoneyType
	// 动态权益
	DynamicRights TRsaFtdcMoneyType
	// 风险度
	Risk TRsaFtdcRatioType
	// 其他费用
	OtherFee TRsaFtdcMoneyType
	// 质押金额
	Mortgage TRsaFtdcMoneyType
	// 币种
	Currency TRsaFtdcCurrencyType
	// 可取资金
	WithdrawQuota TRsaFtdcMoneyType
	// 冻结的资金
	FrozenCash TRsaFtdcMoneyType
	// 账户余额
	CurrentBalance TRsaFtdcMoneyType
	// 总资产
	AssetBalance TRsaFtdcMoneyType
	// 总市值
	MarketValue TRsaFtdcMoneyType
	// 优先资金
	Prioprity TRsaFtdcMoneyType
	// 现金资产
	FundAsset TRsaFtdcMoneyType
	// 上海限购额度
	PurQuotaSh TRsaFtdcMoneyType
	// 上海可用限购额度
	EnPurQuotaSh TRsaFtdcMoneyType
	// 深圳限购额度
	PurQuotaSz TRsaFtdcMoneyType
	// 深圳可用限购额度
	EnPurQuotaSz TRsaFtdcMoneyType
}

// 投资者风控权限通知
type CRsaFtdcRtnInvestorTradingRightRiskField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
	// 交易所代码
	ExchangeID TRsaFtdcExchangeIDType
	// 合约代码
	InstrumentID TRsaFtdcInstrumentIDType
	// 买开权限
	BuyOpen TRsaFtdcBoolType
	// 买平权限
	BuyClose TRsaFtdcBoolType
	// 卖开权限
	SellOpen TRsaFtdcBoolType
	// 卖平权限
	SellClose TRsaFtdcBoolType
	// 产品类型
	ProductClass TRsaFtdcProductClassType
}

// 投资者权限通知
type CRsaFtdcRtnInvestorTradingRightField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
	// 交易所代码
	ExchangeID TRsaFtdcExchangeIDType
	// 合约代码
	InstrumentID TRsaFtdcInstrumentIDType
	// 合约交易权限
	TradingRight TRsaFtdcTradingRightType
	// 产品类型
	ProductClass TRsaFtdcProductClassType
}

// 交易汇率查询
type CRsaFtdcQryExchangeRateField struct {
	// 币种
	InstCurrency TRsaFtdcCurrencyType
	// 币种
	AccCurrency TRsaFtdcCurrencyType
}

// 交易汇率
type CRsaFtdcRspExchangeRateField struct {
	// 币种
	InstCurrency TRsaFtdcCurrencyType
	// 币种
	AccCurrency TRsaFtdcCurrencyType
	// 汇率
	ExchangeRate TRsaFtdcRatioType
}

// 交易汇率通知
type CRsaFtdcRtnExchangeRateField struct {
	// 币种
	InstCurrency TRsaFtdcCurrencyType
	// 币种
	AccCurrency TRsaFtdcCurrencyType
	// 汇率
	ExchangeRate TRsaFtdcRatioType
}

// 新组合持仓明细查询请求
type CRsaFtdcQryCombPositionDetailField struct {
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
}

// 新组合持仓明细查询应答
type CRsaFtdcRspCombPositionDetailField struct {
	// 交易日期
	TradingDay TRsaFtdcTradingDayType
	// 开仓日期
	OpenDate TRsaFtdcDateType
	// 交易所代码
	ExchangeID TRsaFtdcExchangeIDType
	// 结算编号
	SettlementID TRsaFtdcSettlementIDType
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
	// 组合编号
	ComTradeID TRsaFtdcTradeIDType
	// 成交编号
	TradeID TRsaFtdcTradeIDType
	// 合约代码
	InstrumentID TRsaFtdcInstrumentIDType
	// 投机套保标志
	HedgeFlag TRsaFtdcHedgeFlagType
	// 买卖方向
	Direction TRsaFtdcDirectionType
	// 持仓量
	Position TRsaFtdcVolumeType
	// 投资者保证金
	Margin TRsaFtdcMoneyType
	// 交易所保证金
	ExchMargin TRsaFtdcMoneyType
	// 保证金率
	MarginRateByMoney TRsaFtdcMoneyType
	// 保证金率(按手数)
	MarginRateByVolume TRsaFtdcMoneyType
	// 单腿编号
	LegID TRsaFtdcLegIDType
	// 单腿乘数
	LegMultiple TRsaFtdcLegMultipleType
	// 组合持仓合约代码
	CmbInstrumentID TRsaFtdcInstrumentIDType
	// 成交组号
	TradeGroupID TRsaFtdcTradeGroupIDType
}

// 投资者持仓通知
type CRsaFtdcRtnInvestorPositionField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 交易所代码
	ExchangeID TRsaFtdcExchangeIDType
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
	// 合约代码
	InstrumentID TRsaFtdcInstrumentIDType
	// 买卖方向
	Direction TRsaFtdcDirectionType
	// 投机套保标志
	HedgeFlag TRsaFtdcHedgeFlagType
	// 占用保证金
	UsedMargin TRsaFtdcMoneyType
	// 总持仓量
	Position TRsaFtdcVolumeType
	// 总持仓成本
	PositionCost TRsaFtdcPriceType
	// 初始昨持仓量(当日不变)
	YdPosition TRsaFtdcVolumeType
	// 初始昨日持仓成本(当日不变)
	YdPositionCost TRsaFtdcMoneyType
	// 冻结的保证金
	FrozenMargin TRsaFtdcMoneyType
	// 开仓冻结持仓
	FrozenPosition TRsaFtdcVolumeType
	// 平仓冻结持仓
	FrozenClosing TRsaFtdcVolumeType
	// 持仓盈亏
	PositionProfit TRsaFtdcMoneyType
	// 冻结的权利金
	FrozenPremium TRsaFtdcMoneyType
	// 最后一笔成交编号
	LastTradeID TRsaFtdcTradeIDType
	// 最后一笔本地报单编号
	LastOrderLocalID TRsaFtdcOrderLocalIDType
	// 总持仓可平仓数量(包括平仓冻结持仓)
	PositionClose TRsaFtdcVolumeType
	// 昨持仓可平仓数量(包括平仓冻结持仓)
	YdPositionClose TRsaFtdcVolumeType
	// 昨持仓平仓冻结持仓
	YdFrozenClosing TRsaFtdcVolumeType
	// 今日开仓数量(不包括冻结)
	OpenVolume TRsaFtdcVolumeType
	// 今日平仓数量(包括昨持仓的平仓,不包括冻结)
	CloseVolume TRsaFtdcVolumeType
	// 平仓盈亏
	CloseProfit TRsaFtdcMoneyType
	// 执行冻结持仓
	StrikeFrozenPosition TRsaFtdcVolumeType
	// 放弃执行冻结持仓
	AbandonFrozenPosition TRsaFtdcVolumeType
	// 成本价
	CostPrice TRsaFtdcPriceType
	// 开仓成本
	OpenCost TRsaFtdcMoneyType
	// 今持仓(包含冻结)
	TodayPosition TRsaFtdcVolumeType
	// 今持仓冻结持仓
	TodayFrozenClosing TRsaFtdcVolumeType
	// 总持仓量(小数)
	PositionP TRsaFtdcPointVolumeType
}

// 期权自对冲查询
type CRsaFtdcQryOptionSelfCloseField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 投资者编码
	InvestorID TRsaFtdcInvestorIDType
	// 合约代码
	InstrumentID TRsaFtdcInstrumentIDType
	// 交易所代码
	ExchangeID TRsaFtdcExchangeIDType
	// 期权自对冲操作编号
	OptionSelfCloseSysID TRsaFtdcOrderSysIDType
	// 开始时间
	InsertTimeStart TRsaFtdcTimeType
	// 结束时间
	InsertTimeEnd TRsaFtdcTimeType
}

// 锁定证券仓位查询
type CRsaFtdcQryLockPositionField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
	// 合约代码
	InstrumentID TRsaFtdcInstrumentIDType
	// 交易所代码
	ExchangeID TRsaFtdcExchangeIDType
}

// 锁定证券仓位
type CRsaFtdcLockPositionField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
	// 会员编号
	ParticipantID TRsaFtdcParticipantIDType
	// 客户代码
	ClientID TRsaFtdcClientIDType
	// 合约代码
	InstrumentID TRsaFtdcInstrumentIDType
	// 交易所代码
	ExchangeID TRsaFtdcExchangeIDType
	// 锁仓数量
	Volume TRsaFtdcVolumeType
	// 冻结数量
	FrozenVolume TRsaFtdcVolumeType
	// 今日锁定数量
	TodayVolume TRsaFtdcVolumeType
}

// 投资者其他费率查询
type CRsaFtdcQryInvestorOtherFeeField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 用户代码
	UserID TRsaFtdcUserIDType
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
	// 交易所代码
	ExchangeID TRsaFtdcExchangeIDType
	// 合约代码
	InstrumentID TRsaFtdcInstrumentIDType
	// 费率类型
	FeeType TRsaFtdcFeeTypeType
}

// 投资者其他费率
type CRsaFtdcRspInvestorOtherFeeField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 交易所代码
	ExchangeID TRsaFtdcExchangeIDType
	// 投资者编码
	InvestorID TRsaFtdcInvestorIDType
	// 合约代码
	InstrumentID TRsaFtdcInstrumentIDType
	// 开仓手续费按比例
	OpenFeeRate TRsaFtdcRatioType
	// 开仓手续费按手数
	OpenFeeAmt TRsaFtdcRatioType
	// 平仓手续费按比例
	OffsetFeeRate TRsaFtdcRatioType
	// 平仓手续费按手数
	OffsetFeeAmt TRsaFtdcRatioType
	// 平今仓手续费按比例
	OTFeeRate TRsaFtdcRatioType
	// 平今仓手续费按手数
	OTFeeAmt TRsaFtdcRatioType
	// 费率类型
	FeeType TRsaFtdcFeeTypeType
}

// 用户对应的交易员查询请求
type CRsaFtdcQryUserTraderField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 用户代码
	UserID TRsaFtdcUserIDType
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
}

// 用户对应的交易员查询应答
type CRsaFtdcRspUserTraderField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 用户代码
	UserID TRsaFtdcUserIDType
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
	// 父投资者编号
	ParentInvestorID TRsaFtdcInvestorIDType
	// 是否默认
	IsDefault TRsaFtdcBoolType
}

// 合约组
type CRsaFtdcInstrumentGroupField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 交易所代码
	ExchangeID TRsaFtdcExchangeIDType
	// 合约代码
	InstrumentID TRsaFtdcInstrumentIDType
	// 合约组代码
	InstrumentGroupID TRsaFtdcInstrumentGroupIDType
}

// 投资者组合保证金类型
type CRsaFtdcInvestorMarginCombTypeField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
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
	TradingDay TRsaFtdcDateType
}

// 交易节查询应答
type CRsaFtdcRspSessionInfoField struct {
	// 交易节Json数据
	SessionJsonData TRsaFtdcLongJsonDataType
}

// 插件交互数据
type CRsaFtdcQscPluginDataField struct {
	// qsc全局编号
	QscId TRsaFtdcNumberType
	// 插件标识
	PluginTaskId TRsaFtdcNumberType
	// 外部标识
	OthMarkId TRsaFtdcOthMarkIdType
	// 插件字符串标识
	PluginId TRsaFtdcPluginIdType
	// 指令类型
	PluginOrderType TRsaFtdcPluginOrderTypeType
	// 指令编号
	FunctionId TRsaFtdcFunctionIdType
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 用户代码
	UserID TRsaFtdcUserIDType
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
	// 会话编号
	SessionID TRsaFtdcSessionIDType
	// 数据体
	Data TRsaFtdcJsonDataType
}

// 插件服务创建请求
type CRsaFtdcQscRegisterPluginServiceField struct {
	// qsc全局编号
	QscId TRsaFtdcNumberType
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
	QscId TRsaFtdcNumberType
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
	QscId TRsaFtdcNumberType
	// 插件标识
	PluginId TRsaFtdcPluginIdType
	// 服务json描述
	ServiceDesc TRsaFtdcJsonDataType
}

// 插件任务查询请求
type CRsaFtdcQryPluginTaskField struct {
	// qsc全局编号
	QscId TRsaFtdcNumberType
	// 用户代码
	UserID TRsaFtdcUserIDType
}

// 插件任务
type CRsaFtdcPluginTaskField struct {
	// 插件字符串标识
	PluginId TRsaFtdcPluginIdType
	// qsc全局编号
	QscId TRsaFtdcNumberType
	// 插件标识
	PluginTaskId TRsaFtdcNumberType
	// 外部标识
	OthMarkId TRsaFtdcOthMarkIdType
	// 执行功能编号
	ActionId TRsaFtdcFunctionIdType
	// 用户代码
	UserID TRsaFtdcUserIDType
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
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
	QscId TRsaFtdcNumberType
	// 插件标识
	PluginTaskId TRsaFtdcNumberType
	// 外部标识
	OthMarkId TRsaFtdcOthMarkIdType
}

// 插件子任务状态
type CRsaFtdcPluginTaskSubStatusField struct {
	// 插件字符串标识
	PluginId TRsaFtdcPluginIdType
	// qsc全局编号
	QscId TRsaFtdcNumberType
	// 插件标识
	PluginTaskId TRsaFtdcNumberType
	// 外部标识
	OthMarkId TRsaFtdcOthMarkIdType
	// 通知功能编号
	NotifyId TRsaFtdcFunctionIdType
	// 数据关键字
	DataKey TRsaFtdcDataKeyType
	// 用户代码
	UserID TRsaFtdcUserIDType
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
	// 数据体
	Data TRsaFtdcJsonDataType
}

// 持仓
type CRsaFtdcInvestorPositionField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 交易所代码
	ExchangeID TRsaFtdcExchangeIDType
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
	// 合约代码
	InstrumentID TRsaFtdcInstrumentIDType
	// 买卖方向
	Direction TRsaFtdcDirectionType
	// 投机套保标志
	HedgeFlag TRsaFtdcHedgeFlagType
	// 占用保证金
	UsedMargin TRsaFtdcMoneyType
	// 总持仓量
	Position TRsaFtdcVolumeType
	// 总持仓成本
	PositionCost TRsaFtdcPriceType
	// 初始昨持仓量(当日不变)
	YdPosition TRsaFtdcVolumeType
	// 初始昨日持仓成本(当日不变)
	YdPositionCost TRsaFtdcMoneyType
	// 冻结的保证金
	FrozenMargin TRsaFtdcMoneyType
	// 开仓冻结持仓
	FrozenPosition TRsaFtdcVolumeType
	// 平仓冻结持仓
	FrozenClosing TRsaFtdcVolumeType
	// 持仓盈亏
	PositionProfit TRsaFtdcMoneyType
	// 冻结的权利金
	FrozenPremium TRsaFtdcMoneyType
	// 最后一笔成交编号
	LastTradeID TRsaFtdcTradeIDType
	// 最后一笔本地报单编号
	LastOrderLocalID TRsaFtdcOrderLocalIDType
	// 总持仓可平仓数量(包括平仓冻结持仓)
	PositionClose TRsaFtdcVolumeType
	// 昨持仓可平仓数量(包括平仓冻结持仓)
	YdPositionClose TRsaFtdcVolumeType
	// 昨持仓平仓冻结持仓
	YdFrozenClosing TRsaFtdcVolumeType
	// 今日开仓数量(不包括冻结)
	OpenVolume TRsaFtdcVolumeType
	// 今日平仓数量(包括昨持仓的平仓,不包括冻结)
	CloseVolume TRsaFtdcVolumeType
	// 平仓盈亏
	CloseProfit TRsaFtdcMoneyType
	// 执行冻结持仓
	StrikeFrozenPosition TRsaFtdcVolumeType
	// 放弃执行冻结持仓
	AbandonFrozenPosition TRsaFtdcVolumeType
	// 成本价
	CostPrice TRsaFtdcPriceType
	// 开仓成本
	OpenCost TRsaFtdcMoneyType
	// 今持仓(包含冻结)
	TodayPosition TRsaFtdcVolumeType
	// 今持仓冻结持仓
	TodayFrozenClosing TRsaFtdcVolumeType
	// 总持仓量(小数)
	PositionP TRsaFtdcPointVolumeType
	// 资金帐号
	AccountID TRsaFtdcAccountIDType
	// 会员编号
	ParticipantID TRsaFtdcParticipantIDType
	// 客户代码
	ClientID TRsaFtdcClientIDType
	// 计算浮动盈亏最新价
	LastPrice TRsaFtdcPriceType
	// 产品代码
	ProductID TRsaFtdcProductIDType
	// 期权类型
	OptionType TRsaFtdcOptionsTypeType
	// 行权价
	StrikePrice TRsaFtdcPriceType
	// 期权对应的标的合约代码
	OptionInstrumentID TRsaFtdcInstrumentIDType
	// 短线可平仓量
	ShortClose TRsaFtdcVolumeType
}

// 资金
type CRsaFtdcInvestorAccountField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
	// 资金帐号
	AccountID TRsaFtdcAccountIDType
	// 上次结算准备金
	PreBalance TRsaFtdcMoneyType
	// 上日可用资金
	PreAvailable TRsaFtdcMoneyType
	// 入金金额
	Deposit TRsaFtdcMoneyType
	// 出金金额
	Withdraw TRsaFtdcMoneyType
	// 占用保证金
	Margin TRsaFtdcMoneyType
	// 期权权利金收支
	Premium TRsaFtdcMoneyType
	// 手续费
	Fee TRsaFtdcMoneyType
	// 冻结的保证金
	FrozenMargin TRsaFtdcMoneyType
	// 冻结权利金
	FrozenPremium TRsaFtdcMoneyType
	// 冻结手续费
	FrozenFee TRsaFtdcMoneyType
	// 平仓盈亏
	CloseProfit TRsaFtdcMoneyType
	// 持仓盈亏
	PositionProfit TRsaFtdcMoneyType
	// 可用资金
	Available TRsaFtdcMoneyType
	// 结算准备金
	Balance TRsaFtdcMoneyType
	// 多头占用保证金
	LongMargin TRsaFtdcMoneyType
	// 空头占用保证金
	ShortMargin TRsaFtdcMoneyType
	// 多头冻结的保证金
	LongFrozenMargin TRsaFtdcMoneyType
	// 空头冻结的保证金
	ShortFrozenMargin TRsaFtdcMoneyType
	// 动态权益
	DynamicRights TRsaFtdcMoneyType
	// 风险度
	Risk TRsaFtdcRatioType
	// 其他费用
	OtherFee TRsaFtdcMoneyType
	// 质押金额
	Mortgage TRsaFtdcMoneyType
	// 币种
	Currency TRsaFtdcCurrencyType
	// 可取资金
	WithdrawQuota TRsaFtdcMoneyType
	// 冻结的资金
	FrozenCash TRsaFtdcMoneyType
	// 账户余额
	CurrentBalance TRsaFtdcMoneyType
	// 总资产
	AssetBalance TRsaFtdcMoneyType
	// 总市值
	MarketValue TRsaFtdcMoneyType
	// 优先资金
	Prioprity TRsaFtdcMoneyType
	// 现金资产
	FundAsset TRsaFtdcMoneyType
	// 上海限购额度
	PurQuotaSh TRsaFtdcMoneyType
	// 上海可用限购额度
	EnPurQuotaSh TRsaFtdcMoneyType
	// 深圳限购额度
	PurQuotaSz TRsaFtdcMoneyType
	// 深圳可用限购额度
	EnPurQuotaSz TRsaFtdcMoneyType
	// 总冻结持仓
	TotalFrozenPos TRsaFtdcVolumeType
	// 今日出入金
	TodayInOut TRsaFtdcMoneyType
	// 当日释放保证金
	ReleaseMargin TRsaFtdcMoneyType
	// 浮盈是否可用
	IsProfitCanUse TRsaFtdcBoolType
	// 临时冻结资金
	ManualFrozen TRsaFtdcMoneyType
	// 期权权利金收入
	PremiumIn TRsaFtdcMoneyType
	// 期权权利金支出
	PremiumOut TRsaFtdcMoneyType
	// 现货买入冻结资金
	SpotFrozen TRsaFtdcMoneyType
	// 现货买入成交资金冻结
	SpotTradeFrozen TRsaFtdcMoneyType
	// 现货卖出成交回笼资金
	SpotTradeRelease TRsaFtdcMoneyType
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
	QscId TRsaFtdcNumberType
	// 插件标识
	PluginTaskId TRsaFtdcNumberType
	// 外部标识
	OthMarkId TRsaFtdcOthMarkIdType
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
	InvestorID TRsaFtdcInvestorIDType
	// 交易所代码
	ExchangeID TRsaFtdcExchangeIDType
}

// 账户新股申购额度
type CRsaFtdcInvestorAccountIPOField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
	// 交易所代码
	ExchangeID TRsaFtdcExchangeIDType
	// 资金帐号
	AccountID TRsaFtdcAccountIDType
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
	ExchangeID TRsaFtdcExchangeIDType
	// 合约代码
	InstrumentID TRsaFtdcInstrumentIDType
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 用户代码
	UserID TRsaFtdcUserIDType
	// 系统询价编号
	QuoteID TRsaFtdcQuoteIDType
	// 用户本地报单号
	UserOrderLocalID TRsaFtdcUserOrderLocalIDType
	// 询价请求时间
	ReqTime TRsaFtdcTimeType
	// 用户自定义域
	UserCustom TRsaFtdcCustomType
}

// 客户询价请求应答
type CRsaFtdcRspClientReqQuoteField struct {
	// 交易所代码
	ExchangeID TRsaFtdcExchangeIDType
	// 合约代码
	InstrumentID TRsaFtdcInstrumentIDType
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 用户代码
	UserID TRsaFtdcUserIDType
	// 系统询价编号
	QuoteID TRsaFtdcQuoteIDType
	// 用户本地报单号
	UserOrderLocalID TRsaFtdcUserOrderLocalIDType
	// 询价请求时间
	ReqTime TRsaFtdcTimeType
	// 用户自定义域
	UserCustom TRsaFtdcCustomType
	// 价格
	LimitPrice TRsaFtdcPriceType
	// 库存
	Position TRsaFtdcVolumeType
}

// 客户询价请求查询请求
type CRsaFtdcReqQryClientReqQuoteField struct {
	// 交易所代码
	ExchangeID TRsaFtdcExchangeIDType
	// 合约代码
	InstrumentID TRsaFtdcInstrumentIDType
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
	// 系统询价编号
	QuoteID TRsaFtdcQuoteIDType
}

// 客户询价请求应答
type CRsaFtdcRspQryClientReqQuoteField struct {
	// 交易所代码
	ExchangeID TRsaFtdcExchangeIDType
	// 合约代码
	InstrumentID TRsaFtdcInstrumentIDType
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 用户代码
	UserID TRsaFtdcUserIDType
	// 系统询价编号
	QuoteID TRsaFtdcQuoteIDType
	// 用户本地报单号
	UserOrderLocalID TRsaFtdcUserOrderLocalIDType
	// 询价请求时间
	ReqTime TRsaFtdcTimeType
	// 用户自定义域
	UserCustom TRsaFtdcCustomType
	// 价格
	LimitPrice TRsaFtdcPriceType
	// 库存
	Position TRsaFtdcVolumeType
	// 询价应答时间
	RspTime TRsaFtdcTimeType
	// 客户询价状态
	ClientReqQuoteStatus TRsaFtdcClientReqQuoteStatusType
	// 错误代码
	ErrorID TRsaFtdcErrorIDType
	// 错误信息
	ErrorMsg TRsaFtdcErrorMsgType
}

// 发送通用消息请求
type CRsaFtdcComMessageField struct {
	// 用户代码
	UserID TRsaFtdcUserIDType
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
	// 消息类型
	MessageType TRsaFtdcMessageTypeType
	// 信息内容
	MessageContent TRsaFtdcMessageType
}

// 共享内存登录消息
type CRsaFtdcReqShmLoginField struct {
	// 交易用户代码
	UserID TRsaFtdcUserIDType
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 密码
	Password TRsaFtdcPasswordType
}

// 切日通知
type CRsaFtdcUpdTradingDayField struct {
	// 当前交易日（切日前）
	CurrentTradingDay TRsaFtdcTradingDayType
	// 下一交易日（切日后）
	NextTradingDay TRsaFtdcTradingDayType
	// 更新时间
	UpdTime TRsaFtdcTimeType
	// 是否为交易日结束
	IsTradingDayEnd TRsaFtdcBoolType
}

// 行情推送结束通知
type CRsaFtdcNtfMarketDataEndField struct {
	// 行情开始时间
	StartTradingDay TRsaFtdcTradingDayType
	// 行情结束时间
	EndTradingDay TRsaFtdcTradingDayType
	// 时间
	Time TRsaFtdcTimeType
}

// 回测行情订阅
type CRsaFtdcBtSubMarketDataField struct {
	// 交易所代码
	ExchangeID TRsaFtdcExchangeIDType
	// 合约代码
	InstrumentID TRsaFtdcInstrumentIDType
	// Bar精度
	BarPreces TRsaFtdcBarPrecesType
	// Bar周期
	BarPeriod TRsaFtdcNumberType
}

// 主力合约信息查询请求
type CRsaFtdcQryHotInstrumentField struct {
	// 交易日
	TradingDay TRsaFtdcTradingDayType
	// 交易所代码
	ExchangeID TRsaFtdcExchangeIDType
	// 基础品种代码
	UnderlyingID TRsaFtdcInstrumentIDType
}

// 主力合约信息查询应答
type CRsaFtdcRspHotInstrumentField struct {
	// 交易日
	TradingDay TRsaFtdcTradingDayType
	// 交易所代码
	ExchangeID TRsaFtdcExchangeIDType
	// 基础品种代码
	UnderlyingID TRsaFtdcInstrumentIDType
	// 合约代码
	InstrumentID TRsaFtdcInstrumentIDType
	// 产品类型
	ProductClass TRsaFtdcProductClassType
	// 成交量
	Volume TRsaFtdcVolumeType
	// 品种最大成交量
	MaxVolume TRsaFtdcVolumeType
	// 成交金额
	Turnover TRsaFtdcMoneyType
	// 品种最大成交金额
	MaxTurnover TRsaFtdcMoneyType
	// 持仓量
	OpenInterest TRsaFtdcLargeVolumeType
	// 品种最大持仓量
	MaxOpenInterest TRsaFtdcLargeVolumeType
	// 主力合约等级
	Rank TRsaFtdcVolumeType
}

// Cta委托
type CRsaFtdcCtaOrderField struct {
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
	// 交易所代码
	ExchangeID TRsaFtdcExchangeIDType
	// 合约代码
	InstrumentID TRsaFtdcInstrumentIDType
	// 数量
	Volume TRsaFtdcVolumeType
	// 报单类型
	OrderPriceType TRsaFtdcOrderPriceTypeType
	// 开平方式
	OpenCloseMethod TRsaFtdcOpenCloseMethodType
	// 价格
	LimitPrice TRsaFtdcPriceType
	// 业务单元
	BusinessUnit TRsaFtdcBusinessUnitType
	// 业务类别
	BusinessType TRsaFtdcBusinessTypeType
	// 策略限制价格
	AlgoLimitPrice TRsaFtdcPriceType
	// 策略限制tick数量
	AlgoLimitPriceTick TRsaFtdcNumberType
	// 策略限制时间(ms)
	AlgoLimitTime TRsaFtdcNumberType
}

// 订阅组合行情
type CRsaFtdcSubCombMarketDataField struct {
	// 组合合约代码
	CombInstID TRsaFtdcInstrumentIDType
	// 公式
	Formula TRsaFtdcFormulaType
	// 小数点位数
	Precision TRsaFtdcVolumeType
	// 腿数
	LegNum TRsaFtdcVolumeType
	// Leg1交易所代码
	Leg1ExchangeID TRsaFtdcExchangeIDType
	// Leg1合约代码
	Leg1InstrumentID TRsaFtdcInstrumentIDType
	// Leg1买卖方向
	Leg1Direction TRsaFtdcDirectionType
	// Leg2易所代码
	Leg2ExchangeID TRsaFtdcExchangeIDType
	// Leg2合约代码
	Leg2InstrumentID TRsaFtdcInstrumentIDType
	// Leg2买卖方向
	Leg2Direction TRsaFtdcDirectionType
	// Leg3交易所代码
	Leg3ExchangeID TRsaFtdcExchangeIDType
	// Leg3合约代码
	Leg3InstrumentID TRsaFtdcInstrumentIDType
	// Leg3买卖方向
	Leg3Direction TRsaFtdcDirectionType
	// Leg4交易所代码
	Leg4ExchangeID TRsaFtdcExchangeIDType
	// Leg4合约代码
	Leg4InstrumentID TRsaFtdcInstrumentIDType
	// Leg4买卖方向
	Leg4Direction TRsaFtdcDirectionType
	// Leg5交易所代码
	Leg5ExchangeID TRsaFtdcExchangeIDType
	// Leg5合约代码
	Leg5InstrumentID TRsaFtdcInstrumentIDType
	// Leg5买卖方向
	Leg5Direction TRsaFtdcDirectionType
	// Leg6交易所代码
	Leg6ExchangeID TRsaFtdcExchangeIDType
	// Leg6合约代码
	Leg6InstrumentID TRsaFtdcInstrumentIDType
	// Leg6买卖方向
	Leg6Direction TRsaFtdcDirectionType
	// Bar精度
	BarPreces TRsaFtdcBarPrecesType
	// Bar周期
	BarPeriod TRsaFtdcNumberType
}

// 流水状态
type CRsaFtdcFlowStatusField struct {
	// 序列系列号
	SequenceSeries TRsaFtdcSequenceSeriesType
	// 是否就绪
	Ready TRsaFtdcBoolType
}

// 行情订阅完成请求
type CRsaFtdcReqSubMarketDataCompletedField struct {
	// 当前时间
	CurrTime TRsaFtdcInt64Type
}

// 策略消息
type CRsaFtdcStgMsgField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 用户代码
	UserID TRsaFtdcUserIDType
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
	// 系统报单编号
	OrderSysID TRsaFtdcOrderSysIDType
	// 会话编号
	SessionID TRsaFtdcSessionIDType
	// 策略消息类型
	StgMsgType TRsaFtdcStgMsgTypeType
	// 包数量
	Total TRsaFtdcNumberType
	// 包序号
	Index TRsaFtdcNumberType
	// Json数据串
	Data TRsaFtdcJsonDataType
}

// 策略委托消息
type CRsaFtdcStgOrderField struct {
	// 经纪公司编号
	BrokerID TRsaFtdcBrokerIDType
	// 用户代码
	UserID TRsaFtdcUserIDType
	// 投资者编号
	InvestorID TRsaFtdcInvestorIDType
	// 系统报单编号
	OrderSysID TRsaFtdcOrderSysIDType
	// 会话编号
	SessionID TRsaFtdcSessionIDType
	// 交易所代码
	ExchangeID TRsaFtdcExchangeIDType
	// 合约代码
	InstrumentID TRsaFtdcInstrumentIDType
	// 买卖方向
	Direction TRsaFtdcDirectionType
	// 开平标志
	OffsetFlag TRsaFtdcOffsetFlagType
	// 数量
	Volume TRsaFtdcVolumeType
	// 已经成交数量
	VolumeTraded TRsaFtdcVolumeType
	// 报单状态
	OrderStatus TRsaFtdcOrderStatusType
	// 错误代码
	ErrorID TRsaFtdcErrorIDType
	// 错误信息
	ErrorMsg TRsaFtdcErrorMsgType
}

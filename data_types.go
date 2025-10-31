package rmd4go

//go:generate stringer -type Rsa_TE_RESUME_TYPE -linecomment
type Rsa_TE_RESUME_TYPE int

const (
	Rsa_TERT_RESTART Rsa_TE_RESUME_TYPE = iota // 重启
	Rsa_TERT_RESUME                            // 恢复
	Rsa_TERT_QUICK                             // 快速
)

const (
	// 定义交易通道类型类别
	// 连接中金所
	Rsa_EI_CFFEX = "CFFEX"
	// 连接上期所
	Rsa_EI_SHFE = "SHFE"
	// 连接大商所
	Rsa_EI_DCE = "DCE"
	// 连接郑商所
	Rsa_EI_CZCE = "CZCE"
	// 连接上证所
	Rsa_EI_SSE = "SSE"
	// 连接深交所
	Rsa_EI_SZSE = "SZSE"
	// 连接黄金交易所
	Rsa_EI_SGE = "SGE"
	// 连接芝加哥商品期货交易所
	Rsa_EI_CME = "CME"
	// 连接伦敦金属交易所
	Rsa_EI_LME = "LME"
	// 连接国泰君安的通道
	Rsa_EI_GTJA = "GTJA"
	// 连接上海国际能源中心的通道
	Rsa_EI_INE = "INE"
	// 连接无锡不锈钢
	Rsa_EI_BXG = "BXG"
	// 连接新加坡交易所
	Rsa_EI_SGX = "SGX"
	// 连接纽约商业交易所
	Rsa_EI_NYMEX = "NYMEX"
	// 连接香港交易所
	Rsa_EI_HKEX = "HKEX"
	// 连接马来西亚交易所
	Rsa_EI_BMD = "BMD"
	// 连接芝加哥期货交易所
	Rsa_EI_CBOT = "CBOT"
	// 连接纽约期货交易所
	Rsa_EI_NYBOT = "NYBOT"
	// 连接东京商品交易所
	Rsa_EI_TOCOM = "TOCOM"
	// 连接韩国交易所
	Rsa_EI_KRX = "KRX"
	// 连接伦敦国际金融期货交易所
	Rsa_EI_LIFFE = "LIFFE"
)

// //go:generate stringer -type Rsa_TSS_FLOW_TYPE -linecomment
// type Rsa_TSS_FLOW_TYPE int

// const (
// 	Rsa_TSS_DIALOG  Rsa_TSS_FLOW_TYPE = iota + 1 // 对话流
// 	Rsa_TSS_PRIVATE                              // 客户私有流
// 	Rsa_TSS_PUBLIC                               // 公共流
// 	Rsa_TSS_QUERY                                // 查询
// 	Rsa_TSS_USER                                 // 用户私有流
// 	Rsa_TSS_MARKET                               // MBL行情流
// 	Rsa_TSS_SPEC                                 // UserApi特殊流水流
// )

// TFtdcPriceTickType是一个最小变动价位类型
type TRsaFtdcPriceTickType float64

// TFtdcPriceType是一个价格类型
type TRsaFtdcPriceType float64

// TFtdcRatioType是一个比率类型
type TRsaFtdcRatioType float64

// TFtdcMoneyType是一个资金类型
type TRsaFtdcMoneyType float64

// TFtdcLargeVolumeType是一个大额数量类型
type TRsaFtdcLargeVolumeType float64

// TFtdcPointVolumeType是一个带小数点数量类型
type TRsaFtdcPointVolumeType float64

// TFtdcNumberType是一个编号类型
type TRsaFtdcNumberType int

// TFtdcSequenceNoType是一个序列号类型
type TRsaFtdcSequenceNoType int

// TFtdcMillisecType是一个最后修改毫秒类型
type TRsaFtdcMillisecType int

// TFtdcVolumeType是一个数量类型
type TRsaFtdcVolumeType int

// TFtdcUnderlyingMultipleType是一个合约乘数类型
type TRsaFtdcUnderlyingMultipleType int

// TFtdcVolumeMultipleType是一个数量乘数类型
type TRsaFtdcVolumeMultipleType int

// TFtdcSequenceSeriesType是一个序列系列号类型
type TRsaFtdcSequenceSeriesType int

// TFtdcSessionIDType是一个会话编号类型
type TRsaFtdcSessionIDType int

// TFtdcErrorIDType是一个错误代码类型
type TRsaFtdcErrorIDType int

// TFtdcDataCenterIDType是一个数据中心代码类型
type TRsaFtdcDataCenterIDType int

// TFtdcFrontIDType是一个前置编号类型
type TRsaFtdcFrontIDType int

// TFtdcRequestIDType是一个请求编号类型
type TRsaFtdcRequestIDType int

// TFtdcTopicIDType是一个主题代码类型
type TRsaFtdcTopicIDType int

// TFtdcCommandNoType是一个DB命令序号类型
type TRsaFtdcCommandNoType int

// TFtdcSettlementIDType是一个结算编号类型
type TRsaFtdcSettlementIDType int

// TFtdcTotalNumsType是一个累加次数类型
type TRsaFtdcTotalNumsType int

// TFtdcLegIDType是一个单腿编号类型
type TRsaFtdcLegIDType int

// TFtdcLegMultipleType是一个单腿乘数类型
type TRsaFtdcLegMultipleType int

// TFtdcTradeGroupIDType是一个成交组号类型
type TRsaFtdcTradeGroupIDType int

// TFtdcOrderActionRefType是一个报单操作引用类型
type TRsaFtdcOrderActionRefType int

// TFtdcInstallIDType是一个安装编号类型
type TRsaFtdcInstallIDType int

// TFtdcContentLengthType是一个信息长度类型
type TRsaFtdcContentLengthType int

// TFtdcPortType是一个端口类型
type TRsaFtdcPortType int

// TFtdcZBSGEDType是一个主板申购额度类型
type TRsaFtdcZBSGEDType int

// TFtdcKCBSGEDType是一个科创板申购额度类型
type TRsaFtdcKCBSGEDType int

// TFtdcPacketNoType是一个结算编号类型
type TRsaFtdcPacketNoType int

// TFtdcInt64Type是一个longlong类型
type TRsaFtdcInt64Type int64

// TFtdcMonthType是一个月份类型
type TRsaFtdcMonthType int

// TFtdcBoolType是一个布尔型类型
type TRsaFtdcBoolType int

// TFtdcYearType是一个年类型
type TRsaFtdcYearType int

// TFtdcIsUpdateType是一个密码是否更新类型
type TRsaFtdcIsUpdateType uint8

// TFtdcAliasType是一个别名类型
type TRsaFtdcAliasType string

// TFtdcMemTableNameType是一个内存表名类型
type TRsaFtdcMemTableNameType string

// TFtdcOrderSysIDType是一个报单编号类型
type TRsaFtdcOrderSysIDType string

// TFtdcTradeIDType是一个成交编号类型
type TRsaFtdcTradeIDType string

// TFtdcUserIDType是一个用户代码类型
type TRsaFtdcUserIDType string

// TFtdcTraderIDType是一个用户代码类型
type TRsaFtdcTraderIDType string

// TFtdcParticipantIDType是一个会员编号类型
type TRsaFtdcParticipantIDType string

// TFtdcIPAddressType是一个IP地址类型
type TRsaFtdcIPAddressType string

// TFtdcMacAddressType是一个Mac地址类型
type TRsaFtdcMacAddressType string

// TFtdcInstrumentNameType是一个合约名称类型
type TRsaFtdcInstrumentNameType string

// TFtdcBranchIDType是一个营业部代码类型
type TRsaFtdcBranchIDType string

// TFtdcBranchNameType是一个营业部名称类型
type TRsaFtdcBranchNameType string

// TFtdcInstrumentIDType是一个合约编号类型
type TRsaFtdcInstrumentIDType string

// TFtdcExchangeIDType是一个交易(所)通道类型
type TRsaFtdcExchangeIDType string

// TFtdcExchangeNameType是一个交易(所)名称类型
type TRsaFtdcExchangeNameType string

// TFtdcDateType是一个日期类型
type TRsaFtdcDateType string

// TFtdcTimeType是一个时间类型
type TRsaFtdcTimeType string

// TFtdcClientIDType是一个客户编码类型
type TRsaFtdcClientIDType string

// TFtdcAccountIDType是一个资金帐号类型
type TRsaFtdcAccountIDType string

// TFtdcSeatIDType是一个席位号类型
type TRsaFtdcSeatIDType string

// TFtdcProductNameType是一个品种名称类型
type TRsaFtdcProductNameType string

// TFtdcUserOrderLocalIDType是一个用户本地报单号类型
type TRsaFtdcUserOrderLocalIDType string

// TFtdcOrderLocalIDType是一个本地报单编号类型
type TRsaFtdcOrderLocalIDType string

// TFtdcInvestorIDType是一个投资者编号类型
type TRsaFtdcInvestorIDType string

// TFtdcInvestorNameType是一个投资者名称类型
type TRsaFtdcInvestorNameType string

// TFtdcUserNameType是一个用户编码类型
type TRsaFtdcUserNameType string

// TFtdcPasswordType是一个密码类型
type TRsaFtdcPasswordType string

// TFtdcProductInfoType是一个产品信息类型
type TRsaFtdcProductInfoType string

// TFtdcLoginInfoType是一个登录信息类型
type TRsaFtdcLoginInfoType string

// TFtdcProtocolInfoType是一个协议信息类型
type TRsaFtdcProtocolInfoType string

// TFtdcBusinessUnitType是一个业务单元类型
type TRsaFtdcBusinessUnitType string

// TFtdcTradingSystemNameType是一个交易系统名称类型
type TRsaFtdcTradingSystemNameType string

// TFtdcCommandTypeType是一个DB命令类型类型
type TRsaFtdcCommandTypeType string

// TFtdcBrokerIDType是一个经纪公司编号类型
type TRsaFtdcBrokerIDType string

// TFtdcCustomType是一个用户自定义域类型类型
type TRsaFtdcCustomType string

// TFtdcTradingDayType是一个交易日类型
type TRsaFtdcTradingDayType TRsaFtdcDateType

// TFtdcDepartmentType是一个营业部类型
type TRsaFtdcDepartmentType string

// TFtdcGrantFuncSetType是一个授权功能号类型
type TRsaFtdcGrantFuncSetType [5]uint8

// TFtdcProductIDType是一个品种编号类型
type TRsaFtdcProductIDType string

// TFtdcAccountSeqNoType是一个资金流水号类型
type TRsaFtdcAccountSeqNoType string

// TFtdcInstrumentGroupIDType是一个合约组代码类型
type TRsaFtdcInstrumentGroupIDType string

// TFtdcCurrencyType是一个币种代码类型
type TRsaFtdcCurrencyType string

// TFtdcSubInstrumentIDType是一个订阅信息类型
type TRsaFtdcSubInstrumentIDType string

// TFtdcAdviceUserInfoType是一个用户代码类型
type TRsaFtdcAdviceUserInfoType string

// TFtdcHDSerialIDType是一个硬盘序列号类型
type TRsaFtdcHDSerialIDType string

// TFtdcInvestUnitIDType是一个投资单元代码类型
type TRsaFtdcInvestUnitIDType string

// TFtdcSObjectType是一个状态对象类型
type TRsaFtdcSObjectType string

// TFtdcOrderRefType是一个报单引用类型
type TRsaFtdcOrderRefType string

// TFtdcExchangeInstIDType是一个合约在交易所的代码类型
type TRsaFtdcExchangeInstIDType string

// TFtdcCPBHType是一个产品编号类型
type TRsaFtdcCPBHType string

// TFtdcGDDMType是一个股东代码类型
type TRsaFtdcGDDMType string

// TFtdcOthMarkIdType是一个外部标识类型
type TRsaFtdcOthMarkIdType string

// TFtdcFunctionIdType是一个指令编号类型
type TRsaFtdcFunctionIdType [21]uint8

// TFtdcDataKeyType是一个数据关键字类型
type TRsaFtdcDataKeyType string

// TFtdcSysTagType是一个系统标识类型
type TRsaFtdcSysTagType string

// TFtdcPluginIdType是一个插件代码类型
type TRsaFtdcPluginIdType string

// TFtdcSessionNameType是一个交易节名称类型
type TRsaFtdcSessionNameType string

// TFtdcFormulaType是一个公式类型
type TRsaFtdcFormulaType string

// TFtdcTickBSFlagType是一个Tick买卖标志类型
type TRsaFtdcTickBSFlagType [9]uint8

// TFtdcErrorMsgType是一个错误信息类型
type TRsaFtdcErrorMsgType string

// TFtdcAppIDType是一个认证编码类型
type TRsaFtdcAppIDType string

// TFtdcAuthCodeType是一个授权编码类型
type TRsaFtdcAuthCodeType string

// TFtdcLargeErrorMsgType是一个超长错误信息类型
type TRsaFtdcLargeErrorMsgType string

// TFtdcFieldNameType是一个字段名类型
type TRsaFtdcFieldNameType string

// TFtdcFieldContentType是一个字段内容类型
type TRsaFtdcFieldContentType string

// TFtdcContentType是一个字段内容类型
type TRsaFtdcContentType string

// TFtdcStatusMsgType是一个模块状态信息类型
type TRsaFtdcStatusMsgType string

// TFtdcClientSystemInfoType是一个终端采集信息类型
type TRsaFtdcClientSystemInfoType string

// TFtdcJsonDataType是一个Json内容类型
type TRsaFtdcJsonDataType string

// TFtdcLongJsonDataType是一个Json内容类型
type TRsaFtdcLongJsonDataType string

// TFtdcMessageType是一个信息类型
type TRsaFtdcMessageType string

// TFtdcQuoteIDType是一个系统询价编号类型
type TRsaFtdcQuoteIDType string

// TFtdcMessageTypeType是一个信息类型（字符串）类型
type TRsaFtdcMessageTypeType string

// TFtdcVolumeConditionType是一个成交量类型类型
// 任何数量
const Rsa_FTDC_VC_AV uint8 = '1'

// 最小数量
const Rsa_FTDC_VC_MV uint8 = '2'

// 全部数量
const Rsa_FTDC_VC_CV uint8 = '3'

type TRsaFtdcVolumeConditionType uint8

// TFtdcEnterReasonType是一个进入本状态原因类型
// 自动切换
const Rsa_FTDC_ER_Automatic uint8 = '1'

// 手动切换
const Rsa_FTDC_ER_Manual uint8 = '2'

// 熔断
const Rsa_FTDC_ER_Fuse uint8 = '3'

type TRsaFtdcEnterReasonType uint8

// TFtdcForceCloseReasonType是一个强平原因类型
// 非强平
const Rsa_FTDC_FCR_NotForceClose uint8 = '0'

// 资金不足
const Rsa_FTDC_FCR_LackDeposit uint8 = '1'

// 客户超仓
const Rsa_FTDC_FCR_ClientOverPositionLimit uint8 = '2'

// 会员超仓
const Rsa_FTDC_FCR_MemberOverPositionLimit uint8 = '3'

// 持仓非整数倍
const Rsa_FTDC_FCR_NotMultiple uint8 = '4'

type TRsaFtdcForceCloseReasonType uint8

// TFtdcInstrumentStatusType是一个合约交易状态类型
// 开盘前
const Rsa_FTDC_IS_BeforeTrading uint8 = '0'

// 非交易
const Rsa_FTDC_IS_NoTrading uint8 = '1'

// 连续交易
const Rsa_FTDC_IS_Continous uint8 = '2'

// 集合竞价报单
const Rsa_FTDC_IS_AuctionOrdering uint8 = '3'

// 集合竞价价格平衡
const Rsa_FTDC_IS_AuctionBalance uint8 = '4'

// 集合竞价撮合
const Rsa_FTDC_IS_AuctionMatch uint8 = '5'

// 收盘
const Rsa_FTDC_IS_Closed uint8 = '6'

// 金交所交割申报
const Rsa_FTDC_IS_SGE_Dery_App uint8 = '7'

// 金交所交割申报结束
const Rsa_FTDC_IS_SGE_Dery_Match uint8 = '8'

// 金交所中立仓申报
const Rsa_FTDC_IS_SGE_Mid_App uint8 = '9'

// 金交所交割申报撮合
const Rsa_FTDC_IS_SGE_Mid_Match uint8 = 'a'

type TRsaFtdcInstrumentStatusType uint8

// TFtdcOffsetFlagType是一个开平标志类型
// 开仓
const Rsa_FTDC_OF_Open uint8 = '0'

// 平仓
const Rsa_FTDC_OF_Close uint8 = '1'

// 强平
const Rsa_FTDC_OF_ForceClose uint8 = '2'

// 平今
const Rsa_FTDC_OF_CloseToday uint8 = '3'

// 平昨
const Rsa_FTDC_OF_CloseYesterday uint8 = '4'

type TRsaFtdcOffsetFlagType uint8

// TFtdcOrderPriceTypeType是一个报单价格条件类型
// 任意价
const Rsa_FTDC_OPT_AnyPrice uint8 = '1'

// 限价
const Rsa_FTDC_OPT_LimitPrice uint8 = '2'

// 最优价
const Rsa_FTDC_OPT_BestPrice uint8 = '3'

// 五档价
const Rsa_FTDC_OPT_FiveLevelPrice uint8 = '4'

type TRsaFtdcOrderPriceTypeType uint8

// TFtdcOrderStatusType是一个报单状态类型
// 未知类型
const Rsa_FTDC_OS_Unknow uint8 = 'a'

// 全部成交
const Rsa_FTDC_OS_AllTraded uint8 = '0'

// 部分成交还在队列中
const Rsa_FTDC_OS_PartTradedQueueing uint8 = '1'

// 部分成交不在队列中
const Rsa_FTDC_OS_PartTradedNotQueueing uint8 = '2'

// 未成交还在队列中
const Rsa_FTDC_OS_NoTradeQueueing uint8 = '3'

// 未成交不在队列中
const Rsa_FTDC_OS_NoTradeNotQueueing uint8 = '4'

// 撤单(不在队列中)
const Rsa_FTDC_OS_Canceled uint8 = '5'

// 订单已报入交易所未应答
const Rsa_FTDC_OS_AcceptedNoReply uint8 = '6'

// 部分撤单还在队列中
const Rsa_FTDC_OS_PartCanceledQueueing uint8 = '7'

// 部分成交部分撤单还在队列中
const Rsa_FTDC_OS_PTPCQueueing uint8 = '8'

// 待报入
const Rsa_FTDC_OS_OnReport uint8 = '9'

// 投顾报单
const Rsa_FTDC_OS_ADReciveD uint8 = 'B'

// 投资经理驳回
const Rsa_FTDC_OS_ADRefused uint8 = 'C'

// 投资经理通过
const Rsa_FTDC_OS_ADAllowed uint8 = 'D'

// 交易员已报入
const Rsa_FTDC_OS_ADAccepted uint8 = 'E'

// 交易员驳回
const Rsa_FTDC_OS_ADTraderRefused uint8 = 'F'

// 投顾经理保单
const Rsa_FTDC_OS_ADManager uint8 = 'G'

// 报盘机发送报单成功
const Rsa_FTDC_OS_SentSuccess uint8 = 'H'

// 报盘机发送报单失败
const Rsa_FTDC_OS_SentFailed uint8 = 'I'

type TRsaFtdcOrderStatusType uint8

// TFtdcAdviceRunModeType是一个投顾下单类型类型
// 手动模式
const Rsa_FTDC_ARM_HANDMODE uint8 = '0'

// 自动模式
const Rsa_FTDC_ARM_AUTOMODE uint8 = '1'

type TRsaFtdcAdviceRunModeType uint8

// TFtdcAdviceUserOrderModeType是一个下单标示符类型
// 可下单
const Rsa_FTDC_AUOM_ORDERON uint8 = '0'

// 不可下单
const Rsa_FTDC_AUOM_ORDEROFFE uint8 = '1'

type TRsaFtdcAdviceUserOrderModeType uint8

// TFtdcAdviceSwitchFlagType是一个工作流断线处理标示符类型
// 自动转发
const Rsa_FTDC_ASF_AUYOTRANSITION uint8 = '0'

// 手动转发
const Rsa_FTDC_ASF_HANDTRANSITION uint8 = '1'

// 驳回
const Rsa_FTDC_ASF_Refused uint8 = '2'

// 全部自动
const Rsa_FTDC_ASF_ALLAUTO uint8 = '3'

type TRsaFtdcAdviceSwitchFlagType uint8

// TFtdcUserTypeType是一个用户类型类型
// 交易员
const Rsa_FTDC_UT_Trader uint8 = '1'

// 投资经理
const Rsa_FTDC_UT_InvestManager uint8 = '2'

// 风控管理员
const Rsa_FTDC_UT_Manager uint8 = '3'

// 投资顾问
const Rsa_FTDC_UT_InvestAdvisor uint8 = '4'

// 投资交易员
const Rsa_FTDC_UT_InvestTrader uint8 = '5'

// 超级交易员
const Rsa_FTDC_UT_SuperTrader uint8 = '6'

// 策略执行单元
const Rsa_FTDC_UT_RsaPlus uint8 = '7'

type TRsaFtdcUserTypeType uint8

// TFtdcTradingRightType是一个交易权限类型
// 可以交易
const Rsa_FTDC_TR_Allow uint8 = '0'

// 只能平仓
const Rsa_FTDC_TR_CloseOnly uint8 = '1'

// 不能交易
const Rsa_FTDC_TR_Forbidden uint8 = '2'

// 只能做多(买开,卖平)
const Rsa_FTDC_TR_LongOnly uint8 = '3'

// 只能做空(卖开,买平)
const Rsa_FTDC_TR_ShortOnly uint8 = '4'

type TRsaFtdcTradingRightType uint8

// TFtdcTimeConditionType是一个有效期类型类型
// 立即完成，否则撤销
const Rsa_FTDC_TC_IOC uint8 = '1'

// 本节有效
const Rsa_FTDC_TC_GFS uint8 = '2'

// 当日有效
const Rsa_FTDC_TC_GFD uint8 = '3'

// 指定日期前有效
const Rsa_FTDC_TC_GTD uint8 = '4'

// 撤销前有效
const Rsa_FTDC_TC_GTC uint8 = '5'

// 集合竞价有效
const Rsa_FTDC_TC_GFA uint8 = '6'

type TRsaFtdcTimeConditionType uint8

// TFtdcPosiDirectionType是一个持仓多空方向类型
// 净
const Rsa_FTDC_PD_Net uint8 = '1'

// 多头
const Rsa_FTDC_PD_Long uint8 = '2'

// 空头
const Rsa_FTDC_PD_Short uint8 = '3'

type TRsaFtdcPosiDirectionType uint8

// TFtdcOrderSourceType是一个报单来源类型
// 来自参与者
const Rsa_FTDC_OS_Participant uint8 = '0'

// 来自管理员
const Rsa_FTDC_OS_Administrator uint8 = '1'

type TRsaFtdcOrderSourceType uint8

// TFtdcDirectionType是一个买卖方向类型
// 买
const Rsa_FTDC_D_Buy uint8 = '0'

// 卖
const Rsa_FTDC_D_Sell uint8 = '1'

type TRsaFtdcDirectionType uint8

// TFtdcAccountDirectionType是一个出入金方向类型
// 入金
const Rsa_FTDC_AD_In uint8 = '1'

// 出金
const Rsa_FTDC_AD_Out uint8 = '2'

// 优先入金
const Rsa_FTDC_AD_PIn uint8 = '3'

// 优先出金
const Rsa_FTDC_AD_POut uint8 = '4'

// 融资费用
const Rsa_FTDC_AD_FOut uint8 = '5'

type TRsaFtdcAccountDirectionType uint8

// TFtdcHedgeFlagType是一个投机套保标志类型
// 投机
const Rsa_FTDC_CHF_Speculation uint8 = '1'

// 套利
const Rsa_FTDC_CHF_Arbitrage uint8 = '2'

// 套保
const Rsa_FTDC_CHF_Hedge uint8 = '3'

// 做市商
const Rsa_FTDC_CHF_MarketMaker uint8 = '4'

// 备兑
const Rsa_FTDC_CHF_CoveredOption uint8 = '5'

type TRsaFtdcHedgeFlagType uint8

// TFtdcActionFlagType是一个操作标志类型
// 删除
const Rsa_FTDC_AF_Delete uint8 = '0'

// 挂起
const Rsa_FTDC_AF_Suspend uint8 = '1'

// 激活
const Rsa_FTDC_AF_Active uint8 = '2'

// 修改
const Rsa_FTDC_AF_Modify uint8 = '3'

// 驳回
const Rsa_FTDC_AF_Refused uint8 = '4'

// 分发
const Rsa_FTDC_AF_Allowed uint8 = '5'

// 撤全部未成交委托
const Rsa_FTDC_AF_All uint8 = '6'

// 撤算法委托
const Rsa_FTDC_AF_Algo uint8 = 'p'

// 撤算法委托以及对应的未成交委托
const Rsa_FTDC_AF_AlgoAndOrder uint8 = '7'

// 撤全部算法委托
const Rsa_FTDC_AF_AlgoAll uint8 = '8'

// 撤全部算法委托以及对应的未成交委托
const Rsa_FTDC_AF_AlgoAllAndOrder uint8 = '9'

type TRsaFtdcActionFlagType uint8

// TFtdcOrderActionStatusType是一个报单操作状态类型
// 已经提交
const Rsa_FTDC_OAS_Submitted uint8 = 'a'

// 已经接受
const Rsa_FTDC_OAS_Accepted uint8 = 'b'

// 已经被拒绝
const Rsa_FTDC_OAS_Rejected uint8 = 'c'

type TRsaFtdcOrderActionStatusType uint8

// TFtdcDataSyncStatusType是一个数据同步状态类型
// 未同步
const Rsa_FTDC_DS_Asynchronous uint8 = '1'

// 同步中
const Rsa_FTDC_DS_Synchronizing uint8 = '2'

// 已同步
const Rsa_FTDC_DS_Synchronized uint8 = '3'

type TRsaFtdcDataSyncStatusType uint8

// TFtdcPositionTypeType是一个持仓类型类型
// 净持仓
const Rsa_FTDC_PT_Net uint8 = '1'

// 综合持仓
const Rsa_FTDC_PT_Gross uint8 = '2'

type TRsaFtdcPositionTypeType uint8

// TFtdcOptionsTypeType是一个期权类型类型
// 非期权
const Rsa_FTDC_OT_NotOptions uint8 = '0'

// 看涨
const Rsa_FTDC_OT_CallOptions uint8 = '1'

// 看跌
const Rsa_FTDC_OT_PutOptions uint8 = '2'

type TRsaFtdcOptionsTypeType uint8

// TFtdcIsActiveType是一个是否活跃类型
// 不活跃
const Rsa_FTDC_UIA_NoActive uint8 = '0'

// 活跃
const Rsa_FTDC_UIA_Active uint8 = '1'

type TRsaFtdcIsActiveType uint8

// TFtdcAccountStatusType是一个账户状态类型
// 正常
const Rsa_FTDC_AS_Normal uint8 = '0'

// 追保
const Rsa_FTDC_AS_NotOpen uint8 = '1'

// 强平
const Rsa_FTDC_AS_ForceClose uint8 = '2'

type TRsaFtdcAccountStatusType uint8

// TFtdcProductClassType是一个产品类型类型
// 默认类型，API暂不使用，系统风控用作默认类型，代表所有类型
const Rsa_FTDC_PC_Default uint8 = '0'

// 期货
const Rsa_FTDC_PC_Futures uint8 = '1'

// 期权
const Rsa_FTDC_PC_Options uint8 = '2'

// 组合
const Rsa_FTDC_PC_Combination uint8 = '3'

// 即期
const Rsa_FTDC_PC_Spot uint8 = '4'

// 期转现
const Rsa_FTDC_PC_EFP uint8 = '5'

// 未知类型
const Rsa_FTDC_PC_Unknown uint8 = '6'

// 证券
const Rsa_FTDC_PC_Stocks uint8 = '7'

// 股票期权
const Rsa_FTDC_PC_StockOptions uint8 = '8'

// 金交所现货
const Rsa_FTDC_PC_SGE_SPOT uint8 = '9'

// 金交所递延
const Rsa_FTDC_PC_SGE_DEFER uint8 = 'a'

// 金交所远期
const Rsa_FTDC_PC_SGE_FOWARD uint8 = 'b'

// 外汇远期
const Rsa_FTDC_PC_FXFOWARD uint8 = 'c'

// 数字货币
const Rsa_FTDC_PC_BITCOIN uint8 = 'd'

type TRsaFtdcProductClassType uint8

// TFtdcMarginCombTypeType是一个组合保证金类型类型
// 多空保证金都收取，正常保证金算法
const Rsa_FTDC_MCT_AllMargin uint8 = '0'

// 合约组大边保证金组合
const Rsa_FTDC_MCT_MaxMargin uint8 = '1'

// 组合合约保证金
const Rsa_FTDC_MCT_CombMargin uint8 = '2'

// Span保证金
const Rsa_FTDC_MCT_SpanMargin uint8 = '3'

// 组合剩余合约单向大边
const Rsa_FTDC_MCT_CombLeftMaxMargin uint8 = '4'

// 组合剩余合约双边
const Rsa_FTDC_MCT_CombLeftAllMargin uint8 = '5'

type TRsaFtdcMarginCombTypeType uint8

// TFtdcBusinessTypeType是一个业务类别类型
// 普通
const Rsa_FTDC_BT_Normal uint8 = '1'

// 撤单
const Rsa_FTDC_BT_Cancel uint8 = '2'

// 申赎
const Rsa_FTDC_BT_AppliedForRedeemed uint8 = '3'

// 最优五档即时成交剩余撤销
const Rsa_FTDC_BT_FiveLevelIOC uint8 = '4'

// 最优五档即时成交剩余转限价
const Rsa_FTDC_BT_FiveLevelGFD uint8 = '5'

// 即时成交剩余撤销
const Rsa_FTDC_BT_BestPriceIOC uint8 = '6'

// 全额成交或撤销
const Rsa_FTDC_BT_FOK uint8 = '7'

// 本方最优价格
const Rsa_FTDC_BT_SelfGFD uint8 = '8'

// 对方最优价格
const Rsa_FTDC_BT_CpGFD uint8 = '9'

// 套利组合单
const Rsa_FTDC_BT_Combination uint8 = 'a'

// 行权
const Rsa_FTDC_BT_ExerciseOption uint8 = 'c'

// 金交所中立仓申报
const Rsa_FTDC_BT_SGEMidApp uint8 = 'd'

// 金交所递延交割申报
const Rsa_FTDC_BT_SGEDeferDeliApp uint8 = 'e'

// 互换定单
const Rsa_FTDC_BT_SWAP uint8 = 'f'

// 质押
const Rsa_FTDC_BT_MORTGAGE uint8 = 'g'

// 合并分拆
const Rsa_FTDC_BT_MERGESPLIT uint8 = 'h'

// 转股
const Rsa_FTDC_BT_SWAPEQUITY uint8 = 'i'

// 回售
const Rsa_FTDC_BT_SALEBACK uint8 = 'j'

// 报价衍生委托
const Rsa_FTDC_BT_QuoteOrder uint8 = 'k'

// 融券
const Rsa_FTDC_BT_SecuritiesLoan uint8 = 'l'

// 融资
const Rsa_FTDC_BT_Financing uint8 = 'm'

// 新购申购
const Rsa_FTDC_BT_IPO uint8 = 'n'

// 条件单
const Rsa_FTDC_BT_Conditional uint8 = 'o'

// 算法单
const Rsa_FTDC_BT_Algo uint8 = 'p'

// 策略单
const Rsa_FTDC_BT_Strategy uint8 = 'q'

// 预埋单单
const Rsa_FTDC_BT_Packed uint8 = 'r'

type TRsaFtdcBusinessTypeType uint8

// TFtdcInvestorTypeType是一个投资者类型类型
// 一级投资者
const Rsa_FTDC_CT_FirtstClass uint8 = '1'

// 二级投资者
const Rsa_FTDC_CT_SecondClass uint8 = '2'

type TRsaFtdcInvestorTypeType uint8

// TFtdcLinkTypeType是一个交易编码连接类型类型
// 股票市场
const Rsa_FTDC_LT_Stock uint8 = '0'

// 期货市场
const Rsa_FTDC_LT_Future uint8 = '1'

// 外盘
const Rsa_FTDC_LT_Foreign uint8 = '2'

type TRsaFtdcLinkTypeType uint8

// TFtdcSelfTradeAvoidTypeType是一个投资者连接类型类型
// 不检查自成交
const Rsa_FTDC_STAT_NONE uint8 = '0'

// 将之后的订单打回
const Rsa_FTDC_STAT_RejectAfter uint8 = '1'

// 将之前的订单撤单
const Rsa_FTDC_STAT_CancelBefore uint8 = '2'

// 将之前的订单撤单,成交之后补报
const Rsa_FTDC_STAT_CancelAndReInsert uint8 = '3'

type TRsaFtdcSelfTradeAvoidTypeType uint8

// TFtdcErrorLevelType是一个级别类型
// 无
const Rsa_FTDC_EL_NONE uint8 = '0'

// 紧急
const Rsa_FTDC_EL_EMERGENCY uint8 = '1'

// 危险
const Rsa_FTDC_EL_CRITICAL uint8 = '2'

// 错误
const Rsa_FTDC_EL_ERROR uint8 = '3'

// 警告
const Rsa_FTDC_EL_WARNING uint8 = '4'

// 通知
const Rsa_FTDC_EL_INFO uint8 = '5'

// 调试
const Rsa_FTDC_EL_DEBUG uint8 = '6'

type TRsaFtdcErrorLevelType uint8

// TFtdcTradeTypeType是一个成交类型类型
// 普通成交
const Rsa_FTDC_TT_Normal uint8 = '1'

// 组合成交
const Rsa_FTDC_TT_Combination uint8 = '2'

// 指令单成交
const Rsa_FTDC_TT_Advice uint8 = '3'

// 场外成交
const Rsa_FTDC_TT_OTC uint8 = '4'

// 指令单组合成交
const Rsa_FTDC_TT_AdviceCombination uint8 = '5'

// 组合持仓拆分为单一持仓,初始化不应包含该类型的持仓
const Rsa_FTDC_TT_SplitCombination uint8 = '6'

// 期权执行
const Rsa_FTDC_TT_OptionsExecution uint8 = '7'

// 期转现衍生成交
const Rsa_FTDC_TT_EFPDerived uint8 = '8'

type TRsaFtdcTradeTypeType uint8

// TFtdcOptionsModeType是一个期权行权方式类型
// 不行权
const Rsa_FTDC_OM_NullOptions uint8 = '0'

// 欧式行权
const Rsa_FTDC_OM_EuOptions uint8 = 'E'

// 美式行权
const Rsa_FTDC_OM_AmOptions uint8 = 'A'

type TRsaFtdcOptionsModeType uint8

// TFtdcDistributeTypeType是一个投顾选择资金账户类型类型
// 分配给单一资金账户,直接指定的父账户
const Rsa_FTDC_DT_Single uint8 = '0'

// 分配给单一资金账户,找资金最多的父账户
const Rsa_FTDC_DT_SingleMoneyHigh uint8 = '1'

// 分配给所有资金账户，按资金多少分配，资金越多分的越多
const Rsa_FTDC_DT_MultiMoneyHigh uint8 = 'a'

// 分配给所有资金账户,新的分仓算法（机选顺序）
const Rsa_FTDC_DT_MultiAccountAuto uint8 = 'b'

// 分配给所有资金账户,新的分仓算法（指定顺序）
const Rsa_FTDC_DT_MultiAccountManul uint8 = 'c'

type TRsaFtdcDistributeTypeType uint8

// TFtdcActionTypeType是一个执行类型类型
// 执行
const Rsa_FTDC_AT_Exec uint8 = '1'

// 放弃
const Rsa_FTDC_AT_Abandon uint8 = '2'

type TRsaFtdcActionTypeType uint8

// TFtdcReservePositionFlagTypeType是一个期权行权后是否保留期货头寸的标记类型
// 保留
const Rsa_FTDC_RPFT_Reserve uint8 = '0'

// 不保留
const Rsa_FTDC_RPFT_UnReserve uint8 = '1'

type TRsaFtdcReservePositionFlagTypeType uint8

// TFtdcCloseFlagTypeType是一个期权行权后生成的头寸是否自动平仓类型
// 自动平仓
const Rsa_FTDC_CFT_AutoClose uint8 = '0'

// 免于自动平仓
const Rsa_FTDC_CFT_NotToClose uint8 = '1'

type TRsaFtdcCloseFlagTypeType uint8

// TFtdcExecResultTypeType是一个执行结果类型
// 没有执行
const Rsa_FTDC_ERT_NoExec uint8 = 'n'

// 已经取消
const Rsa_FTDC_ERT_Canceled uint8 = 'c'

// 执行成功
const Rsa_FTDC_ERT_OK uint8 = '0'

// 期权持仓不够
const Rsa_FTDC_ERT_NoPosition uint8 = '1'

// 资金不够
const Rsa_FTDC_ERT_NoDeposit uint8 = '2'

// 会员不存在
const Rsa_FTDC_ERT_NoParticipant uint8 = '3'

// 客户不存在
const Rsa_FTDC_ERT_NoClient uint8 = '4'

// 合约不存在
const Rsa_FTDC_ERT_NoInstrument uint8 = '6'

// 没有执行权限
const Rsa_FTDC_ERT_NoRight uint8 = '7'

// 不合理的数量
const Rsa_FTDC_ERT_InvalidVolume uint8 = '8'

// 没有足够的历史成交
const Rsa_FTDC_ERT_NoEnoughHistoryTrade uint8 = '9'

// 未知
const Rsa_FTDC_ERT_Unknown uint8 = 'a'

type TRsaFtdcExecResultTypeType uint8

// TFtdcForQuoteStatusTypeType是一个询价状态类型
// 已经提交
const Rsa_FTDC_FQST_Submitted uint8 = 'a'

// 已经接受
const Rsa_FTDC_FQST_Accepted uint8 = 'b'

// 已经被拒绝
const Rsa_FTDC_FQST_Rejected uint8 = 'c'

type TRsaFtdcForQuoteStatusTypeType uint8

// TFtdcOptSelfCloseFlagType是一个期权行权的头寸是否自对冲类型类型
// 自对冲期权仓位
const Rsa_FTDC_OSCF_CloseSelfOptionPosition uint8 = '1'

// 保留期权仓位
const Rsa_FTDC_OSCF_ReserveOptionPosition uint8 = '2'

// 自对冲卖方履约后的期货仓位
const Rsa_FTDC_OSCF_SellCloseSelfFuturePosition uint8 = '3'

type TRsaFtdcOptSelfCloseFlagType uint8

// TFtdcOrderSubmitStatusType是一个报单提交状态类型
// 已经提交
const Rsa_FTDC_OSS_InsertSubmitted uint8 = '0'

// 撤单已经提交
const Rsa_FTDC_OSS_CancelSubmitted uint8 = '1'

// 修改已经提交
const Rsa_FTDC_OSS_ModifySubmitted uint8 = '2'

// 已经接受
const Rsa_FTDC_OSS_Accepted uint8 = '3'

// 报单已经被拒绝
const Rsa_FTDC_OSS_InsertRejected uint8 = '4'

// 撤单已经被拒绝
const Rsa_FTDC_OSS_CancelRejected uint8 = '5'

// 改单已经被拒绝
const Rsa_FTDC_OSS_ModifyRejected uint8 = '6'

type TRsaFtdcOrderSubmitStatusType uint8

// TFtdcSeatDirectionType是一个席位状态方向类型
// 登录
const Rsa_FTDC_SD_Login uint8 = '0'

// 登出
const Rsa_FTDC_SD_Logout uint8 = '1'

type TRsaFtdcSeatDirectionType uint8

// TFtdcLockTypeType是一个锁定方向类型
// 锁定
const Rsa_FTDC_LCKT_Lock uint8 = '1'

// 解锁
const Rsa_FTDC_LCKT_Unlock uint8 = '2'

type TRsaFtdcLockTypeType uint8

// TFtdcFeeTypeType是一个费率类型类型
// 融券费用
const Rsa_FTDC_FT_S uint8 = '1'

type TRsaFtdcFeeTypeType uint8

// TFtdcPluginOrderTypeType是一个指令类型类型
// 创建
const Rsa_FTDC_POT_Create uint8 = '1'

// 删除
const Rsa_FTDC_POT_Delete uint8 = '2'

// 更新参数
const Rsa_FTDC_POT_Update uint8 = '3'

// 启动
const Rsa_FTDC_POT_Start uint8 = '4'

// 停止
const Rsa_FTDC_POT_Stop uint8 = '5'

// 删除所有
const Rsa_FTDC_POT_DeleteAll uint8 = '6'

// 上传策略
const Rsa_FTDC_POT_LoadPlugin uint8 = '7'

// 删除策略
const Rsa_FTDC_POT_DelPlugin uint8 = '8'

// 实时加载策略
const Rsa_FTDC_POT_RealLoadPlugin uint8 = '9'

type TRsaFtdcPluginOrderTypeType uint8

// TFtdcPluginStatusType是一个插件运行状态类型
// NONE
const Rsa_FTDC_PS_NONE uint8 = '0'

// 删除
const Rsa_FTDC_PS_Delete uint8 = '1'

// 未设置
const Rsa_FTDC_PS_NotSet uint8 = '2'

// 错误
const Rsa_FTDC_PS_Error uint8 = '3'

// 初始化
const Rsa_FTDC_PS_Init uint8 = '4'

// 运行中
const Rsa_FTDC_PS_Running uint8 = '5'

// 停止
const Rsa_FTDC_PS_Stop uint8 = '6'

// 暂停
const Rsa_FTDC_PS_Pause uint8 = '7'

// 完成
const Rsa_FTDC_PS_Finished uint8 = '8'

// 创建
const Rsa_FTDC_PS_Create uint8 = '9'

type TRsaFtdcPluginStatusType uint8

// TFtdcSysTypeType是一个系统类型类型
// 本系统
const Rsa_FTDC_ST_Self uint8 = '1'

// 经纪商系统
const Rsa_FTDC_ST_Broker uint8 = '2'

type TRsaFtdcSysTypeType uint8

// TFtdcSysStatusType是一个系统状态类型
// 未连接
const Rsa_FTDC_SS_UnConnect uint8 = '1'

// 已连接
const Rsa_FTDC_SS_Connected uint8 = '2'

// 登录成功
const Rsa_FTDC_SS_LoginSuccess uint8 = '3'

// 登录失败
const Rsa_FTDC_SS_LoginFailed uint8 = '4'

// 已就绪
const Rsa_FTDC_SS_Ready uint8 = '5'

type TRsaFtdcSysStatusType uint8

// TFtdcPluginStatusTypeType是一个策略状态类型
// 已经加载
const Rsa_FTDC_PST_Load uint8 = '1'

// 未加载
const Rsa_FTDC_PST_Unload uint8 = '2'

// 全部
const Rsa_FTDC_PST_ALL uint8 = '3'

type TRsaFtdcPluginStatusTypeType uint8

// TFtdcClientReqQuoteStatusType是一个客户询价状态类型
// 未应答
const Rsa_FTDC_CRS_NoRsp uint8 = '1'

// 询价失败
const Rsa_FTDC_CRS_Failed uint8 = '2'

// 询价成功
const Rsa_FTDC_CRS_Success uint8 = '3'

type TRsaFtdcClientReqQuoteStatusType uint8

// TFtdcBarPrecesType是一个Bar精度类型
// 秒
const Rsa_FTDC_BP_Second uint8 = '0'

// 分
const Rsa_FTDC_BP_Minute uint8 = '1'

// 日
const Rsa_FTDC_BP_Day uint8 = '2'

//go:generate stringer -type=TRsaFtdcBarPrecesType -linecomment
type TRsaFtdcBarPrecesType uint8

const (
	BarPrecesSecond TRsaFtdcBarPrecesType = TRsaFtdcBarPrecesType(Rsa_FTDC_BP_Second) // s
	BarPrecesMinute TRsaFtdcBarPrecesType = TRsaFtdcBarPrecesType(Rsa_FTDC_BP_Minute) // m
	BarPrecesDay    TRsaFtdcBarPrecesType = TRsaFtdcBarPrecesType(Rsa_FTDC_BP_Day)    // d
)

// TFtdcOpenCloseMethodType是一个开平方式类型
// 自动
const Rsa_FTDC_OCM_Auto uint8 = '0'

// 平今转开仓
const Rsa_FTDC_OCM_CloseTodayToOpen uint8 = '1'

type TRsaFtdcOpenCloseMethodType uint8

// TFtdcTickTypeType是一个Tick类型类型
// 新增委托
const Rsa_FTDC_TT_Add uint8 = 'A'

// 删除委托
const Rsa_FTDC_TT_Delete uint8 = 'D'

// 产品状态
const Rsa_FTDC_TT_Status uint8 = 'S'

// 成交
const Rsa_FTDC_TT_Trade uint8 = 'T'

type TRsaFtdcTickTypeType uint8

// TFtdcStgMsgTypeType是一个策略消息类型类型
// 委托请求消息
const Rsa_FTDC_SMT_OrderReq uint8 = '1'

// 委托状态消息
const Rsa_FTDC_SMT_OrderRet uint8 = '2'

// 账户信息
const Rsa_FTDC_SMT_Account uint8 = '3'

// 持仓信息
const Rsa_FTDC_SMT_Positions uint8 = '4'

// 全量消息
const Rsa_FTDC_SMT_ALL uint8 = '5'

type TRsaFtdcStgMsgTypeType uint8

package epaylinks

var (
	// 测试地址
	testHost = "http://test-efps.epaylinks.cn"
	// 正式地址
	proHost = "https://efps.epaylinks.cn"

	// 统一下单
	unified = "/api/txs/pay/UnifiedPayment"
	// 支付结果查询接口
	paymentQuery = "/api/txs/pay/PaymentQuery"
	// 微信APP支付接口
	appWxPayMent = "/api/txs/pay/appWxPayMent"
	// 微信H5支付接口
	wxH5Payment = "/api/txs/pay/WxH5Payment"
	// 主扫支付接口
	nativePayment = "/api/txs/pay/NativePayment"
	// 微信公众号/小程序支付接口
	wxJSAPIPayment = "/api/txs/pay/WxJSAPIPayment"
	// 支付宝服务窗支付接口
	aliJSAPIPayment = "/api/txs/pay/AliJSAPIPayment"
	// 扫码（被扫）支付接口
	microPayment = "/api/txs/pay/MicroPayment"
	// 申请快捷支付接口
	quickpayApply = "/api/txs/quickpay/apply"
	// 提交快捷支付接口
	quickpayCommit = "/api/txs/quickpay/commit"
	// 网银支付接口
	unionPayMent = "/api/txs/pay/unionPayMent"
	// 协议支付绑卡预交易
	protocolBindCard = "/api/txs/protocol/bindCard"
	// 协议支付绑卡确认
	protocolBindCardConfirm = "/api/txs/protocol/bindCardConfirm"
	// 协议支付预交易
	protocolPayPre = "/api/txs/protocol/protocolPayPre"
	// 协议支付确认交易
	protocolPayConfirm = "/api/txs/protocol/protocolPayConfirm"
	// 协议支付解绑
	protocolunBindCard = "/api/txs/protocol/unBindCard"
	// 订单撤销接口
	payCancel = "/api/txs/pay/Cancel"
	// 订单撤销接口
	payClose = "/api/txs/pay/close"
)

type OrderInfo struct {
	Id           string      `json:"id"`
	BusinessType string      `json:"businessType"`
	GoodsList    []GoodsList `json:"goodsList"`
}

type GoodsList struct {
	Name   string `json:"name"`
	Number string `json:"number"`
	Amount int64  `json:"amount"`
}

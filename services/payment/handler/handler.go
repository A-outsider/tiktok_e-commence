package handler

import (
	"context"
	"fmt"
	"github.com/smartwalle/alipay/v3"
	"go.uber.org/zap"
	"gomall/gateway/types/resp/common"
	order "gomall/kitex_gen/order"
	payment "gomall/kitex_gen/payment"
	"gomall/services/payment/initialize"
	"log"
	"net/url"
)

// PaymentServiceImpl implements the last service interface defined in the IDL.
type PaymentServiceImpl struct{}

func (p PaymentServiceImpl) CreatePayment(ctx context.Context, req *payment.CreatePaymentReq) (res *payment.CreatePaymentResp, _ error) {
	res = new(payment.CreatePaymentResp)
	res.StatusCode = common.CodeServerBusy

	var pay = alipay.TradePagePay{}
	pay.NotifyURL = initialize.GatewayDomain + "/payment/notify" // 提供给支付宝服务器的地址
	pay.ReturnURL = initialize.GatewayDomain + "/payment/callback"
	pay.Subject = "tiktok支付: " + req.Oid
	pay.OutTradeNo = req.Oid
	pay.TotalAmount = fmt.Sprintf("%.2f", req.Amount)
	pay.TimeoutExpress = "15m" // 过期时间
	pay.ProductCode = "FAST_INSTANT_TRADE_PAY"
	pay.PassbackParams = url.QueryEscape(req.GetUserId()) // 填入userid

	url, err := initialize.GetAlipay().TradePagePay(pay)
	if err != nil {
		zap.L().Error("create payment fail", zap.Error(err))
		return
	}

	res.PaymentUrl = url.String()
	res.StatusCode = common.CodeSuccess
	return
}

func (p PaymentServiceImpl) PayCallback(ctx context.Context, req *payment.PayCallbackReq) (res *payment.PayCallbackResp, _ error) {
	res = new(payment.PayCallbackResp)
	res.StatusCode = common.CodeServerBusy

	// 解析为 form 参数
	values, err := url.ParseQuery(string(req.GetRawData()))
	if err != nil {
		res.StatusCode = common.CodeInvalidParams
		return
	}

	// 获取通知参数
	outTradeNo := values.Get("out_trade_no")

	// 调用 client.VerifySign
	if err = initialize.GetAlipay().VerifySign(values); err != nil {
		zap.L().Error(fmt.Sprintf("验证订单 %s 信息发生错误: %v", outTradeNo, err.Error()))
		res.StatusCode = common.CodePaySignatureVerifyFailed
		return
	}

	// 查询订单状态
	var pay = alipay.TradeQuery{}
	pay.OutTradeNo = outTradeNo

	rsp, err := initialize.GetAlipay().TradeQuery(ctx, pay) // 请求访问支付宝网关询问支付状态
	if err != nil {
		zap.L().Error(fmt.Sprintf("验证订单 %s 信息发生错误: %v", outTradeNo, err.Error()))
		res.StatusCode = common.CodePayMsgError
		return
	}

	if rsp.IsFailure() {
		zap.L().Error(fmt.Sprintf("验证订单 %s 信息发生错误: %s-%s", outTradeNo, rsp.Msg, rsp.SubMsg))
		res.StatusCode = common.CodePayMsgError
		return
	}

	res.StatusCode = common.CodeSuccess
	return
}

func (p PaymentServiceImpl) PayNotify(ctx context.Context, req *payment.PayNotifyReq) (res *payment.PayNotifyResp, _ error) {
	res = new(payment.PayNotifyResp)
	res.StatusCode = common.CodeServerBusy

	// 解析为 form 参数
	values, err := url.ParseQuery(string(req.GetRawData()))
	if err != nil {
		res.StatusCode = common.CodeInvalidParams
		return
	}

	// 解析异步通知
	notification, err := initialize.GetAlipay().DecodeNotification(values) // DecodeNotification 内部已调用 VerifySign 方法验证签名
	if err != nil {
		log.Println("解析异步通知发生错误", err)
		res.StatusCode = common.CodePayMsgError
		return
	}

	// 查询订单状态
	outTradeNo := values.Get("out_trade_no")
	var pay = alipay.TradeQuery{}
	pay.OutTradeNo = outTradeNo

	rsp, err := initialize.GetAlipay().TradeQuery(ctx, pay) // 请求访问支付宝网关询问支付状态
	if err != nil {
		zap.L().Error(fmt.Sprintf("异步验证订单 %s 信息发生错误: %v", outTradeNo, err.Error()))
		res.StatusCode = common.CodePayMsgError
		return
	}

	if rsp.IsFailure() {
		zap.L().Error(fmt.Sprintf("异步验证订单 %s 信息发生错误: %s-%s", outTradeNo, rsp.Msg, rsp.SubMsg))
		res.StatusCode = common.CodePayMsgError
		return
	}

	// 处理业务逻辑
	result, _ := initialize.GetOrderClient().MarkOrderPaid(ctx, &order.MarkOrderPaidReq{OrderId: outTradeNo, UserId: notification.PassbackParams})
	if result.StatusCode != common.CodeSuccess {
		res.StatusCode = result.StatusCode
		return
	}

	res.StatusCode = common.CodeSuccess

	return
}

func NewPaymentServiceImpl() *PaymentServiceImpl {
	return &PaymentServiceImpl{}
}

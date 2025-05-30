// Code generated by Kitex v0.10.3. DO NOT EDIT.

package orderservice

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	order "gomall/kitex_gen/order"
	proto "google.golang.org/protobuf/proto"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"PlaceOrder": kitex.NewMethodInfo(
		placeOrderHandler,
		newPlaceOrderArgs,
		newPlaceOrderResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"ListOrder": kitex.NewMethodInfo(
		listOrderHandler,
		newListOrderArgs,
		newListOrderResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"ListOrderFromSeller": kitex.NewMethodInfo(
		listOrderFromSellerHandler,
		newListOrderFromSellerArgs,
		newListOrderFromSellerResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"MarkOrderPaid": kitex.NewMethodInfo(
		markOrderPaidHandler,
		newMarkOrderPaidArgs,
		newMarkOrderPaidResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"MakeSureOrderExpired": kitex.NewMethodInfo(
		makeSureOrderExpiredHandler,
		newMakeSureOrderExpiredArgs,
		newMakeSureOrderExpiredResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"MarkOrderShipped": kitex.NewMethodInfo(
		markOrderShippedHandler,
		newMarkOrderShippedArgs,
		newMarkOrderShippedResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"MarkOrderCompleted": kitex.NewMethodInfo(
		markOrderCompletedHandler,
		newMarkOrderCompletedArgs,
		newMarkOrderCompletedResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
}

var (
	orderServiceServiceInfo                = NewServiceInfo()
	orderServiceServiceInfoForClient       = NewServiceInfoForClient()
	orderServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return orderServiceServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return orderServiceServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return orderServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "OrderService"
	handlerType := (*order.OrderService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "order",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.10.3",
		Extra:           extra,
	}
	return svcInfo
}

func placeOrderHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(order.PlaceOrderReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(order.OrderService).PlaceOrder(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *PlaceOrderArgs:
		success, err := handler.(order.OrderService).PlaceOrder(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*PlaceOrderResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newPlaceOrderArgs() interface{} {
	return &PlaceOrderArgs{}
}

func newPlaceOrderResult() interface{} {
	return &PlaceOrderResult{}
}

type PlaceOrderArgs struct {
	Req *order.PlaceOrderReq
}

func (p *PlaceOrderArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(order.PlaceOrderReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *PlaceOrderArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *PlaceOrderArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *PlaceOrderArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *PlaceOrderArgs) Unmarshal(in []byte) error {
	msg := new(order.PlaceOrderReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var PlaceOrderArgs_Req_DEFAULT *order.PlaceOrderReq

func (p *PlaceOrderArgs) GetReq() *order.PlaceOrderReq {
	if !p.IsSetReq() {
		return PlaceOrderArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *PlaceOrderArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *PlaceOrderArgs) GetFirstArgument() interface{} {
	return p.Req
}

type PlaceOrderResult struct {
	Success *order.PlaceOrderResp
}

var PlaceOrderResult_Success_DEFAULT *order.PlaceOrderResp

func (p *PlaceOrderResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(order.PlaceOrderResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *PlaceOrderResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *PlaceOrderResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *PlaceOrderResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *PlaceOrderResult) Unmarshal(in []byte) error {
	msg := new(order.PlaceOrderResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *PlaceOrderResult) GetSuccess() *order.PlaceOrderResp {
	if !p.IsSetSuccess() {
		return PlaceOrderResult_Success_DEFAULT
	}
	return p.Success
}

func (p *PlaceOrderResult) SetSuccess(x interface{}) {
	p.Success = x.(*order.PlaceOrderResp)
}

func (p *PlaceOrderResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *PlaceOrderResult) GetResult() interface{} {
	return p.Success
}

func listOrderHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(order.ListOrderReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(order.OrderService).ListOrder(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *ListOrderArgs:
		success, err := handler.(order.OrderService).ListOrder(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*ListOrderResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newListOrderArgs() interface{} {
	return &ListOrderArgs{}
}

func newListOrderResult() interface{} {
	return &ListOrderResult{}
}

type ListOrderArgs struct {
	Req *order.ListOrderReq
}

func (p *ListOrderArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(order.ListOrderReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *ListOrderArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *ListOrderArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *ListOrderArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *ListOrderArgs) Unmarshal(in []byte) error {
	msg := new(order.ListOrderReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ListOrderArgs_Req_DEFAULT *order.ListOrderReq

func (p *ListOrderArgs) GetReq() *order.ListOrderReq {
	if !p.IsSetReq() {
		return ListOrderArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ListOrderArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *ListOrderArgs) GetFirstArgument() interface{} {
	return p.Req
}

type ListOrderResult struct {
	Success *order.ListOrderResp
}

var ListOrderResult_Success_DEFAULT *order.ListOrderResp

func (p *ListOrderResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(order.ListOrderResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *ListOrderResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *ListOrderResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *ListOrderResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *ListOrderResult) Unmarshal(in []byte) error {
	msg := new(order.ListOrderResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ListOrderResult) GetSuccess() *order.ListOrderResp {
	if !p.IsSetSuccess() {
		return ListOrderResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ListOrderResult) SetSuccess(x interface{}) {
	p.Success = x.(*order.ListOrderResp)
}

func (p *ListOrderResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ListOrderResult) GetResult() interface{} {
	return p.Success
}

func listOrderFromSellerHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(order.ListOrderFromSellerReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(order.OrderService).ListOrderFromSeller(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *ListOrderFromSellerArgs:
		success, err := handler.(order.OrderService).ListOrderFromSeller(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*ListOrderFromSellerResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newListOrderFromSellerArgs() interface{} {
	return &ListOrderFromSellerArgs{}
}

func newListOrderFromSellerResult() interface{} {
	return &ListOrderFromSellerResult{}
}

type ListOrderFromSellerArgs struct {
	Req *order.ListOrderFromSellerReq
}

func (p *ListOrderFromSellerArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(order.ListOrderFromSellerReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *ListOrderFromSellerArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *ListOrderFromSellerArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *ListOrderFromSellerArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *ListOrderFromSellerArgs) Unmarshal(in []byte) error {
	msg := new(order.ListOrderFromSellerReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var ListOrderFromSellerArgs_Req_DEFAULT *order.ListOrderFromSellerReq

func (p *ListOrderFromSellerArgs) GetReq() *order.ListOrderFromSellerReq {
	if !p.IsSetReq() {
		return ListOrderFromSellerArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *ListOrderFromSellerArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *ListOrderFromSellerArgs) GetFirstArgument() interface{} {
	return p.Req
}

type ListOrderFromSellerResult struct {
	Success *order.ListOrderFromSellerResp
}

var ListOrderFromSellerResult_Success_DEFAULT *order.ListOrderFromSellerResp

func (p *ListOrderFromSellerResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(order.ListOrderFromSellerResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *ListOrderFromSellerResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *ListOrderFromSellerResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *ListOrderFromSellerResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *ListOrderFromSellerResult) Unmarshal(in []byte) error {
	msg := new(order.ListOrderFromSellerResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *ListOrderFromSellerResult) GetSuccess() *order.ListOrderFromSellerResp {
	if !p.IsSetSuccess() {
		return ListOrderFromSellerResult_Success_DEFAULT
	}
	return p.Success
}

func (p *ListOrderFromSellerResult) SetSuccess(x interface{}) {
	p.Success = x.(*order.ListOrderFromSellerResp)
}

func (p *ListOrderFromSellerResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *ListOrderFromSellerResult) GetResult() interface{} {
	return p.Success
}

func markOrderPaidHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(order.MarkOrderPaidReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(order.OrderService).MarkOrderPaid(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *MarkOrderPaidArgs:
		success, err := handler.(order.OrderService).MarkOrderPaid(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*MarkOrderPaidResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newMarkOrderPaidArgs() interface{} {
	return &MarkOrderPaidArgs{}
}

func newMarkOrderPaidResult() interface{} {
	return &MarkOrderPaidResult{}
}

type MarkOrderPaidArgs struct {
	Req *order.MarkOrderPaidReq
}

func (p *MarkOrderPaidArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(order.MarkOrderPaidReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *MarkOrderPaidArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *MarkOrderPaidArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *MarkOrderPaidArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *MarkOrderPaidArgs) Unmarshal(in []byte) error {
	msg := new(order.MarkOrderPaidReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var MarkOrderPaidArgs_Req_DEFAULT *order.MarkOrderPaidReq

func (p *MarkOrderPaidArgs) GetReq() *order.MarkOrderPaidReq {
	if !p.IsSetReq() {
		return MarkOrderPaidArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MarkOrderPaidArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *MarkOrderPaidArgs) GetFirstArgument() interface{} {
	return p.Req
}

type MarkOrderPaidResult struct {
	Success *order.MarkOrderPaidResp
}

var MarkOrderPaidResult_Success_DEFAULT *order.MarkOrderPaidResp

func (p *MarkOrderPaidResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(order.MarkOrderPaidResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *MarkOrderPaidResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *MarkOrderPaidResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *MarkOrderPaidResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *MarkOrderPaidResult) Unmarshal(in []byte) error {
	msg := new(order.MarkOrderPaidResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MarkOrderPaidResult) GetSuccess() *order.MarkOrderPaidResp {
	if !p.IsSetSuccess() {
		return MarkOrderPaidResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MarkOrderPaidResult) SetSuccess(x interface{}) {
	p.Success = x.(*order.MarkOrderPaidResp)
}

func (p *MarkOrderPaidResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MarkOrderPaidResult) GetResult() interface{} {
	return p.Success
}

func makeSureOrderExpiredHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(order.MakeSureOrderExpiredReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(order.OrderService).MakeSureOrderExpired(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *MakeSureOrderExpiredArgs:
		success, err := handler.(order.OrderService).MakeSureOrderExpired(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*MakeSureOrderExpiredResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newMakeSureOrderExpiredArgs() interface{} {
	return &MakeSureOrderExpiredArgs{}
}

func newMakeSureOrderExpiredResult() interface{} {
	return &MakeSureOrderExpiredResult{}
}

type MakeSureOrderExpiredArgs struct {
	Req *order.MakeSureOrderExpiredReq
}

func (p *MakeSureOrderExpiredArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(order.MakeSureOrderExpiredReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *MakeSureOrderExpiredArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *MakeSureOrderExpiredArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *MakeSureOrderExpiredArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *MakeSureOrderExpiredArgs) Unmarshal(in []byte) error {
	msg := new(order.MakeSureOrderExpiredReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var MakeSureOrderExpiredArgs_Req_DEFAULT *order.MakeSureOrderExpiredReq

func (p *MakeSureOrderExpiredArgs) GetReq() *order.MakeSureOrderExpiredReq {
	if !p.IsSetReq() {
		return MakeSureOrderExpiredArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MakeSureOrderExpiredArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *MakeSureOrderExpiredArgs) GetFirstArgument() interface{} {
	return p.Req
}

type MakeSureOrderExpiredResult struct {
	Success *order.MakeSureOrderExpiredResp
}

var MakeSureOrderExpiredResult_Success_DEFAULT *order.MakeSureOrderExpiredResp

func (p *MakeSureOrderExpiredResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(order.MakeSureOrderExpiredResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *MakeSureOrderExpiredResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *MakeSureOrderExpiredResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *MakeSureOrderExpiredResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *MakeSureOrderExpiredResult) Unmarshal(in []byte) error {
	msg := new(order.MakeSureOrderExpiredResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MakeSureOrderExpiredResult) GetSuccess() *order.MakeSureOrderExpiredResp {
	if !p.IsSetSuccess() {
		return MakeSureOrderExpiredResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MakeSureOrderExpiredResult) SetSuccess(x interface{}) {
	p.Success = x.(*order.MakeSureOrderExpiredResp)
}

func (p *MakeSureOrderExpiredResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MakeSureOrderExpiredResult) GetResult() interface{} {
	return p.Success
}

func markOrderShippedHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(order.MarkOrderShippedReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(order.OrderService).MarkOrderShipped(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *MarkOrderShippedArgs:
		success, err := handler.(order.OrderService).MarkOrderShipped(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*MarkOrderShippedResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newMarkOrderShippedArgs() interface{} {
	return &MarkOrderShippedArgs{}
}

func newMarkOrderShippedResult() interface{} {
	return &MarkOrderShippedResult{}
}

type MarkOrderShippedArgs struct {
	Req *order.MarkOrderShippedReq
}

func (p *MarkOrderShippedArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(order.MarkOrderShippedReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *MarkOrderShippedArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *MarkOrderShippedArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *MarkOrderShippedArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *MarkOrderShippedArgs) Unmarshal(in []byte) error {
	msg := new(order.MarkOrderShippedReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var MarkOrderShippedArgs_Req_DEFAULT *order.MarkOrderShippedReq

func (p *MarkOrderShippedArgs) GetReq() *order.MarkOrderShippedReq {
	if !p.IsSetReq() {
		return MarkOrderShippedArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MarkOrderShippedArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *MarkOrderShippedArgs) GetFirstArgument() interface{} {
	return p.Req
}

type MarkOrderShippedResult struct {
	Success *order.MarkOrderShippedResp
}

var MarkOrderShippedResult_Success_DEFAULT *order.MarkOrderShippedResp

func (p *MarkOrderShippedResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(order.MarkOrderShippedResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *MarkOrderShippedResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *MarkOrderShippedResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *MarkOrderShippedResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *MarkOrderShippedResult) Unmarshal(in []byte) error {
	msg := new(order.MarkOrderShippedResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MarkOrderShippedResult) GetSuccess() *order.MarkOrderShippedResp {
	if !p.IsSetSuccess() {
		return MarkOrderShippedResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MarkOrderShippedResult) SetSuccess(x interface{}) {
	p.Success = x.(*order.MarkOrderShippedResp)
}

func (p *MarkOrderShippedResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MarkOrderShippedResult) GetResult() interface{} {
	return p.Success
}

func markOrderCompletedHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(order.MarkOrderCompletedReq)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(order.OrderService).MarkOrderCompleted(ctx, req)
		if err != nil {
			return err
		}
		return st.SendMsg(resp)
	case *MarkOrderCompletedArgs:
		success, err := handler.(order.OrderService).MarkOrderCompleted(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*MarkOrderCompletedResult)
		realResult.Success = success
		return nil
	default:
		return errInvalidMessageType
	}
}
func newMarkOrderCompletedArgs() interface{} {
	return &MarkOrderCompletedArgs{}
}

func newMarkOrderCompletedResult() interface{} {
	return &MarkOrderCompletedResult{}
}

type MarkOrderCompletedArgs struct {
	Req *order.MarkOrderCompletedReq
}

func (p *MarkOrderCompletedArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(order.MarkOrderCompletedReq)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *MarkOrderCompletedArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *MarkOrderCompletedArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *MarkOrderCompletedArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *MarkOrderCompletedArgs) Unmarshal(in []byte) error {
	msg := new(order.MarkOrderCompletedReq)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var MarkOrderCompletedArgs_Req_DEFAULT *order.MarkOrderCompletedReq

func (p *MarkOrderCompletedArgs) GetReq() *order.MarkOrderCompletedReq {
	if !p.IsSetReq() {
		return MarkOrderCompletedArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *MarkOrderCompletedArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *MarkOrderCompletedArgs) GetFirstArgument() interface{} {
	return p.Req
}

type MarkOrderCompletedResult struct {
	Success *order.MarkOrderCompletedResp
}

var MarkOrderCompletedResult_Success_DEFAULT *order.MarkOrderCompletedResp

func (p *MarkOrderCompletedResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(order.MarkOrderCompletedResp)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *MarkOrderCompletedResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *MarkOrderCompletedResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *MarkOrderCompletedResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *MarkOrderCompletedResult) Unmarshal(in []byte) error {
	msg := new(order.MarkOrderCompletedResp)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *MarkOrderCompletedResult) GetSuccess() *order.MarkOrderCompletedResp {
	if !p.IsSetSuccess() {
		return MarkOrderCompletedResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MarkOrderCompletedResult) SetSuccess(x interface{}) {
	p.Success = x.(*order.MarkOrderCompletedResp)
}

func (p *MarkOrderCompletedResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MarkOrderCompletedResult) GetResult() interface{} {
	return p.Success
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) PlaceOrder(ctx context.Context, Req *order.PlaceOrderReq) (r *order.PlaceOrderResp, err error) {
	var _args PlaceOrderArgs
	_args.Req = Req
	var _result PlaceOrderResult
	if err = p.c.Call(ctx, "PlaceOrder", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ListOrder(ctx context.Context, Req *order.ListOrderReq) (r *order.ListOrderResp, err error) {
	var _args ListOrderArgs
	_args.Req = Req
	var _result ListOrderResult
	if err = p.c.Call(ctx, "ListOrder", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) ListOrderFromSeller(ctx context.Context, Req *order.ListOrderFromSellerReq) (r *order.ListOrderFromSellerResp, err error) {
	var _args ListOrderFromSellerArgs
	_args.Req = Req
	var _result ListOrderFromSellerResult
	if err = p.c.Call(ctx, "ListOrderFromSeller", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MarkOrderPaid(ctx context.Context, Req *order.MarkOrderPaidReq) (r *order.MarkOrderPaidResp, err error) {
	var _args MarkOrderPaidArgs
	_args.Req = Req
	var _result MarkOrderPaidResult
	if err = p.c.Call(ctx, "MarkOrderPaid", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MakeSureOrderExpired(ctx context.Context, Req *order.MakeSureOrderExpiredReq) (r *order.MakeSureOrderExpiredResp, err error) {
	var _args MakeSureOrderExpiredArgs
	_args.Req = Req
	var _result MakeSureOrderExpiredResult
	if err = p.c.Call(ctx, "MakeSureOrderExpired", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MarkOrderShipped(ctx context.Context, Req *order.MarkOrderShippedReq) (r *order.MarkOrderShippedResp, err error) {
	var _args MarkOrderShippedArgs
	_args.Req = Req
	var _result MarkOrderShippedResult
	if err = p.c.Call(ctx, "MarkOrderShipped", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MarkOrderCompleted(ctx context.Context, Req *order.MarkOrderCompletedReq) (r *order.MarkOrderCompletedResp, err error) {
	var _args MarkOrderCompletedArgs
	_args.Req = Req
	var _result MarkOrderCompletedResult
	if err = p.c.Call(ctx, "MarkOrderCompleted", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

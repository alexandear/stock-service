// Code generated by protoc-gen-go. DO NOT EDIT.
// source: stock.proto

package devchallenge

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type OrderSide int32

const (
	OrderSide_UNKNOWN_SIDE OrderSide = 0
	OrderSide_BUY          OrderSide = 1
	OrderSide_SELL         OrderSide = 2
)

var OrderSide_name = map[int32]string{
	0: "UNKNOWN_SIDE",
	1: "BUY",
	2: "SELL",
}

var OrderSide_value = map[string]int32{
	"UNKNOWN_SIDE": 0,
	"BUY":          1,
	"SELL":         2,
}

func (x OrderSide) String() string {
	return proto.EnumName(OrderSide_name, int32(x))
}

func (OrderSide) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c87a7814fbd674bd, []int{0}
}

type OrderType int32

const (
	OrderType_UNKNOWN_TYPE OrderType = 0
	OrderType_LIMIT        OrderType = 1
	OrderType_MARKET       OrderType = 2
)

var OrderType_name = map[int32]string{
	0: "UNKNOWN_TYPE",
	1: "LIMIT",
	2: "MARKET",
}

var OrderType_value = map[string]int32{
	"UNKNOWN_TYPE": 0,
	"LIMIT":        1,
	"MARKET":       2,
}

func (x OrderType) String() string {
	return proto.EnumName(OrderType_name, int32(x))
}

func (OrderType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c87a7814fbd674bd, []int{1}
}

type ResponseStatus int32

const (
	ResponseStatus_UNKNOWN_STATUS ResponseStatus = 0
	ResponseStatus_SUCCESS        ResponseStatus = 1
	ResponseStatus_FAILED         ResponseStatus = 2
)

var ResponseStatus_name = map[int32]string{
	0: "UNKNOWN_STATUS",
	1: "SUCCESS",
	2: "FAILED",
}

var ResponseStatus_value = map[string]int32{
	"UNKNOWN_STATUS": 0,
	"SUCCESS":        1,
	"FAILED":         2,
}

func (x ResponseStatus) String() string {
	return proto.EnumName(ResponseStatus_name, int32(x))
}

func (ResponseStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c87a7814fbd674bd, []int{2}
}

// New order
type CreateOrderRequest struct {
	Symbol               string    `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Side                 OrderSide `protobuf:"varint,2,opt,name=side,proto3,enum=devchallenge.OrderSide" json:"side,omitempty"`
	Type                 OrderType `protobuf:"varint,3,opt,name=type,proto3,enum=devchallenge.OrderType" json:"type,omitempty"`
	Price                string    `protobuf:"bytes,4,opt,name=price,proto3" json:"price,omitempty"`
	Quantity             string    `protobuf:"bytes,5,opt,name=quantity,proto3" json:"quantity,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *CreateOrderRequest) Reset()         { *m = CreateOrderRequest{} }
func (m *CreateOrderRequest) String() string { return proto.CompactTextString(m) }
func (*CreateOrderRequest) ProtoMessage()    {}
func (*CreateOrderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c87a7814fbd674bd, []int{0}
}

func (m *CreateOrderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateOrderRequest.Unmarshal(m, b)
}
func (m *CreateOrderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateOrderRequest.Marshal(b, m, deterministic)
}
func (m *CreateOrderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateOrderRequest.Merge(m, src)
}
func (m *CreateOrderRequest) XXX_Size() int {
	return xxx_messageInfo_CreateOrderRequest.Size(m)
}
func (m *CreateOrderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateOrderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateOrderRequest proto.InternalMessageInfo

func (m *CreateOrderRequest) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *CreateOrderRequest) GetSide() OrderSide {
	if m != nil {
		return m.Side
	}
	return OrderSide_UNKNOWN_SIDE
}

func (m *CreateOrderRequest) GetType() OrderType {
	if m != nil {
		return m.Type
	}
	return OrderType_UNKNOWN_TYPE
}

func (m *CreateOrderRequest) GetPrice() string {
	if m != nil {
		return m.Price
	}
	return ""
}

func (m *CreateOrderRequest) GetQuantity() string {
	if m != nil {
		return m.Quantity
	}
	return ""
}

type CreateOrderResponse struct {
	Status               ResponseStatus `protobuf:"varint,1,opt,name=status,proto3,enum=devchallenge.ResponseStatus" json:"status,omitempty"`
	Message              string         `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Id                   string         `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	Time                 uint64         `protobuf:"varint,4,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CreateOrderResponse) Reset()         { *m = CreateOrderResponse{} }
func (m *CreateOrderResponse) String() string { return proto.CompactTextString(m) }
func (*CreateOrderResponse) ProtoMessage()    {}
func (*CreateOrderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c87a7814fbd674bd, []int{1}
}

func (m *CreateOrderResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateOrderResponse.Unmarshal(m, b)
}
func (m *CreateOrderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateOrderResponse.Marshal(b, m, deterministic)
}
func (m *CreateOrderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateOrderResponse.Merge(m, src)
}
func (m *CreateOrderResponse) XXX_Size() int {
	return xxx_messageInfo_CreateOrderResponse.Size(m)
}
func (m *CreateOrderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateOrderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateOrderResponse proto.InternalMessageInfo

func (m *CreateOrderResponse) GetStatus() ResponseStatus {
	if m != nil {
		return m.Status
	}
	return ResponseStatus_UNKNOWN_STATUS
}

func (m *CreateOrderResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *CreateOrderResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CreateOrderResponse) GetTime() uint64 {
	if m != nil {
		return m.Time
	}
	return 0
}

// Cancel order
type CancelOrderRequest struct {
	Symbol               string   `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CancelOrderRequest) Reset()         { *m = CancelOrderRequest{} }
func (m *CancelOrderRequest) String() string { return proto.CompactTextString(m) }
func (*CancelOrderRequest) ProtoMessage()    {}
func (*CancelOrderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c87a7814fbd674bd, []int{2}
}

func (m *CancelOrderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CancelOrderRequest.Unmarshal(m, b)
}
func (m *CancelOrderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CancelOrderRequest.Marshal(b, m, deterministic)
}
func (m *CancelOrderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CancelOrderRequest.Merge(m, src)
}
func (m *CancelOrderRequest) XXX_Size() int {
	return xxx_messageInfo_CancelOrderRequest.Size(m)
}
func (m *CancelOrderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CancelOrderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CancelOrderRequest proto.InternalMessageInfo

func (m *CancelOrderRequest) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *CancelOrderRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type CancelOrderResponse struct {
	Status               ResponseStatus `protobuf:"varint,1,opt,name=status,proto3,enum=devchallenge.ResponseStatus" json:"status,omitempty"`
	Message              string         `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Time                 uint64         `protobuf:"varint,3,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CancelOrderResponse) Reset()         { *m = CancelOrderResponse{} }
func (m *CancelOrderResponse) String() string { return proto.CompactTextString(m) }
func (*CancelOrderResponse) ProtoMessage()    {}
func (*CancelOrderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c87a7814fbd674bd, []int{3}
}

func (m *CancelOrderResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CancelOrderResponse.Unmarshal(m, b)
}
func (m *CancelOrderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CancelOrderResponse.Marshal(b, m, deterministic)
}
func (m *CancelOrderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CancelOrderResponse.Merge(m, src)
}
func (m *CancelOrderResponse) XXX_Size() int {
	return xxx_messageInfo_CancelOrderResponse.Size(m)
}
func (m *CancelOrderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CancelOrderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CancelOrderResponse proto.InternalMessageInfo

func (m *CancelOrderResponse) GetStatus() ResponseStatus {
	if m != nil {
		return m.Status
	}
	return ResponseStatus_UNKNOWN_STATUS
}

func (m *CancelOrderResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *CancelOrderResponse) GetTime() uint64 {
	if m != nil {
		return m.Time
	}
	return 0
}

// Get order book
type GetOrderBookRequest struct {
	Symbol               string   `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Limit                int32    `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetOrderBookRequest) Reset()         { *m = GetOrderBookRequest{} }
func (m *GetOrderBookRequest) String() string { return proto.CompactTextString(m) }
func (*GetOrderBookRequest) ProtoMessage()    {}
func (*GetOrderBookRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c87a7814fbd674bd, []int{4}
}

func (m *GetOrderBookRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOrderBookRequest.Unmarshal(m, b)
}
func (m *GetOrderBookRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOrderBookRequest.Marshal(b, m, deterministic)
}
func (m *GetOrderBookRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOrderBookRequest.Merge(m, src)
}
func (m *GetOrderBookRequest) XXX_Size() int {
	return xxx_messageInfo_GetOrderBookRequest.Size(m)
}
func (m *GetOrderBookRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOrderBookRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetOrderBookRequest proto.InternalMessageInfo

func (m *GetOrderBookRequest) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *GetOrderBookRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type GetOrderBookResponse struct {
	Bids                 []*OrderBookEntity `protobuf:"bytes,1,rep,name=bids,proto3" json:"bids,omitempty"`
	Asks                 []*OrderBookEntity `protobuf:"bytes,2,rep,name=asks,proto3" json:"asks,omitempty"`
	Time                 uint64             `protobuf:"varint,3,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *GetOrderBookResponse) Reset()         { *m = GetOrderBookResponse{} }
func (m *GetOrderBookResponse) String() string { return proto.CompactTextString(m) }
func (*GetOrderBookResponse) ProtoMessage()    {}
func (*GetOrderBookResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c87a7814fbd674bd, []int{5}
}

func (m *GetOrderBookResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOrderBookResponse.Unmarshal(m, b)
}
func (m *GetOrderBookResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOrderBookResponse.Marshal(b, m, deterministic)
}
func (m *GetOrderBookResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOrderBookResponse.Merge(m, src)
}
func (m *GetOrderBookResponse) XXX_Size() int {
	return xxx_messageInfo_GetOrderBookResponse.Size(m)
}
func (m *GetOrderBookResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOrderBookResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetOrderBookResponse proto.InternalMessageInfo

func (m *GetOrderBookResponse) GetBids() []*OrderBookEntity {
	if m != nil {
		return m.Bids
	}
	return nil
}

func (m *GetOrderBookResponse) GetAsks() []*OrderBookEntity {
	if m != nil {
		return m.Asks
	}
	return nil
}

func (m *GetOrderBookResponse) GetTime() uint64 {
	if m != nil {
		return m.Time
	}
	return 0
}

type OrderBookEntity struct {
	Price                string   `protobuf:"bytes,1,opt,name=price,proto3" json:"price,omitempty"`
	Quantity             string   `protobuf:"bytes,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderBookEntity) Reset()         { *m = OrderBookEntity{} }
func (m *OrderBookEntity) String() string { return proto.CompactTextString(m) }
func (*OrderBookEntity) ProtoMessage()    {}
func (*OrderBookEntity) Descriptor() ([]byte, []int) {
	return fileDescriptor_c87a7814fbd674bd, []int{6}
}

func (m *OrderBookEntity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderBookEntity.Unmarshal(m, b)
}
func (m *OrderBookEntity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderBookEntity.Marshal(b, m, deterministic)
}
func (m *OrderBookEntity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderBookEntity.Merge(m, src)
}
func (m *OrderBookEntity) XXX_Size() int {
	return xxx_messageInfo_OrderBookEntity.Size(m)
}
func (m *OrderBookEntity) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderBookEntity.DiscardUnknown(m)
}

var xxx_messageInfo_OrderBookEntity proto.InternalMessageInfo

func (m *OrderBookEntity) GetPrice() string {
	if m != nil {
		return m.Price
	}
	return ""
}

func (m *OrderBookEntity) GetQuantity() string {
	if m != nil {
		return m.Quantity
	}
	return ""
}

// Get trades history (list of executed orders)
type GetTradesRequest struct {
	Symbol               string   `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Limit                int32    `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	FromTime             uint64   `protobuf:"varint,3,opt,name=from_time,json=fromTime,proto3" json:"from_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTradesRequest) Reset()         { *m = GetTradesRequest{} }
func (m *GetTradesRequest) String() string { return proto.CompactTextString(m) }
func (*GetTradesRequest) ProtoMessage()    {}
func (*GetTradesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c87a7814fbd674bd, []int{7}
}

func (m *GetTradesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTradesRequest.Unmarshal(m, b)
}
func (m *GetTradesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTradesRequest.Marshal(b, m, deterministic)
}
func (m *GetTradesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTradesRequest.Merge(m, src)
}
func (m *GetTradesRequest) XXX_Size() int {
	return xxx_messageInfo_GetTradesRequest.Size(m)
}
func (m *GetTradesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTradesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTradesRequest proto.InternalMessageInfo

func (m *GetTradesRequest) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *GetTradesRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *GetTradesRequest) GetFromTime() uint64 {
	if m != nil {
		return m.FromTime
	}
	return 0
}

type GetTradesResponse struct {
	Trades               []*TradeEntity `protobuf:"bytes,1,rep,name=trades,proto3" json:"trades,omitempty"`
	Time                 uint64         `protobuf:"varint,2,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetTradesResponse) Reset()         { *m = GetTradesResponse{} }
func (m *GetTradesResponse) String() string { return proto.CompactTextString(m) }
func (*GetTradesResponse) ProtoMessage()    {}
func (*GetTradesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c87a7814fbd674bd, []int{8}
}

func (m *GetTradesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTradesResponse.Unmarshal(m, b)
}
func (m *GetTradesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTradesResponse.Marshal(b, m, deterministic)
}
func (m *GetTradesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTradesResponse.Merge(m, src)
}
func (m *GetTradesResponse) XXX_Size() int {
	return xxx_messageInfo_GetTradesResponse.Size(m)
}
func (m *GetTradesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTradesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTradesResponse proto.InternalMessageInfo

func (m *GetTradesResponse) GetTrades() []*TradeEntity {
	if m != nil {
		return m.Trades
	}
	return nil
}

func (m *GetTradesResponse) GetTime() uint64 {
	if m != nil {
		return m.Time
	}
	return 0
}

type TradeEntity struct {
	Symbol               string   `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	OrderId              string   `protobuf:"bytes,2,opt,name=orderId,proto3" json:"orderId,omitempty"`
	Price                string   `protobuf:"bytes,3,opt,name=price,proto3" json:"price,omitempty"`
	Quantity             string   `protobuf:"bytes,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Time                 uint64   `protobuf:"varint,5,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TradeEntity) Reset()         { *m = TradeEntity{} }
func (m *TradeEntity) String() string { return proto.CompactTextString(m) }
func (*TradeEntity) ProtoMessage()    {}
func (*TradeEntity) Descriptor() ([]byte, []int) {
	return fileDescriptor_c87a7814fbd674bd, []int{9}
}

func (m *TradeEntity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TradeEntity.Unmarshal(m, b)
}
func (m *TradeEntity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TradeEntity.Marshal(b, m, deterministic)
}
func (m *TradeEntity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TradeEntity.Merge(m, src)
}
func (m *TradeEntity) XXX_Size() int {
	return xxx_messageInfo_TradeEntity.Size(m)
}
func (m *TradeEntity) XXX_DiscardUnknown() {
	xxx_messageInfo_TradeEntity.DiscardUnknown(m)
}

var xxx_messageInfo_TradeEntity proto.InternalMessageInfo

func (m *TradeEntity) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *TradeEntity) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

func (m *TradeEntity) GetPrice() string {
	if m != nil {
		return m.Price
	}
	return ""
}

func (m *TradeEntity) GetQuantity() string {
	if m != nil {
		return m.Quantity
	}
	return ""
}

func (m *TradeEntity) GetTime() uint64 {
	if m != nil {
		return m.Time
	}
	return 0
}

// Get stats of stock
type GetStatsRequest struct {
	Symbol               string   `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetStatsRequest) Reset()         { *m = GetStatsRequest{} }
func (m *GetStatsRequest) String() string { return proto.CompactTextString(m) }
func (*GetStatsRequest) ProtoMessage()    {}
func (*GetStatsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c87a7814fbd674bd, []int{10}
}

func (m *GetStatsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStatsRequest.Unmarshal(m, b)
}
func (m *GetStatsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStatsRequest.Marshal(b, m, deterministic)
}
func (m *GetStatsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStatsRequest.Merge(m, src)
}
func (m *GetStatsRequest) XXX_Size() int {
	return xxx_messageInfo_GetStatsRequest.Size(m)
}
func (m *GetStatsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStatsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetStatsRequest proto.InternalMessageInfo

func (m *GetStatsRequest) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

type GetStatsResponse struct {
	Volume               string   `protobuf:"bytes,1,opt,name=volume,proto3" json:"volume,omitempty"`
	Commission           string   `protobuf:"bytes,2,opt,name=commission,proto3" json:"commission,omitempty"`
	Time                 uint64   `protobuf:"varint,3,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetStatsResponse) Reset()         { *m = GetStatsResponse{} }
func (m *GetStatsResponse) String() string { return proto.CompactTextString(m) }
func (*GetStatsResponse) ProtoMessage()    {}
func (*GetStatsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c87a7814fbd674bd, []int{11}
}

func (m *GetStatsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStatsResponse.Unmarshal(m, b)
}
func (m *GetStatsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStatsResponse.Marshal(b, m, deterministic)
}
func (m *GetStatsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStatsResponse.Merge(m, src)
}
func (m *GetStatsResponse) XXX_Size() int {
	return xxx_messageInfo_GetStatsResponse.Size(m)
}
func (m *GetStatsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStatsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetStatsResponse proto.InternalMessageInfo

func (m *GetStatsResponse) GetVolume() string {
	if m != nil {
		return m.Volume
	}
	return ""
}

func (m *GetStatsResponse) GetCommission() string {
	if m != nil {
		return m.Commission
	}
	return ""
}

func (m *GetStatsResponse) GetTime() uint64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func init() {
	proto.RegisterEnum("devchallenge.OrderSide", OrderSide_name, OrderSide_value)
	proto.RegisterEnum("devchallenge.OrderType", OrderType_name, OrderType_value)
	proto.RegisterEnum("devchallenge.ResponseStatus", ResponseStatus_name, ResponseStatus_value)
	proto.RegisterType((*CreateOrderRequest)(nil), "devchallenge.CreateOrderRequest")
	proto.RegisterType((*CreateOrderResponse)(nil), "devchallenge.CreateOrderResponse")
	proto.RegisterType((*CancelOrderRequest)(nil), "devchallenge.CancelOrderRequest")
	proto.RegisterType((*CancelOrderResponse)(nil), "devchallenge.CancelOrderResponse")
	proto.RegisterType((*GetOrderBookRequest)(nil), "devchallenge.GetOrderBookRequest")
	proto.RegisterType((*GetOrderBookResponse)(nil), "devchallenge.GetOrderBookResponse")
	proto.RegisterType((*OrderBookEntity)(nil), "devchallenge.OrderBookEntity")
	proto.RegisterType((*GetTradesRequest)(nil), "devchallenge.GetTradesRequest")
	proto.RegisterType((*GetTradesResponse)(nil), "devchallenge.GetTradesResponse")
	proto.RegisterType((*TradeEntity)(nil), "devchallenge.TradeEntity")
	proto.RegisterType((*GetStatsRequest)(nil), "devchallenge.GetStatsRequest")
	proto.RegisterType((*GetStatsResponse)(nil), "devchallenge.GetStatsResponse")
}

func init() { proto.RegisterFile("stock.proto", fileDescriptor_c87a7814fbd674bd) }

var fileDescriptor_c87a7814fbd674bd = []byte{
	// 701 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0x4d, 0x4f, 0xdb, 0x58,
	0x14, 0x8d, 0x1d, 0x3b, 0x89, 0x6f, 0x50, 0xf0, 0x5c, 0x10, 0xe3, 0xc9, 0x0c, 0x0c, 0xf5, 0x8a,
	0x52, 0x09, 0xb5, 0x29, 0xcb, 0x76, 0x01, 0x21, 0x45, 0x11, 0x21, 0x54, 0xb6, 0x23, 0x44, 0xa5,
	0x16, 0x99, 0xf8, 0x95, 0x5a, 0xc4, 0x71, 0xf0, 0x7b, 0x41, 0xca, 0x0f, 0xe8, 0xa2, 0xbb, 0xfe,
	0x8d, 0xfe, 0x83, 0xfe, 0xbc, 0xca, 0xcf, 0x1f, 0xd8, 0x18, 0x03, 0xaa, 0xd4, 0x5d, 0xae, 0x7d,
	0xee, 0x79, 0xe7, 0x1e, 0xdf, 0x77, 0x02, 0x4d, 0xca, 0xfc, 0xf1, 0xd5, 0xce, 0x2c, 0xf0, 0x99,
	0x8f, 0x4b, 0x0e, 0xb9, 0x19, 0x7f, 0xb1, 0x27, 0x13, 0x32, 0xbd, 0x24, 0xfa, 0x4f, 0x01, 0xb0,
	0x1b, 0x10, 0x9b, 0x91, 0x93, 0xc0, 0x21, 0x81, 0x41, 0xae, 0xe7, 0x84, 0x32, 0x5c, 0x83, 0x1a,
	0x5d, 0x78, 0x17, 0xfe, 0x44, 0x13, 0x36, 0x85, 0x2d, 0xc5, 0x88, 0x2b, 0x7c, 0x01, 0x12, 0x75,
	0x1d, 0xa2, 0x89, 0x9b, 0xc2, 0x56, 0xab, 0xf3, 0xf7, 0x4e, 0x96, 0x6b, 0x87, 0x33, 0x98, 0xae,
	0x43, 0x0c, 0x0e, 0x0a, 0xc1, 0x6c, 0x31, 0x23, 0x5a, 0xb5, 0x14, 0x6c, 0x2d, 0x66, 0xc4, 0xe0,
	0x20, 0x5c, 0x05, 0x79, 0x16, 0xb8, 0x63, 0xa2, 0x49, 0xfc, 0xc0, 0xa8, 0xc0, 0x36, 0x34, 0xae,
	0xe7, 0xf6, 0x94, 0xb9, 0x6c, 0xa1, 0xc9, 0xfc, 0x45, 0x5a, 0xeb, 0xdf, 0x04, 0x58, 0xc9, 0x49,
	0xa7, 0x33, 0x7f, 0x4a, 0x09, 0xee, 0x42, 0x8d, 0x32, 0x9b, 0xcd, 0x29, 0xd7, 0xde, 0xea, 0xfc,
	0x97, 0x3f, 0x38, 0xc1, 0x99, 0x1c, 0x63, 0xc4, 0x58, 0xd4, 0xa0, 0xee, 0x11, 0x4a, 0xed, 0xcb,
	0x68, 0x38, 0xc5, 0x48, 0x4a, 0x6c, 0x81, 0xe8, 0x3a, 0x7c, 0x08, 0xc5, 0x10, 0x5d, 0x07, 0x11,
	0x24, 0xe6, 0x7a, 0x91, 0x50, 0xc9, 0xe0, 0xbf, 0xf5, 0x37, 0x80, 0x5d, 0x7b, 0x3a, 0x26, 0x93,
	0x27, 0xb9, 0x18, 0x31, 0x8a, 0x09, 0xa3, 0xbe, 0x80, 0x95, 0x5c, 0xf7, 0x1f, 0x1a, 0x24, 0x11,
	0x5e, 0xcd, 0x08, 0xef, 0xc2, 0xca, 0x21, 0x61, 0xfc, 0xdc, 0x7d, 0xdf, 0xbf, 0x7a, 0x4c, 0xf9,
	0x2a, 0xc8, 0x13, 0xd7, 0x73, 0x19, 0xa7, 0x96, 0x8d, 0xa8, 0xd0, 0xbf, 0x0b, 0xb0, 0x9a, 0x67,
	0x89, 0x27, 0x78, 0x05, 0xd2, 0x85, 0xeb, 0x84, 0xfa, 0xab, 0x5b, 0xcd, 0xce, 0xfa, 0x3d, 0x1b,
	0x10, 0xc2, 0x7b, 0xfc, 0x7b, 0x1a, 0x1c, 0x1a, 0xb6, 0xd8, 0xf4, 0x8a, 0x6a, 0xe2, 0x93, 0x5a,
	0x42, 0x68, 0xc9, 0x5c, 0xcb, 0x77, 0xc0, 0xb7, 0x1b, 0x26, 0x94, 0x6d, 0x98, 0x78, 0x67, 0xc3,
	0x3e, 0x82, 0x7a, 0x48, 0x98, 0x15, 0xd8, 0x0e, 0xa1, 0xbf, 0xe5, 0x0c, 0xfe, 0x0b, 0xca, 0xe7,
	0xc0, 0xf7, 0xce, 0x33, 0xfa, 0x1a, 0xe1, 0x03, 0x2b, 0xd4, 0xf8, 0x01, 0xfe, 0xca, 0xd0, 0xa7,
	0x96, 0xd5, 0x18, 0x7f, 0x12, 0x9b, 0xf6, 0x4f, 0xde, 0x01, 0x8e, 0x8e, 0xa7, 0x8f, 0x81, 0xe9,
	0xfc, 0x62, 0x66, 0xfe, 0xaf, 0x02, 0x34, 0x33, 0xd8, 0x52, 0xd9, 0x1a, 0xd4, 0xfd, 0xd0, 0xa7,
	0x7e, 0xb2, 0x8f, 0x49, 0x79, 0x6b, 0x57, 0xb5, 0xcc, 0x2e, 0x29, 0x6f, 0x57, 0xaa, 0x43, 0xce,
	0xe8, 0x78, 0x0e, 0xcb, 0x87, 0x84, 0x85, 0x2b, 0xfa, 0x98, 0x83, 0xfa, 0x27, 0xee, 0x76, 0x0c,
	0x8d, 0xdd, 0x58, 0x83, 0xda, 0x8d, 0x3f, 0x99, 0x7b, 0xc9, 0x47, 0x8b, 0x2b, 0xdc, 0x00, 0x18,
	0xfb, 0x9e, 0xe7, 0x52, 0xea, 0xfa, 0xd3, 0x58, 0x79, 0xe6, 0xc9, 0x7d, 0x2b, 0xb1, 0xfd, 0x12,
	0x94, 0x34, 0xa1, 0x50, 0x85, 0xa5, 0xd1, 0xf0, 0x68, 0x78, 0x72, 0x3a, 0x3c, 0x37, 0xfb, 0x07,
	0x3d, 0xb5, 0x82, 0x75, 0xa8, 0xee, 0x8f, 0xce, 0x54, 0x01, 0x1b, 0x20, 0x99, 0xbd, 0xc1, 0x40,
	0x15, 0xb7, 0x77, 0xe3, 0x8e, 0x30, 0xa6, 0xb2, 0x1d, 0xd6, 0xd9, 0xfb, 0xb0, 0x43, 0x01, 0x79,
	0xd0, 0x3f, 0xee, 0x5b, 0xaa, 0x80, 0x00, 0xb5, 0xe3, 0x3d, 0xe3, 0xa8, 0x67, 0xa9, 0xe2, 0xf6,
	0x5b, 0x68, 0xe5, 0xaf, 0x26, 0x22, 0xb4, 0xd2, 0xc3, 0xac, 0x3d, 0x6b, 0x64, 0xaa, 0x15, 0x6c,
	0x42, 0xdd, 0x1c, 0x75, 0xbb, 0x3d, 0xd3, 0x8c, 0xda, 0xdf, 0xed, 0xf5, 0x07, 0xbd, 0x03, 0x55,
	0xec, 0xfc, 0xa8, 0x82, 0x6c, 0x86, 0x79, 0x8d, 0x16, 0x34, 0x33, 0xf9, 0x86, 0x9b, 0xf9, 0x4d,
	0x28, 0xa6, 0x76, 0xfb, 0xd9, 0x03, 0x88, 0x48, 0x90, 0x5e, 0xe1, 0xac, 0xb7, 0x61, 0x53, 0x60,
	0x2d, 0xa4, 0x58, 0x81, 0xb5, 0x98, 0x54, 0x7a, 0x05, 0x4f, 0x61, 0x29, 0x9b, 0x00, 0x78, 0xa7,
	0xe9, 0x9e, 0x8c, 0x69, 0xeb, 0x0f, 0x41, 0x52, 0xe2, 0x21, 0x28, 0xe9, 0x25, 0xc1, 0x8d, 0x42,
	0x4b, 0xee, 0x72, 0xb6, 0xff, 0x2f, 0x7d, 0x9f, 0xf2, 0x1d, 0x41, 0x23, 0xd9, 0x32, 0x5c, 0x2f,
	0xc0, 0xb3, 0x8b, 0xda, 0xde, 0x28, 0x7b, 0x9d, 0x90, 0x5d, 0xd4, 0xf8, 0x5f, 0xea, 0xeb, 0x5f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xd5, 0xea, 0x4e, 0x95, 0x61, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StockClient is the client API for Stock service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StockClient interface {
	CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderResponse, error)
	CancelOrder(ctx context.Context, in *CancelOrderRequest, opts ...grpc.CallOption) (*CancelOrderResponse, error)
	GetOrderBook(ctx context.Context, in *GetOrderBookRequest, opts ...grpc.CallOption) (*GetOrderBookResponse, error)
	GetTrades(ctx context.Context, in *GetTradesRequest, opts ...grpc.CallOption) (*GetTradesResponse, error)
	GetStats(ctx context.Context, in *GetStatsRequest, opts ...grpc.CallOption) (*GetStatsResponse, error)
}

type stockClient struct {
	cc *grpc.ClientConn
}

func NewStockClient(cc *grpc.ClientConn) StockClient {
	return &stockClient{cc}
}

func (c *stockClient) CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderResponse, error) {
	out := new(CreateOrderResponse)
	err := c.cc.Invoke(ctx, "/devchallenge.Stock/CreateOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockClient) CancelOrder(ctx context.Context, in *CancelOrderRequest, opts ...grpc.CallOption) (*CancelOrderResponse, error) {
	out := new(CancelOrderResponse)
	err := c.cc.Invoke(ctx, "/devchallenge.Stock/CancelOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockClient) GetOrderBook(ctx context.Context, in *GetOrderBookRequest, opts ...grpc.CallOption) (*GetOrderBookResponse, error) {
	out := new(GetOrderBookResponse)
	err := c.cc.Invoke(ctx, "/devchallenge.Stock/GetOrderBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockClient) GetTrades(ctx context.Context, in *GetTradesRequest, opts ...grpc.CallOption) (*GetTradesResponse, error) {
	out := new(GetTradesResponse)
	err := c.cc.Invoke(ctx, "/devchallenge.Stock/GetTrades", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockClient) GetStats(ctx context.Context, in *GetStatsRequest, opts ...grpc.CallOption) (*GetStatsResponse, error) {
	out := new(GetStatsResponse)
	err := c.cc.Invoke(ctx, "/devchallenge.Stock/GetStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StockServer is the server API for Stock service.
type StockServer interface {
	CreateOrder(context.Context, *CreateOrderRequest) (*CreateOrderResponse, error)
	CancelOrder(context.Context, *CancelOrderRequest) (*CancelOrderResponse, error)
	GetOrderBook(context.Context, *GetOrderBookRequest) (*GetOrderBookResponse, error)
	GetTrades(context.Context, *GetTradesRequest) (*GetTradesResponse, error)
	GetStats(context.Context, *GetStatsRequest) (*GetStatsResponse, error)
}

func RegisterStockServer(s *grpc.Server, srv StockServer) {
	s.RegisterService(&_Stock_serviceDesc, srv)
}

func _Stock_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/devchallenge.Stock/CreateOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockServer).CreateOrder(ctx, req.(*CreateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stock_CancelOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockServer).CancelOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/devchallenge.Stock/CancelOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockServer).CancelOrder(ctx, req.(*CancelOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stock_GetOrderBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockServer).GetOrderBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/devchallenge.Stock/GetOrderBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockServer).GetOrderBook(ctx, req.(*GetOrderBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stock_GetTrades_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTradesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockServer).GetTrades(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/devchallenge.Stock/GetTrades",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockServer).GetTrades(ctx, req.(*GetTradesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stock_GetStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockServer).GetStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/devchallenge.Stock/GetStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockServer).GetStats(ctx, req.(*GetStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Stock_serviceDesc = grpc.ServiceDesc{
	ServiceName: "devchallenge.Stock",
	HandlerType: (*StockServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrder",
			Handler:    _Stock_CreateOrder_Handler,
		},
		{
			MethodName: "CancelOrder",
			Handler:    _Stock_CancelOrder_Handler,
		},
		{
			MethodName: "GetOrderBook",
			Handler:    _Stock_GetOrderBook_Handler,
		},
		{
			MethodName: "GetTrades",
			Handler:    _Stock_GetTrades_Handler,
		},
		{
			MethodName: "GetStats",
			Handler:    _Stock_GetStats_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stock.proto",
}

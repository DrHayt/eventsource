// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.4
// source: events.proto

package orders

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type OrderState int32

const (
	OrderState_UNKNOWN   OrderState = 0
	OrderState_CREATED   OrderState = 1
	OrderState_SHIPPED   OrderState = 2
	OrderState_CANCELLED OrderState = 3
)

// Enum value maps for OrderState.
var (
	OrderState_name = map[int32]string{
		0: "UNKNOWN",
		1: "CREATED",
		2: "SHIPPED",
		3: "CANCELLED",
	}
	OrderState_value = map[string]int32{
		"UNKNOWN":   0,
		"CREATED":   1,
		"SHIPPED":   2,
		"CANCELLED": 3,
	}
)

func (x OrderState) Enum() *OrderState {
	p := new(OrderState)
	*p = x
	return p
}

func (x OrderState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OrderState) Descriptor() protoreflect.EnumDescriptor {
	return file_events_proto_enumTypes[0].Descriptor()
}

func (OrderState) Type() protoreflect.EnumType {
	return &file_events_proto_enumTypes[0]
}

func (x OrderState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OrderState.Descriptor instead.
func (OrderState) EnumDescriptor() ([]byte, []int) {
	return file_events_proto_rawDescGZIP(), []int{0}
}

// Might not be safe based on nil event data.
type EventData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Version int32  `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	At      int64  `protobuf:"varint,3,opt,name=at,proto3" json:"at,omitempty"`
}

func (x *EventData) Reset() {
	*x = EventData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_events_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventData) ProtoMessage() {}

func (x *EventData) ProtoReflect() protoreflect.Message {
	mi := &file_events_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventData.ProtoReflect.Descriptor instead.
func (*EventData) Descriptor() ([]byte, []int) {
	return file_events_proto_rawDescGZIP(), []int{0}
}

func (x *EventData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *EventData) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *EventData) GetAt() int64 {
	if x != nil {
		return x.At
	}
	return 0
}

type OrderSnapshot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// EventData ed = 1;
	Id        string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Version   int32      `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	At        int64      `protobuf:"varint,3,opt,name=at,proto3" json:"at,omitempty"`
	Name      string     `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	FirstName string     `protobuf:"bytes,5,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string     `protobuf:"bytes,6,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	State     OrderState `protobuf:"varint,7,opt,name=state,proto3,enum=orders.OrderState" json:"state,omitempty"`
}

func (x *OrderSnapshot) Reset() {
	*x = OrderSnapshot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_events_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderSnapshot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderSnapshot) ProtoMessage() {}

func (x *OrderSnapshot) ProtoReflect() protoreflect.Message {
	mi := &file_events_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderSnapshot.ProtoReflect.Descriptor instead.
func (*OrderSnapshot) Descriptor() ([]byte, []int) {
	return file_events_proto_rawDescGZIP(), []int{1}
}

func (x *OrderSnapshot) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *OrderSnapshot) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *OrderSnapshot) GetAt() int64 {
	if x != nil {
		return x.At
	}
	return 0
}

func (x *OrderSnapshot) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OrderSnapshot) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *OrderSnapshot) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *OrderSnapshot) GetState() OrderState {
	if x != nil {
		return x.State
	}
	return OrderState_UNKNOWN
}

type OrderCreated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Version int32  `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	At      int64  `protobuf:"varint,3,opt,name=at,proto3" json:"at,omitempty"`
	By      string `protobuf:"bytes,4,opt,name=by,proto3" json:"by,omitempty"`
}

func (x *OrderCreated) Reset() {
	*x = OrderCreated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_events_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderCreated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderCreated) ProtoMessage() {}

func (x *OrderCreated) ProtoReflect() protoreflect.Message {
	mi := &file_events_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderCreated.ProtoReflect.Descriptor instead.
func (*OrderCreated) Descriptor() ([]byte, []int) {
	return file_events_proto_rawDescGZIP(), []int{2}
}

func (x *OrderCreated) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *OrderCreated) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *OrderCreated) GetAt() int64 {
	if x != nil {
		return x.At
	}
	return 0
}

func (x *OrderCreated) GetBy() string {
	if x != nil {
		return x.By
	}
	return ""
}

type OrderCancelled struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Version      int32  `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	At           int64  `protobuf:"varint,3,opt,name=at,proto3" json:"at,omitempty"`
	CancelReason string `protobuf:"bytes,4,opt,name=cancel_reason,json=cancelReason,proto3" json:"cancel_reason,omitempty"`
}

func (x *OrderCancelled) Reset() {
	*x = OrderCancelled{}
	if protoimpl.UnsafeEnabled {
		mi := &file_events_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderCancelled) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderCancelled) ProtoMessage() {}

func (x *OrderCancelled) ProtoReflect() protoreflect.Message {
	mi := &file_events_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderCancelled.ProtoReflect.Descriptor instead.
func (*OrderCancelled) Descriptor() ([]byte, []int) {
	return file_events_proto_rawDescGZIP(), []int{3}
}

func (x *OrderCancelled) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *OrderCancelled) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *OrderCancelled) GetAt() int64 {
	if x != nil {
		return x.At
	}
	return 0
}

func (x *OrderCancelled) GetCancelReason() string {
	if x != nil {
		return x.CancelReason
	}
	return ""
}

type OrderAudited struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Version int32  `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	At      int64  `protobuf:"varint,3,opt,name=at,proto3" json:"at,omitempty"`
}

func (x *OrderAudited) Reset() {
	*x = OrderAudited{}
	if protoimpl.UnsafeEnabled {
		mi := &file_events_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderAudited) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderAudited) ProtoMessage() {}

func (x *OrderAudited) ProtoReflect() protoreflect.Message {
	mi := &file_events_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderAudited.ProtoReflect.Descriptor instead.
func (*OrderAudited) Descriptor() ([]byte, []int) {
	return file_events_proto_rawDescGZIP(), []int{4}
}

func (x *OrderAudited) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *OrderAudited) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *OrderAudited) GetAt() int64 {
	if x != nil {
		return x.At
	}
	return 0
}

type OrderNameChanged struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Version   int32  `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	At        int64  `protobuf:"varint,3,opt,name=at,proto3" json:"at,omitempty"`
	FirstName string `protobuf:"bytes,4,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,5,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
}

func (x *OrderNameChanged) Reset() {
	*x = OrderNameChanged{}
	if protoimpl.UnsafeEnabled {
		mi := &file_events_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderNameChanged) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderNameChanged) ProtoMessage() {}

func (x *OrderNameChanged) ProtoReflect() protoreflect.Message {
	mi := &file_events_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderNameChanged.ProtoReflect.Descriptor instead.
func (*OrderNameChanged) Descriptor() ([]byte, []int) {
	return file_events_proto_rawDescGZIP(), []int{5}
}

func (x *OrderNameChanged) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *OrderNameChanged) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *OrderNameChanged) GetAt() int64 {
	if x != nil {
		return x.At
	}
	return 0
}

func (x *OrderNameChanged) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *OrderNameChanged) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

type EventContainer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type             int32             `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	OrderCreated     *OrderCreated     `protobuf:"bytes,2,opt,name=order_created,json=orderCreated,proto3" json:"order_created,omitempty"`
	OrderCancelled   *OrderCancelled   `protobuf:"bytes,3,opt,name=order_cancelled,json=orderCancelled,proto3" json:"order_cancelled,omitempty"`
	OrderAudited     *OrderAudited     `protobuf:"bytes,4,opt,name=order_audited,json=orderAudited,proto3" json:"order_audited,omitempty"`
	OrderNameChanged *OrderNameChanged `protobuf:"bytes,5,opt,name=order_name_changed,json=orderNameChanged,proto3" json:"order_name_changed,omitempty"`
}

func (x *EventContainer) Reset() {
	*x = EventContainer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_events_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventContainer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventContainer) ProtoMessage() {}

func (x *EventContainer) ProtoReflect() protoreflect.Message {
	mi := &file_events_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventContainer.ProtoReflect.Descriptor instead.
func (*EventContainer) Descriptor() ([]byte, []int) {
	return file_events_proto_rawDescGZIP(), []int{6}
}

func (x *EventContainer) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *EventContainer) GetOrderCreated() *OrderCreated {
	if x != nil {
		return x.OrderCreated
	}
	return nil
}

func (x *EventContainer) GetOrderCancelled() *OrderCancelled {
	if x != nil {
		return x.OrderCancelled
	}
	return nil
}

func (x *EventContainer) GetOrderAudited() *OrderAudited {
	if x != nil {
		return x.OrderAudited
	}
	return nil
}

func (x *EventContainer) GetOrderNameChanged() *OrderNameChanged {
	if x != nil {
		return x.OrderNameChanged
	}
	return nil
}

var File_events_proto protoreflect.FileDescriptor

var file_events_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x22, 0x45, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a,
	0x02, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x61, 0x74, 0x22, 0xc3, 0x01,
	0x0a, 0x0d, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x61, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x61, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x22, 0x58, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a,
	0x02, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x61, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x62, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x62, 0x79, 0x22, 0x6f, 0x0a,
	0x0e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x65, 0x64, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x61, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x61, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x61, 0x6e,
	0x63, 0x65, 0x6c, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x22, 0x48,
	0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x41, 0x75, 0x64, 0x69, 0x74, 0x65, 0x64, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x61, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x61, 0x74, 0x22, 0x88, 0x01, 0x0a, 0x10, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x61, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x61, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x22, 0xa4, 0x02, 0x0a, 0x0f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x39, 0x0a, 0x0d, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x52, 0x0c, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x3f, 0x0a, 0x0f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f,
	0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x61,
	0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x65, 0x64, 0x52, 0x0e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x61,
	0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x65, 0x64, 0x12, 0x39, 0x0a, 0x0d, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x5f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x41, 0x75, 0x64,
	0x69, 0x74, 0x65, 0x64, 0x52, 0x0c, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x41, 0x75, 0x64, 0x69, 0x74,
	0x65, 0x64, 0x12, 0x46, 0x0a, 0x12, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x52, 0x10, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x2a, 0x42, 0x0a, 0x0a, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x44,
	0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x48, 0x49, 0x50, 0x50, 0x45, 0x44, 0x10, 0x02, 0x12,
	0x0d, 0x0a, 0x09, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x42, 0x35,
	0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x61, 0x6e,
	0x63, 0x65, 0x6c, 0x6f, 0x6e, 0x67, 0x77, 0x69, 0x6c, 0x6c, 0x2f, 0x5f, 0x65, 0x78, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x73, 0x2f, 0x66, 0x75, 0x6c, 0x6c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_events_proto_rawDescOnce sync.Once
	file_events_proto_rawDescData = file_events_proto_rawDesc
)

func file_events_proto_rawDescGZIP() []byte {
	file_events_proto_rawDescOnce.Do(func() {
		file_events_proto_rawDescData = protoimpl.X.CompressGZIP(file_events_proto_rawDescData)
	})
	return file_events_proto_rawDescData
}

var file_events_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_events_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_events_proto_goTypes = []interface{}{
	(OrderState)(0),          // 0: orders.OrderState
	(*EventData)(nil),        // 1: orders.EventData
	(*OrderSnapshot)(nil),    // 2: orders.OrderSnapshot
	(*OrderCreated)(nil),     // 3: orders.OrderCreated
	(*OrderCancelled)(nil),   // 4: orders.OrderCancelled
	(*OrderAudited)(nil),     // 5: orders.OrderAudited
	(*OrderNameChanged)(nil), // 6: orders.OrderNameChanged
	(*EventContainer)(nil),   // 7: orders.event_container
}
var file_events_proto_depIdxs = []int32{
	0, // 0: orders.OrderSnapshot.state:type_name -> orders.OrderState
	3, // 1: orders.event_container.order_created:type_name -> orders.OrderCreated
	4, // 2: orders.event_container.order_cancelled:type_name -> orders.OrderCancelled
	5, // 3: orders.event_container.order_audited:type_name -> orders.OrderAudited
	6, // 4: orders.event_container.order_name_changed:type_name -> orders.OrderNameChanged
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_events_proto_init() }
func file_events_proto_init() {
	if File_events_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_events_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_events_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderSnapshot); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_events_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderCreated); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_events_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderCancelled); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_events_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderAudited); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_events_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderNameChanged); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_events_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventContainer); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_events_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_events_proto_goTypes,
		DependencyIndexes: file_events_proto_depIdxs,
		EnumInfos:         file_events_proto_enumTypes,
		MessageInfos:      file_events_proto_msgTypes,
	}.Build()
	File_events_proto = out.File
	file_events_proto_rawDesc = nil
	file_events_proto_goTypes = nil
	file_events_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.4
// source: events.proto

package clients

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

type EventContainer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type        int32        `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Created     *Created     `protobuf:"bytes,2,opt,name=created,proto3" json:"created,omitempty"`
	Activated   *Activated   `protobuf:"bytes,3,opt,name=activated,proto3" json:"activated,omitempty"`
	DeActivated *DeActivated `protobuf:"bytes,4,opt,name=de_activated,json=deActivated,proto3" json:"de_activated,omitempty"`
	NameChanged *NameChanged `protobuf:"bytes,5,opt,name=name_changed,json=nameChanged,proto3" json:"name_changed,omitempty"`
}

func (x *EventContainer) Reset() {
	*x = EventContainer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_events_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventContainer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventContainer) ProtoMessage() {}

func (x *EventContainer) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use EventContainer.ProtoReflect.Descriptor instead.
func (*EventContainer) Descriptor() ([]byte, []int) {
	return file_events_proto_rawDescGZIP(), []int{0}
}

func (x *EventContainer) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *EventContainer) GetCreated() *Created {
	if x != nil {
		return x.Created
	}
	return nil
}

func (x *EventContainer) GetActivated() *Activated {
	if x != nil {
		return x.Activated
	}
	return nil
}

func (x *EventContainer) GetDeActivated() *DeActivated {
	if x != nil {
		return x.DeActivated
	}
	return nil
}

func (x *EventContainer) GetNameChanged() *NameChanged {
	if x != nil {
		return x.NameChanged
	}
	return nil
}

type Created struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Version int32  `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	At      int64  `protobuf:"varint,3,opt,name=at,proto3" json:"at,omitempty"`
	By      string `protobuf:"bytes,4,opt,name=by,proto3" json:"by,omitempty"`
}

func (x *Created) Reset() {
	*x = Created{}
	if protoimpl.UnsafeEnabled {
		mi := &file_events_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Created) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Created) ProtoMessage() {}

func (x *Created) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Created.ProtoReflect.Descriptor instead.
func (*Created) Descriptor() ([]byte, []int) {
	return file_events_proto_rawDescGZIP(), []int{1}
}

func (x *Created) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Created) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Created) GetAt() int64 {
	if x != nil {
		return x.At
	}
	return 0
}

func (x *Created) GetBy() string {
	if x != nil {
		return x.By
	}
	return ""
}

type Activated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Version int32  `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	At      int64  `protobuf:"varint,3,opt,name=at,proto3" json:"at,omitempty"`
	By      string `protobuf:"bytes,4,opt,name=by,proto3" json:"by,omitempty"`
}

func (x *Activated) Reset() {
	*x = Activated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_events_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Activated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Activated) ProtoMessage() {}

func (x *Activated) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Activated.ProtoReflect.Descriptor instead.
func (*Activated) Descriptor() ([]byte, []int) {
	return file_events_proto_rawDescGZIP(), []int{2}
}

func (x *Activated) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Activated) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Activated) GetAt() int64 {
	if x != nil {
		return x.At
	}
	return 0
}

func (x *Activated) GetBy() string {
	if x != nil {
		return x.By
	}
	return ""
}

type DeActivated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Version int32  `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	At      int64  `protobuf:"varint,3,opt,name=at,proto3" json:"at,omitempty"`
	By      string `protobuf:"bytes,4,opt,name=by,proto3" json:"by,omitempty"`
}

func (x *DeActivated) Reset() {
	*x = DeActivated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_events_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeActivated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeActivated) ProtoMessage() {}

func (x *DeActivated) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use DeActivated.ProtoReflect.Descriptor instead.
func (*DeActivated) Descriptor() ([]byte, []int) {
	return file_events_proto_rawDescGZIP(), []int{3}
}

func (x *DeActivated) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DeActivated) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *DeActivated) GetAt() int64 {
	if x != nil {
		return x.At
	}
	return 0
}

func (x *DeActivated) GetBy() string {
	if x != nil {
		return x.By
	}
	return ""
}

type NameChanged struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Version int32  `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	At      int64  `protobuf:"varint,3,opt,name=at,proto3" json:"at,omitempty"`
	Name    string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	By      string `protobuf:"bytes,5,opt,name=by,proto3" json:"by,omitempty"`
}

func (x *NameChanged) Reset() {
	*x = NameChanged{}
	if protoimpl.UnsafeEnabled {
		mi := &file_events_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NameChanged) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NameChanged) ProtoMessage() {}

func (x *NameChanged) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use NameChanged.ProtoReflect.Descriptor instead.
func (*NameChanged) Descriptor() ([]byte, []int) {
	return file_events_proto_rawDescGZIP(), []int{4}
}

func (x *NameChanged) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *NameChanged) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *NameChanged) GetAt() int64 {
	if x != nil {
		return x.At
	}
	return 0
}

func (x *NameChanged) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NameChanged) GetBy() string {
	if x != nil {
		return x.By
	}
	return ""
}

var File_events_proto protoreflect.FileDescriptor

var file_events_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x22, 0xf4, 0x01, 0x0a, 0x0e, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x2a,
	0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x30, 0x0a, 0x09, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x64, 0x52, 0x09, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x64, 0x12, 0x37, 0x0a, 0x0c,
	0x64, 0x65, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x44, 0x65, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x64, 0x52, 0x0b, 0x64, 0x65, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x64, 0x12, 0x37, 0x0a, 0x0c, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x63, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65,
	0x64, 0x52, 0x0b, 0x6e, 0x61, 0x6d, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x64, 0x22, 0x53,
	0x0a, 0x07, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x61, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x62, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x62, 0x79, 0x22, 0x55, 0x0a, 0x09, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x64,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x61, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x61, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x62, 0x79,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x62, 0x79, 0x22, 0x57, 0x0a, 0x0b, 0x44, 0x65,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x61, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x62, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x62, 0x79, 0x22, 0x6b, 0x0a, 0x0b, 0x4e, 0x61, 0x6d, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02,
	0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x61, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x62, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x62, 0x79,
	0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76,
	0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6f, 0x6e, 0x67, 0x77, 0x69, 0x6c, 0x6c, 0x2f, 0x5f, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f, 0x66, 0x75, 0x6c, 0x6c, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_events_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_events_proto_goTypes = []interface{}{
	(*EventContainer)(nil), // 0: clients.EventContainer
	(*Created)(nil),        // 1: clients.Created
	(*Activated)(nil),      // 2: clients.Activated
	(*DeActivated)(nil),    // 3: clients.DeActivated
	(*NameChanged)(nil),    // 4: clients.NameChanged
}
var file_events_proto_depIdxs = []int32{
	1, // 0: clients.EventContainer.created:type_name -> clients.Created
	2, // 1: clients.EventContainer.activated:type_name -> clients.Activated
	3, // 2: clients.EventContainer.de_activated:type_name -> clients.DeActivated
	4, // 3: clients.EventContainer.name_changed:type_name -> clients.NameChanged
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_events_proto_init() }
func file_events_proto_init() {
	if File_events_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_events_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_events_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Created); i {
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
			switch v := v.(*Activated); i {
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
			switch v := v.(*DeActivated); i {
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
			switch v := v.(*NameChanged); i {
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
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_events_proto_goTypes,
		DependencyIndexes: file_events_proto_depIdxs,
		MessageInfos:      file_events_proto_msgTypes,
	}.Build()
	File_events_proto = out.File
	file_events_proto_rawDesc = nil
	file_events_proto_goTypes = nil
	file_events_proto_depIdxs = nil
}

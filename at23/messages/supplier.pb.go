// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.2
// source: supplier.proto

package messages

import (
	actor "github.com/asynkron/protoactor-go/actor"
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

type GetItems struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items         []*Item    `protobuf:"bytes,1,rep,name=Items,proto3" json:"Items,omitempty"`
	TransactionId string     `protobuf:"bytes,2,opt,name=TransactionId,proto3" json:"TransactionId,omitempty"`
	Sender        *actor.PID `protobuf:"bytes,3,opt,name=Sender,proto3" json:"Sender,omitempty"`
}

func (x *GetItems) Reset() {
	*x = GetItems{}
	if protoimpl.UnsafeEnabled {
		mi := &file_supplier_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetItems) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetItems) ProtoMessage() {}

func (x *GetItems) ProtoReflect() protoreflect.Message {
	mi := &file_supplier_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetItems.ProtoReflect.Descriptor instead.
func (*GetItems) Descriptor() ([]byte, []int) {
	return file_supplier_proto_rawDescGZIP(), []int{0}
}

func (x *GetItems) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *GetItems) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

func (x *GetItems) GetSender() *actor.PID {
	if x != nil {
		return x.Sender
	}
	return nil
}

type ReturnItems struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items         []*Item `protobuf:"bytes,1,rep,name=Items,proto3" json:"Items,omitempty"`
	TransactionId string  `protobuf:"bytes,2,opt,name=TransactionId,proto3" json:"TransactionId,omitempty"`
}

func (x *ReturnItems) Reset() {
	*x = ReturnItems{}
	if protoimpl.UnsafeEnabled {
		mi := &file_supplier_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReturnItems) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReturnItems) ProtoMessage() {}

func (x *ReturnItems) ProtoReflect() protoreflect.Message {
	mi := &file_supplier_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReturnItems.ProtoReflect.Descriptor instead.
func (*ReturnItems) Descriptor() ([]byte, []int) {
	return file_supplier_proto_rawDescGZIP(), []int{1}
}

func (x *ReturnItems) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ReturnItems) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

type CheckPrice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items  []*Item    `protobuf:"bytes,1,rep,name=Items,proto3" json:"Items,omitempty"`
	Sender *actor.PID `protobuf:"bytes,2,opt,name=Sender,proto3" json:"Sender,omitempty"`
}

func (x *CheckPrice) Reset() {
	*x = CheckPrice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_supplier_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckPrice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckPrice) ProtoMessage() {}

func (x *CheckPrice) ProtoReflect() protoreflect.Message {
	mi := &file_supplier_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckPrice.ProtoReflect.Descriptor instead.
func (*CheckPrice) Descriptor() ([]byte, []int) {
	return file_supplier_proto_rawDescGZIP(), []int{2}
}

func (x *CheckPrice) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *CheckPrice) GetSender() *actor.PID {
	if x != nil {
		return x.Sender
	}
	return nil
}

type ReturnPrice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Price float32 `protobuf:"fixed32,1,opt,name=Price,proto3" json:"Price,omitempty"`
}

func (x *ReturnPrice) Reset() {
	*x = ReturnPrice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_supplier_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReturnPrice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReturnPrice) ProtoMessage() {}

func (x *ReturnPrice) ProtoReflect() protoreflect.Message {
	mi := &file_supplier_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReturnPrice.ProtoReflect.Descriptor instead.
func (*ReturnPrice) Descriptor() ([]byte, []int) {
	return file_supplier_proto_rawDescGZIP(), []int{3}
}

func (x *ReturnPrice) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type RegisterSupplier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=Address,proto3" json:"Address,omitempty"`
}

func (x *RegisterSupplier) Reset() {
	*x = RegisterSupplier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_supplier_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterSupplier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterSupplier) ProtoMessage() {}

func (x *RegisterSupplier) ProtoReflect() protoreflect.Message {
	mi := &file_supplier_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterSupplier.ProtoReflect.Descriptor instead.
func (*RegisterSupplier) Descriptor() ([]byte, []int) {
	return file_supplier_proto_rawDescGZIP(), []int{4}
}

func (x *RegisterSupplier) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

var File_supplier_proto protoreflect.FileDescriptor

var file_supplier_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x1a, 0x0e, 0x63, 0x6f, 0x6e, 0x73,
	0x75, 0x6d, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x61, 0x63, 0x74, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7a, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x49, 0x74,
	0x65, 0x6d, 0x73, 0x12, 0x24, 0x0a, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x22, 0x0a, 0x06, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x50, 0x49, 0x44, 0x52, 0x06, 0x53, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x22, 0x59, 0x0a, 0x0b, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x49, 0x74, 0x65,
	0x6d, 0x73, 0x12, 0x24, 0x0a, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x56,
	0x0a, 0x0a, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x24, 0x0a, 0x05,
	0x49, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x49, 0x74, 0x65,
	0x6d, 0x73, 0x12, 0x22, 0x0a, 0x06, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x50, 0x49, 0x44, 0x52, 0x06,
	0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x22, 0x23, 0x0a, 0x0b, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x22, 0x2c, 0x0a, 0x10, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x12,
	0x18, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42, 0x0f, 0x5a, 0x0d, 0x61, 0x74, 0x32,
	0x33, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_supplier_proto_rawDescOnce sync.Once
	file_supplier_proto_rawDescData = file_supplier_proto_rawDesc
)

func file_supplier_proto_rawDescGZIP() []byte {
	file_supplier_proto_rawDescOnce.Do(func() {
		file_supplier_proto_rawDescData = protoimpl.X.CompressGZIP(file_supplier_proto_rawDescData)
	})
	return file_supplier_proto_rawDescData
}

var file_supplier_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_supplier_proto_goTypes = []interface{}{
	(*GetItems)(nil),         // 0: messages.GetItems
	(*ReturnItems)(nil),      // 1: messages.ReturnItems
	(*CheckPrice)(nil),       // 2: messages.CheckPrice
	(*ReturnPrice)(nil),      // 3: messages.ReturnPrice
	(*RegisterSupplier)(nil), // 4: messages.RegisterSupplier
	(*Item)(nil),             // 5: messages.Item
	(*actor.PID)(nil),        // 6: actor.PID
}
var file_supplier_proto_depIdxs = []int32{
	5, // 0: messages.GetItems.Items:type_name -> messages.Item
	6, // 1: messages.GetItems.Sender:type_name -> actor.PID
	5, // 2: messages.ReturnItems.Items:type_name -> messages.Item
	5, // 3: messages.CheckPrice.Items:type_name -> messages.Item
	6, // 4: messages.CheckPrice.Sender:type_name -> actor.PID
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_supplier_proto_init() }
func file_supplier_proto_init() {
	if File_supplier_proto != nil {
		return
	}
	file_consumer_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_supplier_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetItems); i {
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
		file_supplier_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReturnItems); i {
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
		file_supplier_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckPrice); i {
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
		file_supplier_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReturnPrice); i {
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
		file_supplier_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterSupplier); i {
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
			RawDescriptor: file_supplier_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_supplier_proto_goTypes,
		DependencyIndexes: file_supplier_proto_depIdxs,
		MessageInfos:      file_supplier_proto_msgTypes,
	}.Build()
	File_supplier_proto = out.File
	file_supplier_proto_rawDesc = nil
	file_supplier_proto_goTypes = nil
	file_supplier_proto_depIdxs = nil
}

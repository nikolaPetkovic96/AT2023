// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.2
// source: consumer.proto

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

type Simulate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Simulate) Reset() {
	*x = Simulate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_consumer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Simulate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Simulate) ProtoMessage() {}

func (x *Simulate) ProtoReflect() protoreflect.Message {
	mi := &file_consumer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Simulate.ProtoReflect.Descriptor instead.
func (*Simulate) Descriptor() ([]byte, []int) {
	return file_consumer_proto_rawDescGZIP(), []int{0}
}

type StartSimulation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Item `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *StartSimulation) Reset() {
	*x = StartSimulation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_consumer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartSimulation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartSimulation) ProtoMessage() {}

func (x *StartSimulation) ProtoReflect() protoreflect.Message {
	mi := &file_consumer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartSimulation.ProtoReflect.Descriptor instead.
func (*StartSimulation) Descriptor() ([]byte, []int) {
	return file_consumer_proto_rawDescGZIP(), []int{1}
}

func (x *StartSimulation) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

type Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemId string `protobuf:"bytes,1,opt,name=item_id,json=itemId,proto3" json:"item_id,omitempty"`
	Amount int32  `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *Item) Reset() {
	*x = Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_consumer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Item) ProtoMessage() {}

func (x *Item) ProtoReflect() protoreflect.Message {
	mi := &file_consumer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Item.ProtoReflect.Descriptor instead.
func (*Item) Descriptor() ([]byte, []int) {
	return file_consumer_proto_rawDescGZIP(), []int{2}
}

func (x *Item) GetItemId() string {
	if x != nil {
		return x.ItemId
	}
	return ""
}

func (x *Item) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type BuyProduct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionId string     `protobuf:"bytes,1,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	Items         []*Item    `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	Sender        *actor.PID `protobuf:"bytes,3,opt,name=Sender,proto3" json:"Sender,omitempty"`
}

func (x *BuyProduct) Reset() {
	*x = BuyProduct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_consumer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuyProduct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuyProduct) ProtoMessage() {}

func (x *BuyProduct) ProtoReflect() protoreflect.Message {
	mi := &file_consumer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuyProduct.ProtoReflect.Descriptor instead.
func (*BuyProduct) Descriptor() ([]byte, []int) {
	return file_consumer_proto_rawDescGZIP(), []int{3}
}

func (x *BuyProduct) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

func (x *BuyProduct) GetItems() []*Item {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *BuyProduct) GetSender() *actor.PID {
	if x != nil {
		return x.Sender
	}
	return nil
}

type CompletedTransaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionId string `protobuf:"bytes,1,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
}

func (x *CompletedTransaction) Reset() {
	*x = CompletedTransaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_consumer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompletedTransaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompletedTransaction) ProtoMessage() {}

func (x *CompletedTransaction) ProtoReflect() protoreflect.Message {
	mi := &file_consumer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompletedTransaction.ProtoReflect.Descriptor instead.
func (*CompletedTransaction) Descriptor() ([]byte, []int) {
	return file_consumer_proto_rawDescGZIP(), []int{4}
}

func (x *CompletedTransaction) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

var File_consumer_proto protoreflect.FileDescriptor

var file_consumer_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x1a, 0x0b, 0x61, 0x63, 0x74, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0a, 0x0a, 0x08, 0x53, 0x69, 0x6d, 0x75, 0x6c,
	0x61, 0x74, 0x65, 0x22, 0x37, 0x0a, 0x0f, 0x53, 0x74, 0x61, 0x72, 0x74, 0x53, 0x69, 0x6d, 0x75,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x37, 0x0a, 0x04,
	0x49, 0x74, 0x65, 0x6d, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x7d, 0x0a, 0x0a, 0x42, 0x75, 0x79, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x12, 0x22, 0x0a, 0x06, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0a, 0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x50, 0x49, 0x44, 0x52, 0x06, 0x53, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x22, 0x3d, 0x0a, 0x14, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x0e,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x42, 0x0f, 0x5a, 0x0d, 0x61, 0x74, 0x32, 0x33, 0x2f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_consumer_proto_rawDescOnce sync.Once
	file_consumer_proto_rawDescData = file_consumer_proto_rawDesc
)

func file_consumer_proto_rawDescGZIP() []byte {
	file_consumer_proto_rawDescOnce.Do(func() {
		file_consumer_proto_rawDescData = protoimpl.X.CompressGZIP(file_consumer_proto_rawDescData)
	})
	return file_consumer_proto_rawDescData
}

var file_consumer_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_consumer_proto_goTypes = []interface{}{
	(*Simulate)(nil),             // 0: messages.Simulate
	(*StartSimulation)(nil),      // 1: messages.StartSimulation
	(*Item)(nil),                 // 2: messages.Item
	(*BuyProduct)(nil),           // 3: messages.BuyProduct
	(*CompletedTransaction)(nil), // 4: messages.CompletedTransaction
	(*actor.PID)(nil),            // 5: actor.PID
}
var file_consumer_proto_depIdxs = []int32{
	2, // 0: messages.StartSimulation.items:type_name -> messages.Item
	2, // 1: messages.BuyProduct.items:type_name -> messages.Item
	5, // 2: messages.BuyProduct.Sender:type_name -> actor.PID
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_consumer_proto_init() }
func file_consumer_proto_init() {
	if File_consumer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_consumer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Simulate); i {
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
		file_consumer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartSimulation); i {
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
		file_consumer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Item); i {
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
		file_consumer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuyProduct); i {
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
		file_consumer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompletedTransaction); i {
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
			RawDescriptor: file_consumer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_consumer_proto_goTypes,
		DependencyIndexes: file_consumer_proto_depIdxs,
		MessageInfos:      file_consumer_proto_msgTypes,
	}.Build()
	File_consumer_proto = out.File
	file_consumer_proto_rawDesc = nil
	file_consumer_proto_goTypes = nil
	file_consumer_proto_depIdxs = nil
}

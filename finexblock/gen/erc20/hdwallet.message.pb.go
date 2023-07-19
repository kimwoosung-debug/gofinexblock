// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.23.4
// source: proto/erc20/hdwallet.message.proto

package erc20

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

type CreateWalletInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID uint64 `protobuf:"varint,1,opt,name=UserID,json=user_id,proto3" json:"UserID,omitempty"`
}

func (x *CreateWalletInput) Reset() {
	*x = CreateWalletInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_erc20_hdwallet_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateWalletInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWalletInput) ProtoMessage() {}

func (x *CreateWalletInput) ProtoReflect() protoreflect.Message {
	mi := &file_proto_erc20_hdwallet_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWalletInput.ProtoReflect.Descriptor instead.
func (*CreateWalletInput) Descriptor() ([]byte, []int) {
	return file_proto_erc20_hdwallet_message_proto_rawDescGZIP(), []int{0}
}

func (x *CreateWalletInput) GetUserID() uint64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

type CreateWalletOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=Address,json=address,proto3" json:"Address,omitempty"`
}

func (x *CreateWalletOutput) Reset() {
	*x = CreateWalletOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_erc20_hdwallet_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateWalletOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWalletOutput) ProtoMessage() {}

func (x *CreateWalletOutput) ProtoReflect() protoreflect.Message {
	mi := &file_proto_erc20_hdwallet_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWalletOutput.ProtoReflect.Descriptor instead.
func (*CreateWalletOutput) Descriptor() ([]byte, []int) {
	return file_proto_erc20_hdwallet_message_proto_rawDescGZIP(), []int{1}
}

func (x *CreateWalletOutput) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type GetBalanceInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=Address,json=address,proto3" json:"Address,omitempty"`
}

func (x *GetBalanceInput) Reset() {
	*x = GetBalanceInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_erc20_hdwallet_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBalanceInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBalanceInput) ProtoMessage() {}

func (x *GetBalanceInput) ProtoReflect() protoreflect.Message {
	mi := &file_proto_erc20_hdwallet_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBalanceInput.ProtoReflect.Descriptor instead.
func (*GetBalanceInput) Descriptor() ([]byte, []int) {
	return file_proto_erc20_hdwallet_message_proto_rawDescGZIP(), []int{2}
}

func (x *GetBalanceInput) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type GetBalanceOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Balance string `protobuf:"bytes,1,opt,name=Balance,json=balance,proto3" json:"Balance,omitempty"`
}

func (x *GetBalanceOutput) Reset() {
	*x = GetBalanceOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_erc20_hdwallet_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBalanceOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBalanceOutput) ProtoMessage() {}

func (x *GetBalanceOutput) ProtoReflect() protoreflect.Message {
	mi := &file_proto_erc20_hdwallet_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBalanceOutput.ProtoReflect.Descriptor instead.
func (*GetBalanceOutput) Descriptor() ([]byte, []int) {
	return file_proto_erc20_hdwallet_message_proto_rawDescGZIP(), []int{3}
}

func (x *GetBalanceOutput) GetBalance() string {
	if x != nil {
		return x.Balance
	}
	return ""
}

var File_proto_erc20_hdwallet_message_proto protoreflect.FileDescriptor

var file_proto_erc20_hdwallet_message_proto_rawDesc = []byte{
	0x0a, 0x22, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x72, 0x63, 0x32, 0x30, 0x2f, 0x68, 0x64,
	0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x65, 0x72, 0x63, 0x32, 0x30, 0x22, 0x2c, 0x0a, 0x11, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x12, 0x17, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x22, 0x2e, 0x0a, 0x12, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x2b, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x18, 0x0a, 0x07,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x2c, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x42, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x42, 0x08, 0x5a, 0x06, 0x2f, 0x65, 0x72, 0x63, 0x32, 0x30, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_erc20_hdwallet_message_proto_rawDescOnce sync.Once
	file_proto_erc20_hdwallet_message_proto_rawDescData = file_proto_erc20_hdwallet_message_proto_rawDesc
)

func file_proto_erc20_hdwallet_message_proto_rawDescGZIP() []byte {
	file_proto_erc20_hdwallet_message_proto_rawDescOnce.Do(func() {
		file_proto_erc20_hdwallet_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_erc20_hdwallet_message_proto_rawDescData)
	})
	return file_proto_erc20_hdwallet_message_proto_rawDescData
}

var file_proto_erc20_hdwallet_message_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_erc20_hdwallet_message_proto_goTypes = []interface{}{
	(*CreateWalletInput)(nil),  // 0: erc20.CreateWalletInput
	(*CreateWalletOutput)(nil), // 1: erc20.CreateWalletOutput
	(*GetBalanceInput)(nil),    // 2: erc20.GetBalanceInput
	(*GetBalanceOutput)(nil),   // 3: erc20.GetBalanceOutput
}
var file_proto_erc20_hdwallet_message_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_erc20_hdwallet_message_proto_init() }
func file_proto_erc20_hdwallet_message_proto_init() {
	if File_proto_erc20_hdwallet_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_erc20_hdwallet_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateWalletInput); i {
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
		file_proto_erc20_hdwallet_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateWalletOutput); i {
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
		file_proto_erc20_hdwallet_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBalanceInput); i {
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
		file_proto_erc20_hdwallet_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBalanceOutput); i {
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
			RawDescriptor: file_proto_erc20_hdwallet_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_erc20_hdwallet_message_proto_goTypes,
		DependencyIndexes: file_proto_erc20_hdwallet_message_proto_depIdxs,
		MessageInfos:      file_proto_erc20_hdwallet_message_proto_msgTypes,
	}.Build()
	File_proto_erc20_hdwallet_message_proto = out.File
	file_proto_erc20_hdwallet_message_proto_rawDesc = nil
	file_proto_erc20_hdwallet_message_proto_goTypes = nil
	file_proto_erc20_hdwallet_message_proto_depIdxs = nil
}

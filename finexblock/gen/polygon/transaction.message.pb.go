// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.1
// source: proto/polygon/transaction.message.proto

package polygon

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

type CreateRawTransactionInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From   string `protobuf:"bytes,1,opt,name=From,json=from,proto3" json:"From,omitempty"`
	To     string `protobuf:"bytes,2,opt,name=To,json=to,proto3" json:"To,omitempty"`
	Amount string `protobuf:"bytes,3,opt,name=Amount,json=amount,proto3" json:"Amount,omitempty"`
}

func (x *CreateRawTransactionInput) Reset() {
	*x = CreateRawTransactionInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_polygon_transaction_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRawTransactionInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRawTransactionInput) ProtoMessage() {}

func (x *CreateRawTransactionInput) ProtoReflect() protoreflect.Message {
	mi := &file_proto_polygon_transaction_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRawTransactionInput.ProtoReflect.Descriptor instead.
func (*CreateRawTransactionInput) Descriptor() ([]byte, []int) {
	return file_proto_polygon_transaction_message_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRawTransactionInput) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *CreateRawTransactionInput) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *CreateRawTransactionInput) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

type CreateRawTransactionOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TxHash string `protobuf:"bytes,1,opt,name=TxHash,json=tx_hash,proto3" json:"TxHash,omitempty"`
}

func (x *CreateRawTransactionOutput) Reset() {
	*x = CreateRawTransactionOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_polygon_transaction_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRawTransactionOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRawTransactionOutput) ProtoMessage() {}

func (x *CreateRawTransactionOutput) ProtoReflect() protoreflect.Message {
	mi := &file_proto_polygon_transaction_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRawTransactionOutput.ProtoReflect.Descriptor instead.
func (*CreateRawTransactionOutput) Descriptor() ([]byte, []int) {
	return file_proto_polygon_transaction_message_proto_rawDescGZIP(), []int{1}
}

func (x *CreateRawTransactionOutput) GetTxHash() string {
	if x != nil {
		return x.TxHash
	}
	return ""
}

type SendRawTransactionInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	To     string `protobuf:"bytes,2,opt,name=To,json=to,proto3" json:"To,omitempty"`
	Amount string `protobuf:"bytes,3,opt,name=Amount,json=account,proto3" json:"Amount,omitempty"`
}

func (x *SendRawTransactionInput) Reset() {
	*x = SendRawTransactionInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_polygon_transaction_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendRawTransactionInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendRawTransactionInput) ProtoMessage() {}

func (x *SendRawTransactionInput) ProtoReflect() protoreflect.Message {
	mi := &file_proto_polygon_transaction_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendRawTransactionInput.ProtoReflect.Descriptor instead.
func (*SendRawTransactionInput) Descriptor() ([]byte, []int) {
	return file_proto_polygon_transaction_message_proto_rawDescGZIP(), []int{2}
}

func (x *SendRawTransactionInput) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *SendRawTransactionInput) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

type SendRawTransactionOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=Success,json=success,proto3" json:"Success,omitempty"`
	TxHash  string `protobuf:"bytes,2,opt,name=TxHash,json=tx_hash,proto3" json:"TxHash,omitempty"`
}

func (x *SendRawTransactionOutput) Reset() {
	*x = SendRawTransactionOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_polygon_transaction_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendRawTransactionOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendRawTransactionOutput) ProtoMessage() {}

func (x *SendRawTransactionOutput) ProtoReflect() protoreflect.Message {
	mi := &file_proto_polygon_transaction_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendRawTransactionOutput.ProtoReflect.Descriptor instead.
func (*SendRawTransactionOutput) Descriptor() ([]byte, []int) {
	return file_proto_polygon_transaction_message_proto_rawDescGZIP(), []int{3}
}

func (x *SendRawTransactionOutput) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *SendRawTransactionOutput) GetTxHash() string {
	if x != nil {
		return x.TxHash
	}
	return ""
}

type GetReceiptInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TxHash string `protobuf:"bytes,1,opt,name=TxHash,json=tx_hash,proto3" json:"TxHash,omitempty"`
}

func (x *GetReceiptInput) Reset() {
	*x = GetReceiptInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_polygon_transaction_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReceiptInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReceiptInput) ProtoMessage() {}

func (x *GetReceiptInput) ProtoReflect() protoreflect.Message {
	mi := &file_proto_polygon_transaction_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReceiptInput.ProtoReflect.Descriptor instead.
func (*GetReceiptInput) Descriptor() ([]byte, []int) {
	return file_proto_polygon_transaction_message_proto_rawDescGZIP(), []int{4}
}

func (x *GetReceiptInput) GetTxHash() string {
	if x != nil {
		return x.TxHash
	}
	return ""
}

type GetReceiptOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TxHash           string `protobuf:"bytes,1,opt,name=TxHash,json=tx_hash,proto3" json:"TxHash,omitempty"`
	Status           uint64 `protobuf:"varint,2,opt,name=Status,json=status,proto3" json:"Status,omitempty"`
	BlockHash        string `protobuf:"bytes,3,opt,name=BlockHash,json=block_hash,proto3" json:"BlockHash,omitempty"`
	BlockNumber      string `protobuf:"bytes,4,opt,name=BlockNumber,json=block_number,proto3" json:"BlockNumber,omitempty"`
	GasUsed          uint64 `protobuf:"varint,5,opt,name=GasUsed,json=gas_used,proto3" json:"GasUsed,omitempty"`
	TransactionIndex uint64 `protobuf:"varint,6,opt,name=TransactionIndex,json=transaction_index,proto3" json:"TransactionIndex,omitempty"`
}

func (x *GetReceiptOutput) Reset() {
	*x = GetReceiptOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_polygon_transaction_message_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReceiptOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReceiptOutput) ProtoMessage() {}

func (x *GetReceiptOutput) ProtoReflect() protoreflect.Message {
	mi := &file_proto_polygon_transaction_message_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReceiptOutput.ProtoReflect.Descriptor instead.
func (*GetReceiptOutput) Descriptor() ([]byte, []int) {
	return file_proto_polygon_transaction_message_proto_rawDescGZIP(), []int{5}
}

func (x *GetReceiptOutput) GetTxHash() string {
	if x != nil {
		return x.TxHash
	}
	return ""
}

func (x *GetReceiptOutput) GetStatus() uint64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *GetReceiptOutput) GetBlockHash() string {
	if x != nil {
		return x.BlockHash
	}
	return ""
}

func (x *GetReceiptOutput) GetBlockNumber() string {
	if x != nil {
		return x.BlockNumber
	}
	return ""
}

func (x *GetReceiptOutput) GetGasUsed() uint64 {
	if x != nil {
		return x.GasUsed
	}
	return 0
}

func (x *GetReceiptOutput) GetTransactionIndex() uint64 {
	if x != nil {
		return x.TransactionIndex
	}
	return 0
}

type GetBlockNumberInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetBlockNumberInput) Reset() {
	*x = GetBlockNumberInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_polygon_transaction_message_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlockNumberInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlockNumberInput) ProtoMessage() {}

func (x *GetBlockNumberInput) ProtoReflect() protoreflect.Message {
	mi := &file_proto_polygon_transaction_message_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlockNumberInput.ProtoReflect.Descriptor instead.
func (*GetBlockNumberInput) Descriptor() ([]byte, []int) {
	return file_proto_polygon_transaction_message_proto_rawDescGZIP(), []int{6}
}

type GetBlockNumberOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockNumber uint64 `protobuf:"varint,1,opt,name=BlockNumber,json=block_number,proto3" json:"BlockNumber,omitempty"`
}

func (x *GetBlockNumberOutput) Reset() {
	*x = GetBlockNumberOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_polygon_transaction_message_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlockNumberOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlockNumberOutput) ProtoMessage() {}

func (x *GetBlockNumberOutput) ProtoReflect() protoreflect.Message {
	mi := &file_proto_polygon_transaction_message_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlockNumberOutput.ProtoReflect.Descriptor instead.
func (*GetBlockNumberOutput) Descriptor() ([]byte, []int) {
	return file_proto_polygon_transaction_message_proto_rawDescGZIP(), []int{7}
}

func (x *GetBlockNumberOutput) GetBlockNumber() uint64 {
	if x != nil {
		return x.BlockNumber
	}
	return 0
}

type GetBlocksInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Start uint64 `protobuf:"varint,1,opt,name=Start,json=start,proto3" json:"Start,omitempty"`
	End   uint64 `protobuf:"varint,2,opt,name=End,json=end,proto3" json:"End,omitempty"`
}

func (x *GetBlocksInput) Reset() {
	*x = GetBlocksInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_polygon_transaction_message_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlocksInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlocksInput) ProtoMessage() {}

func (x *GetBlocksInput) ProtoReflect() protoreflect.Message {
	mi := &file_proto_polygon_transaction_message_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlocksInput.ProtoReflect.Descriptor instead.
func (*GetBlocksInput) Descriptor() ([]byte, []int) {
	return file_proto_polygon_transaction_message_proto_rawDescGZIP(), []int{8}
}

func (x *GetBlocksInput) GetStart() uint64 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *GetBlocksInput) GetEnd() uint64 {
	if x != nil {
		return x.End
	}
	return 0
}

type BlockData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	To    string `protobuf:"bytes,1,opt,name=To,json=to,proto3" json:"To,omitempty"`
	Hash  string `protobuf:"bytes,2,opt,name=Hash,json=hash,proto3" json:"Hash,omitempty"`
	Value string `protobuf:"bytes,3,opt,name=Value,json=value,proto3" json:"Value,omitempty"`
}

func (x *BlockData) Reset() {
	*x = BlockData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_polygon_transaction_message_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockData) ProtoMessage() {}

func (x *BlockData) ProtoReflect() protoreflect.Message {
	mi := &file_proto_polygon_transaction_message_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockData.ProtoReflect.Descriptor instead.
func (*BlockData) Descriptor() ([]byte, []int) {
	return file_proto_polygon_transaction_message_proto_rawDescGZIP(), []int{9}
}

func (x *BlockData) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *BlockData) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *BlockData) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type GetBlocksOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result []*BlockData `protobuf:"bytes,1,rep,name=Result,json=result,proto3" json:"Result,omitempty"`
}

func (x *GetBlocksOutput) Reset() {
	*x = GetBlocksOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_polygon_transaction_message_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlocksOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlocksOutput) ProtoMessage() {}

func (x *GetBlocksOutput) ProtoReflect() protoreflect.Message {
	mi := &file_proto_polygon_transaction_message_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlocksOutput.ProtoReflect.Descriptor instead.
func (*GetBlocksOutput) Descriptor() ([]byte, []int) {
	return file_proto_polygon_transaction_message_proto_rawDescGZIP(), []int{10}
}

func (x *GetBlocksOutput) GetResult() []*BlockData {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_proto_polygon_transaction_message_proto protoreflect.FileDescriptor

var file_proto_polygon_transaction_message_proto_rawDesc = []byte{
	0x0a, 0x27, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x2f,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x6f, 0x6c, 0x79, 0x67,
	0x6f, 0x6e, 0x22, 0x57, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x61, 0x77, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x46, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66,
	0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x54, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x74, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x35, 0x0a, 0x1a, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x61, 0x77, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x17, 0x0a, 0x06, 0x54, 0x78, 0x48,
	0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x78, 0x5f, 0x68, 0x61,
	0x73, 0x68, 0x22, 0x42, 0x0a, 0x17, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x61, 0x77, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x54, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x17, 0x0a,
	0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x4d, 0x0a, 0x18, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x61,
	0x77, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x17, 0x0a, 0x06,
	0x54, 0x78, 0x48, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x78,
	0x5f, 0x68, 0x61, 0x73, 0x68, 0x22, 0x2a, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x65,
	0x69, 0x70, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x17, 0x0a, 0x06, 0x54, 0x78, 0x48, 0x61,
	0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x78, 0x5f, 0x68, 0x61, 0x73,
	0x68, 0x22, 0xcd, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x17, 0x0a, 0x06, 0x54, 0x78, 0x48, 0x61, 0x73, 0x68,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x78, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x12,
	0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x09, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x48, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x12, 0x21, 0x0a, 0x0b, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x07, 0x47, 0x61, 0x73,
	0x55, 0x73, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x67, 0x61, 0x73, 0x5f,
	0x75, 0x73, 0x65, 0x64, 0x12, 0x2b, 0x0a, 0x10, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x11,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x22, 0x15, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x22, 0x39, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x12, 0x21, 0x0a, 0x0b, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x22, 0x38, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73,
	0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x45,
	0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x22, 0x45, 0x0a,
	0x09, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x54, 0x6f,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x48, 0x61,
	0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x14,
	0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0x3d, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x2a, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x6f, 0x6c, 0x79, 0x67, 0x6f,
	0x6e, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x52, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x42, 0x0a, 0x5a, 0x08, 0x2f, 0x70, 0x6f, 0x6c, 0x79, 0x67, 0x6f, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_polygon_transaction_message_proto_rawDescOnce sync.Once
	file_proto_polygon_transaction_message_proto_rawDescData = file_proto_polygon_transaction_message_proto_rawDesc
)

func file_proto_polygon_transaction_message_proto_rawDescGZIP() []byte {
	file_proto_polygon_transaction_message_proto_rawDescOnce.Do(func() {
		file_proto_polygon_transaction_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_polygon_transaction_message_proto_rawDescData)
	})
	return file_proto_polygon_transaction_message_proto_rawDescData
}

var file_proto_polygon_transaction_message_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_polygon_transaction_message_proto_goTypes = []interface{}{
	(*CreateRawTransactionInput)(nil),  // 0: polygon.CreateRawTransactionInput
	(*CreateRawTransactionOutput)(nil), // 1: polygon.CreateRawTransactionOutput
	(*SendRawTransactionInput)(nil),    // 2: polygon.SendRawTransactionInput
	(*SendRawTransactionOutput)(nil),   // 3: polygon.SendRawTransactionOutput
	(*GetReceiptInput)(nil),            // 4: polygon.GetReceiptInput
	(*GetReceiptOutput)(nil),           // 5: polygon.GetReceiptOutput
	(*GetBlockNumberInput)(nil),        // 6: polygon.GetBlockNumberInput
	(*GetBlockNumberOutput)(nil),       // 7: polygon.GetBlockNumberOutput
	(*GetBlocksInput)(nil),             // 8: polygon.GetBlocksInput
	(*BlockData)(nil),                  // 9: polygon.BlockData
	(*GetBlocksOutput)(nil),            // 10: polygon.GetBlocksOutput
}
var file_proto_polygon_transaction_message_proto_depIdxs = []int32{
	9, // 0: polygon.GetBlocksOutput.Result:type_name -> polygon.BlockData
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_polygon_transaction_message_proto_init() }
func file_proto_polygon_transaction_message_proto_init() {
	if File_proto_polygon_transaction_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_polygon_transaction_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRawTransactionInput); i {
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
		file_proto_polygon_transaction_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRawTransactionOutput); i {
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
		file_proto_polygon_transaction_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendRawTransactionInput); i {
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
		file_proto_polygon_transaction_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendRawTransactionOutput); i {
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
		file_proto_polygon_transaction_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReceiptInput); i {
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
		file_proto_polygon_transaction_message_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReceiptOutput); i {
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
		file_proto_polygon_transaction_message_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBlockNumberInput); i {
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
		file_proto_polygon_transaction_message_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBlockNumberOutput); i {
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
		file_proto_polygon_transaction_message_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBlocksInput); i {
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
		file_proto_polygon_transaction_message_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockData); i {
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
		file_proto_polygon_transaction_message_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBlocksOutput); i {
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
			RawDescriptor: file_proto_polygon_transaction_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_polygon_transaction_message_proto_goTypes,
		DependencyIndexes: file_proto_polygon_transaction_message_proto_depIdxs,
		MessageInfos:      file_proto_polygon_transaction_message_proto_msgTypes,
	}.Build()
	File_proto_polygon_transaction_message_proto = out.File
	file_proto_polygon_transaction_message_proto_rawDesc = nil
	file_proto_polygon_transaction_message_proto_goTypes = nil
	file_proto_polygon_transaction_message_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.3
// source: proto/ethereum/transaction.message.proto

package ethereum

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
		mi := &file_proto_ethereum_transaction_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendRawTransactionInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendRawTransactionInput) ProtoMessage() {}

func (x *SendRawTransactionInput) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ethereum_transaction_message_proto_msgTypes[0]
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
	return file_proto_ethereum_transaction_message_proto_rawDescGZIP(), []int{0}
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
		mi := &file_proto_ethereum_transaction_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendRawTransactionOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendRawTransactionOutput) ProtoMessage() {}

func (x *SendRawTransactionOutput) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ethereum_transaction_message_proto_msgTypes[1]
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
	return file_proto_ethereum_transaction_message_proto_rawDescGZIP(), []int{1}
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

type TransferInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID string `protobuf:"bytes,1,opt,name=UserID,json=user_id,proto3" json:"UserID,omitempty"`
	From   string `protobuf:"bytes,2,opt,name=From,json=from,proto3" json:"From,omitempty"`
	Amount string `protobuf:"bytes,3,opt,name=Amount,json=success,proto3" json:"Amount,omitempty"`
}

func (x *TransferInput) Reset() {
	*x = TransferInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ethereum_transaction_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransferInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferInput) ProtoMessage() {}

func (x *TransferInput) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ethereum_transaction_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferInput.ProtoReflect.Descriptor instead.
func (*TransferInput) Descriptor() ([]byte, []int) {
	return file_proto_ethereum_transaction_message_proto_rawDescGZIP(), []int{2}
}

func (x *TransferInput) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *TransferInput) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *TransferInput) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

type TransferOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=Success,json=success,proto3" json:"Success,omitempty"`
	TxHash  string `protobuf:"bytes,2,opt,name=TxHash,json=tx_hash,proto3" json:"TxHash,omitempty"`
}

func (x *TransferOutput) Reset() {
	*x = TransferOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ethereum_transaction_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransferOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferOutput) ProtoMessage() {}

func (x *TransferOutput) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ethereum_transaction_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferOutput.ProtoReflect.Descriptor instead.
func (*TransferOutput) Descriptor() ([]byte, []int) {
	return file_proto_ethereum_transaction_message_proto_rawDescGZIP(), []int{3}
}

func (x *TransferOutput) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *TransferOutput) GetTxHash() string {
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
		mi := &file_proto_ethereum_transaction_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReceiptInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReceiptInput) ProtoMessage() {}

func (x *GetReceiptInput) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ethereum_transaction_message_proto_msgTypes[4]
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
	return file_proto_ethereum_transaction_message_proto_rawDescGZIP(), []int{4}
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
		mi := &file_proto_ethereum_transaction_message_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReceiptOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReceiptOutput) ProtoMessage() {}

func (x *GetReceiptOutput) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ethereum_transaction_message_proto_msgTypes[5]
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
	return file_proto_ethereum_transaction_message_proto_rawDescGZIP(), []int{5}
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
		mi := &file_proto_ethereum_transaction_message_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlockNumberInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlockNumberInput) ProtoMessage() {}

func (x *GetBlockNumberInput) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ethereum_transaction_message_proto_msgTypes[6]
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
	return file_proto_ethereum_transaction_message_proto_rawDescGZIP(), []int{6}
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
		mi := &file_proto_ethereum_transaction_message_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlockNumberOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlockNumberOutput) ProtoMessage() {}

func (x *GetBlockNumberOutput) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ethereum_transaction_message_proto_msgTypes[7]
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
	return file_proto_ethereum_transaction_message_proto_rawDescGZIP(), []int{7}
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
		mi := &file_proto_ethereum_transaction_message_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlocksInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlocksInput) ProtoMessage() {}

func (x *GetBlocksInput) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ethereum_transaction_message_proto_msgTypes[8]
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
	return file_proto_ethereum_transaction_message_proto_rawDescGZIP(), []int{8}
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

type GetBlocksOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result [][]byte `protobuf:"bytes,1,rep,name=Result,json=result,proto3" json:"Result,omitempty"`
}

func (x *GetBlocksOutput) Reset() {
	*x = GetBlocksOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_ethereum_transaction_message_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlocksOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlocksOutput) ProtoMessage() {}

func (x *GetBlocksOutput) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ethereum_transaction_message_proto_msgTypes[9]
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
	return file_proto_ethereum_transaction_message_proto_rawDescGZIP(), []int{9}
}

func (x *GetBlocksOutput) GetResult() [][]byte {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_proto_ethereum_transaction_message_proto protoreflect.FileDescriptor

var file_proto_ethereum_transaction_message_proto_rawDesc = []byte{
	0x0a, 0x28, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d,
	0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x65, 0x74, 0x68, 0x65,
	0x72, 0x65, 0x75, 0x6d, 0x22, 0x42, 0x0a, 0x17, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x61, 0x77, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x54, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12,
	0x17, 0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x4d, 0x0a, 0x18, 0x53, 0x65, 0x6e, 0x64,
	0x52, 0x61, 0x77, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x17,
	0x0a, 0x06, 0x54, 0x78, 0x48, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x74, 0x78, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x22, 0x55, 0x0a, 0x0d, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x65, 0x72, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x17, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x46, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x17, 0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x43,
	0x0a, 0x0e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x17, 0x0a, 0x06, 0x54, 0x78,
	0x48, 0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x78, 0x5f, 0x68,
	0x61, 0x73, 0x68, 0x22, 0x2a, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70,
	0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x17, 0x0a, 0x06, 0x54, 0x78, 0x48, 0x61, 0x73, 0x68,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x78, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x22,
	0xcd, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x12, 0x17, 0x0a, 0x06, 0x54, 0x78, 0x48, 0x61, 0x73, 0x68, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x78, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x12, 0x16, 0x0a,
	0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x09, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61,
	0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f,
	0x68, 0x61, 0x73, 0x68, 0x12, 0x21, 0x0a, 0x0b, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x07, 0x47, 0x61, 0x73, 0x55, 0x73,
	0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x67, 0x61, 0x73, 0x5f, 0x75, 0x73,
	0x65, 0x64, 0x12, 0x2b, 0x0a, 0x10, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x11, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x22,
	0x15, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x22, 0x39, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x21,
	0x0a, 0x0b, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x22, 0x38, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x45, 0x6e, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x65, 0x6e, 0x64, 0x22, 0x29, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x65, 0x74, 0x68, 0x65,
	0x72, 0x65, 0x75, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_ethereum_transaction_message_proto_rawDescOnce sync.Once
	file_proto_ethereum_transaction_message_proto_rawDescData = file_proto_ethereum_transaction_message_proto_rawDesc
)

func file_proto_ethereum_transaction_message_proto_rawDescGZIP() []byte {
	file_proto_ethereum_transaction_message_proto_rawDescOnce.Do(func() {
		file_proto_ethereum_transaction_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_ethereum_transaction_message_proto_rawDescData)
	})
	return file_proto_ethereum_transaction_message_proto_rawDescData
}

var file_proto_ethereum_transaction_message_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_ethereum_transaction_message_proto_goTypes = []interface{}{
	(*SendRawTransactionInput)(nil),  // 0: ethereum.SendRawTransactionInput
	(*SendRawTransactionOutput)(nil), // 1: ethereum.SendRawTransactionOutput
	(*TransferInput)(nil),            // 2: ethereum.TransferInput
	(*TransferOutput)(nil),           // 3: ethereum.TransferOutput
	(*GetReceiptInput)(nil),          // 4: ethereum.GetReceiptInput
	(*GetReceiptOutput)(nil),         // 5: ethereum.GetReceiptOutput
	(*GetBlockNumberInput)(nil),      // 6: ethereum.GetBlockNumberInput
	(*GetBlockNumberOutput)(nil),     // 7: ethereum.GetBlockNumberOutput
	(*GetBlocksInput)(nil),           // 8: ethereum.GetBlocksInput
	(*GetBlocksOutput)(nil),          // 9: ethereum.GetBlocksOutput
}
var file_proto_ethereum_transaction_message_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_ethereum_transaction_message_proto_init() }
func file_proto_ethereum_transaction_message_proto_init() {
	if File_proto_ethereum_transaction_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_ethereum_transaction_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_ethereum_transaction_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_ethereum_transaction_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransferInput); i {
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
		file_proto_ethereum_transaction_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransferOutput); i {
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
		file_proto_ethereum_transaction_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_ethereum_transaction_message_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_ethereum_transaction_message_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_ethereum_transaction_message_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_ethereum_transaction_message_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_ethereum_transaction_message_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_proto_ethereum_transaction_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_ethereum_transaction_message_proto_goTypes,
		DependencyIndexes: file_proto_ethereum_transaction_message_proto_depIdxs,
		MessageInfos:      file_proto_ethereum_transaction_message_proto_msgTypes,
	}.Build()
	File_proto_ethereum_transaction_message_proto = out.File
	file_proto_ethereum_transaction_message_proto_rawDesc = nil
	file_proto_ethereum_transaction_message_proto_goTypes = nil
	file_proto_ethereum_transaction_message_proto_depIdxs = nil
}

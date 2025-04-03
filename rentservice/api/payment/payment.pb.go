// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v6.30.1
// source: payment.proto

package payment

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DepositRequest struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	UserId        int64                   `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Amount        *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DepositRequest) Reset() {
	*x = DepositRequest{}
	mi := &file_payment_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DepositRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DepositRequest) ProtoMessage() {}

func (x *DepositRequest) ProtoReflect() protoreflect.Message {
	mi := &file_payment_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DepositRequest.ProtoReflect.Descriptor instead.
func (*DepositRequest) Descriptor() ([]byte, []int) {
	return file_payment_proto_rawDescGZIP(), []int{0}
}

func (x *DepositRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *DepositRequest) GetAmount() *wrapperspb.StringValue {
	if x != nil {
		return x.Amount
	}
	return nil
}

type DepositResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DepositResponse) Reset() {
	*x = DepositResponse{}
	mi := &file_payment_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DepositResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DepositResponse) ProtoMessage() {}

func (x *DepositResponse) ProtoReflect() protoreflect.Message {
	mi := &file_payment_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DepositResponse.ProtoReflect.Descriptor instead.
func (*DepositResponse) Descriptor() ([]byte, []int) {
	return file_payment_proto_rawDescGZIP(), []int{1}
}

func (x *DepositResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type WithDrawRequest struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	UserId        int64                   `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Amount        *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WithDrawRequest) Reset() {
	*x = WithDrawRequest{}
	mi := &file_payment_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WithDrawRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WithDrawRequest) ProtoMessage() {}

func (x *WithDrawRequest) ProtoReflect() protoreflect.Message {
	mi := &file_payment_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WithDrawRequest.ProtoReflect.Descriptor instead.
func (*WithDrawRequest) Descriptor() ([]byte, []int) {
	return file_payment_proto_rawDescGZIP(), []int{2}
}

func (x *WithDrawRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *WithDrawRequest) GetAmount() *wrapperspb.StringValue {
	if x != nil {
		return x.Amount
	}
	return nil
}

type WithDrawResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WithDrawResponse) Reset() {
	*x = WithDrawResponse{}
	mi := &file_payment_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WithDrawResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WithDrawResponse) ProtoMessage() {}

func (x *WithDrawResponse) ProtoReflect() protoreflect.Message {
	mi := &file_payment_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WithDrawResponse.ProtoReflect.Descriptor instead.
func (*WithDrawResponse) Descriptor() ([]byte, []int) {
	return file_payment_proto_rawDescGZIP(), []int{3}
}

func (x *WithDrawResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type HoldRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RenterId      int64                  `protobuf:"varint,1,opt,name=renter_id,json=renterId,proto3" json:"renter_id,omitempty"`
	RentAmount    string                 `protobuf:"bytes,2,opt,name=rent_amount,json=rentAmount,proto3" json:"rent_amount,omitempty"`
	PledgeAmount  string                 `protobuf:"bytes,3,opt,name=pledge_amount,json=pledgeAmount,proto3" json:"pledge_amount,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HoldRequest) Reset() {
	*x = HoldRequest{}
	mi := &file_payment_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HoldRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HoldRequest) ProtoMessage() {}

func (x *HoldRequest) ProtoReflect() protoreflect.Message {
	mi := &file_payment_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HoldRequest.ProtoReflect.Descriptor instead.
func (*HoldRequest) Descriptor() ([]byte, []int) {
	return file_payment_proto_rawDescGZIP(), []int{4}
}

func (x *HoldRequest) GetRenterId() int64 {
	if x != nil {
		return x.RenterId
	}
	return 0
}

func (x *HoldRequest) GetRentAmount() string {
	if x != nil {
		return x.RentAmount
	}
	return ""
}

func (x *HoldRequest) GetPledgeAmount() string {
	if x != nil {
		return x.PledgeAmount
	}
	return ""
}

type HoldResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	HeldFundsID   int64                  `protobuf:"varint,2,opt,name=heldFundsID,proto3" json:"heldFundsID,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HoldResponse) Reset() {
	*x = HoldResponse{}
	mi := &file_payment_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HoldResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HoldResponse) ProtoMessage() {}

func (x *HoldResponse) ProtoReflect() protoreflect.Message {
	mi := &file_payment_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HoldResponse.ProtoReflect.Descriptor instead.
func (*HoldResponse) Descriptor() ([]byte, []int) {
	return file_payment_proto_rawDescGZIP(), []int{5}
}

func (x *HoldResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *HoldResponse) GetHeldFundsID() int64 {
	if x != nil {
		return x.HeldFundsID
	}
	return 0
}

type CompleteRentRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RenterId      int64                  `protobuf:"varint,1,opt,name=renter_id,json=renterId,proto3" json:"renter_id,omitempty"`
	LandlordId    int64                  `protobuf:"varint,2,opt,name=landlord_id,json=landlordId,proto3" json:"landlord_id,omitempty"`
	HeldFundsID   int64                  `protobuf:"varint,3,opt,name=heldFundsID,proto3" json:"heldFundsID,omitempty"`
	RentAmount    string                 `protobuf:"bytes,4,opt,name=rent_amount,json=rentAmount,proto3" json:"rent_amount,omitempty"`
	PledgeAmount  string                 `protobuf:"bytes,5,opt,name=pledge_amount,json=pledgeAmount,proto3" json:"pledge_amount,omitempty"`
	ToLandlord    bool                   `protobuf:"varint,6,opt,name=to_landlord,json=toLandlord,proto3" json:"to_landlord,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CompleteRentRequest) Reset() {
	*x = CompleteRentRequest{}
	mi := &file_payment_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CompleteRentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompleteRentRequest) ProtoMessage() {}

func (x *CompleteRentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_payment_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompleteRentRequest.ProtoReflect.Descriptor instead.
func (*CompleteRentRequest) Descriptor() ([]byte, []int) {
	return file_payment_proto_rawDescGZIP(), []int{6}
}

func (x *CompleteRentRequest) GetRenterId() int64 {
	if x != nil {
		return x.RenterId
	}
	return 0
}

func (x *CompleteRentRequest) GetLandlordId() int64 {
	if x != nil {
		return x.LandlordId
	}
	return 0
}

func (x *CompleteRentRequest) GetHeldFundsID() int64 {
	if x != nil {
		return x.HeldFundsID
	}
	return 0
}

func (x *CompleteRentRequest) GetRentAmount() string {
	if x != nil {
		return x.RentAmount
	}
	return ""
}

func (x *CompleteRentRequest) GetPledgeAmount() string {
	if x != nil {
		return x.PledgeAmount
	}
	return ""
}

func (x *CompleteRentRequest) GetToLandlord() bool {
	if x != nil {
		return x.ToLandlord
	}
	return false
}

type CompleteRentResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CompleteRentResponse) Reset() {
	*x = CompleteRentResponse{}
	mi := &file_payment_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CompleteRentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompleteRentResponse) ProtoMessage() {}

func (x *CompleteRentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_payment_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompleteRentResponse.ProtoReflect.Descriptor instead.
func (*CompleteRentResponse) Descriptor() ([]byte, []int) {
	return file_payment_proto_rawDescGZIP(), []int{7}
}

func (x *CompleteRentResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_payment_proto protoreflect.FileDescriptor

var file_payment_proto_rawDesc = string([]byte{
	0x0a, 0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65,
	0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5f, 0x0a, 0x0e, 0x44, 0x65, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x34, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x2b, 0x0a, 0x0f, 0x44, 0x65, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x60, 0x0a, 0x0f, 0x57, 0x69, 0x74, 0x68, 0x44, 0x72,
	0x61, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x34, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x2c, 0x0a, 0x10, 0x57, 0x69, 0x74, 0x68,
	0x44, 0x72, 0x61, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x70, 0x0a, 0x0b, 0x48, 0x6f, 0x6c, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x72, 0x65, 0x6e, 0x74, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x6e, 0x74, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x5f, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x6c, 0x65, 0x64,
	0x67, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x4a, 0x0a, 0x0c, 0x48, 0x6f, 0x6c, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x68, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6e, 0x64, 0x73, 0x49,
	0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x68, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6e,
	0x64, 0x73, 0x49, 0x44, 0x22, 0xdc, 0x01, 0x0a, 0x13, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x72, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x72, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x61, 0x6e,
	0x64, 0x6c, 0x6f, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x6c, 0x61, 0x6e, 0x64, 0x6c, 0x6f, 0x72, 0x64, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x68, 0x65,
	0x6c, 0x64, 0x46, 0x75, 0x6e, 0x64, 0x73, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0b, 0x68, 0x65, 0x6c, 0x64, 0x46, 0x75, 0x6e, 0x64, 0x73, 0x49, 0x44, 0x12, 0x1f, 0x0a, 0x0b,
	0x72, 0x65, 0x6e, 0x74, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x72, 0x65, 0x6e, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a,
	0x0d, 0x70, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x5f, 0x6c, 0x61, 0x6e, 0x64, 0x6c, 0x6f, 0x72,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x74, 0x6f, 0x4c, 0x61, 0x6e, 0x64, 0x6c,
	0x6f, 0x72, 0x64, 0x22, 0x30, 0x0a, 0x14, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0x91, 0x02, 0x0a, 0x0e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x07, 0x44, 0x65, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x12, 0x17, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x44, 0x65,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x08, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72,
	0x61, 0x77, 0x12, 0x18, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x57, 0x69, 0x74,
	0x68, 0x44, 0x72, 0x61, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x57, 0x69, 0x74, 0x68, 0x44, 0x72, 0x61, 0x77, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x04, 0x48, 0x6f, 0x6c, 0x64, 0x12,
	0x14, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x48, 0x6f, 0x6c, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x48, 0x6f, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x0c,
	0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x2e, 0x70,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_payment_proto_rawDescOnce sync.Once
	file_payment_proto_rawDescData []byte
)

func file_payment_proto_rawDescGZIP() []byte {
	file_payment_proto_rawDescOnce.Do(func() {
		file_payment_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_payment_proto_rawDesc), len(file_payment_proto_rawDesc)))
	})
	return file_payment_proto_rawDescData
}

var file_payment_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_payment_proto_goTypes = []any{
	(*DepositRequest)(nil),         // 0: payment.DepositRequest
	(*DepositResponse)(nil),        // 1: payment.DepositResponse
	(*WithDrawRequest)(nil),        // 2: payment.WithDrawRequest
	(*WithDrawResponse)(nil),       // 3: payment.WithDrawResponse
	(*HoldRequest)(nil),            // 4: payment.HoldRequest
	(*HoldResponse)(nil),           // 5: payment.HoldResponse
	(*CompleteRentRequest)(nil),    // 6: payment.CompleteRentRequest
	(*CompleteRentResponse)(nil),   // 7: payment.CompleteRentResponse
	(*wrapperspb.StringValue)(nil), // 8: google.protobuf.StringValue
}
var file_payment_proto_depIdxs = []int32{
	8, // 0: payment.DepositRequest.amount:type_name -> google.protobuf.StringValue
	8, // 1: payment.WithDrawRequest.amount:type_name -> google.protobuf.StringValue
	0, // 2: payment.PaymentService.Deposit:input_type -> payment.DepositRequest
	2, // 3: payment.PaymentService.Withdraw:input_type -> payment.WithDrawRequest
	4, // 4: payment.PaymentService.Hold:input_type -> payment.HoldRequest
	6, // 5: payment.PaymentService.CompleteRent:input_type -> payment.CompleteRentRequest
	1, // 6: payment.PaymentService.Deposit:output_type -> payment.DepositResponse
	3, // 7: payment.PaymentService.Withdraw:output_type -> payment.WithDrawResponse
	5, // 8: payment.PaymentService.Hold:output_type -> payment.HoldResponse
	7, // 9: payment.PaymentService.CompleteRent:output_type -> payment.CompleteRentResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_payment_proto_init() }
func file_payment_proto_init() {
	if File_payment_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_payment_proto_rawDesc), len(file_payment_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_payment_proto_goTypes,
		DependencyIndexes: file_payment_proto_depIdxs,
		MessageInfos:      file_payment_proto_msgTypes,
	}.Build()
	File_payment_proto = out.File
	file_payment_proto_goTypes = nil
	file_payment_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: pb/user.proto

package __

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

type SignupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *SignupRequest) Reset() {
	*x = SignupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignupRequest) ProtoMessage() {}

func (x *SignupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignupRequest.ProtoReflect.Descriptor instead.
func (*SignupRequest) Descriptor() ([]byte, []int) {
	return file_pb_user_proto_rawDescGZIP(), []int{0}
}

func (x *SignupRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type SignupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SignupResponse) Reset() {
	*x = SignupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignupResponse) ProtoMessage() {}

func (x *SignupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignupResponse.ProtoReflect.Descriptor instead.
func (*SignupResponse) Descriptor() ([]byte, []int) {
	return file_pb_user_proto_rawDescGZIP(), []int{1}
}

func (x *SignupResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type VerifyOTPRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Otp   string `protobuf:"bytes,2,opt,name=otp,proto3" json:"otp,omitempty"`
}

func (x *VerifyOTPRequest) Reset() {
	*x = VerifyOTPRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyOTPRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyOTPRequest) ProtoMessage() {}

func (x *VerifyOTPRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyOTPRequest.ProtoReflect.Descriptor instead.
func (*VerifyOTPRequest) Descriptor() ([]byte, []int) {
	return file_pb_user_proto_rawDescGZIP(), []int{2}
}

func (x *VerifyOTPRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *VerifyOTPRequest) GetOtp() string {
	if x != nil {
		return x.Otp
	}
	return ""
}

type VerifyOTPResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Token   string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *VerifyOTPResponse) Reset() {
	*x = VerifyOTPResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyOTPResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyOTPResponse) ProtoMessage() {}

func (x *VerifyOTPResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyOTPResponse.ProtoReflect.Descriptor instead.
func (*VerifyOTPResponse) Descriptor() ([]byte, []int) {
	return file_pb_user_proto_rawDescGZIP(), []int{3}
}

func (x *VerifyOTPResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *VerifyOTPResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_pb_user_proto_rawDescGZIP(), []int{4}
}

func (x *LoginRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Token string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_pb_user_proto_rawDescGZIP(), []int{5}
}

func (x *LoginResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *LoginResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type RNoParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RNoParam) Reset() {
	*x = RNoParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RNoParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RNoParam) ProtoMessage() {}

func (x *RNoParam) ProtoReflect() protoreflect.Message {
	mi := &file_pb_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RNoParam.ProtoReflect.Descriptor instead.
func (*RNoParam) Descriptor() ([]byte, []int) {
	return file_pb_user_proto_rawDescGZIP(), []int{6}
}

type ProductDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Category    string  `protobuf:"bytes,2,opt,name=category,proto3" json:"category,omitempty"`
	Name        string  `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Price       float64 `protobuf:"fixed64,4,opt,name=price,proto3" json:"price,omitempty"`
	Imagepath   string  `protobuf:"bytes,5,opt,name=imagepath,proto3" json:"imagepath,omitempty"`
	Description string  `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	Size        string  `protobuf:"bytes,7,opt,name=size,proto3" json:"size,omitempty"`
	Quantity    uint32  `protobuf:"varint,8,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *ProductDetails) Reset() {
	*x = ProductDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductDetails) ProtoMessage() {}

func (x *ProductDetails) ProtoReflect() protoreflect.Message {
	mi := &file_pb_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductDetails.ProtoReflect.Descriptor instead.
func (*ProductDetails) Descriptor() ([]byte, []int) {
	return file_pb_user_proto_rawDescGZIP(), []int{7}
}

func (x *ProductDetails) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ProductDetails) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *ProductDetails) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProductDetails) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *ProductDetails) GetImagepath() string {
	if x != nil {
		return x.Imagepath
	}
	return ""
}

func (x *ProductDetails) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ProductDetails) GetSize() string {
	if x != nil {
		return x.Size
	}
	return ""
}

func (x *ProductDetails) GetQuantity() uint32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

type ProductByName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *ProductByName) Reset() {
	*x = ProductByName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_user_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductByName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductByName) ProtoMessage() {}

func (x *ProductByName) ProtoReflect() protoreflect.Message {
	mi := &file_pb_user_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductByName.ProtoReflect.Descriptor instead.
func (*ProductByName) Descriptor() ([]byte, []int) {
	return file_pb_user_proto_rawDescGZIP(), []int{8}
}

func (x *ProductByName) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ProductID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ProductID) Reset() {
	*x = ProductID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_user_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductID) ProtoMessage() {}

func (x *ProductID) ProtoReflect() protoreflect.Message {
	mi := &file_pb_user_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductID.ProtoReflect.Descriptor instead.
func (*ProductID) Descriptor() ([]byte, []int) {
	return file_pb_user_proto_rawDescGZIP(), []int{9}
}

func (x *ProductID) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ProductList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item []*ProductDetails `protobuf:"bytes,1,rep,name=item,proto3" json:"item,omitempty"`
}

func (x *ProductList) Reset() {
	*x = ProductList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_user_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductList) ProtoMessage() {}

func (x *ProductList) ProtoReflect() protoreflect.Message {
	mi := &file_pb_user_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductList.ProtoReflect.Descriptor instead.
func (*ProductList) Descriptor() ([]byte, []int) {
	return file_pb_user_proto_rawDescGZIP(), []int{10}
}

func (x *ProductList) GetItem() []*ProductDetails {
	if x != nil {
		return x.Item
	}
	return nil
}

var File_pb_user_proto protoreflect.FileDescriptor

var file_pb_user_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x62, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x70, 0x62, 0x55, 0x22, 0x25, 0x0a, 0x0d, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x2a, 0x0a, 0x0e, 0x53,
	0x69, 0x67, 0x6e, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x3a, 0x0a, 0x10, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x4f, 0x54, 0x50, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x74, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6f, 0x74, 0x70, 0x22, 0x43, 0x0a, 0x11, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x4f, 0x54, 0x50,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x24, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x3b,
	0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x0a, 0x0a, 0x08, 0x52,
	0x4e, 0x6f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x22, 0xd6, 0x01, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x70, 0x61, 0x74, 0x68, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x70, 0x61, 0x74, 0x68, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x73, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x22, 0x23, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x1b, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x36, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x27, 0x0a, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x70, 0x62, 0x55, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x32, 0xda, 0x02, 0x0a, 0x0c, 0x55,
	0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x31, 0x0a, 0x06, 0x53,
	0x69, 0x67, 0x6e, 0x75, 0x70, 0x12, 0x12, 0x2e, 0x70, 0x62, 0x55, 0x2e, 0x53, 0x69, 0x67, 0x6e,
	0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x62, 0x55, 0x2e,
	0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a,
	0x0a, 0x09, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x4f, 0x54, 0x50, 0x12, 0x15, 0x2e, 0x70, 0x62,
	0x55, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x4f, 0x54, 0x50, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x62, 0x55, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x4f,
	0x54, 0x50, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x05, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x12, 0x11, 0x2e, 0x70, 0x62, 0x55, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x62, 0x55, 0x2e, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x0f, 0x55, 0x73,
	0x65, 0x72, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0d, 0x2e,
	0x70, 0x62, 0x55, 0x2e, 0x52, 0x4e, 0x6f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x13, 0x2e, 0x70,
	0x62, 0x55, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x12, 0x3c, 0x0a, 0x11, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x2e, 0x70, 0x62, 0x55, 0x2e, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x13, 0x2e, 0x70, 0x62, 0x55,
	0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12,
	0x36, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x42, 0x79,
	0x49, 0x44, 0x12, 0x0e, 0x2e, 0x70, 0x62, 0x55, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x49, 0x44, 0x1a, 0x13, 0x2e, 0x70, 0x62, 0x55, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_user_proto_rawDescOnce sync.Once
	file_pb_user_proto_rawDescData = file_pb_user_proto_rawDesc
)

func file_pb_user_proto_rawDescGZIP() []byte {
	file_pb_user_proto_rawDescOnce.Do(func() {
		file_pb_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_user_proto_rawDescData)
	})
	return file_pb_user_proto_rawDescData
}

var file_pb_user_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_pb_user_proto_goTypes = []any{
	(*SignupRequest)(nil),     // 0: pbU.SignupRequest
	(*SignupResponse)(nil),    // 1: pbU.SignupResponse
	(*VerifyOTPRequest)(nil),  // 2: pbU.VerifyOTPRequest
	(*VerifyOTPResponse)(nil), // 3: pbU.VerifyOTPResponse
	(*LoginRequest)(nil),      // 4: pbU.LoginRequest
	(*LoginResponse)(nil),     // 5: pbU.LoginResponse
	(*RNoParam)(nil),          // 6: pbU.RNoParam
	(*ProductDetails)(nil),    // 7: pbU.ProductDetails
	(*ProductByName)(nil),     // 8: pbU.ProductByName
	(*ProductID)(nil),         // 9: pbU.ProductID
	(*ProductList)(nil),       // 10: pbU.ProductList
}
var file_pb_user_proto_depIdxs = []int32{
	7, // 0: pbU.ProductList.item:type_name -> pbU.ProductDetails
	0, // 1: pbU.UserServices.Signup:input_type -> pbU.SignupRequest
	2, // 2: pbU.UserServices.VerifyOTP:input_type -> pbU.VerifyOTPRequest
	4, // 3: pbU.UserServices.Login:input_type -> pbU.LoginRequest
	6, // 4: pbU.UserServices.UserProductList:input_type -> pbU.RNoParam
	8, // 5: pbU.UserServices.UserProductByName:input_type -> pbU.ProductByName
	9, // 6: pbU.UserServices.UserProductByID:input_type -> pbU.ProductID
	1, // 7: pbU.UserServices.Signup:output_type -> pbU.SignupResponse
	3, // 8: pbU.UserServices.VerifyOTP:output_type -> pbU.VerifyOTPResponse
	5, // 9: pbU.UserServices.Login:output_type -> pbU.LoginResponse
	7, // 10: pbU.UserServices.UserProductList:output_type -> pbU.ProductDetails
	7, // 11: pbU.UserServices.UserProductByName:output_type -> pbU.ProductDetails
	7, // 12: pbU.UserServices.UserProductByID:output_type -> pbU.ProductDetails
	7, // [7:13] is the sub-list for method output_type
	1, // [1:7] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pb_user_proto_init() }
func file_pb_user_proto_init() {
	if File_pb_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_user_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*SignupRequest); i {
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
		file_pb_user_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*SignupResponse); i {
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
		file_pb_user_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*VerifyOTPRequest); i {
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
		file_pb_user_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*VerifyOTPResponse); i {
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
		file_pb_user_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*LoginRequest); i {
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
		file_pb_user_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*LoginResponse); i {
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
		file_pb_user_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*RNoParam); i {
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
		file_pb_user_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*ProductDetails); i {
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
		file_pb_user_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*ProductByName); i {
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
		file_pb_user_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*ProductID); i {
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
		file_pb_user_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*ProductList); i {
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
			RawDescriptor: file_pb_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_user_proto_goTypes,
		DependencyIndexes: file_pb_user_proto_depIdxs,
		MessageInfos:      file_pb_user_proto_msgTypes,
	}.Build()
	File_pb_user_proto = out.File
	file_pb_user_proto_rawDesc = nil
	file_pb_user_proto_goTypes = nil
	file_pb_user_proto_depIdxs = nil
}

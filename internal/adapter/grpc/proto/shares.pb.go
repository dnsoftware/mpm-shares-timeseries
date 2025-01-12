// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.12.4
// source: proto/shares.proto

package proto

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

type Share struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid         string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	ServiceId    string `protobuf:"bytes,2,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	CoinId       int64  `protobuf:"varint,3,opt,name=coin_id,json=coinId,proto3" json:"coin_id,omitempty"`
	WorkerId     int64  `protobuf:"varint,4,opt,name=worker_id,json=workerId,proto3" json:"worker_id,omitempty"`
	WalletId     int64  `protobuf:"varint,5,opt,name=wallet_id,json=walletId,proto3" json:"wallet_id,omitempty"`
	ShareDate    string `protobuf:"bytes,6,opt,name=share_date,json=shareDate,proto3" json:"share_date,omitempty"`
	Difficulty   string `protobuf:"bytes,7,opt,name=difficulty,proto3" json:"difficulty,omitempty"`
	ShareDif     string `protobuf:"bytes,8,opt,name=share_dif,json=shareDif,proto3" json:"share_dif,omitempty"`
	Nonce        string `protobuf:"bytes,9,opt,name=nonce,proto3" json:"nonce,omitempty"`
	IsSolo       bool   `protobuf:"varint,10,opt,name=is_solo,json=isSolo,proto3" json:"is_solo,omitempty"`
	RewardMethod string `protobuf:"bytes,11,opt,name=reward_method,json=rewardMethod,proto3" json:"reward_method,omitempty"`
	Cost         string `protobuf:"bytes,12,opt,name=cost,proto3" json:"cost,omitempty"`
}

func (x *Share) Reset() {
	*x = Share{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shares_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Share) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Share) ProtoMessage() {}

func (x *Share) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shares_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Share.ProtoReflect.Descriptor instead.
func (*Share) Descriptor() ([]byte, []int) {
	return file_proto_shares_proto_rawDescGZIP(), []int{0}
}

func (x *Share) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *Share) GetServiceId() string {
	if x != nil {
		return x.ServiceId
	}
	return ""
}

func (x *Share) GetCoinId() int64 {
	if x != nil {
		return x.CoinId
	}
	return 0
}

func (x *Share) GetWorkerId() int64 {
	if x != nil {
		return x.WorkerId
	}
	return 0
}

func (x *Share) GetWalletId() int64 {
	if x != nil {
		return x.WalletId
	}
	return 0
}

func (x *Share) GetShareDate() string {
	if x != nil {
		return x.ShareDate
	}
	return ""
}

func (x *Share) GetDifficulty() string {
	if x != nil {
		return x.Difficulty
	}
	return ""
}

func (x *Share) GetShareDif() string {
	if x != nil {
		return x.ShareDif
	}
	return ""
}

func (x *Share) GetNonce() string {
	if x != nil {
		return x.Nonce
	}
	return ""
}

func (x *Share) GetIsSolo() bool {
	if x != nil {
		return x.IsSolo
	}
	return false
}

func (x *Share) GetRewardMethod() string {
	if x != nil {
		return x.RewardMethod
	}
	return ""
}

func (x *Share) GetCost() string {
	if x != nil {
		return x.Cost
	}
	return ""
}

type AddSharesBatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Share []*Share `protobuf:"bytes,1,rep,name=share,proto3" json:"share,omitempty"`
}

func (x *AddSharesBatchRequest) Reset() {
	*x = AddSharesBatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shares_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddSharesBatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddSharesBatchRequest) ProtoMessage() {}

func (x *AddSharesBatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shares_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddSharesBatchRequest.ProtoReflect.Descriptor instead.
func (*AddSharesBatchRequest) Descriptor() ([]byte, []int) {
	return file_proto_shares_proto_rawDescGZIP(), []int{1}
}

func (x *AddSharesBatchRequest) GetShare() []*Share {
	if x != nil {
		return x.Share
	}
	return nil
}

type AddSharesBatchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AddedCount int64 `protobuf:"varint,1,opt,name=added_count,json=addedCount,proto3" json:"added_count,omitempty"`
}

func (x *AddSharesBatchResponse) Reset() {
	*x = AddSharesBatchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_shares_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddSharesBatchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddSharesBatchResponse) ProtoMessage() {}

func (x *AddSharesBatchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_shares_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddSharesBatchResponse.ProtoReflect.Descriptor instead.
func (*AddSharesBatchResponse) Descriptor() ([]byte, []int) {
	return file_proto_shares_proto_rawDescGZIP(), []int{2}
}

func (x *AddSharesBatchResponse) GetAddedCount() int64 {
	if x != nil {
		return x.AddedCount
	}
	return 0
}

var File_proto_shares_proto protoreflect.FileDescriptor

var file_proto_shares_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x72, 0x70, 0x63, 0x22, 0xd1, 0x02, 0x0a, 0x05, 0x53,
	0x68, 0x61, 0x72, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x6f, 0x69, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x63, 0x6f, 0x69, 0x6e, 0x49, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x69, 0x66,
	0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64,
	0x69, 0x66, 0x66, 0x69, 0x63, 0x75, 0x6c, 0x74, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x5f, 0x64, 0x69, 0x66, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x44, 0x69, 0x66, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x17, 0x0a, 0x07,
	0x69, 0x73, 0x5f, 0x73, 0x6f, 0x6c, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69,
	0x73, 0x53, 0x6f, 0x6c, 0x6f, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x5f,
	0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x73, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x73, 0x74, 0x22, 0x3a,
	0x0a, 0x15, 0x41, 0x64, 0x64, 0x53, 0x68, 0x61, 0x72, 0x65, 0x73, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x05, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x68,
	0x61, 0x72, 0x65, 0x52, 0x05, 0x73, 0x68, 0x61, 0x72, 0x65, 0x22, 0x39, 0x0a, 0x16, 0x41, 0x64,
	0x64, 0x53, 0x68, 0x61, 0x72, 0x65, 0x73, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x64, 0x64, 0x65, 0x64, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x61, 0x64, 0x64, 0x65, 0x64,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0x5c, 0x0a, 0x0d, 0x53, 0x68, 0x61, 0x72, 0x65, 0x73, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x53, 0x68, 0x61,
	0x72, 0x65, 0x73, 0x42, 0x61, 0x74, 0x63, 0x68, 0x12, 0x1b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x41, 0x64, 0x64, 0x53, 0x68, 0x61, 0x72, 0x65, 0x73, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x64, 0x64,
	0x53, 0x68, 0x61, 0x72, 0x65, 0x73, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x1d, 0x5a, 0x1b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x61, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_shares_proto_rawDescOnce sync.Once
	file_proto_shares_proto_rawDescData = file_proto_shares_proto_rawDesc
)

func file_proto_shares_proto_rawDescGZIP() []byte {
	file_proto_shares_proto_rawDescOnce.Do(func() {
		file_proto_shares_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_shares_proto_rawDescData)
	})
	return file_proto_shares_proto_rawDescData
}

var file_proto_shares_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_shares_proto_goTypes = []interface{}{
	(*Share)(nil),                  // 0: grpc.Share
	(*AddSharesBatchRequest)(nil),  // 1: grpc.AddSharesBatchRequest
	(*AddSharesBatchResponse)(nil), // 2: grpc.AddSharesBatchResponse
}
var file_proto_shares_proto_depIdxs = []int32{
	0, // 0: grpc.AddSharesBatchRequest.share:type_name -> grpc.Share
	1, // 1: grpc.SharesService.AddSharesBatch:input_type -> grpc.AddSharesBatchRequest
	2, // 2: grpc.SharesService.AddSharesBatch:output_type -> grpc.AddSharesBatchResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_shares_proto_init() }
func file_proto_shares_proto_init() {
	if File_proto_shares_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_shares_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Share); i {
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
		file_proto_shares_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddSharesBatchRequest); i {
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
		file_proto_shares_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddSharesBatchResponse); i {
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
			RawDescriptor: file_proto_shares_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_shares_proto_goTypes,
		DependencyIndexes: file_proto_shares_proto_depIdxs,
		MessageInfos:      file_proto_shares_proto_msgTypes,
	}.Build()
	File_proto_shares_proto = out.File
	file_proto_shares_proto_rawDesc = nil
	file_proto_shares_proto_goTypes = nil
	file_proto_shares_proto_depIdxs = nil
}

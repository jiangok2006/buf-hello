// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: album_detail_service/v1/album_detail_service.proto

package album_detail_servicev1

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

type GetAlbumDetailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetAlbumDetailRequest) Reset() {
	*x = GetAlbumDetailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_album_detail_service_v1_album_detail_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAlbumDetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAlbumDetailRequest) ProtoMessage() {}

func (x *GetAlbumDetailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_album_detail_service_v1_album_detail_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAlbumDetailRequest.ProtoReflect.Descriptor instead.
func (*GetAlbumDetailRequest) Descriptor() ([]byte, []int) {
	return file_album_detail_service_v1_album_detail_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetAlbumDetailRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetAlbumDetailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title string  `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Price float64 `protobuf:"fixed64,3,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *GetAlbumDetailResponse) Reset() {
	*x = GetAlbumDetailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_album_detail_service_v1_album_detail_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAlbumDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAlbumDetailResponse) ProtoMessage() {}

func (x *GetAlbumDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_album_detail_service_v1_album_detail_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAlbumDetailResponse.ProtoReflect.Descriptor instead.
func (*GetAlbumDetailResponse) Descriptor() ([]byte, []int) {
	return file_album_detail_service_v1_album_detail_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetAlbumDetailResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetAlbumDetailResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *GetAlbumDetailResponse) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

var File_album_detail_service_v1_album_detail_service_proto protoreflect.FileDescriptor

var file_album_detail_service_v1_album_detail_service_proto_rawDesc = []byte{
	0x0a, 0x32, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x5f,
	0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x5f, 0x64, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x22, 0x27, 0x0a,
	0x15, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x54, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x62,
	0x75, 0x6d, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x32, 0x89, 0x01, 0x0a,
	0x12, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x73, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x2e, 0x2e, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x5f, 0x64, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x5f, 0x64, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xfc, 0x01, 0x0a, 0x1b, 0x63, 0x6f, 0x6d,
	0x2e, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x17, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x4f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x62, 0x75, 0x66, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x62, 0x75, 0x66, 0x2d, 0x74, 0x6f, 0x75,
	0x72, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x5f, 0x64, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x6c,
	0x62, 0x75, 0x6d, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x41, 0x58, 0x58, 0xaa, 0x02, 0x15, 0x41, 0x6c, 0x62,
	0x75, 0x6d, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x56, 0x31, 0xca, 0x02, 0x15, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x21, 0x41, 0x6c, 0x62,
	0x75, 0x6d, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c,
	0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x16, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_album_detail_service_v1_album_detail_service_proto_rawDescOnce sync.Once
	file_album_detail_service_v1_album_detail_service_proto_rawDescData = file_album_detail_service_v1_album_detail_service_proto_rawDesc
)

func file_album_detail_service_v1_album_detail_service_proto_rawDescGZIP() []byte {
	file_album_detail_service_v1_album_detail_service_proto_rawDescOnce.Do(func() {
		file_album_detail_service_v1_album_detail_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_album_detail_service_v1_album_detail_service_proto_rawDescData)
	})
	return file_album_detail_service_v1_album_detail_service_proto_rawDescData
}

var file_album_detail_service_v1_album_detail_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_album_detail_service_v1_album_detail_service_proto_goTypes = []any{
	(*GetAlbumDetailRequest)(nil),  // 0: album_detail_service.v1.GetAlbumDetailRequest
	(*GetAlbumDetailResponse)(nil), // 1: album_detail_service.v1.GetAlbumDetailResponse
}
var file_album_detail_service_v1_album_detail_service_proto_depIdxs = []int32{
	0, // 0: album_detail_service.v1.AlbumDetailService.GetAlbumDetail:input_type -> album_detail_service.v1.GetAlbumDetailRequest
	1, // 1: album_detail_service.v1.AlbumDetailService.GetAlbumDetail:output_type -> album_detail_service.v1.GetAlbumDetailResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_album_detail_service_v1_album_detail_service_proto_init() }
func file_album_detail_service_v1_album_detail_service_proto_init() {
	if File_album_detail_service_v1_album_detail_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_album_detail_service_v1_album_detail_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetAlbumDetailRequest); i {
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
		file_album_detail_service_v1_album_detail_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetAlbumDetailResponse); i {
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
			RawDescriptor: file_album_detail_service_v1_album_detail_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_album_detail_service_v1_album_detail_service_proto_goTypes,
		DependencyIndexes: file_album_detail_service_v1_album_detail_service_proto_depIdxs,
		MessageInfos:      file_album_detail_service_v1_album_detail_service_proto_msgTypes,
	}.Build()
	File_album_detail_service_v1_album_detail_service_proto = out.File
	file_album_detail_service_v1_album_detail_service_proto_rawDesc = nil
	file_album_detail_service_v1_album_detail_service_proto_goTypes = nil
	file_album_detail_service_v1_album_detail_service_proto_depIdxs = nil
}

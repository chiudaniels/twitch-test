// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0-devel
// 	protoc        v3.12.3
// source: rpc/service.proto

package rpc

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// A Hat is a piece of headwear made by a Haberdasher.
type Movie struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"` // anything but "invisible"
	Year  int32  `protobuf:"varint,2,opt,name=year,proto3" json:"year,omitempty"`  // i.e. "bowler"
}

func (x *Movie) Reset() {
	*x = Movie{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Movie) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Movie) ProtoMessage() {}

func (x *Movie) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Movie.ProtoReflect.Descriptor instead.
func (*Movie) Descriptor() ([]byte, []int) {
	return file_rpc_service_proto_rawDescGZIP(), []int{0}
}

func (x *Movie) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Movie) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

type Movies struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*Movie `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *Movies) Reset() {
	*x = Movies{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Movies) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Movies) ProtoMessage() {}

func (x *Movies) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Movies.ProtoReflect.Descriptor instead.
func (*Movies) Descriptor() ([]byte, []int) {
	return file_rpc_service_proto_rawDescGZIP(), []int{1}
}

func (x *Movies) GetData() []*Movie {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetAllQueryParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllQueryParam) Reset() {
	*x = GetAllQueryParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpc_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllQueryParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllQueryParam) ProtoMessage() {}

func (x *GetAllQueryParam) ProtoReflect() protoreflect.Message {
	mi := &file_rpc_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllQueryParam.ProtoReflect.Descriptor instead.
func (*GetAllQueryParam) Descriptor() ([]byte, []int) {
	return file_rpc_service_proto_rawDescGZIP(), []int{2}
}

var File_rpc_service_proto protoreflect.FileDescriptor

var file_rpc_service_proto_rawDesc = []byte{
	0x0a, 0x11, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x03, 0x72, 0x70, 0x63, 0x22, 0x31, 0x0a, 0x05, 0x4d, 0x6f, 0x76, 0x69,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x22, 0x28, 0x0a, 0x06, 0x4d,
	0x6f, 0x76, 0x69, 0x65, 0x73, 0x12, 0x1e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x12, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x32, 0x37, 0x0a, 0x07, 0x54, 0x65, 0x73,
	0x74, 0x41, 0x50, 0x49, 0x12, 0x2c, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x15,
	0x2e, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x0b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x4d, 0x6f, 0x76, 0x69,
	0x65, 0x73, 0x42, 0x05, 0x5a, 0x03, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_rpc_service_proto_rawDescOnce sync.Once
	file_rpc_service_proto_rawDescData = file_rpc_service_proto_rawDesc
)

func file_rpc_service_proto_rawDescGZIP() []byte {
	file_rpc_service_proto_rawDescOnce.Do(func() {
		file_rpc_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpc_service_proto_rawDescData)
	})
	return file_rpc_service_proto_rawDescData
}

var file_rpc_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_rpc_service_proto_goTypes = []interface{}{
	(*Movie)(nil),            // 0: rpc.Movie
	(*Movies)(nil),           // 1: rpc.Movies
	(*GetAllQueryParam)(nil), // 2: rpc.GetAllQueryParam
}
var file_rpc_service_proto_depIdxs = []int32{
	0, // 0: rpc.Movies.data:type_name -> rpc.Movie
	2, // 1: rpc.TestAPI.GetAll:input_type -> rpc.GetAllQueryParam
	1, // 2: rpc.TestAPI.GetAll:output_type -> rpc.Movies
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_rpc_service_proto_init() }
func file_rpc_service_proto_init() {
	if File_rpc_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpc_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Movie); i {
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
		file_rpc_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Movies); i {
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
		file_rpc_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllQueryParam); i {
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
			RawDescriptor: file_rpc_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpc_service_proto_goTypes,
		DependencyIndexes: file_rpc_service_proto_depIdxs,
		MessageInfos:      file_rpc_service_proto_msgTypes,
	}.Build()
	File_rpc_service_proto = out.File
	file_rpc_service_proto_rawDesc = nil
	file_rpc_service_proto_goTypes = nil
	file_rpc_service_proto_depIdxs = nil
}
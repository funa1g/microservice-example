// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: pkg/petshop/delivery/petshop.proto

package delivery

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

type GetPetListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit int32 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *GetPetListRequest) Reset() {
	*x = GetPetListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_petshop_delivery_petshop_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPetListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPetListRequest) ProtoMessage() {}

func (x *GetPetListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_petshop_delivery_petshop_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPetListRequest.ProtoReflect.Descriptor instead.
func (*GetPetListRequest) Descriptor() ([]byte, []int) {
	return file_pkg_petshop_delivery_petshop_proto_rawDescGZIP(), []int{0}
}

func (x *GetPetListRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type PostPetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Tag  string `protobuf:"bytes,2,opt,name=tag,proto3" json:"tag,omitempty"`
}

func (x *PostPetRequest) Reset() {
	*x = PostPetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_petshop_delivery_petshop_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostPetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostPetRequest) ProtoMessage() {}

func (x *PostPetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_petshop_delivery_petshop_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostPetRequest.ProtoReflect.Descriptor instead.
func (*PostPetRequest) Descriptor() ([]byte, []int) {
	return file_pkg_petshop_delivery_petshop_proto_rawDescGZIP(), []int{1}
}

func (x *PostPetRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PostPetRequest) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

type GetPetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetPetRequest) Reset() {
	*x = GetPetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_petshop_delivery_petshop_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPetRequest) ProtoMessage() {}

func (x *GetPetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_petshop_delivery_petshop_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPetRequest.ProtoReflect.Descriptor instead.
func (*GetPetRequest) Descriptor() ([]byte, []int) {
	return file_pkg_petshop_delivery_petshop_proto_rawDescGZIP(), []int{2}
}

func (x *GetPetRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type PetListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*PetResponse `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *PetListResponse) Reset() {
	*x = PetListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_petshop_delivery_petshop_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PetListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PetListResponse) ProtoMessage() {}

func (x *PetListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_petshop_delivery_petshop_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PetListResponse.ProtoReflect.Descriptor instead.
func (*PetListResponse) Descriptor() ([]byte, []int) {
	return file_pkg_petshop_delivery_petshop_proto_rawDescGZIP(), []int{3}
}

func (x *PetListResponse) GetResults() []*PetResponse {
	if x != nil {
		return x.Results
	}
	return nil
}

type PetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Tag  string `protobuf:"bytes,3,opt,name=tag,proto3" json:"tag,omitempty"`
}

func (x *PetResponse) Reset() {
	*x = PetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_petshop_delivery_petshop_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PetResponse) ProtoMessage() {}

func (x *PetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_petshop_delivery_petshop_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PetResponse.ProtoReflect.Descriptor instead.
func (*PetResponse) Descriptor() ([]byte, []int) {
	return file_pkg_petshop_delivery_petshop_proto_rawDescGZIP(), []int{4}
}

func (x *PetResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PetResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PetResponse) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

var File_pkg_petshop_delivery_petshop_proto protoreflect.FileDescriptor

var file_pkg_petshop_delivery_petshop_proto_rawDesc = []byte{
	0x0a, 0x22, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x65, 0x74, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x64, 0x65,
	0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x70, 0x65, 0x74, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x65, 0x74, 0x73, 0x68, 0x6f, 0x70, 0x22, 0x29, 0x0a,
	0x11, 0x47, 0x65, 0x74, 0x50, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x36, 0x0a, 0x0e, 0x50, 0x6f, 0x73, 0x74,
	0x50, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67,
	0x22, 0x1f, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x50, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x41, 0x0a, 0x0f, 0x50, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x65, 0x74, 0x73, 0x68, 0x6f, 0x70, 0x2e,
	0x50, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x07, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x73, 0x22, 0x43, 0x0a, 0x0b, 0x50, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x61, 0x67, 0x32, 0xc9, 0x01, 0x0a, 0x07, 0x50, 0x65,
	0x74, 0x73, 0x68, 0x6f, 0x70, 0x12, 0x44, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50, 0x65, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x1a, 0x2e, 0x70, 0x65, 0x74, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x47, 0x65,
	0x74, 0x50, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x18, 0x2e, 0x70, 0x65, 0x74, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x50, 0x65, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x07, 0x50,
	0x6f, 0x73, 0x74, 0x50, 0x65, 0x74, 0x12, 0x17, 0x2e, 0x70, 0x65, 0x74, 0x73, 0x68, 0x6f, 0x70,
	0x2e, 0x50, 0x6f, 0x73, 0x74, 0x50, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x70, 0x65, 0x74, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x50, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50, 0x65,
	0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x16, 0x2e, 0x70, 0x65, 0x74, 0x73, 0x68, 0x6f, 0x70, 0x2e,
	0x47, 0x65, 0x74, 0x50, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e,
	0x70, 0x65, 0x74, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x50, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x75, 0x6e, 0x61, 0x31, 0x67, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_petshop_delivery_petshop_proto_rawDescOnce sync.Once
	file_pkg_petshop_delivery_petshop_proto_rawDescData = file_pkg_petshop_delivery_petshop_proto_rawDesc
)

func file_pkg_petshop_delivery_petshop_proto_rawDescGZIP() []byte {
	file_pkg_petshop_delivery_petshop_proto_rawDescOnce.Do(func() {
		file_pkg_petshop_delivery_petshop_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_petshop_delivery_petshop_proto_rawDescData)
	})
	return file_pkg_petshop_delivery_petshop_proto_rawDescData
}

var file_pkg_petshop_delivery_petshop_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pkg_petshop_delivery_petshop_proto_goTypes = []interface{}{
	(*GetPetListRequest)(nil), // 0: petshop.GetPetListRequest
	(*PostPetRequest)(nil),    // 1: petshop.PostPetRequest
	(*GetPetRequest)(nil),     // 2: petshop.GetPetRequest
	(*PetListResponse)(nil),   // 3: petshop.PetListResponse
	(*PetResponse)(nil),       // 4: petshop.PetResponse
}
var file_pkg_petshop_delivery_petshop_proto_depIdxs = []int32{
	4, // 0: petshop.PetListResponse.results:type_name -> petshop.PetResponse
	0, // 1: petshop.Petshop.GetPetList:input_type -> petshop.GetPetListRequest
	1, // 2: petshop.Petshop.PostPet:input_type -> petshop.PostPetRequest
	2, // 3: petshop.Petshop.GetPetById:input_type -> petshop.GetPetRequest
	3, // 4: petshop.Petshop.GetPetList:output_type -> petshop.PetListResponse
	4, // 5: petshop.Petshop.PostPet:output_type -> petshop.PetResponse
	4, // 6: petshop.Petshop.GetPetById:output_type -> petshop.PetResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_petshop_delivery_petshop_proto_init() }
func file_pkg_petshop_delivery_petshop_proto_init() {
	if File_pkg_petshop_delivery_petshop_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_petshop_delivery_petshop_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPetListRequest); i {
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
		file_pkg_petshop_delivery_petshop_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostPetRequest); i {
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
		file_pkg_petshop_delivery_petshop_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPetRequest); i {
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
		file_pkg_petshop_delivery_petshop_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PetListResponse); i {
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
		file_pkg_petshop_delivery_petshop_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PetResponse); i {
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
			RawDescriptor: file_pkg_petshop_delivery_petshop_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_petshop_delivery_petshop_proto_goTypes,
		DependencyIndexes: file_pkg_petshop_delivery_petshop_proto_depIdxs,
		MessageInfos:      file_pkg_petshop_delivery_petshop_proto_msgTypes,
	}.Build()
	File_pkg_petshop_delivery_petshop_proto = out.File
	file_pkg_petshop_delivery_petshop_proto_rawDesc = nil
	file_pkg_petshop_delivery_petshop_proto_goTypes = nil
	file_pkg_petshop_delivery_petshop_proto_depIdxs = nil
}
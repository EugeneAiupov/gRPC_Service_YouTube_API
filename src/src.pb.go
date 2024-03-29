// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: .proto

// Имя пакета, используется для генерации пространства имен в коде.

package src

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

// Сообщение запроса, содержащее URL видеоролика.
type GetThumbnailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VideoUrl string `protobuf:"bytes,1,opt,name=video_url,json=videoUrl,proto3" json:"video_url,omitempty"`
}

func (x *GetThumbnailRequest) Reset() {
	*x = GetThumbnailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file___proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetThumbnailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetThumbnailRequest) ProtoMessage() {}

func (x *GetThumbnailRequest) ProtoReflect() protoreflect.Message {
	mi := &file___proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetThumbnailRequest.ProtoReflect.Descriptor instead.
func (*GetThumbnailRequest) Descriptor() ([]byte, []int) {
	return file___proto_rawDescGZIP(), []int{0}
}

func (x *GetThumbnailRequest) GetVideoUrl() string {
	if x != nil {
		return x.VideoUrl
	}
	return ""
}

// Сообщение ответа
type GetThumbnailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ThumbnailPath string `protobuf:"bytes,1,opt,name=thumbnail_path,json=thumbnailPath,proto3" json:"thumbnail_path,omitempty"`
}

func (x *GetThumbnailResponse) Reset() {
	*x = GetThumbnailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file___proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetThumbnailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetThumbnailResponse) ProtoMessage() {}

func (x *GetThumbnailResponse) ProtoReflect() protoreflect.Message {
	mi := &file___proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetThumbnailResponse.ProtoReflect.Descriptor instead.
func (*GetThumbnailResponse) Descriptor() ([]byte, []int) {
	return file___proto_rawDescGZIP(), []int{1}
}

func (x *GetThumbnailResponse) GetThumbnailPath() string {
	if x != nil {
		return x.ThumbnailPath
	}
	return ""
}

var File___proto protoreflect.FileDescriptor

var file___proto_rawDesc = []byte{
	0x0a, 0x06, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x32,
	0x0a, 0x13, 0x47, 0x65, 0x74, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x55,
	0x72, 0x6c, 0x22, 0x3d, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x68,
	0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x50, 0x61, 0x74,
	0x68, 0x32, 0x5d, 0x0a, 0x12, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x54, 0x68,
	0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x12, 0x19, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x47,
	0x65, 0x74, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x68, 0x75,
	0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x73, 0x72, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file___proto_rawDescOnce sync.Once
	file___proto_rawDescData = file___proto_rawDesc
)

func file___proto_rawDescGZIP() []byte {
	file___proto_rawDescOnce.Do(func() {
		file___proto_rawDescData = protoimpl.X.CompressGZIP(file___proto_rawDescData)
	})
	return file___proto_rawDescData
}

var file___proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file___proto_goTypes = []interface{}{
	(*GetThumbnailRequest)(nil),  // 0: main.GetThumbnailRequest
	(*GetThumbnailResponse)(nil), // 1: main.GetThumbnailResponse
}
var file___proto_depIdxs = []int32{
	0, // 0: main.ThumbnailerService.GetThumbnail:input_type -> main.GetThumbnailRequest
	1, // 1: main.ThumbnailerService.GetThumbnail:output_type -> main.GetThumbnailResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file___proto_init() }
func file___proto_init() {
	if File___proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file___proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetThumbnailRequest); i {
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
		file___proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetThumbnailResponse); i {
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
			RawDescriptor: file___proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file___proto_goTypes,
		DependencyIndexes: file___proto_depIdxs,
		MessageInfos:      file___proto_msgTypes,
	}.Build()
	File___proto = out.File
	file___proto_rawDesc = nil
	file___proto_goTypes = nil
	file___proto_depIdxs = nil
}

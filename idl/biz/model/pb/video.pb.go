// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.5.1
// source: video.proto

package pb

import (
	_ "go_tiktok_project/idl/biz/model/api"
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

type DouyinPublishListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id" query:"user_id"` // 用户id
	Token  string `protobuf:"bytes,2,opt,name=token,proto3" json:"token" query:"token"`                    // 用户鉴权token
}

func (x *DouyinPublishListRequest) Reset() {
	*x = DouyinPublishListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DouyinPublishListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DouyinPublishListRequest) ProtoMessage() {}

func (x *DouyinPublishListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DouyinPublishListRequest.ProtoReflect.Descriptor instead.
func (*DouyinPublishListRequest) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{0}
}

func (x *DouyinPublishListRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *DouyinPublishListRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type DouyinPublishListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32    `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code" form:"status_code" query:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string   `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg" form:"status_msg" query:"status_msg"`       // 返回状态描述
	VideoList  []*Video `protobuf:"bytes,3,rep,name=video_list,json=videoList,proto3" json:"video_list" form:"video_list" query:"video_list"`       // 用户发布的视频列表
}

func (x *DouyinPublishListResponse) Reset() {
	*x = DouyinPublishListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DouyinPublishListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DouyinPublishListResponse) ProtoMessage() {}

func (x *DouyinPublishListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DouyinPublishListResponse.ProtoReflect.Descriptor instead.
func (*DouyinPublishListResponse) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{1}
}

func (x *DouyinPublishListResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *DouyinPublishListResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *DouyinPublishListResponse) GetVideoList() []*Video {
	if x != nil {
		return x.VideoList
	}
	return nil
}

type DouyinPublishActionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token" form:"token" query:"token"` // 用户鉴权token
	Data  []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data" form:"data" query:"data"`     // 视频数据
	Title string `protobuf:"bytes,3,opt,name=title,proto3" json:"title" form:"title" query:"title"` // 视频标题
}

func (x *DouyinPublishActionRequest) Reset() {
	*x = DouyinPublishActionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DouyinPublishActionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DouyinPublishActionRequest) ProtoMessage() {}

func (x *DouyinPublishActionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DouyinPublishActionRequest.ProtoReflect.Descriptor instead.
func (*DouyinPublishActionRequest) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{2}
}

func (x *DouyinPublishActionRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *DouyinPublishActionRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *DouyinPublishActionRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

type DouyinPublishActionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code" form:"status_code" query:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg" form:"status_msg" query:"status_msg"`       // 返回状态描述
}

func (x *DouyinPublishActionResponse) Reset() {
	*x = DouyinPublishActionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DouyinPublishActionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DouyinPublishActionResponse) ProtoMessage() {}

func (x *DouyinPublishActionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_video_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DouyinPublishActionResponse.ProtoReflect.Descriptor instead.
func (*DouyinPublishActionResponse) Descriptor() ([]byte, []int) {
	return file_video_proto_rawDescGZIP(), []int{3}
}

func (x *DouyinPublishActionResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *DouyinPublishActionResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

var File_video_proto protoreflect.FileDescriptor

var file_video_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x64,
	0x6f, 0x75, 0x79, 0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x64, 0x0a, 0x1b, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x5f, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x24, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x42, 0x0b, 0xb2, 0xbb, 0x18, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xb2, 0xbb, 0x18, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x91, 0x01, 0x0a, 0x1c, 0x64, 0x6f,
	0x75, 0x79, 0x69, 0x6e, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x12, 0x31, 0x0a, 0x0a, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x52, 0x09, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x5f, 0x0a,
	0x1d, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x5f,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x60,
	0x0a, 0x1e, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68,
	0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67,
	0x42, 0x25, 0x5a, 0x23, 0x67, 0x6f, 0x5f, 0x74, 0x69, 0x6b, 0x74, 0x6f, 0x6b, 0x5f, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x69, 0x64, 0x6c, 0x2f, 0x62, 0x69, 0x7a, 0x2f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2f, 0x70, 0x62, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_video_proto_rawDescOnce sync.Once
	file_video_proto_rawDescData = file_video_proto_rawDesc
)

func file_video_proto_rawDescGZIP() []byte {
	file_video_proto_rawDescOnce.Do(func() {
		file_video_proto_rawDescData = protoimpl.X.CompressGZIP(file_video_proto_rawDescData)
	})
	return file_video_proto_rawDescData
}

var file_video_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_video_proto_goTypes = []interface{}{
	(*DouyinPublishListRequest)(nil),    // 0: douyin.core.douyin_publish_list_request
	(*DouyinPublishListResponse)(nil),   // 1: douyin.core.douyin_publish_list_response
	(*DouyinPublishActionRequest)(nil),  // 2: douyin.core.douyin_publish_action_request
	(*DouyinPublishActionResponse)(nil), // 3: douyin.core.douyin_publish_action_response
	(*Video)(nil),                       // 4: douyin.core.Video
}
var file_video_proto_depIdxs = []int32{
	4, // 0: douyin.core.douyin_publish_list_response.video_list:type_name -> douyin.core.Video
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_video_proto_init() }
func file_video_proto_init() {
	if File_video_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_video_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DouyinPublishListRequest); i {
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
		file_video_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DouyinPublishListResponse); i {
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
		file_video_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DouyinPublishActionRequest); i {
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
		file_video_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DouyinPublishActionResponse); i {
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
			RawDescriptor: file_video_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_video_proto_goTypes,
		DependencyIndexes: file_video_proto_depIdxs,
		MessageInfos:      file_video_proto_msgTypes,
	}.Build()
	File_video_proto = out.File
	file_video_proto_rawDesc = nil
	file_video_proto_goTypes = nil
	file_video_proto_depIdxs = nil
}
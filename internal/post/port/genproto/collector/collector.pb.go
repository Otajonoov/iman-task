// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.6.1
// source: collector/collector.proto

package collector

import (
	empty "github.com/golang/protobuf/ptypes/empty"
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

type Post struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostId int64  `protobuf:"varint,1,opt,name=postId,proto3" json:"postId,omitempty"`
	UserId int64  `protobuf:"varint,2,opt,name=userId,proto3" json:"userId,omitempty"`
	Title  string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Body   string `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	Page   int64  `protobuf:"varint,5,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *Post) Reset() {
	*x = Post{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collector_collector_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Post) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Post) ProtoMessage() {}

func (x *Post) ProtoReflect() protoreflect.Message {
	mi := &file_collector_collector_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Post.ProtoReflect.Descriptor instead.
func (*Post) Descriptor() ([]byte, []int) {
	return file_collector_collector_proto_rawDescGZIP(), []int{0}
}

func (x *Post) GetPostId() int64 {
	if x != nil {
		return x.PostId
	}
	return 0
}

func (x *Post) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Post) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Post) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *Post) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

type Posts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Posts []*Post `protobuf:"bytes,1,rep,name=posts,proto3" json:"posts,omitempty"`
}

func (x *Posts) Reset() {
	*x = Posts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_collector_collector_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Posts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Posts) ProtoMessage() {}

func (x *Posts) ProtoReflect() protoreflect.Message {
	mi := &file_collector_collector_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Posts.ProtoReflect.Descriptor instead.
func (*Posts) Descriptor() ([]byte, []int) {
	return file_collector_collector_proto_rawDescGZIP(), []int{1}
}

func (x *Posts) GetPosts() []*Post {
	if x != nil {
		return x.Posts
	}
	return nil
}

var File_collector_collector_proto protoreflect.FileDescriptor

var file_collector_collector_proto_rawDesc = []byte{
	0x0a, 0x19, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2f, 0x63, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x63, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x74, 0x0a, 0x04, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70,
	0x6f, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x70, 0x6f, 0x73,
	0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x2e, 0x0a, 0x05, 0x50, 0x6f, 0x73,
	0x74, 0x73, 0x12, 0x25, 0x0a, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x50, 0x6f,
	0x73, 0x74, 0x52, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x32, 0x47, 0x0a, 0x09, 0x43, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x3a, 0x0a, 0x0c, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x10,
	0x2e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x73,
	0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_collector_collector_proto_rawDescOnce sync.Once
	file_collector_collector_proto_rawDescData = file_collector_collector_proto_rawDesc
)

func file_collector_collector_proto_rawDescGZIP() []byte {
	file_collector_collector_proto_rawDescOnce.Do(func() {
		file_collector_collector_proto_rawDescData = protoimpl.X.CompressGZIP(file_collector_collector_proto_rawDescData)
	})
	return file_collector_collector_proto_rawDescData
}

var file_collector_collector_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_collector_collector_proto_goTypes = []interface{}{
	(*Post)(nil),        // 0: collector.Post
	(*Posts)(nil),       // 1: collector.Posts
	(*empty.Empty)(nil), // 2: google.protobuf.Empty
}
var file_collector_collector_proto_depIdxs = []int32{
	0, // 0: collector.Posts.posts:type_name -> collector.Post
	2, // 1: collector.Collector.CollectPosts:input_type -> google.protobuf.Empty
	1, // 2: collector.Collector.CollectPosts:output_type -> collector.Posts
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_collector_collector_proto_init() }
func file_collector_collector_proto_init() {
	if File_collector_collector_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_collector_collector_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Post); i {
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
		file_collector_collector_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Posts); i {
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
			RawDescriptor: file_collector_collector_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_collector_collector_proto_goTypes,
		DependencyIndexes: file_collector_collector_proto_depIdxs,
		MessageInfos:      file_collector_collector_proto_msgTypes,
	}.Build()
	File_collector_collector_proto = out.File
	file_collector_collector_proto_rawDesc = nil
	file_collector_collector_proto_goTypes = nil
	file_collector_collector_proto_depIdxs = nil
}

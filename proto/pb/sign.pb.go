// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: proto/sign.proto

package pb

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

type SignArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count     uint32   `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Threshold uint32   `protobuf:"varint,2,opt,name=threshold,proto3" json:"threshold,omitempty"`
	Alg       string   `protobuf:"bytes,3,opt,name=alg,proto3" json:"alg,omitempty"`
	Party     []string `protobuf:"bytes,4,rep,name=party,proto3" json:"party,omitempty"`
}

func (x *SignArgs) Reset() {
	*x = SignArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sign_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignArgs) ProtoMessage() {}

func (x *SignArgs) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sign_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignArgs.ProtoReflect.Descriptor instead.
func (*SignArgs) Descriptor() ([]byte, []int) {
	return file_proto_sign_proto_rawDescGZIP(), []int{0}
}

func (x *SignArgs) GetCount() uint32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *SignArgs) GetThreshold() uint32 {
	if x != nil {
		return x.Threshold
	}
	return 0
}

func (x *SignArgs) GetAlg() string {
	if x != nil {
		return x.Alg
	}
	return ""
}

func (x *SignArgs) GetParty() []string {
	if x != nil {
		return x.Party
	}
	return nil
}

type SignReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signature string `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *SignReply) Reset() {
	*x = SignReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sign_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignReply) ProtoMessage() {}

func (x *SignReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sign_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignReply.ProtoReflect.Descriptor instead.
func (*SignReply) Descriptor() ([]byte, []int) {
	return file_proto_sign_proto_rawDescGZIP(), []int{1}
}

func (x *SignReply) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

var File_proto_sign_proto protoreflect.FileDescriptor

var file_proto_sign_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x69, 0x67, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x73, 0x69, 0x67, 0x6e, 0x22, 0x66, 0x0a, 0x08, 0x53, 0x69, 0x67, 0x6e,
	0x41, 0x72, 0x67, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x68,
	0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x74,
	0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x6c, 0x67, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x6c, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61,
	0x72, 0x74, 0x79, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x70, 0x61, 0x72, 0x74, 0x79,
	0x22, 0x29, 0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1c, 0x0a,
	0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x42, 0x05, 0x5a, 0x03, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_sign_proto_rawDescOnce sync.Once
	file_proto_sign_proto_rawDescData = file_proto_sign_proto_rawDesc
)

func file_proto_sign_proto_rawDescGZIP() []byte {
	file_proto_sign_proto_rawDescOnce.Do(func() {
		file_proto_sign_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_sign_proto_rawDescData)
	})
	return file_proto_sign_proto_rawDescData
}

var file_proto_sign_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_sign_proto_goTypes = []interface{}{
	(*SignArgs)(nil),  // 0: sign.SignArgs
	(*SignReply)(nil), // 1: sign.SignReply
}
var file_proto_sign_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_sign_proto_init() }
func file_proto_sign_proto_init() {
	if File_proto_sign_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_sign_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignArgs); i {
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
		file_proto_sign_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignReply); i {
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
			RawDescriptor: file_proto_sign_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_sign_proto_goTypes,
		DependencyIndexes: file_proto_sign_proto_depIdxs,
		MessageInfos:      file_proto_sign_proto_msgTypes,
	}.Build()
	File_proto_sign_proto = out.File
	file_proto_sign_proto_rawDesc = nil
	file_proto_sign_proto_goTypes = nil
	file_proto_sign_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.14.0
// source: model/protoc/message.proto

package model

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

type Msg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topic    string      `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Format   string      `protobuf:"bytes,2,opt,name=format,proto3" json:"format,omitempty"`
	Databyte []*DataByte `protobuf:"bytes,3,rep,name=databyte,proto3" json:"databyte,omitempty"`
	Commands []*Commands `protobuf:"bytes,4,rep,name=commands,proto3" json:"commands,omitempty"`
	Talks    []*Talks    `protobuf:"bytes,5,rep,name=talks,proto3" json:"talks,omitempty"`
}

func (x *Msg) Reset() {
	*x = Msg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_protoc_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Msg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Msg) ProtoMessage() {}

func (x *Msg) ProtoReflect() protoreflect.Message {
	mi := &file_model_protoc_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Msg.ProtoReflect.Descriptor instead.
func (*Msg) Descriptor() ([]byte, []int) {
	return file_model_protoc_message_proto_rawDescGZIP(), []int{0}
}

func (x *Msg) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *Msg) GetFormat() string {
	if x != nil {
		return x.Format
	}
	return ""
}

func (x *Msg) GetDatabyte() []*DataByte {
	if x != nil {
		return x.Databyte
	}
	return nil
}

func (x *Msg) GetCommands() []*Commands {
	if x != nil {
		return x.Commands
	}
	return nil
}

func (x *Msg) GetTalks() []*Talks {
	if x != nil {
		return x.Talks
	}
	return nil
}

type DataByte struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data     []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Filename string   `protobuf:"bytes,2,opt,name=filename,proto3" json:"filename,omitempty"`
	Execute  []string `protobuf:"bytes,3,rep,name=execute,proto3" json:"execute,omitempty"`
}

func (x *DataByte) Reset() {
	*x = DataByte{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_protoc_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataByte) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataByte) ProtoMessage() {}

func (x *DataByte) ProtoReflect() protoreflect.Message {
	mi := &file_model_protoc_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataByte.ProtoReflect.Descriptor instead.
func (*DataByte) Descriptor() ([]byte, []int) {
	return file_model_protoc_message_proto_rawDescGZIP(), []int{1}
}

func (x *DataByte) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *DataByte) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *DataByte) GetExecute() []string {
	if x != nil {
		return x.Execute
	}
	return nil
}

type Commands struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Command []string `protobuf:"bytes,2,rep,name=command,proto3" json:"command,omitempty"`
}

func (x *Commands) Reset() {
	*x = Commands{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_protoc_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Commands) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Commands) ProtoMessage() {}

func (x *Commands) ProtoReflect() protoreflect.Message {
	mi := &file_model_protoc_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Commands.ProtoReflect.Descriptor instead.
func (*Commands) Descriptor() ([]byte, []int) {
	return file_model_protoc_message_proto_rawDescGZIP(), []int{2}
}

func (x *Commands) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Commands) GetCommand() []string {
	if x != nil {
		return x.Command
	}
	return nil
}

type Talks struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Text string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *Talks) Reset() {
	*x = Talks{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_protoc_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Talks) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Talks) ProtoMessage() {}

func (x *Talks) ProtoReflect() protoreflect.Message {
	mi := &file_model_protoc_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Talks.ProtoReflect.Descriptor instead.
func (*Talks) Descriptor() ([]byte, []int) {
	return file_model_protoc_message_proto_rawDescGZIP(), []int{3}
}

func (x *Talks) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Talks) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

var File_model_protoc_message_proto protoreflect.FileDescriptor

var file_model_protoc_message_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x22, 0xb4, 0x01, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70,
	0x69, 0x63, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x2c, 0x0a, 0x08, 0x64, 0x61,
	0x74, 0x61, 0x62, 0x79, 0x74, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x42, 0x79, 0x74, 0x65, 0x52, 0x08,
	0x64, 0x61, 0x74, 0x61, 0x62, 0x79, 0x74, 0x65, 0x12, 0x2c, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x52, 0x08, 0x63, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x12, 0x23, 0x0a, 0x05, 0x74, 0x61, 0x6c, 0x6b, 0x73, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x54,
	0x61, 0x6c, 0x6b, 0x73, 0x52, 0x05, 0x74, 0x61, 0x6c, 0x6b, 0x73, 0x22, 0x54, 0x0a, 0x08, 0x44,
	0x61, 0x74, 0x61, 0x42, 0x79, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x66,
	0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66,
	0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x65, 0x22, 0x38, 0x0a, 0x08, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x22, 0x2f, 0x0a, 0x05, 0x54,
	0x61, 0x6c, 0x6b, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x42, 0x0a, 0x5a, 0x08,
	0x2e, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_model_protoc_message_proto_rawDescOnce sync.Once
	file_model_protoc_message_proto_rawDescData = file_model_protoc_message_proto_rawDesc
)

func file_model_protoc_message_proto_rawDescGZIP() []byte {
	file_model_protoc_message_proto_rawDescOnce.Do(func() {
		file_model_protoc_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_model_protoc_message_proto_rawDescData)
	})
	return file_model_protoc_message_proto_rawDescData
}

var file_model_protoc_message_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_model_protoc_message_proto_goTypes = []interface{}{
	(*Msg)(nil),      // 0: protoc.Msg
	(*DataByte)(nil), // 1: protoc.DataByte
	(*Commands)(nil), // 2: protoc.Commands
	(*Talks)(nil),    // 3: protoc.Talks
}
var file_model_protoc_message_proto_depIdxs = []int32{
	1, // 0: protoc.Msg.databyte:type_name -> protoc.DataByte
	2, // 1: protoc.Msg.commands:type_name -> protoc.Commands
	3, // 2: protoc.Msg.talks:type_name -> protoc.Talks
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_model_protoc_message_proto_init() }
func file_model_protoc_message_proto_init() {
	if File_model_protoc_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_model_protoc_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Msg); i {
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
		file_model_protoc_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataByte); i {
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
		file_model_protoc_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Commands); i {
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
		file_model_protoc_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Talks); i {
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
			RawDescriptor: file_model_protoc_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_model_protoc_message_proto_goTypes,
		DependencyIndexes: file_model_protoc_message_proto_depIdxs,
		MessageInfos:      file_model_protoc_message_proto_msgTypes,
	}.Build()
	File_model_protoc_message_proto = out.File
	file_model_protoc_message_proto_rawDesc = nil
	file_model_protoc_message_proto_goTypes = nil
	file_model_protoc_message_proto_depIdxs = nil
}
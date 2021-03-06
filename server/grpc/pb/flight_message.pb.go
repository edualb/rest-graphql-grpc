// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: flight_message.proto

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

type Flight struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	From         string   `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	To           string   `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	PassengerIds []string `protobuf:"bytes,4,rep,name=passengerIds,proto3" json:"passengerIds,omitempty"`
}

func (x *Flight) Reset() {
	*x = Flight{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flight_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Flight) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Flight) ProtoMessage() {}

func (x *Flight) ProtoReflect() protoreflect.Message {
	mi := &file_flight_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Flight.ProtoReflect.Descriptor instead.
func (*Flight) Descriptor() ([]byte, []int) {
	return file_flight_message_proto_rawDescGZIP(), []int{0}
}

func (x *Flight) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Flight) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *Flight) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *Flight) GetPassengerIds() []string {
	if x != nil {
		return x.PassengerIds
	}
	return nil
}

var File_flight_message_proto protoreflect.FileDescriptor

var file_flight_message_proto_rawDesc = []byte{
	0x0a, 0x14, 0x66, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x72, 0x65, 0x73, 0x74, 0x5f, 0x67, 0x72, 0x70,
	0x63, 0x5f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x22, 0x60,
	0x0a, 0x06, 0x46, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02,
	0x74, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x22, 0x0a, 0x0c,
	0x70, 0x61, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x49, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0c, 0x70, 0x61, 0x73, 0x73, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x49, 0x64, 0x73,
	0x42, 0x05, 0x5a, 0x03, 0x70, 0x62, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_flight_message_proto_rawDescOnce sync.Once
	file_flight_message_proto_rawDescData = file_flight_message_proto_rawDesc
)

func file_flight_message_proto_rawDescGZIP() []byte {
	file_flight_message_proto_rawDescOnce.Do(func() {
		file_flight_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_flight_message_proto_rawDescData)
	})
	return file_flight_message_proto_rawDescData
}

var file_flight_message_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_flight_message_proto_goTypes = []interface{}{
	(*Flight)(nil), // 0: rest_grpc_graphql.grpc.Flight
}
var file_flight_message_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_flight_message_proto_init() }
func file_flight_message_proto_init() {
	if File_flight_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_flight_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Flight); i {
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
			RawDescriptor: file_flight_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_flight_message_proto_goTypes,
		DependencyIndexes: file_flight_message_proto_depIdxs,
		MessageInfos:      file_flight_message_proto_msgTypes,
	}.Build()
	File_flight_message_proto = out.File
	file_flight_message_proto_rawDesc = nil
	file_flight_message_proto_goTypes = nil
	file_flight_message_proto_depIdxs = nil
}

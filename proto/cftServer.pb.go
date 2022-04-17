// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.20.0--rc1
// source: proto/cftServer.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_proto_cftServer_proto protoreflect.FileDescriptor

var file_proto_cftServer_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x66, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x6e, 0x61,
	0x70, 0x73, 0x68, 0x6f, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x92, 0x01, 0x0a, 0x09,
	0x43, 0x66, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x43, 0x0a, 0x0f, 0x48, 0x61, 0x6e,
	0x64, 0x6c, 0x65, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x12, 0x17, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x48, 0x65,
	0x61, 0x72, 0x74, 0x42, 0x65, 0x61, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x40,
	0x0a, 0x0e, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74,
	0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00,
	0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var file_proto_cftServer_proto_goTypes = []interface{}{
	(*HeartbeatRequest)(nil), // 0: proto.HeartbeatRequest
	(*SnapshotRequest)(nil),  // 1: proto.SnapshotRequest
	(*HeartBeatReply)(nil),   // 2: proto.HeartBeatReply
	(*SnapshotReply)(nil),    // 3: proto.SnapshotReply
}
var file_proto_cftServer_proto_depIdxs = []int32{
	0, // 0: proto.CftServer.HandleHeartbeat:input_type -> proto.HeartbeatRequest
	1, // 1: proto.CftServer.HandleSnapshot:input_type -> proto.SnapshotRequest
	2, // 2: proto.CftServer.HandleHeartbeat:output_type -> proto.HeartBeatReply
	3, // 3: proto.CftServer.HandleSnapshot:output_type -> proto.SnapshotReply
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_cftServer_proto_init() }
func file_proto_cftServer_proto_init() {
	if File_proto_cftServer_proto != nil {
		return
	}
	file_proto_heartbeat_proto_init()
	file_proto_snapshot_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_cftServer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_cftServer_proto_goTypes,
		DependencyIndexes: file_proto_cftServer_proto_depIdxs,
	}.Build()
	File_proto_cftServer_proto = out.File
	file_proto_cftServer_proto_rawDesc = nil
	file_proto_cftServer_proto_goTypes = nil
	file_proto_cftServer_proto_depIdxs = nil
}

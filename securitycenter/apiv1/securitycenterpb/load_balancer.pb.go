// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.3
// source: google/cloud/securitycenter/v1/load_balancer.proto

package securitycenterpb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Contains information related to the load balancer associated with the
// finding.
type LoadBalancer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the load balancer associated with the finding.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *LoadBalancer) Reset() {
	*x = LoadBalancer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_securitycenter_v1_load_balancer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoadBalancer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoadBalancer) ProtoMessage() {}

func (x *LoadBalancer) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_securitycenter_v1_load_balancer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoadBalancer.ProtoReflect.Descriptor instead.
func (*LoadBalancer) Descriptor() ([]byte, []int) {
	return file_google_cloud_securitycenter_v1_load_balancer_proto_rawDescGZIP(), []int{0}
}

func (x *LoadBalancer) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_google_cloud_securitycenter_v1_load_balancer_proto protoreflect.FileDescriptor

var file_google_cloud_securitycenter_v1_load_balancer_proto_rawDesc = []byte{
	0x0a, 0x32, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31,
	0x2f, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x22, 0x22, 0x0a, 0x0c, 0x4c, 0x6f, 0x61, 0x64, 0x42, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0xeb, 0x01, 0x0a, 0x22, 0x63, 0x6f, 0x6d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42,
	0x11, 0x4c, 0x6f, 0x61, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4a, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69,
	0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x62, 0x3b,
	0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x70, 0x62,
	0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x56,
	0x31, 0xca, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x5c, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x5c,
	0x56, 0x31, 0xea, 0x02, 0x21, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x3a, 0x3a, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x65, 0x6e, 0x74,
	0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_securitycenter_v1_load_balancer_proto_rawDescOnce sync.Once
	file_google_cloud_securitycenter_v1_load_balancer_proto_rawDescData = file_google_cloud_securitycenter_v1_load_balancer_proto_rawDesc
)

func file_google_cloud_securitycenter_v1_load_balancer_proto_rawDescGZIP() []byte {
	file_google_cloud_securitycenter_v1_load_balancer_proto_rawDescOnce.Do(func() {
		file_google_cloud_securitycenter_v1_load_balancer_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_securitycenter_v1_load_balancer_proto_rawDescData)
	})
	return file_google_cloud_securitycenter_v1_load_balancer_proto_rawDescData
}

var file_google_cloud_securitycenter_v1_load_balancer_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_cloud_securitycenter_v1_load_balancer_proto_goTypes = []any{
	(*LoadBalancer)(nil), // 0: google.cloud.securitycenter.v1.LoadBalancer
}
var file_google_cloud_securitycenter_v1_load_balancer_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_cloud_securitycenter_v1_load_balancer_proto_init() }
func file_google_cloud_securitycenter_v1_load_balancer_proto_init() {
	if File_google_cloud_securitycenter_v1_load_balancer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_securitycenter_v1_load_balancer_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*LoadBalancer); i {
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
			RawDescriptor: file_google_cloud_securitycenter_v1_load_balancer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_securitycenter_v1_load_balancer_proto_goTypes,
		DependencyIndexes: file_google_cloud_securitycenter_v1_load_balancer_proto_depIdxs,
		MessageInfos:      file_google_cloud_securitycenter_v1_load_balancer_proto_msgTypes,
	}.Build()
	File_google_cloud_securitycenter_v1_load_balancer_proto = out.File
	file_google_cloud_securitycenter_v1_load_balancer_proto_rawDesc = nil
	file_google_cloud_securitycenter_v1_load_balancer_proto_goTypes = nil
	file_google_cloud_securitycenter_v1_load_balancer_proto_depIdxs = nil
}

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
// source: google/cloud/datacatalog/v1/usage.proto

package datacatalogpb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Detailed statistics on the entry's usage.
//
// Usage statistics have the following limitations:
//
//   - Only BigQuery tables have them.
//   - They only include BigQuery query jobs.
//   - They might be underestimated because wildcard table references
//     are not yet counted. For more information, see
//     [Querying multiple tables using a wildcard table]
//     (https://cloud.google.com/bigquery/docs/querying-wildcard-tables)
type UsageStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The number of successful uses of the underlying entry.
	TotalCompletions float32 `protobuf:"fixed32,1,opt,name=total_completions,json=totalCompletions,proto3" json:"total_completions,omitempty"`
	// The number of failed attempts to use the underlying entry.
	TotalFailures float32 `protobuf:"fixed32,2,opt,name=total_failures,json=totalFailures,proto3" json:"total_failures,omitempty"`
	// The number of cancelled attempts to use the underlying entry.
	TotalCancellations float32 `protobuf:"fixed32,3,opt,name=total_cancellations,json=totalCancellations,proto3" json:"total_cancellations,omitempty"`
	// Total time spent only on successful uses, in milliseconds.
	TotalExecutionTimeForCompletionsMillis float32 `protobuf:"fixed32,4,opt,name=total_execution_time_for_completions_millis,json=totalExecutionTimeForCompletionsMillis,proto3" json:"total_execution_time_for_completions_millis,omitempty"`
}

func (x *UsageStats) Reset() {
	*x = UsageStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_datacatalog_v1_usage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsageStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsageStats) ProtoMessage() {}

func (x *UsageStats) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_datacatalog_v1_usage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsageStats.ProtoReflect.Descriptor instead.
func (*UsageStats) Descriptor() ([]byte, []int) {
	return file_google_cloud_datacatalog_v1_usage_proto_rawDescGZIP(), []int{0}
}

func (x *UsageStats) GetTotalCompletions() float32 {
	if x != nil {
		return x.TotalCompletions
	}
	return 0
}

func (x *UsageStats) GetTotalFailures() float32 {
	if x != nil {
		return x.TotalFailures
	}
	return 0
}

func (x *UsageStats) GetTotalCancellations() float32 {
	if x != nil {
		return x.TotalCancellations
	}
	return 0
}

func (x *UsageStats) GetTotalExecutionTimeForCompletionsMillis() float32 {
	if x != nil {
		return x.TotalExecutionTimeForCompletionsMillis
	}
	return 0
}

// Common statistics on the entry's usage.
//
// They can be set on any system.
type CommonUsageStats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// View count in source system.
	ViewCount *int64 `protobuf:"varint,1,opt,name=view_count,json=viewCount,proto3,oneof" json:"view_count,omitempty"`
}

func (x *CommonUsageStats) Reset() {
	*x = CommonUsageStats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_datacatalog_v1_usage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonUsageStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonUsageStats) ProtoMessage() {}

func (x *CommonUsageStats) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_datacatalog_v1_usage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonUsageStats.ProtoReflect.Descriptor instead.
func (*CommonUsageStats) Descriptor() ([]byte, []int) {
	return file_google_cloud_datacatalog_v1_usage_proto_rawDescGZIP(), []int{1}
}

func (x *CommonUsageStats) GetViewCount() int64 {
	if x != nil && x.ViewCount != nil {
		return *x.ViewCount
	}
	return 0
}

// The set of all usage signals that Data Catalog stores.
//
// Note: Usually, these signals are updated daily. In rare cases, an update may
// fail but will be performed again on the next day.
type UsageSignal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The end timestamp of the duration of usage statistics.
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	// Output only. BigQuery usage statistics over each of the predefined time
	// ranges.
	//
	// Supported time ranges are `{"24H", "7D", "30D"}`.
	UsageWithinTimeRange map[string]*UsageStats `protobuf:"bytes,2,rep,name=usage_within_time_range,json=usageWithinTimeRange,proto3" json:"usage_within_time_range,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Common usage statistics over each of the predefined time ranges.
	//
	// Supported time ranges are `{"24H", "7D", "30D", "Lifetime"}`.
	CommonUsageWithinTimeRange map[string]*CommonUsageStats `protobuf:"bytes,3,rep,name=common_usage_within_time_range,json=commonUsageWithinTimeRange,proto3" json:"common_usage_within_time_range,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Favorite count in the source system.
	FavoriteCount *int64 `protobuf:"varint,4,opt,name=favorite_count,json=favoriteCount,proto3,oneof" json:"favorite_count,omitempty"`
}

func (x *UsageSignal) Reset() {
	*x = UsageSignal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_datacatalog_v1_usage_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsageSignal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsageSignal) ProtoMessage() {}

func (x *UsageSignal) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_datacatalog_v1_usage_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsageSignal.ProtoReflect.Descriptor instead.
func (*UsageSignal) Descriptor() ([]byte, []int) {
	return file_google_cloud_datacatalog_v1_usage_proto_rawDescGZIP(), []int{2}
}

func (x *UsageSignal) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

func (x *UsageSignal) GetUsageWithinTimeRange() map[string]*UsageStats {
	if x != nil {
		return x.UsageWithinTimeRange
	}
	return nil
}

func (x *UsageSignal) GetCommonUsageWithinTimeRange() map[string]*CommonUsageStats {
	if x != nil {
		return x.CommonUsageWithinTimeRange
	}
	return nil
}

func (x *UsageSignal) GetFavoriteCount() int64 {
	if x != nil && x.FavoriteCount != nil {
		return *x.FavoriteCount
	}
	return 0
}

var File_google_cloud_datacatalog_v1_usage_proto protoreflect.FileDescriptor

var file_google_cloud_datacatalog_v1_usage_proto_rawDesc = []byte{
	0x0a, 0x27, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x61, 0x74, 0x61,
	0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xee, 0x01, 0x0a, 0x0a, 0x55, 0x73, 0x61,
	0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x2b, 0x0a, 0x11, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x10, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x66, 0x61,
	0x69, 0x6c, 0x75, 0x72, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0d, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x73, 0x12, 0x2f, 0x0a, 0x13, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x12, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43,
	0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x5b, 0x0a, 0x2b,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x5f, 0x66, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x6d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x26, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x69, 0x6d, 0x65, 0x46, 0x6f, 0x72, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x4d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x22, 0x45, 0x0a, 0x10, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x55, 0x73, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x22, 0x0a,
	0x0a, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x48, 0x00, 0x52, 0x09, 0x76, 0x69, 0x65, 0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01,
	0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x22, 0x88, 0x05, 0x0a, 0x0b, 0x55, 0x73, 0x61, 0x67, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c,
	0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x7e, 0x0a,
	0x17, 0x75, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x69, 0x6e, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x42,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x61,
	0x67, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x2e, 0x55, 0x73, 0x61, 0x67, 0x65, 0x57, 0x69,
	0x74, 0x68, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x14, 0x75, 0x73, 0x61, 0x67, 0x65, 0x57, 0x69,
	0x74, 0x68, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x8c, 0x01,
	0x0a, 0x1e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x77,
	0x69, 0x74, 0x68, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x48, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f,
	0x67, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x61, 0x67, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x55, 0x73, 0x61, 0x67, 0x65, 0x57, 0x69, 0x74, 0x68,
	0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x1a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x55, 0x73, 0x61, 0x67, 0x65, 0x57, 0x69, 0x74,
	0x68, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x2a, 0x0a, 0x0e,
	0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x0d, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x1a, 0x70, 0x0a, 0x19, 0x55, 0x73, 0x61, 0x67,
	0x65, 0x57, 0x69, 0x74, 0x68, 0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x3d, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f,
	0x67, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x7c, 0x0a, 0x1f, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x55, 0x73, 0x61, 0x67, 0x65, 0x57, 0x69, 0x74, 0x68, 0x69, 0x6e, 0x54,
	0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x43, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x55, 0x73, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x66, 0x61, 0x76,
	0x6f, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0xc6, 0x01, 0x0a, 0x1f,
	0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x50,
	0x01, 0x5a, 0x41, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x63, 0x61, 0x74, 0x61, 0x6c,
	0x6f, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x63, 0x61, 0x74,
	0x61, 0x6c, 0x6f, 0x67, 0x70, 0x62, 0x3b, 0x64, 0x61, 0x74, 0x61, 0x63, 0x61, 0x74, 0x61, 0x6c,
	0x6f, 0x67, 0x70, 0x62, 0xf8, 0x01, 0x01, 0xaa, 0x02, 0x1b, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x43, 0x61, 0x74, 0x61, 0x6c,
	0x6f, 0x67, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x1b, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43,
	0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x44, 0x61, 0x74, 0x61, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67,
	0x5c, 0x56, 0x31, 0xea, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x44, 0x61, 0x74, 0x61, 0x43, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67,
	0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_datacatalog_v1_usage_proto_rawDescOnce sync.Once
	file_google_cloud_datacatalog_v1_usage_proto_rawDescData = file_google_cloud_datacatalog_v1_usage_proto_rawDesc
)

func file_google_cloud_datacatalog_v1_usage_proto_rawDescGZIP() []byte {
	file_google_cloud_datacatalog_v1_usage_proto_rawDescOnce.Do(func() {
		file_google_cloud_datacatalog_v1_usage_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_datacatalog_v1_usage_proto_rawDescData)
	})
	return file_google_cloud_datacatalog_v1_usage_proto_rawDescData
}

var file_google_cloud_datacatalog_v1_usage_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_google_cloud_datacatalog_v1_usage_proto_goTypes = []any{
	(*UsageStats)(nil),            // 0: google.cloud.datacatalog.v1.UsageStats
	(*CommonUsageStats)(nil),      // 1: google.cloud.datacatalog.v1.CommonUsageStats
	(*UsageSignal)(nil),           // 2: google.cloud.datacatalog.v1.UsageSignal
	nil,                           // 3: google.cloud.datacatalog.v1.UsageSignal.UsageWithinTimeRangeEntry
	nil,                           // 4: google.cloud.datacatalog.v1.UsageSignal.CommonUsageWithinTimeRangeEntry
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_google_cloud_datacatalog_v1_usage_proto_depIdxs = []int32{
	5, // 0: google.cloud.datacatalog.v1.UsageSignal.update_time:type_name -> google.protobuf.Timestamp
	3, // 1: google.cloud.datacatalog.v1.UsageSignal.usage_within_time_range:type_name -> google.cloud.datacatalog.v1.UsageSignal.UsageWithinTimeRangeEntry
	4, // 2: google.cloud.datacatalog.v1.UsageSignal.common_usage_within_time_range:type_name -> google.cloud.datacatalog.v1.UsageSignal.CommonUsageWithinTimeRangeEntry
	0, // 3: google.cloud.datacatalog.v1.UsageSignal.UsageWithinTimeRangeEntry.value:type_name -> google.cloud.datacatalog.v1.UsageStats
	1, // 4: google.cloud.datacatalog.v1.UsageSignal.CommonUsageWithinTimeRangeEntry.value:type_name -> google.cloud.datacatalog.v1.CommonUsageStats
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_google_cloud_datacatalog_v1_usage_proto_init() }
func file_google_cloud_datacatalog_v1_usage_proto_init() {
	if File_google_cloud_datacatalog_v1_usage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_datacatalog_v1_usage_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*UsageStats); i {
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
		file_google_cloud_datacatalog_v1_usage_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CommonUsageStats); i {
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
		file_google_cloud_datacatalog_v1_usage_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*UsageSignal); i {
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
	file_google_cloud_datacatalog_v1_usage_proto_msgTypes[1].OneofWrappers = []any{}
	file_google_cloud_datacatalog_v1_usage_proto_msgTypes[2].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_datacatalog_v1_usage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_datacatalog_v1_usage_proto_goTypes,
		DependencyIndexes: file_google_cloud_datacatalog_v1_usage_proto_depIdxs,
		MessageInfos:      file_google_cloud_datacatalog_v1_usage_proto_msgTypes,
	}.Build()
	File_google_cloud_datacatalog_v1_usage_proto = out.File
	file_google_cloud_datacatalog_v1_usage_proto_rawDesc = nil
	file_google_cloud_datacatalog_v1_usage_proto_goTypes = nil
	file_google_cloud_datacatalog_v1_usage_proto_depIdxs = nil
}
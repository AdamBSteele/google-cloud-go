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
// source: google/cloud/talent/v4/completion_service.proto

package talentpb

import (
	context "context"
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Enum to specify the scope of completion.
type CompleteQueryRequest_CompletionScope int32

const (
	// Default value.
	CompleteQueryRequest_COMPLETION_SCOPE_UNSPECIFIED CompleteQueryRequest_CompletionScope = 0
	// Suggestions are based only on the data provided by the client.
	CompleteQueryRequest_TENANT CompleteQueryRequest_CompletionScope = 1
	// Suggestions are based on all jobs data in the system that's visible to
	// the client
	CompleteQueryRequest_PUBLIC CompleteQueryRequest_CompletionScope = 2
)

// Enum value maps for CompleteQueryRequest_CompletionScope.
var (
	CompleteQueryRequest_CompletionScope_name = map[int32]string{
		0: "COMPLETION_SCOPE_UNSPECIFIED",
		1: "TENANT",
		2: "PUBLIC",
	}
	CompleteQueryRequest_CompletionScope_value = map[string]int32{
		"COMPLETION_SCOPE_UNSPECIFIED": 0,
		"TENANT":                       1,
		"PUBLIC":                       2,
	}
)

func (x CompleteQueryRequest_CompletionScope) Enum() *CompleteQueryRequest_CompletionScope {
	p := new(CompleteQueryRequest_CompletionScope)
	*p = x
	return p
}

func (x CompleteQueryRequest_CompletionScope) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CompleteQueryRequest_CompletionScope) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_talent_v4_completion_service_proto_enumTypes[0].Descriptor()
}

func (CompleteQueryRequest_CompletionScope) Type() protoreflect.EnumType {
	return &file_google_cloud_talent_v4_completion_service_proto_enumTypes[0]
}

func (x CompleteQueryRequest_CompletionScope) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CompleteQueryRequest_CompletionScope.Descriptor instead.
func (CompleteQueryRequest_CompletionScope) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_talent_v4_completion_service_proto_rawDescGZIP(), []int{0, 0}
}

// Enum to specify auto-completion topics.
type CompleteQueryRequest_CompletionType int32

const (
	// Default value.
	CompleteQueryRequest_COMPLETION_TYPE_UNSPECIFIED CompleteQueryRequest_CompletionType = 0
	// Suggest job titles for jobs autocomplete.
	//
	// For
	// [CompletionType.JOB_TITLE][google.cloud.talent.v4.CompleteQueryRequest.CompletionType.JOB_TITLE]
	// type, only open jobs with the same
	// [language_codes][google.cloud.talent.v4.CompleteQueryRequest.language_codes]
	// are returned.
	CompleteQueryRequest_JOB_TITLE CompleteQueryRequest_CompletionType = 1
	// Suggest company names for jobs autocomplete.
	//
	// For
	// [CompletionType.COMPANY_NAME][google.cloud.talent.v4.CompleteQueryRequest.CompletionType.COMPANY_NAME]
	// type, only companies having open jobs with the same
	// [language_codes][google.cloud.talent.v4.CompleteQueryRequest.language_codes]
	// are returned.
	CompleteQueryRequest_COMPANY_NAME CompleteQueryRequest_CompletionType = 2
	// Suggest both job titles and company names for jobs autocomplete.
	//
	// For
	// [CompletionType.COMBINED][google.cloud.talent.v4.CompleteQueryRequest.CompletionType.COMBINED]
	// type, only open jobs with the same
	// [language_codes][google.cloud.talent.v4.CompleteQueryRequest.language_codes]
	// or companies having open jobs with the same
	// [language_codes][google.cloud.talent.v4.CompleteQueryRequest.language_codes]
	// are returned.
	CompleteQueryRequest_COMBINED CompleteQueryRequest_CompletionType = 3
)

// Enum value maps for CompleteQueryRequest_CompletionType.
var (
	CompleteQueryRequest_CompletionType_name = map[int32]string{
		0: "COMPLETION_TYPE_UNSPECIFIED",
		1: "JOB_TITLE",
		2: "COMPANY_NAME",
		3: "COMBINED",
	}
	CompleteQueryRequest_CompletionType_value = map[string]int32{
		"COMPLETION_TYPE_UNSPECIFIED": 0,
		"JOB_TITLE":                   1,
		"COMPANY_NAME":                2,
		"COMBINED":                    3,
	}
)

func (x CompleteQueryRequest_CompletionType) Enum() *CompleteQueryRequest_CompletionType {
	p := new(CompleteQueryRequest_CompletionType)
	*p = x
	return p
}

func (x CompleteQueryRequest_CompletionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CompleteQueryRequest_CompletionType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_talent_v4_completion_service_proto_enumTypes[1].Descriptor()
}

func (CompleteQueryRequest_CompletionType) Type() protoreflect.EnumType {
	return &file_google_cloud_talent_v4_completion_service_proto_enumTypes[1]
}

func (x CompleteQueryRequest_CompletionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CompleteQueryRequest_CompletionType.Descriptor instead.
func (CompleteQueryRequest_CompletionType) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_talent_v4_completion_service_proto_rawDescGZIP(), []int{0, 1}
}

// Auto-complete parameters.
type CompleteQueryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Resource name of tenant the completion is performed within.
	//
	// The format is "projects/{project_id}/tenants/{tenant_id}", for example,
	// "projects/foo/tenants/bar".
	Tenant string `protobuf:"bytes,1,opt,name=tenant,proto3" json:"tenant,omitempty"`
	// Required. The query used to generate suggestions.
	//
	// The maximum number of allowed characters is 255.
	Query string `protobuf:"bytes,2,opt,name=query,proto3" json:"query,omitempty"`
	// The list of languages of the query. This is
	// the BCP-47 language code, such as "en-US" or "sr-Latn".
	// For more information, see
	// [Tags for Identifying Languages](https://tools.ietf.org/html/bcp47).
	//
	// The maximum number of allowed characters is 255.
	LanguageCodes []string `protobuf:"bytes,3,rep,name=language_codes,json=languageCodes,proto3" json:"language_codes,omitempty"`
	// Required. Completion result count.
	//
	// The maximum allowed page size is 10.
	PageSize int32 `protobuf:"varint,4,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// If provided, restricts completion to specified company.
	//
	// The format is
	// "projects/{project_id}/tenants/{tenant_id}/companies/{company_id}", for
	// example, "projects/foo/tenants/bar/companies/baz".
	Company string `protobuf:"bytes,5,opt,name=company,proto3" json:"company,omitempty"`
	// The scope of the completion. The defaults is
	// [CompletionScope.PUBLIC][google.cloud.talent.v4.CompleteQueryRequest.CompletionScope.PUBLIC].
	Scope CompleteQueryRequest_CompletionScope `protobuf:"varint,6,opt,name=scope,proto3,enum=google.cloud.talent.v4.CompleteQueryRequest_CompletionScope" json:"scope,omitempty"`
	// The completion topic. The default is
	// [CompletionType.COMBINED][google.cloud.talent.v4.CompleteQueryRequest.CompletionType.COMBINED].
	Type CompleteQueryRequest_CompletionType `protobuf:"varint,7,opt,name=type,proto3,enum=google.cloud.talent.v4.CompleteQueryRequest_CompletionType" json:"type,omitempty"`
}

func (x *CompleteQueryRequest) Reset() {
	*x = CompleteQueryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_talent_v4_completion_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompleteQueryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompleteQueryRequest) ProtoMessage() {}

func (x *CompleteQueryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_talent_v4_completion_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompleteQueryRequest.ProtoReflect.Descriptor instead.
func (*CompleteQueryRequest) Descriptor() ([]byte, []int) {
	return file_google_cloud_talent_v4_completion_service_proto_rawDescGZIP(), []int{0}
}

func (x *CompleteQueryRequest) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *CompleteQueryRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *CompleteQueryRequest) GetLanguageCodes() []string {
	if x != nil {
		return x.LanguageCodes
	}
	return nil
}

func (x *CompleteQueryRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *CompleteQueryRequest) GetCompany() string {
	if x != nil {
		return x.Company
	}
	return ""
}

func (x *CompleteQueryRequest) GetScope() CompleteQueryRequest_CompletionScope {
	if x != nil {
		return x.Scope
	}
	return CompleteQueryRequest_COMPLETION_SCOPE_UNSPECIFIED
}

func (x *CompleteQueryRequest) GetType() CompleteQueryRequest_CompletionType {
	if x != nil {
		return x.Type
	}
	return CompleteQueryRequest_COMPLETION_TYPE_UNSPECIFIED
}

// Response of auto-complete query.
type CompleteQueryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Results of the matching job/company candidates.
	CompletionResults []*CompleteQueryResponse_CompletionResult `protobuf:"bytes,1,rep,name=completion_results,json=completionResults,proto3" json:"completion_results,omitempty"`
	// Additional information for the API invocation, such as the request
	// tracking id.
	Metadata *ResponseMetadata `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *CompleteQueryResponse) Reset() {
	*x = CompleteQueryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_talent_v4_completion_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompleteQueryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompleteQueryResponse) ProtoMessage() {}

func (x *CompleteQueryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_talent_v4_completion_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompleteQueryResponse.ProtoReflect.Descriptor instead.
func (*CompleteQueryResponse) Descriptor() ([]byte, []int) {
	return file_google_cloud_talent_v4_completion_service_proto_rawDescGZIP(), []int{1}
}

func (x *CompleteQueryResponse) GetCompletionResults() []*CompleteQueryResponse_CompletionResult {
	if x != nil {
		return x.CompletionResults
	}
	return nil
}

func (x *CompleteQueryResponse) GetMetadata() *ResponseMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// Resource that represents completion results.
type CompleteQueryResponse_CompletionResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The suggestion for the query.
	Suggestion string `protobuf:"bytes,1,opt,name=suggestion,proto3" json:"suggestion,omitempty"`
	// The completion topic.
	Type CompleteQueryRequest_CompletionType `protobuf:"varint,2,opt,name=type,proto3,enum=google.cloud.talent.v4.CompleteQueryRequest_CompletionType" json:"type,omitempty"`
	// The URI of the company image for
	// [COMPANY_NAME][google.cloud.talent.v4.CompleteQueryRequest.CompletionType.COMPANY_NAME].
	ImageUri string `protobuf:"bytes,3,opt,name=image_uri,json=imageUri,proto3" json:"image_uri,omitempty"`
}

func (x *CompleteQueryResponse_CompletionResult) Reset() {
	*x = CompleteQueryResponse_CompletionResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_talent_v4_completion_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompleteQueryResponse_CompletionResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompleteQueryResponse_CompletionResult) ProtoMessage() {}

func (x *CompleteQueryResponse_CompletionResult) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_talent_v4_completion_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompleteQueryResponse_CompletionResult.ProtoReflect.Descriptor instead.
func (*CompleteQueryResponse_CompletionResult) Descriptor() ([]byte, []int) {
	return file_google_cloud_talent_v4_completion_service_proto_rawDescGZIP(), []int{1, 0}
}

func (x *CompleteQueryResponse_CompletionResult) GetSuggestion() string {
	if x != nil {
		return x.Suggestion
	}
	return ""
}

func (x *CompleteQueryResponse_CompletionResult) GetType() CompleteQueryRequest_CompletionType {
	if x != nil {
		return x.Type
	}
	return CompleteQueryRequest_COMPLETION_TYPE_UNSPECIFIED
}

func (x *CompleteQueryResponse_CompletionResult) GetImageUri() string {
	if x != nil {
		return x.ImageUri
	}
	return ""
}

var File_google_cloud_talent_v4_completion_service_proto protoreflect.FileDescriptor

var file_google_cloud_talent_v4_completion_service_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x74,
	0x61, 0x6c, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x34, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x16, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x74, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x34, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x74, 0x61, 0x6c, 0x65, 0x6e,
	0x74, 0x2f, 0x76, 0x34, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xc6, 0x04, 0x0a, 0x14, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x06, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x22, 0xe0, 0x41, 0x02, 0xfa,
	0x41, 0x1c, 0x0a, 0x1a, 0x6a, 0x6f, 0x62, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x54, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x52, 0x06,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x12, 0x25, 0x0a, 0x0e, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x6c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x42, 0x03, 0xe0, 0x41, 0x02,
	0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x3a, 0x0a, 0x07, 0x63, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x20, 0xfa, 0x41, 0x1d,
	0x0a, 0x1b, 0x6a, 0x6f, 0x62, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x52, 0x07, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x52, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x34, 0x2e, 0x43,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x63,
	0x6f, 0x70, 0x65, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x4f, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x2e, 0x76,
	0x34, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x4b, 0x0a, 0x0f, 0x43,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x20,
	0x0a, 0x1c, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x43, 0x4f,
	0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x0a, 0x0a, 0x06, 0x54, 0x45, 0x4e, 0x41, 0x4e, 0x54, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06,
	0x50, 0x55, 0x42, 0x4c, 0x49, 0x43, 0x10, 0x02, 0x22, 0x60, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x1b, 0x43, 0x4f,
	0x4d, 0x50, 0x4c, 0x45, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x4a,
	0x4f, 0x42, 0x5f, 0x54, 0x49, 0x54, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x43, 0x4f,
	0x4d, 0x50, 0x41, 0x4e, 0x59, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08,
	0x43, 0x4f, 0x4d, 0x42, 0x49, 0x4e, 0x45, 0x44, 0x10, 0x03, 0x22, 0xef, 0x02, 0x0a, 0x15, 0x43,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6d, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x3e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x74, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x34, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x74, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x52, 0x11, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x12, 0x44, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x34, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52,
	0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0xa0, 0x01, 0x0a, 0x10, 0x43, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1e,
	0x0a, 0x0a, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x73, 0x75, 0x67, 0x67, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4f,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3b, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x61, 0x6c, 0x65,
	0x6e, 0x74, 0x2e, 0x76, 0x34, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x69, 0x32, 0xa2, 0x02, 0x0a,
	0x0a, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0xa5, 0x01, 0x0a, 0x0d,
	0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x2c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x61, 0x6c,
	0x65, 0x6e, 0x74, 0x2e, 0x76, 0x34, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x61, 0x6c, 0x65, 0x6e,
	0x74, 0x2e, 0x76, 0x34, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x37, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x31, 0x12, 0x2f, 0x2f, 0x76, 0x34, 0x2f, 0x7b, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x3d,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x74, 0x65, 0x6e, 0x61, 0x6e,
	0x74, 0x73, 0x2f, 0x2a, 0x7d, 0x3a, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x1a, 0x6c, 0xca, 0x41, 0x13, 0x6a, 0x6f, 0x62, 0x73, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x53, 0x68, 0x74,
	0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2c, 0x68, 0x74, 0x74,
	0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x6a, 0x6f, 0x62,
	0x73, 0x42, 0x70, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x34, 0x42,
	0x16, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x74,
	0x61, 0x6c, 0x65, 0x6e, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x34, 0x2f, 0x74, 0x61, 0x6c, 0x65,
	0x6e, 0x74, 0x70, 0x62, 0x3b, 0x74, 0x61, 0x6c, 0x65, 0x6e, 0x74, 0x70, 0x62, 0xa2, 0x02, 0x03,
	0x43, 0x54, 0x53, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_talent_v4_completion_service_proto_rawDescOnce sync.Once
	file_google_cloud_talent_v4_completion_service_proto_rawDescData = file_google_cloud_talent_v4_completion_service_proto_rawDesc
)

func file_google_cloud_talent_v4_completion_service_proto_rawDescGZIP() []byte {
	file_google_cloud_talent_v4_completion_service_proto_rawDescOnce.Do(func() {
		file_google_cloud_talent_v4_completion_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_talent_v4_completion_service_proto_rawDescData)
	})
	return file_google_cloud_talent_v4_completion_service_proto_rawDescData
}

var file_google_cloud_talent_v4_completion_service_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_google_cloud_talent_v4_completion_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_google_cloud_talent_v4_completion_service_proto_goTypes = []any{
	(CompleteQueryRequest_CompletionScope)(0),      // 0: google.cloud.talent.v4.CompleteQueryRequest.CompletionScope
	(CompleteQueryRequest_CompletionType)(0),       // 1: google.cloud.talent.v4.CompleteQueryRequest.CompletionType
	(*CompleteQueryRequest)(nil),                   // 2: google.cloud.talent.v4.CompleteQueryRequest
	(*CompleteQueryResponse)(nil),                  // 3: google.cloud.talent.v4.CompleteQueryResponse
	(*CompleteQueryResponse_CompletionResult)(nil), // 4: google.cloud.talent.v4.CompleteQueryResponse.CompletionResult
	(*ResponseMetadata)(nil),                       // 5: google.cloud.talent.v4.ResponseMetadata
}
var file_google_cloud_talent_v4_completion_service_proto_depIdxs = []int32{
	0, // 0: google.cloud.talent.v4.CompleteQueryRequest.scope:type_name -> google.cloud.talent.v4.CompleteQueryRequest.CompletionScope
	1, // 1: google.cloud.talent.v4.CompleteQueryRequest.type:type_name -> google.cloud.talent.v4.CompleteQueryRequest.CompletionType
	4, // 2: google.cloud.talent.v4.CompleteQueryResponse.completion_results:type_name -> google.cloud.talent.v4.CompleteQueryResponse.CompletionResult
	5, // 3: google.cloud.talent.v4.CompleteQueryResponse.metadata:type_name -> google.cloud.talent.v4.ResponseMetadata
	1, // 4: google.cloud.talent.v4.CompleteQueryResponse.CompletionResult.type:type_name -> google.cloud.talent.v4.CompleteQueryRequest.CompletionType
	2, // 5: google.cloud.talent.v4.Completion.CompleteQuery:input_type -> google.cloud.talent.v4.CompleteQueryRequest
	3, // 6: google.cloud.talent.v4.Completion.CompleteQuery:output_type -> google.cloud.talent.v4.CompleteQueryResponse
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_google_cloud_talent_v4_completion_service_proto_init() }
func file_google_cloud_talent_v4_completion_service_proto_init() {
	if File_google_cloud_talent_v4_completion_service_proto != nil {
		return
	}
	file_google_cloud_talent_v4_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_talent_v4_completion_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CompleteQueryRequest); i {
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
		file_google_cloud_talent_v4_completion_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CompleteQueryResponse); i {
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
		file_google_cloud_talent_v4_completion_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CompleteQueryResponse_CompletionResult); i {
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
			RawDescriptor: file_google_cloud_talent_v4_completion_service_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_cloud_talent_v4_completion_service_proto_goTypes,
		DependencyIndexes: file_google_cloud_talent_v4_completion_service_proto_depIdxs,
		EnumInfos:         file_google_cloud_talent_v4_completion_service_proto_enumTypes,
		MessageInfos:      file_google_cloud_talent_v4_completion_service_proto_msgTypes,
	}.Build()
	File_google_cloud_talent_v4_completion_service_proto = out.File
	file_google_cloud_talent_v4_completion_service_proto_rawDesc = nil
	file_google_cloud_talent_v4_completion_service_proto_goTypes = nil
	file_google_cloud_talent_v4_completion_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CompletionClient is the client API for Completion service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CompletionClient interface {
	// Completes the specified prefix with keyword suggestions.
	// Intended for use by a job search auto-complete search box.
	CompleteQuery(ctx context.Context, in *CompleteQueryRequest, opts ...grpc.CallOption) (*CompleteQueryResponse, error)
}

type completionClient struct {
	cc grpc.ClientConnInterface
}

func NewCompletionClient(cc grpc.ClientConnInterface) CompletionClient {
	return &completionClient{cc}
}

func (c *completionClient) CompleteQuery(ctx context.Context, in *CompleteQueryRequest, opts ...grpc.CallOption) (*CompleteQueryResponse, error) {
	out := new(CompleteQueryResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.talent.v4.Completion/CompleteQuery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CompletionServer is the server API for Completion service.
type CompletionServer interface {
	// Completes the specified prefix with keyword suggestions.
	// Intended for use by a job search auto-complete search box.
	CompleteQuery(context.Context, *CompleteQueryRequest) (*CompleteQueryResponse, error)
}

// UnimplementedCompletionServer can be embedded to have forward compatible implementations.
type UnimplementedCompletionServer struct {
}

func (*UnimplementedCompletionServer) CompleteQuery(context.Context, *CompleteQueryRequest) (*CompleteQueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompleteQuery not implemented")
}

func RegisterCompletionServer(s *grpc.Server, srv CompletionServer) {
	s.RegisterService(&_Completion_serviceDesc, srv)
}

func _Completion_CompleteQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompleteQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompletionServer).CompleteQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.talent.v4.Completion/CompleteQuery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompletionServer).CompleteQuery(ctx, req.(*CompleteQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Completion_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.talent.v4.Completion",
	HandlerType: (*CompletionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CompleteQuery",
			Handler:    _Completion_CompleteQuery_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/talent/v4/completion_service.proto",
}

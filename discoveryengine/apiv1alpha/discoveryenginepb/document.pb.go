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
// source: google/cloud/discoveryengine/v1alpha/document.proto

package discoveryenginepb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Document captures all raw metadata information of items to be recommended or
// searched.
type Document struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Data representation. One of
	// [struct_data][google.cloud.discoveryengine.v1alpha.Document.struct_data] or
	// [json_data][google.cloud.discoveryengine.v1alpha.Document.json_data] should
	// be provided otherwise an `INVALID_ARGUMENT` error is thrown.
	//
	// Types that are assignable to Data:
	//
	//	*Document_StructData
	//	*Document_JsonData
	Data isDocument_Data `protobuf_oneof:"data"`
	// Immutable. The full resource name of the document.
	// Format:
	// `projects/{project}/locations/{location}/collections/{collection}/dataStores/{data_store}/branches/{branch}/documents/{document_id}`.
	//
	// This field must be a UTF-8 encoded string with a length limit of 1024
	// characters.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Immutable. The identifier of the document.
	//
	// Id should conform to [RFC-1034](https://tools.ietf.org/html/rfc1034)
	// standard with a length limit of 63 characters.
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// The identifier of the schema located in the same data store.
	SchemaId string `protobuf:"bytes,3,opt,name=schema_id,json=schemaId,proto3" json:"schema_id,omitempty"`
	// The unstructured data linked to this document. Content must be set if this
	// document is under a
	// `CONTENT_REQUIRED` data store.
	Content *Document_Content `protobuf:"bytes,10,opt,name=content,proto3" json:"content,omitempty"`
	// The identifier of the parent document. Currently supports at most two level
	// document hierarchy.
	//
	// Id should conform to [RFC-1034](https://tools.ietf.org/html/rfc1034)
	// standard with a length limit of 63 characters.
	ParentDocumentId string `protobuf:"bytes,7,opt,name=parent_document_id,json=parentDocumentId,proto3" json:"parent_document_id,omitempty"`
	// Output only. This field is OUTPUT_ONLY.
	// It contains derived data that are not in the original input document.
	DerivedStructData *structpb.Struct `protobuf:"bytes,6,opt,name=derived_struct_data,json=derivedStructData,proto3" json:"derived_struct_data,omitempty"`
	// Access control information for the document.
	AclInfo *Document_AclInfo `protobuf:"bytes,11,opt,name=acl_info,json=aclInfo,proto3" json:"acl_info,omitempty"`
	// Output only. The last time the document was indexed. If this field is set,
	// the document could be returned in search results.
	//
	// This field is OUTPUT_ONLY. If this field is not populated, it means the
	// document has never been indexed.
	IndexTime *timestamppb.Timestamp `protobuf:"bytes,13,opt,name=index_time,json=indexTime,proto3" json:"index_time,omitempty"`
}

func (x *Document) Reset() {
	*x = Document{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_discoveryengine_v1alpha_document_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Document) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Document) ProtoMessage() {}

func (x *Document) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_discoveryengine_v1alpha_document_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Document.ProtoReflect.Descriptor instead.
func (*Document) Descriptor() ([]byte, []int) {
	return file_google_cloud_discoveryengine_v1alpha_document_proto_rawDescGZIP(), []int{0}
}

func (m *Document) GetData() isDocument_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *Document) GetStructData() *structpb.Struct {
	if x, ok := x.GetData().(*Document_StructData); ok {
		return x.StructData
	}
	return nil
}

func (x *Document) GetJsonData() string {
	if x, ok := x.GetData().(*Document_JsonData); ok {
		return x.JsonData
	}
	return ""
}

func (x *Document) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Document) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Document) GetSchemaId() string {
	if x != nil {
		return x.SchemaId
	}
	return ""
}

func (x *Document) GetContent() *Document_Content {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *Document) GetParentDocumentId() string {
	if x != nil {
		return x.ParentDocumentId
	}
	return ""
}

func (x *Document) GetDerivedStructData() *structpb.Struct {
	if x != nil {
		return x.DerivedStructData
	}
	return nil
}

func (x *Document) GetAclInfo() *Document_AclInfo {
	if x != nil {
		return x.AclInfo
	}
	return nil
}

func (x *Document) GetIndexTime() *timestamppb.Timestamp {
	if x != nil {
		return x.IndexTime
	}
	return nil
}

type isDocument_Data interface {
	isDocument_Data()
}

type Document_StructData struct {
	// The structured JSON data for the document. It should conform to the
	// registered [Schema][google.cloud.discoveryengine.v1alpha.Schema] or an
	// `INVALID_ARGUMENT` error is thrown.
	StructData *structpb.Struct `protobuf:"bytes,4,opt,name=struct_data,json=structData,proto3,oneof"`
}

type Document_JsonData struct {
	// The JSON string representation of the document. It should conform to the
	// registered [Schema][google.cloud.discoveryengine.v1alpha.Schema] or an
	// `INVALID_ARGUMENT` error is thrown.
	JsonData string `protobuf:"bytes,5,opt,name=json_data,json=jsonData,proto3,oneof"`
}

func (*Document_StructData) isDocument_Data() {}

func (*Document_JsonData) isDocument_Data() {}

// Document captures all raw metadata information of items to be recommended or
// searched.
type ProcessedDocument struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output format of the processed document.
	//
	// Types that are assignable to ProcessedDataFormat:
	//
	//	*ProcessedDocument_JsonData
	ProcessedDataFormat isProcessedDocument_ProcessedDataFormat `protobuf_oneof:"processed_data_format"`
	// Required. Full resource name of the referenced document, in the format
	// `projects/*/locations/*/collections/*/dataStores/*/branches/*/documents/*`.
	Document string `protobuf:"bytes,1,opt,name=document,proto3" json:"document,omitempty"`
}

func (x *ProcessedDocument) Reset() {
	*x = ProcessedDocument{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_discoveryengine_v1alpha_document_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProcessedDocument) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessedDocument) ProtoMessage() {}

func (x *ProcessedDocument) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_discoveryengine_v1alpha_document_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessedDocument.ProtoReflect.Descriptor instead.
func (*ProcessedDocument) Descriptor() ([]byte, []int) {
	return file_google_cloud_discoveryengine_v1alpha_document_proto_rawDescGZIP(), []int{1}
}

func (m *ProcessedDocument) GetProcessedDataFormat() isProcessedDocument_ProcessedDataFormat {
	if m != nil {
		return m.ProcessedDataFormat
	}
	return nil
}

func (x *ProcessedDocument) GetJsonData() string {
	if x, ok := x.GetProcessedDataFormat().(*ProcessedDocument_JsonData); ok {
		return x.JsonData
	}
	return ""
}

func (x *ProcessedDocument) GetDocument() string {
	if x != nil {
		return x.Document
	}
	return ""
}

type isProcessedDocument_ProcessedDataFormat interface {
	isProcessedDocument_ProcessedDataFormat()
}

type ProcessedDocument_JsonData struct {
	// The JSON string representation of the processed document.
	JsonData string `protobuf:"bytes,2,opt,name=json_data,json=jsonData,proto3,oneof"`
}

func (*ProcessedDocument_JsonData) isProcessedDocument_ProcessedDataFormat() {}

// Unstructured data linked to this document.
type Document_Content struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Content:
	//
	//	*Document_Content_RawBytes
	//	*Document_Content_Uri
	Content isDocument_Content_Content `protobuf_oneof:"content"`
	// The MIME type of the content. Supported types:
	//
	// * `application/pdf` (PDF, only native PDFs are supported for now)
	// * `text/html` (HTML)
	// * `application/vnd.openxmlformats-officedocument.wordprocessingml.document` (DOCX)
	// * `application/vnd.openxmlformats-officedocument.presentationml.presentation` (PPTX)
	// * `text/plain` (TXT)
	//
	// See https://www.iana.org/assignments/media-types/media-types.xhtml.
	MimeType string `protobuf:"bytes,1,opt,name=mime_type,json=mimeType,proto3" json:"mime_type,omitempty"`
}

func (x *Document_Content) Reset() {
	*x = Document_Content{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_discoveryengine_v1alpha_document_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Document_Content) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Document_Content) ProtoMessage() {}

func (x *Document_Content) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_discoveryengine_v1alpha_document_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Document_Content.ProtoReflect.Descriptor instead.
func (*Document_Content) Descriptor() ([]byte, []int) {
	return file_google_cloud_discoveryengine_v1alpha_document_proto_rawDescGZIP(), []int{0, 0}
}

func (m *Document_Content) GetContent() isDocument_Content_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (x *Document_Content) GetRawBytes() []byte {
	if x, ok := x.GetContent().(*Document_Content_RawBytes); ok {
		return x.RawBytes
	}
	return nil
}

func (x *Document_Content) GetUri() string {
	if x, ok := x.GetContent().(*Document_Content_Uri); ok {
		return x.Uri
	}
	return ""
}

func (x *Document_Content) GetMimeType() string {
	if x != nil {
		return x.MimeType
	}
	return ""
}

type isDocument_Content_Content interface {
	isDocument_Content_Content()
}

type Document_Content_RawBytes struct {
	// The content represented as a stream of bytes. The maximum length is
	// 1,000,000 bytes (1 MB / ~0.95 MiB).
	//
	// Note: As with all `bytes` fields, this field is represented as pure
	// binary in Protocol Buffers and base64-encoded string in JSON. For
	// example, `abc123!?$*&()'-=@~` should be represented as
	// `YWJjMTIzIT8kKiYoKSctPUB+` in JSON. See
	// https://developers.google.com/protocol-buffers/docs/proto3#json.
	RawBytes []byte `protobuf:"bytes,2,opt,name=raw_bytes,json=rawBytes,proto3,oneof"`
}

type Document_Content_Uri struct {
	// The URI of the content. Only Cloud Storage URIs (e.g.
	// `gs://bucket-name/path/to/file`) are supported. The maximum file size
	// is 2.5 MB for text-based formats, 100 MB for other formats.
	Uri string `protobuf:"bytes,3,opt,name=uri,proto3,oneof"`
}

func (*Document_Content_RawBytes) isDocument_Content_Content() {}

func (*Document_Content_Uri) isDocument_Content_Content() {}

// ACL Information of the Document.
type Document_AclInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Readers of the document.
	Readers []*Document_AclInfo_AccessRestriction `protobuf:"bytes,1,rep,name=readers,proto3" json:"readers,omitempty"`
}

func (x *Document_AclInfo) Reset() {
	*x = Document_AclInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_discoveryengine_v1alpha_document_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Document_AclInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Document_AclInfo) ProtoMessage() {}

func (x *Document_AclInfo) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_discoveryengine_v1alpha_document_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Document_AclInfo.ProtoReflect.Descriptor instead.
func (*Document_AclInfo) Descriptor() ([]byte, []int) {
	return file_google_cloud_discoveryengine_v1alpha_document_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Document_AclInfo) GetReaders() []*Document_AclInfo_AccessRestriction {
	if x != nil {
		return x.Readers
	}
	return nil
}

// AclRestriction to model complex inheritance restrictions.
//
// Example: Modeling a "Both Permit" inheritance, where to access a
// child document, user needs to have access to parent document.
//
// Document Hierarchy - Space_S --> Page_P.
//
// Readers:
//
//	Space_S: group_1, user_1
//	Page_P: group_2, group_3, user_2
//
// Space_S ACL Restriction -
//
//	{
//	  "acl_info": {
//	    "readers": [
//	      {
//	        "principals": [
//	          {
//	            "group_id": "group_1"
//	          },
//	          {
//	            "user_id": "user_1"
//	          }
//	        ]
//	      }
//	    ]
//	  }
//	}
//
// Page_P ACL Restriction.
//
//	{
//	  "acl_info": {
//	    "readers": [
//	      {
//	        "principals": [
//	          {
//	            "group_id": "group_2"
//	          },
//	          {
//	            "group_id": "group_3"
//	          },
//	          {
//	            "user_id": "user_2"
//	          }
//	        ],
//	      },
//	      {
//	        "principals": [
//	          {
//	            "group_id": "group_1"
//	          },
//	          {
//	            "user_id": "user_1"
//	          }
//	        ],
//	      }
//	    ]
//	  }
//	}
type Document_AclInfo_AccessRestriction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of principals.
	Principals []*Principal `protobuf:"bytes,1,rep,name=principals,proto3" json:"principals,omitempty"`
}

func (x *Document_AclInfo_AccessRestriction) Reset() {
	*x = Document_AclInfo_AccessRestriction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_discoveryengine_v1alpha_document_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Document_AclInfo_AccessRestriction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Document_AclInfo_AccessRestriction) ProtoMessage() {}

func (x *Document_AclInfo_AccessRestriction) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_discoveryengine_v1alpha_document_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Document_AclInfo_AccessRestriction.ProtoReflect.Descriptor instead.
func (*Document_AclInfo_AccessRestriction) Descriptor() ([]byte, []int) {
	return file_google_cloud_discoveryengine_v1alpha_document_proto_rawDescGZIP(), []int{0, 1, 0}
}

func (x *Document_AclInfo_AccessRestriction) GetPrincipals() []*Principal {
	if x != nil {
		return x.Principals
	}
	return nil
}

var File_google_cloud_discoveryengine_v1alpha_document_proto protoreflect.FileDescriptor

var file_google_cloud_discoveryengine_v1alpha_document_proto_rawDesc = []byte{
	0x0a, 0x33, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x24, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65,
	0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x31, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xee, 0x08, 0x0a, 0x08, 0x44, 0x6f,
	0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x3a, 0x0a, 0x0b, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x48, 0x00, 0x52, 0x0a, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x1d, 0x0a, 0x09, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x6a, 0x73, 0x6f, 0x6e, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x03, 0xe0, 0x41, 0x05, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x49, 0x64, 0x12, 0x50, 0x0a, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x2c,
	0x0a, 0x12, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x4c, 0x0a, 0x13,
	0x64, 0x65, 0x72, 0x69, 0x76, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x42, 0x03, 0xe0, 0x41, 0x03, 0x52, 0x11, 0x64, 0x65, 0x72, 0x69, 0x76, 0x65, 0x64,
	0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x51, 0x0a, 0x08, 0x61, 0x63,
	0x6c, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x63, 0x6c,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x61, 0x63, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3e, 0x0a,
	0x0a, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0,
	0x41, 0x03, 0x52, 0x09, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x54, 0x69, 0x6d, 0x65, 0x1a, 0x64, 0x0a,
	0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x09, 0x72, 0x61, 0x77, 0x5f,
	0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x08, 0x72,
	0x61, 0x77, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x03, 0x75, 0x72, 0x69, 0x12, 0x1b, 0x0a, 0x09, 0x6d,
	0x69, 0x6d, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6d, 0x69, 0x6d, 0x65, 0x54, 0x79, 0x70, 0x65, 0x42, 0x09, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x1a, 0xd3, 0x01, 0x0a, 0x07, 0x41, 0x63, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x62, 0x0a, 0x07, 0x72, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x48, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x41, 0x63, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52,
	0x65, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x72, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x73, 0x1a, 0x64, 0x0a, 0x11, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73,
	0x74, 0x72, 0x69, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4f, 0x0a, 0x0a, 0x70, 0x72, 0x69, 0x6e,
	0x63, 0x69, 0x70, 0x61, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2e, 0x50, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x52, 0x0a, 0x70,
	0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x73, 0x3a, 0x96, 0x02, 0xea, 0x41, 0x92, 0x02,
	0x0a, 0x27, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x66, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x7d, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x2f, 0x7b, 0x64,
	0x61, 0x74, 0x61, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x7d, 0x2f, 0x62, 0x72, 0x61, 0x6e, 0x63,
	0x68, 0x65, 0x73, 0x2f, 0x7b, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x7d, 0x2f, 0x64, 0x6f, 0x63,
	0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x7d, 0x12, 0x7f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x7d, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x73, 0x2f, 0x7b,
	0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x7d, 0x2f, 0x62, 0x72, 0x61, 0x6e,
	0x63, 0x68, 0x65, 0x73, 0x2f, 0x7b, 0x62, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x7d, 0x2f, 0x64, 0x6f,
	0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e,
	0x74, 0x7d, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x98, 0x01, 0x0a, 0x11, 0x50,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x1d, 0x0a, 0x09, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x6a, 0x73, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x4b, 0x0a, 0x08, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x2f, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x29, 0x0a, 0x27, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x08, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x17, 0x0a, 0x15,
	0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x42, 0x99, 0x02, 0x0a, 0x28, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x42, 0x0d, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x52, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x79, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0x70, 0x62, 0x3b, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x65,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x70, 0x62, 0xa2, 0x02, 0x0f, 0x44, 0x49, 0x53, 0x43, 0x4f, 0x56,
	0x45, 0x52, 0x59, 0x45, 0x4e, 0x47, 0x49, 0x4e, 0x45, 0xaa, 0x02, 0x24, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x79, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x56, 0x31, 0x41, 0x6c, 0x70, 0x68, 0x61,
	0xca, 0x02, 0x24, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c,
	0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x5c,
	0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0xea, 0x02, 0x27, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x79, 0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_discoveryengine_v1alpha_document_proto_rawDescOnce sync.Once
	file_google_cloud_discoveryengine_v1alpha_document_proto_rawDescData = file_google_cloud_discoveryengine_v1alpha_document_proto_rawDesc
)

func file_google_cloud_discoveryengine_v1alpha_document_proto_rawDescGZIP() []byte {
	file_google_cloud_discoveryengine_v1alpha_document_proto_rawDescOnce.Do(func() {
		file_google_cloud_discoveryengine_v1alpha_document_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_discoveryengine_v1alpha_document_proto_rawDescData)
	})
	return file_google_cloud_discoveryengine_v1alpha_document_proto_rawDescData
}

var file_google_cloud_discoveryengine_v1alpha_document_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_google_cloud_discoveryengine_v1alpha_document_proto_goTypes = []any{
	(*Document)(nil),                           // 0: google.cloud.discoveryengine.v1alpha.Document
	(*ProcessedDocument)(nil),                  // 1: google.cloud.discoveryengine.v1alpha.ProcessedDocument
	(*Document_Content)(nil),                   // 2: google.cloud.discoveryengine.v1alpha.Document.Content
	(*Document_AclInfo)(nil),                   // 3: google.cloud.discoveryengine.v1alpha.Document.AclInfo
	(*Document_AclInfo_AccessRestriction)(nil), // 4: google.cloud.discoveryengine.v1alpha.Document.AclInfo.AccessRestriction
	(*structpb.Struct)(nil),                    // 5: google.protobuf.Struct
	(*timestamppb.Timestamp)(nil),              // 6: google.protobuf.Timestamp
	(*Principal)(nil),                          // 7: google.cloud.discoveryengine.v1alpha.Principal
}
var file_google_cloud_discoveryengine_v1alpha_document_proto_depIdxs = []int32{
	5, // 0: google.cloud.discoveryengine.v1alpha.Document.struct_data:type_name -> google.protobuf.Struct
	2, // 1: google.cloud.discoveryengine.v1alpha.Document.content:type_name -> google.cloud.discoveryengine.v1alpha.Document.Content
	5, // 2: google.cloud.discoveryengine.v1alpha.Document.derived_struct_data:type_name -> google.protobuf.Struct
	3, // 3: google.cloud.discoveryengine.v1alpha.Document.acl_info:type_name -> google.cloud.discoveryengine.v1alpha.Document.AclInfo
	6, // 4: google.cloud.discoveryengine.v1alpha.Document.index_time:type_name -> google.protobuf.Timestamp
	4, // 5: google.cloud.discoveryengine.v1alpha.Document.AclInfo.readers:type_name -> google.cloud.discoveryengine.v1alpha.Document.AclInfo.AccessRestriction
	7, // 6: google.cloud.discoveryengine.v1alpha.Document.AclInfo.AccessRestriction.principals:type_name -> google.cloud.discoveryengine.v1alpha.Principal
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_google_cloud_discoveryengine_v1alpha_document_proto_init() }
func file_google_cloud_discoveryengine_v1alpha_document_proto_init() {
	if File_google_cloud_discoveryengine_v1alpha_document_proto != nil {
		return
	}
	file_google_cloud_discoveryengine_v1alpha_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_discoveryengine_v1alpha_document_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Document); i {
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
		file_google_cloud_discoveryengine_v1alpha_document_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ProcessedDocument); i {
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
		file_google_cloud_discoveryengine_v1alpha_document_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Document_Content); i {
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
		file_google_cloud_discoveryengine_v1alpha_document_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Document_AclInfo); i {
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
		file_google_cloud_discoveryengine_v1alpha_document_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*Document_AclInfo_AccessRestriction); i {
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
	file_google_cloud_discoveryengine_v1alpha_document_proto_msgTypes[0].OneofWrappers = []any{
		(*Document_StructData)(nil),
		(*Document_JsonData)(nil),
	}
	file_google_cloud_discoveryengine_v1alpha_document_proto_msgTypes[1].OneofWrappers = []any{
		(*ProcessedDocument_JsonData)(nil),
	}
	file_google_cloud_discoveryengine_v1alpha_document_proto_msgTypes[2].OneofWrappers = []any{
		(*Document_Content_RawBytes)(nil),
		(*Document_Content_Uri)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_discoveryengine_v1alpha_document_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_discoveryengine_v1alpha_document_proto_goTypes,
		DependencyIndexes: file_google_cloud_discoveryengine_v1alpha_document_proto_depIdxs,
		MessageInfos:      file_google_cloud_discoveryengine_v1alpha_document_proto_msgTypes,
	}.Build()
	File_google_cloud_discoveryengine_v1alpha_document_proto = out.File
	file_google_cloud_discoveryengine_v1alpha_document_proto_rawDesc = nil
	file_google_cloud_discoveryengine_v1alpha_document_proto_goTypes = nil
	file_google_cloud_discoveryengine_v1alpha_document_proto_depIdxs = nil
}

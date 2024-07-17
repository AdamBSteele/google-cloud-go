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
// source: google/shopping/merchant/lfp/v1beta/lfpinventory.proto

package lfppb

import (
	context "context"
	reflect "reflect"
	sync "sync"

	typepb "cloud.google.com/go/shopping/type/typepb"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// Local Inventory for the merchant.
type LfpInventory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. Identifier. The name for the `LfpInventory` resource.
	// Format:
	// `accounts/{account}/lfpInventories/{target_merchant}~{store_code}~{offer}`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. The Merchant Center ID of the merchant to submit the inventory
	// for.
	TargetAccount int64 `protobuf:"varint,2,opt,name=target_account,json=targetAccount,proto3" json:"target_account,omitempty"`
	// Required. The identifier of the merchant's store. Either the store code
	// inserted through `InsertLfpStore` or the store code in the Business
	// Profile.
	StoreCode string `protobuf:"bytes,3,opt,name=store_code,json=storeCode,proto3" json:"store_code,omitempty"`
	// Required. Immutable. A unique identifier for the product. If both
	// inventories and sales are submitted for a merchant, this id should match
	// for the same product.
	//
	// **Note**: if the merchant sells the same product new and used, they should
	// have different IDs.
	OfferId string `protobuf:"bytes,4,opt,name=offer_id,json=offerId,proto3" json:"offer_id,omitempty"`
	// Required. The [CLDR territory
	// code](https://github.com/unicode-org/cldr/blob/latest/common/main/en.xml)
	// for the country where the product is sold.
	RegionCode string `protobuf:"bytes,5,opt,name=region_code,json=regionCode,proto3" json:"region_code,omitempty"`
	// Required. The two-letter ISO 639-1 language code for the item.
	ContentLanguage string `protobuf:"bytes,6,opt,name=content_language,json=contentLanguage,proto3" json:"content_language,omitempty"`
	// Optional. The Global Trade Item Number of the product.
	Gtin *string `protobuf:"bytes,7,opt,name=gtin,proto3,oneof" json:"gtin,omitempty"`
	// Optional. The current price of the product.
	Price *typepb.Price `protobuf:"bytes,8,opt,name=price,proto3" json:"price,omitempty"`
	// Required. Availability of the product at this store.
	// For accepted attribute values, see the [local product inventory data
	// specification](https://support.google.com/merchants/answer/3061342)
	Availability string `protobuf:"bytes,9,opt,name=availability,proto3" json:"availability,omitempty"`
	// Optional. Quantity of the product available at this store. Must be greater
	// than or equal to zero.
	Quantity *int64 `protobuf:"varint,10,opt,name=quantity,proto3,oneof" json:"quantity,omitempty"`
	// Optional. The time when the inventory is collected. If not set, it will be
	// set to the time when the inventory is submitted.
	CollectionTime *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=collection_time,json=collectionTime,proto3" json:"collection_time,omitempty"`
	// Optional. Supported pickup method for this offer. Unless the value is "not
	// supported", this field must be submitted together with `pickupSla`. For
	// accepted attribute values, see the [local product inventory data
	// specification](https://support.google.com/merchants/answer/3061342).
	PickupMethod *string `protobuf:"bytes,12,opt,name=pickup_method,json=pickupMethod,proto3,oneof" json:"pickup_method,omitempty"`
	// Optional. Expected date that an order will be ready for pickup relative to
	// the order date. Must be submitted together with `pickupMethod`. For
	// accepted attribute values, see the [local product inventory data
	// specification](https://support.google.com/merchants/answer/3061342).
	PickupSla *string `protobuf:"bytes,13,opt,name=pickup_sla,json=pickupSla,proto3,oneof" json:"pickup_sla,omitempty"`
	// Optional. The [feed
	// label](https://developers.google.com/shopping-content/guides/products/feed-labels)
	// for the product. If this is not set, it will default to `regionCode`.
	FeedLabel *string `protobuf:"bytes,14,opt,name=feed_label,json=feedLabel,proto3,oneof" json:"feed_label,omitempty"`
}

func (x *LfpInventory) Reset() {
	*x = LfpInventory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_shopping_merchant_lfp_v1beta_lfpinventory_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LfpInventory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LfpInventory) ProtoMessage() {}

func (x *LfpInventory) ProtoReflect() protoreflect.Message {
	mi := &file_google_shopping_merchant_lfp_v1beta_lfpinventory_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LfpInventory.ProtoReflect.Descriptor instead.
func (*LfpInventory) Descriptor() ([]byte, []int) {
	return file_google_shopping_merchant_lfp_v1beta_lfpinventory_proto_rawDescGZIP(), []int{0}
}

func (x *LfpInventory) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LfpInventory) GetTargetAccount() int64 {
	if x != nil {
		return x.TargetAccount
	}
	return 0
}

func (x *LfpInventory) GetStoreCode() string {
	if x != nil {
		return x.StoreCode
	}
	return ""
}

func (x *LfpInventory) GetOfferId() string {
	if x != nil {
		return x.OfferId
	}
	return ""
}

func (x *LfpInventory) GetRegionCode() string {
	if x != nil {
		return x.RegionCode
	}
	return ""
}

func (x *LfpInventory) GetContentLanguage() string {
	if x != nil {
		return x.ContentLanguage
	}
	return ""
}

func (x *LfpInventory) GetGtin() string {
	if x != nil && x.Gtin != nil {
		return *x.Gtin
	}
	return ""
}

func (x *LfpInventory) GetPrice() *typepb.Price {
	if x != nil {
		return x.Price
	}
	return nil
}

func (x *LfpInventory) GetAvailability() string {
	if x != nil {
		return x.Availability
	}
	return ""
}

func (x *LfpInventory) GetQuantity() int64 {
	if x != nil && x.Quantity != nil {
		return *x.Quantity
	}
	return 0
}

func (x *LfpInventory) GetCollectionTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CollectionTime
	}
	return nil
}

func (x *LfpInventory) GetPickupMethod() string {
	if x != nil && x.PickupMethod != nil {
		return *x.PickupMethod
	}
	return ""
}

func (x *LfpInventory) GetPickupSla() string {
	if x != nil && x.PickupSla != nil {
		return *x.PickupSla
	}
	return ""
}

func (x *LfpInventory) GetFeedLabel() string {
	if x != nil && x.FeedLabel != nil {
		return *x.FeedLabel
	}
	return ""
}

// Request message for the `InsertLfpInventory` method.
type InsertLfpInventoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The LFP provider account.
	// Format: `accounts/{account}`
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Required. The inventory to insert.
	LfpInventory *LfpInventory `protobuf:"bytes,2,opt,name=lfp_inventory,json=lfpInventory,proto3" json:"lfp_inventory,omitempty"`
}

func (x *InsertLfpInventoryRequest) Reset() {
	*x = InsertLfpInventoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_shopping_merchant_lfp_v1beta_lfpinventory_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertLfpInventoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertLfpInventoryRequest) ProtoMessage() {}

func (x *InsertLfpInventoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_shopping_merchant_lfp_v1beta_lfpinventory_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertLfpInventoryRequest.ProtoReflect.Descriptor instead.
func (*InsertLfpInventoryRequest) Descriptor() ([]byte, []int) {
	return file_google_shopping_merchant_lfp_v1beta_lfpinventory_proto_rawDescGZIP(), []int{1}
}

func (x *InsertLfpInventoryRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *InsertLfpInventoryRequest) GetLfpInventory() *LfpInventory {
	if x != nil {
		return x.LfpInventory
	}
	return nil
}

var File_google_shopping_merchant_lfp_v1beta_lfpinventory_proto protoreflect.FileDescriptor

var file_google_shopping_merchant_lfp_v1beta_lfpinventory_proto_rawDesc = []byte{
	0x0a, 0x36, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x2f, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2f, 0x6c, 0x66, 0x70, 0x2f, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x6c, 0x66, 0x70, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x23, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61,
	0x6e, 0x74, 0x2e, 0x6c, 0x66, 0x70, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xc1, 0x06, 0x0a, 0x0c, 0x4c, 0x66, 0x70, 0x49, 0x6e, 0x76, 0x65, 0x6e,
	0x74, 0x6f, 0x72, 0x79, 0x12, 0x1a, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x06, 0xe0, 0x41, 0x03, 0xe0, 0x41, 0x08, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x2a, 0x0a, 0x0e, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0d, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0a,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x09, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x21, 0x0a, 0x08, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x06, 0xe0, 0x41, 0x02, 0xe0, 0x41, 0x05, 0x52, 0x07, 0x6f, 0x66, 0x66, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0b, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0a, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x2e, 0x0a, 0x10, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x04, 0x67, 0x74, 0x69,
	0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x48, 0x00, 0x52, 0x04,
	0x67, 0x74, 0x69, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x36, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x27, 0x0a, 0x0c, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0c, 0x61, 0x76, 0x61, 0x69,
	0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x24, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x48,
	0x01, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x88, 0x01, 0x01, 0x12, 0x48,
	0x0a, 0x0f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x0e, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2d, 0x0a, 0x0d, 0x70, 0x69, 0x63, 0x6b,
	0x75, 0x70, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x03, 0xe0, 0x41, 0x01, 0x48, 0x02, 0x52, 0x0c, 0x70, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x88, 0x01, 0x01, 0x12, 0x27, 0x0a, 0x0a, 0x70, 0x69, 0x63, 0x6b, 0x75,
	0x70, 0x5f, 0x73, 0x6c, 0x61, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01,
	0x48, 0x03, 0x52, 0x09, 0x70, 0x69, 0x63, 0x6b, 0x75, 0x70, 0x53, 0x6c, 0x61, 0x88, 0x01, 0x01,
	0x12, 0x27, 0x0a, 0x0a, 0x66, 0x65, 0x65, 0x64, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x48, 0x04, 0x52, 0x09, 0x66, 0x65, 0x65,
	0x64, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x88, 0x01, 0x01, 0x3a, 0x95, 0x01, 0xea, 0x41, 0x91, 0x01,
	0x0a, 0x27, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4c, 0x66, 0x70,
	0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x48, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x7d, 0x2f, 0x6c, 0x66,
	0x70, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x2f, 0x7b, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x5f, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x7d, 0x7e, 0x7b,
	0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x7d, 0x7e, 0x7b, 0x6f, 0x66, 0x66,
	0x65, 0x72, 0x7d, 0x2a, 0x0e, 0x6c, 0x66, 0x70, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72,
	0x69, 0x65, 0x73, 0x32, 0x0c, 0x6c, 0x66, 0x70, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72,
	0x79, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x67, 0x74, 0x69, 0x6e, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x71,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x70, 0x69, 0x63, 0x6b,
	0x75, 0x70, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x70, 0x69,
	0x63, 0x6b, 0x75, 0x70, 0x5f, 0x73, 0x6c, 0x61, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x66, 0x65, 0x65,
	0x64, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x22, 0xc1, 0x01, 0x0a, 0x19, 0x49, 0x6e, 0x73, 0x65,
	0x72, 0x74, 0x4c, 0x66, 0x70, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x47, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2f, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x29, 0x12, 0x27, 0x6d,
	0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4c, 0x66, 0x70, 0x49, 0x6e, 0x76,
	0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x5b,
	0x0a, 0x0d, 0x6c, 0x66, 0x70, 0x5f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73,
	0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74,
	0x2e, 0x6c, 0x66, 0x70, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x4c, 0x66, 0x70, 0x49,
	0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0c, 0x6c,
	0x66, 0x70, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x32, 0xb6, 0x02, 0x0a, 0x13,
	0x4c, 0x66, 0x70, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0xd5, 0x01, 0x0a, 0x12, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x4c, 0x66,
	0x70, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x3e, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x6d, 0x65, 0x72,
	0x63, 0x68, 0x61, 0x6e, 0x74, 0x2e, 0x6c, 0x66, 0x70, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x4c, 0x66, 0x70, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74,
	0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x6d, 0x65, 0x72,
	0x63, 0x68, 0x61, 0x6e, 0x74, 0x2e, 0x6c, 0x66, 0x70, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x2e, 0x4c, 0x66, 0x70, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x22, 0x4c, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x46, 0x3a, 0x0d, 0x6c, 0x66, 0x70, 0x5f, 0x69, 0x6e, 0x76, 0x65, 0x6e,
	0x74, 0x6f, 0x72, 0x79, 0x22, 0x35, 0x2f, 0x6c, 0x66, 0x70, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x6c, 0x66, 0x70, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f,
	0x72, 0x69, 0x65, 0x73, 0x3a, 0x69, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x1a, 0x47, 0xca, 0x41, 0x1a,
	0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0xd2, 0x41, 0x27, 0x68, 0x74, 0x74,
	0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61,
	0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x42, 0xba, 0x01, 0xea, 0x41, 0x38, 0x0a, 0x22, 0x6d, 0x65, 0x72, 0x63,
	0x68, 0x61, 0x6e, 0x74, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2f, 0x7b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x7d, 0x0a, 0x27, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x73,
	0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74,
	0x2e, 0x6c, 0x66, 0x70, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x42, 0x11, 0x4c, 0x66, 0x70,
	0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x3f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2f, 0x6d,
	0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2f, 0x6c, 0x66, 0x70, 0x2f, 0x61, 0x70, 0x69, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x6c, 0x66, 0x70, 0x70, 0x62, 0x3b, 0x6c, 0x66, 0x70, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_shopping_merchant_lfp_v1beta_lfpinventory_proto_rawDescOnce sync.Once
	file_google_shopping_merchant_lfp_v1beta_lfpinventory_proto_rawDescData = file_google_shopping_merchant_lfp_v1beta_lfpinventory_proto_rawDesc
)

func file_google_shopping_merchant_lfp_v1beta_lfpinventory_proto_rawDescGZIP() []byte {
	file_google_shopping_merchant_lfp_v1beta_lfpinventory_proto_rawDescOnce.Do(func() {
		file_google_shopping_merchant_lfp_v1beta_lfpinventory_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_shopping_merchant_lfp_v1beta_lfpinventory_proto_rawDescData)
	})
	return file_google_shopping_merchant_lfp_v1beta_lfpinventory_proto_rawDescData
}

var file_google_shopping_merchant_lfp_v1beta_lfpinventory_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_shopping_merchant_lfp_v1beta_lfpinventory_proto_goTypes = []any{
	(*LfpInventory)(nil),              // 0: google.shopping.merchant.lfp.v1beta.LfpInventory
	(*InsertLfpInventoryRequest)(nil), // 1: google.shopping.merchant.lfp.v1beta.InsertLfpInventoryRequest
	(*typepb.Price)(nil),              // 2: google.shopping.type.Price
	(*timestamppb.Timestamp)(nil),     // 3: google.protobuf.Timestamp
}
var file_google_shopping_merchant_lfp_v1beta_lfpinventory_proto_depIdxs = []int32{
	2, // 0: google.shopping.merchant.lfp.v1beta.LfpInventory.price:type_name -> google.shopping.type.Price
	3, // 1: google.shopping.merchant.lfp.v1beta.LfpInventory.collection_time:type_name -> google.protobuf.Timestamp
	0, // 2: google.shopping.merchant.lfp.v1beta.InsertLfpInventoryRequest.lfp_inventory:type_name -> google.shopping.merchant.lfp.v1beta.LfpInventory
	1, // 3: google.shopping.merchant.lfp.v1beta.LfpInventoryService.InsertLfpInventory:input_type -> google.shopping.merchant.lfp.v1beta.InsertLfpInventoryRequest
	0, // 4: google.shopping.merchant.lfp.v1beta.LfpInventoryService.InsertLfpInventory:output_type -> google.shopping.merchant.lfp.v1beta.LfpInventory
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_google_shopping_merchant_lfp_v1beta_lfpinventory_proto_init() }
func file_google_shopping_merchant_lfp_v1beta_lfpinventory_proto_init() {
	if File_google_shopping_merchant_lfp_v1beta_lfpinventory_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_shopping_merchant_lfp_v1beta_lfpinventory_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*LfpInventory); i {
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
		file_google_shopping_merchant_lfp_v1beta_lfpinventory_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*InsertLfpInventoryRequest); i {
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
	file_google_shopping_merchant_lfp_v1beta_lfpinventory_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_shopping_merchant_lfp_v1beta_lfpinventory_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_shopping_merchant_lfp_v1beta_lfpinventory_proto_goTypes,
		DependencyIndexes: file_google_shopping_merchant_lfp_v1beta_lfpinventory_proto_depIdxs,
		MessageInfos:      file_google_shopping_merchant_lfp_v1beta_lfpinventory_proto_msgTypes,
	}.Build()
	File_google_shopping_merchant_lfp_v1beta_lfpinventory_proto = out.File
	file_google_shopping_merchant_lfp_v1beta_lfpinventory_proto_rawDesc = nil
	file_google_shopping_merchant_lfp_v1beta_lfpinventory_proto_goTypes = nil
	file_google_shopping_merchant_lfp_v1beta_lfpinventory_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LfpInventoryServiceClient is the client API for LfpInventoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LfpInventoryServiceClient interface {
	// Inserts a `LfpInventory` resource for the given target merchant account. If
	// the resource already exists, it will be replaced. The inventory
	// automatically expires after 30 days.
	InsertLfpInventory(ctx context.Context, in *InsertLfpInventoryRequest, opts ...grpc.CallOption) (*LfpInventory, error)
}

type lfpInventoryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLfpInventoryServiceClient(cc grpc.ClientConnInterface) LfpInventoryServiceClient {
	return &lfpInventoryServiceClient{cc}
}

func (c *lfpInventoryServiceClient) InsertLfpInventory(ctx context.Context, in *InsertLfpInventoryRequest, opts ...grpc.CallOption) (*LfpInventory, error) {
	out := new(LfpInventory)
	err := c.cc.Invoke(ctx, "/google.shopping.merchant.lfp.v1beta.LfpInventoryService/InsertLfpInventory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LfpInventoryServiceServer is the server API for LfpInventoryService service.
type LfpInventoryServiceServer interface {
	// Inserts a `LfpInventory` resource for the given target merchant account. If
	// the resource already exists, it will be replaced. The inventory
	// automatically expires after 30 days.
	InsertLfpInventory(context.Context, *InsertLfpInventoryRequest) (*LfpInventory, error)
}

// UnimplementedLfpInventoryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLfpInventoryServiceServer struct {
}

func (*UnimplementedLfpInventoryServiceServer) InsertLfpInventory(context.Context, *InsertLfpInventoryRequest) (*LfpInventory, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertLfpInventory not implemented")
}

func RegisterLfpInventoryServiceServer(s *grpc.Server, srv LfpInventoryServiceServer) {
	s.RegisterService(&_LfpInventoryService_serviceDesc, srv)
}

func _LfpInventoryService_InsertLfpInventory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertLfpInventoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LfpInventoryServiceServer).InsertLfpInventory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.shopping.merchant.lfp.v1beta.LfpInventoryService/InsertLfpInventory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LfpInventoryServiceServer).InsertLfpInventory(ctx, req.(*InsertLfpInventoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LfpInventoryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.shopping.merchant.lfp.v1beta.LfpInventoryService",
	HandlerType: (*LfpInventoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InsertLfpInventory",
			Handler:    _LfpInventoryService_InsertLfpInventory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/shopping/merchant/lfp/v1beta/lfpinventory.proto",
}
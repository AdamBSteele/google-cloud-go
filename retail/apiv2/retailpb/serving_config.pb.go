// Copyright 2021 Google LLC
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
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: google/cloud/retail/v2/serving_config.proto

package retailpb

import (
	reflect "reflect"
	sync "sync"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// What type of diversity - data or rule based.
type ServingConfig_DiversityType int32

const (
	// Default value.
	ServingConfig_DIVERSITY_TYPE_UNSPECIFIED ServingConfig_DiversityType = 0
	// Rule based diversity.
	ServingConfig_RULE_BASED_DIVERSITY ServingConfig_DiversityType = 2
	// Data driven diversity.
	ServingConfig_DATA_DRIVEN_DIVERSITY ServingConfig_DiversityType = 3
)

// Enum value maps for ServingConfig_DiversityType.
var (
	ServingConfig_DiversityType_name = map[int32]string{
		0: "DIVERSITY_TYPE_UNSPECIFIED",
		2: "RULE_BASED_DIVERSITY",
		3: "DATA_DRIVEN_DIVERSITY",
	}
	ServingConfig_DiversityType_value = map[string]int32{
		"DIVERSITY_TYPE_UNSPECIFIED": 0,
		"RULE_BASED_DIVERSITY":       2,
		"DATA_DRIVEN_DIVERSITY":      3,
	}
)

func (x ServingConfig_DiversityType) Enum() *ServingConfig_DiversityType {
	p := new(ServingConfig_DiversityType)
	*p = x
	return p
}

func (x ServingConfig_DiversityType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ServingConfig_DiversityType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_retail_v2_serving_config_proto_enumTypes[0].Descriptor()
}

func (ServingConfig_DiversityType) Type() protoreflect.EnumType {
	return &file_google_cloud_retail_v2_serving_config_proto_enumTypes[0]
}

func (x ServingConfig_DiversityType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ServingConfig_DiversityType.Descriptor instead.
func (ServingConfig_DiversityType) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_retail_v2_serving_config_proto_rawDescGZIP(), []int{0, 0}
}

// Configures metadata that is used to generate serving time results (e.g.
// search results or recommendation predictions).
type ServingConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Immutable. Fully qualified name
	// `projects/*/locations/global/catalogs/*/servingConfig/*`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. The human readable serving config display name. Used in Retail
	// UI.
	//
	// This field must be a UTF-8 encoded string with a length limit of 128
	// characters. Otherwise, an INVALID_ARGUMENT error is returned.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// The id of the model in the same [Catalog][google.cloud.retail.v2.Catalog]
	// to use at serving time. Currently only RecommendationModels are supported:
	// https://cloud.google.com/retail/recommendations-ai/docs/create-models
	// Can be changed but only to a compatible model (e.g.
	// others-you-may-like CTR to others-you-may-like CVR).
	//
	// Required when
	// [solution_types][google.cloud.retail.v2.ServingConfig.solution_types] is
	// [SOLUTION_TYPE_RECOMMENDATION][google.cloud.retail.v2main.SolutionType.SOLUTION_TYPE_RECOMMENDATION].
	ModelId string `protobuf:"bytes,3,opt,name=model_id,json=modelId,proto3" json:"model_id,omitempty"`
	// How much price ranking we want in serving results.
	// Price reranking causes product items with a similar
	// recommendation probability to be ordered by price, with the
	// highest-priced items first. This setting could result in a decrease in
	// click-through and conversion rates.
	//
	//	Allowed values are:
	//
	// * `no-price-reranking`
	// * `low-price-reranking`
	// * `medium-price-reranking`
	// * `high-price-reranking`
	//
	// If not specified, we choose default based on model type. Default value:
	// `no-price-reranking`.
	//
	// Can only be set if
	// [solution_types][google.cloud.retail.v2.ServingConfig.solution_types] is
	// [SOLUTION_TYPE_RECOMMENDATION][google.cloud.retail.v2main.SolutionType.SOLUTION_TYPE_RECOMMENDATION].
	PriceRerankingLevel string `protobuf:"bytes,4,opt,name=price_reranking_level,json=priceRerankingLevel,proto3" json:"price_reranking_level,omitempty"`
	// Facet specifications for faceted search. If empty, no facets are returned.
	// The ids refer to the ids of [Control][google.cloud.retail.v2.Control]
	// resources with only the Facet control set. These controls are assumed to be
	// in the same [Catalog][google.cloud.retail.v2.Catalog] as the
	// [ServingConfig][google.cloud.retail.v2.ServingConfig].
	// A maximum of 100 values are allowed. Otherwise, an INVALID_ARGUMENT error
	// is returned.
	//
	// Can only be set if
	// [solution_types][google.cloud.retail.v2.ServingConfig.solution_types] is
	// [SOLUTION_TYPE_SEARCH][google.cloud.retail.v2main.SolutionType.SOLUTION_TYPE_SEARCH].
	FacetControlIds []string `protobuf:"bytes,5,rep,name=facet_control_ids,json=facetControlIds,proto3" json:"facet_control_ids,omitempty"`
	// The specification for dynamically generated facets. Notice that only
	// textual facets can be dynamically generated.
	//
	// Can only be set if
	// [solution_types][google.cloud.retail.v2.ServingConfig.solution_types] is
	// [SOLUTION_TYPE_SEARCH][google.cloud.retail.v2main.SolutionType.SOLUTION_TYPE_SEARCH].
	DynamicFacetSpec *SearchRequest_DynamicFacetSpec `protobuf:"bytes,6,opt,name=dynamic_facet_spec,json=dynamicFacetSpec,proto3" json:"dynamic_facet_spec,omitempty"`
	// Condition boost specifications. If a product matches multiple conditions
	// in the specifications, boost scores from these specifications are all
	// applied and combined in a non-linear way. Maximum number of
	// specifications is 100.
	//
	// Notice that if both
	// [ServingConfig.boost_control_ids][google.cloud.retail.v2.ServingConfig.boost_control_ids]
	// and
	// [SearchRequest.boost_spec][google.cloud.retail.v2.SearchRequest.boost_spec]
	// are set, the boost conditions from both places are evaluated. If a search
	// request matches multiple boost conditions, the final boost score is equal
	// to the sum of the boost scores from all matched boost conditions.
	//
	// Can only be set if
	// [solution_types][google.cloud.retail.v2.ServingConfig.solution_types] is
	// [SOLUTION_TYPE_SEARCH][google.cloud.retail.v2main.SolutionType.SOLUTION_TYPE_SEARCH].
	BoostControlIds []string `protobuf:"bytes,7,rep,name=boost_control_ids,json=boostControlIds,proto3" json:"boost_control_ids,omitempty"`
	// Condition filter specifications. If a product matches multiple conditions
	// in the specifications, filters from these specifications are all
	// applied and combined via the AND operator. Maximum number of
	// specifications is 100.
	//
	// Can only be set if
	// [solution_types][google.cloud.retail.v2.ServingConfig.solution_types] is
	// [SOLUTION_TYPE_SEARCH][google.cloud.retail.v2main.SolutionType.SOLUTION_TYPE_SEARCH].
	FilterControlIds []string `protobuf:"bytes,9,rep,name=filter_control_ids,json=filterControlIds,proto3" json:"filter_control_ids,omitempty"`
	// Condition redirect specifications. Only the first triggered redirect action
	// is applied, even if multiple apply. Maximum number of specifications is
	// 1000.
	//
	// Can only be set if
	// [solution_types][google.cloud.retail.v2.ServingConfig.solution_types] is
	// [SOLUTION_TYPE_SEARCH][google.cloud.retail.v2main.SolutionType.SOLUTION_TYPE_SEARCH].
	RedirectControlIds []string `protobuf:"bytes,10,rep,name=redirect_control_ids,json=redirectControlIds,proto3" json:"redirect_control_ids,omitempty"`
	// Condition synonyms specifications. If multiple syonyms conditions match,
	// all matching synonyms control in the list will execute. Order of controls
	// in the list will not matter. Maximum number of specifications is
	// 100.
	//
	// Can only be set if
	// [solution_types][google.cloud.retail.v2.ServingConfig.solution_types] is
	// [SOLUTION_TYPE_SEARCH][google.cloud.retail.v2main.SolutionType.SOLUTION_TYPE_SEARCH].
	TwowaySynonymsControlIds []string `protobuf:"bytes,18,rep,name=twoway_synonyms_control_ids,json=twowaySynonymsControlIds,proto3" json:"twoway_synonyms_control_ids,omitempty"`
	// Condition oneway synonyms specifications. If multiple oneway synonyms
	// conditions match, all matching oneway synonyms controls in the list will
	// execute. Order of controls in the list will not matter. Maximum number of
	// specifications is 100.
	//
	// Can only be set if
	// [solution_types][google.cloud.retail.v2.ServingConfig.solution_types] is
	// [SOLUTION_TYPE_SEARCH][google.cloud.retail.v2main.SolutionType.SOLUTION_TYPE_SEARCH].
	OnewaySynonymsControlIds []string `protobuf:"bytes,12,rep,name=oneway_synonyms_control_ids,json=onewaySynonymsControlIds,proto3" json:"oneway_synonyms_control_ids,omitempty"`
	// Condition do not associate specifications. If multiple do not associate
	// conditions match, all matching do not associate controls in the list will
	// execute.
	// - Order does not matter.
	// - Maximum number of specifications is 100.
	//
	// Can only be set if
	// [solution_types][google.cloud.retail.v2.ServingConfig.solution_types] is
	// [SOLUTION_TYPE_SEARCH][google.cloud.retail.v2main.SolutionType.SOLUTION_TYPE_SEARCH].
	DoNotAssociateControlIds []string `protobuf:"bytes,13,rep,name=do_not_associate_control_ids,json=doNotAssociateControlIds,proto3" json:"do_not_associate_control_ids,omitempty"`
	// Condition replacement specifications.
	// - Applied according to the order in the list.
	// - A previously replaced term can not be re-replaced.
	// - Maximum number of specifications is 100.
	//
	// Can only be set if
	// [solution_types][google.cloud.retail.v2.ServingConfig.solution_types] is
	// [SOLUTION_TYPE_SEARCH][google.cloud.retail.v2main.SolutionType.SOLUTION_TYPE_SEARCH].
	ReplacementControlIds []string `protobuf:"bytes,14,rep,name=replacement_control_ids,json=replacementControlIds,proto3" json:"replacement_control_ids,omitempty"`
	// Condition ignore specifications. If multiple ignore
	// conditions match, all matching ignore controls in the list will
	// execute.
	// - Order does not matter.
	// - Maximum number of specifications is 100.
	//
	// Can only be set if
	// [solution_types][google.cloud.retail.v2.ServingConfig.solution_types] is
	// [SOLUTION_TYPE_SEARCH][google.cloud.retail.v2main.SolutionType.SOLUTION_TYPE_SEARCH].
	IgnoreControlIds []string `protobuf:"bytes,15,rep,name=ignore_control_ids,json=ignoreControlIds,proto3" json:"ignore_control_ids,omitempty"`
	// How much diversity to use in recommendation model results e.g.
	// `medium-diversity` or `high-diversity`. Currently supported values:
	//
	// * `no-diversity`
	// * `low-diversity`
	// * `medium-diversity`
	// * `high-diversity`
	// * `auto-diversity`
	//
	// If not specified, we choose default based on recommendation model
	// type. Default value: `no-diversity`.
	//
	// Can only be set if
	// [solution_types][google.cloud.retail.v2.ServingConfig.solution_types] is
	// [SOLUTION_TYPE_RECOMMENDATION][google.cloud.retail.v2main.SolutionType.SOLUTION_TYPE_RECOMMENDATION].
	DiversityLevel string `protobuf:"bytes,8,opt,name=diversity_level,json=diversityLevel,proto3" json:"diversity_level,omitempty"`
	// What kind of diversity to use - data driven or rule based. If unset, the
	// server behavior defaults to
	// [RULE_BASED_DIVERSITY][google.cloud.retail.v2.ServingConfig.DiversityType.RULE_BASED_DIVERSITY].
	DiversityType ServingConfig_DiversityType `protobuf:"varint,20,opt,name=diversity_type,json=diversityType,proto3,enum=google.cloud.retail.v2.ServingConfig_DiversityType" json:"diversity_type,omitempty"`
	// Whether to add additional category filters on the `similar-items` model.
	// If not specified, we enable it by default.
	//
	//	Allowed values are:
	//
	//   - `no-category-match`: No additional filtering of original results from
	//     the model and the customer's filters.
	//   - `relaxed-category-match`: Only keep results with categories that match
	//     at least one item categories in the PredictRequests's context item.
	//   - If customer also sends filters in the PredictRequest, then the results
	//     will satisfy both conditions (user given and category match).
	//
	// Can only be set if
	// [solution_types][google.cloud.retail.v2.ServingConfig.solution_types] is
	// [SOLUTION_TYPE_RECOMMENDATION][google.cloud.retail.v2main.SolutionType.SOLUTION_TYPE_RECOMMENDATION].
	EnableCategoryFilterLevel string `protobuf:"bytes,16,opt,name=enable_category_filter_level,json=enableCategoryFilterLevel,proto3" json:"enable_category_filter_level,omitempty"`
	// The specification for personalization spec.
	//
	// Can only be set if
	// [solution_types][google.cloud.retail.v2.ServingConfig.solution_types] is
	// [SOLUTION_TYPE_SEARCH][google.cloud.retail.v2main.SolutionType.SOLUTION_TYPE_SEARCH].
	//
	// Notice that if both
	// [ServingConfig.personalization_spec][google.cloud.retail.v2.ServingConfig.personalization_spec]
	// and
	// [SearchRequest.personalization_spec][google.cloud.retail.v2.SearchRequest.personalization_spec]
	// are set.
	// [SearchRequest.personalization_spec][google.cloud.retail.v2.SearchRequest.personalization_spec]
	// will override
	// [ServingConfig.personalization_spec][google.cloud.retail.v2.ServingConfig.personalization_spec].
	PersonalizationSpec *SearchRequest_PersonalizationSpec `protobuf:"bytes,21,opt,name=personalization_spec,json=personalizationSpec,proto3" json:"personalization_spec,omitempty"`
	// Required. Immutable. Specifies the solution types that a serving config can
	// be associated with. Currently we support setting only one type of solution.
	SolutionTypes []SolutionType `protobuf:"varint,19,rep,packed,name=solution_types,json=solutionTypes,proto3,enum=google.cloud.retail.v2.SolutionType" json:"solution_types,omitempty"`
}

func (x *ServingConfig) Reset() {
	*x = ServingConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_retail_v2_serving_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServingConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServingConfig) ProtoMessage() {}

func (x *ServingConfig) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_retail_v2_serving_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServingConfig.ProtoReflect.Descriptor instead.
func (*ServingConfig) Descriptor() ([]byte, []int) {
	return file_google_cloud_retail_v2_serving_config_proto_rawDescGZIP(), []int{0}
}

func (x *ServingConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ServingConfig) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *ServingConfig) GetModelId() string {
	if x != nil {
		return x.ModelId
	}
	return ""
}

func (x *ServingConfig) GetPriceRerankingLevel() string {
	if x != nil {
		return x.PriceRerankingLevel
	}
	return ""
}

func (x *ServingConfig) GetFacetControlIds() []string {
	if x != nil {
		return x.FacetControlIds
	}
	return nil
}

func (x *ServingConfig) GetDynamicFacetSpec() *SearchRequest_DynamicFacetSpec {
	if x != nil {
		return x.DynamicFacetSpec
	}
	return nil
}

func (x *ServingConfig) GetBoostControlIds() []string {
	if x != nil {
		return x.BoostControlIds
	}
	return nil
}

func (x *ServingConfig) GetFilterControlIds() []string {
	if x != nil {
		return x.FilterControlIds
	}
	return nil
}

func (x *ServingConfig) GetRedirectControlIds() []string {
	if x != nil {
		return x.RedirectControlIds
	}
	return nil
}

func (x *ServingConfig) GetTwowaySynonymsControlIds() []string {
	if x != nil {
		return x.TwowaySynonymsControlIds
	}
	return nil
}

func (x *ServingConfig) GetOnewaySynonymsControlIds() []string {
	if x != nil {
		return x.OnewaySynonymsControlIds
	}
	return nil
}

func (x *ServingConfig) GetDoNotAssociateControlIds() []string {
	if x != nil {
		return x.DoNotAssociateControlIds
	}
	return nil
}

func (x *ServingConfig) GetReplacementControlIds() []string {
	if x != nil {
		return x.ReplacementControlIds
	}
	return nil
}

func (x *ServingConfig) GetIgnoreControlIds() []string {
	if x != nil {
		return x.IgnoreControlIds
	}
	return nil
}

func (x *ServingConfig) GetDiversityLevel() string {
	if x != nil {
		return x.DiversityLevel
	}
	return ""
}

func (x *ServingConfig) GetDiversityType() ServingConfig_DiversityType {
	if x != nil {
		return x.DiversityType
	}
	return ServingConfig_DIVERSITY_TYPE_UNSPECIFIED
}

func (x *ServingConfig) GetEnableCategoryFilterLevel() string {
	if x != nil {
		return x.EnableCategoryFilterLevel
	}
	return ""
}

func (x *ServingConfig) GetPersonalizationSpec() *SearchRequest_PersonalizationSpec {
	if x != nil {
		return x.PersonalizationSpec
	}
	return nil
}

func (x *ServingConfig) GetSolutionTypes() []SolutionType {
	if x != nil {
		return x.SolutionTypes
	}
	return nil
}

var File_google_cloud_retail_v2_serving_config_proto protoreflect.FileDescriptor

var file_google_cloud_retail_v2_serving_config_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x2e, 0x76, 0x32, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x23, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2f, 0x76, 0x32, 0x2f, 0x73,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xd8, 0x0a, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x05, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x26,
	0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x49,
	0x64, 0x12, 0x32, 0x0a, 0x15, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x72, 0x65, 0x72, 0x61, 0x6e,
	0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x13, 0x70, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x72, 0x61, 0x6e, 0x6b, 0x69, 0x6e, 0x67,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x2a, 0x0a, 0x11, 0x66, 0x61, 0x63, 0x65, 0x74, 0x5f, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0f, 0x66, 0x61, 0x63, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x49, 0x64,
	0x73, 0x12, 0x64, 0x0a, 0x12, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x5f, 0x66, 0x61, 0x63,
	0x65, 0x74, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x46, 0x61, 0x63, 0x65,
	0x74, 0x53, 0x70, 0x65, 0x63, 0x52, 0x10, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x69, 0x63, 0x46, 0x61,
	0x63, 0x65, 0x74, 0x53, 0x70, 0x65, 0x63, 0x12, 0x2a, 0x0a, 0x11, 0x62, 0x6f, 0x6f, 0x73, 0x74,
	0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x07, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0f, 0x62, 0x6f, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x49, 0x64, 0x73, 0x12, 0x2c, 0x0a, 0x12, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x10, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x49, 0x64,
	0x73, 0x12, 0x30, 0x0a, 0x14, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x12, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x49, 0x64, 0x73, 0x12, 0x3d, 0x0a, 0x1b, 0x74, 0x77, 0x6f, 0x77, 0x61, 0x79, 0x5f, 0x73, 0x79,
	0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x69,
	0x64, 0x73, 0x18, 0x12, 0x20, 0x03, 0x28, 0x09, 0x52, 0x18, 0x74, 0x77, 0x6f, 0x77, 0x61, 0x79,
	0x53, 0x79, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x49,
	0x64, 0x73, 0x12, 0x3d, 0x0a, 0x1b, 0x6f, 0x6e, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x73, 0x79, 0x6e,
	0x6f, 0x6e, 0x79, 0x6d, 0x73, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x69, 0x64,
	0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x09, 0x52, 0x18, 0x6f, 0x6e, 0x65, 0x77, 0x61, 0x79, 0x53,
	0x79, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x49, 0x64,
	0x73, 0x12, 0x3e, 0x0a, 0x1c, 0x64, 0x6f, 0x5f, 0x6e, 0x6f, 0x74, 0x5f, 0x61, 0x73, 0x73, 0x6f,
	0x63, 0x69, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x69, 0x64,
	0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x09, 0x52, 0x18, 0x64, 0x6f, 0x4e, 0x6f, 0x74, 0x41, 0x73,
	0x73, 0x6f, 0x63, 0x69, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x49, 0x64,
	0x73, 0x12, 0x36, 0x0a, 0x17, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x0e, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x15, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x49, 0x64, 0x73, 0x12, 0x2c, 0x0a, 0x12, 0x69, 0x67, 0x6e,
	0x6f, 0x72, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x69, 0x64, 0x73, 0x18,
	0x0f, 0x20, 0x03, 0x28, 0x09, 0x52, 0x10, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x49, 0x64, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x64, 0x69, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x74, 0x79, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x64, 0x69, 0x76, 0x65, 0x72, 0x73, 0x69, 0x74, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x12, 0x5a, 0x0a, 0x0e, 0x64, 0x69, 0x76, 0x65, 0x72, 0x73, 0x69, 0x74, 0x79, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x33, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x76,
	0x32, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x44, 0x69, 0x76, 0x65, 0x72, 0x73, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0d, 0x64,
	0x69, 0x76, 0x65, 0x72, 0x73, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3f, 0x0a, 0x1c,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x10, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x19, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x6c, 0x0a,
	0x14, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x73, 0x70, 0x65, 0x63, 0x18, 0x15, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x52, 0x13, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x12, 0x53, 0x0a, 0x0e, 0x73,
	0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x18, 0x13, 0x20,
	0x03, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x6f, 0x6c,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x42, 0x06, 0xe0, 0x41, 0x02, 0xe0, 0x41,
	0x05, 0x52, 0x0d, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x73,
	0x22, 0x64, 0x0a, 0x0d, 0x44, 0x69, 0x76, 0x65, 0x72, 0x73, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x1e, 0x0a, 0x1a, 0x44, 0x49, 0x56, 0x45, 0x52, 0x53, 0x49, 0x54, 0x59, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x18, 0x0a, 0x14, 0x52, 0x55, 0x4c, 0x45, 0x5f, 0x42, 0x41, 0x53, 0x45, 0x44, 0x5f,
	0x44, 0x49, 0x56, 0x45, 0x52, 0x53, 0x49, 0x54, 0x59, 0x10, 0x02, 0x12, 0x19, 0x0a, 0x15, 0x44,
	0x41, 0x54, 0x41, 0x5f, 0x44, 0x52, 0x49, 0x56, 0x45, 0x4e, 0x5f, 0x44, 0x49, 0x56, 0x45, 0x52,
	0x53, 0x49, 0x54, 0x59, 0x10, 0x03, 0x3a, 0x85, 0x01, 0xea, 0x41, 0x81, 0x01, 0x0a, 0x23, 0x72,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x5a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x63, 0x61, 0x74, 0x61,
	0x6c, 0x6f, 0x67, 0x73, 0x2f, 0x7b, 0x63, 0x61, 0x74, 0x61, 0x6c, 0x6f, 0x67, 0x7d, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x2f, 0x7b, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x7d, 0x42, 0xbd,
	0x01, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x32, 0x42, 0x12, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x32, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2f,
	0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x72, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x70, 0x62, 0x3b, 0x72,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x70, 0x62, 0xa2, 0x02, 0x06, 0x52, 0x45, 0x54, 0x41, 0x49, 0x4c,
	0xaa, 0x02, 0x16, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x52, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x2e, 0x56, 0x32, 0xca, 0x02, 0x16, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x52, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x5c,
	0x56, 0x32, 0xea, 0x02, 0x19, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x3a, 0x3a, 0x52, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x3a, 0x3a, 0x56, 0x32, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_retail_v2_serving_config_proto_rawDescOnce sync.Once
	file_google_cloud_retail_v2_serving_config_proto_rawDescData = file_google_cloud_retail_v2_serving_config_proto_rawDesc
)

func file_google_cloud_retail_v2_serving_config_proto_rawDescGZIP() []byte {
	file_google_cloud_retail_v2_serving_config_proto_rawDescOnce.Do(func() {
		file_google_cloud_retail_v2_serving_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_retail_v2_serving_config_proto_rawDescData)
	})
	return file_google_cloud_retail_v2_serving_config_proto_rawDescData
}

var file_google_cloud_retail_v2_serving_config_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_retail_v2_serving_config_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_cloud_retail_v2_serving_config_proto_goTypes = []interface{}{
	(ServingConfig_DiversityType)(0),          // 0: google.cloud.retail.v2.ServingConfig.DiversityType
	(*ServingConfig)(nil),                     // 1: google.cloud.retail.v2.ServingConfig
	(*SearchRequest_DynamicFacetSpec)(nil),    // 2: google.cloud.retail.v2.SearchRequest.DynamicFacetSpec
	(*SearchRequest_PersonalizationSpec)(nil), // 3: google.cloud.retail.v2.SearchRequest.PersonalizationSpec
	(SolutionType)(0),                         // 4: google.cloud.retail.v2.SolutionType
}
var file_google_cloud_retail_v2_serving_config_proto_depIdxs = []int32{
	2, // 0: google.cloud.retail.v2.ServingConfig.dynamic_facet_spec:type_name -> google.cloud.retail.v2.SearchRequest.DynamicFacetSpec
	0, // 1: google.cloud.retail.v2.ServingConfig.diversity_type:type_name -> google.cloud.retail.v2.ServingConfig.DiversityType
	3, // 2: google.cloud.retail.v2.ServingConfig.personalization_spec:type_name -> google.cloud.retail.v2.SearchRequest.PersonalizationSpec
	4, // 3: google.cloud.retail.v2.ServingConfig.solution_types:type_name -> google.cloud.retail.v2.SolutionType
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_google_cloud_retail_v2_serving_config_proto_init() }
func file_google_cloud_retail_v2_serving_config_proto_init() {
	if File_google_cloud_retail_v2_serving_config_proto != nil {
		return
	}
	file_google_cloud_retail_v2_common_proto_init()
	file_google_cloud_retail_v2_search_service_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_retail_v2_serving_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServingConfig); i {
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
			RawDescriptor: file_google_cloud_retail_v2_serving_config_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_retail_v2_serving_config_proto_goTypes,
		DependencyIndexes: file_google_cloud_retail_v2_serving_config_proto_depIdxs,
		EnumInfos:         file_google_cloud_retail_v2_serving_config_proto_enumTypes,
		MessageInfos:      file_google_cloud_retail_v2_serving_config_proto_msgTypes,
	}.Build()
	File_google_cloud_retail_v2_serving_config_proto = out.File
	file_google_cloud_retail_v2_serving_config_proto_rawDesc = nil
	file_google_cloud_retail_v2_serving_config_proto_goTypes = nil
	file_google_cloud_retail_v2_serving_config_proto_depIdxs = nil
}

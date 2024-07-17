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
// source: google/cloud/speech/v1p1beta1/resource.proto

package speechpb

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

// A set of words or phrases that represents a common concept likely to appear
// in your audio, for example a list of passenger ship names. CustomClass items
// can be substituted into placeholders that you set in PhraseSet phrases.
type CustomClass struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource name of the custom class.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// If this custom class is a resource, the custom_class_id is the resource id
	// of the CustomClass. Case sensitive.
	CustomClassId string `protobuf:"bytes,2,opt,name=custom_class_id,json=customClassId,proto3" json:"custom_class_id,omitempty"`
	// A collection of class items.
	Items []*CustomClass_ClassItem `protobuf:"bytes,3,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *CustomClass) Reset() {
	*x = CustomClass{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_speech_v1p1beta1_resource_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomClass) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomClass) ProtoMessage() {}

func (x *CustomClass) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_speech_v1p1beta1_resource_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomClass.ProtoReflect.Descriptor instead.
func (*CustomClass) Descriptor() ([]byte, []int) {
	return file_google_cloud_speech_v1p1beta1_resource_proto_rawDescGZIP(), []int{0}
}

func (x *CustomClass) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CustomClass) GetCustomClassId() string {
	if x != nil {
		return x.CustomClassId
	}
	return ""
}

func (x *CustomClass) GetItems() []*CustomClass_ClassItem {
	if x != nil {
		return x.Items
	}
	return nil
}

// Provides "hints" to the speech recognizer to favor specific words and phrases
// in the results.
type PhraseSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource name of the phrase set.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// A list of word and phrases.
	Phrases []*PhraseSet_Phrase `protobuf:"bytes,2,rep,name=phrases,proto3" json:"phrases,omitempty"`
	// Hint Boost. Positive value will increase the probability that a specific
	// phrase will be recognized over other similar sounding phrases. The higher
	// the boost, the higher the chance of false positive recognition as well.
	// Negative boost values would correspond to anti-biasing. Anti-biasing is not
	// enabled, so negative boost will simply be ignored. Though `boost` can
	// accept a wide range of positive values, most use cases are best served with
	// values between 0 (exclusive) and 20. We recommend using a binary search
	// approach to finding the optimal value for your use case as well as adding
	// phrases both with and without boost to your requests.
	Boost float32 `protobuf:"fixed32,4,opt,name=boost,proto3" json:"boost,omitempty"`
}

func (x *PhraseSet) Reset() {
	*x = PhraseSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_speech_v1p1beta1_resource_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PhraseSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PhraseSet) ProtoMessage() {}

func (x *PhraseSet) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_speech_v1p1beta1_resource_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PhraseSet.ProtoReflect.Descriptor instead.
func (*PhraseSet) Descriptor() ([]byte, []int) {
	return file_google_cloud_speech_v1p1beta1_resource_proto_rawDescGZIP(), []int{1}
}

func (x *PhraseSet) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PhraseSet) GetPhrases() []*PhraseSet_Phrase {
	if x != nil {
		return x.Phrases
	}
	return nil
}

func (x *PhraseSet) GetBoost() float32 {
	if x != nil {
		return x.Boost
	}
	return 0
}

// Speech adaptation configuration.
type SpeechAdaptation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A collection of phrase sets. To specify the hints inline, leave the
	// phrase set's `name` blank and fill in the rest of its fields. Any
	// phrase set can use any custom class.
	PhraseSets []*PhraseSet `protobuf:"bytes,1,rep,name=phrase_sets,json=phraseSets,proto3" json:"phrase_sets,omitempty"`
	// A collection of phrase set resource names to use.
	PhraseSetReferences []string `protobuf:"bytes,2,rep,name=phrase_set_references,json=phraseSetReferences,proto3" json:"phrase_set_references,omitempty"`
	// A collection of custom classes. To specify the classes inline, leave the
	// class' `name` blank and fill in the rest of its fields, giving it a unique
	// `custom_class_id`. Refer to the inline defined class in phrase hints by its
	// `custom_class_id`.
	CustomClasses []*CustomClass `protobuf:"bytes,3,rep,name=custom_classes,json=customClasses,proto3" json:"custom_classes,omitempty"`
	// Augmented Backus-Naur form (ABNF) is a standardized grammar notation
	// comprised by a set of derivation rules.
	// See specifications: https://www.w3.org/TR/speech-grammar
	AbnfGrammar *SpeechAdaptation_ABNFGrammar `protobuf:"bytes,4,opt,name=abnf_grammar,json=abnfGrammar,proto3" json:"abnf_grammar,omitempty"`
}

func (x *SpeechAdaptation) Reset() {
	*x = SpeechAdaptation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_speech_v1p1beta1_resource_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpeechAdaptation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpeechAdaptation) ProtoMessage() {}

func (x *SpeechAdaptation) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_speech_v1p1beta1_resource_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpeechAdaptation.ProtoReflect.Descriptor instead.
func (*SpeechAdaptation) Descriptor() ([]byte, []int) {
	return file_google_cloud_speech_v1p1beta1_resource_proto_rawDescGZIP(), []int{2}
}

func (x *SpeechAdaptation) GetPhraseSets() []*PhraseSet {
	if x != nil {
		return x.PhraseSets
	}
	return nil
}

func (x *SpeechAdaptation) GetPhraseSetReferences() []string {
	if x != nil {
		return x.PhraseSetReferences
	}
	return nil
}

func (x *SpeechAdaptation) GetCustomClasses() []*CustomClass {
	if x != nil {
		return x.CustomClasses
	}
	return nil
}

func (x *SpeechAdaptation) GetAbnfGrammar() *SpeechAdaptation_ABNFGrammar {
	if x != nil {
		return x.AbnfGrammar
	}
	return nil
}

// Transcription normalization configuration. Use transcription normalization
// to automatically replace parts of the transcript with phrases of your
// choosing. For StreamingRecognize, this normalization only applies to stable
// partial transcripts (stability > 0.8) and final transcripts.
type TranscriptNormalization struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A list of replacement entries. We will perform replacement with one entry
	// at a time. For example, the second entry in ["cat" => "dog", "mountain cat"
	// => "mountain dog"] will never be applied because we will always process the
	// first entry before it. At most 100 entries.
	Entries []*TranscriptNormalization_Entry `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
}

func (x *TranscriptNormalization) Reset() {
	*x = TranscriptNormalization{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_speech_v1p1beta1_resource_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TranscriptNormalization) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TranscriptNormalization) ProtoMessage() {}

func (x *TranscriptNormalization) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_speech_v1p1beta1_resource_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TranscriptNormalization.ProtoReflect.Descriptor instead.
func (*TranscriptNormalization) Descriptor() ([]byte, []int) {
	return file_google_cloud_speech_v1p1beta1_resource_proto_rawDescGZIP(), []int{3}
}

func (x *TranscriptNormalization) GetEntries() []*TranscriptNormalization_Entry {
	if x != nil {
		return x.Entries
	}
	return nil
}

// An item of the class.
type CustomClass_ClassItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The class item's value.
	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *CustomClass_ClassItem) Reset() {
	*x = CustomClass_ClassItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_speech_v1p1beta1_resource_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomClass_ClassItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomClass_ClassItem) ProtoMessage() {}

func (x *CustomClass_ClassItem) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_speech_v1p1beta1_resource_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomClass_ClassItem.ProtoReflect.Descriptor instead.
func (*CustomClass_ClassItem) Descriptor() ([]byte, []int) {
	return file_google_cloud_speech_v1p1beta1_resource_proto_rawDescGZIP(), []int{0, 0}
}

func (x *CustomClass_ClassItem) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// A phrases containing words and phrase "hints" so that
// the speech recognition is more likely to recognize them. This can be used
// to improve the accuracy for specific words and phrases, for example, if
// specific commands are typically spoken by the user. This can also be used
// to add additional words to the vocabulary of the recognizer. See
// [usage limits](https://cloud.google.com/speech-to-text/quotas#content).
//
// List items can also include pre-built or custom classes containing groups
// of words that represent common concepts that occur in natural language. For
// example, rather than providing a phrase hint for every month of the
// year (e.g. "i was born in january", "i was born in febuary", ...), use the
// pre-built `$MONTH` class improves the likelihood of correctly transcribing
// audio that includes months (e.g. "i was born in $month").
// To refer to pre-built classes, use the class' symbol prepended with `$`
// e.g. `$MONTH`. To refer to custom classes that were defined inline in the
// request, set the class's `custom_class_id` to a string unique to all class
// resources and inline classes. Then use the class' id wrapped in $`{...}`
// e.g. "${my-months}". To refer to custom classes resources, use the class'
// id wrapped in `${}` (e.g. `${my-months}`).
//
// Speech-to-Text supports three locations: `global`, `us` (US North America),
// and `eu` (Europe). If you are calling the `speech.googleapis.com`
// endpoint, use the `global` location. To specify a region, use a
// [regional endpoint](https://cloud.google.com/speech-to-text/docs/endpoints)
// with matching `us` or `eu` location value.
type PhraseSet_Phrase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The phrase itself.
	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	// Hint Boost. Overrides the boost set at the phrase set level.
	// Positive value will increase the probability that a specific phrase will
	// be recognized over other similar sounding phrases. The higher the boost,
	// the higher the chance of false positive recognition as well. Negative
	// boost will simply be ignored. Though `boost` can accept a wide range of
	// positive values, most use cases are best served
	// with values between 0 and 20. We recommend using a binary search approach
	// to finding the optimal value for your use case as well as adding
	// phrases both with and without boost to your requests.
	Boost float32 `protobuf:"fixed32,2,opt,name=boost,proto3" json:"boost,omitempty"`
}

func (x *PhraseSet_Phrase) Reset() {
	*x = PhraseSet_Phrase{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_speech_v1p1beta1_resource_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PhraseSet_Phrase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PhraseSet_Phrase) ProtoMessage() {}

func (x *PhraseSet_Phrase) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_speech_v1p1beta1_resource_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PhraseSet_Phrase.ProtoReflect.Descriptor instead.
func (*PhraseSet_Phrase) Descriptor() ([]byte, []int) {
	return file_google_cloud_speech_v1p1beta1_resource_proto_rawDescGZIP(), []int{1, 0}
}

func (x *PhraseSet_Phrase) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *PhraseSet_Phrase) GetBoost() float32 {
	if x != nil {
		return x.Boost
	}
	return 0
}

type SpeechAdaptation_ABNFGrammar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// All declarations and rules of an ABNF grammar broken up into multiple
	// strings that will end up concatenated.
	AbnfStrings []string `protobuf:"bytes,1,rep,name=abnf_strings,json=abnfStrings,proto3" json:"abnf_strings,omitempty"`
}

func (x *SpeechAdaptation_ABNFGrammar) Reset() {
	*x = SpeechAdaptation_ABNFGrammar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_speech_v1p1beta1_resource_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpeechAdaptation_ABNFGrammar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpeechAdaptation_ABNFGrammar) ProtoMessage() {}

func (x *SpeechAdaptation_ABNFGrammar) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_speech_v1p1beta1_resource_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpeechAdaptation_ABNFGrammar.ProtoReflect.Descriptor instead.
func (*SpeechAdaptation_ABNFGrammar) Descriptor() ([]byte, []int) {
	return file_google_cloud_speech_v1p1beta1_resource_proto_rawDescGZIP(), []int{2, 0}
}

func (x *SpeechAdaptation_ABNFGrammar) GetAbnfStrings() []string {
	if x != nil {
		return x.AbnfStrings
	}
	return nil
}

// A single replacement configuration.
type TranscriptNormalization_Entry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// What to replace. Max length is 100 characters.
	Search string `protobuf:"bytes,1,opt,name=search,proto3" json:"search,omitempty"`
	// What to replace with. Max length is 100 characters.
	Replace string `protobuf:"bytes,2,opt,name=replace,proto3" json:"replace,omitempty"`
	// Whether the search is case sensitive.
	CaseSensitive bool `protobuf:"varint,3,opt,name=case_sensitive,json=caseSensitive,proto3" json:"case_sensitive,omitempty"`
}

func (x *TranscriptNormalization_Entry) Reset() {
	*x = TranscriptNormalization_Entry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_cloud_speech_v1p1beta1_resource_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TranscriptNormalization_Entry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TranscriptNormalization_Entry) ProtoMessage() {}

func (x *TranscriptNormalization_Entry) ProtoReflect() protoreflect.Message {
	mi := &file_google_cloud_speech_v1p1beta1_resource_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TranscriptNormalization_Entry.ProtoReflect.Descriptor instead.
func (*TranscriptNormalization_Entry) Descriptor() ([]byte, []int) {
	return file_google_cloud_speech_v1p1beta1_resource_proto_rawDescGZIP(), []int{3, 0}
}

func (x *TranscriptNormalization_Entry) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

func (x *TranscriptNormalization_Entry) GetReplace() string {
	if x != nil {
		return x.Replace
	}
	return ""
}

func (x *TranscriptNormalization_Entry) GetCaseSensitive() bool {
	if x != nil {
		return x.CaseSensitive
	}
	return false
}

var File_google_cloud_speech_v1p1beta1_resource_proto protoreflect.FileDescriptor

var file_google_cloud_speech_v1p1beta1_resource_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73,
	0x70, 0x65, 0x65, 0x63, 0x68, 0x2f, 0x76, 0x31, 0x70, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x70, 0x65,
	0x65, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x70, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa6, 0x02, 0x0a, 0x0b, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0f,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x43, 0x6c, 0x61,
	0x73, 0x73, 0x49, 0x64, 0x12, 0x4a, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x70, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x2e,
	0x43, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x1a, 0x21, 0x0a, 0x09, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x6c, 0xea, 0x41, 0x69, 0x0a, 0x21, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x44, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x7d, 0x2f,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x43, 0x6c, 0x61, 0x73, 0x73,
	0x65, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73,
	0x7d, 0x22, 0x9d, 0x02, 0x0a, 0x09, 0x50, 0x68, 0x72, 0x61, 0x73, 0x65, 0x53, 0x65, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x49, 0x0a, 0x07, 0x70, 0x68, 0x72, 0x61, 0x73, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x70, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x68, 0x72, 0x61, 0x73, 0x65, 0x53, 0x65, 0x74, 0x2e, 0x50,
	0x68, 0x72, 0x61, 0x73, 0x65, 0x52, 0x07, 0x70, 0x68, 0x72, 0x61, 0x73, 0x65, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x62, 0x6f, 0x6f, 0x73, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x62,
	0x6f, 0x6f, 0x73, 0x74, 0x1a, 0x34, 0x0a, 0x06, 0x50, 0x68, 0x72, 0x61, 0x73, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x6f, 0x6f, 0x73, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x05, 0x62, 0x6f, 0x6f, 0x73, 0x74, 0x3a, 0x65, 0xea, 0x41, 0x62, 0x0a,
	0x1f, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70,
	0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x68, 0x72, 0x61, 0x73, 0x65, 0x53, 0x65, 0x74,
	0x12, 0x3f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x7d, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2f, 0x70, 0x68, 0x72, 0x61, 0x73, 0x65,
	0x53, 0x65, 0x74, 0x73, 0x2f, 0x7b, 0x70, 0x68, 0x72, 0x61, 0x73, 0x65, 0x5f, 0x73, 0x65, 0x74,
	0x7d, 0x22, 0x9c, 0x03, 0x0a, 0x10, 0x53, 0x70, 0x65, 0x65, 0x63, 0x68, 0x41, 0x64, 0x61, 0x70,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x49, 0x0a, 0x0b, 0x70, 0x68, 0x72, 0x61, 0x73, 0x65,
	0x5f, 0x73, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x70, 0x65, 0x65, 0x63,
	0x68, 0x2e, 0x76, 0x31, 0x70, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x68, 0x72, 0x61,
	0x73, 0x65, 0x53, 0x65, 0x74, 0x52, 0x0a, 0x70, 0x68, 0x72, 0x61, 0x73, 0x65, 0x53, 0x65, 0x74,
	0x73, 0x12, 0x58, 0x0a, 0x15, 0x70, 0x68, 0x72, 0x61, 0x73, 0x65, 0x5f, 0x73, 0x65, 0x74, 0x5f,
	0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09,
	0x42, 0x24, 0xfa, 0x41, 0x21, 0x0a, 0x1f, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x68, 0x72,
	0x61, 0x73, 0x65, 0x53, 0x65, 0x74, 0x52, 0x13, 0x70, 0x68, 0x72, 0x61, 0x73, 0x65, 0x53, 0x65,
	0x74, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x51, 0x0a, 0x0e, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x70, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52,
	0x0d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x73, 0x12, 0x5e,
	0x0a, 0x0c, 0x61, 0x62, 0x6e, 0x66, 0x5f, 0x67, 0x72, 0x61, 0x6d, 0x6d, 0x61, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x70, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x53, 0x70, 0x65, 0x65, 0x63, 0x68, 0x41, 0x64, 0x61, 0x70, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x42, 0x4e, 0x46, 0x47, 0x72, 0x61, 0x6d, 0x6d, 0x61,
	0x72, 0x52, 0x0b, 0x61, 0x62, 0x6e, 0x66, 0x47, 0x72, 0x61, 0x6d, 0x6d, 0x61, 0x72, 0x1a, 0x30,
	0x0a, 0x0b, 0x41, 0x42, 0x4e, 0x46, 0x47, 0x72, 0x61, 0x6d, 0x6d, 0x61, 0x72, 0x12, 0x21, 0x0a,
	0x0c, 0x61, 0x62, 0x6e, 0x66, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x62, 0x6e, 0x66, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x73,
	0x22, 0xd3, 0x01, 0x0a, 0x17, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x4e,
	0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x56, 0x0a, 0x07,
	0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x70, 0x65,
	0x65, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x70, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x65, 0x6e, 0x74,
	0x72, 0x69, 0x65, 0x73, 0x1a, 0x60, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x12,
	0x25, 0x0a, 0x0e, 0x63, 0x61, 0x73, 0x65, 0x5f, 0x73, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x63, 0x61, 0x73, 0x65, 0x53, 0x65, 0x6e,
	0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x42, 0x7e, 0x0a, 0x21, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x70, 0x65, 0x65, 0x63,
	0x68, 0x2e, 0x76, 0x31, 0x70, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x13, 0x53, 0x70, 0x65,
	0x65, 0x63, 0x68, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x39, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x2f, 0x61,
	0x70, 0x69, 0x76, 0x31, 0x70, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x73, 0x70, 0x65, 0x65,
	0x63, 0x68, 0x70, 0x62, 0x3b, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x70, 0x62, 0xf8, 0x01, 0x01,
	0xa2, 0x02, 0x03, 0x47, 0x43, 0x53, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_speech_v1p1beta1_resource_proto_rawDescOnce sync.Once
	file_google_cloud_speech_v1p1beta1_resource_proto_rawDescData = file_google_cloud_speech_v1p1beta1_resource_proto_rawDesc
)

func file_google_cloud_speech_v1p1beta1_resource_proto_rawDescGZIP() []byte {
	file_google_cloud_speech_v1p1beta1_resource_proto_rawDescOnce.Do(func() {
		file_google_cloud_speech_v1p1beta1_resource_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_speech_v1p1beta1_resource_proto_rawDescData)
	})
	return file_google_cloud_speech_v1p1beta1_resource_proto_rawDescData
}

var file_google_cloud_speech_v1p1beta1_resource_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_google_cloud_speech_v1p1beta1_resource_proto_goTypes = []any{
	(*CustomClass)(nil),                   // 0: google.cloud.speech.v1p1beta1.CustomClass
	(*PhraseSet)(nil),                     // 1: google.cloud.speech.v1p1beta1.PhraseSet
	(*SpeechAdaptation)(nil),              // 2: google.cloud.speech.v1p1beta1.SpeechAdaptation
	(*TranscriptNormalization)(nil),       // 3: google.cloud.speech.v1p1beta1.TranscriptNormalization
	(*CustomClass_ClassItem)(nil),         // 4: google.cloud.speech.v1p1beta1.CustomClass.ClassItem
	(*PhraseSet_Phrase)(nil),              // 5: google.cloud.speech.v1p1beta1.PhraseSet.Phrase
	(*SpeechAdaptation_ABNFGrammar)(nil),  // 6: google.cloud.speech.v1p1beta1.SpeechAdaptation.ABNFGrammar
	(*TranscriptNormalization_Entry)(nil), // 7: google.cloud.speech.v1p1beta1.TranscriptNormalization.Entry
}
var file_google_cloud_speech_v1p1beta1_resource_proto_depIdxs = []int32{
	4, // 0: google.cloud.speech.v1p1beta1.CustomClass.items:type_name -> google.cloud.speech.v1p1beta1.CustomClass.ClassItem
	5, // 1: google.cloud.speech.v1p1beta1.PhraseSet.phrases:type_name -> google.cloud.speech.v1p1beta1.PhraseSet.Phrase
	1, // 2: google.cloud.speech.v1p1beta1.SpeechAdaptation.phrase_sets:type_name -> google.cloud.speech.v1p1beta1.PhraseSet
	0, // 3: google.cloud.speech.v1p1beta1.SpeechAdaptation.custom_classes:type_name -> google.cloud.speech.v1p1beta1.CustomClass
	6, // 4: google.cloud.speech.v1p1beta1.SpeechAdaptation.abnf_grammar:type_name -> google.cloud.speech.v1p1beta1.SpeechAdaptation.ABNFGrammar
	7, // 5: google.cloud.speech.v1p1beta1.TranscriptNormalization.entries:type_name -> google.cloud.speech.v1p1beta1.TranscriptNormalization.Entry
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_google_cloud_speech_v1p1beta1_resource_proto_init() }
func file_google_cloud_speech_v1p1beta1_resource_proto_init() {
	if File_google_cloud_speech_v1p1beta1_resource_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_cloud_speech_v1p1beta1_resource_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CustomClass); i {
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
		file_google_cloud_speech_v1p1beta1_resource_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*PhraseSet); i {
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
		file_google_cloud_speech_v1p1beta1_resource_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*SpeechAdaptation); i {
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
		file_google_cloud_speech_v1p1beta1_resource_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*TranscriptNormalization); i {
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
		file_google_cloud_speech_v1p1beta1_resource_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*CustomClass_ClassItem); i {
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
		file_google_cloud_speech_v1p1beta1_resource_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*PhraseSet_Phrase); i {
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
		file_google_cloud_speech_v1p1beta1_resource_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*SpeechAdaptation_ABNFGrammar); i {
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
		file_google_cloud_speech_v1p1beta1_resource_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*TranscriptNormalization_Entry); i {
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
			RawDescriptor: file_google_cloud_speech_v1p1beta1_resource_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_speech_v1p1beta1_resource_proto_goTypes,
		DependencyIndexes: file_google_cloud_speech_v1p1beta1_resource_proto_depIdxs,
		MessageInfos:      file_google_cloud_speech_v1p1beta1_resource_proto_msgTypes,
	}.Build()
	File_google_cloud_speech_v1p1beta1_resource_proto = out.File
	file_google_cloud_speech_v1p1beta1_resource_proto_rawDesc = nil
	file_google_cloud_speech_v1p1beta1_resource_proto_goTypes = nil
	file_google_cloud_speech_v1p1beta1_resource_proto_depIdxs = nil
}
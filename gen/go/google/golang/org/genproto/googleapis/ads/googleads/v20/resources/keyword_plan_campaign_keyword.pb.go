// Copyright 2025 Google LLC
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
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: google/ads/googleads/v20/resources/keyword_plan_campaign_keyword.proto

package resources

import (
	enums "google.golang.org/genproto/googleapis/ads/googleads/v20/enums"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A Keyword Plan Campaign keyword.
// Only negative keywords are supported for Campaign Keyword.
type KeywordPlanCampaignKeyword struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Immutable. The resource name of the Keyword Plan Campaign keyword.
	// KeywordPlanCampaignKeyword resource names have the form:
	//
	// `customers/{customer_id}/keywordPlanCampaignKeywords/{kp_campaign_keyword_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The Keyword Plan campaign to which this negative keyword belongs.
	KeywordPlanCampaign *string `protobuf:"bytes,8,opt,name=keyword_plan_campaign,json=keywordPlanCampaign,proto3,oneof" json:"keyword_plan_campaign,omitempty"`
	// Output only. The ID of the Keyword Plan negative keyword.
	Id *int64 `protobuf:"varint,9,opt,name=id,proto3,oneof" json:"id,omitempty"`
	// The keyword text.
	Text *string `protobuf:"bytes,10,opt,name=text,proto3,oneof" json:"text,omitempty"`
	// The keyword match type.
	MatchType enums.KeywordMatchTypeEnum_KeywordMatchType `protobuf:"varint,5,opt,name=match_type,json=matchType,proto3,enum=google.ads.googleads.v20.enums.KeywordMatchTypeEnum_KeywordMatchType" json:"match_type,omitempty"`
	// Immutable. If true, the keyword is negative.
	// Must be set to true. Only negative campaign keywords are supported.
	Negative      *bool `protobuf:"varint,11,opt,name=negative,proto3,oneof" json:"negative,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *KeywordPlanCampaignKeyword) Reset() {
	*x = KeywordPlanCampaignKeyword{}
	mi := &file_google_ads_googleads_v20_resources_keyword_plan_campaign_keyword_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *KeywordPlanCampaignKeyword) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeywordPlanCampaignKeyword) ProtoMessage() {}

func (x *KeywordPlanCampaignKeyword) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_resources_keyword_plan_campaign_keyword_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeywordPlanCampaignKeyword.ProtoReflect.Descriptor instead.
func (*KeywordPlanCampaignKeyword) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_resources_keyword_plan_campaign_keyword_proto_rawDescGZIP(), []int{0}
}

func (x *KeywordPlanCampaignKeyword) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *KeywordPlanCampaignKeyword) GetKeywordPlanCampaign() string {
	if x != nil && x.KeywordPlanCampaign != nil {
		return *x.KeywordPlanCampaign
	}
	return ""
}

func (x *KeywordPlanCampaignKeyword) GetId() int64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *KeywordPlanCampaignKeyword) GetText() string {
	if x != nil && x.Text != nil {
		return *x.Text
	}
	return ""
}

func (x *KeywordPlanCampaignKeyword) GetMatchType() enums.KeywordMatchTypeEnum_KeywordMatchType {
	if x != nil {
		return x.MatchType
	}
	return enums.KeywordMatchTypeEnum_KeywordMatchType(0)
}

func (x *KeywordPlanCampaignKeyword) GetNegative() bool {
	if x != nil && x.Negative != nil {
		return *x.Negative
	}
	return false
}

var File_google_ads_googleads_v20_resources_keyword_plan_campaign_keyword_proto protoreflect.FileDescriptor

const file_google_ads_googleads_v20_resources_keyword_plan_campaign_keyword_proto_rawDesc = "" +
	"\n" +
	"Fgoogle/ads/googleads/v20/resources/keyword_plan_campaign_keyword.proto\x12\"google.ads.googleads.v20.resources\x1a7google/ads/googleads/v20/enums/keyword_match_type.proto\x1a\x1fgoogle/api/field_behavior.proto\x1a\x19google/api/resource.proto\"\xf4\x04\n" +
	"\x1aKeywordPlanCampaignKeyword\x12`\n" +
	"\rresource_name\x18\x01 \x01(\tB;\xe0A\x05\xfaA5\n" +
	"3googleads.googleapis.com/KeywordPlanCampaignKeywordR\fresourceName\x12j\n" +
	"\x15keyword_plan_campaign\x18\b \x01(\tB1\xfaA.\n" +
	",googleads.googleapis.com/KeywordPlanCampaignH\x00R\x13keywordPlanCampaign\x88\x01\x01\x12\x18\n" +
	"\x02id\x18\t \x01(\x03B\x03\xe0A\x03H\x01R\x02id\x88\x01\x01\x12\x17\n" +
	"\x04text\x18\n" +
	" \x01(\tH\x02R\x04text\x88\x01\x01\x12d\n" +
	"\n" +
	"match_type\x18\x05 \x01(\x0e2E.google.ads.googleads.v20.enums.KeywordMatchTypeEnum.KeywordMatchTypeR\tmatchType\x12$\n" +
	"\bnegative\x18\v \x01(\bB\x03\xe0A\x05H\x03R\bnegative\x88\x01\x01:\x91\x01\xeaA\x8d\x01\n" +
	"3googleads.googleapis.com/KeywordPlanCampaignKeyword\x12Vcustomers/{customer_id}/keywordPlanCampaignKeywords/{keyword_plan_campaign_keyword_id}B\x18\n" +
	"\x16_keyword_plan_campaignB\x05\n" +
	"\x03_idB\a\n" +
	"\x05_textB\v\n" +
	"\t_negativeB\x91\x02\n" +
	"&com.google.ads.googleads.v20.resourcesB\x1fKeywordPlanCampaignKeywordProtoP\x01ZKgoogle.golang.org/genproto/googleapis/ads/googleads/v20/resources;resources\xa2\x02\x03GAA\xaa\x02\"Google.Ads.GoogleAds.V20.Resources\xca\x02\"Google\\Ads\\GoogleAds\\V20\\Resources\xea\x02&Google::Ads::GoogleAds::V20::Resourcesb\x06proto3"

var (
	file_google_ads_googleads_v20_resources_keyword_plan_campaign_keyword_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v20_resources_keyword_plan_campaign_keyword_proto_rawDescData []byte
)

func file_google_ads_googleads_v20_resources_keyword_plan_campaign_keyword_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v20_resources_keyword_plan_campaign_keyword_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v20_resources_keyword_plan_campaign_keyword_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_resources_keyword_plan_campaign_keyword_proto_rawDesc), len(file_google_ads_googleads_v20_resources_keyword_plan_campaign_keyword_proto_rawDesc)))
	})
	return file_google_ads_googleads_v20_resources_keyword_plan_campaign_keyword_proto_rawDescData
}

var file_google_ads_googleads_v20_resources_keyword_plan_campaign_keyword_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v20_resources_keyword_plan_campaign_keyword_proto_goTypes = []any{
	(*KeywordPlanCampaignKeyword)(nil),               // 0: google.ads.googleads.v20.resources.KeywordPlanCampaignKeyword
	(enums.KeywordMatchTypeEnum_KeywordMatchType)(0), // 1: google.ads.googleads.v20.enums.KeywordMatchTypeEnum.KeywordMatchType
}
var file_google_ads_googleads_v20_resources_keyword_plan_campaign_keyword_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v20.resources.KeywordPlanCampaignKeyword.match_type:type_name -> google.ads.googleads.v20.enums.KeywordMatchTypeEnum.KeywordMatchType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v20_resources_keyword_plan_campaign_keyword_proto_init() }
func file_google_ads_googleads_v20_resources_keyword_plan_campaign_keyword_proto_init() {
	if File_google_ads_googleads_v20_resources_keyword_plan_campaign_keyword_proto != nil {
		return
	}
	file_google_ads_googleads_v20_resources_keyword_plan_campaign_keyword_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_resources_keyword_plan_campaign_keyword_proto_rawDesc), len(file_google_ads_googleads_v20_resources_keyword_plan_campaign_keyword_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v20_resources_keyword_plan_campaign_keyword_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v20_resources_keyword_plan_campaign_keyword_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v20_resources_keyword_plan_campaign_keyword_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v20_resources_keyword_plan_campaign_keyword_proto = out.File
	file_google_ads_googleads_v20_resources_keyword_plan_campaign_keyword_proto_goTypes = nil
	file_google_ads_googleads_v20_resources_keyword_plan_campaign_keyword_proto_depIdxs = nil
}

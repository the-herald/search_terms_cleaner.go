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
// source: google/ads/googleads/v20/enums/shopping_add_products_to_campaign_recommendation_enum.proto

package enums

import (
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

// Issues that results in a shopping campaign targeting zero products.
type ShoppingAddProductsToCampaignRecommendationEnum_Reason int32

const (
	// Not specified.
	ShoppingAddProductsToCampaignRecommendationEnum_UNSPECIFIED ShoppingAddProductsToCampaignRecommendationEnum_Reason = 0
	// Used for return value only. Represents value unknown in this version.
	ShoppingAddProductsToCampaignRecommendationEnum_UNKNOWN ShoppingAddProductsToCampaignRecommendationEnum_Reason = 1
	// The Merchant Center account does not have any submitted products.
	ShoppingAddProductsToCampaignRecommendationEnum_MERCHANT_CENTER_ACCOUNT_HAS_NO_SUBMITTED_PRODUCTS ShoppingAddProductsToCampaignRecommendationEnum_Reason = 2
	// The Merchant Center account does not have any submitted products in the
	// feed.
	ShoppingAddProductsToCampaignRecommendationEnum_MERCHANT_CENTER_ACCOUNT_HAS_NO_SUBMITTED_PRODUCTS_IN_FEED ShoppingAddProductsToCampaignRecommendationEnum_Reason = 3
	// The Google Ads account has active campaign filters that prevents
	// inclusion of offers in the campaign.
	ShoppingAddProductsToCampaignRecommendationEnum_ADS_ACCOUNT_EXCLUDES_OFFERS_FROM_CAMPAIGN ShoppingAddProductsToCampaignRecommendationEnum_Reason = 4
	// All products available have been explicitly excluded from
	// being targeted by the campaign.
	ShoppingAddProductsToCampaignRecommendationEnum_ALL_PRODUCTS_ARE_EXCLUDED_FROM_CAMPAIGN ShoppingAddProductsToCampaignRecommendationEnum_Reason = 5
)

// Enum value maps for ShoppingAddProductsToCampaignRecommendationEnum_Reason.
var (
	ShoppingAddProductsToCampaignRecommendationEnum_Reason_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "MERCHANT_CENTER_ACCOUNT_HAS_NO_SUBMITTED_PRODUCTS",
		3: "MERCHANT_CENTER_ACCOUNT_HAS_NO_SUBMITTED_PRODUCTS_IN_FEED",
		4: "ADS_ACCOUNT_EXCLUDES_OFFERS_FROM_CAMPAIGN",
		5: "ALL_PRODUCTS_ARE_EXCLUDED_FROM_CAMPAIGN",
	}
	ShoppingAddProductsToCampaignRecommendationEnum_Reason_value = map[string]int32{
		"UNSPECIFIED": 0,
		"UNKNOWN":     1,
		"MERCHANT_CENTER_ACCOUNT_HAS_NO_SUBMITTED_PRODUCTS":         2,
		"MERCHANT_CENTER_ACCOUNT_HAS_NO_SUBMITTED_PRODUCTS_IN_FEED": 3,
		"ADS_ACCOUNT_EXCLUDES_OFFERS_FROM_CAMPAIGN":                 4,
		"ALL_PRODUCTS_ARE_EXCLUDED_FROM_CAMPAIGN":                   5,
	}
)

func (x ShoppingAddProductsToCampaignRecommendationEnum_Reason) Enum() *ShoppingAddProductsToCampaignRecommendationEnum_Reason {
	p := new(ShoppingAddProductsToCampaignRecommendationEnum_Reason)
	*p = x
	return p
}

func (x ShoppingAddProductsToCampaignRecommendationEnum_Reason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ShoppingAddProductsToCampaignRecommendationEnum_Reason) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v20_enums_shopping_add_products_to_campaign_recommendation_enum_proto_enumTypes[0].Descriptor()
}

func (ShoppingAddProductsToCampaignRecommendationEnum_Reason) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v20_enums_shopping_add_products_to_campaign_recommendation_enum_proto_enumTypes[0]
}

func (x ShoppingAddProductsToCampaignRecommendationEnum_Reason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ShoppingAddProductsToCampaignRecommendationEnum_Reason.Descriptor instead.
func (ShoppingAddProductsToCampaignRecommendationEnum_Reason) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_enums_shopping_add_products_to_campaign_recommendation_enum_proto_rawDescGZIP(), []int{0, 0}
}

// Indicates the key issue that results in a shopping campaign targeting zero
// products.
type ShoppingAddProductsToCampaignRecommendationEnum struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ShoppingAddProductsToCampaignRecommendationEnum) Reset() {
	*x = ShoppingAddProductsToCampaignRecommendationEnum{}
	mi := &file_google_ads_googleads_v20_enums_shopping_add_products_to_campaign_recommendation_enum_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ShoppingAddProductsToCampaignRecommendationEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShoppingAddProductsToCampaignRecommendationEnum) ProtoMessage() {}

func (x *ShoppingAddProductsToCampaignRecommendationEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_enums_shopping_add_products_to_campaign_recommendation_enum_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShoppingAddProductsToCampaignRecommendationEnum.ProtoReflect.Descriptor instead.
func (*ShoppingAddProductsToCampaignRecommendationEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_enums_shopping_add_products_to_campaign_recommendation_enum_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v20_enums_shopping_add_products_to_campaign_recommendation_enum_proto protoreflect.FileDescriptor

const file_google_ads_googleads_v20_enums_shopping_add_products_to_campaign_recommendation_enum_proto_rawDesc = "" +
	"\n" +
	"Zgoogle/ads/googleads/v20/enums/shopping_add_products_to_campaign_recommendation_enum.proto\x12\x1egoogle.ads.googleads.v20.enums\"\xac\x02\n" +
	"/ShoppingAddProductsToCampaignRecommendationEnum\"\xf8\x01\n" +
	"\x06Reason\x12\x0f\n" +
	"\vUNSPECIFIED\x10\x00\x12\v\n" +
	"\aUNKNOWN\x10\x01\x125\n" +
	"1MERCHANT_CENTER_ACCOUNT_HAS_NO_SUBMITTED_PRODUCTS\x10\x02\x12=\n" +
	"9MERCHANT_CENTER_ACCOUNT_HAS_NO_SUBMITTED_PRODUCTS_IN_FEED\x10\x03\x12-\n" +
	")ADS_ACCOUNT_EXCLUDES_OFFERS_FROM_CAMPAIGN\x10\x04\x12+\n" +
	"'ALL_PRODUCTS_ARE_EXCLUDED_FROM_CAMPAIGN\x10\x05B\x8e\x02\n" +
	"\"com.google.ads.googleads.v20.enumsB4ShoppingAddProductsToCampaignRecommendationEnumProtoP\x01ZCgoogle.golang.org/genproto/googleapis/ads/googleads/v20/enums;enums\xa2\x02\x03GAA\xaa\x02\x1eGoogle.Ads.GoogleAds.V20.Enums\xca\x02\x1eGoogle\\Ads\\GoogleAds\\V20\\Enums\xea\x02\"Google::Ads::GoogleAds::V20::Enumsb\x06proto3"

var (
	file_google_ads_googleads_v20_enums_shopping_add_products_to_campaign_recommendation_enum_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v20_enums_shopping_add_products_to_campaign_recommendation_enum_proto_rawDescData []byte
)

func file_google_ads_googleads_v20_enums_shopping_add_products_to_campaign_recommendation_enum_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v20_enums_shopping_add_products_to_campaign_recommendation_enum_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v20_enums_shopping_add_products_to_campaign_recommendation_enum_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_enums_shopping_add_products_to_campaign_recommendation_enum_proto_rawDesc), len(file_google_ads_googleads_v20_enums_shopping_add_products_to_campaign_recommendation_enum_proto_rawDesc)))
	})
	return file_google_ads_googleads_v20_enums_shopping_add_products_to_campaign_recommendation_enum_proto_rawDescData
}

var file_google_ads_googleads_v20_enums_shopping_add_products_to_campaign_recommendation_enum_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v20_enums_shopping_add_products_to_campaign_recommendation_enum_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v20_enums_shopping_add_products_to_campaign_recommendation_enum_proto_goTypes = []any{
	(ShoppingAddProductsToCampaignRecommendationEnum_Reason)(0), // 0: google.ads.googleads.v20.enums.ShoppingAddProductsToCampaignRecommendationEnum.Reason
	(*ShoppingAddProductsToCampaignRecommendationEnum)(nil),     // 1: google.ads.googleads.v20.enums.ShoppingAddProductsToCampaignRecommendationEnum
}
var file_google_ads_googleads_v20_enums_shopping_add_products_to_campaign_recommendation_enum_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_google_ads_googleads_v20_enums_shopping_add_products_to_campaign_recommendation_enum_proto_init()
}
func file_google_ads_googleads_v20_enums_shopping_add_products_to_campaign_recommendation_enum_proto_init() {
	if File_google_ads_googleads_v20_enums_shopping_add_products_to_campaign_recommendation_enum_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_enums_shopping_add_products_to_campaign_recommendation_enum_proto_rawDesc), len(file_google_ads_googleads_v20_enums_shopping_add_products_to_campaign_recommendation_enum_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v20_enums_shopping_add_products_to_campaign_recommendation_enum_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v20_enums_shopping_add_products_to_campaign_recommendation_enum_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v20_enums_shopping_add_products_to_campaign_recommendation_enum_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v20_enums_shopping_add_products_to_campaign_recommendation_enum_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v20_enums_shopping_add_products_to_campaign_recommendation_enum_proto = out.File
	file_google_ads_googleads_v20_enums_shopping_add_products_to_campaign_recommendation_enum_proto_goTypes = nil
	file_google_ads_googleads_v20_enums_shopping_add_products_to_campaign_recommendation_enum_proto_depIdxs = nil
}

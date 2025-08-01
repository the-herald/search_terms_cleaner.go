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
// source: google/ads/googleads/v20/enums/asset_field_type.proto

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

// Enum describing the possible placements of an asset.
type AssetFieldTypeEnum_AssetFieldType int32

const (
	// Not specified.
	AssetFieldTypeEnum_UNSPECIFIED AssetFieldTypeEnum_AssetFieldType = 0
	// Used for return value only. Represents value unknown in this version.
	AssetFieldTypeEnum_UNKNOWN AssetFieldTypeEnum_AssetFieldType = 1
	// The asset is linked for use as a headline.
	AssetFieldTypeEnum_HEADLINE AssetFieldTypeEnum_AssetFieldType = 2
	// The asset is linked for use as a description.
	AssetFieldTypeEnum_DESCRIPTION AssetFieldTypeEnum_AssetFieldType = 3
	// The asset is linked for use as mandatory ad text.
	AssetFieldTypeEnum_MANDATORY_AD_TEXT AssetFieldTypeEnum_AssetFieldType = 4
	// The asset is linked for use as a marketing image.
	AssetFieldTypeEnum_MARKETING_IMAGE AssetFieldTypeEnum_AssetFieldType = 5
	// The asset is linked for use as a media bundle.
	AssetFieldTypeEnum_MEDIA_BUNDLE AssetFieldTypeEnum_AssetFieldType = 6
	// The asset is linked for use as a YouTube video.
	AssetFieldTypeEnum_YOUTUBE_VIDEO AssetFieldTypeEnum_AssetFieldType = 7
	// The asset is linked to indicate that a hotels campaign is "Book on
	// Google" enabled.
	AssetFieldTypeEnum_BOOK_ON_GOOGLE AssetFieldTypeEnum_AssetFieldType = 8
	// The asset is linked for use as a Lead Form extension.
	AssetFieldTypeEnum_LEAD_FORM AssetFieldTypeEnum_AssetFieldType = 9
	// The asset is linked for use as a Promotion extension.
	AssetFieldTypeEnum_PROMOTION AssetFieldTypeEnum_AssetFieldType = 10
	// The asset is linked for use as a Callout extension.
	AssetFieldTypeEnum_CALLOUT AssetFieldTypeEnum_AssetFieldType = 11
	// The asset is linked for use as a Structured Snippet extension.
	AssetFieldTypeEnum_STRUCTURED_SNIPPET AssetFieldTypeEnum_AssetFieldType = 12
	// The asset is linked for use as a Sitelink.
	AssetFieldTypeEnum_SITELINK AssetFieldTypeEnum_AssetFieldType = 13
	// The asset is linked for use as a Mobile App extension.
	AssetFieldTypeEnum_MOBILE_APP AssetFieldTypeEnum_AssetFieldType = 14
	// The asset is linked for use as a Hotel Callout extension.
	AssetFieldTypeEnum_HOTEL_CALLOUT AssetFieldTypeEnum_AssetFieldType = 15
	// The asset is linked for use as a Call extension.
	AssetFieldTypeEnum_CALL AssetFieldTypeEnum_AssetFieldType = 16
	// The asset is linked for use as a Price extension.
	AssetFieldTypeEnum_PRICE AssetFieldTypeEnum_AssetFieldType = 24
	// The asset is linked for use as a long headline.
	AssetFieldTypeEnum_LONG_HEADLINE AssetFieldTypeEnum_AssetFieldType = 17
	// The asset is linked for use as a business name.
	AssetFieldTypeEnum_BUSINESS_NAME AssetFieldTypeEnum_AssetFieldType = 18
	// The asset is linked for use as a square marketing image.
	AssetFieldTypeEnum_SQUARE_MARKETING_IMAGE AssetFieldTypeEnum_AssetFieldType = 19
	// The asset is linked for use as a portrait marketing image.
	AssetFieldTypeEnum_PORTRAIT_MARKETING_IMAGE AssetFieldTypeEnum_AssetFieldType = 20
	// The asset is linked for use as a logo.
	AssetFieldTypeEnum_LOGO AssetFieldTypeEnum_AssetFieldType = 21
	// The asset is linked for use as a landscape logo.
	AssetFieldTypeEnum_LANDSCAPE_LOGO AssetFieldTypeEnum_AssetFieldType = 22
	// The asset is linked for use as a non YouTube logo.
	AssetFieldTypeEnum_VIDEO AssetFieldTypeEnum_AssetFieldType = 23
	// The asset is linked for use to select a call-to-action.
	AssetFieldTypeEnum_CALL_TO_ACTION_SELECTION AssetFieldTypeEnum_AssetFieldType = 25
	// The asset is linked for use to select an ad image.
	AssetFieldTypeEnum_AD_IMAGE AssetFieldTypeEnum_AssetFieldType = 26
	// The asset is linked for use as a business logo.
	AssetFieldTypeEnum_BUSINESS_LOGO AssetFieldTypeEnum_AssetFieldType = 27
	// The asset is linked for use as a hotel property in a Performance Max for
	// travel goals campaign.
	AssetFieldTypeEnum_HOTEL_PROPERTY AssetFieldTypeEnum_AssetFieldType = 28
	// The asset is linked for use as a Demand Gen carousel card.
	AssetFieldTypeEnum_DEMAND_GEN_CAROUSEL_CARD AssetFieldTypeEnum_AssetFieldType = 30
	// The asset is linked for use as a Business Message.
	AssetFieldTypeEnum_BUSINESS_MESSAGE AssetFieldTypeEnum_AssetFieldType = 31
	// The asset is linked for use as a tall portrait marketing image.
	AssetFieldTypeEnum_TALL_PORTRAIT_MARKETING_IMAGE AssetFieldTypeEnum_AssetFieldType = 32
)

// Enum value maps for AssetFieldTypeEnum_AssetFieldType.
var (
	AssetFieldTypeEnum_AssetFieldType_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "HEADLINE",
		3:  "DESCRIPTION",
		4:  "MANDATORY_AD_TEXT",
		5:  "MARKETING_IMAGE",
		6:  "MEDIA_BUNDLE",
		7:  "YOUTUBE_VIDEO",
		8:  "BOOK_ON_GOOGLE",
		9:  "LEAD_FORM",
		10: "PROMOTION",
		11: "CALLOUT",
		12: "STRUCTURED_SNIPPET",
		13: "SITELINK",
		14: "MOBILE_APP",
		15: "HOTEL_CALLOUT",
		16: "CALL",
		24: "PRICE",
		17: "LONG_HEADLINE",
		18: "BUSINESS_NAME",
		19: "SQUARE_MARKETING_IMAGE",
		20: "PORTRAIT_MARKETING_IMAGE",
		21: "LOGO",
		22: "LANDSCAPE_LOGO",
		23: "VIDEO",
		25: "CALL_TO_ACTION_SELECTION",
		26: "AD_IMAGE",
		27: "BUSINESS_LOGO",
		28: "HOTEL_PROPERTY",
		30: "DEMAND_GEN_CAROUSEL_CARD",
		31: "BUSINESS_MESSAGE",
		32: "TALL_PORTRAIT_MARKETING_IMAGE",
	}
	AssetFieldTypeEnum_AssetFieldType_value = map[string]int32{
		"UNSPECIFIED":                   0,
		"UNKNOWN":                       1,
		"HEADLINE":                      2,
		"DESCRIPTION":                   3,
		"MANDATORY_AD_TEXT":             4,
		"MARKETING_IMAGE":               5,
		"MEDIA_BUNDLE":                  6,
		"YOUTUBE_VIDEO":                 7,
		"BOOK_ON_GOOGLE":                8,
		"LEAD_FORM":                     9,
		"PROMOTION":                     10,
		"CALLOUT":                       11,
		"STRUCTURED_SNIPPET":            12,
		"SITELINK":                      13,
		"MOBILE_APP":                    14,
		"HOTEL_CALLOUT":                 15,
		"CALL":                          16,
		"PRICE":                         24,
		"LONG_HEADLINE":                 17,
		"BUSINESS_NAME":                 18,
		"SQUARE_MARKETING_IMAGE":        19,
		"PORTRAIT_MARKETING_IMAGE":      20,
		"LOGO":                          21,
		"LANDSCAPE_LOGO":                22,
		"VIDEO":                         23,
		"CALL_TO_ACTION_SELECTION":      25,
		"AD_IMAGE":                      26,
		"BUSINESS_LOGO":                 27,
		"HOTEL_PROPERTY":                28,
		"DEMAND_GEN_CAROUSEL_CARD":      30,
		"BUSINESS_MESSAGE":              31,
		"TALL_PORTRAIT_MARKETING_IMAGE": 32,
	}
)

func (x AssetFieldTypeEnum_AssetFieldType) Enum() *AssetFieldTypeEnum_AssetFieldType {
	p := new(AssetFieldTypeEnum_AssetFieldType)
	*p = x
	return p
}

func (x AssetFieldTypeEnum_AssetFieldType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AssetFieldTypeEnum_AssetFieldType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v20_enums_asset_field_type_proto_enumTypes[0].Descriptor()
}

func (AssetFieldTypeEnum_AssetFieldType) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v20_enums_asset_field_type_proto_enumTypes[0]
}

func (x AssetFieldTypeEnum_AssetFieldType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AssetFieldTypeEnum_AssetFieldType.Descriptor instead.
func (AssetFieldTypeEnum_AssetFieldType) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_enums_asset_field_type_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing the possible placements of an asset.
type AssetFieldTypeEnum struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AssetFieldTypeEnum) Reset() {
	*x = AssetFieldTypeEnum{}
	mi := &file_google_ads_googleads_v20_enums_asset_field_type_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AssetFieldTypeEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssetFieldTypeEnum) ProtoMessage() {}

func (x *AssetFieldTypeEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_enums_asset_field_type_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssetFieldTypeEnum.ProtoReflect.Descriptor instead.
func (*AssetFieldTypeEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_enums_asset_field_type_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v20_enums_asset_field_type_proto protoreflect.FileDescriptor

const file_google_ads_googleads_v20_enums_asset_field_type_proto_rawDesc = "" +
	"\n" +
	"5google/ads/googleads/v20/enums/asset_field_type.proto\x12\x1egoogle.ads.googleads.v20.enums\"\x85\x05\n" +
	"\x12AssetFieldTypeEnum\"\xee\x04\n" +
	"\x0eAssetFieldType\x12\x0f\n" +
	"\vUNSPECIFIED\x10\x00\x12\v\n" +
	"\aUNKNOWN\x10\x01\x12\f\n" +
	"\bHEADLINE\x10\x02\x12\x0f\n" +
	"\vDESCRIPTION\x10\x03\x12\x15\n" +
	"\x11MANDATORY_AD_TEXT\x10\x04\x12\x13\n" +
	"\x0fMARKETING_IMAGE\x10\x05\x12\x10\n" +
	"\fMEDIA_BUNDLE\x10\x06\x12\x11\n" +
	"\rYOUTUBE_VIDEO\x10\a\x12\x12\n" +
	"\x0eBOOK_ON_GOOGLE\x10\b\x12\r\n" +
	"\tLEAD_FORM\x10\t\x12\r\n" +
	"\tPROMOTION\x10\n" +
	"\x12\v\n" +
	"\aCALLOUT\x10\v\x12\x16\n" +
	"\x12STRUCTURED_SNIPPET\x10\f\x12\f\n" +
	"\bSITELINK\x10\r\x12\x0e\n" +
	"\n" +
	"MOBILE_APP\x10\x0e\x12\x11\n" +
	"\rHOTEL_CALLOUT\x10\x0f\x12\b\n" +
	"\x04CALL\x10\x10\x12\t\n" +
	"\x05PRICE\x10\x18\x12\x11\n" +
	"\rLONG_HEADLINE\x10\x11\x12\x11\n" +
	"\rBUSINESS_NAME\x10\x12\x12\x1a\n" +
	"\x16SQUARE_MARKETING_IMAGE\x10\x13\x12\x1c\n" +
	"\x18PORTRAIT_MARKETING_IMAGE\x10\x14\x12\b\n" +
	"\x04LOGO\x10\x15\x12\x12\n" +
	"\x0eLANDSCAPE_LOGO\x10\x16\x12\t\n" +
	"\x05VIDEO\x10\x17\x12\x1c\n" +
	"\x18CALL_TO_ACTION_SELECTION\x10\x19\x12\f\n" +
	"\bAD_IMAGE\x10\x1a\x12\x11\n" +
	"\rBUSINESS_LOGO\x10\x1b\x12\x12\n" +
	"\x0eHOTEL_PROPERTY\x10\x1c\x12\x1c\n" +
	"\x18DEMAND_GEN_CAROUSEL_CARD\x10\x1e\x12\x14\n" +
	"\x10BUSINESS_MESSAGE\x10\x1f\x12!\n" +
	"\x1dTALL_PORTRAIT_MARKETING_IMAGE\x10 B\xed\x01\n" +
	"\"com.google.ads.googleads.v20.enumsB\x13AssetFieldTypeProtoP\x01ZCgoogle.golang.org/genproto/googleapis/ads/googleads/v20/enums;enums\xa2\x02\x03GAA\xaa\x02\x1eGoogle.Ads.GoogleAds.V20.Enums\xca\x02\x1eGoogle\\Ads\\GoogleAds\\V20\\Enums\xea\x02\"Google::Ads::GoogleAds::V20::Enumsb\x06proto3"

var (
	file_google_ads_googleads_v20_enums_asset_field_type_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v20_enums_asset_field_type_proto_rawDescData []byte
)

func file_google_ads_googleads_v20_enums_asset_field_type_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v20_enums_asset_field_type_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v20_enums_asset_field_type_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_enums_asset_field_type_proto_rawDesc), len(file_google_ads_googleads_v20_enums_asset_field_type_proto_rawDesc)))
	})
	return file_google_ads_googleads_v20_enums_asset_field_type_proto_rawDescData
}

var file_google_ads_googleads_v20_enums_asset_field_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v20_enums_asset_field_type_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v20_enums_asset_field_type_proto_goTypes = []any{
	(AssetFieldTypeEnum_AssetFieldType)(0), // 0: google.ads.googleads.v20.enums.AssetFieldTypeEnum.AssetFieldType
	(*AssetFieldTypeEnum)(nil),             // 1: google.ads.googleads.v20.enums.AssetFieldTypeEnum
}
var file_google_ads_googleads_v20_enums_asset_field_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v20_enums_asset_field_type_proto_init() }
func file_google_ads_googleads_v20_enums_asset_field_type_proto_init() {
	if File_google_ads_googleads_v20_enums_asset_field_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_enums_asset_field_type_proto_rawDesc), len(file_google_ads_googleads_v20_enums_asset_field_type_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v20_enums_asset_field_type_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v20_enums_asset_field_type_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v20_enums_asset_field_type_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v20_enums_asset_field_type_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v20_enums_asset_field_type_proto = out.File
	file_google_ads_googleads_v20_enums_asset_field_type_proto_goTypes = nil
	file_google_ads_googleads_v20_enums_asset_field_type_proto_depIdxs = nil
}

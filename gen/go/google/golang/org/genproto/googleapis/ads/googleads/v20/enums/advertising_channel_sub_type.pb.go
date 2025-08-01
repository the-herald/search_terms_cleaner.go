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
// source: google/ads/googleads/v20/enums/advertising_channel_sub_type.proto

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

// Enum describing the different channel subtypes.
type AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType int32

const (
	// Not specified.
	AdvertisingChannelSubTypeEnum_UNSPECIFIED AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 0
	// Used as a return value only. Represents value unknown in this version.
	AdvertisingChannelSubTypeEnum_UNKNOWN AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 1
	// Mobile app campaigns for Search.
	AdvertisingChannelSubTypeEnum_SEARCH_MOBILE_APP AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 2
	// Mobile app campaigns for Display.
	AdvertisingChannelSubTypeEnum_DISPLAY_MOBILE_APP AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 3
	// AdWords express campaigns for search.
	AdvertisingChannelSubTypeEnum_SEARCH_EXPRESS AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 4
	// AdWords Express campaigns for display.
	AdvertisingChannelSubTypeEnum_DISPLAY_EXPRESS AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 5
	// Smart Shopping campaigns.
	AdvertisingChannelSubTypeEnum_SHOPPING_SMART_ADS AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 6
	// Gmail Ad campaigns.
	AdvertisingChannelSubTypeEnum_DISPLAY_GMAIL_AD AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 7
	// Smart display campaigns. New campaigns of this sub type cannot be
	// created.
	AdvertisingChannelSubTypeEnum_DISPLAY_SMART_CAMPAIGN AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 8
	// Video TrueView for Action campaigns.
	AdvertisingChannelSubTypeEnum_VIDEO_ACTION AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 10
	// Video campaigns with non-skippable video ads.
	AdvertisingChannelSubTypeEnum_VIDEO_NON_SKIPPABLE AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 11
	// App Campaign that lets you easily promote your Android or iOS app
	// across Google's top properties including Search, Play, YouTube, and the
	// Google Display Network.
	AdvertisingChannelSubTypeEnum_APP_CAMPAIGN AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 12
	// App Campaign for engagement, focused on driving re-engagement with the
	// app across several of Google's top properties including Search, YouTube,
	// and the Google Display Network.
	AdvertisingChannelSubTypeEnum_APP_CAMPAIGN_FOR_ENGAGEMENT AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 13
	// Campaigns specialized for local advertising.
	AdvertisingChannelSubTypeEnum_LOCAL_CAMPAIGN AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 14
	// Shopping Comparison Listing campaigns.
	AdvertisingChannelSubTypeEnum_SHOPPING_COMPARISON_LISTING_ADS AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 15
	// Standard Smart campaigns.
	AdvertisingChannelSubTypeEnum_SMART_CAMPAIGN AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 16
	// Video campaigns with sequence video ads.
	AdvertisingChannelSubTypeEnum_VIDEO_SEQUENCE AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 17
	// App Campaign for pre registration, specialized for advertising mobile
	// app pre-registration, that targets multiple advertising channels across
	// Google Play, YouTube and Display Network. See
	// https://support.google.com/google-ads/answer/9441344 to learn more.
	AdvertisingChannelSubTypeEnum_APP_CAMPAIGN_FOR_PRE_REGISTRATION AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 18
	// Video reach campaign with Target Frequency bidding strategy.
	AdvertisingChannelSubTypeEnum_VIDEO_REACH_TARGET_FREQUENCY AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 19
	// Travel Activities campaigns.
	AdvertisingChannelSubTypeEnum_TRAVEL_ACTIVITIES AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 20
	// YouTube Audio campaigns.
	AdvertisingChannelSubTypeEnum_YOUTUBE_AUDIO AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType = 22
)

// Enum value maps for AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType.
var (
	AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "SEARCH_MOBILE_APP",
		3:  "DISPLAY_MOBILE_APP",
		4:  "SEARCH_EXPRESS",
		5:  "DISPLAY_EXPRESS",
		6:  "SHOPPING_SMART_ADS",
		7:  "DISPLAY_GMAIL_AD",
		8:  "DISPLAY_SMART_CAMPAIGN",
		10: "VIDEO_ACTION",
		11: "VIDEO_NON_SKIPPABLE",
		12: "APP_CAMPAIGN",
		13: "APP_CAMPAIGN_FOR_ENGAGEMENT",
		14: "LOCAL_CAMPAIGN",
		15: "SHOPPING_COMPARISON_LISTING_ADS",
		16: "SMART_CAMPAIGN",
		17: "VIDEO_SEQUENCE",
		18: "APP_CAMPAIGN_FOR_PRE_REGISTRATION",
		19: "VIDEO_REACH_TARGET_FREQUENCY",
		20: "TRAVEL_ACTIVITIES",
		22: "YOUTUBE_AUDIO",
	}
	AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType_value = map[string]int32{
		"UNSPECIFIED":                       0,
		"UNKNOWN":                           1,
		"SEARCH_MOBILE_APP":                 2,
		"DISPLAY_MOBILE_APP":                3,
		"SEARCH_EXPRESS":                    4,
		"DISPLAY_EXPRESS":                   5,
		"SHOPPING_SMART_ADS":                6,
		"DISPLAY_GMAIL_AD":                  7,
		"DISPLAY_SMART_CAMPAIGN":            8,
		"VIDEO_ACTION":                      10,
		"VIDEO_NON_SKIPPABLE":               11,
		"APP_CAMPAIGN":                      12,
		"APP_CAMPAIGN_FOR_ENGAGEMENT":       13,
		"LOCAL_CAMPAIGN":                    14,
		"SHOPPING_COMPARISON_LISTING_ADS":   15,
		"SMART_CAMPAIGN":                    16,
		"VIDEO_SEQUENCE":                    17,
		"APP_CAMPAIGN_FOR_PRE_REGISTRATION": 18,
		"VIDEO_REACH_TARGET_FREQUENCY":      19,
		"TRAVEL_ACTIVITIES":                 20,
		"YOUTUBE_AUDIO":                     22,
	}
)

func (x AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType) Enum() *AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType {
	p := new(AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType)
	*p = x
	return p
}

func (x AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v20_enums_advertising_channel_sub_type_proto_enumTypes[0].Descriptor()
}

func (AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v20_enums_advertising_channel_sub_type_proto_enumTypes[0]
}

func (x AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType.Descriptor instead.
func (AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_enums_advertising_channel_sub_type_proto_rawDescGZIP(), []int{0, 0}
}

// An immutable specialization of an Advertising Channel.
type AdvertisingChannelSubTypeEnum struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AdvertisingChannelSubTypeEnum) Reset() {
	*x = AdvertisingChannelSubTypeEnum{}
	mi := &file_google_ads_googleads_v20_enums_advertising_channel_sub_type_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AdvertisingChannelSubTypeEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdvertisingChannelSubTypeEnum) ProtoMessage() {}

func (x *AdvertisingChannelSubTypeEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_enums_advertising_channel_sub_type_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdvertisingChannelSubTypeEnum.ProtoReflect.Descriptor instead.
func (*AdvertisingChannelSubTypeEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_enums_advertising_channel_sub_type_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v20_enums_advertising_channel_sub_type_proto protoreflect.FileDescriptor

const file_google_ads_googleads_v20_enums_advertising_channel_sub_type_proto_rawDesc = "" +
	"\n" +
	"Agoogle/ads/googleads/v20/enums/advertising_channel_sub_type.proto\x12\x1egoogle.ads.googleads.v20.enums\"\xaf\x04\n" +
	"\x1dAdvertisingChannelSubTypeEnum\"\x8d\x04\n" +
	"\x19AdvertisingChannelSubType\x12\x0f\n" +
	"\vUNSPECIFIED\x10\x00\x12\v\n" +
	"\aUNKNOWN\x10\x01\x12\x15\n" +
	"\x11SEARCH_MOBILE_APP\x10\x02\x12\x16\n" +
	"\x12DISPLAY_MOBILE_APP\x10\x03\x12\x12\n" +
	"\x0eSEARCH_EXPRESS\x10\x04\x12\x13\n" +
	"\x0fDISPLAY_EXPRESS\x10\x05\x12\x16\n" +
	"\x12SHOPPING_SMART_ADS\x10\x06\x12\x14\n" +
	"\x10DISPLAY_GMAIL_AD\x10\a\x12\x1a\n" +
	"\x16DISPLAY_SMART_CAMPAIGN\x10\b\x12\x10\n" +
	"\fVIDEO_ACTION\x10\n" +
	"\x12\x17\n" +
	"\x13VIDEO_NON_SKIPPABLE\x10\v\x12\x10\n" +
	"\fAPP_CAMPAIGN\x10\f\x12\x1f\n" +
	"\x1bAPP_CAMPAIGN_FOR_ENGAGEMENT\x10\r\x12\x12\n" +
	"\x0eLOCAL_CAMPAIGN\x10\x0e\x12#\n" +
	"\x1fSHOPPING_COMPARISON_LISTING_ADS\x10\x0f\x12\x12\n" +
	"\x0eSMART_CAMPAIGN\x10\x10\x12\x12\n" +
	"\x0eVIDEO_SEQUENCE\x10\x11\x12%\n" +
	"!APP_CAMPAIGN_FOR_PRE_REGISTRATION\x10\x12\x12 \n" +
	"\x1cVIDEO_REACH_TARGET_FREQUENCY\x10\x13\x12\x15\n" +
	"\x11TRAVEL_ACTIVITIES\x10\x14\x12\x11\n" +
	"\rYOUTUBE_AUDIO\x10\x16B\xf8\x01\n" +
	"\"com.google.ads.googleads.v20.enumsB\x1eAdvertisingChannelSubTypeProtoP\x01ZCgoogle.golang.org/genproto/googleapis/ads/googleads/v20/enums;enums\xa2\x02\x03GAA\xaa\x02\x1eGoogle.Ads.GoogleAds.V20.Enums\xca\x02\x1eGoogle\\Ads\\GoogleAds\\V20\\Enums\xea\x02\"Google::Ads::GoogleAds::V20::Enumsb\x06proto3"

var (
	file_google_ads_googleads_v20_enums_advertising_channel_sub_type_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v20_enums_advertising_channel_sub_type_proto_rawDescData []byte
)

func file_google_ads_googleads_v20_enums_advertising_channel_sub_type_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v20_enums_advertising_channel_sub_type_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v20_enums_advertising_channel_sub_type_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_enums_advertising_channel_sub_type_proto_rawDesc), len(file_google_ads_googleads_v20_enums_advertising_channel_sub_type_proto_rawDesc)))
	})
	return file_google_ads_googleads_v20_enums_advertising_channel_sub_type_proto_rawDescData
}

var file_google_ads_googleads_v20_enums_advertising_channel_sub_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v20_enums_advertising_channel_sub_type_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v20_enums_advertising_channel_sub_type_proto_goTypes = []any{
	(AdvertisingChannelSubTypeEnum_AdvertisingChannelSubType)(0), // 0: google.ads.googleads.v20.enums.AdvertisingChannelSubTypeEnum.AdvertisingChannelSubType
	(*AdvertisingChannelSubTypeEnum)(nil),                        // 1: google.ads.googleads.v20.enums.AdvertisingChannelSubTypeEnum
}
var file_google_ads_googleads_v20_enums_advertising_channel_sub_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v20_enums_advertising_channel_sub_type_proto_init() }
func file_google_ads_googleads_v20_enums_advertising_channel_sub_type_proto_init() {
	if File_google_ads_googleads_v20_enums_advertising_channel_sub_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_enums_advertising_channel_sub_type_proto_rawDesc), len(file_google_ads_googleads_v20_enums_advertising_channel_sub_type_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v20_enums_advertising_channel_sub_type_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v20_enums_advertising_channel_sub_type_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v20_enums_advertising_channel_sub_type_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v20_enums_advertising_channel_sub_type_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v20_enums_advertising_channel_sub_type_proto = out.File
	file_google_ads_googleads_v20_enums_advertising_channel_sub_type_proto_goTypes = nil
	file_google_ads_googleads_v20_enums_advertising_channel_sub_type_proto_depIdxs = nil
}

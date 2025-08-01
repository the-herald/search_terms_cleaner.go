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
// source: google/ads/googleads/v20/enums/conversion_action_type.proto

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

// Possible types of a conversion action.
type ConversionActionTypeEnum_ConversionActionType int32

const (
	// Not specified.
	ConversionActionTypeEnum_UNSPECIFIED ConversionActionTypeEnum_ConversionActionType = 0
	// Used for return value only. Represents value unknown in this version.
	ConversionActionTypeEnum_UNKNOWN ConversionActionTypeEnum_ConversionActionType = 1
	// Conversions that occur when a user clicks on an ad's call extension.
	ConversionActionTypeEnum_AD_CALL ConversionActionTypeEnum_ConversionActionType = 2
	// Conversions that occur when a user on a mobile device clicks a phone
	// number.
	ConversionActionTypeEnum_CLICK_TO_CALL ConversionActionTypeEnum_ConversionActionType = 3
	// Conversions that occur when a user downloads a mobile app from the Google
	// Play Store.
	ConversionActionTypeEnum_GOOGLE_PLAY_DOWNLOAD ConversionActionTypeEnum_ConversionActionType = 4
	// Conversions that occur when a user makes a purchase in an app through
	// Android billing.
	ConversionActionTypeEnum_GOOGLE_PLAY_IN_APP_PURCHASE ConversionActionTypeEnum_ConversionActionType = 5
	// Call conversions that are tracked by the advertiser and uploaded.
	ConversionActionTypeEnum_UPLOAD_CALLS ConversionActionTypeEnum_ConversionActionType = 6
	// Conversions that are tracked by the advertiser and uploaded with
	// attributed clicks.
	ConversionActionTypeEnum_UPLOAD_CLICKS ConversionActionTypeEnum_ConversionActionType = 7
	// Conversions that occur on a webpage.
	ConversionActionTypeEnum_WEBPAGE ConversionActionTypeEnum_ConversionActionType = 8
	// Conversions that occur when a user calls a dynamically-generated phone
	// number from an advertiser's website.
	ConversionActionTypeEnum_WEBSITE_CALL ConversionActionTypeEnum_ConversionActionType = 9
	// Store Sales conversion based on first-party or third-party merchant
	// data uploads.
	// Only customers on the allowlist can use store sales direct upload types.
	ConversionActionTypeEnum_STORE_SALES_DIRECT_UPLOAD ConversionActionTypeEnum_ConversionActionType = 10
	// Store Sales conversion based on first-party or third-party merchant
	// data uploads and/or from in-store purchases using cards from payment
	// networks.
	// Only customers on the allowlist can use store sales types.
	// Read only.
	ConversionActionTypeEnum_STORE_SALES ConversionActionTypeEnum_ConversionActionType = 11
	// Android app first open conversions tracked through Firebase.
	ConversionActionTypeEnum_FIREBASE_ANDROID_FIRST_OPEN ConversionActionTypeEnum_ConversionActionType = 12
	// Android app in app purchase conversions tracked through Firebase.
	ConversionActionTypeEnum_FIREBASE_ANDROID_IN_APP_PURCHASE ConversionActionTypeEnum_ConversionActionType = 13
	// Android app custom conversions tracked through Firebase.
	ConversionActionTypeEnum_FIREBASE_ANDROID_CUSTOM ConversionActionTypeEnum_ConversionActionType = 14
	// iOS app first open conversions tracked through Firebase.
	ConversionActionTypeEnum_FIREBASE_IOS_FIRST_OPEN ConversionActionTypeEnum_ConversionActionType = 15
	// iOS app in app purchase conversions tracked through Firebase.
	ConversionActionTypeEnum_FIREBASE_IOS_IN_APP_PURCHASE ConversionActionTypeEnum_ConversionActionType = 16
	// iOS app custom conversions tracked through Firebase.
	ConversionActionTypeEnum_FIREBASE_IOS_CUSTOM ConversionActionTypeEnum_ConversionActionType = 17
	// Android app first open conversions tracked through Third Party App
	// Analytics.
	ConversionActionTypeEnum_THIRD_PARTY_APP_ANALYTICS_ANDROID_FIRST_OPEN ConversionActionTypeEnum_ConversionActionType = 18
	// Android app in app purchase conversions tracked through Third Party App
	// Analytics.
	ConversionActionTypeEnum_THIRD_PARTY_APP_ANALYTICS_ANDROID_IN_APP_PURCHASE ConversionActionTypeEnum_ConversionActionType = 19
	// Android app custom conversions tracked through Third Party App Analytics.
	ConversionActionTypeEnum_THIRD_PARTY_APP_ANALYTICS_ANDROID_CUSTOM ConversionActionTypeEnum_ConversionActionType = 20
	// iOS app first open conversions tracked through Third Party App Analytics.
	ConversionActionTypeEnum_THIRD_PARTY_APP_ANALYTICS_IOS_FIRST_OPEN ConversionActionTypeEnum_ConversionActionType = 21
	// iOS app in app purchase conversions tracked through Third Party App
	// Analytics.
	ConversionActionTypeEnum_THIRD_PARTY_APP_ANALYTICS_IOS_IN_APP_PURCHASE ConversionActionTypeEnum_ConversionActionType = 22
	// iOS app custom conversions tracked through Third Party App Analytics.
	ConversionActionTypeEnum_THIRD_PARTY_APP_ANALYTICS_IOS_CUSTOM ConversionActionTypeEnum_ConversionActionType = 23
	// Conversions that occur when a user pre-registers a mobile app from the
	// Google Play Store. Read only.
	ConversionActionTypeEnum_ANDROID_APP_PRE_REGISTRATION ConversionActionTypeEnum_ConversionActionType = 24
	// Conversions that track all Google Play downloads which aren't tracked
	// by an app-specific type. Read only.
	ConversionActionTypeEnum_ANDROID_INSTALLS_ALL_OTHER_APPS ConversionActionTypeEnum_ConversionActionType = 25
	// Floodlight activity that counts the number of times that users have
	// visited a particular webpage after seeing or clicking on one of
	// an advertiser's ads. Read only.
	ConversionActionTypeEnum_FLOODLIGHT_ACTION ConversionActionTypeEnum_ConversionActionType = 26
	// Floodlight activity that tracks the number of sales made or the number
	// of items purchased. Can also capture the total value of each sale.
	// Read only.
	ConversionActionTypeEnum_FLOODLIGHT_TRANSACTION ConversionActionTypeEnum_ConversionActionType = 27
	// Conversions that track local actions from Google's products and
	// services after interacting with an ad. Read only.
	ConversionActionTypeEnum_GOOGLE_HOSTED ConversionActionTypeEnum_ConversionActionType = 28
	// Conversions reported when a user submits a lead form. Read only.
	ConversionActionTypeEnum_LEAD_FORM_SUBMIT ConversionActionTypeEnum_ConversionActionType = 29
	// Deprecated: The Salesforce integration will be going away and
	// replaced with an improved way to import your conversions from Salesforce.
	// - see https://support.google.com/google-ads/answer/14728349
	//
	// Deprecated: Marked as deprecated in google/ads/googleads/v20/enums/conversion_action_type.proto.
	ConversionActionTypeEnum_SALESFORCE ConversionActionTypeEnum_ConversionActionType = 30
	// Conversions imported from Search Ads 360 Floodlight data. Read only.
	ConversionActionTypeEnum_SEARCH_ADS_360 ConversionActionTypeEnum_ConversionActionType = 31
	// Call conversions that occur on Smart campaign Ads without call tracking
	// setup, using Smart campaign custom criteria. Read only.
	ConversionActionTypeEnum_SMART_CAMPAIGN_AD_CLICKS_TO_CALL ConversionActionTypeEnum_ConversionActionType = 32
	// The user clicks on a call element within Google Maps. Smart campaign
	// only. Read only.
	ConversionActionTypeEnum_SMART_CAMPAIGN_MAP_CLICKS_TO_CALL ConversionActionTypeEnum_ConversionActionType = 33
	// The user requests directions to a business location within Google Maps.
	// Smart campaign only. Read only.
	ConversionActionTypeEnum_SMART_CAMPAIGN_MAP_DIRECTIONS ConversionActionTypeEnum_ConversionActionType = 34
	// Call conversions that occur on Smart campaign Ads with call tracking
	// setup, using Smart campaign custom criteria. Read only.
	ConversionActionTypeEnum_SMART_CAMPAIGN_TRACKED_CALLS ConversionActionTypeEnum_ConversionActionType = 35
	// Conversions that occur when a user visits an advertiser's retail store.
	// Read only.
	ConversionActionTypeEnum_STORE_VISITS ConversionActionTypeEnum_ConversionActionType = 36
	// Conversions created from website events (such as form submissions or page
	// loads), that don't use individually coded event snippets. Read only.
	ConversionActionTypeEnum_WEBPAGE_CODELESS ConversionActionTypeEnum_ConversionActionType = 37
	// Conversions that come from linked Universal Analytics goals.
	ConversionActionTypeEnum_UNIVERSAL_ANALYTICS_GOAL ConversionActionTypeEnum_ConversionActionType = 38
	// Conversions that come from linked Universal Analytics transactions.
	ConversionActionTypeEnum_UNIVERSAL_ANALYTICS_TRANSACTION ConversionActionTypeEnum_ConversionActionType = 39
	// Conversions that come from linked Google Analytics 4 custom event
	// conversions.
	ConversionActionTypeEnum_GOOGLE_ANALYTICS_4_CUSTOM ConversionActionTypeEnum_ConversionActionType = 40
	// Conversions that come from linked Google Analytics 4 purchase
	// conversions.
	ConversionActionTypeEnum_GOOGLE_ANALYTICS_4_PURCHASE ConversionActionTypeEnum_ConversionActionType = 41
)

// Enum value maps for ConversionActionTypeEnum_ConversionActionType.
var (
	ConversionActionTypeEnum_ConversionActionType_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "AD_CALL",
		3:  "CLICK_TO_CALL",
		4:  "GOOGLE_PLAY_DOWNLOAD",
		5:  "GOOGLE_PLAY_IN_APP_PURCHASE",
		6:  "UPLOAD_CALLS",
		7:  "UPLOAD_CLICKS",
		8:  "WEBPAGE",
		9:  "WEBSITE_CALL",
		10: "STORE_SALES_DIRECT_UPLOAD",
		11: "STORE_SALES",
		12: "FIREBASE_ANDROID_FIRST_OPEN",
		13: "FIREBASE_ANDROID_IN_APP_PURCHASE",
		14: "FIREBASE_ANDROID_CUSTOM",
		15: "FIREBASE_IOS_FIRST_OPEN",
		16: "FIREBASE_IOS_IN_APP_PURCHASE",
		17: "FIREBASE_IOS_CUSTOM",
		18: "THIRD_PARTY_APP_ANALYTICS_ANDROID_FIRST_OPEN",
		19: "THIRD_PARTY_APP_ANALYTICS_ANDROID_IN_APP_PURCHASE",
		20: "THIRD_PARTY_APP_ANALYTICS_ANDROID_CUSTOM",
		21: "THIRD_PARTY_APP_ANALYTICS_IOS_FIRST_OPEN",
		22: "THIRD_PARTY_APP_ANALYTICS_IOS_IN_APP_PURCHASE",
		23: "THIRD_PARTY_APP_ANALYTICS_IOS_CUSTOM",
		24: "ANDROID_APP_PRE_REGISTRATION",
		25: "ANDROID_INSTALLS_ALL_OTHER_APPS",
		26: "FLOODLIGHT_ACTION",
		27: "FLOODLIGHT_TRANSACTION",
		28: "GOOGLE_HOSTED",
		29: "LEAD_FORM_SUBMIT",
		30: "SALESFORCE",
		31: "SEARCH_ADS_360",
		32: "SMART_CAMPAIGN_AD_CLICKS_TO_CALL",
		33: "SMART_CAMPAIGN_MAP_CLICKS_TO_CALL",
		34: "SMART_CAMPAIGN_MAP_DIRECTIONS",
		35: "SMART_CAMPAIGN_TRACKED_CALLS",
		36: "STORE_VISITS",
		37: "WEBPAGE_CODELESS",
		38: "UNIVERSAL_ANALYTICS_GOAL",
		39: "UNIVERSAL_ANALYTICS_TRANSACTION",
		40: "GOOGLE_ANALYTICS_4_CUSTOM",
		41: "GOOGLE_ANALYTICS_4_PURCHASE",
	}
	ConversionActionTypeEnum_ConversionActionType_value = map[string]int32{
		"UNSPECIFIED":                      0,
		"UNKNOWN":                          1,
		"AD_CALL":                          2,
		"CLICK_TO_CALL":                    3,
		"GOOGLE_PLAY_DOWNLOAD":             4,
		"GOOGLE_PLAY_IN_APP_PURCHASE":      5,
		"UPLOAD_CALLS":                     6,
		"UPLOAD_CLICKS":                    7,
		"WEBPAGE":                          8,
		"WEBSITE_CALL":                     9,
		"STORE_SALES_DIRECT_UPLOAD":        10,
		"STORE_SALES":                      11,
		"FIREBASE_ANDROID_FIRST_OPEN":      12,
		"FIREBASE_ANDROID_IN_APP_PURCHASE": 13,
		"FIREBASE_ANDROID_CUSTOM":          14,
		"FIREBASE_IOS_FIRST_OPEN":          15,
		"FIREBASE_IOS_IN_APP_PURCHASE":     16,
		"FIREBASE_IOS_CUSTOM":              17,
		"THIRD_PARTY_APP_ANALYTICS_ANDROID_FIRST_OPEN":      18,
		"THIRD_PARTY_APP_ANALYTICS_ANDROID_IN_APP_PURCHASE": 19,
		"THIRD_PARTY_APP_ANALYTICS_ANDROID_CUSTOM":          20,
		"THIRD_PARTY_APP_ANALYTICS_IOS_FIRST_OPEN":          21,
		"THIRD_PARTY_APP_ANALYTICS_IOS_IN_APP_PURCHASE":     22,
		"THIRD_PARTY_APP_ANALYTICS_IOS_CUSTOM":              23,
		"ANDROID_APP_PRE_REGISTRATION":                      24,
		"ANDROID_INSTALLS_ALL_OTHER_APPS":                   25,
		"FLOODLIGHT_ACTION":                                 26,
		"FLOODLIGHT_TRANSACTION":                            27,
		"GOOGLE_HOSTED":                                     28,
		"LEAD_FORM_SUBMIT":                                  29,
		"SALESFORCE":                                        30,
		"SEARCH_ADS_360":                                    31,
		"SMART_CAMPAIGN_AD_CLICKS_TO_CALL":                  32,
		"SMART_CAMPAIGN_MAP_CLICKS_TO_CALL":                 33,
		"SMART_CAMPAIGN_MAP_DIRECTIONS":                     34,
		"SMART_CAMPAIGN_TRACKED_CALLS":                      35,
		"STORE_VISITS":                                      36,
		"WEBPAGE_CODELESS":                                  37,
		"UNIVERSAL_ANALYTICS_GOAL":                          38,
		"UNIVERSAL_ANALYTICS_TRANSACTION":                   39,
		"GOOGLE_ANALYTICS_4_CUSTOM":                         40,
		"GOOGLE_ANALYTICS_4_PURCHASE":                       41,
	}
)

func (x ConversionActionTypeEnum_ConversionActionType) Enum() *ConversionActionTypeEnum_ConversionActionType {
	p := new(ConversionActionTypeEnum_ConversionActionType)
	*p = x
	return p
}

func (x ConversionActionTypeEnum_ConversionActionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConversionActionTypeEnum_ConversionActionType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v20_enums_conversion_action_type_proto_enumTypes[0].Descriptor()
}

func (ConversionActionTypeEnum_ConversionActionType) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v20_enums_conversion_action_type_proto_enumTypes[0]
}

func (x ConversionActionTypeEnum_ConversionActionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConversionActionTypeEnum_ConversionActionType.Descriptor instead.
func (ConversionActionTypeEnum_ConversionActionType) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_enums_conversion_action_type_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible types of a conversion action.
type ConversionActionTypeEnum struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConversionActionTypeEnum) Reset() {
	*x = ConversionActionTypeEnum{}
	mi := &file_google_ads_googleads_v20_enums_conversion_action_type_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConversionActionTypeEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConversionActionTypeEnum) ProtoMessage() {}

func (x *ConversionActionTypeEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_enums_conversion_action_type_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConversionActionTypeEnum.ProtoReflect.Descriptor instead.
func (*ConversionActionTypeEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_enums_conversion_action_type_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v20_enums_conversion_action_type_proto protoreflect.FileDescriptor

const file_google_ads_googleads_v20_enums_conversion_action_type_proto_rawDesc = "" +
	"\n" +
	";google/ads/googleads/v20/enums/conversion_action_type.proto\x12\x1egoogle.ads.googleads.v20.enums\"\x86\n" +
	"\n" +
	"\x18ConversionActionTypeEnum\"\xe9\t\n" +
	"\x14ConversionActionType\x12\x0f\n" +
	"\vUNSPECIFIED\x10\x00\x12\v\n" +
	"\aUNKNOWN\x10\x01\x12\v\n" +
	"\aAD_CALL\x10\x02\x12\x11\n" +
	"\rCLICK_TO_CALL\x10\x03\x12\x18\n" +
	"\x14GOOGLE_PLAY_DOWNLOAD\x10\x04\x12\x1f\n" +
	"\x1bGOOGLE_PLAY_IN_APP_PURCHASE\x10\x05\x12\x10\n" +
	"\fUPLOAD_CALLS\x10\x06\x12\x11\n" +
	"\rUPLOAD_CLICKS\x10\a\x12\v\n" +
	"\aWEBPAGE\x10\b\x12\x10\n" +
	"\fWEBSITE_CALL\x10\t\x12\x1d\n" +
	"\x19STORE_SALES_DIRECT_UPLOAD\x10\n" +
	"\x12\x0f\n" +
	"\vSTORE_SALES\x10\v\x12\x1f\n" +
	"\x1bFIREBASE_ANDROID_FIRST_OPEN\x10\f\x12$\n" +
	" FIREBASE_ANDROID_IN_APP_PURCHASE\x10\r\x12\x1b\n" +
	"\x17FIREBASE_ANDROID_CUSTOM\x10\x0e\x12\x1b\n" +
	"\x17FIREBASE_IOS_FIRST_OPEN\x10\x0f\x12 \n" +
	"\x1cFIREBASE_IOS_IN_APP_PURCHASE\x10\x10\x12\x17\n" +
	"\x13FIREBASE_IOS_CUSTOM\x10\x11\x120\n" +
	",THIRD_PARTY_APP_ANALYTICS_ANDROID_FIRST_OPEN\x10\x12\x125\n" +
	"1THIRD_PARTY_APP_ANALYTICS_ANDROID_IN_APP_PURCHASE\x10\x13\x12,\n" +
	"(THIRD_PARTY_APP_ANALYTICS_ANDROID_CUSTOM\x10\x14\x12,\n" +
	"(THIRD_PARTY_APP_ANALYTICS_IOS_FIRST_OPEN\x10\x15\x121\n" +
	"-THIRD_PARTY_APP_ANALYTICS_IOS_IN_APP_PURCHASE\x10\x16\x12(\n" +
	"$THIRD_PARTY_APP_ANALYTICS_IOS_CUSTOM\x10\x17\x12 \n" +
	"\x1cANDROID_APP_PRE_REGISTRATION\x10\x18\x12#\n" +
	"\x1fANDROID_INSTALLS_ALL_OTHER_APPS\x10\x19\x12\x15\n" +
	"\x11FLOODLIGHT_ACTION\x10\x1a\x12\x1a\n" +
	"\x16FLOODLIGHT_TRANSACTION\x10\x1b\x12\x11\n" +
	"\rGOOGLE_HOSTED\x10\x1c\x12\x14\n" +
	"\x10LEAD_FORM_SUBMIT\x10\x1d\x12\x12\n" +
	"\n" +
	"SALESFORCE\x10\x1e\x1a\x02\b\x01\x12\x12\n" +
	"\x0eSEARCH_ADS_360\x10\x1f\x12$\n" +
	" SMART_CAMPAIGN_AD_CLICKS_TO_CALL\x10 \x12%\n" +
	"!SMART_CAMPAIGN_MAP_CLICKS_TO_CALL\x10!\x12!\n" +
	"\x1dSMART_CAMPAIGN_MAP_DIRECTIONS\x10\"\x12 \n" +
	"\x1cSMART_CAMPAIGN_TRACKED_CALLS\x10#\x12\x10\n" +
	"\fSTORE_VISITS\x10$\x12\x14\n" +
	"\x10WEBPAGE_CODELESS\x10%\x12\x1c\n" +
	"\x18UNIVERSAL_ANALYTICS_GOAL\x10&\x12#\n" +
	"\x1fUNIVERSAL_ANALYTICS_TRANSACTION\x10'\x12\x1d\n" +
	"\x19GOOGLE_ANALYTICS_4_CUSTOM\x10(\x12\x1f\n" +
	"\x1bGOOGLE_ANALYTICS_4_PURCHASE\x10)B\xf3\x01\n" +
	"\"com.google.ads.googleads.v20.enumsB\x19ConversionActionTypeProtoP\x01ZCgoogle.golang.org/genproto/googleapis/ads/googleads/v20/enums;enums\xa2\x02\x03GAA\xaa\x02\x1eGoogle.Ads.GoogleAds.V20.Enums\xca\x02\x1eGoogle\\Ads\\GoogleAds\\V20\\Enums\xea\x02\"Google::Ads::GoogleAds::V20::Enumsb\x06proto3"

var (
	file_google_ads_googleads_v20_enums_conversion_action_type_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v20_enums_conversion_action_type_proto_rawDescData []byte
)

func file_google_ads_googleads_v20_enums_conversion_action_type_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v20_enums_conversion_action_type_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v20_enums_conversion_action_type_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_enums_conversion_action_type_proto_rawDesc), len(file_google_ads_googleads_v20_enums_conversion_action_type_proto_rawDesc)))
	})
	return file_google_ads_googleads_v20_enums_conversion_action_type_proto_rawDescData
}

var file_google_ads_googleads_v20_enums_conversion_action_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v20_enums_conversion_action_type_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v20_enums_conversion_action_type_proto_goTypes = []any{
	(ConversionActionTypeEnum_ConversionActionType)(0), // 0: google.ads.googleads.v20.enums.ConversionActionTypeEnum.ConversionActionType
	(*ConversionActionTypeEnum)(nil),                   // 1: google.ads.googleads.v20.enums.ConversionActionTypeEnum
}
var file_google_ads_googleads_v20_enums_conversion_action_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v20_enums_conversion_action_type_proto_init() }
func file_google_ads_googleads_v20_enums_conversion_action_type_proto_init() {
	if File_google_ads_googleads_v20_enums_conversion_action_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_enums_conversion_action_type_proto_rawDesc), len(file_google_ads_googleads_v20_enums_conversion_action_type_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v20_enums_conversion_action_type_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v20_enums_conversion_action_type_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v20_enums_conversion_action_type_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v20_enums_conversion_action_type_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v20_enums_conversion_action_type_proto = out.File
	file_google_ads_googleads_v20_enums_conversion_action_type_proto_goTypes = nil
	file_google_ads_googleads_v20_enums_conversion_action_type_proto_depIdxs = nil
}

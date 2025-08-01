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
// source: google/ads/googleads/v20/errors/campaign_criterion_error.proto

package errors

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

// Enum describing possible campaign criterion errors.
type CampaignCriterionErrorEnum_CampaignCriterionError int32

const (
	// Enum unspecified.
	CampaignCriterionErrorEnum_UNSPECIFIED CampaignCriterionErrorEnum_CampaignCriterionError = 0
	// The received error code is not known in this version.
	CampaignCriterionErrorEnum_UNKNOWN CampaignCriterionErrorEnum_CampaignCriterionError = 1
	// Concrete type of criterion (keyword v.s. placement) is required for
	// CREATE and UPDATE operations.
	CampaignCriterionErrorEnum_CONCRETE_TYPE_REQUIRED CampaignCriterionErrorEnum_CampaignCriterionError = 2
	// Invalid placement URL.
	CampaignCriterionErrorEnum_INVALID_PLACEMENT_URL CampaignCriterionErrorEnum_CampaignCriterionError = 3
	// Criteria type can not be excluded for the campaign by the customer. like
	// AOL account type cannot target site type criteria
	CampaignCriterionErrorEnum_CANNOT_EXCLUDE_CRITERIA_TYPE CampaignCriterionErrorEnum_CampaignCriterionError = 4
	// Cannot set the campaign criterion status for this criteria type.
	CampaignCriterionErrorEnum_CANNOT_SET_STATUS_FOR_CRITERIA_TYPE CampaignCriterionErrorEnum_CampaignCriterionError = 5
	// Cannot set the campaign criterion status for an excluded criteria.
	CampaignCriterionErrorEnum_CANNOT_SET_STATUS_FOR_EXCLUDED_CRITERIA CampaignCriterionErrorEnum_CampaignCriterionError = 6
	// Cannot target and exclude the same criterion.
	CampaignCriterionErrorEnum_CANNOT_TARGET_AND_EXCLUDE CampaignCriterionErrorEnum_CampaignCriterionError = 7
	// The mutate contained too many operations.
	CampaignCriterionErrorEnum_TOO_MANY_OPERATIONS CampaignCriterionErrorEnum_CampaignCriterionError = 8
	// This operator cannot be applied to a criterion of this type.
	CampaignCriterionErrorEnum_OPERATOR_NOT_SUPPORTED_FOR_CRITERION_TYPE CampaignCriterionErrorEnum_CampaignCriterionError = 9
	// The Shopping campaign sales country is not supported for
	// ProductSalesChannel targeting.
	CampaignCriterionErrorEnum_SHOPPING_CAMPAIGN_SALES_COUNTRY_NOT_SUPPORTED_FOR_SALES_CHANNEL CampaignCriterionErrorEnum_CampaignCriterionError = 10
	// The existing field can't be updated with CREATE operation. It can be
	// updated with UPDATE operation only.
	CampaignCriterionErrorEnum_CANNOT_ADD_EXISTING_FIELD CampaignCriterionErrorEnum_CampaignCriterionError = 11
	// Negative criteria are immutable, so updates are not allowed.
	CampaignCriterionErrorEnum_CANNOT_UPDATE_NEGATIVE_CRITERION CampaignCriterionErrorEnum_CampaignCriterionError = 12
	// Only free form names are allowed for negative Smart campaign keyword
	// theme.
	CampaignCriterionErrorEnum_CANNOT_SET_NEGATIVE_KEYWORD_THEME_CONSTANT_CRITERION CampaignCriterionErrorEnum_CampaignCriterionError = 13
	// Invalid Smart campaign keyword theme constant criterion.
	CampaignCriterionErrorEnum_INVALID_KEYWORD_THEME_CONSTANT CampaignCriterionErrorEnum_CampaignCriterionError = 14
	// A Smart campaign keyword theme constant or free-form Smart campaign
	// keyword theme is required.
	CampaignCriterionErrorEnum_MISSING_KEYWORD_THEME_CONSTANT_OR_FREE_FORM_KEYWORD_THEME CampaignCriterionErrorEnum_CampaignCriterionError = 15
	// A Smart campaign may not target proximity and location criteria
	// simultaneously.
	CampaignCriterionErrorEnum_CANNOT_TARGET_BOTH_PROXIMITY_AND_LOCATION_CRITERIA_FOR_SMART_CAMPAIGN CampaignCriterionErrorEnum_CampaignCriterionError = 16
	// A Smart campaign may not target multiple proximity criteria.
	CampaignCriterionErrorEnum_CANNOT_TARGET_MULTIPLE_PROXIMITY_CRITERIA_FOR_SMART_CAMPAIGN CampaignCriterionErrorEnum_CampaignCriterionError = 17
	// Location is not launched for Local Services Campaigns.
	CampaignCriterionErrorEnum_LOCATION_NOT_LAUNCHED_FOR_LOCAL_SERVICES_CAMPAIGN CampaignCriterionErrorEnum_CampaignCriterionError = 18
	// A Local Services campaign may not target certain criteria types.
	CampaignCriterionErrorEnum_LOCATION_INVALID_FOR_LOCAL_SERVICES_CAMPAIGN CampaignCriterionErrorEnum_CampaignCriterionError = 19
	// Country locations are not supported for Local Services campaign.
	CampaignCriterionErrorEnum_CANNOT_TARGET_COUNTRY_FOR_LOCAL_SERVICES_CAMPAIGN CampaignCriterionErrorEnum_CampaignCriterionError = 20
	// Location is not within the home country of Local Services campaign.
	CampaignCriterionErrorEnum_LOCATION_NOT_IN_HOME_COUNTRY_FOR_LOCAL_SERVICES_CAMPAIGN CampaignCriterionErrorEnum_CampaignCriterionError = 21
	// Local Services profile does not exist for a particular Local Services
	// campaign.
	CampaignCriterionErrorEnum_CANNOT_ADD_OR_REMOVE_LOCATION_FOR_LOCAL_SERVICES_CAMPAIGN CampaignCriterionErrorEnum_CampaignCriterionError = 22
	// Local Services campaign must have at least one target location.
	CampaignCriterionErrorEnum_AT_LEAST_ONE_POSITIVE_LOCATION_REQUIRED_FOR_LOCAL_SERVICES_CAMPAIGN CampaignCriterionErrorEnum_CampaignCriterionError = 23
	// At least one positive local service ID criterion is required for a Local
	// Services campaign.
	CampaignCriterionErrorEnum_AT_LEAST_ONE_LOCAL_SERVICE_ID_CRITERION_REQUIRED_FOR_LOCAL_SERVICES_CAMPAIGN CampaignCriterionErrorEnum_CampaignCriterionError = 24
	// Local service ID is not found under selected categories in local
	// services campaign setting.
	CampaignCriterionErrorEnum_LOCAL_SERVICE_ID_NOT_FOUND_FOR_CATEGORY CampaignCriterionErrorEnum_CampaignCriterionError = 25
	// For search advertising channel, brand lists can only be applied to
	// exclusive targeting, broad match campaigns for inclusive targeting or
	// PMax generated campaigns.
	CampaignCriterionErrorEnum_CANNOT_ATTACH_BRAND_LIST_TO_NON_QUALIFIED_SEARCH_CAMPAIGN CampaignCriterionErrorEnum_CampaignCriterionError = 26
	// Campaigns that target all countries and territories are limited to a
	// certain number of top-level location exclusions. If removing a criterion
	// causes the campaign to target all countries and territories and the
	// campaign has more top-level location exclusions than the limit allows,
	// then this error is returned.
	CampaignCriterionErrorEnum_CANNOT_REMOVE_ALL_LOCATIONS_DUE_TO_TOO_MANY_COUNTRY_EXCLUSIONS CampaignCriterionErrorEnum_CampaignCriterionError = 27
)

// Enum value maps for CampaignCriterionErrorEnum_CampaignCriterionError.
var (
	CampaignCriterionErrorEnum_CampaignCriterionError_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "CONCRETE_TYPE_REQUIRED",
		3:  "INVALID_PLACEMENT_URL",
		4:  "CANNOT_EXCLUDE_CRITERIA_TYPE",
		5:  "CANNOT_SET_STATUS_FOR_CRITERIA_TYPE",
		6:  "CANNOT_SET_STATUS_FOR_EXCLUDED_CRITERIA",
		7:  "CANNOT_TARGET_AND_EXCLUDE",
		8:  "TOO_MANY_OPERATIONS",
		9:  "OPERATOR_NOT_SUPPORTED_FOR_CRITERION_TYPE",
		10: "SHOPPING_CAMPAIGN_SALES_COUNTRY_NOT_SUPPORTED_FOR_SALES_CHANNEL",
		11: "CANNOT_ADD_EXISTING_FIELD",
		12: "CANNOT_UPDATE_NEGATIVE_CRITERION",
		13: "CANNOT_SET_NEGATIVE_KEYWORD_THEME_CONSTANT_CRITERION",
		14: "INVALID_KEYWORD_THEME_CONSTANT",
		15: "MISSING_KEYWORD_THEME_CONSTANT_OR_FREE_FORM_KEYWORD_THEME",
		16: "CANNOT_TARGET_BOTH_PROXIMITY_AND_LOCATION_CRITERIA_FOR_SMART_CAMPAIGN",
		17: "CANNOT_TARGET_MULTIPLE_PROXIMITY_CRITERIA_FOR_SMART_CAMPAIGN",
		18: "LOCATION_NOT_LAUNCHED_FOR_LOCAL_SERVICES_CAMPAIGN",
		19: "LOCATION_INVALID_FOR_LOCAL_SERVICES_CAMPAIGN",
		20: "CANNOT_TARGET_COUNTRY_FOR_LOCAL_SERVICES_CAMPAIGN",
		21: "LOCATION_NOT_IN_HOME_COUNTRY_FOR_LOCAL_SERVICES_CAMPAIGN",
		22: "CANNOT_ADD_OR_REMOVE_LOCATION_FOR_LOCAL_SERVICES_CAMPAIGN",
		23: "AT_LEAST_ONE_POSITIVE_LOCATION_REQUIRED_FOR_LOCAL_SERVICES_CAMPAIGN",
		24: "AT_LEAST_ONE_LOCAL_SERVICE_ID_CRITERION_REQUIRED_FOR_LOCAL_SERVICES_CAMPAIGN",
		25: "LOCAL_SERVICE_ID_NOT_FOUND_FOR_CATEGORY",
		26: "CANNOT_ATTACH_BRAND_LIST_TO_NON_QUALIFIED_SEARCH_CAMPAIGN",
		27: "CANNOT_REMOVE_ALL_LOCATIONS_DUE_TO_TOO_MANY_COUNTRY_EXCLUSIONS",
	}
	CampaignCriterionErrorEnum_CampaignCriterionError_value = map[string]int32{
		"UNSPECIFIED":                               0,
		"UNKNOWN":                                   1,
		"CONCRETE_TYPE_REQUIRED":                    2,
		"INVALID_PLACEMENT_URL":                     3,
		"CANNOT_EXCLUDE_CRITERIA_TYPE":              4,
		"CANNOT_SET_STATUS_FOR_CRITERIA_TYPE":       5,
		"CANNOT_SET_STATUS_FOR_EXCLUDED_CRITERIA":   6,
		"CANNOT_TARGET_AND_EXCLUDE":                 7,
		"TOO_MANY_OPERATIONS":                       8,
		"OPERATOR_NOT_SUPPORTED_FOR_CRITERION_TYPE": 9,
		"SHOPPING_CAMPAIGN_SALES_COUNTRY_NOT_SUPPORTED_FOR_SALES_CHANNEL":              10,
		"CANNOT_ADD_EXISTING_FIELD":                                                    11,
		"CANNOT_UPDATE_NEGATIVE_CRITERION":                                             12,
		"CANNOT_SET_NEGATIVE_KEYWORD_THEME_CONSTANT_CRITERION":                         13,
		"INVALID_KEYWORD_THEME_CONSTANT":                                               14,
		"MISSING_KEYWORD_THEME_CONSTANT_OR_FREE_FORM_KEYWORD_THEME":                    15,
		"CANNOT_TARGET_BOTH_PROXIMITY_AND_LOCATION_CRITERIA_FOR_SMART_CAMPAIGN":        16,
		"CANNOT_TARGET_MULTIPLE_PROXIMITY_CRITERIA_FOR_SMART_CAMPAIGN":                 17,
		"LOCATION_NOT_LAUNCHED_FOR_LOCAL_SERVICES_CAMPAIGN":                            18,
		"LOCATION_INVALID_FOR_LOCAL_SERVICES_CAMPAIGN":                                 19,
		"CANNOT_TARGET_COUNTRY_FOR_LOCAL_SERVICES_CAMPAIGN":                            20,
		"LOCATION_NOT_IN_HOME_COUNTRY_FOR_LOCAL_SERVICES_CAMPAIGN":                     21,
		"CANNOT_ADD_OR_REMOVE_LOCATION_FOR_LOCAL_SERVICES_CAMPAIGN":                    22,
		"AT_LEAST_ONE_POSITIVE_LOCATION_REQUIRED_FOR_LOCAL_SERVICES_CAMPAIGN":          23,
		"AT_LEAST_ONE_LOCAL_SERVICE_ID_CRITERION_REQUIRED_FOR_LOCAL_SERVICES_CAMPAIGN": 24,
		"LOCAL_SERVICE_ID_NOT_FOUND_FOR_CATEGORY":                                      25,
		"CANNOT_ATTACH_BRAND_LIST_TO_NON_QUALIFIED_SEARCH_CAMPAIGN":                    26,
		"CANNOT_REMOVE_ALL_LOCATIONS_DUE_TO_TOO_MANY_COUNTRY_EXCLUSIONS":               27,
	}
)

func (x CampaignCriterionErrorEnum_CampaignCriterionError) Enum() *CampaignCriterionErrorEnum_CampaignCriterionError {
	p := new(CampaignCriterionErrorEnum_CampaignCriterionError)
	*p = x
	return p
}

func (x CampaignCriterionErrorEnum_CampaignCriterionError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CampaignCriterionErrorEnum_CampaignCriterionError) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v20_errors_campaign_criterion_error_proto_enumTypes[0].Descriptor()
}

func (CampaignCriterionErrorEnum_CampaignCriterionError) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v20_errors_campaign_criterion_error_proto_enumTypes[0]
}

func (x CampaignCriterionErrorEnum_CampaignCriterionError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CampaignCriterionErrorEnum_CampaignCriterionError.Descriptor instead.
func (CampaignCriterionErrorEnum_CampaignCriterionError) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_errors_campaign_criterion_error_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible campaign criterion errors.
type CampaignCriterionErrorEnum struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CampaignCriterionErrorEnum) Reset() {
	*x = CampaignCriterionErrorEnum{}
	mi := &file_google_ads_googleads_v20_errors_campaign_criterion_error_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CampaignCriterionErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CampaignCriterionErrorEnum) ProtoMessage() {}

func (x *CampaignCriterionErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_errors_campaign_criterion_error_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CampaignCriterionErrorEnum.ProtoReflect.Descriptor instead.
func (*CampaignCriterionErrorEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_errors_campaign_criterion_error_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v20_errors_campaign_criterion_error_proto protoreflect.FileDescriptor

const file_google_ads_googleads_v20_errors_campaign_criterion_error_proto_rawDesc = "" +
	"\n" +
	">google/ads/googleads/v20/errors/campaign_criterion_error.proto\x12\x1fgoogle.ads.googleads.v20.errors\"\x87\v\n" +
	"\x1aCampaignCriterionErrorEnum\"\xe8\n" +
	"\n" +
	"\x16CampaignCriterionError\x12\x0f\n" +
	"\vUNSPECIFIED\x10\x00\x12\v\n" +
	"\aUNKNOWN\x10\x01\x12\x1a\n" +
	"\x16CONCRETE_TYPE_REQUIRED\x10\x02\x12\x19\n" +
	"\x15INVALID_PLACEMENT_URL\x10\x03\x12 \n" +
	"\x1cCANNOT_EXCLUDE_CRITERIA_TYPE\x10\x04\x12'\n" +
	"#CANNOT_SET_STATUS_FOR_CRITERIA_TYPE\x10\x05\x12+\n" +
	"'CANNOT_SET_STATUS_FOR_EXCLUDED_CRITERIA\x10\x06\x12\x1d\n" +
	"\x19CANNOT_TARGET_AND_EXCLUDE\x10\a\x12\x17\n" +
	"\x13TOO_MANY_OPERATIONS\x10\b\x12-\n" +
	")OPERATOR_NOT_SUPPORTED_FOR_CRITERION_TYPE\x10\t\x12C\n" +
	"?SHOPPING_CAMPAIGN_SALES_COUNTRY_NOT_SUPPORTED_FOR_SALES_CHANNEL\x10\n" +
	"\x12\x1d\n" +
	"\x19CANNOT_ADD_EXISTING_FIELD\x10\v\x12$\n" +
	" CANNOT_UPDATE_NEGATIVE_CRITERION\x10\f\x128\n" +
	"4CANNOT_SET_NEGATIVE_KEYWORD_THEME_CONSTANT_CRITERION\x10\r\x12\"\n" +
	"\x1eINVALID_KEYWORD_THEME_CONSTANT\x10\x0e\x12=\n" +
	"9MISSING_KEYWORD_THEME_CONSTANT_OR_FREE_FORM_KEYWORD_THEME\x10\x0f\x12I\n" +
	"ECANNOT_TARGET_BOTH_PROXIMITY_AND_LOCATION_CRITERIA_FOR_SMART_CAMPAIGN\x10\x10\x12@\n" +
	"<CANNOT_TARGET_MULTIPLE_PROXIMITY_CRITERIA_FOR_SMART_CAMPAIGN\x10\x11\x125\n" +
	"1LOCATION_NOT_LAUNCHED_FOR_LOCAL_SERVICES_CAMPAIGN\x10\x12\x120\n" +
	",LOCATION_INVALID_FOR_LOCAL_SERVICES_CAMPAIGN\x10\x13\x125\n" +
	"1CANNOT_TARGET_COUNTRY_FOR_LOCAL_SERVICES_CAMPAIGN\x10\x14\x12<\n" +
	"8LOCATION_NOT_IN_HOME_COUNTRY_FOR_LOCAL_SERVICES_CAMPAIGN\x10\x15\x12=\n" +
	"9CANNOT_ADD_OR_REMOVE_LOCATION_FOR_LOCAL_SERVICES_CAMPAIGN\x10\x16\x12G\n" +
	"CAT_LEAST_ONE_POSITIVE_LOCATION_REQUIRED_FOR_LOCAL_SERVICES_CAMPAIGN\x10\x17\x12P\n" +
	"LAT_LEAST_ONE_LOCAL_SERVICE_ID_CRITERION_REQUIRED_FOR_LOCAL_SERVICES_CAMPAIGN\x10\x18\x12+\n" +
	"'LOCAL_SERVICE_ID_NOT_FOUND_FOR_CATEGORY\x10\x19\x12=\n" +
	"9CANNOT_ATTACH_BRAND_LIST_TO_NON_QUALIFIED_SEARCH_CAMPAIGN\x10\x1a\x12B\n" +
	">CANNOT_REMOVE_ALL_LOCATIONS_DUE_TO_TOO_MANY_COUNTRY_EXCLUSIONS\x10\x1bB\xfb\x01\n" +
	"#com.google.ads.googleads.v20.errorsB\x1bCampaignCriterionErrorProtoP\x01ZEgoogle.golang.org/genproto/googleapis/ads/googleads/v20/errors;errors\xa2\x02\x03GAA\xaa\x02\x1fGoogle.Ads.GoogleAds.V20.Errors\xca\x02\x1fGoogle\\Ads\\GoogleAds\\V20\\Errors\xea\x02#Google::Ads::GoogleAds::V20::Errorsb\x06proto3"

var (
	file_google_ads_googleads_v20_errors_campaign_criterion_error_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v20_errors_campaign_criterion_error_proto_rawDescData []byte
)

func file_google_ads_googleads_v20_errors_campaign_criterion_error_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v20_errors_campaign_criterion_error_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v20_errors_campaign_criterion_error_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_errors_campaign_criterion_error_proto_rawDesc), len(file_google_ads_googleads_v20_errors_campaign_criterion_error_proto_rawDesc)))
	})
	return file_google_ads_googleads_v20_errors_campaign_criterion_error_proto_rawDescData
}

var file_google_ads_googleads_v20_errors_campaign_criterion_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v20_errors_campaign_criterion_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v20_errors_campaign_criterion_error_proto_goTypes = []any{
	(CampaignCriterionErrorEnum_CampaignCriterionError)(0), // 0: google.ads.googleads.v20.errors.CampaignCriterionErrorEnum.CampaignCriterionError
	(*CampaignCriterionErrorEnum)(nil),                     // 1: google.ads.googleads.v20.errors.CampaignCriterionErrorEnum
}
var file_google_ads_googleads_v20_errors_campaign_criterion_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v20_errors_campaign_criterion_error_proto_init() }
func file_google_ads_googleads_v20_errors_campaign_criterion_error_proto_init() {
	if File_google_ads_googleads_v20_errors_campaign_criterion_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_errors_campaign_criterion_error_proto_rawDesc), len(file_google_ads_googleads_v20_errors_campaign_criterion_error_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v20_errors_campaign_criterion_error_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v20_errors_campaign_criterion_error_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v20_errors_campaign_criterion_error_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v20_errors_campaign_criterion_error_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v20_errors_campaign_criterion_error_proto = out.File
	file_google_ads_googleads_v20_errors_campaign_criterion_error_proto_goTypes = nil
	file_google_ads_googleads_v20_errors_campaign_criterion_error_proto_depIdxs = nil
}

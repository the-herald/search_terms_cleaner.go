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
// source: google/ads/googleads/v20/errors/setting_error.proto

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

// Enum describing possible setting errors.
type SettingErrorEnum_SettingError int32

const (
	// Enum unspecified.
	SettingErrorEnum_UNSPECIFIED SettingErrorEnum_SettingError = 0
	// The received error code is not known in this version.
	SettingErrorEnum_UNKNOWN SettingErrorEnum_SettingError = 1
	// The campaign setting is not available for this Google Ads account.
	SettingErrorEnum_SETTING_TYPE_IS_NOT_AVAILABLE SettingErrorEnum_SettingError = 3
	// The setting is not compatible with the campaign.
	SettingErrorEnum_SETTING_TYPE_IS_NOT_COMPATIBLE_WITH_CAMPAIGN SettingErrorEnum_SettingError = 4
	// The supplied TargetingSetting contains an invalid CriterionTypeGroup. See
	// CriterionTypeGroup documentation for CriterionTypeGroups allowed
	// in Campaign or AdGroup TargetingSettings.
	SettingErrorEnum_TARGETING_SETTING_CONTAINS_INVALID_CRITERION_TYPE_GROUP SettingErrorEnum_SettingError = 5
	// TargetingSetting must not explicitly
	// set any of the Demographic CriterionTypeGroups (AGE_RANGE, GENDER,
	// PARENT, INCOME_RANGE) to false (it's okay to not set them at all, in
	// which case the system will set them to true automatically).
	SettingErrorEnum_TARGETING_SETTING_DEMOGRAPHIC_CRITERION_TYPE_GROUPS_MUST_BE_SET_TO_TARGET_ALL SettingErrorEnum_SettingError = 6
	// TargetingSetting cannot change any of
	// the Demographic CriterionTypeGroups (AGE_RANGE, GENDER, PARENT,
	// INCOME_RANGE) from true to false.
	SettingErrorEnum_TARGETING_SETTING_CANNOT_CHANGE_TARGET_ALL_TO_FALSE_FOR_DEMOGRAPHIC_CRITERION_TYPE_GROUP SettingErrorEnum_SettingError = 7
	// At least one feed id should be present.
	SettingErrorEnum_DYNAMIC_SEARCH_ADS_SETTING_AT_LEAST_ONE_FEED_ID_MUST_BE_PRESENT SettingErrorEnum_SettingError = 8
	// The supplied DynamicSearchAdsSetting contains an invalid domain name.
	SettingErrorEnum_DYNAMIC_SEARCH_ADS_SETTING_CONTAINS_INVALID_DOMAIN_NAME SettingErrorEnum_SettingError = 9
	// The supplied DynamicSearchAdsSetting contains a subdomain name.
	SettingErrorEnum_DYNAMIC_SEARCH_ADS_SETTING_CONTAINS_SUBDOMAIN_NAME SettingErrorEnum_SettingError = 10
	// The supplied DynamicSearchAdsSetting contains an invalid language code.
	SettingErrorEnum_DYNAMIC_SEARCH_ADS_SETTING_CONTAINS_INVALID_LANGUAGE_CODE SettingErrorEnum_SettingError = 11
	// TargetingSettings in search campaigns should not have
	// CriterionTypeGroup.PLACEMENT set to targetAll.
	SettingErrorEnum_TARGET_ALL_IS_NOT_ALLOWED_FOR_PLACEMENT_IN_SEARCH_CAMPAIGN SettingErrorEnum_SettingError = 12
	// The setting value is not compatible with the campaign type.
	SettingErrorEnum_SETTING_VALUE_NOT_COMPATIBLE_WITH_CAMPAIGN SettingErrorEnum_SettingError = 20
	// Switching from observation setting to targeting setting is not allowed
	// for Customer Match lists. See
	// https://support.google.com/google-ads/answer/6299717.
	SettingErrorEnum_BID_ONLY_IS_NOT_ALLOWED_TO_BE_MODIFIED_WITH_CUSTOMER_MATCH_TARGETING SettingErrorEnum_SettingError = 21
)

// Enum value maps for SettingErrorEnum_SettingError.
var (
	SettingErrorEnum_SettingError_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		3:  "SETTING_TYPE_IS_NOT_AVAILABLE",
		4:  "SETTING_TYPE_IS_NOT_COMPATIBLE_WITH_CAMPAIGN",
		5:  "TARGETING_SETTING_CONTAINS_INVALID_CRITERION_TYPE_GROUP",
		6:  "TARGETING_SETTING_DEMOGRAPHIC_CRITERION_TYPE_GROUPS_MUST_BE_SET_TO_TARGET_ALL",
		7:  "TARGETING_SETTING_CANNOT_CHANGE_TARGET_ALL_TO_FALSE_FOR_DEMOGRAPHIC_CRITERION_TYPE_GROUP",
		8:  "DYNAMIC_SEARCH_ADS_SETTING_AT_LEAST_ONE_FEED_ID_MUST_BE_PRESENT",
		9:  "DYNAMIC_SEARCH_ADS_SETTING_CONTAINS_INVALID_DOMAIN_NAME",
		10: "DYNAMIC_SEARCH_ADS_SETTING_CONTAINS_SUBDOMAIN_NAME",
		11: "DYNAMIC_SEARCH_ADS_SETTING_CONTAINS_INVALID_LANGUAGE_CODE",
		12: "TARGET_ALL_IS_NOT_ALLOWED_FOR_PLACEMENT_IN_SEARCH_CAMPAIGN",
		20: "SETTING_VALUE_NOT_COMPATIBLE_WITH_CAMPAIGN",
		21: "BID_ONLY_IS_NOT_ALLOWED_TO_BE_MODIFIED_WITH_CUSTOMER_MATCH_TARGETING",
	}
	SettingErrorEnum_SettingError_value = map[string]int32{
		"UNSPECIFIED":                   0,
		"UNKNOWN":                       1,
		"SETTING_TYPE_IS_NOT_AVAILABLE": 3,
		"SETTING_TYPE_IS_NOT_COMPATIBLE_WITH_CAMPAIGN":                                             4,
		"TARGETING_SETTING_CONTAINS_INVALID_CRITERION_TYPE_GROUP":                                  5,
		"TARGETING_SETTING_DEMOGRAPHIC_CRITERION_TYPE_GROUPS_MUST_BE_SET_TO_TARGET_ALL":            6,
		"TARGETING_SETTING_CANNOT_CHANGE_TARGET_ALL_TO_FALSE_FOR_DEMOGRAPHIC_CRITERION_TYPE_GROUP": 7,
		"DYNAMIC_SEARCH_ADS_SETTING_AT_LEAST_ONE_FEED_ID_MUST_BE_PRESENT":                          8,
		"DYNAMIC_SEARCH_ADS_SETTING_CONTAINS_INVALID_DOMAIN_NAME":                                  9,
		"DYNAMIC_SEARCH_ADS_SETTING_CONTAINS_SUBDOMAIN_NAME":                                       10,
		"DYNAMIC_SEARCH_ADS_SETTING_CONTAINS_INVALID_LANGUAGE_CODE":                                11,
		"TARGET_ALL_IS_NOT_ALLOWED_FOR_PLACEMENT_IN_SEARCH_CAMPAIGN":                               12,
		"SETTING_VALUE_NOT_COMPATIBLE_WITH_CAMPAIGN":                                               20,
		"BID_ONLY_IS_NOT_ALLOWED_TO_BE_MODIFIED_WITH_CUSTOMER_MATCH_TARGETING":                     21,
	}
)

func (x SettingErrorEnum_SettingError) Enum() *SettingErrorEnum_SettingError {
	p := new(SettingErrorEnum_SettingError)
	*p = x
	return p
}

func (x SettingErrorEnum_SettingError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SettingErrorEnum_SettingError) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v20_errors_setting_error_proto_enumTypes[0].Descriptor()
}

func (SettingErrorEnum_SettingError) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v20_errors_setting_error_proto_enumTypes[0]
}

func (x SettingErrorEnum_SettingError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SettingErrorEnum_SettingError.Descriptor instead.
func (SettingErrorEnum_SettingError) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_errors_setting_error_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible setting errors.
type SettingErrorEnum struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SettingErrorEnum) Reset() {
	*x = SettingErrorEnum{}
	mi := &file_google_ads_googleads_v20_errors_setting_error_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SettingErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SettingErrorEnum) ProtoMessage() {}

func (x *SettingErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_errors_setting_error_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SettingErrorEnum.ProtoReflect.Descriptor instead.
func (*SettingErrorEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_errors_setting_error_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v20_errors_setting_error_proto protoreflect.FileDescriptor

const file_google_ads_googleads_v20_errors_setting_error_proto_rawDesc = "" +
	"\n" +
	"3google/ads/googleads/v20/errors/setting_error.proto\x12\x1fgoogle.ads.googleads.v20.errors\"\xb7\x06\n" +
	"\x10SettingErrorEnum\"\xa2\x06\n" +
	"\fSettingError\x12\x0f\n" +
	"\vUNSPECIFIED\x10\x00\x12\v\n" +
	"\aUNKNOWN\x10\x01\x12!\n" +
	"\x1dSETTING_TYPE_IS_NOT_AVAILABLE\x10\x03\x120\n" +
	",SETTING_TYPE_IS_NOT_COMPATIBLE_WITH_CAMPAIGN\x10\x04\x12;\n" +
	"7TARGETING_SETTING_CONTAINS_INVALID_CRITERION_TYPE_GROUP\x10\x05\x12Q\n" +
	"MTARGETING_SETTING_DEMOGRAPHIC_CRITERION_TYPE_GROUPS_MUST_BE_SET_TO_TARGET_ALL\x10\x06\x12\\\n" +
	"XTARGETING_SETTING_CANNOT_CHANGE_TARGET_ALL_TO_FALSE_FOR_DEMOGRAPHIC_CRITERION_TYPE_GROUP\x10\a\x12C\n" +
	"?DYNAMIC_SEARCH_ADS_SETTING_AT_LEAST_ONE_FEED_ID_MUST_BE_PRESENT\x10\b\x12;\n" +
	"7DYNAMIC_SEARCH_ADS_SETTING_CONTAINS_INVALID_DOMAIN_NAME\x10\t\x126\n" +
	"2DYNAMIC_SEARCH_ADS_SETTING_CONTAINS_SUBDOMAIN_NAME\x10\n" +
	"\x12=\n" +
	"9DYNAMIC_SEARCH_ADS_SETTING_CONTAINS_INVALID_LANGUAGE_CODE\x10\v\x12>\n" +
	":TARGET_ALL_IS_NOT_ALLOWED_FOR_PLACEMENT_IN_SEARCH_CAMPAIGN\x10\f\x12.\n" +
	"*SETTING_VALUE_NOT_COMPATIBLE_WITH_CAMPAIGN\x10\x14\x12H\n" +
	"DBID_ONLY_IS_NOT_ALLOWED_TO_BE_MODIFIED_WITH_CUSTOMER_MATCH_TARGETING\x10\x15B\xf1\x01\n" +
	"#com.google.ads.googleads.v20.errorsB\x11SettingErrorProtoP\x01ZEgoogle.golang.org/genproto/googleapis/ads/googleads/v20/errors;errors\xa2\x02\x03GAA\xaa\x02\x1fGoogle.Ads.GoogleAds.V20.Errors\xca\x02\x1fGoogle\\Ads\\GoogleAds\\V20\\Errors\xea\x02#Google::Ads::GoogleAds::V20::Errorsb\x06proto3"

var (
	file_google_ads_googleads_v20_errors_setting_error_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v20_errors_setting_error_proto_rawDescData []byte
)

func file_google_ads_googleads_v20_errors_setting_error_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v20_errors_setting_error_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v20_errors_setting_error_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_errors_setting_error_proto_rawDesc), len(file_google_ads_googleads_v20_errors_setting_error_proto_rawDesc)))
	})
	return file_google_ads_googleads_v20_errors_setting_error_proto_rawDescData
}

var file_google_ads_googleads_v20_errors_setting_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v20_errors_setting_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v20_errors_setting_error_proto_goTypes = []any{
	(SettingErrorEnum_SettingError)(0), // 0: google.ads.googleads.v20.errors.SettingErrorEnum.SettingError
	(*SettingErrorEnum)(nil),           // 1: google.ads.googleads.v20.errors.SettingErrorEnum
}
var file_google_ads_googleads_v20_errors_setting_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v20_errors_setting_error_proto_init() }
func file_google_ads_googleads_v20_errors_setting_error_proto_init() {
	if File_google_ads_googleads_v20_errors_setting_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_errors_setting_error_proto_rawDesc), len(file_google_ads_googleads_v20_errors_setting_error_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v20_errors_setting_error_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v20_errors_setting_error_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v20_errors_setting_error_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v20_errors_setting_error_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v20_errors_setting_error_proto = out.File
	file_google_ads_googleads_v20_errors_setting_error_proto_goTypes = nil
	file_google_ads_googleads_v20_errors_setting_error_proto_depIdxs = nil
}

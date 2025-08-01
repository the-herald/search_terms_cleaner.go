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
// source: google/ads/googleads/v20/errors/user_list_error.proto

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

// Enum describing possible user list errors.
type UserListErrorEnum_UserListError int32

const (
	// Enum unspecified.
	UserListErrorEnum_UNSPECIFIED UserListErrorEnum_UserListError = 0
	// The received error code is not known in this version.
	UserListErrorEnum_UNKNOWN UserListErrorEnum_UserListError = 1
	// Creating and updating external remarketing user lists is not supported.
	UserListErrorEnum_EXTERNAL_REMARKETING_USER_LIST_MUTATE_NOT_SUPPORTED UserListErrorEnum_UserListError = 2
	// Concrete type of user list is required.
	UserListErrorEnum_CONCRETE_TYPE_REQUIRED UserListErrorEnum_UserListError = 3
	// Creating/updating user list conversion types requires specifying the
	// conversion type Id.
	UserListErrorEnum_CONVERSION_TYPE_ID_REQUIRED UserListErrorEnum_UserListError = 4
	// Remarketing user list cannot have duplicate conversion types.
	UserListErrorEnum_DUPLICATE_CONVERSION_TYPES UserListErrorEnum_UserListError = 5
	// Conversion type is invalid/unknown.
	UserListErrorEnum_INVALID_CONVERSION_TYPE UserListErrorEnum_UserListError = 6
	// User list description is empty or invalid.
	UserListErrorEnum_INVALID_DESCRIPTION UserListErrorEnum_UserListError = 7
	// User list name is empty or invalid.
	UserListErrorEnum_INVALID_NAME UserListErrorEnum_UserListError = 8
	// Type of the UserList does not match.
	UserListErrorEnum_INVALID_TYPE UserListErrorEnum_UserListError = 9
	// Embedded logical user lists are not allowed.
	UserListErrorEnum_CAN_NOT_ADD_LOGICAL_LIST_AS_LOGICAL_LIST_OPERAND UserListErrorEnum_UserListError = 10
	// User list rule operand is invalid.
	UserListErrorEnum_INVALID_USER_LIST_LOGICAL_RULE_OPERAND UserListErrorEnum_UserListError = 11
	// Name is already being used for another user list for the account.
	UserListErrorEnum_NAME_ALREADY_USED UserListErrorEnum_UserListError = 12
	// Name is required when creating a new conversion type.
	UserListErrorEnum_NEW_CONVERSION_TYPE_NAME_REQUIRED UserListErrorEnum_UserListError = 13
	// The given conversion type name has been used.
	UserListErrorEnum_CONVERSION_TYPE_NAME_ALREADY_USED UserListErrorEnum_UserListError = 14
	// Only an owner account may edit a user list.
	UserListErrorEnum_OWNERSHIP_REQUIRED_FOR_SET UserListErrorEnum_UserListError = 15
	// Creating user list without setting type in oneof user_list field, or
	// creating/updating read-only user list types is not allowed.
	UserListErrorEnum_USER_LIST_MUTATE_NOT_SUPPORTED UserListErrorEnum_UserListError = 16
	// Rule is invalid.
	UserListErrorEnum_INVALID_RULE UserListErrorEnum_UserListError = 17
	// The specified date range is empty.
	UserListErrorEnum_INVALID_DATE_RANGE UserListErrorEnum_UserListError = 27
	// A UserList which is privacy sensitive or legal rejected cannot be mutated
	// by external users.
	UserListErrorEnum_CAN_NOT_MUTATE_SENSITIVE_USERLIST UserListErrorEnum_UserListError = 28
	// Maximum number of rulebased user lists a customer can have.
	UserListErrorEnum_MAX_NUM_RULEBASED_USERLISTS UserListErrorEnum_UserListError = 29
	// BasicUserList's billable record field cannot be modified once it is set.
	UserListErrorEnum_CANNOT_MODIFY_BILLABLE_RECORD_COUNT UserListErrorEnum_UserListError = 30
	// crm_based_user_list.app_id field must be set when upload_key_type is
	// MOBILE_ADVERTISING_ID.
	UserListErrorEnum_APP_ID_NOT_SET UserListErrorEnum_UserListError = 31
	// Name of the user list is reserved for system generated lists and cannot
	// be used.
	UserListErrorEnum_USERLIST_NAME_IS_RESERVED_FOR_SYSTEM_LIST UserListErrorEnum_UserListError = 32
	// Advertiser needs to be on the allow-list to use remarketing lists created
	// from advertiser uploaded data (for example, Customer Match lists).
	UserListErrorEnum_ADVERTISER_NOT_ON_ALLOWLIST_FOR_USING_UPLOADED_DATA UserListErrorEnum_UserListError = 37
	// The provided rule_type is not supported for the user list.
	UserListErrorEnum_RULE_TYPE_IS_NOT_SUPPORTED UserListErrorEnum_UserListError = 34
	// Similar user list cannot be used as a logical user list operand.
	UserListErrorEnum_CAN_NOT_ADD_A_SIMILAR_USERLIST_AS_LOGICAL_LIST_OPERAND UserListErrorEnum_UserListError = 35
	// Logical user list should not have a mix of CRM based user list and other
	// types of lists in its rules.
	UserListErrorEnum_CAN_NOT_MIX_CRM_BASED_IN_LOGICAL_LIST_WITH_OTHER_LISTS UserListErrorEnum_UserListError = 36
	// crm_based_user_list.app_id field can only be set when upload_key_type is
	// MOBILE_ADVERTISING_ID.
	UserListErrorEnum_APP_ID_NOT_ALLOWED UserListErrorEnum_UserListError = 39
	// Google system generated user lists cannot be mutated.
	UserListErrorEnum_CANNOT_MUTATE_SYSTEM_LIST UserListErrorEnum_UserListError = 40
	// The mobile app associated with the remarketing list is sensitive.
	UserListErrorEnum_MOBILE_APP_IS_SENSITIVE UserListErrorEnum_UserListError = 41
	// One or more given seed lists do not exist.
	UserListErrorEnum_SEED_LIST_DOES_NOT_EXIST UserListErrorEnum_UserListError = 42
	// One or more given seed lists are not accessible to the current user.
	UserListErrorEnum_INVALID_SEED_LIST_ACCESS_REASON UserListErrorEnum_UserListError = 43
	// One or more given seed lists have an unsupported type.
	UserListErrorEnum_INVALID_SEED_LIST_TYPE UserListErrorEnum_UserListError = 44
	// One or more invalid country codes are added to Lookalike UserList.
	UserListErrorEnum_INVALID_COUNTRY_CODES UserListErrorEnum_UserListError = 45
)

// Enum value maps for UserListErrorEnum_UserListError.
var (
	UserListErrorEnum_UserListError_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "EXTERNAL_REMARKETING_USER_LIST_MUTATE_NOT_SUPPORTED",
		3:  "CONCRETE_TYPE_REQUIRED",
		4:  "CONVERSION_TYPE_ID_REQUIRED",
		5:  "DUPLICATE_CONVERSION_TYPES",
		6:  "INVALID_CONVERSION_TYPE",
		7:  "INVALID_DESCRIPTION",
		8:  "INVALID_NAME",
		9:  "INVALID_TYPE",
		10: "CAN_NOT_ADD_LOGICAL_LIST_AS_LOGICAL_LIST_OPERAND",
		11: "INVALID_USER_LIST_LOGICAL_RULE_OPERAND",
		12: "NAME_ALREADY_USED",
		13: "NEW_CONVERSION_TYPE_NAME_REQUIRED",
		14: "CONVERSION_TYPE_NAME_ALREADY_USED",
		15: "OWNERSHIP_REQUIRED_FOR_SET",
		16: "USER_LIST_MUTATE_NOT_SUPPORTED",
		17: "INVALID_RULE",
		27: "INVALID_DATE_RANGE",
		28: "CAN_NOT_MUTATE_SENSITIVE_USERLIST",
		29: "MAX_NUM_RULEBASED_USERLISTS",
		30: "CANNOT_MODIFY_BILLABLE_RECORD_COUNT",
		31: "APP_ID_NOT_SET",
		32: "USERLIST_NAME_IS_RESERVED_FOR_SYSTEM_LIST",
		37: "ADVERTISER_NOT_ON_ALLOWLIST_FOR_USING_UPLOADED_DATA",
		34: "RULE_TYPE_IS_NOT_SUPPORTED",
		35: "CAN_NOT_ADD_A_SIMILAR_USERLIST_AS_LOGICAL_LIST_OPERAND",
		36: "CAN_NOT_MIX_CRM_BASED_IN_LOGICAL_LIST_WITH_OTHER_LISTS",
		39: "APP_ID_NOT_ALLOWED",
		40: "CANNOT_MUTATE_SYSTEM_LIST",
		41: "MOBILE_APP_IS_SENSITIVE",
		42: "SEED_LIST_DOES_NOT_EXIST",
		43: "INVALID_SEED_LIST_ACCESS_REASON",
		44: "INVALID_SEED_LIST_TYPE",
		45: "INVALID_COUNTRY_CODES",
	}
	UserListErrorEnum_UserListError_value = map[string]int32{
		"UNSPECIFIED": 0,
		"UNKNOWN":     1,
		"EXTERNAL_REMARKETING_USER_LIST_MUTATE_NOT_SUPPORTED":    2,
		"CONCRETE_TYPE_REQUIRED":                                 3,
		"CONVERSION_TYPE_ID_REQUIRED":                            4,
		"DUPLICATE_CONVERSION_TYPES":                             5,
		"INVALID_CONVERSION_TYPE":                                6,
		"INVALID_DESCRIPTION":                                    7,
		"INVALID_NAME":                                           8,
		"INVALID_TYPE":                                           9,
		"CAN_NOT_ADD_LOGICAL_LIST_AS_LOGICAL_LIST_OPERAND":       10,
		"INVALID_USER_LIST_LOGICAL_RULE_OPERAND":                 11,
		"NAME_ALREADY_USED":                                      12,
		"NEW_CONVERSION_TYPE_NAME_REQUIRED":                      13,
		"CONVERSION_TYPE_NAME_ALREADY_USED":                      14,
		"OWNERSHIP_REQUIRED_FOR_SET":                             15,
		"USER_LIST_MUTATE_NOT_SUPPORTED":                         16,
		"INVALID_RULE":                                           17,
		"INVALID_DATE_RANGE":                                     27,
		"CAN_NOT_MUTATE_SENSITIVE_USERLIST":                      28,
		"MAX_NUM_RULEBASED_USERLISTS":                            29,
		"CANNOT_MODIFY_BILLABLE_RECORD_COUNT":                    30,
		"APP_ID_NOT_SET":                                         31,
		"USERLIST_NAME_IS_RESERVED_FOR_SYSTEM_LIST":              32,
		"ADVERTISER_NOT_ON_ALLOWLIST_FOR_USING_UPLOADED_DATA":    37,
		"RULE_TYPE_IS_NOT_SUPPORTED":                             34,
		"CAN_NOT_ADD_A_SIMILAR_USERLIST_AS_LOGICAL_LIST_OPERAND": 35,
		"CAN_NOT_MIX_CRM_BASED_IN_LOGICAL_LIST_WITH_OTHER_LISTS": 36,
		"APP_ID_NOT_ALLOWED":                                     39,
		"CANNOT_MUTATE_SYSTEM_LIST":                              40,
		"MOBILE_APP_IS_SENSITIVE":                                41,
		"SEED_LIST_DOES_NOT_EXIST":                               42,
		"INVALID_SEED_LIST_ACCESS_REASON":                        43,
		"INVALID_SEED_LIST_TYPE":                                 44,
		"INVALID_COUNTRY_CODES":                                  45,
	}
)

func (x UserListErrorEnum_UserListError) Enum() *UserListErrorEnum_UserListError {
	p := new(UserListErrorEnum_UserListError)
	*p = x
	return p
}

func (x UserListErrorEnum_UserListError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserListErrorEnum_UserListError) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v20_errors_user_list_error_proto_enumTypes[0].Descriptor()
}

func (UserListErrorEnum_UserListError) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v20_errors_user_list_error_proto_enumTypes[0]
}

func (x UserListErrorEnum_UserListError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserListErrorEnum_UserListError.Descriptor instead.
func (UserListErrorEnum_UserListError) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_errors_user_list_error_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible user list errors.
type UserListErrorEnum struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserListErrorEnum) Reset() {
	*x = UserListErrorEnum{}
	mi := &file_google_ads_googleads_v20_errors_user_list_error_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserListErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserListErrorEnum) ProtoMessage() {}

func (x *UserListErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_errors_user_list_error_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserListErrorEnum.ProtoReflect.Descriptor instead.
func (*UserListErrorEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_errors_user_list_error_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v20_errors_user_list_error_proto protoreflect.FileDescriptor

const file_google_ads_googleads_v20_errors_user_list_error_proto_rawDesc = "" +
	"\n" +
	"5google/ads/googleads/v20/errors/user_list_error.proto\x12\x1fgoogle.ads.googleads.v20.errors\"\xbb\t\n" +
	"\x11UserListErrorEnum\"\xa5\t\n" +
	"\rUserListError\x12\x0f\n" +
	"\vUNSPECIFIED\x10\x00\x12\v\n" +
	"\aUNKNOWN\x10\x01\x127\n" +
	"3EXTERNAL_REMARKETING_USER_LIST_MUTATE_NOT_SUPPORTED\x10\x02\x12\x1a\n" +
	"\x16CONCRETE_TYPE_REQUIRED\x10\x03\x12\x1f\n" +
	"\x1bCONVERSION_TYPE_ID_REQUIRED\x10\x04\x12\x1e\n" +
	"\x1aDUPLICATE_CONVERSION_TYPES\x10\x05\x12\x1b\n" +
	"\x17INVALID_CONVERSION_TYPE\x10\x06\x12\x17\n" +
	"\x13INVALID_DESCRIPTION\x10\a\x12\x10\n" +
	"\fINVALID_NAME\x10\b\x12\x10\n" +
	"\fINVALID_TYPE\x10\t\x124\n" +
	"0CAN_NOT_ADD_LOGICAL_LIST_AS_LOGICAL_LIST_OPERAND\x10\n" +
	"\x12*\n" +
	"&INVALID_USER_LIST_LOGICAL_RULE_OPERAND\x10\v\x12\x15\n" +
	"\x11NAME_ALREADY_USED\x10\f\x12%\n" +
	"!NEW_CONVERSION_TYPE_NAME_REQUIRED\x10\r\x12%\n" +
	"!CONVERSION_TYPE_NAME_ALREADY_USED\x10\x0e\x12\x1e\n" +
	"\x1aOWNERSHIP_REQUIRED_FOR_SET\x10\x0f\x12\"\n" +
	"\x1eUSER_LIST_MUTATE_NOT_SUPPORTED\x10\x10\x12\x10\n" +
	"\fINVALID_RULE\x10\x11\x12\x16\n" +
	"\x12INVALID_DATE_RANGE\x10\x1b\x12%\n" +
	"!CAN_NOT_MUTATE_SENSITIVE_USERLIST\x10\x1c\x12\x1f\n" +
	"\x1bMAX_NUM_RULEBASED_USERLISTS\x10\x1d\x12'\n" +
	"#CANNOT_MODIFY_BILLABLE_RECORD_COUNT\x10\x1e\x12\x12\n" +
	"\x0eAPP_ID_NOT_SET\x10\x1f\x12-\n" +
	")USERLIST_NAME_IS_RESERVED_FOR_SYSTEM_LIST\x10 \x127\n" +
	"3ADVERTISER_NOT_ON_ALLOWLIST_FOR_USING_UPLOADED_DATA\x10%\x12\x1e\n" +
	"\x1aRULE_TYPE_IS_NOT_SUPPORTED\x10\"\x12:\n" +
	"6CAN_NOT_ADD_A_SIMILAR_USERLIST_AS_LOGICAL_LIST_OPERAND\x10#\x12:\n" +
	"6CAN_NOT_MIX_CRM_BASED_IN_LOGICAL_LIST_WITH_OTHER_LISTS\x10$\x12\x16\n" +
	"\x12APP_ID_NOT_ALLOWED\x10'\x12\x1d\n" +
	"\x19CANNOT_MUTATE_SYSTEM_LIST\x10(\x12\x1b\n" +
	"\x17MOBILE_APP_IS_SENSITIVE\x10)\x12\x1c\n" +
	"\x18SEED_LIST_DOES_NOT_EXIST\x10*\x12#\n" +
	"\x1fINVALID_SEED_LIST_ACCESS_REASON\x10+\x12\x1a\n" +
	"\x16INVALID_SEED_LIST_TYPE\x10,\x12\x19\n" +
	"\x15INVALID_COUNTRY_CODES\x10-B\xf2\x01\n" +
	"#com.google.ads.googleads.v20.errorsB\x12UserListErrorProtoP\x01ZEgoogle.golang.org/genproto/googleapis/ads/googleads/v20/errors;errors\xa2\x02\x03GAA\xaa\x02\x1fGoogle.Ads.GoogleAds.V20.Errors\xca\x02\x1fGoogle\\Ads\\GoogleAds\\V20\\Errors\xea\x02#Google::Ads::GoogleAds::V20::Errorsb\x06proto3"

var (
	file_google_ads_googleads_v20_errors_user_list_error_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v20_errors_user_list_error_proto_rawDescData []byte
)

func file_google_ads_googleads_v20_errors_user_list_error_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v20_errors_user_list_error_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v20_errors_user_list_error_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_errors_user_list_error_proto_rawDesc), len(file_google_ads_googleads_v20_errors_user_list_error_proto_rawDesc)))
	})
	return file_google_ads_googleads_v20_errors_user_list_error_proto_rawDescData
}

var file_google_ads_googleads_v20_errors_user_list_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v20_errors_user_list_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v20_errors_user_list_error_proto_goTypes = []any{
	(UserListErrorEnum_UserListError)(0), // 0: google.ads.googleads.v20.errors.UserListErrorEnum.UserListError
	(*UserListErrorEnum)(nil),            // 1: google.ads.googleads.v20.errors.UserListErrorEnum
}
var file_google_ads_googleads_v20_errors_user_list_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v20_errors_user_list_error_proto_init() }
func file_google_ads_googleads_v20_errors_user_list_error_proto_init() {
	if File_google_ads_googleads_v20_errors_user_list_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_errors_user_list_error_proto_rawDesc), len(file_google_ads_googleads_v20_errors_user_list_error_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v20_errors_user_list_error_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v20_errors_user_list_error_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v20_errors_user_list_error_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v20_errors_user_list_error_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v20_errors_user_list_error_proto = out.File
	file_google_ads_googleads_v20_errors_user_list_error_proto_goTypes = nil
	file_google_ads_googleads_v20_errors_user_list_error_proto_depIdxs = nil
}

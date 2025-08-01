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
// source: google/ads/googleads/v20/enums/user_list_size_range.proto

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

// Enum containing possible user list size ranges.
type UserListSizeRangeEnum_UserListSizeRange int32

const (
	// Not specified.
	UserListSizeRangeEnum_UNSPECIFIED UserListSizeRangeEnum_UserListSizeRange = 0
	// Used for return value only. Represents value unknown in this version.
	UserListSizeRangeEnum_UNKNOWN UserListSizeRangeEnum_UserListSizeRange = 1
	// User list has less than 500 users.
	UserListSizeRangeEnum_LESS_THAN_FIVE_HUNDRED UserListSizeRangeEnum_UserListSizeRange = 2
	// User list has number of users in range of 500 to 1000.
	UserListSizeRangeEnum_LESS_THAN_ONE_THOUSAND UserListSizeRangeEnum_UserListSizeRange = 3
	// User list has number of users in range of 1000 to 10000.
	UserListSizeRangeEnum_ONE_THOUSAND_TO_TEN_THOUSAND UserListSizeRangeEnum_UserListSizeRange = 4
	// User list has number of users in range of 10000 to 50000.
	UserListSizeRangeEnum_TEN_THOUSAND_TO_FIFTY_THOUSAND UserListSizeRangeEnum_UserListSizeRange = 5
	// User list has number of users in range of 50000 to 100000.
	UserListSizeRangeEnum_FIFTY_THOUSAND_TO_ONE_HUNDRED_THOUSAND UserListSizeRangeEnum_UserListSizeRange = 6
	// User list has number of users in range of 100000 to 300000.
	UserListSizeRangeEnum_ONE_HUNDRED_THOUSAND_TO_THREE_HUNDRED_THOUSAND UserListSizeRangeEnum_UserListSizeRange = 7
	// User list has number of users in range of 300000 to 500000.
	UserListSizeRangeEnum_THREE_HUNDRED_THOUSAND_TO_FIVE_HUNDRED_THOUSAND UserListSizeRangeEnum_UserListSizeRange = 8
	// User list has number of users in range of 500000 to 1 million.
	UserListSizeRangeEnum_FIVE_HUNDRED_THOUSAND_TO_ONE_MILLION UserListSizeRangeEnum_UserListSizeRange = 9
	// User list has number of users in range of 1 to 2 millions.
	UserListSizeRangeEnum_ONE_MILLION_TO_TWO_MILLION UserListSizeRangeEnum_UserListSizeRange = 10
	// User list has number of users in range of 2 to 3 millions.
	UserListSizeRangeEnum_TWO_MILLION_TO_THREE_MILLION UserListSizeRangeEnum_UserListSizeRange = 11
	// User list has number of users in range of 3 to 5 millions.
	UserListSizeRangeEnum_THREE_MILLION_TO_FIVE_MILLION UserListSizeRangeEnum_UserListSizeRange = 12
	// User list has number of users in range of 5 to 10 millions.
	UserListSizeRangeEnum_FIVE_MILLION_TO_TEN_MILLION UserListSizeRangeEnum_UserListSizeRange = 13
	// User list has number of users in range of 10 to 20 millions.
	UserListSizeRangeEnum_TEN_MILLION_TO_TWENTY_MILLION UserListSizeRangeEnum_UserListSizeRange = 14
	// User list has number of users in range of 20 to 30 millions.
	UserListSizeRangeEnum_TWENTY_MILLION_TO_THIRTY_MILLION UserListSizeRangeEnum_UserListSizeRange = 15
	// User list has number of users in range of 30 to 50 millions.
	UserListSizeRangeEnum_THIRTY_MILLION_TO_FIFTY_MILLION UserListSizeRangeEnum_UserListSizeRange = 16
	// User list has over 50 million users.
	UserListSizeRangeEnum_OVER_FIFTY_MILLION UserListSizeRangeEnum_UserListSizeRange = 17
)

// Enum value maps for UserListSizeRangeEnum_UserListSizeRange.
var (
	UserListSizeRangeEnum_UserListSizeRange_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "LESS_THAN_FIVE_HUNDRED",
		3:  "LESS_THAN_ONE_THOUSAND",
		4:  "ONE_THOUSAND_TO_TEN_THOUSAND",
		5:  "TEN_THOUSAND_TO_FIFTY_THOUSAND",
		6:  "FIFTY_THOUSAND_TO_ONE_HUNDRED_THOUSAND",
		7:  "ONE_HUNDRED_THOUSAND_TO_THREE_HUNDRED_THOUSAND",
		8:  "THREE_HUNDRED_THOUSAND_TO_FIVE_HUNDRED_THOUSAND",
		9:  "FIVE_HUNDRED_THOUSAND_TO_ONE_MILLION",
		10: "ONE_MILLION_TO_TWO_MILLION",
		11: "TWO_MILLION_TO_THREE_MILLION",
		12: "THREE_MILLION_TO_FIVE_MILLION",
		13: "FIVE_MILLION_TO_TEN_MILLION",
		14: "TEN_MILLION_TO_TWENTY_MILLION",
		15: "TWENTY_MILLION_TO_THIRTY_MILLION",
		16: "THIRTY_MILLION_TO_FIFTY_MILLION",
		17: "OVER_FIFTY_MILLION",
	}
	UserListSizeRangeEnum_UserListSizeRange_value = map[string]int32{
		"UNSPECIFIED":                                     0,
		"UNKNOWN":                                         1,
		"LESS_THAN_FIVE_HUNDRED":                          2,
		"LESS_THAN_ONE_THOUSAND":                          3,
		"ONE_THOUSAND_TO_TEN_THOUSAND":                    4,
		"TEN_THOUSAND_TO_FIFTY_THOUSAND":                  5,
		"FIFTY_THOUSAND_TO_ONE_HUNDRED_THOUSAND":          6,
		"ONE_HUNDRED_THOUSAND_TO_THREE_HUNDRED_THOUSAND":  7,
		"THREE_HUNDRED_THOUSAND_TO_FIVE_HUNDRED_THOUSAND": 8,
		"FIVE_HUNDRED_THOUSAND_TO_ONE_MILLION":            9,
		"ONE_MILLION_TO_TWO_MILLION":                      10,
		"TWO_MILLION_TO_THREE_MILLION":                    11,
		"THREE_MILLION_TO_FIVE_MILLION":                   12,
		"FIVE_MILLION_TO_TEN_MILLION":                     13,
		"TEN_MILLION_TO_TWENTY_MILLION":                   14,
		"TWENTY_MILLION_TO_THIRTY_MILLION":                15,
		"THIRTY_MILLION_TO_FIFTY_MILLION":                 16,
		"OVER_FIFTY_MILLION":                              17,
	}
)

func (x UserListSizeRangeEnum_UserListSizeRange) Enum() *UserListSizeRangeEnum_UserListSizeRange {
	p := new(UserListSizeRangeEnum_UserListSizeRange)
	*p = x
	return p
}

func (x UserListSizeRangeEnum_UserListSizeRange) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserListSizeRangeEnum_UserListSizeRange) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v20_enums_user_list_size_range_proto_enumTypes[0].Descriptor()
}

func (UserListSizeRangeEnum_UserListSizeRange) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v20_enums_user_list_size_range_proto_enumTypes[0]
}

func (x UserListSizeRangeEnum_UserListSizeRange) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserListSizeRangeEnum_UserListSizeRange.Descriptor instead.
func (UserListSizeRangeEnum_UserListSizeRange) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_enums_user_list_size_range_proto_rawDescGZIP(), []int{0, 0}
}

// Size range in terms of number of users of a UserList.
type UserListSizeRangeEnum struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserListSizeRangeEnum) Reset() {
	*x = UserListSizeRangeEnum{}
	mi := &file_google_ads_googleads_v20_enums_user_list_size_range_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserListSizeRangeEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserListSizeRangeEnum) ProtoMessage() {}

func (x *UserListSizeRangeEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_enums_user_list_size_range_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserListSizeRangeEnum.ProtoReflect.Descriptor instead.
func (*UserListSizeRangeEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_enums_user_list_size_range_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v20_enums_user_list_size_range_proto protoreflect.FileDescriptor

const file_google_ads_googleads_v20_enums_user_list_size_range_proto_rawDesc = "" +
	"\n" +
	"9google/ads/googleads/v20/enums/user_list_size_range.proto\x12\x1egoogle.ads.googleads.v20.enums\"\x94\x05\n" +
	"\x15UserListSizeRangeEnum\"\xfa\x04\n" +
	"\x11UserListSizeRange\x12\x0f\n" +
	"\vUNSPECIFIED\x10\x00\x12\v\n" +
	"\aUNKNOWN\x10\x01\x12\x1a\n" +
	"\x16LESS_THAN_FIVE_HUNDRED\x10\x02\x12\x1a\n" +
	"\x16LESS_THAN_ONE_THOUSAND\x10\x03\x12 \n" +
	"\x1cONE_THOUSAND_TO_TEN_THOUSAND\x10\x04\x12\"\n" +
	"\x1eTEN_THOUSAND_TO_FIFTY_THOUSAND\x10\x05\x12*\n" +
	"&FIFTY_THOUSAND_TO_ONE_HUNDRED_THOUSAND\x10\x06\x122\n" +
	".ONE_HUNDRED_THOUSAND_TO_THREE_HUNDRED_THOUSAND\x10\a\x123\n" +
	"/THREE_HUNDRED_THOUSAND_TO_FIVE_HUNDRED_THOUSAND\x10\b\x12(\n" +
	"$FIVE_HUNDRED_THOUSAND_TO_ONE_MILLION\x10\t\x12\x1e\n" +
	"\x1aONE_MILLION_TO_TWO_MILLION\x10\n" +
	"\x12 \n" +
	"\x1cTWO_MILLION_TO_THREE_MILLION\x10\v\x12!\n" +
	"\x1dTHREE_MILLION_TO_FIVE_MILLION\x10\f\x12\x1f\n" +
	"\x1bFIVE_MILLION_TO_TEN_MILLION\x10\r\x12!\n" +
	"\x1dTEN_MILLION_TO_TWENTY_MILLION\x10\x0e\x12$\n" +
	" TWENTY_MILLION_TO_THIRTY_MILLION\x10\x0f\x12#\n" +
	"\x1fTHIRTY_MILLION_TO_FIFTY_MILLION\x10\x10\x12\x16\n" +
	"\x12OVER_FIFTY_MILLION\x10\x11B\xf0\x01\n" +
	"\"com.google.ads.googleads.v20.enumsB\x16UserListSizeRangeProtoP\x01ZCgoogle.golang.org/genproto/googleapis/ads/googleads/v20/enums;enums\xa2\x02\x03GAA\xaa\x02\x1eGoogle.Ads.GoogleAds.V20.Enums\xca\x02\x1eGoogle\\Ads\\GoogleAds\\V20\\Enums\xea\x02\"Google::Ads::GoogleAds::V20::Enumsb\x06proto3"

var (
	file_google_ads_googleads_v20_enums_user_list_size_range_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v20_enums_user_list_size_range_proto_rawDescData []byte
)

func file_google_ads_googleads_v20_enums_user_list_size_range_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v20_enums_user_list_size_range_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v20_enums_user_list_size_range_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_enums_user_list_size_range_proto_rawDesc), len(file_google_ads_googleads_v20_enums_user_list_size_range_proto_rawDesc)))
	})
	return file_google_ads_googleads_v20_enums_user_list_size_range_proto_rawDescData
}

var file_google_ads_googleads_v20_enums_user_list_size_range_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v20_enums_user_list_size_range_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v20_enums_user_list_size_range_proto_goTypes = []any{
	(UserListSizeRangeEnum_UserListSizeRange)(0), // 0: google.ads.googleads.v20.enums.UserListSizeRangeEnum.UserListSizeRange
	(*UserListSizeRangeEnum)(nil),                // 1: google.ads.googleads.v20.enums.UserListSizeRangeEnum
}
var file_google_ads_googleads_v20_enums_user_list_size_range_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v20_enums_user_list_size_range_proto_init() }
func file_google_ads_googleads_v20_enums_user_list_size_range_proto_init() {
	if File_google_ads_googleads_v20_enums_user_list_size_range_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_enums_user_list_size_range_proto_rawDesc), len(file_google_ads_googleads_v20_enums_user_list_size_range_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v20_enums_user_list_size_range_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v20_enums_user_list_size_range_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v20_enums_user_list_size_range_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v20_enums_user_list_size_range_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v20_enums_user_list_size_range_proto = out.File
	file_google_ads_googleads_v20_enums_user_list_size_range_proto_goTypes = nil
	file_google_ads_googleads_v20_enums_user_list_size_range_proto_depIdxs = nil
}

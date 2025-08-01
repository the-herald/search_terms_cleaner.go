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
// source: google/ads/googleads/v20/errors/feed_item_error.proto

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

// Enum describing possible feed item errors.
type FeedItemErrorEnum_FeedItemError int32

const (
	// Enum unspecified.
	FeedItemErrorEnum_UNSPECIFIED FeedItemErrorEnum_FeedItemError = 0
	// The received error code is not known in this version.
	FeedItemErrorEnum_UNKNOWN FeedItemErrorEnum_FeedItemError = 1
	// Cannot convert the feed attribute value from string to its real type.
	FeedItemErrorEnum_CANNOT_CONVERT_ATTRIBUTE_VALUE_FROM_STRING FeedItemErrorEnum_FeedItemError = 2
	// Cannot operate on removed feed item.
	FeedItemErrorEnum_CANNOT_OPERATE_ON_REMOVED_FEED_ITEM FeedItemErrorEnum_FeedItemError = 3
	// Date time zone does not match the account's time zone.
	FeedItemErrorEnum_DATE_TIME_MUST_BE_IN_ACCOUNT_TIME_ZONE FeedItemErrorEnum_FeedItemError = 4
	// Feed item with the key attributes could not be found.
	FeedItemErrorEnum_KEY_ATTRIBUTES_NOT_FOUND FeedItemErrorEnum_FeedItemError = 5
	// Url feed attribute value is not valid.
	FeedItemErrorEnum_INVALID_URL FeedItemErrorEnum_FeedItemError = 6
	// Some key attributes are missing.
	FeedItemErrorEnum_MISSING_KEY_ATTRIBUTES FeedItemErrorEnum_FeedItemError = 7
	// Feed item has same key attributes as another feed item.
	FeedItemErrorEnum_KEY_ATTRIBUTES_NOT_UNIQUE FeedItemErrorEnum_FeedItemError = 8
	// Cannot modify key attributes on an existing feed item.
	FeedItemErrorEnum_CANNOT_MODIFY_KEY_ATTRIBUTE_VALUE FeedItemErrorEnum_FeedItemError = 9
	// The feed attribute value is too large.
	FeedItemErrorEnum_SIZE_TOO_LARGE_FOR_MULTI_VALUE_ATTRIBUTE FeedItemErrorEnum_FeedItemError = 10
	// Feed is read only.
	FeedItemErrorEnum_LEGACY_FEED_TYPE_READ_ONLY FeedItemErrorEnum_FeedItemError = 11
)

// Enum value maps for FeedItemErrorEnum_FeedItemError.
var (
	FeedItemErrorEnum_FeedItemError_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "CANNOT_CONVERT_ATTRIBUTE_VALUE_FROM_STRING",
		3:  "CANNOT_OPERATE_ON_REMOVED_FEED_ITEM",
		4:  "DATE_TIME_MUST_BE_IN_ACCOUNT_TIME_ZONE",
		5:  "KEY_ATTRIBUTES_NOT_FOUND",
		6:  "INVALID_URL",
		7:  "MISSING_KEY_ATTRIBUTES",
		8:  "KEY_ATTRIBUTES_NOT_UNIQUE",
		9:  "CANNOT_MODIFY_KEY_ATTRIBUTE_VALUE",
		10: "SIZE_TOO_LARGE_FOR_MULTI_VALUE_ATTRIBUTE",
		11: "LEGACY_FEED_TYPE_READ_ONLY",
	}
	FeedItemErrorEnum_FeedItemError_value = map[string]int32{
		"UNSPECIFIED": 0,
		"UNKNOWN":     1,
		"CANNOT_CONVERT_ATTRIBUTE_VALUE_FROM_STRING": 2,
		"CANNOT_OPERATE_ON_REMOVED_FEED_ITEM":        3,
		"DATE_TIME_MUST_BE_IN_ACCOUNT_TIME_ZONE":     4,
		"KEY_ATTRIBUTES_NOT_FOUND":                   5,
		"INVALID_URL":                                6,
		"MISSING_KEY_ATTRIBUTES":                     7,
		"KEY_ATTRIBUTES_NOT_UNIQUE":                  8,
		"CANNOT_MODIFY_KEY_ATTRIBUTE_VALUE":          9,
		"SIZE_TOO_LARGE_FOR_MULTI_VALUE_ATTRIBUTE":   10,
		"LEGACY_FEED_TYPE_READ_ONLY":                 11,
	}
)

func (x FeedItemErrorEnum_FeedItemError) Enum() *FeedItemErrorEnum_FeedItemError {
	p := new(FeedItemErrorEnum_FeedItemError)
	*p = x
	return p
}

func (x FeedItemErrorEnum_FeedItemError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FeedItemErrorEnum_FeedItemError) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v20_errors_feed_item_error_proto_enumTypes[0].Descriptor()
}

func (FeedItemErrorEnum_FeedItemError) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v20_errors_feed_item_error_proto_enumTypes[0]
}

func (x FeedItemErrorEnum_FeedItemError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FeedItemErrorEnum_FeedItemError.Descriptor instead.
func (FeedItemErrorEnum_FeedItemError) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_errors_feed_item_error_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible feed item errors.
type FeedItemErrorEnum struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FeedItemErrorEnum) Reset() {
	*x = FeedItemErrorEnum{}
	mi := &file_google_ads_googleads_v20_errors_feed_item_error_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FeedItemErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedItemErrorEnum) ProtoMessage() {}

func (x *FeedItemErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_errors_feed_item_error_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeedItemErrorEnum.ProtoReflect.Descriptor instead.
func (*FeedItemErrorEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_errors_feed_item_error_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v20_errors_feed_item_error_proto protoreflect.FileDescriptor

const file_google_ads_googleads_v20_errors_feed_item_error_proto_rawDesc = "" +
	"\n" +
	"5google/ads/googleads/v20/errors/feed_item_error.proto\x12\x1fgoogle.ads.googleads.v20.errors\"\xa7\x03\n" +
	"\x11FeedItemErrorEnum\"\x91\x03\n" +
	"\rFeedItemError\x12\x0f\n" +
	"\vUNSPECIFIED\x10\x00\x12\v\n" +
	"\aUNKNOWN\x10\x01\x12.\n" +
	"*CANNOT_CONVERT_ATTRIBUTE_VALUE_FROM_STRING\x10\x02\x12'\n" +
	"#CANNOT_OPERATE_ON_REMOVED_FEED_ITEM\x10\x03\x12*\n" +
	"&DATE_TIME_MUST_BE_IN_ACCOUNT_TIME_ZONE\x10\x04\x12\x1c\n" +
	"\x18KEY_ATTRIBUTES_NOT_FOUND\x10\x05\x12\x0f\n" +
	"\vINVALID_URL\x10\x06\x12\x1a\n" +
	"\x16MISSING_KEY_ATTRIBUTES\x10\a\x12\x1d\n" +
	"\x19KEY_ATTRIBUTES_NOT_UNIQUE\x10\b\x12%\n" +
	"!CANNOT_MODIFY_KEY_ATTRIBUTE_VALUE\x10\t\x12,\n" +
	"(SIZE_TOO_LARGE_FOR_MULTI_VALUE_ATTRIBUTE\x10\n" +
	"\x12\x1e\n" +
	"\x1aLEGACY_FEED_TYPE_READ_ONLY\x10\vB\xf2\x01\n" +
	"#com.google.ads.googleads.v20.errorsB\x12FeedItemErrorProtoP\x01ZEgoogle.golang.org/genproto/googleapis/ads/googleads/v20/errors;errors\xa2\x02\x03GAA\xaa\x02\x1fGoogle.Ads.GoogleAds.V20.Errors\xca\x02\x1fGoogle\\Ads\\GoogleAds\\V20\\Errors\xea\x02#Google::Ads::GoogleAds::V20::Errorsb\x06proto3"

var (
	file_google_ads_googleads_v20_errors_feed_item_error_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v20_errors_feed_item_error_proto_rawDescData []byte
)

func file_google_ads_googleads_v20_errors_feed_item_error_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v20_errors_feed_item_error_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v20_errors_feed_item_error_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_errors_feed_item_error_proto_rawDesc), len(file_google_ads_googleads_v20_errors_feed_item_error_proto_rawDesc)))
	})
	return file_google_ads_googleads_v20_errors_feed_item_error_proto_rawDescData
}

var file_google_ads_googleads_v20_errors_feed_item_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v20_errors_feed_item_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v20_errors_feed_item_error_proto_goTypes = []any{
	(FeedItemErrorEnum_FeedItemError)(0), // 0: google.ads.googleads.v20.errors.FeedItemErrorEnum.FeedItemError
	(*FeedItemErrorEnum)(nil),            // 1: google.ads.googleads.v20.errors.FeedItemErrorEnum
}
var file_google_ads_googleads_v20_errors_feed_item_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v20_errors_feed_item_error_proto_init() }
func file_google_ads_googleads_v20_errors_feed_item_error_proto_init() {
	if File_google_ads_googleads_v20_errors_feed_item_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_errors_feed_item_error_proto_rawDesc), len(file_google_ads_googleads_v20_errors_feed_item_error_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v20_errors_feed_item_error_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v20_errors_feed_item_error_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v20_errors_feed_item_error_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v20_errors_feed_item_error_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v20_errors_feed_item_error_proto = out.File
	file_google_ads_googleads_v20_errors_feed_item_error_proto_goTypes = nil
	file_google_ads_googleads_v20_errors_feed_item_error_proto_depIdxs = nil
}

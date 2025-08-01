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
// source: google/ads/googleads/v20/errors/asset_set_error.proto

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

// Enum describing possible asset set errors.
type AssetSetErrorEnum_AssetSetError int32

const (
	// Enum unspecified.
	AssetSetErrorEnum_UNSPECIFIED AssetSetErrorEnum_AssetSetError = 0
	// The received error code is not known in this version.
	AssetSetErrorEnum_UNKNOWN AssetSetErrorEnum_AssetSetError = 1
	// The asset set name matches that of another enabled asset set.
	AssetSetErrorEnum_DUPLICATE_ASSET_SET_NAME AssetSetErrorEnum_AssetSetError = 2
	// The type of AssetSet.asset_set_source does not match the type of
	// AssetSet.location_set.source in its parent AssetSet.
	AssetSetErrorEnum_INVALID_PARENT_ASSET_SET_TYPE AssetSetErrorEnum_AssetSetError = 3
	// The asset set source doesn't match its parent AssetSet's data.
	AssetSetErrorEnum_ASSET_SET_SOURCE_INCOMPATIBLE_WITH_PARENT_ASSET_SET AssetSetErrorEnum_AssetSetError = 4
	// This AssetSet type cannot be linked to CustomerAssetSet.
	AssetSetErrorEnum_ASSET_SET_TYPE_CANNOT_BE_LINKED_TO_CUSTOMER AssetSetErrorEnum_AssetSetError = 5
	// The chain id(s) in ChainSet of a LOCATION_SYNC typed AssetSet is invalid.
	AssetSetErrorEnum_INVALID_CHAIN_IDS AssetSetErrorEnum_AssetSetError = 6
	// The relationship type in ChainSet of a LOCATION_SYNC typed AssetSet is
	// not supported.
	AssetSetErrorEnum_LOCATION_SYNC_ASSET_SET_DOES_NOT_SUPPORT_RELATIONSHIP_TYPE AssetSetErrorEnum_AssetSetError = 7
	// There is more than one enabled LocationSync typed AssetSet under one
	// customer.
	AssetSetErrorEnum_NOT_UNIQUE_ENABLED_LOCATION_SYNC_TYPED_ASSET_SET AssetSetErrorEnum_AssetSetError = 8
	// The place id(s) in a LocationSync typed AssetSet is invalid and can't be
	// decoded.
	AssetSetErrorEnum_INVALID_PLACE_IDS AssetSetErrorEnum_AssetSetError = 9
	// The Google Business Profile OAuth info is invalid.
	AssetSetErrorEnum_OAUTH_INFO_INVALID AssetSetErrorEnum_AssetSetError = 11
	// The Google Business Profile OAuth info is missing.
	AssetSetErrorEnum_OAUTH_INFO_MISSING AssetSetErrorEnum_AssetSetError = 12
	// Can't delete an AssetSet if it has any enabled linkages (e.g.
	// CustomerAssetSet), or AssetSet is a parent AssetSet and has enabled child
	// AssetSet associated.
	AssetSetErrorEnum_CANNOT_DELETE_AS_ENABLED_LINKAGES_EXIST AssetSetErrorEnum_AssetSetError = 10
)

// Enum value maps for AssetSetErrorEnum_AssetSetError.
var (
	AssetSetErrorEnum_AssetSetError_name = map[int32]string{
		0:  "UNSPECIFIED",
		1:  "UNKNOWN",
		2:  "DUPLICATE_ASSET_SET_NAME",
		3:  "INVALID_PARENT_ASSET_SET_TYPE",
		4:  "ASSET_SET_SOURCE_INCOMPATIBLE_WITH_PARENT_ASSET_SET",
		5:  "ASSET_SET_TYPE_CANNOT_BE_LINKED_TO_CUSTOMER",
		6:  "INVALID_CHAIN_IDS",
		7:  "LOCATION_SYNC_ASSET_SET_DOES_NOT_SUPPORT_RELATIONSHIP_TYPE",
		8:  "NOT_UNIQUE_ENABLED_LOCATION_SYNC_TYPED_ASSET_SET",
		9:  "INVALID_PLACE_IDS",
		11: "OAUTH_INFO_INVALID",
		12: "OAUTH_INFO_MISSING",
		10: "CANNOT_DELETE_AS_ENABLED_LINKAGES_EXIST",
	}
	AssetSetErrorEnum_AssetSetError_value = map[string]int32{
		"UNSPECIFIED":                   0,
		"UNKNOWN":                       1,
		"DUPLICATE_ASSET_SET_NAME":      2,
		"INVALID_PARENT_ASSET_SET_TYPE": 3,
		"ASSET_SET_SOURCE_INCOMPATIBLE_WITH_PARENT_ASSET_SET": 4,
		"ASSET_SET_TYPE_CANNOT_BE_LINKED_TO_CUSTOMER":         5,
		"INVALID_CHAIN_IDS": 6,
		"LOCATION_SYNC_ASSET_SET_DOES_NOT_SUPPORT_RELATIONSHIP_TYPE": 7,
		"NOT_UNIQUE_ENABLED_LOCATION_SYNC_TYPED_ASSET_SET":           8,
		"INVALID_PLACE_IDS":                       9,
		"OAUTH_INFO_INVALID":                      11,
		"OAUTH_INFO_MISSING":                      12,
		"CANNOT_DELETE_AS_ENABLED_LINKAGES_EXIST": 10,
	}
)

func (x AssetSetErrorEnum_AssetSetError) Enum() *AssetSetErrorEnum_AssetSetError {
	p := new(AssetSetErrorEnum_AssetSetError)
	*p = x
	return p
}

func (x AssetSetErrorEnum_AssetSetError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AssetSetErrorEnum_AssetSetError) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v20_errors_asset_set_error_proto_enumTypes[0].Descriptor()
}

func (AssetSetErrorEnum_AssetSetError) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v20_errors_asset_set_error_proto_enumTypes[0]
}

func (x AssetSetErrorEnum_AssetSetError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AssetSetErrorEnum_AssetSetError.Descriptor instead.
func (AssetSetErrorEnum_AssetSetError) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_errors_asset_set_error_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible asset set errors.
type AssetSetErrorEnum struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AssetSetErrorEnum) Reset() {
	*x = AssetSetErrorEnum{}
	mi := &file_google_ads_googleads_v20_errors_asset_set_error_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AssetSetErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssetSetErrorEnum) ProtoMessage() {}

func (x *AssetSetErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_errors_asset_set_error_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssetSetErrorEnum.ProtoReflect.Descriptor instead.
func (*AssetSetErrorEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_errors_asset_set_error_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v20_errors_asset_set_error_proto protoreflect.FileDescriptor

const file_google_ads_googleads_v20_errors_asset_set_error_proto_rawDesc = "" +
	"\n" +
	"5google/ads/googleads/v20/errors/asset_set_error.proto\x12\x1fgoogle.ads.googleads.v20.errors\"\xef\x03\n" +
	"\x11AssetSetErrorEnum\"\xd9\x03\n" +
	"\rAssetSetError\x12\x0f\n" +
	"\vUNSPECIFIED\x10\x00\x12\v\n" +
	"\aUNKNOWN\x10\x01\x12\x1c\n" +
	"\x18DUPLICATE_ASSET_SET_NAME\x10\x02\x12!\n" +
	"\x1dINVALID_PARENT_ASSET_SET_TYPE\x10\x03\x127\n" +
	"3ASSET_SET_SOURCE_INCOMPATIBLE_WITH_PARENT_ASSET_SET\x10\x04\x12/\n" +
	"+ASSET_SET_TYPE_CANNOT_BE_LINKED_TO_CUSTOMER\x10\x05\x12\x15\n" +
	"\x11INVALID_CHAIN_IDS\x10\x06\x12>\n" +
	":LOCATION_SYNC_ASSET_SET_DOES_NOT_SUPPORT_RELATIONSHIP_TYPE\x10\a\x124\n" +
	"0NOT_UNIQUE_ENABLED_LOCATION_SYNC_TYPED_ASSET_SET\x10\b\x12\x15\n" +
	"\x11INVALID_PLACE_IDS\x10\t\x12\x16\n" +
	"\x12OAUTH_INFO_INVALID\x10\v\x12\x16\n" +
	"\x12OAUTH_INFO_MISSING\x10\f\x12+\n" +
	"'CANNOT_DELETE_AS_ENABLED_LINKAGES_EXIST\x10\n" +
	"B\xf2\x01\n" +
	"#com.google.ads.googleads.v20.errorsB\x12AssetSetErrorProtoP\x01ZEgoogle.golang.org/genproto/googleapis/ads/googleads/v20/errors;errors\xa2\x02\x03GAA\xaa\x02\x1fGoogle.Ads.GoogleAds.V20.Errors\xca\x02\x1fGoogle\\Ads\\GoogleAds\\V20\\Errors\xea\x02#Google::Ads::GoogleAds::V20::Errorsb\x06proto3"

var (
	file_google_ads_googleads_v20_errors_asset_set_error_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v20_errors_asset_set_error_proto_rawDescData []byte
)

func file_google_ads_googleads_v20_errors_asset_set_error_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v20_errors_asset_set_error_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v20_errors_asset_set_error_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_errors_asset_set_error_proto_rawDesc), len(file_google_ads_googleads_v20_errors_asset_set_error_proto_rawDesc)))
	})
	return file_google_ads_googleads_v20_errors_asset_set_error_proto_rawDescData
}

var file_google_ads_googleads_v20_errors_asset_set_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v20_errors_asset_set_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v20_errors_asset_set_error_proto_goTypes = []any{
	(AssetSetErrorEnum_AssetSetError)(0), // 0: google.ads.googleads.v20.errors.AssetSetErrorEnum.AssetSetError
	(*AssetSetErrorEnum)(nil),            // 1: google.ads.googleads.v20.errors.AssetSetErrorEnum
}
var file_google_ads_googleads_v20_errors_asset_set_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v20_errors_asset_set_error_proto_init() }
func file_google_ads_googleads_v20_errors_asset_set_error_proto_init() {
	if File_google_ads_googleads_v20_errors_asset_set_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_errors_asset_set_error_proto_rawDesc), len(file_google_ads_googleads_v20_errors_asset_set_error_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v20_errors_asset_set_error_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v20_errors_asset_set_error_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v20_errors_asset_set_error_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v20_errors_asset_set_error_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v20_errors_asset_set_error_proto = out.File
	file_google_ads_googleads_v20_errors_asset_set_error_proto_goTypes = nil
	file_google_ads_googleads_v20_errors_asset_set_error_proto_depIdxs = nil
}

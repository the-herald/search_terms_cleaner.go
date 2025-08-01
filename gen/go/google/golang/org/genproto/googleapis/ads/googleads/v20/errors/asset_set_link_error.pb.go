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
// source: google/ads/googleads/v20/errors/asset_set_link_error.proto

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

// Enum describing possible asset set link errors.
type AssetSetLinkErrorEnum_AssetSetLinkError int32

const (
	// Enum unspecified.
	AssetSetLinkErrorEnum_UNSPECIFIED AssetSetLinkErrorEnum_AssetSetLinkError = 0
	// The received error code is not known in this version.
	AssetSetLinkErrorEnum_UNKNOWN AssetSetLinkErrorEnum_AssetSetLinkError = 1
	// Advertising channel type cannot be attached to the asset set due to
	// channel-based restrictions.
	AssetSetLinkErrorEnum_INCOMPATIBLE_ADVERTISING_CHANNEL_TYPE AssetSetLinkErrorEnum_AssetSetLinkError = 2
	// For this asset set type, only one campaign to feed linkage is allowed.
	AssetSetLinkErrorEnum_DUPLICATE_FEED_LINK AssetSetLinkErrorEnum_AssetSetLinkError = 3
	// The asset set type and campaign type are incompatible.
	AssetSetLinkErrorEnum_INCOMPATIBLE_ASSET_SET_TYPE_WITH_CAMPAIGN_TYPE AssetSetLinkErrorEnum_AssetSetLinkError = 4
	// Cannot link duplicate asset sets to the same campaign.
	AssetSetLinkErrorEnum_DUPLICATE_ASSET_SET_LINK AssetSetLinkErrorEnum_AssetSetLinkError = 5
	// Cannot remove the asset set link. If a campaign is linked with only one
	// asset set and you attempt to unlink them, this error will be triggered.
	AssetSetLinkErrorEnum_ASSET_SET_LINK_CANNOT_BE_REMOVED AssetSetLinkErrorEnum_AssetSetLinkError = 6
)

// Enum value maps for AssetSetLinkErrorEnum_AssetSetLinkError.
var (
	AssetSetLinkErrorEnum_AssetSetLinkError_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "INCOMPATIBLE_ADVERTISING_CHANNEL_TYPE",
		3: "DUPLICATE_FEED_LINK",
		4: "INCOMPATIBLE_ASSET_SET_TYPE_WITH_CAMPAIGN_TYPE",
		5: "DUPLICATE_ASSET_SET_LINK",
		6: "ASSET_SET_LINK_CANNOT_BE_REMOVED",
	}
	AssetSetLinkErrorEnum_AssetSetLinkError_value = map[string]int32{
		"UNSPECIFIED":                           0,
		"UNKNOWN":                               1,
		"INCOMPATIBLE_ADVERTISING_CHANNEL_TYPE": 2,
		"DUPLICATE_FEED_LINK":                   3,
		"INCOMPATIBLE_ASSET_SET_TYPE_WITH_CAMPAIGN_TYPE": 4,
		"DUPLICATE_ASSET_SET_LINK":                       5,
		"ASSET_SET_LINK_CANNOT_BE_REMOVED":               6,
	}
)

func (x AssetSetLinkErrorEnum_AssetSetLinkError) Enum() *AssetSetLinkErrorEnum_AssetSetLinkError {
	p := new(AssetSetLinkErrorEnum_AssetSetLinkError)
	*p = x
	return p
}

func (x AssetSetLinkErrorEnum_AssetSetLinkError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AssetSetLinkErrorEnum_AssetSetLinkError) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v20_errors_asset_set_link_error_proto_enumTypes[0].Descriptor()
}

func (AssetSetLinkErrorEnum_AssetSetLinkError) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v20_errors_asset_set_link_error_proto_enumTypes[0]
}

func (x AssetSetLinkErrorEnum_AssetSetLinkError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AssetSetLinkErrorEnum_AssetSetLinkError.Descriptor instead.
func (AssetSetLinkErrorEnum_AssetSetLinkError) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_errors_asset_set_link_error_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible asset set link errors.
type AssetSetLinkErrorEnum struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AssetSetLinkErrorEnum) Reset() {
	*x = AssetSetLinkErrorEnum{}
	mi := &file_google_ads_googleads_v20_errors_asset_set_link_error_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AssetSetLinkErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssetSetLinkErrorEnum) ProtoMessage() {}

func (x *AssetSetLinkErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_errors_asset_set_link_error_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssetSetLinkErrorEnum.ProtoReflect.Descriptor instead.
func (*AssetSetLinkErrorEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_errors_asset_set_link_error_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v20_errors_asset_set_link_error_proto protoreflect.FileDescriptor

const file_google_ads_googleads_v20_errors_asset_set_link_error_proto_rawDesc = "" +
	"\n" +
	":google/ads/googleads/v20/errors/asset_set_link_error.proto\x12\x1fgoogle.ads.googleads.v20.errors\"\x87\x02\n" +
	"\x15AssetSetLinkErrorEnum\"\xed\x01\n" +
	"\x11AssetSetLinkError\x12\x0f\n" +
	"\vUNSPECIFIED\x10\x00\x12\v\n" +
	"\aUNKNOWN\x10\x01\x12)\n" +
	"%INCOMPATIBLE_ADVERTISING_CHANNEL_TYPE\x10\x02\x12\x17\n" +
	"\x13DUPLICATE_FEED_LINK\x10\x03\x122\n" +
	".INCOMPATIBLE_ASSET_SET_TYPE_WITH_CAMPAIGN_TYPE\x10\x04\x12\x1c\n" +
	"\x18DUPLICATE_ASSET_SET_LINK\x10\x05\x12$\n" +
	" ASSET_SET_LINK_CANNOT_BE_REMOVED\x10\x06B\xf6\x01\n" +
	"#com.google.ads.googleads.v20.errorsB\x16AssetSetLinkErrorProtoP\x01ZEgoogle.golang.org/genproto/googleapis/ads/googleads/v20/errors;errors\xa2\x02\x03GAA\xaa\x02\x1fGoogle.Ads.GoogleAds.V20.Errors\xca\x02\x1fGoogle\\Ads\\GoogleAds\\V20\\Errors\xea\x02#Google::Ads::GoogleAds::V20::Errorsb\x06proto3"

var (
	file_google_ads_googleads_v20_errors_asset_set_link_error_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v20_errors_asset_set_link_error_proto_rawDescData []byte
)

func file_google_ads_googleads_v20_errors_asset_set_link_error_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v20_errors_asset_set_link_error_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v20_errors_asset_set_link_error_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_errors_asset_set_link_error_proto_rawDesc), len(file_google_ads_googleads_v20_errors_asset_set_link_error_proto_rawDesc)))
	})
	return file_google_ads_googleads_v20_errors_asset_set_link_error_proto_rawDescData
}

var file_google_ads_googleads_v20_errors_asset_set_link_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v20_errors_asset_set_link_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v20_errors_asset_set_link_error_proto_goTypes = []any{
	(AssetSetLinkErrorEnum_AssetSetLinkError)(0), // 0: google.ads.googleads.v20.errors.AssetSetLinkErrorEnum.AssetSetLinkError
	(*AssetSetLinkErrorEnum)(nil),                // 1: google.ads.googleads.v20.errors.AssetSetLinkErrorEnum
}
var file_google_ads_googleads_v20_errors_asset_set_link_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v20_errors_asset_set_link_error_proto_init() }
func file_google_ads_googleads_v20_errors_asset_set_link_error_proto_init() {
	if File_google_ads_googleads_v20_errors_asset_set_link_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_errors_asset_set_link_error_proto_rawDesc), len(file_google_ads_googleads_v20_errors_asset_set_link_error_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v20_errors_asset_set_link_error_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v20_errors_asset_set_link_error_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v20_errors_asset_set_link_error_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v20_errors_asset_set_link_error_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v20_errors_asset_set_link_error_proto = out.File
	file_google_ads_googleads_v20_errors_asset_set_link_error_proto_goTypes = nil
	file_google_ads_googleads_v20_errors_asset_set_link_error_proto_depIdxs = nil
}

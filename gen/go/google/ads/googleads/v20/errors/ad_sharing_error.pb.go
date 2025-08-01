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
// source: google/ads/googleads/v20/errors/ad_sharing_error.proto

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

// Enum describing possible ad sharing errors.
type AdSharingErrorEnum_AdSharingError int32

const (
	// Enum unspecified.
	AdSharingErrorEnum_UNSPECIFIED AdSharingErrorEnum_AdSharingError = 0
	// The received error code is not known in this version.
	AdSharingErrorEnum_UNKNOWN AdSharingErrorEnum_AdSharingError = 1
	// Error resulting in attempting to add an Ad to an AdGroup that already
	// contains the Ad.
	AdSharingErrorEnum_AD_GROUP_ALREADY_CONTAINS_AD AdSharingErrorEnum_AdSharingError = 2
	// Ad is not compatible with the AdGroup it is being shared with.
	AdSharingErrorEnum_INCOMPATIBLE_AD_UNDER_AD_GROUP AdSharingErrorEnum_AdSharingError = 3
	// Cannot add AdGroupAd on inactive Ad.
	AdSharingErrorEnum_CANNOT_SHARE_INACTIVE_AD AdSharingErrorEnum_AdSharingError = 4
)

// Enum value maps for AdSharingErrorEnum_AdSharingError.
var (
	AdSharingErrorEnum_AdSharingError_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "AD_GROUP_ALREADY_CONTAINS_AD",
		3: "INCOMPATIBLE_AD_UNDER_AD_GROUP",
		4: "CANNOT_SHARE_INACTIVE_AD",
	}
	AdSharingErrorEnum_AdSharingError_value = map[string]int32{
		"UNSPECIFIED":                    0,
		"UNKNOWN":                        1,
		"AD_GROUP_ALREADY_CONTAINS_AD":   2,
		"INCOMPATIBLE_AD_UNDER_AD_GROUP": 3,
		"CANNOT_SHARE_INACTIVE_AD":       4,
	}
)

func (x AdSharingErrorEnum_AdSharingError) Enum() *AdSharingErrorEnum_AdSharingError {
	p := new(AdSharingErrorEnum_AdSharingError)
	*p = x
	return p
}

func (x AdSharingErrorEnum_AdSharingError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AdSharingErrorEnum_AdSharingError) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v20_errors_ad_sharing_error_proto_enumTypes[0].Descriptor()
}

func (AdSharingErrorEnum_AdSharingError) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v20_errors_ad_sharing_error_proto_enumTypes[0]
}

func (x AdSharingErrorEnum_AdSharingError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AdSharingErrorEnum_AdSharingError.Descriptor instead.
func (AdSharingErrorEnum_AdSharingError) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_errors_ad_sharing_error_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible ad sharing errors.
type AdSharingErrorEnum struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AdSharingErrorEnum) Reset() {
	*x = AdSharingErrorEnum{}
	mi := &file_google_ads_googleads_v20_errors_ad_sharing_error_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AdSharingErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdSharingErrorEnum) ProtoMessage() {}

func (x *AdSharingErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_errors_ad_sharing_error_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdSharingErrorEnum.ProtoReflect.Descriptor instead.
func (*AdSharingErrorEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_errors_ad_sharing_error_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v20_errors_ad_sharing_error_proto protoreflect.FileDescriptor

const file_google_ads_googleads_v20_errors_ad_sharing_error_proto_rawDesc = "" +
	"\n" +
	"6google/ads/googleads/v20/errors/ad_sharing_error.proto\x12\x1fgoogle.ads.googleads.v20.errors\"\xa9\x01\n" +
	"\x12AdSharingErrorEnum\"\x92\x01\n" +
	"\x0eAdSharingError\x12\x0f\n" +
	"\vUNSPECIFIED\x10\x00\x12\v\n" +
	"\aUNKNOWN\x10\x01\x12 \n" +
	"\x1cAD_GROUP_ALREADY_CONTAINS_AD\x10\x02\x12\"\n" +
	"\x1eINCOMPATIBLE_AD_UNDER_AD_GROUP\x10\x03\x12\x1c\n" +
	"\x18CANNOT_SHARE_INACTIVE_AD\x10\x04B\xf3\x01\n" +
	"#com.google.ads.googleads.v20.errorsB\x13AdSharingErrorProtoP\x01ZEgoogle.golang.org/genproto/googleapis/ads/googleads/v20/errors;errors\xa2\x02\x03GAA\xaa\x02\x1fGoogle.Ads.GoogleAds.V20.Errors\xca\x02\x1fGoogle\\Ads\\GoogleAds\\V20\\Errors\xea\x02#Google::Ads::GoogleAds::V20::Errorsb\x06proto3"

var (
	file_google_ads_googleads_v20_errors_ad_sharing_error_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v20_errors_ad_sharing_error_proto_rawDescData []byte
)

func file_google_ads_googleads_v20_errors_ad_sharing_error_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v20_errors_ad_sharing_error_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v20_errors_ad_sharing_error_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_errors_ad_sharing_error_proto_rawDesc), len(file_google_ads_googleads_v20_errors_ad_sharing_error_proto_rawDesc)))
	})
	return file_google_ads_googleads_v20_errors_ad_sharing_error_proto_rawDescData
}

var file_google_ads_googleads_v20_errors_ad_sharing_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v20_errors_ad_sharing_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v20_errors_ad_sharing_error_proto_goTypes = []any{
	(AdSharingErrorEnum_AdSharingError)(0), // 0: google.ads.googleads.v20.errors.AdSharingErrorEnum.AdSharingError
	(*AdSharingErrorEnum)(nil),             // 1: google.ads.googleads.v20.errors.AdSharingErrorEnum
}
var file_google_ads_googleads_v20_errors_ad_sharing_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v20_errors_ad_sharing_error_proto_init() }
func file_google_ads_googleads_v20_errors_ad_sharing_error_proto_init() {
	if File_google_ads_googleads_v20_errors_ad_sharing_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_errors_ad_sharing_error_proto_rawDesc), len(file_google_ads_googleads_v20_errors_ad_sharing_error_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v20_errors_ad_sharing_error_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v20_errors_ad_sharing_error_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v20_errors_ad_sharing_error_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v20_errors_ad_sharing_error_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v20_errors_ad_sharing_error_proto = out.File
	file_google_ads_googleads_v20_errors_ad_sharing_error_proto_goTypes = nil
	file_google_ads_googleads_v20_errors_ad_sharing_error_proto_depIdxs = nil
}

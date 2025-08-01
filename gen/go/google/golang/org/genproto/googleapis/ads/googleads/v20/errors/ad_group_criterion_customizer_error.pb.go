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
// source: google/ads/googleads/v20/errors/ad_group_criterion_customizer_error.proto

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

// Enum describing possible ad group criterion customizer errors.
type AdGroupCriterionCustomizerErrorEnum_AdGroupCriterionCustomizerError int32

const (
	// Enum unspecified.
	AdGroupCriterionCustomizerErrorEnum_UNSPECIFIED AdGroupCriterionCustomizerErrorEnum_AdGroupCriterionCustomizerError = 0
	// The received error code is not known in this version.
	AdGroupCriterionCustomizerErrorEnum_UNKNOWN AdGroupCriterionCustomizerErrorEnum_AdGroupCriterionCustomizerError = 1
	// Only keyword type criterion is allowed to link customizer attribute.
	AdGroupCriterionCustomizerErrorEnum_CRITERION_IS_NOT_KEYWORD AdGroupCriterionCustomizerErrorEnum_AdGroupCriterionCustomizerError = 2
)

// Enum value maps for AdGroupCriterionCustomizerErrorEnum_AdGroupCriterionCustomizerError.
var (
	AdGroupCriterionCustomizerErrorEnum_AdGroupCriterionCustomizerError_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "CRITERION_IS_NOT_KEYWORD",
	}
	AdGroupCriterionCustomizerErrorEnum_AdGroupCriterionCustomizerError_value = map[string]int32{
		"UNSPECIFIED":              0,
		"UNKNOWN":                  1,
		"CRITERION_IS_NOT_KEYWORD": 2,
	}
)

func (x AdGroupCriterionCustomizerErrorEnum_AdGroupCriterionCustomizerError) Enum() *AdGroupCriterionCustomizerErrorEnum_AdGroupCriterionCustomizerError {
	p := new(AdGroupCriterionCustomizerErrorEnum_AdGroupCriterionCustomizerError)
	*p = x
	return p
}

func (x AdGroupCriterionCustomizerErrorEnum_AdGroupCriterionCustomizerError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AdGroupCriterionCustomizerErrorEnum_AdGroupCriterionCustomizerError) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v20_errors_ad_group_criterion_customizer_error_proto_enumTypes[0].Descriptor()
}

func (AdGroupCriterionCustomizerErrorEnum_AdGroupCriterionCustomizerError) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v20_errors_ad_group_criterion_customizer_error_proto_enumTypes[0]
}

func (x AdGroupCriterionCustomizerErrorEnum_AdGroupCriterionCustomizerError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AdGroupCriterionCustomizerErrorEnum_AdGroupCriterionCustomizerError.Descriptor instead.
func (AdGroupCriterionCustomizerErrorEnum_AdGroupCriterionCustomizerError) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_errors_ad_group_criterion_customizer_error_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible ad group criterion customizer errors.
type AdGroupCriterionCustomizerErrorEnum struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AdGroupCriterionCustomizerErrorEnum) Reset() {
	*x = AdGroupCriterionCustomizerErrorEnum{}
	mi := &file_google_ads_googleads_v20_errors_ad_group_criterion_customizer_error_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AdGroupCriterionCustomizerErrorEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdGroupCriterionCustomizerErrorEnum) ProtoMessage() {}

func (x *AdGroupCriterionCustomizerErrorEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_errors_ad_group_criterion_customizer_error_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdGroupCriterionCustomizerErrorEnum.ProtoReflect.Descriptor instead.
func (*AdGroupCriterionCustomizerErrorEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_errors_ad_group_criterion_customizer_error_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v20_errors_ad_group_criterion_customizer_error_proto protoreflect.FileDescriptor

const file_google_ads_googleads_v20_errors_ad_group_criterion_customizer_error_proto_rawDesc = "" +
	"\n" +
	"Igoogle/ads/googleads/v20/errors/ad_group_criterion_customizer_error.proto\x12\x1fgoogle.ads.googleads.v20.errors\"\x84\x01\n" +
	"#AdGroupCriterionCustomizerErrorEnum\"]\n" +
	"\x1fAdGroupCriterionCustomizerError\x12\x0f\n" +
	"\vUNSPECIFIED\x10\x00\x12\v\n" +
	"\aUNKNOWN\x10\x01\x12\x1c\n" +
	"\x18CRITERION_IS_NOT_KEYWORD\x10\x02B\x84\x02\n" +
	"#com.google.ads.googleads.v20.errorsB$AdGroupCriterionCustomizerErrorProtoP\x01ZEgoogle.golang.org/genproto/googleapis/ads/googleads/v20/errors;errors\xa2\x02\x03GAA\xaa\x02\x1fGoogle.Ads.GoogleAds.V20.Errors\xca\x02\x1fGoogle\\Ads\\GoogleAds\\V20\\Errors\xea\x02#Google::Ads::GoogleAds::V20::Errorsb\x06proto3"

var (
	file_google_ads_googleads_v20_errors_ad_group_criterion_customizer_error_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v20_errors_ad_group_criterion_customizer_error_proto_rawDescData []byte
)

func file_google_ads_googleads_v20_errors_ad_group_criterion_customizer_error_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v20_errors_ad_group_criterion_customizer_error_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v20_errors_ad_group_criterion_customizer_error_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_errors_ad_group_criterion_customizer_error_proto_rawDesc), len(file_google_ads_googleads_v20_errors_ad_group_criterion_customizer_error_proto_rawDesc)))
	})
	return file_google_ads_googleads_v20_errors_ad_group_criterion_customizer_error_proto_rawDescData
}

var file_google_ads_googleads_v20_errors_ad_group_criterion_customizer_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v20_errors_ad_group_criterion_customizer_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v20_errors_ad_group_criterion_customizer_error_proto_goTypes = []any{
	(AdGroupCriterionCustomizerErrorEnum_AdGroupCriterionCustomizerError)(0), // 0: google.ads.googleads.v20.errors.AdGroupCriterionCustomizerErrorEnum.AdGroupCriterionCustomizerError
	(*AdGroupCriterionCustomizerErrorEnum)(nil),                              // 1: google.ads.googleads.v20.errors.AdGroupCriterionCustomizerErrorEnum
}
var file_google_ads_googleads_v20_errors_ad_group_criterion_customizer_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v20_errors_ad_group_criterion_customizer_error_proto_init() }
func file_google_ads_googleads_v20_errors_ad_group_criterion_customizer_error_proto_init() {
	if File_google_ads_googleads_v20_errors_ad_group_criterion_customizer_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_errors_ad_group_criterion_customizer_error_proto_rawDesc), len(file_google_ads_googleads_v20_errors_ad_group_criterion_customizer_error_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v20_errors_ad_group_criterion_customizer_error_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v20_errors_ad_group_criterion_customizer_error_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v20_errors_ad_group_criterion_customizer_error_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v20_errors_ad_group_criterion_customizer_error_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v20_errors_ad_group_criterion_customizer_error_proto = out.File
	file_google_ads_googleads_v20_errors_ad_group_criterion_customizer_error_proto_goTypes = nil
	file_google_ads_googleads_v20_errors_ad_group_criterion_customizer_error_proto_depIdxs = nil
}

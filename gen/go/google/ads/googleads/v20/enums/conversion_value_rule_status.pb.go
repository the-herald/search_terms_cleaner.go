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
// source: google/ads/googleads/v20/enums/conversion_value_rule_status.proto

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

// Possible statuses of a conversion value rule.
type ConversionValueRuleStatusEnum_ConversionValueRuleStatus int32

const (
	// Not specified.
	ConversionValueRuleStatusEnum_UNSPECIFIED ConversionValueRuleStatusEnum_ConversionValueRuleStatus = 0
	// Used for return value only. Represents value unknown in this version.
	ConversionValueRuleStatusEnum_UNKNOWN ConversionValueRuleStatusEnum_ConversionValueRuleStatus = 1
	// Conversion Value Rule is enabled and can be applied.
	ConversionValueRuleStatusEnum_ENABLED ConversionValueRuleStatusEnum_ConversionValueRuleStatus = 2
	// Conversion Value Rule is permanently deleted and can't be applied.
	ConversionValueRuleStatusEnum_REMOVED ConversionValueRuleStatusEnum_ConversionValueRuleStatus = 3
	// Conversion Value Rule is paused, but can be re-enabled.
	ConversionValueRuleStatusEnum_PAUSED ConversionValueRuleStatusEnum_ConversionValueRuleStatus = 4
)

// Enum value maps for ConversionValueRuleStatusEnum_ConversionValueRuleStatus.
var (
	ConversionValueRuleStatusEnum_ConversionValueRuleStatus_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "ENABLED",
		3: "REMOVED",
		4: "PAUSED",
	}
	ConversionValueRuleStatusEnum_ConversionValueRuleStatus_value = map[string]int32{
		"UNSPECIFIED": 0,
		"UNKNOWN":     1,
		"ENABLED":     2,
		"REMOVED":     3,
		"PAUSED":      4,
	}
)

func (x ConversionValueRuleStatusEnum_ConversionValueRuleStatus) Enum() *ConversionValueRuleStatusEnum_ConversionValueRuleStatus {
	p := new(ConversionValueRuleStatusEnum_ConversionValueRuleStatus)
	*p = x
	return p
}

func (x ConversionValueRuleStatusEnum_ConversionValueRuleStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConversionValueRuleStatusEnum_ConversionValueRuleStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v20_enums_conversion_value_rule_status_proto_enumTypes[0].Descriptor()
}

func (ConversionValueRuleStatusEnum_ConversionValueRuleStatus) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v20_enums_conversion_value_rule_status_proto_enumTypes[0]
}

func (x ConversionValueRuleStatusEnum_ConversionValueRuleStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConversionValueRuleStatusEnum_ConversionValueRuleStatus.Descriptor instead.
func (ConversionValueRuleStatusEnum_ConversionValueRuleStatus) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_enums_conversion_value_rule_status_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible statuses of a conversion value rule.
type ConversionValueRuleStatusEnum struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConversionValueRuleStatusEnum) Reset() {
	*x = ConversionValueRuleStatusEnum{}
	mi := &file_google_ads_googleads_v20_enums_conversion_value_rule_status_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConversionValueRuleStatusEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConversionValueRuleStatusEnum) ProtoMessage() {}

func (x *ConversionValueRuleStatusEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_enums_conversion_value_rule_status_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConversionValueRuleStatusEnum.ProtoReflect.Descriptor instead.
func (*ConversionValueRuleStatusEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_enums_conversion_value_rule_status_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v20_enums_conversion_value_rule_status_proto protoreflect.FileDescriptor

const file_google_ads_googleads_v20_enums_conversion_value_rule_status_proto_rawDesc = "" +
	"\n" +
	"Agoogle/ads/googleads/v20/enums/conversion_value_rule_status.proto\x12\x1egoogle.ads.googleads.v20.enums\"\x80\x01\n" +
	"\x1dConversionValueRuleStatusEnum\"_\n" +
	"\x19ConversionValueRuleStatus\x12\x0f\n" +
	"\vUNSPECIFIED\x10\x00\x12\v\n" +
	"\aUNKNOWN\x10\x01\x12\v\n" +
	"\aENABLED\x10\x02\x12\v\n" +
	"\aREMOVED\x10\x03\x12\n" +
	"\n" +
	"\x06PAUSED\x10\x04B\xf8\x01\n" +
	"\"com.google.ads.googleads.v20.enumsB\x1eConversionValueRuleStatusProtoP\x01ZCgoogle.golang.org/genproto/googleapis/ads/googleads/v20/enums;enums\xa2\x02\x03GAA\xaa\x02\x1eGoogle.Ads.GoogleAds.V20.Enums\xca\x02\x1eGoogle\\Ads\\GoogleAds\\V20\\Enums\xea\x02\"Google::Ads::GoogleAds::V20::Enumsb\x06proto3"

var (
	file_google_ads_googleads_v20_enums_conversion_value_rule_status_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v20_enums_conversion_value_rule_status_proto_rawDescData []byte
)

func file_google_ads_googleads_v20_enums_conversion_value_rule_status_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v20_enums_conversion_value_rule_status_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v20_enums_conversion_value_rule_status_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_enums_conversion_value_rule_status_proto_rawDesc), len(file_google_ads_googleads_v20_enums_conversion_value_rule_status_proto_rawDesc)))
	})
	return file_google_ads_googleads_v20_enums_conversion_value_rule_status_proto_rawDescData
}

var file_google_ads_googleads_v20_enums_conversion_value_rule_status_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v20_enums_conversion_value_rule_status_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v20_enums_conversion_value_rule_status_proto_goTypes = []any{
	(ConversionValueRuleStatusEnum_ConversionValueRuleStatus)(0), // 0: google.ads.googleads.v20.enums.ConversionValueRuleStatusEnum.ConversionValueRuleStatus
	(*ConversionValueRuleStatusEnum)(nil),                        // 1: google.ads.googleads.v20.enums.ConversionValueRuleStatusEnum
}
var file_google_ads_googleads_v20_enums_conversion_value_rule_status_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v20_enums_conversion_value_rule_status_proto_init() }
func file_google_ads_googleads_v20_enums_conversion_value_rule_status_proto_init() {
	if File_google_ads_googleads_v20_enums_conversion_value_rule_status_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_enums_conversion_value_rule_status_proto_rawDesc), len(file_google_ads_googleads_v20_enums_conversion_value_rule_status_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v20_enums_conversion_value_rule_status_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v20_enums_conversion_value_rule_status_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v20_enums_conversion_value_rule_status_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v20_enums_conversion_value_rule_status_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v20_enums_conversion_value_rule_status_proto = out.File
	file_google_ads_googleads_v20_enums_conversion_value_rule_status_proto_goTypes = nil
	file_google_ads_googleads_v20_enums_conversion_value_rule_status_proto_depIdxs = nil
}

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
// source: google/ads/googleads/v20/enums/conversion_value_rule_primary_dimension.proto

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

// Identifies the primary dimension for conversion value rule stats.
type ConversionValueRulePrimaryDimensionEnum_ConversionValueRulePrimaryDimension int32

const (
	// Not specified.
	ConversionValueRulePrimaryDimensionEnum_UNSPECIFIED ConversionValueRulePrimaryDimensionEnum_ConversionValueRulePrimaryDimension = 0
	// Used for return value only. Represents value unknown in this version.
	ConversionValueRulePrimaryDimensionEnum_UNKNOWN ConversionValueRulePrimaryDimensionEnum_ConversionValueRulePrimaryDimension = 1
	// For no-value-rule-applied conversions after value rule is enabled.
	ConversionValueRulePrimaryDimensionEnum_NO_RULE_APPLIED ConversionValueRulePrimaryDimensionEnum_ConversionValueRulePrimaryDimension = 2
	// Below are for value-rule-applied conversions:
	// The original stats.
	ConversionValueRulePrimaryDimensionEnum_ORIGINAL ConversionValueRulePrimaryDimensionEnum_ConversionValueRulePrimaryDimension = 3
	// When a new or returning customer condition is satisfied.
	ConversionValueRulePrimaryDimensionEnum_NEW_VS_RETURNING_USER ConversionValueRulePrimaryDimensionEnum_ConversionValueRulePrimaryDimension = 4
	// When a query-time Geo location condition is satisfied.
	ConversionValueRulePrimaryDimensionEnum_GEO_LOCATION ConversionValueRulePrimaryDimensionEnum_ConversionValueRulePrimaryDimension = 5
	// When a query-time browsing device condition is satisfied.
	ConversionValueRulePrimaryDimensionEnum_DEVICE ConversionValueRulePrimaryDimensionEnum_ConversionValueRulePrimaryDimension = 6
	// When a query-time audience condition is satisfied.
	ConversionValueRulePrimaryDimensionEnum_AUDIENCE ConversionValueRulePrimaryDimensionEnum_ConversionValueRulePrimaryDimension = 7
	// When multiple rules are applied.
	ConversionValueRulePrimaryDimensionEnum_MULTIPLE ConversionValueRulePrimaryDimensionEnum_ConversionValueRulePrimaryDimension = 8
	// When a query-time itinerary condition is satisfied.
	ConversionValueRulePrimaryDimensionEnum_ITINERARY ConversionValueRulePrimaryDimensionEnum_ConversionValueRulePrimaryDimension = 9
)

// Enum value maps for ConversionValueRulePrimaryDimensionEnum_ConversionValueRulePrimaryDimension.
var (
	ConversionValueRulePrimaryDimensionEnum_ConversionValueRulePrimaryDimension_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "NO_RULE_APPLIED",
		3: "ORIGINAL",
		4: "NEW_VS_RETURNING_USER",
		5: "GEO_LOCATION",
		6: "DEVICE",
		7: "AUDIENCE",
		8: "MULTIPLE",
		9: "ITINERARY",
	}
	ConversionValueRulePrimaryDimensionEnum_ConversionValueRulePrimaryDimension_value = map[string]int32{
		"UNSPECIFIED":           0,
		"UNKNOWN":               1,
		"NO_RULE_APPLIED":       2,
		"ORIGINAL":              3,
		"NEW_VS_RETURNING_USER": 4,
		"GEO_LOCATION":          5,
		"DEVICE":                6,
		"AUDIENCE":              7,
		"MULTIPLE":              8,
		"ITINERARY":             9,
	}
)

func (x ConversionValueRulePrimaryDimensionEnum_ConversionValueRulePrimaryDimension) Enum() *ConversionValueRulePrimaryDimensionEnum_ConversionValueRulePrimaryDimension {
	p := new(ConversionValueRulePrimaryDimensionEnum_ConversionValueRulePrimaryDimension)
	*p = x
	return p
}

func (x ConversionValueRulePrimaryDimensionEnum_ConversionValueRulePrimaryDimension) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConversionValueRulePrimaryDimensionEnum_ConversionValueRulePrimaryDimension) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v20_enums_conversion_value_rule_primary_dimension_proto_enumTypes[0].Descriptor()
}

func (ConversionValueRulePrimaryDimensionEnum_ConversionValueRulePrimaryDimension) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v20_enums_conversion_value_rule_primary_dimension_proto_enumTypes[0]
}

func (x ConversionValueRulePrimaryDimensionEnum_ConversionValueRulePrimaryDimension) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConversionValueRulePrimaryDimensionEnum_ConversionValueRulePrimaryDimension.Descriptor instead.
func (ConversionValueRulePrimaryDimensionEnum_ConversionValueRulePrimaryDimension) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_enums_conversion_value_rule_primary_dimension_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing value rule primary dimension for stats.
type ConversionValueRulePrimaryDimensionEnum struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConversionValueRulePrimaryDimensionEnum) Reset() {
	*x = ConversionValueRulePrimaryDimensionEnum{}
	mi := &file_google_ads_googleads_v20_enums_conversion_value_rule_primary_dimension_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConversionValueRulePrimaryDimensionEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConversionValueRulePrimaryDimensionEnum) ProtoMessage() {}

func (x *ConversionValueRulePrimaryDimensionEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_enums_conversion_value_rule_primary_dimension_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConversionValueRulePrimaryDimensionEnum.ProtoReflect.Descriptor instead.
func (*ConversionValueRulePrimaryDimensionEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_enums_conversion_value_rule_primary_dimension_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v20_enums_conversion_value_rule_primary_dimension_proto protoreflect.FileDescriptor

const file_google_ads_googleads_v20_enums_conversion_value_rule_primary_dimension_proto_rawDesc = "" +
	"\n" +
	"Lgoogle/ads/googleads/v20/enums/conversion_value_rule_primary_dimension.proto\x12\x1egoogle.ads.googleads.v20.enums\"\xf6\x01\n" +
	"'ConversionValueRulePrimaryDimensionEnum\"\xca\x01\n" +
	"#ConversionValueRulePrimaryDimension\x12\x0f\n" +
	"\vUNSPECIFIED\x10\x00\x12\v\n" +
	"\aUNKNOWN\x10\x01\x12\x13\n" +
	"\x0fNO_RULE_APPLIED\x10\x02\x12\f\n" +
	"\bORIGINAL\x10\x03\x12\x19\n" +
	"\x15NEW_VS_RETURNING_USER\x10\x04\x12\x10\n" +
	"\fGEO_LOCATION\x10\x05\x12\n" +
	"\n" +
	"\x06DEVICE\x10\x06\x12\f\n" +
	"\bAUDIENCE\x10\a\x12\f\n" +
	"\bMULTIPLE\x10\b\x12\r\n" +
	"\tITINERARY\x10\tB\x82\x02\n" +
	"\"com.google.ads.googleads.v20.enumsB(ConversionValueRulePrimaryDimensionProtoP\x01ZCgoogle.golang.org/genproto/googleapis/ads/googleads/v20/enums;enums\xa2\x02\x03GAA\xaa\x02\x1eGoogle.Ads.GoogleAds.V20.Enums\xca\x02\x1eGoogle\\Ads\\GoogleAds\\V20\\Enums\xea\x02\"Google::Ads::GoogleAds::V20::Enumsb\x06proto3"

var (
	file_google_ads_googleads_v20_enums_conversion_value_rule_primary_dimension_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v20_enums_conversion_value_rule_primary_dimension_proto_rawDescData []byte
)

func file_google_ads_googleads_v20_enums_conversion_value_rule_primary_dimension_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v20_enums_conversion_value_rule_primary_dimension_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v20_enums_conversion_value_rule_primary_dimension_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_enums_conversion_value_rule_primary_dimension_proto_rawDesc), len(file_google_ads_googleads_v20_enums_conversion_value_rule_primary_dimension_proto_rawDesc)))
	})
	return file_google_ads_googleads_v20_enums_conversion_value_rule_primary_dimension_proto_rawDescData
}

var file_google_ads_googleads_v20_enums_conversion_value_rule_primary_dimension_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v20_enums_conversion_value_rule_primary_dimension_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v20_enums_conversion_value_rule_primary_dimension_proto_goTypes = []any{
	(ConversionValueRulePrimaryDimensionEnum_ConversionValueRulePrimaryDimension)(0), // 0: google.ads.googleads.v20.enums.ConversionValueRulePrimaryDimensionEnum.ConversionValueRulePrimaryDimension
	(*ConversionValueRulePrimaryDimensionEnum)(nil),                                  // 1: google.ads.googleads.v20.enums.ConversionValueRulePrimaryDimensionEnum
}
var file_google_ads_googleads_v20_enums_conversion_value_rule_primary_dimension_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v20_enums_conversion_value_rule_primary_dimension_proto_init() }
func file_google_ads_googleads_v20_enums_conversion_value_rule_primary_dimension_proto_init() {
	if File_google_ads_googleads_v20_enums_conversion_value_rule_primary_dimension_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_enums_conversion_value_rule_primary_dimension_proto_rawDesc), len(file_google_ads_googleads_v20_enums_conversion_value_rule_primary_dimension_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v20_enums_conversion_value_rule_primary_dimension_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v20_enums_conversion_value_rule_primary_dimension_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v20_enums_conversion_value_rule_primary_dimension_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v20_enums_conversion_value_rule_primary_dimension_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v20_enums_conversion_value_rule_primary_dimension_proto = out.File
	file_google_ads_googleads_v20_enums_conversion_value_rule_primary_dimension_proto_goTypes = nil
	file_google_ads_googleads_v20_enums_conversion_value_rule_primary_dimension_proto_depIdxs = nil
}

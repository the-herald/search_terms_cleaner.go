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
// source: google/ads/googleads/v20/enums/conversion_adjustment_type.proto

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

// The different actions advertisers can take to adjust the conversions that
// they already reported. Retractions negate a conversion. Restatements change
// the value of a conversion.
type ConversionAdjustmentTypeEnum_ConversionAdjustmentType int32

const (
	// Not specified.
	ConversionAdjustmentTypeEnum_UNSPECIFIED ConversionAdjustmentTypeEnum_ConversionAdjustmentType = 0
	// Represents value unknown in this version.
	ConversionAdjustmentTypeEnum_UNKNOWN ConversionAdjustmentTypeEnum_ConversionAdjustmentType = 1
	// Negates a conversion so that its total value and count are both zero.
	ConversionAdjustmentTypeEnum_RETRACTION ConversionAdjustmentTypeEnum_ConversionAdjustmentType = 2
	// Changes the value of a conversion.
	ConversionAdjustmentTypeEnum_RESTATEMENT ConversionAdjustmentTypeEnum_ConversionAdjustmentType = 3
	// Supplements an existing conversion with provided user identifiers and
	// user agent, which can be used by Google to enhance the conversion count.
	ConversionAdjustmentTypeEnum_ENHANCEMENT ConversionAdjustmentTypeEnum_ConversionAdjustmentType = 4
)

// Enum value maps for ConversionAdjustmentTypeEnum_ConversionAdjustmentType.
var (
	ConversionAdjustmentTypeEnum_ConversionAdjustmentType_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "RETRACTION",
		3: "RESTATEMENT",
		4: "ENHANCEMENT",
	}
	ConversionAdjustmentTypeEnum_ConversionAdjustmentType_value = map[string]int32{
		"UNSPECIFIED": 0,
		"UNKNOWN":     1,
		"RETRACTION":  2,
		"RESTATEMENT": 3,
		"ENHANCEMENT": 4,
	}
)

func (x ConversionAdjustmentTypeEnum_ConversionAdjustmentType) Enum() *ConversionAdjustmentTypeEnum_ConversionAdjustmentType {
	p := new(ConversionAdjustmentTypeEnum_ConversionAdjustmentType)
	*p = x
	return p
}

func (x ConversionAdjustmentTypeEnum_ConversionAdjustmentType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConversionAdjustmentTypeEnum_ConversionAdjustmentType) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v20_enums_conversion_adjustment_type_proto_enumTypes[0].Descriptor()
}

func (ConversionAdjustmentTypeEnum_ConversionAdjustmentType) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v20_enums_conversion_adjustment_type_proto_enumTypes[0]
}

func (x ConversionAdjustmentTypeEnum_ConversionAdjustmentType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConversionAdjustmentTypeEnum_ConversionAdjustmentType.Descriptor instead.
func (ConversionAdjustmentTypeEnum_ConversionAdjustmentType) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_enums_conversion_adjustment_type_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing conversion adjustment types.
type ConversionAdjustmentTypeEnum struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConversionAdjustmentTypeEnum) Reset() {
	*x = ConversionAdjustmentTypeEnum{}
	mi := &file_google_ads_googleads_v20_enums_conversion_adjustment_type_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConversionAdjustmentTypeEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConversionAdjustmentTypeEnum) ProtoMessage() {}

func (x *ConversionAdjustmentTypeEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_enums_conversion_adjustment_type_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConversionAdjustmentTypeEnum.ProtoReflect.Descriptor instead.
func (*ConversionAdjustmentTypeEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_enums_conversion_adjustment_type_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v20_enums_conversion_adjustment_type_proto protoreflect.FileDescriptor

const file_google_ads_googleads_v20_enums_conversion_adjustment_type_proto_rawDesc = "" +
	"\n" +
	"?google/ads/googleads/v20/enums/conversion_adjustment_type.proto\x12\x1egoogle.ads.googleads.v20.enums\"\x8a\x01\n" +
	"\x1cConversionAdjustmentTypeEnum\"j\n" +
	"\x18ConversionAdjustmentType\x12\x0f\n" +
	"\vUNSPECIFIED\x10\x00\x12\v\n" +
	"\aUNKNOWN\x10\x01\x12\x0e\n" +
	"\n" +
	"RETRACTION\x10\x02\x12\x0f\n" +
	"\vRESTATEMENT\x10\x03\x12\x0f\n" +
	"\vENHANCEMENT\x10\x04B\xf7\x01\n" +
	"\"com.google.ads.googleads.v20.enumsB\x1dConversionAdjustmentTypeProtoP\x01ZCgoogle.golang.org/genproto/googleapis/ads/googleads/v20/enums;enums\xa2\x02\x03GAA\xaa\x02\x1eGoogle.Ads.GoogleAds.V20.Enums\xca\x02\x1eGoogle\\Ads\\GoogleAds\\V20\\Enums\xea\x02\"Google::Ads::GoogleAds::V20::Enumsb\x06proto3"

var (
	file_google_ads_googleads_v20_enums_conversion_adjustment_type_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v20_enums_conversion_adjustment_type_proto_rawDescData []byte
)

func file_google_ads_googleads_v20_enums_conversion_adjustment_type_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v20_enums_conversion_adjustment_type_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v20_enums_conversion_adjustment_type_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_enums_conversion_adjustment_type_proto_rawDesc), len(file_google_ads_googleads_v20_enums_conversion_adjustment_type_proto_rawDesc)))
	})
	return file_google_ads_googleads_v20_enums_conversion_adjustment_type_proto_rawDescData
}

var file_google_ads_googleads_v20_enums_conversion_adjustment_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v20_enums_conversion_adjustment_type_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v20_enums_conversion_adjustment_type_proto_goTypes = []any{
	(ConversionAdjustmentTypeEnum_ConversionAdjustmentType)(0), // 0: google.ads.googleads.v20.enums.ConversionAdjustmentTypeEnum.ConversionAdjustmentType
	(*ConversionAdjustmentTypeEnum)(nil),                       // 1: google.ads.googleads.v20.enums.ConversionAdjustmentTypeEnum
}
var file_google_ads_googleads_v20_enums_conversion_adjustment_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v20_enums_conversion_adjustment_type_proto_init() }
func file_google_ads_googleads_v20_enums_conversion_adjustment_type_proto_init() {
	if File_google_ads_googleads_v20_enums_conversion_adjustment_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_enums_conversion_adjustment_type_proto_rawDesc), len(file_google_ads_googleads_v20_enums_conversion_adjustment_type_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v20_enums_conversion_adjustment_type_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v20_enums_conversion_adjustment_type_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v20_enums_conversion_adjustment_type_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v20_enums_conversion_adjustment_type_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v20_enums_conversion_adjustment_type_proto = out.File
	file_google_ads_googleads_v20_enums_conversion_adjustment_type_proto_goTypes = nil
	file_google_ads_googleads_v20_enums_conversion_adjustment_type_proto_depIdxs = nil
}

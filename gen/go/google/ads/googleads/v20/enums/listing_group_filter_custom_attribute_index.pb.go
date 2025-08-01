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
// source: google/ads/googleads/v20/enums/listing_group_filter_custom_attribute_index.proto

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

// The index of customer attributes.
type ListingGroupFilterCustomAttributeIndexEnum_ListingGroupFilterCustomAttributeIndex int32

const (
	// Not specified.
	ListingGroupFilterCustomAttributeIndexEnum_UNSPECIFIED ListingGroupFilterCustomAttributeIndexEnum_ListingGroupFilterCustomAttributeIndex = 0
	// Used for return value only. Represents value unknown in this version.
	ListingGroupFilterCustomAttributeIndexEnum_UNKNOWN ListingGroupFilterCustomAttributeIndexEnum_ListingGroupFilterCustomAttributeIndex = 1
	// First listing group filter custom attribute.
	ListingGroupFilterCustomAttributeIndexEnum_INDEX0 ListingGroupFilterCustomAttributeIndexEnum_ListingGroupFilterCustomAttributeIndex = 2
	// Second listing group filter custom attribute.
	ListingGroupFilterCustomAttributeIndexEnum_INDEX1 ListingGroupFilterCustomAttributeIndexEnum_ListingGroupFilterCustomAttributeIndex = 3
	// Third listing group filter custom attribute.
	ListingGroupFilterCustomAttributeIndexEnum_INDEX2 ListingGroupFilterCustomAttributeIndexEnum_ListingGroupFilterCustomAttributeIndex = 4
	// Fourth listing group filter custom attribute.
	ListingGroupFilterCustomAttributeIndexEnum_INDEX3 ListingGroupFilterCustomAttributeIndexEnum_ListingGroupFilterCustomAttributeIndex = 5
	// Fifth listing group filter custom attribute.
	ListingGroupFilterCustomAttributeIndexEnum_INDEX4 ListingGroupFilterCustomAttributeIndexEnum_ListingGroupFilterCustomAttributeIndex = 6
)

// Enum value maps for ListingGroupFilterCustomAttributeIndexEnum_ListingGroupFilterCustomAttributeIndex.
var (
	ListingGroupFilterCustomAttributeIndexEnum_ListingGroupFilterCustomAttributeIndex_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "INDEX0",
		3: "INDEX1",
		4: "INDEX2",
		5: "INDEX3",
		6: "INDEX4",
	}
	ListingGroupFilterCustomAttributeIndexEnum_ListingGroupFilterCustomAttributeIndex_value = map[string]int32{
		"UNSPECIFIED": 0,
		"UNKNOWN":     1,
		"INDEX0":      2,
		"INDEX1":      3,
		"INDEX2":      4,
		"INDEX3":      5,
		"INDEX4":      6,
	}
)

func (x ListingGroupFilterCustomAttributeIndexEnum_ListingGroupFilterCustomAttributeIndex) Enum() *ListingGroupFilterCustomAttributeIndexEnum_ListingGroupFilterCustomAttributeIndex {
	p := new(ListingGroupFilterCustomAttributeIndexEnum_ListingGroupFilterCustomAttributeIndex)
	*p = x
	return p
}

func (x ListingGroupFilterCustomAttributeIndexEnum_ListingGroupFilterCustomAttributeIndex) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ListingGroupFilterCustomAttributeIndexEnum_ListingGroupFilterCustomAttributeIndex) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v20_enums_listing_group_filter_custom_attribute_index_proto_enumTypes[0].Descriptor()
}

func (ListingGroupFilterCustomAttributeIndexEnum_ListingGroupFilterCustomAttributeIndex) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v20_enums_listing_group_filter_custom_attribute_index_proto_enumTypes[0]
}

func (x ListingGroupFilterCustomAttributeIndexEnum_ListingGroupFilterCustomAttributeIndex) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ListingGroupFilterCustomAttributeIndexEnum_ListingGroupFilterCustomAttributeIndex.Descriptor instead.
func (ListingGroupFilterCustomAttributeIndexEnum_ListingGroupFilterCustomAttributeIndex) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_enums_listing_group_filter_custom_attribute_index_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing the indexes of custom attribute used in
// ListingGroupFilterDimension.
type ListingGroupFilterCustomAttributeIndexEnum struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListingGroupFilterCustomAttributeIndexEnum) Reset() {
	*x = ListingGroupFilterCustomAttributeIndexEnum{}
	mi := &file_google_ads_googleads_v20_enums_listing_group_filter_custom_attribute_index_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListingGroupFilterCustomAttributeIndexEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListingGroupFilterCustomAttributeIndexEnum) ProtoMessage() {}

func (x *ListingGroupFilterCustomAttributeIndexEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_enums_listing_group_filter_custom_attribute_index_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListingGroupFilterCustomAttributeIndexEnum.ProtoReflect.Descriptor instead.
func (*ListingGroupFilterCustomAttributeIndexEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_enums_listing_group_filter_custom_attribute_index_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v20_enums_listing_group_filter_custom_attribute_index_proto protoreflect.FileDescriptor

const file_google_ads_googleads_v20_enums_listing_group_filter_custom_attribute_index_proto_rawDesc = "" +
	"\n" +
	"Pgoogle/ads/googleads/v20/enums/listing_group_filter_custom_attribute_index.proto\x12\x1egoogle.ads.googleads.v20.enums\"\xb1\x01\n" +
	"*ListingGroupFilterCustomAttributeIndexEnum\"\x82\x01\n" +
	"&ListingGroupFilterCustomAttributeIndex\x12\x0f\n" +
	"\vUNSPECIFIED\x10\x00\x12\v\n" +
	"\aUNKNOWN\x10\x01\x12\n" +
	"\n" +
	"\x06INDEX0\x10\x02\x12\n" +
	"\n" +
	"\x06INDEX1\x10\x03\x12\n" +
	"\n" +
	"\x06INDEX2\x10\x04\x12\n" +
	"\n" +
	"\x06INDEX3\x10\x05\x12\n" +
	"\n" +
	"\x06INDEX4\x10\x06B\x85\x02\n" +
	"\"com.google.ads.googleads.v20.enumsB+ListingGroupFilterCustomAttributeIndexProtoP\x01ZCgoogle.golang.org/genproto/googleapis/ads/googleads/v20/enums;enums\xa2\x02\x03GAA\xaa\x02\x1eGoogle.Ads.GoogleAds.V20.Enums\xca\x02\x1eGoogle\\Ads\\GoogleAds\\V20\\Enums\xea\x02\"Google::Ads::GoogleAds::V20::Enumsb\x06proto3"

var (
	file_google_ads_googleads_v20_enums_listing_group_filter_custom_attribute_index_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v20_enums_listing_group_filter_custom_attribute_index_proto_rawDescData []byte
)

func file_google_ads_googleads_v20_enums_listing_group_filter_custom_attribute_index_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v20_enums_listing_group_filter_custom_attribute_index_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v20_enums_listing_group_filter_custom_attribute_index_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_enums_listing_group_filter_custom_attribute_index_proto_rawDesc), len(file_google_ads_googleads_v20_enums_listing_group_filter_custom_attribute_index_proto_rawDesc)))
	})
	return file_google_ads_googleads_v20_enums_listing_group_filter_custom_attribute_index_proto_rawDescData
}

var file_google_ads_googleads_v20_enums_listing_group_filter_custom_attribute_index_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v20_enums_listing_group_filter_custom_attribute_index_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v20_enums_listing_group_filter_custom_attribute_index_proto_goTypes = []any{
	(ListingGroupFilterCustomAttributeIndexEnum_ListingGroupFilterCustomAttributeIndex)(0), // 0: google.ads.googleads.v20.enums.ListingGroupFilterCustomAttributeIndexEnum.ListingGroupFilterCustomAttributeIndex
	(*ListingGroupFilterCustomAttributeIndexEnum)(nil),                                     // 1: google.ads.googleads.v20.enums.ListingGroupFilterCustomAttributeIndexEnum
}
var file_google_ads_googleads_v20_enums_listing_group_filter_custom_attribute_index_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_google_ads_googleads_v20_enums_listing_group_filter_custom_attribute_index_proto_init()
}
func file_google_ads_googleads_v20_enums_listing_group_filter_custom_attribute_index_proto_init() {
	if File_google_ads_googleads_v20_enums_listing_group_filter_custom_attribute_index_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_enums_listing_group_filter_custom_attribute_index_proto_rawDesc), len(file_google_ads_googleads_v20_enums_listing_group_filter_custom_attribute_index_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v20_enums_listing_group_filter_custom_attribute_index_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v20_enums_listing_group_filter_custom_attribute_index_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v20_enums_listing_group_filter_custom_attribute_index_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v20_enums_listing_group_filter_custom_attribute_index_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v20_enums_listing_group_filter_custom_attribute_index_proto = out.File
	file_google_ads_googleads_v20_enums_listing_group_filter_custom_attribute_index_proto_goTypes = nil
	file_google_ads_googleads_v20_enums_listing_group_filter_custom_attribute_index_proto_depIdxs = nil
}

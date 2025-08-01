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
// source: google/ads/googleads/v20/enums/product_channel.proto

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

// Enum describing the locality of a product offer.
type ProductChannelEnum_ProductChannel int32

const (
	// Not specified.
	ProductChannelEnum_UNSPECIFIED ProductChannelEnum_ProductChannel = 0
	// Used for return value only. Represents value unknown in this version.
	ProductChannelEnum_UNKNOWN ProductChannelEnum_ProductChannel = 1
	// The item is sold online.
	ProductChannelEnum_ONLINE ProductChannelEnum_ProductChannel = 2
	// The item is sold in local stores.
	ProductChannelEnum_LOCAL ProductChannelEnum_ProductChannel = 3
)

// Enum value maps for ProductChannelEnum_ProductChannel.
var (
	ProductChannelEnum_ProductChannel_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "ONLINE",
		3: "LOCAL",
	}
	ProductChannelEnum_ProductChannel_value = map[string]int32{
		"UNSPECIFIED": 0,
		"UNKNOWN":     1,
		"ONLINE":      2,
		"LOCAL":       3,
	}
)

func (x ProductChannelEnum_ProductChannel) Enum() *ProductChannelEnum_ProductChannel {
	p := new(ProductChannelEnum_ProductChannel)
	*p = x
	return p
}

func (x ProductChannelEnum_ProductChannel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProductChannelEnum_ProductChannel) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v20_enums_product_channel_proto_enumTypes[0].Descriptor()
}

func (ProductChannelEnum_ProductChannel) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v20_enums_product_channel_proto_enumTypes[0]
}

func (x ProductChannelEnum_ProductChannel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProductChannelEnum_ProductChannel.Descriptor instead.
func (ProductChannelEnum_ProductChannel) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_enums_product_channel_proto_rawDescGZIP(), []int{0, 0}
}

// Locality of a product offer.
type ProductChannelEnum struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProductChannelEnum) Reset() {
	*x = ProductChannelEnum{}
	mi := &file_google_ads_googleads_v20_enums_product_channel_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProductChannelEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductChannelEnum) ProtoMessage() {}

func (x *ProductChannelEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_enums_product_channel_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductChannelEnum.ProtoReflect.Descriptor instead.
func (*ProductChannelEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_enums_product_channel_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v20_enums_product_channel_proto protoreflect.FileDescriptor

const file_google_ads_googleads_v20_enums_product_channel_proto_rawDesc = "" +
	"\n" +
	"4google/ads/googleads/v20/enums/product_channel.proto\x12\x1egoogle.ads.googleads.v20.enums\"[\n" +
	"\x12ProductChannelEnum\"E\n" +
	"\x0eProductChannel\x12\x0f\n" +
	"\vUNSPECIFIED\x10\x00\x12\v\n" +
	"\aUNKNOWN\x10\x01\x12\n" +
	"\n" +
	"\x06ONLINE\x10\x02\x12\t\n" +
	"\x05LOCAL\x10\x03B\xed\x01\n" +
	"\"com.google.ads.googleads.v20.enumsB\x13ProductChannelProtoP\x01ZCgoogle.golang.org/genproto/googleapis/ads/googleads/v20/enums;enums\xa2\x02\x03GAA\xaa\x02\x1eGoogle.Ads.GoogleAds.V20.Enums\xca\x02\x1eGoogle\\Ads\\GoogleAds\\V20\\Enums\xea\x02\"Google::Ads::GoogleAds::V20::Enumsb\x06proto3"

var (
	file_google_ads_googleads_v20_enums_product_channel_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v20_enums_product_channel_proto_rawDescData []byte
)

func file_google_ads_googleads_v20_enums_product_channel_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v20_enums_product_channel_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v20_enums_product_channel_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_enums_product_channel_proto_rawDesc), len(file_google_ads_googleads_v20_enums_product_channel_proto_rawDesc)))
	})
	return file_google_ads_googleads_v20_enums_product_channel_proto_rawDescData
}

var file_google_ads_googleads_v20_enums_product_channel_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v20_enums_product_channel_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v20_enums_product_channel_proto_goTypes = []any{
	(ProductChannelEnum_ProductChannel)(0), // 0: google.ads.googleads.v20.enums.ProductChannelEnum.ProductChannel
	(*ProductChannelEnum)(nil),             // 1: google.ads.googleads.v20.enums.ProductChannelEnum
}
var file_google_ads_googleads_v20_enums_product_channel_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v20_enums_product_channel_proto_init() }
func file_google_ads_googleads_v20_enums_product_channel_proto_init() {
	if File_google_ads_googleads_v20_enums_product_channel_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_enums_product_channel_proto_rawDesc), len(file_google_ads_googleads_v20_enums_product_channel_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v20_enums_product_channel_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v20_enums_product_channel_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v20_enums_product_channel_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v20_enums_product_channel_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v20_enums_product_channel_proto = out.File
	file_google_ads_googleads_v20_enums_product_channel_proto_goTypes = nil
	file_google_ads_googleads_v20_enums_product_channel_proto_depIdxs = nil
}

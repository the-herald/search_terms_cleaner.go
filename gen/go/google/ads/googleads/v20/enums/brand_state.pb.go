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
// source: google/ads/googleads/v20/enums/brand_state.proto

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

// Enumeration of different brand states.
type BrandStateEnum_BrandState int32

const (
	// No value has been specified.
	BrandStateEnum_UNSPECIFIED BrandStateEnum_BrandState = 0
	// Used for return value only. Represents value unknown in this version.
	BrandStateEnum_UNKNOWN BrandStateEnum_BrandState = 1
	// Brand is verified and globally available for selection
	BrandStateEnum_ENABLED BrandStateEnum_BrandState = 2
	// Brand was globally available in past but is no longer a valid brand
	// (based on business criteria)
	BrandStateEnum_DEPRECATED BrandStateEnum_BrandState = 3
	// Brand is unverified and customer scoped, but can be selected by customer
	// (only who requested for same) for targeting
	BrandStateEnum_UNVERIFIED BrandStateEnum_BrandState = 4
	// Was a customer-scoped (unverified) brand, which got approved by business
	// and added to the global list. Its assigned CKG MID should be used instead
	// of this
	BrandStateEnum_APPROVED BrandStateEnum_BrandState = 5
	// Was a customer-scoped (unverified) brand, but the request was canceled by
	// customer and this brand id is no longer valid
	BrandStateEnum_CANCELLED BrandStateEnum_BrandState = 6
	// Was a customer-scoped (unverified) brand, but the request was rejected by
	// internal business team and this brand id is no longer valid
	BrandStateEnum_REJECTED BrandStateEnum_BrandState = 7
)

// Enum value maps for BrandStateEnum_BrandState.
var (
	BrandStateEnum_BrandState_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "ENABLED",
		3: "DEPRECATED",
		4: "UNVERIFIED",
		5: "APPROVED",
		6: "CANCELLED",
		7: "REJECTED",
	}
	BrandStateEnum_BrandState_value = map[string]int32{
		"UNSPECIFIED": 0,
		"UNKNOWN":     1,
		"ENABLED":     2,
		"DEPRECATED":  3,
		"UNVERIFIED":  4,
		"APPROVED":    5,
		"CANCELLED":   6,
		"REJECTED":    7,
	}
)

func (x BrandStateEnum_BrandState) Enum() *BrandStateEnum_BrandState {
	p := new(BrandStateEnum_BrandState)
	*p = x
	return p
}

func (x BrandStateEnum_BrandState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BrandStateEnum_BrandState) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v20_enums_brand_state_proto_enumTypes[0].Descriptor()
}

func (BrandStateEnum_BrandState) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v20_enums_brand_state_proto_enumTypes[0]
}

func (x BrandStateEnum_BrandState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BrandStateEnum_BrandState.Descriptor instead.
func (BrandStateEnum_BrandState) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_enums_brand_state_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible brand states.
type BrandStateEnum struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BrandStateEnum) Reset() {
	*x = BrandStateEnum{}
	mi := &file_google_ads_googleads_v20_enums_brand_state_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BrandStateEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BrandStateEnum) ProtoMessage() {}

func (x *BrandStateEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_enums_brand_state_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BrandStateEnum.ProtoReflect.Descriptor instead.
func (*BrandStateEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_enums_brand_state_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v20_enums_brand_state_proto protoreflect.FileDescriptor

const file_google_ads_googleads_v20_enums_brand_state_proto_rawDesc = "" +
	"\n" +
	"0google/ads/googleads/v20/enums/brand_state.proto\x12\x1egoogle.ads.googleads.v20.enums\"\x95\x01\n" +
	"\x0eBrandStateEnum\"\x82\x01\n" +
	"\n" +
	"BrandState\x12\x0f\n" +
	"\vUNSPECIFIED\x10\x00\x12\v\n" +
	"\aUNKNOWN\x10\x01\x12\v\n" +
	"\aENABLED\x10\x02\x12\x0e\n" +
	"\n" +
	"DEPRECATED\x10\x03\x12\x0e\n" +
	"\n" +
	"UNVERIFIED\x10\x04\x12\f\n" +
	"\bAPPROVED\x10\x05\x12\r\n" +
	"\tCANCELLED\x10\x06\x12\f\n" +
	"\bREJECTED\x10\aB\xe9\x01\n" +
	"\"com.google.ads.googleads.v20.enumsB\x0fBrandStateProtoP\x01ZCgoogle.golang.org/genproto/googleapis/ads/googleads/v20/enums;enums\xa2\x02\x03GAA\xaa\x02\x1eGoogle.Ads.GoogleAds.V20.Enums\xca\x02\x1eGoogle\\Ads\\GoogleAds\\V20\\Enums\xea\x02\"Google::Ads::GoogleAds::V20::Enumsb\x06proto3"

var (
	file_google_ads_googleads_v20_enums_brand_state_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v20_enums_brand_state_proto_rawDescData []byte
)

func file_google_ads_googleads_v20_enums_brand_state_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v20_enums_brand_state_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v20_enums_brand_state_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_enums_brand_state_proto_rawDesc), len(file_google_ads_googleads_v20_enums_brand_state_proto_rawDesc)))
	})
	return file_google_ads_googleads_v20_enums_brand_state_proto_rawDescData
}

var file_google_ads_googleads_v20_enums_brand_state_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v20_enums_brand_state_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v20_enums_brand_state_proto_goTypes = []any{
	(BrandStateEnum_BrandState)(0), // 0: google.ads.googleads.v20.enums.BrandStateEnum.BrandState
	(*BrandStateEnum)(nil),         // 1: google.ads.googleads.v20.enums.BrandStateEnum
}
var file_google_ads_googleads_v20_enums_brand_state_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v20_enums_brand_state_proto_init() }
func file_google_ads_googleads_v20_enums_brand_state_proto_init() {
	if File_google_ads_googleads_v20_enums_brand_state_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_enums_brand_state_proto_rawDesc), len(file_google_ads_googleads_v20_enums_brand_state_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v20_enums_brand_state_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v20_enums_brand_state_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v20_enums_brand_state_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v20_enums_brand_state_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v20_enums_brand_state_proto = out.File
	file_google_ads_googleads_v20_enums_brand_state_proto_goTypes = nil
	file_google_ads_googleads_v20_enums_brand_state_proto_depIdxs = nil
}

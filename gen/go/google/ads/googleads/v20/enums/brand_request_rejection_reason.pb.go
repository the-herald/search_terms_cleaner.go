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
// source: google/ads/googleads/v20/enums/brand_request_rejection_reason.proto

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

// Enumeration of different brand request rejection reasons.
type BrandRequestRejectionReasonEnum_BrandRequestRejectionReason int32

const (
	// No value has been specified.
	BrandRequestRejectionReasonEnum_UNSPECIFIED BrandRequestRejectionReasonEnum_BrandRequestRejectionReason = 0
	// Used for return value only. Represents value unknown in this version.
	BrandRequestRejectionReasonEnum_UNKNOWN BrandRequestRejectionReasonEnum_BrandRequestRejectionReason = 1
	// Brand is already present in the commercial brand set.
	BrandRequestRejectionReasonEnum_EXISTING_BRAND BrandRequestRejectionReasonEnum_BrandRequestRejectionReason = 2
	// Brand is already present in the commercial brand set, but is a variant.
	BrandRequestRejectionReasonEnum_EXISTING_BRAND_VARIANT BrandRequestRejectionReasonEnum_BrandRequestRejectionReason = 3
	// Brand information is not correct (eg: URL and name don't match).
	BrandRequestRejectionReasonEnum_INCORRECT_INFORMATION BrandRequestRejectionReasonEnum_BrandRequestRejectionReason = 4
	// Not a valid brand as per Google policy.
	BrandRequestRejectionReasonEnum_NOT_A_BRAND BrandRequestRejectionReasonEnum_BrandRequestRejectionReason = 5
)

// Enum value maps for BrandRequestRejectionReasonEnum_BrandRequestRejectionReason.
var (
	BrandRequestRejectionReasonEnum_BrandRequestRejectionReason_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "EXISTING_BRAND",
		3: "EXISTING_BRAND_VARIANT",
		4: "INCORRECT_INFORMATION",
		5: "NOT_A_BRAND",
	}
	BrandRequestRejectionReasonEnum_BrandRequestRejectionReason_value = map[string]int32{
		"UNSPECIFIED":            0,
		"UNKNOWN":                1,
		"EXISTING_BRAND":         2,
		"EXISTING_BRAND_VARIANT": 3,
		"INCORRECT_INFORMATION":  4,
		"NOT_A_BRAND":            5,
	}
)

func (x BrandRequestRejectionReasonEnum_BrandRequestRejectionReason) Enum() *BrandRequestRejectionReasonEnum_BrandRequestRejectionReason {
	p := new(BrandRequestRejectionReasonEnum_BrandRequestRejectionReason)
	*p = x
	return p
}

func (x BrandRequestRejectionReasonEnum_BrandRequestRejectionReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BrandRequestRejectionReasonEnum_BrandRequestRejectionReason) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v20_enums_brand_request_rejection_reason_proto_enumTypes[0].Descriptor()
}

func (BrandRequestRejectionReasonEnum_BrandRequestRejectionReason) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v20_enums_brand_request_rejection_reason_proto_enumTypes[0]
}

func (x BrandRequestRejectionReasonEnum_BrandRequestRejectionReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BrandRequestRejectionReasonEnum_BrandRequestRejectionReason.Descriptor instead.
func (BrandRequestRejectionReasonEnum_BrandRequestRejectionReason) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_enums_brand_request_rejection_reason_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing rejection reasons for the customer brand
// requests.
type BrandRequestRejectionReasonEnum struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BrandRequestRejectionReasonEnum) Reset() {
	*x = BrandRequestRejectionReasonEnum{}
	mi := &file_google_ads_googleads_v20_enums_brand_request_rejection_reason_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BrandRequestRejectionReasonEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BrandRequestRejectionReasonEnum) ProtoMessage() {}

func (x *BrandRequestRejectionReasonEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_enums_brand_request_rejection_reason_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BrandRequestRejectionReasonEnum.ProtoReflect.Descriptor instead.
func (*BrandRequestRejectionReasonEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_enums_brand_request_rejection_reason_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v20_enums_brand_request_rejection_reason_proto protoreflect.FileDescriptor

const file_google_ads_googleads_v20_enums_brand_request_rejection_reason_proto_rawDesc = "" +
	"\n" +
	"Cgoogle/ads/googleads/v20/enums/brand_request_rejection_reason.proto\x12\x1egoogle.ads.googleads.v20.enums\"\xbb\x01\n" +
	"\x1fBrandRequestRejectionReasonEnum\"\x97\x01\n" +
	"\x1bBrandRequestRejectionReason\x12\x0f\n" +
	"\vUNSPECIFIED\x10\x00\x12\v\n" +
	"\aUNKNOWN\x10\x01\x12\x12\n" +
	"\x0eEXISTING_BRAND\x10\x02\x12\x1a\n" +
	"\x16EXISTING_BRAND_VARIANT\x10\x03\x12\x19\n" +
	"\x15INCORRECT_INFORMATION\x10\x04\x12\x0f\n" +
	"\vNOT_A_BRAND\x10\x05B\xfa\x01\n" +
	"\"com.google.ads.googleads.v20.enumsB BrandRequestRejectionReasonProtoP\x01ZCgoogle.golang.org/genproto/googleapis/ads/googleads/v20/enums;enums\xa2\x02\x03GAA\xaa\x02\x1eGoogle.Ads.GoogleAds.V20.Enums\xca\x02\x1eGoogle\\Ads\\GoogleAds\\V20\\Enums\xea\x02\"Google::Ads::GoogleAds::V20::Enumsb\x06proto3"

var (
	file_google_ads_googleads_v20_enums_brand_request_rejection_reason_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v20_enums_brand_request_rejection_reason_proto_rawDescData []byte
)

func file_google_ads_googleads_v20_enums_brand_request_rejection_reason_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v20_enums_brand_request_rejection_reason_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v20_enums_brand_request_rejection_reason_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_enums_brand_request_rejection_reason_proto_rawDesc), len(file_google_ads_googleads_v20_enums_brand_request_rejection_reason_proto_rawDesc)))
	})
	return file_google_ads_googleads_v20_enums_brand_request_rejection_reason_proto_rawDescData
}

var file_google_ads_googleads_v20_enums_brand_request_rejection_reason_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v20_enums_brand_request_rejection_reason_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v20_enums_brand_request_rejection_reason_proto_goTypes = []any{
	(BrandRequestRejectionReasonEnum_BrandRequestRejectionReason)(0), // 0: google.ads.googleads.v20.enums.BrandRequestRejectionReasonEnum.BrandRequestRejectionReason
	(*BrandRequestRejectionReasonEnum)(nil),                          // 1: google.ads.googleads.v20.enums.BrandRequestRejectionReasonEnum
}
var file_google_ads_googleads_v20_enums_brand_request_rejection_reason_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v20_enums_brand_request_rejection_reason_proto_init() }
func file_google_ads_googleads_v20_enums_brand_request_rejection_reason_proto_init() {
	if File_google_ads_googleads_v20_enums_brand_request_rejection_reason_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_enums_brand_request_rejection_reason_proto_rawDesc), len(file_google_ads_googleads_v20_enums_brand_request_rejection_reason_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v20_enums_brand_request_rejection_reason_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v20_enums_brand_request_rejection_reason_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v20_enums_brand_request_rejection_reason_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v20_enums_brand_request_rejection_reason_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v20_enums_brand_request_rejection_reason_proto = out.File
	file_google_ads_googleads_v20_enums_brand_request_rejection_reason_proto_goTypes = nil
	file_google_ads_googleads_v20_enums_brand_request_rejection_reason_proto_depIdxs = nil
}

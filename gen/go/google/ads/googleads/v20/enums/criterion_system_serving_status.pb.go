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
// source: google/ads/googleads/v20/enums/criterion_system_serving_status.proto

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

// Enumerates criterion system serving statuses.
type CriterionSystemServingStatusEnum_CriterionSystemServingStatus int32

const (
	// Not specified.
	CriterionSystemServingStatusEnum_UNSPECIFIED CriterionSystemServingStatusEnum_CriterionSystemServingStatus = 0
	// The value is unknown in this version.
	CriterionSystemServingStatusEnum_UNKNOWN CriterionSystemServingStatusEnum_CriterionSystemServingStatus = 1
	// Eligible.
	CriterionSystemServingStatusEnum_ELIGIBLE CriterionSystemServingStatusEnum_CriterionSystemServingStatus = 2
	// Low search volume.
	CriterionSystemServingStatusEnum_RARELY_SERVED CriterionSystemServingStatusEnum_CriterionSystemServingStatus = 3
)

// Enum value maps for CriterionSystemServingStatusEnum_CriterionSystemServingStatus.
var (
	CriterionSystemServingStatusEnum_CriterionSystemServingStatus_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "ELIGIBLE",
		3: "RARELY_SERVED",
	}
	CriterionSystemServingStatusEnum_CriterionSystemServingStatus_value = map[string]int32{
		"UNSPECIFIED":   0,
		"UNKNOWN":       1,
		"ELIGIBLE":      2,
		"RARELY_SERVED": 3,
	}
)

func (x CriterionSystemServingStatusEnum_CriterionSystemServingStatus) Enum() *CriterionSystemServingStatusEnum_CriterionSystemServingStatus {
	p := new(CriterionSystemServingStatusEnum_CriterionSystemServingStatus)
	*p = x
	return p
}

func (x CriterionSystemServingStatusEnum_CriterionSystemServingStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CriterionSystemServingStatusEnum_CriterionSystemServingStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v20_enums_criterion_system_serving_status_proto_enumTypes[0].Descriptor()
}

func (CriterionSystemServingStatusEnum_CriterionSystemServingStatus) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v20_enums_criterion_system_serving_status_proto_enumTypes[0]
}

func (x CriterionSystemServingStatusEnum_CriterionSystemServingStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CriterionSystemServingStatusEnum_CriterionSystemServingStatus.Descriptor instead.
func (CriterionSystemServingStatusEnum_CriterionSystemServingStatus) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_enums_criterion_system_serving_status_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible criterion system serving statuses.
type CriterionSystemServingStatusEnum struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CriterionSystemServingStatusEnum) Reset() {
	*x = CriterionSystemServingStatusEnum{}
	mi := &file_google_ads_googleads_v20_enums_criterion_system_serving_status_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CriterionSystemServingStatusEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CriterionSystemServingStatusEnum) ProtoMessage() {}

func (x *CriterionSystemServingStatusEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_enums_criterion_system_serving_status_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CriterionSystemServingStatusEnum.ProtoReflect.Descriptor instead.
func (*CriterionSystemServingStatusEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_enums_criterion_system_serving_status_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v20_enums_criterion_system_serving_status_proto protoreflect.FileDescriptor

const file_google_ads_googleads_v20_enums_criterion_system_serving_status_proto_rawDesc = "" +
	"\n" +
	"Dgoogle/ads/googleads/v20/enums/criterion_system_serving_status.proto\x12\x1egoogle.ads.googleads.v20.enums\"\x81\x01\n" +
	" CriterionSystemServingStatusEnum\"]\n" +
	"\x1cCriterionSystemServingStatus\x12\x0f\n" +
	"\vUNSPECIFIED\x10\x00\x12\v\n" +
	"\aUNKNOWN\x10\x01\x12\f\n" +
	"\bELIGIBLE\x10\x02\x12\x11\n" +
	"\rRARELY_SERVED\x10\x03B\xfb\x01\n" +
	"\"com.google.ads.googleads.v20.enumsB!CriterionSystemServingStatusProtoP\x01ZCgoogle.golang.org/genproto/googleapis/ads/googleads/v20/enums;enums\xa2\x02\x03GAA\xaa\x02\x1eGoogle.Ads.GoogleAds.V20.Enums\xca\x02\x1eGoogle\\Ads\\GoogleAds\\V20\\Enums\xea\x02\"Google::Ads::GoogleAds::V20::Enumsb\x06proto3"

var (
	file_google_ads_googleads_v20_enums_criterion_system_serving_status_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v20_enums_criterion_system_serving_status_proto_rawDescData []byte
)

func file_google_ads_googleads_v20_enums_criterion_system_serving_status_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v20_enums_criterion_system_serving_status_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v20_enums_criterion_system_serving_status_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_enums_criterion_system_serving_status_proto_rawDesc), len(file_google_ads_googleads_v20_enums_criterion_system_serving_status_proto_rawDesc)))
	})
	return file_google_ads_googleads_v20_enums_criterion_system_serving_status_proto_rawDescData
}

var file_google_ads_googleads_v20_enums_criterion_system_serving_status_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v20_enums_criterion_system_serving_status_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v20_enums_criterion_system_serving_status_proto_goTypes = []any{
	(CriterionSystemServingStatusEnum_CriterionSystemServingStatus)(0), // 0: google.ads.googleads.v20.enums.CriterionSystemServingStatusEnum.CriterionSystemServingStatus
	(*CriterionSystemServingStatusEnum)(nil),                           // 1: google.ads.googleads.v20.enums.CriterionSystemServingStatusEnum
}
var file_google_ads_googleads_v20_enums_criterion_system_serving_status_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v20_enums_criterion_system_serving_status_proto_init() }
func file_google_ads_googleads_v20_enums_criterion_system_serving_status_proto_init() {
	if File_google_ads_googleads_v20_enums_criterion_system_serving_status_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_enums_criterion_system_serving_status_proto_rawDesc), len(file_google_ads_googleads_v20_enums_criterion_system_serving_status_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v20_enums_criterion_system_serving_status_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v20_enums_criterion_system_serving_status_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v20_enums_criterion_system_serving_status_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v20_enums_criterion_system_serving_status_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v20_enums_criterion_system_serving_status_proto = out.File
	file_google_ads_googleads_v20_enums_criterion_system_serving_status_proto_goTypes = nil
	file_google_ads_googleads_v20_enums_criterion_system_serving_status_proto_depIdxs = nil
}

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
// source: google/ads/googleads/v20/enums/campaign_serving_status.proto

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

// Possible serving statuses of a campaign.
type CampaignServingStatusEnum_CampaignServingStatus int32

const (
	// No value has been specified.
	CampaignServingStatusEnum_UNSPECIFIED CampaignServingStatusEnum_CampaignServingStatus = 0
	// The received value is not known in this version.
	//
	// This is a response-only value.
	CampaignServingStatusEnum_UNKNOWN CampaignServingStatusEnum_CampaignServingStatus = 1
	// Serving.
	CampaignServingStatusEnum_SERVING CampaignServingStatusEnum_CampaignServingStatus = 2
	// None.
	CampaignServingStatusEnum_NONE CampaignServingStatusEnum_CampaignServingStatus = 3
	// Ended.
	CampaignServingStatusEnum_ENDED CampaignServingStatusEnum_CampaignServingStatus = 4
	// Pending.
	CampaignServingStatusEnum_PENDING CampaignServingStatusEnum_CampaignServingStatus = 5
	// Suspended.
	CampaignServingStatusEnum_SUSPENDED CampaignServingStatusEnum_CampaignServingStatus = 6
)

// Enum value maps for CampaignServingStatusEnum_CampaignServingStatus.
var (
	CampaignServingStatusEnum_CampaignServingStatus_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "SERVING",
		3: "NONE",
		4: "ENDED",
		5: "PENDING",
		6: "SUSPENDED",
	}
	CampaignServingStatusEnum_CampaignServingStatus_value = map[string]int32{
		"UNSPECIFIED": 0,
		"UNKNOWN":     1,
		"SERVING":     2,
		"NONE":        3,
		"ENDED":       4,
		"PENDING":     5,
		"SUSPENDED":   6,
	}
)

func (x CampaignServingStatusEnum_CampaignServingStatus) Enum() *CampaignServingStatusEnum_CampaignServingStatus {
	p := new(CampaignServingStatusEnum_CampaignServingStatus)
	*p = x
	return p
}

func (x CampaignServingStatusEnum_CampaignServingStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CampaignServingStatusEnum_CampaignServingStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v20_enums_campaign_serving_status_proto_enumTypes[0].Descriptor()
}

func (CampaignServingStatusEnum_CampaignServingStatus) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v20_enums_campaign_serving_status_proto_enumTypes[0]
}

func (x CampaignServingStatusEnum_CampaignServingStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CampaignServingStatusEnum_CampaignServingStatus.Descriptor instead.
func (CampaignServingStatusEnum_CampaignServingStatus) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_enums_campaign_serving_status_proto_rawDescGZIP(), []int{0, 0}
}

// Message describing Campaign serving statuses.
type CampaignServingStatusEnum struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CampaignServingStatusEnum) Reset() {
	*x = CampaignServingStatusEnum{}
	mi := &file_google_ads_googleads_v20_enums_campaign_serving_status_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CampaignServingStatusEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CampaignServingStatusEnum) ProtoMessage() {}

func (x *CampaignServingStatusEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_enums_campaign_serving_status_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CampaignServingStatusEnum.ProtoReflect.Descriptor instead.
func (*CampaignServingStatusEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_enums_campaign_serving_status_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v20_enums_campaign_serving_status_proto protoreflect.FileDescriptor

const file_google_ads_googleads_v20_enums_campaign_serving_status_proto_rawDesc = "" +
	"\n" +
	"<google/ads/googleads/v20/enums/campaign_serving_status.proto\x12\x1egoogle.ads.googleads.v20.enums\"\x90\x01\n" +
	"\x19CampaignServingStatusEnum\"s\n" +
	"\x15CampaignServingStatus\x12\x0f\n" +
	"\vUNSPECIFIED\x10\x00\x12\v\n" +
	"\aUNKNOWN\x10\x01\x12\v\n" +
	"\aSERVING\x10\x02\x12\b\n" +
	"\x04NONE\x10\x03\x12\t\n" +
	"\x05ENDED\x10\x04\x12\v\n" +
	"\aPENDING\x10\x05\x12\r\n" +
	"\tSUSPENDED\x10\x06B\xf4\x01\n" +
	"\"com.google.ads.googleads.v20.enumsB\x1aCampaignServingStatusProtoP\x01ZCgoogle.golang.org/genproto/googleapis/ads/googleads/v20/enums;enums\xa2\x02\x03GAA\xaa\x02\x1eGoogle.Ads.GoogleAds.V20.Enums\xca\x02\x1eGoogle\\Ads\\GoogleAds\\V20\\Enums\xea\x02\"Google::Ads::GoogleAds::V20::Enumsb\x06proto3"

var (
	file_google_ads_googleads_v20_enums_campaign_serving_status_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v20_enums_campaign_serving_status_proto_rawDescData []byte
)

func file_google_ads_googleads_v20_enums_campaign_serving_status_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v20_enums_campaign_serving_status_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v20_enums_campaign_serving_status_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_enums_campaign_serving_status_proto_rawDesc), len(file_google_ads_googleads_v20_enums_campaign_serving_status_proto_rawDesc)))
	})
	return file_google_ads_googleads_v20_enums_campaign_serving_status_proto_rawDescData
}

var file_google_ads_googleads_v20_enums_campaign_serving_status_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v20_enums_campaign_serving_status_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v20_enums_campaign_serving_status_proto_goTypes = []any{
	(CampaignServingStatusEnum_CampaignServingStatus)(0), // 0: google.ads.googleads.v20.enums.CampaignServingStatusEnum.CampaignServingStatus
	(*CampaignServingStatusEnum)(nil),                    // 1: google.ads.googleads.v20.enums.CampaignServingStatusEnum
}
var file_google_ads_googleads_v20_enums_campaign_serving_status_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v20_enums_campaign_serving_status_proto_init() }
func file_google_ads_googleads_v20_enums_campaign_serving_status_proto_init() {
	if File_google_ads_googleads_v20_enums_campaign_serving_status_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_enums_campaign_serving_status_proto_rawDesc), len(file_google_ads_googleads_v20_enums_campaign_serving_status_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v20_enums_campaign_serving_status_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v20_enums_campaign_serving_status_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v20_enums_campaign_serving_status_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v20_enums_campaign_serving_status_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v20_enums_campaign_serving_status_proto = out.File
	file_google_ads_googleads_v20_enums_campaign_serving_status_proto_goTypes = nil
	file_google_ads_googleads_v20_enums_campaign_serving_status_proto_depIdxs = nil
}

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
// source: google/ads/googleads/v20/enums/campaign_draft_status.proto

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

// Possible statuses of a campaign draft.
type CampaignDraftStatusEnum_CampaignDraftStatus int32

const (
	// The status has not been specified.
	CampaignDraftStatusEnum_UNSPECIFIED CampaignDraftStatusEnum_CampaignDraftStatus = 0
	// Used for return value only. Represents value unknown in this version.
	CampaignDraftStatusEnum_UNKNOWN CampaignDraftStatusEnum_CampaignDraftStatus = 1
	// Initial state of the draft, the advertiser can start adding changes with
	// no effect on serving.
	CampaignDraftStatusEnum_PROPOSED CampaignDraftStatusEnum_CampaignDraftStatus = 2
	// The campaign draft is removed.
	CampaignDraftStatusEnum_REMOVED CampaignDraftStatusEnum_CampaignDraftStatus = 3
	// Advertiser requested to promote draft's changes back into the original
	// campaign. Advertiser can poll the long running operation returned by
	// the promote action to see the status of the promotion.
	CampaignDraftStatusEnum_PROMOTING CampaignDraftStatusEnum_CampaignDraftStatus = 5
	// The process to merge changes in the draft back to the original campaign
	// has completed successfully.
	CampaignDraftStatusEnum_PROMOTED CampaignDraftStatusEnum_CampaignDraftStatus = 4
	// The promotion failed after it was partially applied. Promote cannot be
	// attempted again safely, so the issue must be corrected in the original
	// campaign.
	CampaignDraftStatusEnum_PROMOTE_FAILED CampaignDraftStatusEnum_CampaignDraftStatus = 6
)

// Enum value maps for CampaignDraftStatusEnum_CampaignDraftStatus.
var (
	CampaignDraftStatusEnum_CampaignDraftStatus_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "UNKNOWN",
		2: "PROPOSED",
		3: "REMOVED",
		5: "PROMOTING",
		4: "PROMOTED",
		6: "PROMOTE_FAILED",
	}
	CampaignDraftStatusEnum_CampaignDraftStatus_value = map[string]int32{
		"UNSPECIFIED":    0,
		"UNKNOWN":        1,
		"PROPOSED":       2,
		"REMOVED":        3,
		"PROMOTING":      5,
		"PROMOTED":       4,
		"PROMOTE_FAILED": 6,
	}
)

func (x CampaignDraftStatusEnum_CampaignDraftStatus) Enum() *CampaignDraftStatusEnum_CampaignDraftStatus {
	p := new(CampaignDraftStatusEnum_CampaignDraftStatus)
	*p = x
	return p
}

func (x CampaignDraftStatusEnum_CampaignDraftStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CampaignDraftStatusEnum_CampaignDraftStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_google_ads_googleads_v20_enums_campaign_draft_status_proto_enumTypes[0].Descriptor()
}

func (CampaignDraftStatusEnum_CampaignDraftStatus) Type() protoreflect.EnumType {
	return &file_google_ads_googleads_v20_enums_campaign_draft_status_proto_enumTypes[0]
}

func (x CampaignDraftStatusEnum_CampaignDraftStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CampaignDraftStatusEnum_CampaignDraftStatus.Descriptor instead.
func (CampaignDraftStatusEnum_CampaignDraftStatus) EnumDescriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_enums_campaign_draft_status_proto_rawDescGZIP(), []int{0, 0}
}

// Container for enum describing possible statuses of a campaign draft.
type CampaignDraftStatusEnum struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CampaignDraftStatusEnum) Reset() {
	*x = CampaignDraftStatusEnum{}
	mi := &file_google_ads_googleads_v20_enums_campaign_draft_status_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CampaignDraftStatusEnum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CampaignDraftStatusEnum) ProtoMessage() {}

func (x *CampaignDraftStatusEnum) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_enums_campaign_draft_status_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CampaignDraftStatusEnum.ProtoReflect.Descriptor instead.
func (*CampaignDraftStatusEnum) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_enums_campaign_draft_status_proto_rawDescGZIP(), []int{0}
}

var File_google_ads_googleads_v20_enums_campaign_draft_status_proto protoreflect.FileDescriptor

const file_google_ads_googleads_v20_enums_campaign_draft_status_proto_rawDesc = "" +
	"\n" +
	":google/ads/googleads/v20/enums/campaign_draft_status.proto\x12\x1egoogle.ads.googleads.v20.enums\"\x9a\x01\n" +
	"\x17CampaignDraftStatusEnum\"\x7f\n" +
	"\x13CampaignDraftStatus\x12\x0f\n" +
	"\vUNSPECIFIED\x10\x00\x12\v\n" +
	"\aUNKNOWN\x10\x01\x12\f\n" +
	"\bPROPOSED\x10\x02\x12\v\n" +
	"\aREMOVED\x10\x03\x12\r\n" +
	"\tPROMOTING\x10\x05\x12\f\n" +
	"\bPROMOTED\x10\x04\x12\x12\n" +
	"\x0ePROMOTE_FAILED\x10\x06B\xf2\x01\n" +
	"\"com.google.ads.googleads.v20.enumsB\x18CampaignDraftStatusProtoP\x01ZCgoogle.golang.org/genproto/googleapis/ads/googleads/v20/enums;enums\xa2\x02\x03GAA\xaa\x02\x1eGoogle.Ads.GoogleAds.V20.Enums\xca\x02\x1eGoogle\\Ads\\GoogleAds\\V20\\Enums\xea\x02\"Google::Ads::GoogleAds::V20::Enumsb\x06proto3"

var (
	file_google_ads_googleads_v20_enums_campaign_draft_status_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v20_enums_campaign_draft_status_proto_rawDescData []byte
)

func file_google_ads_googleads_v20_enums_campaign_draft_status_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v20_enums_campaign_draft_status_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v20_enums_campaign_draft_status_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_enums_campaign_draft_status_proto_rawDesc), len(file_google_ads_googleads_v20_enums_campaign_draft_status_proto_rawDesc)))
	})
	return file_google_ads_googleads_v20_enums_campaign_draft_status_proto_rawDescData
}

var file_google_ads_googleads_v20_enums_campaign_draft_status_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_ads_googleads_v20_enums_campaign_draft_status_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v20_enums_campaign_draft_status_proto_goTypes = []any{
	(CampaignDraftStatusEnum_CampaignDraftStatus)(0), // 0: google.ads.googleads.v20.enums.CampaignDraftStatusEnum.CampaignDraftStatus
	(*CampaignDraftStatusEnum)(nil),                  // 1: google.ads.googleads.v20.enums.CampaignDraftStatusEnum
}
var file_google_ads_googleads_v20_enums_campaign_draft_status_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v20_enums_campaign_draft_status_proto_init() }
func file_google_ads_googleads_v20_enums_campaign_draft_status_proto_init() {
	if File_google_ads_googleads_v20_enums_campaign_draft_status_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_enums_campaign_draft_status_proto_rawDesc), len(file_google_ads_googleads_v20_enums_campaign_draft_status_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v20_enums_campaign_draft_status_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v20_enums_campaign_draft_status_proto_depIdxs,
		EnumInfos:         file_google_ads_googleads_v20_enums_campaign_draft_status_proto_enumTypes,
		MessageInfos:      file_google_ads_googleads_v20_enums_campaign_draft_status_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v20_enums_campaign_draft_status_proto = out.File
	file_google_ads_googleads_v20_enums_campaign_draft_status_proto_goTypes = nil
	file_google_ads_googleads_v20_enums_campaign_draft_status_proto_depIdxs = nil
}

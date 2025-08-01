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
// source: google/ads/googleads/v20/resources/campaign_shared_set.proto

package resources

import (
	enums "google.golang.org/genproto/googleapis/ads/googleads/v20/enums"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// CampaignSharedSets are used for managing the shared sets associated with a
// campaign.
type CampaignSharedSet struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Immutable. The resource name of the campaign shared set.
	// Campaign shared set resource names have the form:
	//
	// `customers/{customer_id}/campaignSharedSets/{campaign_id}~{shared_set_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Immutable. The campaign to which the campaign shared set belongs.
	Campaign *string `protobuf:"bytes,5,opt,name=campaign,proto3,oneof" json:"campaign,omitempty"`
	// Immutable. The shared set associated with the campaign. This may be a
	// negative keyword shared set of another customer. This customer should be a
	// manager of the other customer, otherwise the campaign shared set will exist
	// but have no serving effect. Only negative keyword shared sets can be
	// associated with Shopping campaigns. Only negative placement shared sets can
	// be associated with Display mobile app campaigns.
	SharedSet *string `protobuf:"bytes,6,opt,name=shared_set,json=sharedSet,proto3,oneof" json:"shared_set,omitempty"`
	// Output only. The status of this campaign shared set. Read only.
	Status        enums.CampaignSharedSetStatusEnum_CampaignSharedSetStatus `protobuf:"varint,2,opt,name=status,proto3,enum=google.ads.googleads.v20.enums.CampaignSharedSetStatusEnum_CampaignSharedSetStatus" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CampaignSharedSet) Reset() {
	*x = CampaignSharedSet{}
	mi := &file_google_ads_googleads_v20_resources_campaign_shared_set_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CampaignSharedSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CampaignSharedSet) ProtoMessage() {}

func (x *CampaignSharedSet) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_resources_campaign_shared_set_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CampaignSharedSet.ProtoReflect.Descriptor instead.
func (*CampaignSharedSet) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_resources_campaign_shared_set_proto_rawDescGZIP(), []int{0}
}

func (x *CampaignSharedSet) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *CampaignSharedSet) GetCampaign() string {
	if x != nil && x.Campaign != nil {
		return *x.Campaign
	}
	return ""
}

func (x *CampaignSharedSet) GetSharedSet() string {
	if x != nil && x.SharedSet != nil {
		return *x.SharedSet
	}
	return ""
}

func (x *CampaignSharedSet) GetStatus() enums.CampaignSharedSetStatusEnum_CampaignSharedSetStatus {
	if x != nil {
		return x.Status
	}
	return enums.CampaignSharedSetStatusEnum_CampaignSharedSetStatus(0)
}

var File_google_ads_googleads_v20_resources_campaign_shared_set_proto protoreflect.FileDescriptor

const file_google_ads_googleads_v20_resources_campaign_shared_set_proto_rawDesc = "" +
	"\n" +
	"<google/ads/googleads/v20/resources/campaign_shared_set.proto\x12\"google.ads.googleads.v20.resources\x1a?google/ads/googleads/v20/enums/campaign_shared_set_status.proto\x1a\x1fgoogle/api/field_behavior.proto\x1a\x19google/api/resource.proto\"\x91\x04\n" +
	"\x11CampaignSharedSet\x12W\n" +
	"\rresource_name\x18\x01 \x01(\tB2\xe0A\x05\xfaA,\n" +
	"*googleads.googleapis.com/CampaignSharedSetR\fresourceName\x12J\n" +
	"\bcampaign\x18\x05 \x01(\tB)\xe0A\x05\xfaA#\n" +
	"!googleads.googleapis.com/CampaignH\x00R\bcampaign\x88\x01\x01\x12N\n" +
	"\n" +
	"shared_set\x18\x06 \x01(\tB*\xe0A\x05\xfaA$\n" +
	"\"googleads.googleapis.com/SharedSetH\x01R\tsharedSet\x88\x01\x01\x12p\n" +
	"\x06status\x18\x02 \x01(\x0e2S.google.ads.googleads.v20.enums.CampaignSharedSetStatusEnum.CampaignSharedSetStatusB\x03\xe0A\x03R\x06status:y\xeaAv\n" +
	"*googleads.googleapis.com/CampaignSharedSet\x12Hcustomers/{customer_id}/campaignSharedSets/{campaign_id}~{shared_set_id}B\v\n" +
	"\t_campaignB\r\n" +
	"\v_shared_setB\x88\x02\n" +
	"&com.google.ads.googleads.v20.resourcesB\x16CampaignSharedSetProtoP\x01ZKgoogle.golang.org/genproto/googleapis/ads/googleads/v20/resources;resources\xa2\x02\x03GAA\xaa\x02\"Google.Ads.GoogleAds.V20.Resources\xca\x02\"Google\\Ads\\GoogleAds\\V20\\Resources\xea\x02&Google::Ads::GoogleAds::V20::Resourcesb\x06proto3"

var (
	file_google_ads_googleads_v20_resources_campaign_shared_set_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v20_resources_campaign_shared_set_proto_rawDescData []byte
)

func file_google_ads_googleads_v20_resources_campaign_shared_set_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v20_resources_campaign_shared_set_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v20_resources_campaign_shared_set_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_resources_campaign_shared_set_proto_rawDesc), len(file_google_ads_googleads_v20_resources_campaign_shared_set_proto_rawDesc)))
	})
	return file_google_ads_googleads_v20_resources_campaign_shared_set_proto_rawDescData
}

var file_google_ads_googleads_v20_resources_campaign_shared_set_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v20_resources_campaign_shared_set_proto_goTypes = []any{
	(*CampaignSharedSet)(nil),                                      // 0: google.ads.googleads.v20.resources.CampaignSharedSet
	(enums.CampaignSharedSetStatusEnum_CampaignSharedSetStatus)(0), // 1: google.ads.googleads.v20.enums.CampaignSharedSetStatusEnum.CampaignSharedSetStatus
}
var file_google_ads_googleads_v20_resources_campaign_shared_set_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v20.resources.CampaignSharedSet.status:type_name -> google.ads.googleads.v20.enums.CampaignSharedSetStatusEnum.CampaignSharedSetStatus
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v20_resources_campaign_shared_set_proto_init() }
func file_google_ads_googleads_v20_resources_campaign_shared_set_proto_init() {
	if File_google_ads_googleads_v20_resources_campaign_shared_set_proto != nil {
		return
	}
	file_google_ads_googleads_v20_resources_campaign_shared_set_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_resources_campaign_shared_set_proto_rawDesc), len(file_google_ads_googleads_v20_resources_campaign_shared_set_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v20_resources_campaign_shared_set_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v20_resources_campaign_shared_set_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v20_resources_campaign_shared_set_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v20_resources_campaign_shared_set_proto = out.File
	file_google_ads_googleads_v20_resources_campaign_shared_set_proto_goTypes = nil
	file_google_ads_googleads_v20_resources_campaign_shared_set_proto_depIdxs = nil
}

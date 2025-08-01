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
// source: google/ads/googleads/v20/resources/campaign_group.proto

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

// A campaign group.
type CampaignGroup struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Immutable. The resource name of the campaign group.
	// Campaign group resource names have the form:
	//
	// `customers/{customer_id}/campaignGroups/{campaign_group_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The ID of the campaign group.
	Id int64 `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	// The name of the campaign group.
	//
	// This field is required and should not be empty when creating new campaign
	// groups.
	//
	// It must not contain any null (code point 0x0), NL line feed
	// (code point 0xA) or carriage return (code point 0xD) characters.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// The status of the campaign group.
	//
	// When a new campaign group is added, the status defaults to ENABLED.
	Status        enums.CampaignGroupStatusEnum_CampaignGroupStatus `protobuf:"varint,5,opt,name=status,proto3,enum=google.ads.googleads.v20.enums.CampaignGroupStatusEnum_CampaignGroupStatus" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CampaignGroup) Reset() {
	*x = CampaignGroup{}
	mi := &file_google_ads_googleads_v20_resources_campaign_group_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CampaignGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CampaignGroup) ProtoMessage() {}

func (x *CampaignGroup) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_resources_campaign_group_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CampaignGroup.ProtoReflect.Descriptor instead.
func (*CampaignGroup) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_resources_campaign_group_proto_rawDescGZIP(), []int{0}
}

func (x *CampaignGroup) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *CampaignGroup) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CampaignGroup) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CampaignGroup) GetStatus() enums.CampaignGroupStatusEnum_CampaignGroupStatus {
	if x != nil {
		return x.Status
	}
	return enums.CampaignGroupStatusEnum_CampaignGroupStatus(0)
}

var File_google_ads_googleads_v20_resources_campaign_group_proto protoreflect.FileDescriptor

const file_google_ads_googleads_v20_resources_campaign_group_proto_rawDesc = "" +
	"\n" +
	"7google/ads/googleads/v20/resources/campaign_group.proto\x12\"google.ads.googleads.v20.resources\x1a:google/ads/googleads/v20/enums/campaign_group_status.proto\x1a\x1fgoogle/api/field_behavior.proto\x1a\x19google/api/resource.proto\"\xdb\x02\n" +
	"\rCampaignGroup\x12S\n" +
	"\rresource_name\x18\x01 \x01(\tB.\xe0A\x05\xfaA(\n" +
	"&googleads.googleapis.com/CampaignGroupR\fresourceName\x12\x13\n" +
	"\x02id\x18\x03 \x01(\x03B\x03\xe0A\x03R\x02id\x12\x12\n" +
	"\x04name\x18\x04 \x01(\tR\x04name\x12c\n" +
	"\x06status\x18\x05 \x01(\x0e2K.google.ads.googleads.v20.enums.CampaignGroupStatusEnum.CampaignGroupStatusR\x06status:g\xeaAd\n" +
	"&googleads.googleapis.com/CampaignGroup\x12:customers/{customer_id}/campaignGroups/{campaign_group_id}B\x84\x02\n" +
	"&com.google.ads.googleads.v20.resourcesB\x12CampaignGroupProtoP\x01ZKgoogle.golang.org/genproto/googleapis/ads/googleads/v20/resources;resources\xa2\x02\x03GAA\xaa\x02\"Google.Ads.GoogleAds.V20.Resources\xca\x02\"Google\\Ads\\GoogleAds\\V20\\Resources\xea\x02&Google::Ads::GoogleAds::V20::Resourcesb\x06proto3"

var (
	file_google_ads_googleads_v20_resources_campaign_group_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v20_resources_campaign_group_proto_rawDescData []byte
)

func file_google_ads_googleads_v20_resources_campaign_group_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v20_resources_campaign_group_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v20_resources_campaign_group_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_resources_campaign_group_proto_rawDesc), len(file_google_ads_googleads_v20_resources_campaign_group_proto_rawDesc)))
	})
	return file_google_ads_googleads_v20_resources_campaign_group_proto_rawDescData
}

var file_google_ads_googleads_v20_resources_campaign_group_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v20_resources_campaign_group_proto_goTypes = []any{
	(*CampaignGroup)(nil), // 0: google.ads.googleads.v20.resources.CampaignGroup
	(enums.CampaignGroupStatusEnum_CampaignGroupStatus)(0), // 1: google.ads.googleads.v20.enums.CampaignGroupStatusEnum.CampaignGroupStatus
}
var file_google_ads_googleads_v20_resources_campaign_group_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v20.resources.CampaignGroup.status:type_name -> google.ads.googleads.v20.enums.CampaignGroupStatusEnum.CampaignGroupStatus
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v20_resources_campaign_group_proto_init() }
func file_google_ads_googleads_v20_resources_campaign_group_proto_init() {
	if File_google_ads_googleads_v20_resources_campaign_group_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_resources_campaign_group_proto_rawDesc), len(file_google_ads_googleads_v20_resources_campaign_group_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v20_resources_campaign_group_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v20_resources_campaign_group_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v20_resources_campaign_group_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v20_resources_campaign_group_proto = out.File
	file_google_ads_googleads_v20_resources_campaign_group_proto_goTypes = nil
	file_google_ads_googleads_v20_resources_campaign_group_proto_depIdxs = nil
}

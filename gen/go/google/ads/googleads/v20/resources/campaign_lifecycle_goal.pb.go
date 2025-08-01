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
// source: google/ads/googleads/v20/resources/campaign_lifecycle_goal.proto

package resources

import (
	common "google.golang.org/genproto/googleapis/ads/googleads/v20/common"
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

// Campaign level customer lifecycle goal settings.
type CampaignLifecycleGoal struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Immutable. The resource name of the customer lifecycle goal of a campaign.
	//
	// `customers/{customer_id}/campaignLifecycleGoal/{campaign_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The campaign where the goal is attached.
	Campaign string `protobuf:"bytes,2,opt,name=campaign,proto3" json:"campaign,omitempty"`
	// Output only. The customer acquisition goal settings for the campaign. The
	// customer acquisition goal is described in this article:
	// https://support.google.com/google-ads/answer/12080169
	CustomerAcquisitionGoalSettings *CustomerAcquisitionGoalSettings `protobuf:"bytes,3,opt,name=customer_acquisition_goal_settings,json=customerAcquisitionGoalSettings,proto3" json:"customer_acquisition_goal_settings,omitempty"`
	unknownFields                   protoimpl.UnknownFields
	sizeCache                       protoimpl.SizeCache
}

func (x *CampaignLifecycleGoal) Reset() {
	*x = CampaignLifecycleGoal{}
	mi := &file_google_ads_googleads_v20_resources_campaign_lifecycle_goal_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CampaignLifecycleGoal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CampaignLifecycleGoal) ProtoMessage() {}

func (x *CampaignLifecycleGoal) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_resources_campaign_lifecycle_goal_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CampaignLifecycleGoal.ProtoReflect.Descriptor instead.
func (*CampaignLifecycleGoal) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_resources_campaign_lifecycle_goal_proto_rawDescGZIP(), []int{0}
}

func (x *CampaignLifecycleGoal) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *CampaignLifecycleGoal) GetCampaign() string {
	if x != nil {
		return x.Campaign
	}
	return ""
}

func (x *CampaignLifecycleGoal) GetCustomerAcquisitionGoalSettings() *CustomerAcquisitionGoalSettings {
	if x != nil {
		return x.CustomerAcquisitionGoalSettings
	}
	return nil
}

// The customer acquisition goal settings for the campaign.
type CustomerAcquisitionGoalSettings struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Output only. Customer acquisition optimization mode of this campaign.
	OptimizationMode enums.CustomerAcquisitionOptimizationModeEnum_CustomerAcquisitionOptimizationMode `protobuf:"varint,1,opt,name=optimization_mode,json=optimizationMode,proto3,enum=google.ads.googleads.v20.enums.CustomerAcquisitionOptimizationModeEnum_CustomerAcquisitionOptimizationMode" json:"optimization_mode,omitempty"`
	// Output only. Campaign specific values for the customer acquisition goal.
	ValueSettings *common.LifecycleGoalValueSettings `protobuf:"bytes,2,opt,name=value_settings,json=valueSettings,proto3" json:"value_settings,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CustomerAcquisitionGoalSettings) Reset() {
	*x = CustomerAcquisitionGoalSettings{}
	mi := &file_google_ads_googleads_v20_resources_campaign_lifecycle_goal_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CustomerAcquisitionGoalSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerAcquisitionGoalSettings) ProtoMessage() {}

func (x *CustomerAcquisitionGoalSettings) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_resources_campaign_lifecycle_goal_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerAcquisitionGoalSettings.ProtoReflect.Descriptor instead.
func (*CustomerAcquisitionGoalSettings) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_resources_campaign_lifecycle_goal_proto_rawDescGZIP(), []int{1}
}

func (x *CustomerAcquisitionGoalSettings) GetOptimizationMode() enums.CustomerAcquisitionOptimizationModeEnum_CustomerAcquisitionOptimizationMode {
	if x != nil {
		return x.OptimizationMode
	}
	return enums.CustomerAcquisitionOptimizationModeEnum_CustomerAcquisitionOptimizationMode(0)
}

func (x *CustomerAcquisitionGoalSettings) GetValueSettings() *common.LifecycleGoalValueSettings {
	if x != nil {
		return x.ValueSettings
	}
	return nil
}

var File_google_ads_googleads_v20_resources_campaign_lifecycle_goal_proto protoreflect.FileDescriptor

const file_google_ads_googleads_v20_resources_campaign_lifecycle_goal_proto_rawDesc = "" +
	"\n" +
	"@google/ads/googleads/v20/resources/campaign_lifecycle_goal.proto\x12\"google.ads.googleads.v20.resources\x1a5google/ads/googleads/v20/common/lifecycle_goals.proto\x1aKgoogle/ads/googleads/v20/enums/customer_acquisition_optimization_mode.proto\x1a\x1fgoogle/api/field_behavior.proto\x1a\x19google/api/resource.proto\"\xc6\x03\n" +
	"\x15CampaignLifecycleGoal\x12[\n" +
	"\rresource_name\x18\x01 \x01(\tB6\xe0A\x05\xfaA0\n" +
	".googleads.googleapis.com/CampaignLifecycleGoalR\fresourceName\x12E\n" +
	"\bcampaign\x18\x02 \x01(\tB)\xe0A\x03\xfaA#\n" +
	"!googleads.googleapis.com/CampaignR\bcampaign\x12\x95\x01\n" +
	"\"customer_acquisition_goal_settings\x18\x03 \x01(\v2C.google.ads.googleads.v20.resources.CustomerAcquisitionGoalSettingsB\x03\xe0A\x03R\x1fcustomerAcquisitionGoalSettings:q\xeaAn\n" +
	".googleads.googleapis.com/CampaignLifecycleGoal\x12<customers/{customer_id}/campaignLifecycleGoals/{campaign_id}\"\xaa\x02\n" +
	"\x1fCustomerAcquisitionGoalSettings\x12\x9d\x01\n" +
	"\x11optimization_mode\x18\x01 \x01(\x0e2k.google.ads.googleads.v20.enums.CustomerAcquisitionOptimizationModeEnum.CustomerAcquisitionOptimizationModeB\x03\xe0A\x03R\x10optimizationMode\x12g\n" +
	"\x0evalue_settings\x18\x02 \x01(\v2;.google.ads.googleads.v20.common.LifecycleGoalValueSettingsB\x03\xe0A\x03R\rvalueSettingsB\x8c\x02\n" +
	"&com.google.ads.googleads.v20.resourcesB\x1aCampaignLifecycleGoalProtoP\x01ZKgoogle.golang.org/genproto/googleapis/ads/googleads/v20/resources;resources\xa2\x02\x03GAA\xaa\x02\"Google.Ads.GoogleAds.V20.Resources\xca\x02\"Google\\Ads\\GoogleAds\\V20\\Resources\xea\x02&Google::Ads::GoogleAds::V20::Resourcesb\x06proto3"

var (
	file_google_ads_googleads_v20_resources_campaign_lifecycle_goal_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v20_resources_campaign_lifecycle_goal_proto_rawDescData []byte
)

func file_google_ads_googleads_v20_resources_campaign_lifecycle_goal_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v20_resources_campaign_lifecycle_goal_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v20_resources_campaign_lifecycle_goal_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_resources_campaign_lifecycle_goal_proto_rawDesc), len(file_google_ads_googleads_v20_resources_campaign_lifecycle_goal_proto_rawDesc)))
	})
	return file_google_ads_googleads_v20_resources_campaign_lifecycle_goal_proto_rawDescData
}

var file_google_ads_googleads_v20_resources_campaign_lifecycle_goal_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_ads_googleads_v20_resources_campaign_lifecycle_goal_proto_goTypes = []any{
	(*CampaignLifecycleGoal)(nil),           // 0: google.ads.googleads.v20.resources.CampaignLifecycleGoal
	(*CustomerAcquisitionGoalSettings)(nil), // 1: google.ads.googleads.v20.resources.CustomerAcquisitionGoalSettings
	(enums.CustomerAcquisitionOptimizationModeEnum_CustomerAcquisitionOptimizationMode)(0), // 2: google.ads.googleads.v20.enums.CustomerAcquisitionOptimizationModeEnum.CustomerAcquisitionOptimizationMode
	(*common.LifecycleGoalValueSettings)(nil),                                              // 3: google.ads.googleads.v20.common.LifecycleGoalValueSettings
}
var file_google_ads_googleads_v20_resources_campaign_lifecycle_goal_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v20.resources.CampaignLifecycleGoal.customer_acquisition_goal_settings:type_name -> google.ads.googleads.v20.resources.CustomerAcquisitionGoalSettings
	2, // 1: google.ads.googleads.v20.resources.CustomerAcquisitionGoalSettings.optimization_mode:type_name -> google.ads.googleads.v20.enums.CustomerAcquisitionOptimizationModeEnum.CustomerAcquisitionOptimizationMode
	3, // 2: google.ads.googleads.v20.resources.CustomerAcquisitionGoalSettings.value_settings:type_name -> google.ads.googleads.v20.common.LifecycleGoalValueSettings
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v20_resources_campaign_lifecycle_goal_proto_init() }
func file_google_ads_googleads_v20_resources_campaign_lifecycle_goal_proto_init() {
	if File_google_ads_googleads_v20_resources_campaign_lifecycle_goal_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_resources_campaign_lifecycle_goal_proto_rawDesc), len(file_google_ads_googleads_v20_resources_campaign_lifecycle_goal_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v20_resources_campaign_lifecycle_goal_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v20_resources_campaign_lifecycle_goal_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v20_resources_campaign_lifecycle_goal_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v20_resources_campaign_lifecycle_goal_proto = out.File
	file_google_ads_googleads_v20_resources_campaign_lifecycle_goal_proto_goTypes = nil
	file_google_ads_googleads_v20_resources_campaign_lifecycle_goal_proto_depIdxs = nil
}

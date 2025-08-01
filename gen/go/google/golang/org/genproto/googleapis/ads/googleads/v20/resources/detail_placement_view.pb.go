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
// source: google/ads/googleads/v20/resources/detail_placement_view.proto

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

// A view with metrics aggregated by ad group and URL or YouTube video.
type DetailPlacementView struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Output only. The resource name of the detail placement view.
	// Detail placement view resource names have the form:
	//
	// `customers/{customer_id}/detailPlacementViews/{ad_group_id}~{base64_placement}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// Output only. The automatic placement string at detail level, e. g. website
	// URL, mobile application ID, or a YouTube video ID.
	Placement *string `protobuf:"bytes,7,opt,name=placement,proto3,oneof" json:"placement,omitempty"`
	// Output only. The display name is URL name for websites, YouTube video name
	// for YouTube videos, and translated mobile app name for mobile apps.
	DisplayName *string `protobuf:"bytes,8,opt,name=display_name,json=displayName,proto3,oneof" json:"display_name,omitempty"`
	// Output only. URL of the group placement, for example, domain, link to the
	// mobile application in app store, or a YouTube channel URL.
	GroupPlacementTargetUrl *string `protobuf:"bytes,9,opt,name=group_placement_target_url,json=groupPlacementTargetUrl,proto3,oneof" json:"group_placement_target_url,omitempty"`
	// Output only. URL of the placement, for example, website, link to the mobile
	// application in app store, or a YouTube video URL.
	TargetUrl *string `protobuf:"bytes,10,opt,name=target_url,json=targetUrl,proto3,oneof" json:"target_url,omitempty"`
	// Output only. Type of the placement, for example, Website, YouTube Video,
	// and Mobile Application.
	PlacementType enums.PlacementTypeEnum_PlacementType `protobuf:"varint,6,opt,name=placement_type,json=placementType,proto3,enum=google.ads.googleads.v20.enums.PlacementTypeEnum_PlacementType" json:"placement_type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DetailPlacementView) Reset() {
	*x = DetailPlacementView{}
	mi := &file_google_ads_googleads_v20_resources_detail_placement_view_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DetailPlacementView) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DetailPlacementView) ProtoMessage() {}

func (x *DetailPlacementView) ProtoReflect() protoreflect.Message {
	mi := &file_google_ads_googleads_v20_resources_detail_placement_view_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DetailPlacementView.ProtoReflect.Descriptor instead.
func (*DetailPlacementView) Descriptor() ([]byte, []int) {
	return file_google_ads_googleads_v20_resources_detail_placement_view_proto_rawDescGZIP(), []int{0}
}

func (x *DetailPlacementView) GetResourceName() string {
	if x != nil {
		return x.ResourceName
	}
	return ""
}

func (x *DetailPlacementView) GetPlacement() string {
	if x != nil && x.Placement != nil {
		return *x.Placement
	}
	return ""
}

func (x *DetailPlacementView) GetDisplayName() string {
	if x != nil && x.DisplayName != nil {
		return *x.DisplayName
	}
	return ""
}

func (x *DetailPlacementView) GetGroupPlacementTargetUrl() string {
	if x != nil && x.GroupPlacementTargetUrl != nil {
		return *x.GroupPlacementTargetUrl
	}
	return ""
}

func (x *DetailPlacementView) GetTargetUrl() string {
	if x != nil && x.TargetUrl != nil {
		return *x.TargetUrl
	}
	return ""
}

func (x *DetailPlacementView) GetPlacementType() enums.PlacementTypeEnum_PlacementType {
	if x != nil {
		return x.PlacementType
	}
	return enums.PlacementTypeEnum_PlacementType(0)
}

var File_google_ads_googleads_v20_resources_detail_placement_view_proto protoreflect.FileDescriptor

const file_google_ads_googleads_v20_resources_detail_placement_view_proto_rawDesc = "" +
	"\n" +
	">google/ads/googleads/v20/resources/detail_placement_view.proto\x12\"google.ads.googleads.v20.resources\x1a3google/ads/googleads/v20/enums/placement_type.proto\x1a\x1fgoogle/api/field_behavior.proto\x1a\x19google/api/resource.proto\"\xf2\x04\n" +
	"\x13DetailPlacementView\x12Y\n" +
	"\rresource_name\x18\x01 \x01(\tB4\xe0A\x03\xfaA.\n" +
	",googleads.googleapis.com/DetailPlacementViewR\fresourceName\x12&\n" +
	"\tplacement\x18\a \x01(\tB\x03\xe0A\x03H\x00R\tplacement\x88\x01\x01\x12+\n" +
	"\fdisplay_name\x18\b \x01(\tB\x03\xe0A\x03H\x01R\vdisplayName\x88\x01\x01\x12E\n" +
	"\x1agroup_placement_target_url\x18\t \x01(\tB\x03\xe0A\x03H\x02R\x17groupPlacementTargetUrl\x88\x01\x01\x12'\n" +
	"\n" +
	"target_url\x18\n" +
	" \x01(\tB\x03\xe0A\x03H\x03R\ttargetUrl\x88\x01\x01\x12k\n" +
	"\x0eplacement_type\x18\x06 \x01(\x0e2?.google.ads.googleads.v20.enums.PlacementTypeEnum.PlacementTypeB\x03\xe0A\x03R\rplacementType:\x80\x01\xeaA}\n" +
	",googleads.googleapis.com/DetailPlacementView\x12Mcustomers/{customer_id}/detailPlacementViews/{ad_group_id}~{base64_placement}B\f\n" +
	"\n" +
	"_placementB\x0f\n" +
	"\r_display_nameB\x1d\n" +
	"\x1b_group_placement_target_urlB\r\n" +
	"\v_target_urlB\x8a\x02\n" +
	"&com.google.ads.googleads.v20.resourcesB\x18DetailPlacementViewProtoP\x01ZKgoogle.golang.org/genproto/googleapis/ads/googleads/v20/resources;resources\xa2\x02\x03GAA\xaa\x02\"Google.Ads.GoogleAds.V20.Resources\xca\x02\"Google\\Ads\\GoogleAds\\V20\\Resources\xea\x02&Google::Ads::GoogleAds::V20::Resourcesb\x06proto3"

var (
	file_google_ads_googleads_v20_resources_detail_placement_view_proto_rawDescOnce sync.Once
	file_google_ads_googleads_v20_resources_detail_placement_view_proto_rawDescData []byte
)

func file_google_ads_googleads_v20_resources_detail_placement_view_proto_rawDescGZIP() []byte {
	file_google_ads_googleads_v20_resources_detail_placement_view_proto_rawDescOnce.Do(func() {
		file_google_ads_googleads_v20_resources_detail_placement_view_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_resources_detail_placement_view_proto_rawDesc), len(file_google_ads_googleads_v20_resources_detail_placement_view_proto_rawDesc)))
	})
	return file_google_ads_googleads_v20_resources_detail_placement_view_proto_rawDescData
}

var file_google_ads_googleads_v20_resources_detail_placement_view_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_ads_googleads_v20_resources_detail_placement_view_proto_goTypes = []any{
	(*DetailPlacementView)(nil),                // 0: google.ads.googleads.v20.resources.DetailPlacementView
	(enums.PlacementTypeEnum_PlacementType)(0), // 1: google.ads.googleads.v20.enums.PlacementTypeEnum.PlacementType
}
var file_google_ads_googleads_v20_resources_detail_placement_view_proto_depIdxs = []int32{
	1, // 0: google.ads.googleads.v20.resources.DetailPlacementView.placement_type:type_name -> google.ads.googleads.v20.enums.PlacementTypeEnum.PlacementType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_ads_googleads_v20_resources_detail_placement_view_proto_init() }
func file_google_ads_googleads_v20_resources_detail_placement_view_proto_init() {
	if File_google_ads_googleads_v20_resources_detail_placement_view_proto != nil {
		return
	}
	file_google_ads_googleads_v20_resources_detail_placement_view_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_google_ads_googleads_v20_resources_detail_placement_view_proto_rawDesc), len(file_google_ads_googleads_v20_resources_detail_placement_view_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_ads_googleads_v20_resources_detail_placement_view_proto_goTypes,
		DependencyIndexes: file_google_ads_googleads_v20_resources_detail_placement_view_proto_depIdxs,
		MessageInfos:      file_google_ads_googleads_v20_resources_detail_placement_view_proto_msgTypes,
	}.Build()
	File_google_ads_googleads_v20_resources_detail_placement_view_proto = out.File
	file_google_ads_googleads_v20_resources_detail_placement_view_proto_goTypes = nil
	file_google_ads_googleads_v20_resources_detail_placement_view_proto_depIdxs = nil
}
